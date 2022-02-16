package user_test

import (
	"fmt"
	"regexp"
	"testing"

	sdk "github.com/harness/harness-go-sdk"
	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceUserGroupPermissions_AccountPermissions(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_user_group_permissions.test"

	defer func() {
		c := acctest.TestAccGetApiClientFromProvider()
		ug, err := c.CDClient.UserClient.GetUserGroupByName(expectedName)
		require.NoError(t, err)
		require.NotNil(t, ug)
		c.CDClient.UserClient.DeleteUserGroup(ug.Id)
	}()

	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			c := acctest.TestAccGetApiClientFromProvider()
			c.CDClient.UserClient.CreateUserGroup(&graphql.UserGroup{
				Name: expectedName,
			})
		},
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupPermissionsDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroupPermissions_AccountPermissions(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "account_permissions.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					return primary.Attributes["user_group_id"], nil
				},
			},
		},
	})
}

func TestAccResourceUserGroupPermissions_DeleteUnderlyingResource(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))

	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			c := acctest.TestAccGetApiClientFromProvider()
			c.CDClient.UserClient.CreateUserGroup(&graphql.UserGroup{
				Name: expectedName,
			})
		},
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroupPermissions_AccountPermissions(expectedName),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c := acctest.TestAccProvider.Meta().(*sdk.Session)

					grp, err := c.CDClient.UserClient.GetUserGroupByName(expectedName)
					require.NoError(t, err)
					require.NotNil(t, grp)

					err = c.CDClient.UserClient.DeleteUserGroup(grp.Id)
					require.NoError(t, err)
				},
				Config:   testAccResourceUserGroupAccountPermissions(expectedName),
				PlanOnly: true,
				// ExpectNonEmptyPlan: true,
				ExpectError: regexp.MustCompile("user group .* does not exist"),
			},
		},
	})
}

func TestAccResourceUserGroupPermissions_AppPermissions(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_user_group_permissions.test"

	defer func() {
		c := acctest.TestAccGetApiClientFromProvider()
		ug, err := c.CDClient.UserClient.GetUserGroupByName(expectedName)
		require.NoError(t, err)
		require.NotNil(t, ug)
		c.CDClient.UserClient.DeleteUserGroup(ug.Id)
	}()

	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			c := acctest.TestAccGetApiClientFromProvider()
			c.CDClient.UserClient.CreateUserGroup(&graphql.UserGroup{
				Name: expectedName,
			})
		},
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserGroupPermissionsDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserGroupPermissionsAppPermissions(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "app_permissions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "app_permissions.0.all.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "app_permissions.0.deployment.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "app_permissions.0.environment.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "app_permissions.0.pipeline.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "app_permissions.0.provisioner.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "app_permissions.0.service.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "app_permissions.0.template.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "app_permissions.0.workflow.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					return primary.Attributes["user_group_id"], nil
				},
			},
		},
	})
}

func testAccResourceUserGroupPermissionsAppPermissions(name string) string {
	return fmt.Sprintf(`
		data "harness_user_group" "test" {
			name = "%s"
		}

		resource "harness_user_group_permissions" "test" {
			user_group_id = data.harness_user_group.test.id

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
					actions = ["CREATE", "READ"]
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
`, name)
}

func testAccResourceUserGroupPermissions_AccountPermissions(name string) string {
	return fmt.Sprintf(`
		data "harness_user_group" "test" {
			name = "%s"
		}
		
		resource "harness_user_group_permissions" "test" {
			user_group_id = data.harness_user_group.test.id

			account_permissions = ["ADMINISTER_OTHER_ACCOUNT_FUNCTIONS", "MANAGE_API_KEYS"]
		}
`, name)
}

func testAccUserGroupPermissionsDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		r := acctest.TestAccGetResource(resourceName, state)
		c := acctest.TestAccGetApiClientFromProvider()

		id := r.Primary.ID

		ug, err := c.CDClient.UserClient.GetUserGroupById(id)
		if err != nil {
			return err
		}

		if len(ug.Permissions.AccountPermissions.AccountPermissionTypes) == 0 && len(ug.Permissions.AppPermissions) == 0 {
			return nil
		}

		return fmt.Errorf("User group permissions not destroyed: %s", ug.Id)
	}
}
