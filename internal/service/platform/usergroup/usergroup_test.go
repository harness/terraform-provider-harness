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

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
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
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
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
				Config: testAccResourceUserGroup(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccProvider.Meta().(*internal.Session).GetPlatformClient()
					resp, _, err := c.UserGroupApi.DeleteUserGroup(ctx, c.AccountId, id, &nextgen.UserGroupApiDeleteUserGroupOpts{
						OrgIdentifier:     optional.NewString(id),
						ProjectIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)
					require.NotNil(t, resp)
					require.Equal(t, resp.Status, nextgen.ResponseStatusTypes.Success)
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
	c, ctx := acctest.TestAccGetApiClientFromProvider().GetPlatformClient()
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
		}
`, id, name)
}
