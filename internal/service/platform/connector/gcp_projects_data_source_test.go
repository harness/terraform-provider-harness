package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGcpProjects(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_gcp_projects.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGcpProjects(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "connector_id", name),
					resource.TestCheckResourceAttrSet(resourceName, "projects.#"),
				),
			},
		},
	})
}

func testAccDataSourceGcpProjects(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_gcp" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			inherit_from_delegate {
				delegate_selectors = ["harness-delegate"]
			}
		}

		data "harness_platform_gcp_projects" "test" {
			connector_id = harness_platform_connector_gcp.test.identifier
		}
	`, name)
}
