package central_notification_channel_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCentralNotificationChannel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_central_notification_channel.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCentralNotificationChannel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testAccDataSourceCentralNotificationChannel(id string, name string) string {
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
				org_id     = harness_platform_organization.test.id
				project_id = harness_platform_project.test.id
			 name                      = "%[2]s"
			 notification_channel_type = "EMAIL"
			 status                    = "ENABLED"
			
			 channel {
			   email_ids            = ["notify@harness.io"]
			   api_key              = "dummy-api-key"
			   execute_on_delegate  = true
			   user_groups {
				 identifier = "account.test"
			   }
			   headers {
				 key   = "X-Custom-Header"
				 value = "HeaderValue"
			   }
			 }
			}

		data "harness_platform_central_notification_channel" "test" {
			identifier = harness_platform_central_notification_channel.test.identifier
			org_id     = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}
