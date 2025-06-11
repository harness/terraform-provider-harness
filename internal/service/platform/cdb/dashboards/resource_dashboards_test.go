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
	title := t.Name()
	dashboard_id := fmt.Sprintf("%s_%s", title, utils.RandStringBytes(5))
	folder_id := fmt.Sprintf("%s_%s", title, utils.RandStringBytes(5))
	updatedTitle := fmt.Sprintf("%s_updated", title)
	resourceName := "harness_platform_dashboards.dashboard"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDashboardDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDashboard(dashboard_id, description, folder_id, title),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dashboard_id", dashboard_id),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "folder_id", folder_id),
					resource.TestCheckResourceAttr(resourceName, "title", title),
				),
			},
			{
				Config: testAccResourceDashboard(dashboard_id, description, folder_id, updatedTitle),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dashboard_id", dashboard_id),
					resource.TestCheckResourceAttr(resourceName, "description", description),
					resource.TestCheckResourceAttr(resourceName, "folder_id", folder_id),
					resource.TestCheckResourceAttr(resourceName, "title", updatedTitle),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetDashboard(resourceName string, state *terraform.State) (*nextgen.Dashboard, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	accId := r.Primary.Attributes["account_id"]

	resp, _, err := c.DashboardsApi.GetDashboard(ctx, id, &nextgen.DashboardsApiGetDashboardOpts{AccountId: optional.NewString(accId)})
	if err != nil {
		return nil, err
	}

	return resp.Resource, nil
}

func testAccDashboardDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		dashboard, _ := testAccGetDashboard(resourceName, state)
		if dashboard != nil {
			return fmt.Errorf("Found dashboard: %s", dashboard.Id)
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
	}
	`, dashboard_id, description, folder_id, title)
}
