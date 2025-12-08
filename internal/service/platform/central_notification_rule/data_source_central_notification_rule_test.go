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
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_refs.#", "1"),
				),
			},
		},
	})
}

func TestAccDataSourceCentralNotificationRule_multipleConditions(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "data.harness_platform_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCentralNotificationRuleMultipleConditions(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.condition_name", "pipeline-condition"),
				),
			},
		},
	})
}

func TestAccDataSourceCentralNotificationRule_multipleChannels(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "data.harness_platform_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCentralNotificationRuleMultipleChannels(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_refs.#", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceCentralNotificationRule_disabled(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "data.harness_platform_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCentralNotificationRuleDisabled(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
				),
			},
		},
	})
}

func TestAccDataSourceCentralNotificationRule_deploymentEvents(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "data.harness_platform_central_notification_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCentralNotificationRuleDeploymentEvents(rName, id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_entity", "PIPELINE"),
					resource.TestCheckResourceAttr(resourceName, "notification_conditions.0.notification_event_configs.0.notification_event", "PIPELINE_FAILED"),
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
		  entity_identifiers = []
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

func testAccDataSourceCentralNotificationRuleMultipleConditions(name, id string) string {
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
			   email_ids = ["notify@harness.io"]
			 }
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
		condition_name = "pipeline-condition"
	
		notification_event_configs {
		  notification_entity = "PIPELINE"
		  notification_event  = "PIPELINE_FAILED"
		  entity_identifiers = []
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

func testAccDataSourceCentralNotificationRuleMultipleChannels(name, id string) string {
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
		
		resource "harness_platform_central_notification_channel" "email_test" {
             depends_on = [
				harness_platform_organization.test,
				harness_platform_project.test,
			]
			 identifier                = "%[1]s_email"
             org                       = harness_platform_organization.test.id
             project                   = harness_platform_project.test.id
			 name                      = "%[2]s_email"
			 notification_channel_type = "EMAIL"
			 status                    = "ENABLED"
			
			 channel {
			   email_ids = ["notify@harness.io"]
			 }
		}

		resource "harness_platform_central_notification_channel" "email2_test" {
             depends_on = [
				harness_platform_organization.test,
				harness_platform_project.test,
			]
			 identifier                = "%[1]s_email2"
             org                       = harness_platform_organization.test.id
             project                   = harness_platform_project.test.id
			 name                      = "%[2]s_email2"
			 notification_channel_type = "EMAIL"
			 status                    = "ENABLED"
			
			 channel {
			   email_ids = ["notify2@harness.io"]
			 }
		}

	resource "harness_platform_central_notification_rule" "test" {
	  depends_on = [
					harness_platform_central_notification_channel.email_test,
					harness_platform_central_notification_channel.email2_test
				]
	  identifier                 = "%[1]s"
	  name                       = "%[2]s"
	  org                        = harness_platform_organization.test.id
	  project                    = harness_platform_project.test.id
	  status                     = "ENABLED"
	  notification_channel_refs  = [
		harness_platform_central_notification_channel.email_test.identifier,
		harness_platform_central_notification_channel.email2_test.identifier
	  ]

	  notification_conditions {
		condition_name = "test-condition"
	
		notification_event_configs {
		  notification_entity = "PIPELINE"
		  notification_event  = "PIPELINE_FAILED"
		  entity_identifiers = []
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

func testAccDataSourceCentralNotificationRuleDisabled(name, id string) string {
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
			   email_ids = ["notify@harness.io"]
			 }
		}

	resource "harness_platform_central_notification_rule" "test" {
	  depends_on = [
					harness_platform_central_notification_channel.test
				]
	  identifier                 = "%[1]s"
	  name                       = "%[2]s"
	  org                        = harness_platform_organization.test.id
	  project                    = harness_platform_project.test.id
	  status                     = "DISABLED"
	  notification_channel_refs  = [harness_platform_central_notification_channel.test.identifier]

	  notification_conditions {
		condition_name = "test-condition"
	
		notification_event_configs {
		  notification_entity = "PIPELINE"
		  notification_event  = "PIPELINE_FAILED"
		  entity_identifiers = []
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

func testAccDataSourceCentralNotificationRuleDeploymentEvents(name, id string) string {
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
			   email_ids = ["notify@harness.io"]
			 }
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
		condition_name = "deployment-condition"
	
		notification_event_configs {
		  notification_entity = "PIPELINE"
		  notification_event  = "PIPELINE_FAILED"
		  entity_identifiers = []
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
