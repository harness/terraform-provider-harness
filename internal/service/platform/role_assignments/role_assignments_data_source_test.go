package role_assignments_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceRoleAssignments(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	resourceName := "data.harness_platform_role_assignments.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRoleAssignments(id, name, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "resource_group_identifier", "_all_account_level_resources"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "managed", "false"),
					resource.TestCheckResourceAttr(resourceName, "principal.0.type", "SERVICE_ACCOUNT"),
				),
			},
		},
	})
}

func testAccDataSourceRoleAssignments(id string, name string, accountId string) string {
	return fmt.Sprintf(`
	resource "harness_platform_service_account" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		email = "email@service.harness.io"
		description = "test"
		tags = ["foo:bar"]
		account_id = "%[3]s"
	}

	resource "harness_platform_role_assignments" "test" {
		identifier = "%[1]s"
		resource_group_identifier = "_all_account_level_resources"
		role_identifier = "_account_viewer"
		principal {
			identifier = harness_platform_service_account.test.id
			type = "SERVICE_ACCOUNT"
		}
		disabled = false
		managed = false
	}

	data "harness_platform_role_assignments" "test" {
		identifier = harness_platform_role_assignments.test.id
	}
	`, id, name, accountId)
}
