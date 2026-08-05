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
	resourceName := "data.harness_platform_dashboard_folder.folder"
	resourceNameCreated := "harness_platform_dashboard_folder.folder"

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
	resource "harness_platform_dashboard_folder" "folder" {
		name = "%[1]s"
	}

	data "harness_platform_dashboard_folder" "folder" {
		id = harness_platform_dashboard_folder.folder.id
		depends_on = [harness_platform_dashboard_folder.folder]
	}
	`, name)
}

func TestAccDataSourceDashboardFolders_List(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	listDataSourceName := "data.harness_platform_dashboard_folders.all"
	createdResource := "harness_platform_dashboard_folder.folder"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFolderList(name),
				Check: resource.ComposeTestCheckFunc(
					// The list should contain at least the created folder (count >=1)
					resource.TestCheckResourceAttrWith(listDataSourceName, "folders.#", func(value string) error {
						if value == "0" {
							return fmt.Errorf("expected at least one folder, got 0")
						}
						return nil
					}),
					// Also verify singular by name works
					resource.TestCheckResourceAttrPair("data.harness_platform_dashboard_folder.by_name", "id", createdResource, "id"),
					resource.TestCheckResourceAttrPair("data.harness_platform_dashboard_folder.by_name", "name", createdResource, "name"),
				),
			},
		},
	})
}

func testAccDataSourceFolderList(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_dashboard_folder" "folder" {
		name = "%[1]s"
	}

	data "harness_platform_dashboard_folders" "all" {
		depends_on = [harness_platform_dashboard_folder.folder]
	}

	data "harness_platform_dashboard_folder" "by_name" {
		name       = "%[1]s"
		depends_on = [harness_platform_dashboard_folder.folder]
	}
	`, name)
}
