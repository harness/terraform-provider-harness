package alert_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestResourceAlert(t *testing.T) {
	name := "terraform-alert-test"
	resourceName := "harness_autostopping_alert.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAlertConfig(name, true, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "applicable_to_all_rules", "true"),
				),
			},
			{
				Config: testAlertConfig(name, false, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
		},
	})
}

func testAlertConfig(name string, enabled, applicableToAll bool) string {
	applicableBlock := `applicable_to_all_rules = true`
	if !applicableToAll {
		applicableBlock = `rule_id_list = [1234]`
	}
	return fmt.Sprintf(`
resource "harness_autostopping_alert" "test" {
  name    = "%[1]s"
  enabled = %[2]t
  recipients {
    email = ["user@example.com"]
  }
  events = [
    "autostopping_rule_created",
    "autostopping_warmup_failed",
    "autostopping_cooldown_failed"
  ]
  %[3]s
}
`, name, enabled, applicableBlock)
}
