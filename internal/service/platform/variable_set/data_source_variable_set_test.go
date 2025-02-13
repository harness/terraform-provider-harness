package variable_set_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceVariableSetProjLevel(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "data.harness_platform_infra_variable_set.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVariableSetProjLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func TestAccDataSourceVariableSetOrgLevel(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "data.harness_platform_infra_variable_set.testorg"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVariableSetOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func TestAccDataSourceVariableSetAccLevel(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "data.harness_platform_infra_variable_set.testacc"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVariableSetAccLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func testAccDataSourceVariableSetProjLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}

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

		resource "harness_platform_infra_variable_set" "test" {
			identifier              = "%[1]s"
			name                    = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description             = "desc"

			environment_variable {
				key = "key1"
				value = "value1"
				value_type = "string"
			}

			terraform_variable {
				key = "key1"
				value = "1111"
				value_type = "string"
			}

			terraform_variable_file {
				repository              = "https://github.com/org/repo"
				repository_branch       = "main"
				repository_path         = "tf/aws/basic"
				repository_connector    = "account.${harness_platform_connector_github.test.id}"
			}     

			connector {
				connector_ref = "account.${harness_platform_connector_github.test.id}"
				type = "aws"
			}

		}		

		data "harness_platform_infra_variable_set" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			depends_on = [harness_platform_infra_variable_set.test]
		}

`, id, name)
}

func testAccDataSourceVariableSetOrgLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "testorg" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_secret_text" "testorg" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "testorg"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "testorg" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "testorg"
			tags = ["foo:bar"]

			url = "https://github.com/account"
			connection_type = "Account"
			validation_repo = "some_repo"
			delegate_selectors = ["harness-delegate"]
			credentials {
				http {
					username = "admin"
					token_ref = "account.${harness_platform_secret_text.testorg.id}"
				}
			}
		}

		resource "harness_platform_infra_variable_set" "testorg" {
			identifier              = "%[1]s"
			name                    = "%[2]s"
			org_id = harness_platform_organization.testorg.id
			description             = "desc"

			environment_variable {
				key = "key1"
				value = "value1"
				value_type = "string"
			}

			terraform_variable {
				key = "key1"
				value = "1111"
				value_type = "string"
			}

			terraform_variable_file {
				repository              = "https://github.com/org/repo"
				repository_branch       = "main"
				repository_path         = "tf/aws/basic"
				repository_connector    = "account.${harness_platform_connector_github.testorg.id}"
			}     

			connector {
				connector_ref = "account.${harness_platform_connector_github.testorg.id}"
				type = "aws"
			}

		}		

		data "harness_platform_infra_variable_set" "testorg" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.testorg.id
			depends_on = [harness_platform_infra_variable_set.testorg]
		}

`, id, name)
}

func testAccDataSourceVariableSetAccLevel(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "testacc" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "testacc"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "testacc" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "testacc"
			tags = ["foo:bar"]

			url = "https://github.com/account"
			connection_type = "Account"
			validation_repo = "some_repo"
			delegate_selectors = ["harness-delegate"]
			credentials {
				http {
					username = "admin"
					token_ref = "account.${harness_platform_secret_text.testacc.id}"
				}
			}
		}

		resource "harness_platform_infra_variable_set" "testacc" {
			identifier              = "%[1]s"
			name                    = "%[2]s"
			description             = "desc"

			environment_variable {
				key = "key1"
				value = "value1"
				value_type = "string"
			}

			terraform_variable {
				key = "key1"
				value = "1111"
				value_type = "string"
			}

			terraform_variable_file {
				repository              = "https://github.com/org/repo"
				repository_branch       = "main"
				repository_path         = "tf/aws/basic"
				repository_connector    = "account.${harness_platform_connector_github.testacc.id}"
			}     

			connector {
				connector_ref = "account.${harness_platform_connector_github.testacc.id}"
				type = "aws"
			}

		}		

		data "harness_platform_infra_variable_set" "testacc" {
			identifier = "%[1]s"
			depends_on = [harness_platform_infra_variable_set.testacc]
		}

`, id, name)
}
