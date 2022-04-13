package user_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func init() {
	resource.AddTestSweepers("harness_user_group", &resource.Sweeper{
		Name:         "harness_user_group",
		F:            testSweepUserGroups,
		Dependencies: []string{"harness_users"},
	})
}

func TestAccResourceUserGroup_LDAP(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_user_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroupLDAP(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccUserGroupCreation(t, resourceName, expectedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceUserGroup_SAML(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_user_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroupSAML(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccUserGroupCreation(t, resourceName, expectedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceUserGroup_NotificationsSettings(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_user_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroupNotificationSettings(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccUserGroupCreation(t, resourceName, expectedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceUserGroup_AccountPermissions(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_user_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroupAccountPermissions(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccUserGroupCreation(t, resourceName, expectedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceUserGroup_DeleteUnderlyingResource(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_user_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroupAccountPermissions(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*internal.Session).CDClient

					grp, err := c.UserClient.GetUserGroupByName(expectedName)
					require.NoError(t, err)
					require.NotNil(t, grp)

					err = c.UserClient.DeleteUserGroup(grp.Id)
					require.NoError(t, err)
				},
				Config:             testAccResourceUserGroupAccountPermissions(expectedName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccResourceUserGroup_AppPermissions(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_user_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroupAppPermissions(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccUserGroupCreation(t, resourceName, expectedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccUserGroupCreation(t *testing.T, resourceName string, name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		group, err := testAccGetUserGroup(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, group)
		require.Equal(t, name, group.Name)

		return nil
	}
}

func testAccGetUserGroup(resourceName string, state *terraform.State) (*graphql.UserGroup, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c := acctest.TestAccGetApiClientFromProvider().CDClient
	id := r.Primary.ID

	return c.UserClient.GetUserGroupById(id)
}

func testAccUserGroupDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		app, _ := testAccGetUserGroup(resourceName, state)
		if app != nil {
			return fmt.Errorf("Found user group: %s", app.Id)
		}

		return nil
	}
}

func testAccResourceUserGroupAppPermissions(name string) string {
	return fmt.Sprintf(`
		resource "harness_user_group" "test" {
			name = "%s"
			description = "my description"

			permissions {
				app_permissions {
					
					all {
						app_ids = []
						actions = ["CREATE", "READ", "UPDATE", "DELETE"]
					}

					deployment {
						actions = ["READ", "ROLLBACK_WORKFLOW", "EXECUTE_PIPELINE", "EXECUTE_WORKFLOW"]
						filters = ["NON_PRODUCTION_ENVIRONMENTS"]
					}	

					deployment {
						actions = ["READ"]
						filters = ["PRODUCTION_ENVIRONMENTS"]
					}	

					environment {
						actions = ["CREATE", "READ", "UPDATE", "DELETE"]
						filters = ["NON_PRODUCTION_ENVIRONMENTS"]
					}

					environment {
						actions = ["READ"]
						filters = ["PRODUCTION_ENVIRONMENTS"]
					}

					pipeline {
						actions = ["CREATE", "READ", "UPDATE", "DELETE"]
						filters = ["NON_PRODUCTION_PIPELINES"]
					}

					pipeline {
						actions = ["READ"]
						filters = ["PRODUCTION_PIPELINES"]
					}

					provisioner {
						actions = ["UPDATE", "DELETE"]
					}

					provisioner {
						actions = ["CREATE", "READ"]
					}

					service {
						actions = ["UPDATE", "DELETE"]
					}

					service {
						actions = ["UPDATE", "DELETE"]
					}
					
					template {
						actions = ["CREATE", "READ", "UPDATE", "DELETE"]
					}

					workflow {
						actions = ["UPDATE", "DELETE"]
						filters = ["NON_PRODUCTION_WORKFLOWS",]
					}

					workflow {
						actions = ["CREATE", "READ"]
						filters = ["PRODUCTION_WORKFLOWS", "WORKFLOW_TEMPLATES"]
					}

				}
			}
		}
`, name)
}

func testAccResourceUserGroupAccountPermissions(name string) string {
	return fmt.Sprintf(`
		resource "harness_user_group" "test" {
			name = "%s"
			description = "my description"

			permissions {
				account_permissions = ["VIEW_CE", "ADMINISTER_OTHER_ACCOUNT_FUNCTIONS", "MANAGE_API_KEYS"]
			}
		}
`, name)
}

func testAccResourceUserGroupNotificationSettings(name string) string {
	return fmt.Sprintf(`
		resource "harness_user_group" "test" {
			name = "%s"
			description = "my description"

			notification_settings {
				group_email_addresses = ["test@example.com", "foo@example.com"]
				microsoft_teams_webhook_url = "https://notifications.example.com"
				send_mail_to_new_members = true
				send_notifications_to_members = true
				slack_channel = "test"
				slack_webhook_url = "https://slack.webhooks.example.com"
			}
		}
`, name)
}

func testAccResourceUserGroupLDAP(name string) string {
	return fmt.Sprintf(`
	  data "harness_sso_provider" "ldap" {
			name = "ldap-test"
		}

		resource "harness_user_group" "test" {
			name = "%s"
			description = "my description"

			ldap_settings {
				group_dn = "groupdn"
				group_name = "group name"
				sso_provider_id = data.harness_sso_provider.ldap.id
			}
		}
`, name)
}

func testAccResourceUserGroupSAML(name string) string {
	return fmt.Sprintf(`
	  data "harness_sso_provider" "saml" {
			name = "saml-test"
		}

		resource "harness_user_group" "test" {
			name = "%s"
			description = "my description"

			saml_settings {
				group_name = "group name"
				sso_provider_id = data.harness_sso_provider.saml.id
			}
		}
`, name)
}

func testSweepUserGroups(r string) error {
	c := acctest.TestAccGetApiClientFromProvider().CDClient

	limit := 100
	offset := 0
	hasMore := true

	for hasMore {

		groups, _, err := c.UserClient.ListUserGroups(limit, offset)
		if err != nil {
			return err
		}

		for _, group := range groups {
			// Only delete user groups that start with 'Test'
			if strings.HasPrefix(group.Name, "Test") {
				if err = c.UserClient.DeleteUserGroup(group.Id); err != nil {
					return err
				}
			}
		}

		hasMore = len(groups) == limit
		offset += limit
	}

	return nil
}
