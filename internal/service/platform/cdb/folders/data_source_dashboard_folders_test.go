package folders_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceDashboardFolders(t *testing.T) {

	name := "test_folder_name"
	resourceName := "data.harness_platform_dashboard_folders.folder"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFolder(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
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
	`, name)
}
