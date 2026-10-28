/*
Acceptance tests only for AutoStopping rule data sources (K8s and VM).
All other resource tests are symmetric and vary only by kind param passed like "instance" or "k8s"
Looks up pre-existing rules by ID and by name regex.

Prerequisites:

	export HARNESS_ACCOUNT_ID="your_account_id"
	export HARNESS_PLATFORM_API_KEY="your_api_key"
	export HARNESS_ENDPOINT="https://app.harness.io/gateway"

	# K8s rule (must exist in the account)
	export AS_DS_TEST_K8S_RULE_ID="40641"
	export AS_DS_TEST_K8S_RULE_NAME="k8scluster23sep2146pm"

	# VM/instance rule (must exist in the account)
	export AS_DS_TEST_VM_RULE_ID="40640"
	export AS_DS_TEST_VM_RULE_NAME="ccm-35727-crosacc-tcp-htkcw2"

Run all acceptance tests:

	TF_ACC=1 go test -v ./internal/service/platform/autostopping/rule/... \
	  -run "TestAccDataSource" -timeout 120m

Run K8s only:

	TF_ACC=1 go test -v ./internal/service/platform/autostopping/rule/... \
	  -run "TestAccDataSourceK8sRule" -timeout 120m

Run VM only:

	TF_ACC=1 go test -v ./internal/service/platform/autostopping/rule/... \
	  -run "TestAccDataSourceVMRule" -timeout 120m

*/

package as_rule_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccDataSourceEnv(t *testing.T, idKey, nameKey string) (ruleID, ruleName string) {
	t.Helper()
	ruleID = os.Getenv(idKey)
	if ruleID == "" {
		t.Skipf("Skipping: %s not set", idKey)
	}
	ruleName = os.Getenv(nameKey)
	if ruleName == "" {
		t.Skipf("Skipping: %s not set", nameKey)
	}
	return
}

// --- K8s data source acceptance tests ---

func TestAccDataSourceK8sRuleById(t *testing.T) {
	ruleID, _ := testAccDataSourceEnv(t, "AS_DS_TEST_K8S_RULE_ID", "AS_DS_TEST_K8S_RULE_NAME")
	dataSourceName := "data.harness_autostopping_rule_k8s.by_id"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "harness_autostopping_rule_k8s" "by_id" {
  identifier = %q
}
`, ruleID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "identifier", ruleID),
					resource.TestCheckResourceAttrSet(dataSourceName, "name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "cloud_connector_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "idle_time_mins"),
				),
			},
		},
	})
}

func TestAccDataSourceK8sRuleByName(t *testing.T) {
	_, ruleName := testAccDataSourceEnv(t, "AS_DS_TEST_K8S_RULE_ID", "AS_DS_TEST_K8S_RULE_NAME")
	namePattern := fmt.Sprintf("^%s$", ruleName)
	dataSourceName := "data.harness_autostopping_rule_k8s.by_name"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "harness_autostopping_rule_k8s" "by_name" {
  name = %q
}
`, namePattern),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "identifier"),
					resource.TestCheckResourceAttr(dataSourceName, "name", ruleName),
					resource.TestCheckResourceAttrSet(dataSourceName, "cloud_connector_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "idle_time_mins"),
				),
			},
		},
	})
}

// --- VM data source acceptance tests ---

func TestAccDataSourceVMRuleById(t *testing.T) {
	ruleID, _ := testAccDataSourceEnv(t, "AS_DS_TEST_VM_RULE_ID", "AS_DS_TEST_VM_RULE_NAME")
	dataSourceName := "data.harness_autostopping_rule_vm.by_id"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "harness_autostopping_rule_vm" "by_id" {
  identifier = %q
}
`, ruleID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "identifier", ruleID),
					resource.TestCheckResourceAttrSet(dataSourceName, "name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "cloud_connector_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "idle_time_mins"),
				),
			},
		},
	})
}

func TestAccDataSourceVMRuleByName(t *testing.T) {
	_, ruleName := testAccDataSourceEnv(t, "AS_DS_TEST_VM_RULE_ID", "AS_DS_TEST_VM_RULE_NAME")
	namePattern := fmt.Sprintf("^%s$", ruleName)
	dataSourceName := "data.harness_autostopping_rule_vm.by_name"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "harness_autostopping_rule_vm" "by_name" {
  name = %q
}
`, namePattern),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "identifier"),
					resource.TestCheckResourceAttr(dataSourceName, "name", ruleName),
					resource.TestCheckResourceAttrSet(dataSourceName, "cloud_connector_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "idle_time_mins"),
				),
			},
		},
	})
}
