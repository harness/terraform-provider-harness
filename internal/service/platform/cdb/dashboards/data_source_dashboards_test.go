package dashboards_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDashboards(t *testing.T) {

	dashboard_id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	description := "test_description"
	folder_id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	title := "test_title"
	resourceName := "data.harness_platform_dashboards.dashboard"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDashboard(dashboard_id, description, folder_id, title),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dashboard_id", dashboard_id),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "resource_identifier", folder_id),
					resource.TestCheckResourceAttr(resourceName, "title", title),
				),
			},
		},
	})
}

func testAccDataSourceDashboard(dashboard_id string, description string, folder_id string, title string) string {
	return fmt.Sprintf(`
	resource "harness_platform_dashboards" "dashboard" {
		dashboard_id = "%[1]s"
		description = "%[2]s"
		resource_identifier = "%[3]s"
		title = "%[4]s"
	}
	`, dashboard_id, description, folder_id, title)
}
