package folders_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDashboardFolders(t *testing.T) {

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "data.harness_platform_dashboard_folders.folder"
	resourceNameCreated := "harness_platform_dashboard_folders.folder"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFolder(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "id", resourceNameCreated, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "name", resourceNameCreated, "name"),
				),
			},
		},
	})
}

func testAccDataSourceFolder(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_dashboard_folders" "folder" {
		name = "%[1]s"
	}

	data "harness_platform_dashboard_folders" "folder" {
		id = harness_platform_dashboard_folders.folder.id
		depends_on = [harness_platform_dashboard_folders.folder]
	}
	`, name)
}
