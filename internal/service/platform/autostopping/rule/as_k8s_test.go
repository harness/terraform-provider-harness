package as_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceK8sRule(t *testing.T) {
	name := utils.RandStringBytes(5)
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
  cloud_connector_id  = %[2]q
  k8s_connector_id    = %[3]q
  k8s_namespace       = "default"
  idle_time_mins      = 10
  dry_run             = %[4]t

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
`, name, cloudConnectorIDK8s, k8sConnectorID, dryRun)
}

func testK8sRuleUpdate(name, idleTime string, dryRun bool) string {
	return fmt.Sprintf(`
resource "harness_autostopping_rule_k8s" "test" {
  name                = "%[1]s"
  cloud_connector_id  = %[2]q
  k8s_connector_id    = %[3]q
  k8s_namespace       = "default"
  idle_time_mins      = %[4]s
  dry_run             = %[5]t

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
  idleTimeMins: %[4]s
  hideProgressPage: false
  dependencies: []
EOT
}
`, name, cloudConnectorIDK8s, k8sConnectorID, idleTime, dryRun)
}
