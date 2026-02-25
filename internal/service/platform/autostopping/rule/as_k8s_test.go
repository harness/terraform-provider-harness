package as_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceK8sRule(t *testing.T) {
	name := "terraform-rule-test-k8s"
	resourceName := "harness_autostopping_rule_k8s.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testK8sRule(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
					resource.TestCheckResourceAttr(resourceName, "k8s_namespace", "default"),
				),
			},
			{
				Config: testK8sRule(name, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
			{
				Config: testK8sRuleUpdate(name, "15", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "idle_time_mins", "15"),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "true"),
				),
			},
			{
				Config: testK8sRuleUpdate(name, "20", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "idle_time_mins", "20"),
					resource.TestCheckResourceAttr(resourceName, "dry_run", "false"),
				),
			},
		},
	})
}

func testK8sRule(name string, dryRun bool) string {
	return fmt.Sprintf(`
resource "harness_autostopping_rule_k8s" "test" {
  name                = "%[1]s"
  cloud_connector_id  = "granpermissions"
  k8s_connector_id    = "account.k8s_connector"
  k8s_namespace       = "default"
  idle_time_mins      = 10
  dry_run             = %[2]t

  rule_yaml = <<-EOT
apiVersion: ccm.harness.io/v1
kind: AutoStoppingRule
metadata:
  name: %[1]s
  namespace: default
spec:
  service:
    name: nginx
    port: 80
  idleTimeMins: 10
  hideProgressPage: false
  dependencies: []
EOT
}
`, name, dryRun)
}

func testK8sRuleUpdate(name, idleTime string, dryRun bool) string {
	return fmt.Sprintf(`
resource "harness_autostopping_rule_k8s" "test" {
  name                = "%[1]s"
  cloud_connector_id  = "granpermissions"
  k8s_connector_id    = "account.k8s_connector"
  k8s_namespace       = "default"
  idle_time_mins      = %[2]s
  dry_run             = %[3]t

  rule_yaml = <<-EOT
apiVersion: ccm.harness.io/v1
kind: AutoStoppingRule
metadata:
  name: %[1]s
  namespace: default
spec:
  service:
    name: nginx
    port: 80
  idleTimeMins: %[2]s
  hideProgressPage: false
  dependencies: []
EOT
}
`, name, idleTime, dryRun)
}
