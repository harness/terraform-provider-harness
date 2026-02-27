package as_rule_test

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

/*

AS_RULES_TEST_PROVIDER_BLOCK="true"  to be set only for testing with dev testing only
AS_RULES_TEST_DRY_RUN="false" is required to really run the test cases

# Run the test case only
AS_RULES_TEST_CLOUD_CONNECTOR_ID="awsconn05" \
AS_RULES_TEST_VM_ID="i-0c3591bd03f1f649e" \
AS_RULES_TEST_REGION="ap-south-1" \
AS_RULES_TEST_DRY_RUN="false" \
AS_RULES_TEST_PROVIDER_BLOCK="true" \
go test -v   \
    ./internal/service/platform/autostopping/rule/...  \
    -run TestAccDSRulesAll

*/

func TestAccDSRulesAll(t *testing.T) {
	TestAccDSRulesNoFilter(t)
	if t.Failed() {
		t.Log("TestAccDSRulesNoFilter FAILED – stopping")
		return
	}

	TestAccDSRulesKindInstance(t)
	if t.Failed() {
		t.Log("TestAccDSRulesKindInstance FAILED – stopping")
		return
	}

	TestAccDSRulesNamePrefix(t)
	if t.Failed() {
		t.Log("TestAccDSRulesNamePrefix FAILED – stopping")
		return
	}

	TestAccDSRulesNameRegex(t)
	if t.Failed() {
		t.Log("TestAccDSRulesNameRegex FAILED – stopping")
	}
}

// TestAccDSRulesNoFilter creates one VM rule and queries the data source with no filter.
func TestAccDSRulesNoFilter(t *testing.T) {
	isDryRunOnly := os.Getenv("AS_RULES_TEST_DRY_RUN") == "true"
	cloudConnectorID, vmID, region := testAccDSRulesEnvVars(t, isDryRunOnly)

	suffix := randLower(6)
	name := fmt.Sprintf("ds-nofilter-%s", suffix)

	cfg := getNewVMRuleBlock("svc", name, cloudConnectorID, vmID, region) + `
data "harness_autostopping_rules" "all" {
  depends_on = [harness_autostopping_rule_vm.svc]
}
`
	if isDryRunOnly {
		t.Logf("\n=== generated Terraform source ===\n%s", cfg)
		return
	}

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.harness_autostopping_rules.all", "id"),
					resource.TestCheckResourceAttrSet("data.harness_autostopping_rules.all", "rules.#"),
				),
			},
		},
	})
}

func TestAccDSRulesKindInstance(t *testing.T) {
	isDryRunOnly := os.Getenv("AS_RULES_TEST_DRY_RUN") == "true"
	cloudConnectorID, vmID, region := testAccDSRulesEnvVars(t, isDryRunOnly)

	suffix := randLower(6)
	name := fmt.Sprintf("service-%s-02", suffix)

	cfg := getNewVMRuleBlock("svc", name, cloudConnectorID, vmID, region) + `
data "harness_autostopping_rules" "by_instance_kind" {
  kind       = "instance"
  depends_on = [harness_autostopping_rule_vm.svc]
}
`
	if isDryRunOnly {
		t.Logf("\n=== generated Terraform source ===\n%s", cfg)
		return
	}

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.harness_autostopping_rules.by_instance_kind", "kind", "instance"),
					resource.TestCheckResourceAttrSet("data.harness_autostopping_rules.by_instance_kind", "rules.#"),
				),
			},
		},
	})
}

func TestAccDSRulesNamePrefix(t *testing.T) {
	isDryRunOnly := os.Getenv("AS_RULES_TEST_DRY_RUN") == "true"
	cloudConnectorID, vmID, region := testAccDSRulesEnvVars(t, isDryRunOnly)

	suffix := randLower(6)
	name := fmt.Sprintf("test-svc-%s-01", suffix)
	nameFilter := fmt.Sprintf("^test-svc-%s.*", suffix)

	cfg := getNewVMRuleBlock("svc", name, cloudConnectorID, vmID, region) + fmt.Sprintf(`
data "harness_autostopping_rules" "by_name_prefix" {
  name       = %q
  depends_on = [harness_autostopping_rule_vm.svc]
}
`, nameFilter)

	if isDryRunOnly {
		t.Logf("\n=== generated Terraform source ===\n%s", cfg)
		return
	}

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.harness_autostopping_rules.by_name_prefix", "name", nameFilter),
					resource.TestCheckResourceAttrSet("data.harness_autostopping_rules.by_name_prefix", "rules.#"),
				),
			},
		},
	})
}

func TestAccDSRulesNameRegex(t *testing.T) {
	isDryRunOnly := os.Getenv("AS_RULES_TEST_DRY_RUN") == "true"
	cloudConnectorID, vmID, region := testAccDSRulesEnvVars(t, isDryRunOnly)

	suffix := randLower(6)
	name := fmt.Sprintf("test-svc-%s-02", suffix)
	nameFilter := fmt.Sprintf("^(test-app|test-svc)-%s.*", suffix)

	cfg := getNewVMRuleBlock("svc", name, cloudConnectorID, vmID, region) + fmt.Sprintf(`
data "harness_autostopping_rules" "by_name_regex" {
  name       = %q
  depends_on = [harness_autostopping_rule_vm.svc]
}
`, nameFilter)

	if isDryRunOnly {
		t.Logf("\n=== generated Terraform source ===\n%s", cfg)
		return
	}

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.harness_autostopping_rules.by_name_regex", "name", nameFilter),
					resource.TestCheckResourceAttrSet("data.harness_autostopping_rules.by_name_regex", "rules.#"),
				),
			},
		},
	})
}

const providerBlockStr = `terraform {
  required_providers {
    harness = {
      source = "harness/harness"
      version = "0.40.2"
    }
  }
}

`
const vmRuleConstStr = `
resource "harness_autostopping_rule_vm" %[1]q {
  name               = %[2]q
  cloud_connector_id = %[3]q
  idle_time_mins     = 15

  filter {
    vm_ids  = [%[4]q]
    regions = [%[5]q]
  }
}
`

func randLower(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func testAccDSRulesEnvVars(t *testing.T, isDryRunOnly bool) (cloudConnectorID, vmID, region string) {
	t.Helper()
	if isDryRunOnly {
		return "<cloud-connector-id>", "<vm-id>", "<region>"
	}
	cloudConnectorID = os.Getenv("AS_RULES_TEST_CLOUD_CONNECTOR_ID")
	vmID = os.Getenv("AS_RULES_TEST_VM_ID")
	region = os.Getenv("AS_RULES_TEST_REGION")
	if cloudConnectorID == "" {
		t.Fatalf("AS_RULES_TEST_CLOUD_CONNECTOR_ID must be set for acceptance tests")
	}
	if vmID == "" {
		t.Fatalf("AS_RULES_TEST_VM_ID must be set for acceptance tests")
	}
	if region == "" {
		t.Fatalf("AS_RULES_TEST_REGION must be set for acceptance tests")
	}
	return
}

func getNewVMRuleBlock(resourceLabel, name, cloudConnectorID, vmID, region string) string {
	vmRuleStr := strings.Clone(vmRuleConstStr)
	vmRuleStr = fmt.Sprintf(vmRuleStr, resourceLabel, name, cloudConnectorID, vmID, region)

	if os.Getenv("AS_RULES_TEST_PROVIDER_BLOCK") == "true" {
		return providerBlockStr + vmRuleStr
	}

	return vmRuleStr
}
