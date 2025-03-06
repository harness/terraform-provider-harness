package variable_set_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceVariableSetProjLevel(t *testing.T) {
	resourceName := "harness_platform_infra_variable_set.test"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceVariableSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceVariableSetProjLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceVariableSetProjLevel(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceVariableSetOrgLevel(t *testing.T) {
	resourceName := "harness_platform_infra_variable_set.testorg"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceVariableSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceVariableSetOrgLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceVariableSetOrgLevel(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceVariableSetAccLevel(t *testing.T) {
	resourceName := "harness_platform_infra_variable_set.testacc"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceVariableSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceVariableSetAccLevel(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceVariableSetAccLevel(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceVariableSetDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		ws, _ := testAccGetPlatformVariableSet(resourceName, state)
		if ws != nil {
			return fmt.Errorf("VariableSet found: %s", ws.Identifier)
		}
		return nil
	}
}

func testAccGetPlatformVariableSet(resourceName string, state *terraform.State) (*nextgen.VariableSetsGetVariableSetAccountLevelResponseBody, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	org := r.Primary.Attributes["org_id"]
	project := r.Primary.Attributes["project_id"]

	var vs nextgen.VariableSetsGetVariableSetAccountLevelResponseBody
	var resp *http.Response
	var err error

	if org == "" {
		vs, resp, err = c.VariableSetsApi.VariableSetsGetVariableSetAccountLevel(ctx, id, c.AccountId)
	} else if project == "" {
		vs, resp, err = c.VariableSetsApi.VariableSetsGetVariableSetOrgLevel(ctx, org, id, c.AccountId)
	} else {
		vs, resp, err = c.VariableSetsApi.VariableSetsGetVariableSetProjLevel(ctx, org, project, id, c.AccountId)
	}

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return &vs, nil
}

func testAccResourceVariableSetProjLevel(id string, name string) string {
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
			environment_variable {
				key = "key2"
				value = "value2up"
				value_type = "secret"
			}

			terraform_variable {
				key = "key1"
				value = "1111"
				value_type = "string"
			}
			terraform_variable {
				key = "key2"
				value = "1111u"
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
			connector {
				connector_ref = "azureCon"
				type = "azure"
			}

		}		

`, id, name)
}

func testAccResourceVariableSetOrgLevel(id string, name string) string {
	return fmt.Sprintf(`

		resource "harness_platform_organization" "testorg" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_secret_text" "testorg" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "testorg" {
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
			environment_variable {
				key = "key2"
				value = "value2up"
				value_type = "secret"
			}

			terraform_variable {
				key = "key1"
				value = "1111"
				value_type = "string"
			}
			terraform_variable {
				key = "key2"
				value = "1111u"
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
			connector {
				connector_ref = "account.azureCon"
				type = "azure"
			}

		}		

`, id, name)
}

func testAccResourceVariableSetAccLevel(id string, name string) string {
	return fmt.Sprintf(`

		resource "harness_platform_secret_text" "testacc" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "testacc" {
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
			environment_variable {
				key = "key2"
				value = "value2up"
				value_type = "secret"
			}

			terraform_variable {
				key = "key1"
				value = "1111"
				value_type = "string"
			}
			terraform_variable {
				key = "key2"
				value = "1111u"
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

`, id, name)
}
