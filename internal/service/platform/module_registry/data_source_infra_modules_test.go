package module_registry_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceInfraModules(t *testing.T) {
	resourceName := "data.harness_platform_infra_modules.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceInfraModules(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "modules.#"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func TestAccDataSourceInfraModulesWithModules(t *testing.T) {
	dataSourceName := "data.harness_platform_infra_modules.test"
	resourceName := "harness_platform_infra_module.test"
	name := fmt.Sprintf("testmodule_%d", time.Now().Unix())
	system := "terraform"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceInfraModulesWithModule(name, system),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "modules.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "system", system),
				),
			},
		},
	})
}

func testAccDataSourceInfraModules() string {
	return `
		data "harness_platform_infra_modules" "test" {}
	`
}

func testAccDataSourceInfraModulesWithModule(name, system string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "secret%[1]s"
			name = "secret%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier = "github%[1]s"
			name = "github%[1]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://github.com/account"
			connection_type = "Account"
			validation_repo = "some_repo"
			delegate_selectors = ["harness-delegate"]
			credentials {
				http {
					username = "admin"
					token_ref = "account.${harness_platform_secret_text.test.id}"
				}
			}
		}

		resource "harness_platform_infra_module" "test" {
			name = "%[1]s"
			system = "%[2]s"
			description = "Test module for data source"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "tf/aws/basic"
			repository_connector = "account.${harness_platform_connector_github.test.id}"
		}

		data "harness_platform_infra_modules" "test" {
			depends_on = [harness_platform_infra_module.test]
		}
	`, name, system)
}
