package secret_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSecretFile(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_secret_file.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSecret_file(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
				)},
		},
	})
}
func TestAccDataSourceSecretFileProjectLevel(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_secret_file.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSecret_fileProjectLevel(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
					resource.TestCheckResourceAttr(resourceName, "org_id", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", name),
				)},
		},
	})
}
func TestAccDataSourceSecretFileOrgLevel(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_secret_file.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSecret_fileOrgLevel(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_identifier", "harnessSecretManager"),
				)},
		},
	})
}

func testAccDataSourceSecret_file(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		file_path = "%[2]s"
		secret_manager_identifier = "harnessSecretManager"
	}
	data "harness_platform_secret_file" "test"{
		identifier = harness_platform_secret_file.test.identifier
		
	}
	`, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccDataSourceSecret_fileProjectLevel(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		color = "#0063F7"
		org_id = harness_platform_organization.test.id
	}
	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		tags = ["foo:bar"]
		file_path = "%[2]s"
		secret_manager_identifier = "harnessSecretManager"
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_project.test]
		create_duration = "4s" 
	}

	data "harness_platform_secret_file" "test"{
		identifier = harness_platform_secret_file.test.identifier
		
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
	`, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}

func testAccDataSourceSecret_fileOrgLevel(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_secret_file" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		file_path = "%[2]s"
		org_id = harness_platform_organization.test.id
		secret_manager_identifier = "harnessSecretManager"
		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_organization.test]
		create_duration = "4s" 
	}

	data "harness_platform_secret_file" "test"{
		identifier = harness_platform_secret_file.test.identifier
		
		org_id = harness_platform_organization.test.id
	}
	`, name, getAbsFilePath("../../../acctest/secret_files/secret.txt"))
}
