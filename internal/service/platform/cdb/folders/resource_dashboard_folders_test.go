package folders_test

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

func TestAccResourceDashboardFolder(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_dashboard_folders.folder"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccDashboardFolderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceDashboardFolder(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceDashboardFolder(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
		},
	})
}

func testAccDashboardFolderDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		r, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return nil
		}

		c, ctx := acctest.TestAccGetPlatformClientWithContext()
		id := r.Primary.ID

		resp, httpResp, err := c.DashboardsFolderApi.GetFolder(ctx, id, &nextgen.DashboardsFoldersApiGetFolderOpts{
			AccountId: optional.NewString(c.AccountId),
		})

		if err != nil && httpResp != nil && httpResp.StatusCode == 404 {
			return nil
		}

		if err == nil && resp.Resource != nil {
			return fmt.Errorf("Found folder: %s", id)
		}

		if err != nil {
			return err
		}

		return nil
	}
}

func testAccResourceDashboardFolder(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_dashboard_folders" "folder" {
		name = "%[1]s"
	}
	`, name)
}
