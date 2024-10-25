package file_store_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceFileStoreFile(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_file_store_file.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFileStoreDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFileStore_File(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test file"),
					resource.TestCheckResourceAttr(resourceName, "content", "file content"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceFileStoreFileOrgLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_file_store_file.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccFileStoreDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFileStore_FileOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test file"),
					resource.TestCheckResourceAttr(resourceName, "content", "file content"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
		},
	})
}

func TestAccDataSourceFileStoreFileProjectLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	resourceName := "harness_platform_file_store_file.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccFileStoreDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFileStore_FileProjectLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test file"),
					resource.TestCheckResourceAttr(resourceName, "content", "file content"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
		},
	})
}

func testAccDataSourceFileStore_FileProjectLevel(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.identifier
	}

	resource "harness_platform_file_store_file" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = "%[1]s"
		project_id  = "%[1]s"
		description = "test file"
		tags = ["foo:bar", "bar:foo"]
		parent_identifier = "Root"
		mime_type = "text"
		file_usage = "SCRIPT"
		file_content_path = "%[3]s"
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_project.test]
		create_duration = "4s"
	}
		`, id, name, getAbsFilePath("../../../acctest/file_store_files/file.txt"))
}

func testAccDataSourceFileStore_FileOrgLevel(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_file_store_file" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = "%[1]s"
		description = "test file"
		tags = ["foo:bar"]
		parent_identifier = "Root"
		mime_type = "text"
		file_usage = "SCRIPT"
		file_content_path = "%[3]s"
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_organization.test]
		create_duration = "4s"
	}
		`, id, name, getAbsFilePath("../../../acctest/file_store_files/file.txt"))
}

func testAccDataSourceFileStore_File(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_file_store_file" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test file"
		tags = ["foo:bar", "bar:foo"]
		parent_identifier = "Root"
		mime_type = "text"
		file_usage = "SCRIPT"
		file_content_path =  "%[3]s"
	}
		`, id, name, getAbsFilePath("../../../acctest/file_store_files/file.txt"))
}

// common methods for file and folder
func getAbsFilePath(file_path string) string {
	absPath, _ := filepath.Abs(file_path)
	return absPath
}
