package default_notification_template_set_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceDefaultNotificationTemplateSet_basic(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_default_notification_template_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDefaultNotificationTemplateSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDefaultNotificationTemplateSetBasic(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_entity", "PIPELINE"),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "event_template_configuration_set.#", "1"),
				),
			},
			{
				Config: testAccResourceDefaultNotificationTemplateSetBasic(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceDefaultNotificationTemplateSet_projectLevel(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_default_notification_template_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDefaultNotificationTemplateSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDefaultNotificationTemplateSetProjectLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceDefaultNotificationTemplateSet_orgLevel(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_default_notification_template_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDefaultNotificationTemplateSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDefaultNotificationTemplateSetOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckNoResourceAttr(resourceName, "project_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceDefaultNotificationTemplateSet_multipleChannelTypes(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_default_notification_template_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDefaultNotificationTemplateSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDefaultNotificationTemplateSetSlack(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "SLACK"),
				),
			},
			{
				Config: testAccResourceDefaultNotificationTemplateSetMSTeams(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "notification_channel_type", "MSTEAMS"),
				),
			},
		},
	})
}

func TestAccResourceDefaultNotificationTemplateSet_multipleEventConfigurations(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_default_notification_template_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDefaultNotificationTemplateSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDefaultNotificationTemplateSetMultipleEvents(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "notification_entity", "DEPLOYMENT"),
					resource.TestCheckResourceAttr(resourceName, "event_template_configuration_set.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetDefaultNotificationTemplateSet(resourceName string, state *terraform.State) (*nextgen.DefaultNotificationTemplateSetResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	var resp nextgen.DefaultNotificationTemplateSetResponse
	var err error

	org := r.Primary.Attributes["org_id"]
	if org == "" {
		org = r.Primary.Attributes["org"]
	}
	project := r.Primary.Attributes["project_id"]
	if project == "" {
		project = r.Primary.Attributes["project"]
	}

	if org != "" && project != "" {
		resp, _, err = c.ProjectDefaultNotificationTemplateSetApi.GetProjectDefaultNotificationTemplateSet(ctx, id, org, project, &nextgen.ProjectDefaultNotificationTemplateSetApiGetProjectDefaultNotificationTemplateSetOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
	} else if org != "" {
		resp, _, err = c.OrgDefaultNotificationTemplateSetApi.GetOrgDefaultNotificationTemplateSet(ctx, id, org, &nextgen.OrgDefaultNotificationTemplateSetApiGetOrgDefaultNotificationTemplateSetOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
	} else {
		resp, _, err = c.AccountDefaultNotificationTemplateSetApi.GetAccountDefaultNotificationTemplateSet(ctx, id, &nextgen.AccountDefaultNotificationTemplateSetApiGetAccountDefaultNotificationTemplateSetOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
	}

	if err != nil {
		return nil, err
	}

	if resp.DefaultNotificationTemplateSet == nil {
		return nil, fmt.Errorf("empty template set received in response")
	}

	return &resp, nil
}

func testAccDefaultNotificationTemplateSetDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		templateSetResp, _ := testAccGetDefaultNotificationTemplateSet(resourceName, state)
		if templateSetResp != nil && templateSetResp.DefaultNotificationTemplateSet != nil {
			return fmt.Errorf("found default notification template set: %s", templateSetResp.DefaultNotificationTemplateSet.Identifier)
		}

		return nil
	}
}

func testAccResourceDefaultNotificationTemplateSetBasic(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_default_notification_template_set" "test" {
			identifier                  = "%[1]s"
			name                       = "%[2]s"
			description                = "Test default notification template set"
			notification_entity        = "PIPELINE"
			notification_channel_type  = "EMAIL"
			
			event_template_configuration_set {
				notification_events = ["PIPELINE_START", "PIPELINE_SUCCESS", "PIPELINE_FAILED"]
				
				template {
					template_ref   = "email_pipeline_template"
					version_label  = "v1"
				}
			}
		}
	`, id, name)
}

func testAccResourceDefaultNotificationTemplateSetProjectLevel(id string, name string) string {
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

		resource "harness_platform_default_notification_template_set" "test" {
			depends_on = [
				harness_platform_organization.test,
				harness_platform_project.test,
			]
			identifier                  = "%[1]s"
			name                       = "%[2]s"
			org_id                     = harness_platform_organization.test.id
			project_id                 = harness_platform_project.test.id
			description                = "Test project level notification template set"
			notification_entity        = "PIPELINE"
			notification_channel_type  = "EMAIL"
			
			event_template_configuration_set {
				notification_events = ["PIPELINE_START"]
				
				template {
					template_ref   = "project_pipeline_template"
					version_label  = "v1"
				}
			}
		}
	`, id, name)
}

func testAccResourceDefaultNotificationTemplateSetOrgLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		resource "harness_platform_default_notification_template_set" "test" {
			depends_on = [
				harness_platform_organization.test,
			]
			identifier                  = "%[1]s"
			name                       = "%[2]s"
			org_id                     = harness_platform_organization.test.id
			description                = "Test org level notification template set"
			notification_entity        = "PIPELINE"
			notification_channel_type  = "EMAIL"
			
			event_template_configuration_set {
				notification_events = ["PIPELINE_START"]
				
				template {
					template_ref   = "org_pipeline_template"
					version_label  = "v1"
				}
			}
		}
	`, id, name)
}

func testAccResourceDefaultNotificationTemplateSetSlack(id string, name string) string {
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
					template_ref   = "slack_pipeline_template"
					version_label  = "v1"
				}
			}
		}
	`, id, name)
}

func testAccResourceDefaultNotificationTemplateSetMSTeams(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_default_notification_template_set" "test" {
			identifier                  = "%[1]s"
			name                       = "%[2]s"
			description                = "Test MS Teams notification template set"
			notification_entity        = "PIPELINE"
			notification_channel_type  = "MSTEAMS"
			
			event_template_configuration_set {
				notification_events = ["PIPELINE_START", "PIPELINE_FAILED"]
				
				template {
					template_ref   = "msteams_pipeline_template"
					version_label  = "v1"
				}
			}
		}
	`, id, name)
}

func testAccResourceDefaultNotificationTemplateSetMultipleEvents(id string, name string) string {
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
	`, id, name)
}
