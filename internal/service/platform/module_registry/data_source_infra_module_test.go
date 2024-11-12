package module_registry_test

import (
	"fmt"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strings"
	"testing"
)

func TestAccDataSourceInfraModule(t *testing.T) {
	name := strings.ToLower(t.Name())
	system := strings.ToLower(t.Name())
	resourceName := "harness_platform_infra_module.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceInfraModule(name, system),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "system", system),
				),
			},
		},
	})
}

func testAccDataSourceInfraModule(name string, system string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}
		resource "harness_platform_connector_github" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
			description = "description"
			repository              = "https://github.com/org/repo"
			repository_branch       = "main"
			repository_path         = "tf/aws/basic"
			repository_connector    = "account.${harness_platform_connector_github.test.id}"
  		}	
	`, name, system)
}
