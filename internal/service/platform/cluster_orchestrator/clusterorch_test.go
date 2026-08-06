package cluster_orchestrator_test

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/harness/terraform-provider-harness/internal/service/platform/cluster_orchestrator"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func funcName(f interface{}) string {
	if f == nil {
		return ""
	}
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func TestResourceClusterOrchestratorLifecycleWiring(t *testing.T) {
	r := cluster_orchestrator.ResourceClusterOrchestrator()

	if err := r.InternalValidate(nil, true); err != nil {
		t.Fatalf("schema failed InternalValidate: %v", err)
	}

	createName := funcName(r.CreateContext)
	readName := funcName(r.ReadContext)
	deleteName := funcName(r.DeleteContext)

	if readName == "" {
		t.Fatal("ReadContext must be set")
	}
	if readName == createName {
		t.Fatalf("ReadContext must not be the Create handler, got %s", readName)
	}
	if r.UpdateContext != nil && funcName(r.UpdateContext) == createName {
		t.Fatalf("UpdateContext must not be the Create handler, got %s", funcName(r.UpdateContext))
	}
	if deleteName == createName {
		t.Fatalf("DeleteContext must not be the Create handler, got %s", deleteName)
	}

	for _, field := range []string{"name", "cluster_endpoint", "k8s_connector_id", "region"} {
		s, ok := r.Schema[field]
		if !ok {
			t.Fatalf("expected schema field %q to exist", field)
		}
		if r.UpdateContext == nil && !s.ForceNew {
			t.Errorf("field %q must be ForceNew since no UpdateContext is defined and the API has no update endpoint for it", field)
		}
	}
}

func TestResourceClusterOrchestrator(t *testing.T) {
	name := "terraform-clusterorch-test"
	resourceName := "harness_cluster_orchestrator.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		// CheckDestroy:      testRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testClusterOrch(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

// TestResourceClusterOrchestrator_CCM32336_OutOfBandDeleteRecreates verifies that
// when a Cluster Orchestrator is deleted out-of-band (UI / direct API), the next
// terraform refresh does not error and re-plans a create.
//
// Regression test for CCM-32336 (lwd GET returning HTTP 500 for a deleted entity
// causes terraform plan to fail with "giving up after 11 attempt(s)").
//
// Skipped until CCM-34967: DELETE of a missing orchestrator still returns HTTP 500
// ("invalid cluster id"), so post-test destroy fails after OOB delete.
func TestResourceClusterOrchestrator_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	t.Skip("CCM-34967: cluster orchestrator DELETE returns 500 for missing ID; re-enable after fix")

	name := "terraform-co-ccm32336-test"
	resourceName := "harness_cluster_orchestrator.test"

	var orchIDBefore string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testClusterOrch(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrWith(resourceName, "id", func(value string) error {
						orchIDBefore = value
						return nil
					}),
				),
			},
			{
				PreConfig: func() {
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					if _, err := c.CloudCostClusterOrchestratorApi.DeleteClusterOrchestrator(
						ctx, c.AccountId, orchIDBefore,
					); err != nil {
						t.Fatalf("CCM-32336: out-of-band delete failed: %v", err)
					}
				},
				Config: testClusterOrch(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testClusterOrch(name string) string {
	return fmt.Sprintf(`
	resource "harness_cluster_orchestrator" "test" {
		name = "%s"  
		cluster_endpoint = "http://test.com" 
		k8s_connector_id = "TestDoNotDelete"                    
	}
`, name)
}
