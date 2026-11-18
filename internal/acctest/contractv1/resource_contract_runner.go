package contractv1

import (
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var contractLocks sync.Map

func (c ResourceAccTestContractV1) newTestCase(t *testing.T) resource.TestCase {
	return c.TestCase(t)
}

func (c ResourceAccTestContractV1) prerequisites(t *testing.T) {
	t.Helper()
	if os.Getenv("TF_ACC") != "1" {
		t.Skip("set TF_ACC=1")
	}
	if os.Getenv("HARNESS_ACC_ALLOW_DESTRUCTIVE") != "yes" {
		t.Skip("set HARNESS_ACC_ALLOW_DESTRUCTIVE=yes")
	}
	for _, k := range c.GetRequiredEnv() {
		if os.Getenv(k) == "" {
			t.Skipf("missing required environment variable %s", k)
		}
	}
}

func (c ResourceAccTestContractV1) executeContractDefinition(t *testing.T, d ContractDefinition) {
	t.Helper()
	t.Run(string(d.TestName), func(t *testing.T) {
		t.Logf("[contractv1] START %s", d.TestName)
		if !d.Supported {
			t.Logf("[contractv1] SKIP %s: %s", d.TestName, d.UnsupportedReason)
			t.Skip(d.UnsupportedReason)
		}
		if d.PreTest != nil {
			t.Logf("[contractv1] %s: running PreTest", d.TestName)
			if err := d.PreTest(); err != nil {
				t.Fatalf("PreTest: %v", err)
			}
		}
		if d.PostTest != nil {
			t.Cleanup(func() {
				t.Logf("[contractv1] %s: running PostTest", d.TestName)
				if err := d.PostTest(); err != nil {
					t.Errorf("PostTest: %v", err)
				}
			})
		}
		tc := c.newTestCase(t)
		t.Logf("[contractv1] %s: building test steps", d.TestName)
		switch d.TestName {
		case TestCreateResource:
			runCreateResourceTest(t, &tc, d)
		case TestReadResource:
			runReadResourceTest(t, &tc, d)
		case TestUpdateResource:
			runUpdateResourceTest(t, &tc, d)
		case TestDeleteResource:
			runDeleteResourceTest(t, &tc, d)
		case TestCreateUpdateDeleteResource:
			runCreateUpdateDeleteResourceTest(t, &tc, d)
		case TestImportResource:
			runImportResourceTest(t, &tc, d)
		case TestIdempotencyResource:
			runIdempotencyResourceTest(t, &tc, d)
		case TestExternalDeleteResource:
			runExternalDeleteResourceTest(t, &tc, d)
		case TestExternalUpdateResource:
			runExternalUpdateResourceTest(t, &tc, d)
		case TestForceNewResource:
			runForceNewResourceTest(t, &tc, d)
		case TestInvalidConfigurationResource:
			runInvalidConfigurationResourceTest(t, &tc, d)
		case TestAPIErrorResource:
			runAPIErrorResourceTest(t, &tc, d)
		case TestAsyncResource:
			runAsyncResourceTest(t, &tc, d)
		case TestTimeoutResource:
			runTimeoutResourceTest(t, &tc, d)
		case TestImportNonExistentResource:
			runImportNonExistentResourceTest(t, &tc, d)
		}
		t.Logf("[contractv1] %s: running resource.Test with %d steps", d.TestName, len(tc.Steps))
		resource.Test(t, tc)
		t.Logf("[contractv1] END %s", d.TestName)
	})
}

func (c ResourceAccTestContractV1) run(t *testing.T) {
	t.Helper()
	c.prerequisites(t)
	key := c.GetSerializerKey()
	if key == "" {
		t.Fatal("GetSerializerKey returned an empty string")
	}
	v, _ := contractLocks.LoadOrStore(key, &sync.Mutex{})
	mu := v.(*sync.Mutex)
	mu.Lock()
	defer mu.Unlock()
	only := map[TestName]bool{}
	for _, n := range c.DebugSettings.RunOnlyTests {
		only[n] = true
	}
	for _, n := range TestOrder {
		d, _ := c.Definition(n)
		if len(only) > 0 && !only[n] {
			continue
		}
		c.executeContractDefinition(t, d)
	}
}

func runCreateResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig(), Check: d.Verifiers.VerifyCreated},
	}
}

func runReadResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig()},
		{RefreshState: true, Check: d.Verifiers.VerifyRead},
	}
}

func runUpdateResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig()},
		{Config: d.GetUpdatedTerraformConfig(), Check: d.Verifiers.VerifyUpdated},
	}
}

func runDeleteResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.CheckDestroy = d.Verifiers.VerifyDeleted
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig()},
	}
}

func runCreateUpdateDeleteResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.CheckDestroy = d.Verifiers.VerifyDeleted
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig(), Check: d.Verifiers.VerifyCreated},
		{Config: d.GetUpdatedTerraformConfig(), Check: d.Verifiers.VerifyUpdated},
	}
}

func runImportResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	importStep := resource.TestStep{
		ResourceName:            d.ResourceName,
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateVerifyIgnore: d.ImportStateVerifyIgnore,
	}
	if d.GetImportID != nil {
		importStep.ImportStateId = d.GetImportID()
	}
	tc.Steps = []resource.TestStep{{Config: d.GetTerraformConfig()}, importStep}
}

func runIdempotencyResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig()},
		{Config: d.GetTerraformConfig(), PlanOnly: true, ExpectNonEmptyPlan: false},
	}
}

func runExternalDeleteResourceTest(t *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig()},
		{
			PreConfig: func() {
				if err := d.ExternalAction(); err != nil {
					t.Fatalf("ExternalAction: %v", err)
				}
			},
			RefreshState: true, ExpectNonEmptyPlan: true,
		},
	}
}

func runExternalUpdateResourceTest(t *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig()},
		{
			PreConfig: func() {
				if err := d.ExternalAction(); err != nil {
					t.Fatalf("ExternalAction: %v", err)
				}
			},
			RefreshState: true, ExpectNonEmptyPlan: true,
		},
	}
}

func runForceNewResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig()},
		{Config: d.GetUpdatedTerraformConfig(), Check: d.Verifiers.VerifyUpdated},
	}
}

func runInvalidConfigurationResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig(), ExpectError: d.ExpectError},
	}
}

func runAPIErrorResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig(), ExpectError: d.ExpectError},
	}
}

func runAsyncResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig(), Check: d.Verifiers.VerifyAsync},
	}
}

func runTimeoutResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig(), ExpectError: d.ExpectError},
	}
}

func runImportNonExistentResourceTest(_ *testing.T, tc *resource.TestCase, d ContractDefinition) {
	tc.Steps = []resource.TestStep{
		{Config: d.GetTerraformConfig()},
		{
			ResourceName:  d.ResourceName,
			ImportState:   true,
			ImportStateId: d.GetImportID(),
			ExpectError:   d.ExpectError,
		},
	}
}
