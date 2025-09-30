package default_notification_template_set_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDefaultNotificationTemplateSet_basic(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_default_notification_template_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDefaultNotificationTemplateSetBasic(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_entity", "PIPELINE"),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "event_template_configuration_set.#", "1"),
				),
			},
		},
	})
}

func TestAccDataSourceDefaultNotificationTemplateSet_slack(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_default_notification_template_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDefaultNotificationTemplateSetSlack(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_entity", "PIPELINE"),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "SLACK"),
					resource.TestCheckResourceAttr(resourceName, "event_template_configuration_set.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "event_template_configuration_set.0.notification_events.#", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceDefaultNotificationTemplateSet_multipleEvents(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_default_notification_template_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDefaultNotificationTemplateSetMultipleEvents(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_entity", "DEPLOYMENT"),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "event_template_configuration_set.#", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceDefaultNotificationTemplateSet_withVariables(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_default_notification_template_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDefaultNotificationTemplateSetWithVariables(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_entity", "PIPELINE"),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "event_template_configuration_set.0.template.0.variables.#", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceDefaultNotificationTemplateSet_withTags(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_default_notification_template_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDefaultNotificationTemplateSetWithTags(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "tags.environment", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.team", "platform"),
				),
			},
		},
	})
}

func testAccDataSourceDefaultNotificationTemplateSetBasic(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_default_notification_template_set" "test" {
			identifier                  = "%[1]s"
			name                       = "%[2]s"
			description                = "Test default notification template set"
			notification_entity        = "PIPELINE"
			notification_channel_type  = "EMAIL"
			
			event_template_configuration_set {
				notification_events = ["PIPELINE_START"]
				
				template {
					template_ref   = "test_template"
					version_label  = "v1"
				}
			}
		}

		data "harness_platform_default_notification_template_set" "test" {
			identifier                  = harness_platform_default_notification_template_set.test.identifier
			name                       = harness_platform_default_notification_template_set.test.name
			notification_entity        = harness_platform_default_notification_template_set.test.notification_entity
			notification_channel_type  = harness_platform_default_notification_template_set.test.notification_channel_type
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}

func testAccDataSourceDefaultNotificationTemplateSetSlack(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_default_notification_template_set" "test" {
			identifier                  = "%[1]s"
			name                       = "%[2]s"
			description                = "Test slack notification template set"
			notification_entity        = "PIPELINE"
			notification_channel_type  = "SLACK"
			
			event_template_configuration_set {
				notification_events = ["PIPELINE_START", "PIPELINE_SUCCESS"]
				
				template {
					template_ref   = "slack_template"
					version_label  = "v1"
				}
			}
		}

		data "harness_platform_default_notification_template_set" "test" {
			identifier                  = harness_platform_default_notification_template_set.test.identifier
			name                       = harness_platform_default_notification_template_set.test.name
			notification_entity        = harness_platform_default_notification_template_set.test.notification_entity
			notification_channel_type  = harness_platform_default_notification_template_set.test.notification_channel_type
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}

func testAccDataSourceDefaultNotificationTemplateSetMultipleEvents(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_default_notification_template_set" "test" {
			identifier                  = "%[1]s"
			name                       = "%[2]s"
			description                = "Test multiple events notification template set"
			notification_entity        = "DEPLOYMENT"
			notification_channel_type  = "EMAIL"
			
			event_template_configuration_set {
				notification_events = ["DEPLOYMENT_START"]
				
				template {
					template_ref   = "deployment_start_template"
					version_label  = "v1"
				}
			}

			event_template_configuration_set {
				notification_events = ["DEPLOYMENT_FAILED", "DEPLOYMENT_SUCCESS"]
				
				template {
					template_ref   = "deployment_end_template"
					version_label  = "v1"
				}
			}
		}

		data "harness_platform_default_notification_template_set" "test" {
			identifier                  = harness_platform_default_notification_template_set.test.identifier
			name                       = harness_platform_default_notification_template_set.test.name
			notification_entity        = harness_platform_default_notification_template_set.test.notification_entity
			notification_channel_type  = harness_platform_default_notification_template_set.test.notification_channel_type
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}

func testAccDataSourceDefaultNotificationTemplateSetWithVariables(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_default_notification_template_set" "test" {
			identifier                  = "%[1]s"
			name                       = "%[2]s"
			description                = "Test notification template set with variables"
			notification_entity        = "PIPELINE"
			notification_channel_type  = "EMAIL"
			
			event_template_configuration_set {
				notification_events = ["PIPELINE_FAILED"]
				
				template {
					template_ref   = "pipeline_failed_template"
					version_label  = "v1"
					
					variables {
						name  = "pipeline_name"
						value = "${pipeline.name}"
						type  = "STRING"
					}
					
					variables {
						name  = "failure_reason"
						value = "${pipeline.failure.reason}"
						type  = "STRING"
					}
				}
			}
		}

		data "harness_platform_default_notification_template_set" "test" {
			identifier                  = harness_platform_default_notification_template_set.test.identifier
			name                       = harness_platform_default_notification_template_set.test.name
			notification_entity        = harness_platform_default_notification_template_set.test.notification_entity
			notification_channel_type  = harness_platform_default_notification_template_set.test.notification_channel_type
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}

func testAccDataSourceDefaultNotificationTemplateSetWithTags(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_default_notification_template_set" "test" {
			identifier                  = "%[1]s"
			name                       = "%[2]s"
			description                = "Test notification template set with tags"
			notification_entity        = "PIPELINE"
			notification_channel_type  = "EMAIL"
			
			event_template_configuration_set {
				notification_events = ["PIPELINE_START"]
				
				template {
					template_ref   = "tagged_template"
					version_label  = "v1"
				}
			}
			
			tags = {
				environment = "test"
				team        = "platform"
			}
		}

		data "harness_platform_default_notification_template_set" "test" {
			identifier                  = harness_platform_default_notification_template_set.test.identifier
			name                       = harness_platform_default_notification_template_set.test.name
			notification_entity        = harness_platform_default_notification_template_set.test.notification_entity
			notification_channel_type  = harness_platform_default_notification_template_set.test.notification_channel_type
		}

		resource "time_sleep" "wait_4_seconds" {
			destroy_duration = "4s"
		}
	`, id, name)
}