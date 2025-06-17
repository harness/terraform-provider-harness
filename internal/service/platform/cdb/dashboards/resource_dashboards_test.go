package dashboards_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceDashboard(t *testing.T) {

	description := "test_description"
	title := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	dashboardId := "48507" // DashBoard ID is present in QA : rXUXvbFqRr2XwcjBu3Oq-Q Account
	folderId := "shared"
	updatedTitle := fmt.Sprintf("%s_updated", title)
	resourceName := "harness_platform_dashboards.dashboard"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDashboardDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDashboard(dashboardId, description, folderId, title),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dashboard_id", dashboardId),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "resource_identifier", folderId),
					resource.TestCheckResourceAttr(resourceName, "title", title),
				),
			},
			{
				Config: testAccResourceDashboard(dashboardId, description, folderId, updatedTitle),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dashboard_id", dashboardId),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "resource_identifier", folderId),
					resource.TestCheckResourceAttr(resourceName, "title", updatedTitle),
				),
			},
		},
	})
}

func testAccDashboardDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		r, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return nil
		}

		c, ctx := acctest.TestAccGetPlatformClientWithContext()
		id := r.Primary.ID

		resp, httpResp, err := c.DashboardsApi.GetDashboard(ctx, id, &nextgen.DashboardsApiGetDashboardOpts{
			AccountId: optional.NewString(c.AccountId),
		})

		if err != nil && httpResp != nil && httpResp.StatusCode == 404 {
			return nil
		}

		if err == nil && resp.Resource != nil {
			return fmt.Errorf("Found dashboard: %s", id)
		}

		if err != nil {
			return err
		}

		return nil
	}
}

func testAccResourceDashboard(dashboard_id string, description string, folder_id string, title string) string {
	return fmt.Sprintf(`
	resource "harness_platform_dashboards" "dashboard" {
		dashboard_id = "%[1]s"
		description = "%[2]s"
		resource_identifier = "%[3]s"
		title = "%[4]s"
		data_source = []
		models = []
	}
	`, dashboard_id, description, folder_id, title)
}
