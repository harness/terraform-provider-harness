package folders_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceDashboardFolder(t *testing.T) {

	name := t.Name()
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_dashboard_folders.dashboard"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFolderDestroy(resourceName),
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
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccGetFolder(resourceName string, state *terraform.State) (*nextgen.Folder, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	accId := r.Primary.Attributes["account_id"]

	resp, _, err := c.DashboardsFolderApi.GetFolder(ctx, id, &nextgen.DashboardsFoldersApiGetFolderOpts{AccountId: optional.NewString(accId)})
	if err != nil {
		return nil, err
	}

	return resp.Resource, nil
}

func testAccFolderDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		folder, _ := testAccGetFolder(resourceName, state)
		if folder != nil {
			return fmt.Errorf("Found folder: %s", folder.Id)
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
