package dashboards_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDashboards(t *testing.T) {

	dashboardId := "48507" // DashBoard ID is present in QA : rXUXvbFqRr2XwcjBu3Oq-Q Account
	description := "test_description"
	title := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	folderId := "shared"
	resourceName := "data.harness_platform_dashboards.dashboard"
	resourceNameCreated := "harness_platform_dashboards.dashboard"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDashboard(dashboardId, description, folderId, title),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "id", resourceNameCreated, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "description", resourceNameCreated, "description"),
					resource.TestCheckResourceAttrPair(resourceName, "resource_identifier", resourceNameCreated, "resource_identifier"),
					resource.TestCheckResourceAttrPair(resourceName, "title", resourceNameCreated, "title"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "type"),
					resource.TestCheckResourceAttrSet(resourceName, "view_count"),
				),
			},
		},
	})
}

func testAccDataSourceDashboard(dashboardId string, description string, folderId string, title string) string {
	return fmt.Sprintf(`
	resource "harness_platform_dashboards" "dashboard" {
		dashboard_id = "%[1]s"
		description = "%[2]s"
		resource_identifier = "%[3]s"
		title = "%[4]s"
		data_source = []
		models = []
	}

	data "harness_platform_dashboards" "dashboard" {
		id = harness_platform_dashboards.dashboard.id
		depends_on = [harness_platform_dashboards.dashboard]
	}
	`, dashboardId, description, folderId, title)
}
