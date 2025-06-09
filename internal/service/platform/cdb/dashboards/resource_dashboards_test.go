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

	title := t.Name()
	id := fmt.Sprintf("%s_%s", title, utils.RandStringBytes(5))
	updatedTitle := fmt.Sprintf("%s_updated", title)
	resourceName := "harness_platform_dashboard.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDashboardDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDashboard(id, title),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dashboard_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", title),
				),
			},
			{
				Config: testAccResourceDashboard(id, updatedTitle),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "dashboard_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedTitle),
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

func testAccResourceDashboard(id string, title string) string {
	return fmt.Sprintf(`
		resource "harness_platform_dashboard" "test" {
			dashboard_id = "%[1]s"
			name = "%[2]s"
		}
`, id, title)
}
