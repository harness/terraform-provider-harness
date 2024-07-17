package notification_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNotificationRule(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_notification_rule.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNotificationRule(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testAccDataSourceNotificationRule(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_notification_rule" "test" {
			depends_on = [
				harness_platform_project.test,
			]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			identifier = "%[1]s"
			request {
				  name = "%[2]s"
				  notification_method {
					type                  = "Slack"
					spec = jsonencode({
						webhook_url = "http://myslackwebhookurl.com"
						user_groups = ["account.dsd"]
					})
				  }
				  type = "ServiceLevelObjective"
				  conditions {
					type       = "ErrorBudgetRemainingPercentage"
					spec = jsonencode({
					threshold = 30
					})
				  }
			}
		}

		data "harness_platform_notification_rule" "test" {
			identifier = harness_platform_notification_rule.test.identifier
			org_id = harness_platform_notification_rule.test.org_id
			project_id = harness_platform_notification_rule.test.project_id
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
`, id, name)
}
