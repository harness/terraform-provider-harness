package user_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceUser(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceUser(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Rajendra Baviskar_updated"),
				),
			},
		},
	})
}

func testAccDataSourceUser(id string, name string) string {
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

	resource "harness_platform_user" "test" {
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		emails = ["rajendra.baviskar@harness.io"]
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
				name,
			]
		}
	}

	data "harness_platform_user" "test" {
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		email = "rajendra.baviskar@harness.io"
	}
	`, id, name)
}
