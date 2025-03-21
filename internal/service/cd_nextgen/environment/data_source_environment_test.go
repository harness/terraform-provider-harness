package environment_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceEnvironment(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_environment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEnvironment(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func TestAccDataSourceEnvironmentByName(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_environment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEnvironmentByName(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func TestAccDataSourceEnvironmentAtAccountLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_environment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEnvironmentAccountLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func TestAccDataSourceEnvironmentAtOrgLevel(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_environment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEnvironmentOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func TestDataSourceRemoteEnvironment(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_environment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSourceRemoteEnvironment(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testSourceRemoteEnvironment(id string, name string) string {
	return fmt.Sprintf(`
			resource "harness_platform_organization" "test" {
				identifier = "%[1]s"
				name = "%[2]s"
			}
			resource "harness_platform_project" "test" {
				identifier = "%[1]s"
				name = "%[2]s"
				org_id = harness_platform_organization.test.id
				color = "#472848"
			}
			resource "harness_platform_environment" "test" {
				identifier = "%[1]s"
				name = "%[2]s"
				org_id = harness_platform_project.test.org_id
				project_id = harness_platform_project.test.id
				type = "PreProduction"
				git_details {
					store_type = "REMOTE"
					connector_ref = "account.TF_GitX_connector"
					repo_name = "pcf_practice"
					file_path = ".harness/automation/%[1]s.yaml"
					branch = "main"
				}
			}
			data "harness_platform_environment" "test" {
				identifier = harness_platform_environment.test.identifier
				org_id = harness_platform_environment.test.org_id
				project_id = harness_platform_environment.test.project_id
				git_details {
					repo_name = "pcf_practice"
					branch = "main"
				}
			}
			`, id, name)
}

func testAccDataSourceEnvironment(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			color = "#0063F7"
			type = "PreProduction"
		}

		data "harness_platform_environment" "test" {
			identifier = harness_platform_environment.test.id
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_environment.test.project_id
		}
`, id, name)
}

func testAccDataSourceEnvironmentAccountLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			color = "#0063F7"
			type = "PreProduction"
		}

		data "harness_platform_environment" "test" {
			identifier = harness_platform_environment.test.id
		}
`, id, name)
}

func testAccDataSourceEnvironmentOrgLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			name = "%[2]s"
			color = "#0063F7"
			type = "PreProduction"
		}

		data "harness_platform_environment" "test" {
			identifier = harness_platform_environment.test.id
			org_id = harness_platform_environment.test.org_id
		}
`, id, name)
}

func testAccDataSourceEnvironmentByName(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			color = "#0063F7"
			type = "PreProduction"
		}

		data "harness_platform_environment" "test" {
			identifier = harness_platform_environment.test.id
			name = harness_platform_environment.test.name
			org_id = harness_platform_environment.test.org_id
			project_id = harness_platform_environment.test.project_id
		}
`, id, name)
}
