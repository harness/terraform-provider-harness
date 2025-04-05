package filters_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitOpsFilters(t *testing.T) {
	id := strings.ToLower(fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(5)))
	id = strings.ReplaceAll(id, "_", "")
	name := id
	resourceName := "data.harness_platform_gitops_filters.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGitOpsFilters(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", "APPLICATION"),
					resource.TestCheckResourceAttr(resourceName, "filter_visibility", "OnlyCreator"),
				),
			},
		},
	})
}

func testAccDataSourceGitOpsFilters(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_gitops_filters" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			type = "APPLICATION"
			filter_properties = jsonencode({
				"healthStatus": ["Suspended"],
				"syncStatus": ["Synced"]
			})
			filter_visibility = "OnlyCreator"
		}

		data "harness_platform_gitops_filters" "test" {
			identifier = harness_platform_gitops_filters.test.identifier
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			type = "APPLICATION"
			depends_on = [harness_platform_gitops_filters.test]
		}
	`, id, name)
}
