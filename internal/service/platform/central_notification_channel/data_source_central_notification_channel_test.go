package central_notification_channel_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCentralNotificationChannel_email(t *testing.T) {
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
					resource.TestCheckResourceAttr(resourceName, "org", id),
					resource.TestCheckResourceAttr(resourceName, "project", id),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
				),
			},
		},
	})
}

func TestAccDataSourceCentralNotificationChannel_slack(t *testing.T) {
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
				Config: testAccDataSourceCentralNotificationChannelSlack(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "SLACK"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.slack_webhook_urls.0", "https://hooks.slack.com/services/test"),
				),
			},
		},
	})
}

func TestAccDataSourceCentralNotificationChannel_msteams(t *testing.T) {
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
				Config: testAccDataSourceCentralNotificationChannelMSTeams(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "MSTEAMS"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.ms_team_keys.0", "https://outlook.office.com/webhook/test"),
				),
			},
		},
	})
}

func TestAccDataSourceCentralNotificationChannel_pagerduty(t *testing.T) {
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
				Config: testAccDataSourceCentralNotificationChannelPagerDuty(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "PAGERDUTY"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.pager_duty_integration_keys.0", "test-integration-key"),
				),
			},
		},
	})
}

func TestAccDataSourceCentralNotificationChannel_webhook(t *testing.T) {
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
				Config: testAccDataSourceCentralNotificationChannelWebhook(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "WEBHOOK"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.webhook_urls.0", "https://webhook.example.com/test"),
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
				org     = harness_platform_organization.test.id
				project = harness_platform_project.test.id
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

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}

func testAccDataSourceCentralNotificationChannelSlack(id string, name string) string {
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
			 notification_channel_type = "SLACK"
			 status                    = "ENABLED"
			
			 channel {
			   slack_webhook_urls = ["https://hooks.slack.com/services/test"]
			 }
		}

		data "harness_platform_central_notification_channel" "test" {
			identifier = harness_platform_central_notification_channel.test.identifier
			org        = harness_platform_organization.test.id
			project    = harness_platform_project.test.id
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}

func testAccDataSourceCentralNotificationChannelMSTeams(id string, name string) string {
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
			 notification_channel_type = "MSTEAMS"
			 status                    = "ENABLED"
			
			 channel {
			   ms_team_keys = ["https://outlook.office.com/webhook/test"]
			 }
		}

		data "harness_platform_central_notification_channel" "test" {
			identifier = harness_platform_central_notification_channel.test.identifier
			org        = harness_platform_organization.test.id
			project    = harness_platform_project.test.id
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}

func testAccDataSourceCentralNotificationChannelPagerDuty(id string, name string) string {
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
			 notification_channel_type = "PAGERDUTY"
			 status                    = "ENABLED"
			
			 channel {
			   pager_duty_integration_keys = ["test-integration-key"]
			 }
		}

		data "harness_platform_central_notification_channel" "test" {
			identifier = harness_platform_central_notification_channel.test.identifier
			org        = harness_platform_organization.test.id
			project    = harness_platform_project.test.id
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}

func testAccDataSourceCentralNotificationChannelWebhook(id string, name string) string {
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
			 notification_channel_type = "WEBHOOK"
			 status                    = "ENABLED"
			
			 channel {
			   webhook_urls = ["https://webhook.example.com/test"]
			 }
		}

		data "harness_platform_central_notification_channel" "test" {
			identifier = harness_platform_central_notification_channel.test.identifier
			org        = harness_platform_organization.test.id
			project    = harness_platform_project.test.id
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}

func TestAccDataSourceCentralNotificationChannel_datadog(t *testing.T) {
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
				Config: testAccDataSourceCentralNotificationChannelDatadog(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "DATADOG"),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.datadog_urls.0", "https://api.datadoghq.com/api/v1/events"),
					resource.TestCheckResourceAttr(resourceName, "channel.0.api_key", "test-api-key"),
				),
			},
		},
	})
}

func testAccDataSourceCentralNotificationChannelDatadog(id string, name string) string {
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
		 notification_channel_type = "DATADOG"
		 status                    = "ENABLED"
		 
		 channel {
		   datadog_urls = ["https://api.datadoghq.com/api/v1/events"]
		   api_key = "test-api-key"
		 }
		}

		data "harness_platform_central_notification_channel" "test" {
			identifier = harness_platform_central_notification_channel.test.identifier
			org        = harness_platform_organization.test.id
			project    = harness_platform_project.test.id
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}

