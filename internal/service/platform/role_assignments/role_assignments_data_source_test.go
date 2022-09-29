package role_assignments_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceRoleAssignments(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_role_assignments.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRoleAssignments(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "resource_group_identifier", "_all_project_level_resources"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "managed", "false"),
					resource.TestCheckResourceAttr(resourceName, "principal.0.type", "SERVICE_ACCOUNT"),
				),
			},
		},
	})
}

func testAccDataSourceRoleAssignments(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}

	resource "harness_platform_service_account" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		email = "email@service.harness.io"
		description = "test"
		tags = ["foo:bar"]
		account_id = "UKh5Yts7THSMAbccG3HrLA"
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
	}

	resource "harness_platform_role_assignments" "test1"{
		identifier = "%[1]s"
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		resource_group_identifier = "_all_project_level_resources"
		role_identifier = "_project_viewer"
		principal {
			identifier = harness_platform_service_account.test.id
			type = "SERVICE_ACCOUNT"
		}
		disabled = false
		managed = false
	}

	data "harness_platform_role_assignments" "test" {
		identifier = "%[1]s"
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
	}
	`, id, name)
}
