package environment_group_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceEnvironmentGroup(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_environment_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccEnvironmentGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironmentGroup(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "color", "#0063F7"),
				),
			},
		},
	})
}

func TestAccResourceEnvironmentGRoup_DeleteUnderlyingResource(t *testing.T) {
	t.Skip()
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_environment_group.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceEnvironmentGroup(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetPlatformClientWithContext()

					OrgIdentifier := id
					ProjectIdentifier := id

					resp, _, err := c.EnvironmentGroupApi.DeleteEnvironmentGroup(ctx, id, c.AccountId, OrgIdentifier, ProjectIdentifier, &nextgen.EnvironmentGroupApiDeleteEnvironmentGroupOpts{
						Branch:         optional.NewString(id),
						RepoIdentifier: optional.NewString(id),
					})
					require.NoError(t, err)
					require.True(t, resp.Data.Deleted)
				},
				Config:             testAccResourceEnvironmentGroup(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGetPlatformEnvironmentGroup(resourceName string, state *terraform.State) (*nextgen.EnvironmentGroupResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	branch := r.Primary.Attributes["branch"]
	repoIdentifier := r.Primary.Attributes["repoIdentifier"]
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]

	resp, _, err := c.EnvironmentGroupApi.GetEnvironmentGroup((ctx), id, c.AccountId, orgId, projId, &nextgen.EnvironmentGroupApiGetEnvironmentGroupOpts{
		Branch:         optional.NewString(branch),
		RepoIdentifier: optional.NewString(repoIdentifier),
	})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil || resp.Data.EnvGroup == nil {
		return nil, nil
	}

	return resp.Data.EnvGroup, nil
}

func testAccEnvironmentGroupDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		env, _ := testAccGetPlatformEnvironmentGroup(resourceName, state)
		if env != nil {
			return fmt.Errorf("Found environment group: %s", env.Identifier)
		}

		return nil
	}
}

func testAccResourceEnvironmentGroup(id string, name string) string {
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

		resource "harness_platform_environment_group" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			color = "#0063F7"
			yaml = <<-EOT
			     environmentGroup:
			                 name: "%[1]s"
			                 identifier: "%[1]s"
			                 description: "temp"
			                 orgIdentifier: ${harness_platform_project.test.org_id}
			                 projectIdentifier: ${harness_platform_project.test.id}
			                 envIdentifiers: []
		  EOT
		}
`, id, name)
}
