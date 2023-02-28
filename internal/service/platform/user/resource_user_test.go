package user_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceUserProjectLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	email := strings.ToLower(id) + "@harness.io"
	resourceName := "harness_platform_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserProjectLevel(id, name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"role_bindings", "user_groups"},
				ImportStateIdFunc:       acctest.UserResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceUserAccountLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	email := strings.ToLower(id) + "@harness.io"
	resourceName := "harness_platform_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserAccountLevel(id, name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "email", email),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"role_bindings", "user_groups"},
				ImportStateIdFunc:       acctest.UserResourceImportStateIdFuncAccountLevel(resourceName),
			},
		},
	})
}

func TestAccResourceUserOrgLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	email := strings.ToLower(id) + "@harness.io"
	resourceName := "harness_platform_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccUserDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserOrgLevel(id, name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "email", email),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"role_bindings", "user_groups"},
				ImportStateIdFunc:       acctest.UserResourceImportStateIdFuncOrgLevel(resourceName),
			},
		},
	})
}

func TestAccResourceUser_DeleteUnderlyingResourceProjectLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	email := strings.ToLower(id) + "@harness.io"
	resourceName := "harness_platform_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserProjectLevel(id, name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					_, _, err := c.UserApi.RemoveUser(ctx, email, c.AccountId, &nextgen.UserApiRemoveUserOpts{
						OrgIdentifier:     optional.NewString(id),
						ProjectIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)
				},
				Config:   testAccResourceUserProjectLevel(id, name, email),
				PlanOnly: true,
			},
		},
	})
}

func TestAccResourceUser_DeleteUnderlyingResourceAccountLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	email := strings.ToLower(id) + "@harness.io"
	resourceName := "harness_platform_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserAccountLevel(id, name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "email", email),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					_, _, err := c.UserApi.RemoveUser(ctx, email, c.AccountId, &nextgen.UserApiRemoveUserOpts{})
					require.NoError(t, err)
				},
				Config:   testAccResourceUserAccountLevel(id, name, email),
				PlanOnly: true,
			},
		},
	})
}

func TestAccResourceUser_DeleteUnderlyingResourceOrgLevel(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	email := strings.ToLower(id) + "@harness.io"
	resourceName := "harness_platform_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserOrgLevel(id, name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					_, _, err := c.UserApi.RemoveUser(ctx, email, c.AccountId, &nextgen.UserApiRemoveUserOpts{
						OrgIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)
				},
				Config:   testAccResourceUserOrgLevel(id, name, email),
				PlanOnly: true,
			},
		},
	})
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccGetPlatformUser(resourceName string, state *terraform.State) (*nextgen.UserAggregate, error) {

	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	email := r.Primary.Attributes["email"]

	resp, _, err := c.UserApi.GetAggregatedUsers(ctx, c.AccountId, &nextgen.UserApiGetAggregatedUsersOpts{
		OrgIdentifier:     buildField(r, "org_id"),
		ProjectIdentifier: buildField(r, "project_id"),
		SearchTerm:        optional.NewString(email),
	})

	if err != nil || resp.Data.Empty {
		return nil, err
	}

	return &resp.Data.Content[0], nil
}

func testAccUserDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		user, _ := testAccGetPlatformUser(resourceName, state)
		if user != nil {
			return fmt.Errorf("Found user: %s", user.User.Uuid)
		}

		return nil
	}
}

func testAccResourceUserProjectLevel(id string, name string, email string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#0063F7"
		}

		resource "harness_platform_user" "test" {
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			email = "%[3]s"
			user_groups = ["_project_all_users"]
			role_bindings {
				resource_group_identifier = "_all_project_level_resources"
				role_identifier = "_project_viewer"
				role_name = "Project Viewer"
				resource_group_name = "All Project Level Resources"
				managed_role = true
			}
		}
`, id, name, email)
}

func testAccResourceUserAccountLevel(id string, name string, email string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#0063F7"
		}

		resource "harness_platform_user" "test" {
			email = "%[3]s"
			user_groups = ["_project_all_users"]
			role_bindings {
				resource_group_identifier = "_all_project_level_resources"
				role_identifier = "_project_viewer"
				role_name = "Project Viewer"
				resource_group_name = "All Project Level Resources"
				managed_role = true
			}
			lifecycle {
				ignore_changes = [
					org_id, project_id
				]
			}
		}
`, id, name, email)
}

func testAccResourceUserOrgLevel(id string, name string, email string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#0063F7"
		}

		resource "harness_platform_user" "test" {
			org_id = harness_platform_project.test.org_id
			email = "%[3]s"
			user_groups = ["_project_all_users"]
			role_bindings {
				resource_group_identifier = "_all_project_level_resources"
				role_identifier = "_project_viewer"
				role_name = "Project Viewer"
				resource_group_name = "All Project Level Resources"
				managed_role = true
			}
			lifecycle {
				ignore_changes = [
					org_id, project_id
				]
			}
		}
`, id, name, email)
}
