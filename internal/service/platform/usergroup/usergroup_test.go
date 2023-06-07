package usergroup_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceUserGroup(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_usergroup.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroup(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceUserGroup(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"users"},
			},
		},
	})
}

func TestProjectResourceUserGroup(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_usergroup.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceUserGroup(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testProjectResourceUserGroup(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"users"},
			},
		},
	})
}

func TestOrgResourceUserGroup(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_usergroup.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceUserGroup(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testOrgResourceUserGroup(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"users"},
			},
		},
	})
}

func TestAccResourceUserGroup_emails(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_usergroup.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroup_emails(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "user_emails.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.0", "meet.meet@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.1", "richaterraformtest@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.2", "jajooricha@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.3", "aditya.kashyap@harness.io"),
				),
			},
			{
				Config: testAccResourceUserGroup_emails(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "user_emails.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.0", "meet.meet@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.1", "richaterraformtest@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.2", "jajooricha@gmail.com"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.3", "aditya.kashyap@harness.io"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"user_emails"},
			},
		},
	})
}

func TestProjectResourceUserGroup_emails(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_usergroup.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceUserGroup_emails(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "user_emails.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.0", "rathod.meetsatish@harness.io"),
				),
			},
			{
				Config: testProjectResourceUserGroup_emails(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "user_emails.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.0", "rathod.meetsatish@harness.io"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"user_emails"},
			},
		},
	})
}

func TestOrgResourceUserGroup_emails(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_usergroup.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceUserGroup_emails(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "user_emails.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.0", "rathod.meetsatish@harness.io"),
				),
			},
			{
				Config: testOrgResourceUserGroup_emails(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "user_emails.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "user_emails.0", "rathod.meetsatish@harness.io"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"user_emails"},
			},
		},
	})
}

func TestAccResourceUserGroup_DeleteUnderlyingResource(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_usergroup.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceUserGroup(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccProvider.Meta().(*internal.Session).GetPlatformClient()
					_, _, err := c.UserGroupApi.DeleteUserGroup(ctx, c.AccountId, id, &nextgen.UserGroupApiDeleteUserGroupOpts{
						OrgIdentifier:     optional.NewString(id),
						ProjectIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)

				},
				Config:             testAccResourceUserGroup(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGetPlatformUserGroup(resourceName string, state *terraform.State) (*nextgen.UserGroup, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]

	resp, _, err := c.UserGroupApi.GetUserGroup(ctx, c.AccountId, id, &nextgen.UserGroupApiGetUserGroupOpts{
		OrgIdentifier:     optional.NewString(orgId),
		ProjectIdentifier: optional.NewString(projId),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return resp.Data, nil
}

func testAccUserGroupDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetPlatformUserGroup(resourceName, state)
		if env != nil {
			return fmt.Errorf("Found usergroup: %s", env.Identifier)
		}

		return nil
	}
}

func testAccResourceUserGroup(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_usergroup" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			linked_sso_id = "linked_sso_id"
			externally_managed = false
			users = []
			notification_configs {
				type = "SLACK"
				slack_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "EMAIL"
				group_email = "email@email.com"
				send_email_to_all_users = true
			}
			notification_configs {
				type = "MSTEAMS"
				microsoft_teams_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "PAGERDUTY"
				pager_duty_key = "pagerDutyKey"
			}
			linked_sso_display_name = "linked_sso_display_name"
			sso_group_id = "sso_group_id"
			sso_group_name = "sso_group_name"
			linked_sso_type = "SAML"
			sso_linked = true
		}
`, id, name)
}

func testProjectResourceUserGroup(id string, name string) string {
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

		resource "harness_platform_usergroup" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			linked_sso_id = "linked_sso_id"
			externally_managed = false
			users = []
			notification_configs {
				type = "SLACK"
				slack_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "EMAIL"
				group_email = "email@email.com"
				send_email_to_all_users = true
			}
			notification_configs {
				type = "MSTEAMS"
				microsoft_teams_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "PAGERDUTY"
				pager_duty_key = "pagerDutyKey"
			}
			linked_sso_display_name = "linked_sso_display_name"
			sso_group_id = "sso_group_id"
			sso_group_name = "sso_group_name"
			linked_sso_type = "SAML"
			sso_linked = true
		}
`, id, name)
}

func testOrgResourceUserGroup(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_usergroup" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			linked_sso_id = "linked_sso_id"
			externally_managed = false
			users = []
			notification_configs {
				type = "SLACK"
				slack_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "EMAIL"
				group_email = "email@email.com"
				send_email_to_all_users = true
			}
			notification_configs {
				type = "MSTEAMS"
				microsoft_teams_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "PAGERDUTY"
				pager_duty_key = "pagerDutyKey"
			}
			linked_sso_display_name = "linked_sso_display_name"
			sso_group_id = "sso_group_id"
			sso_group_name = "sso_group_name"
			linked_sso_type = "SAML"
			sso_linked = true
		}
`, id, name)
}

func testAccResourceUserGroup_emails(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_usergroup" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			linked_sso_id = "linked_sso_id"
			externally_managed = false
			user_emails = ["meet.meet@gmail.com", "richaterraformtest@gmail.com", "jajooricha@gmail.com", "aditya.kashyap@harness.io"]
			notification_configs {
				type = "SLACK"
				slack_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "EMAIL"
				group_email = "email@email.com"
				send_email_to_all_users = true
			}
			notification_configs {
				type = "MSTEAMS"
				microsoft_teams_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "PAGERDUTY"
				pager_duty_key = "pagerDutyKey"
			}
			linked_sso_display_name = "linked_sso_display_name"
			sso_group_id = "sso_group_id"
			sso_group_name = "sso_group_name"
			linked_sso_type = "SAML"
			sso_linked = true
		}
`, id, name)
}

func testProjectResourceUserGroup_emails(id string, name string) string {
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

		resource "harness_platform_usergroup" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			linked_sso_id = "linked_sso_id"
			externally_managed = false
			user_emails = ["rathod.meetsatish@harness.io"]
			notification_configs {
				type = "SLACK"
				slack_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "EMAIL"
				group_email = "email@email.com"
				send_email_to_all_users = true
			}
			notification_configs {
				type = "MSTEAMS"
				microsoft_teams_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "PAGERDUTY"
				pager_duty_key = "pagerDutyKey"
			}
			linked_sso_display_name = "linked_sso_display_name"
			sso_group_id = "sso_group_id"
			sso_group_name = "sso_group_name"
			linked_sso_type = "SAML"
			sso_linked = true
		}
`, id, name)
}

func testOrgResourceUserGroup_emails(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_usergroup" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			linked_sso_id = "linked_sso_id"
			externally_managed = false
			user_emails = ["rathod.meetsatish@harness.io"]
			notification_configs {
				type = "SLACK"
				slack_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "EMAIL"
				group_email = "email@email.com"
				send_email_to_all_users = true
			}
			notification_configs {
				type = "MSTEAMS"
				microsoft_teams_webhook_url = "https://google.com"
			}
			notification_configs {
				type = "PAGERDUTY"
				pager_duty_key = "pagerDutyKey"
			}
			linked_sso_display_name = "linked_sso_display_name"
			sso_group_id = "sso_group_id"
			sso_group_name = "sso_group_name"
			linked_sso_type = "SAML"
			sso_linked = true
		}
`, id, name)
}
