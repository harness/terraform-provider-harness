package central_notification_rule_test

import (
	"fmt"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCentralNotificationRule_basic(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "data.harness_platform_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCentralNotificationRuleConfig(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
		},
	})
}

func testAccDataSourceCentralNotificationRuleConfig(name, id string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
			color      = "#472848"
		}
		
		resource "harness_platform_central_notification_channel" "test" {
             depends_on = [
				harness_platform_organization.test,
				harness_platform_project.test,
			]
			 identifier                = "%[1]s"
             org                       = harness_platform_organization.test.id
             project                   = harness_platform_project.test.id
			 name                      = "%[2]s"
			 notification_channel_type = "EMAIL"
			 status                    = "ENABLED"
			
			 channel {
			   email_ids            = ["notify@harness.io"]
			 }
			}

		data "harness_platform_central_notification_channel" "test" {
			identifier = harness_platform_central_notification_channel.test.identifier
			org     = harness_platform_organization.test.id
			project = harness_platform_project.test.id
		}
	resource "harness_platform_central_notification_rule" "test" {
	  depends_on = [
					harness_platform_central_notification_channel.test
				]
	  identifier                 = "%[1]s"
	  name                       = "%[2]s"
	  org                        = harness_platform_organization.test.id
	  project                    = harness_platform_project.test.id
	  status                     = "ENABLED"
	  notification_channel_refs  = [harness_platform_central_notification_channel.test.identifier]

	  notification_conditions {
		condition_name = "test-condition"
	
		notification_event_configs {
		  notification_entity = "PIPELINE"
		  notification_event  = "PIPELINE_FAILED"
		  notification_event_data = {
			foo = "bar"
		  }
		}
	  }
	}

	data "harness_platform_central_notification_rule" "test" {
	  identifier = harness_platform_central_notification_rule.test.identifier
	  org        = harness_platform_central_notification_rule.test.org
	  project    = harness_platform_central_notification_rule.test.project
	}
`, id, name)
}
