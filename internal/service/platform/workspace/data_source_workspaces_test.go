package workspace_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceWorkspaces(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	allDataSourceName := "data.harness_platform_workspaces.all"
	filteredDataSourceName := "data.harness_platform_workspaces.filtered"
	filteredID := fmt.Sprintf("%s-2", id)
	filteredName := fmt.Sprintf("%s-2", name)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceWorkspaces(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(allDataSourceName, "workspaces.#", "2"),
					resource.TestCheckResourceAttr(allDataSourceName, "identifiers.#", "2"),
					resource.TestCheckResourceAttr(filteredDataSourceName, "workspaces.#", "1"),
					resource.TestCheckResourceAttr(filteredDataSourceName, "identifiers.#", "1"),
					resource.TestCheckResourceAttr(filteredDataSourceName, "workspaces.0.identifier", filteredID),
					resource.TestCheckResourceAttr(filteredDataSourceName, "workspaces.0.name", filteredName),
					resource.TestCheckResourceAttr(filteredDataSourceName, "identifiers.0", filteredID),
				),
			},
		},
	})
}

func testAccDataSourceWorkspaces(id string, name string) string {
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

		resource "harness_platform_connector_aws" "test" {
			identifier = "aws%[1]s"
			name = "%[2]s-aws"
			description = "test"
			tags = ["foo:bar"]

			manual {
				access_key_ref = "account.${harness_platform_secret_text.test.id}"
				secret_key_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		resource "harness_platform_workspace" "test_1" {
			identifier = "%[1]s-1"
			name = "%[2]s-1"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "description"
			provisioner_type        = "terraform"
			provisioner_version     = "1.5.6"
			repository              = "https://github.com/org/repo"
			repository_branch       = "main"
			repository_path         = "tf/aws/basic"
			cost_estimation_enabled = true
			provider_connector      = "account.${harness_platform_connector_aws.test.id}"
			repository_connector    = "account.${harness_platform_connector_github.test.id}"
			connector {
				connector_ref = "account.${harness_platform_connector_aws.test.id}"
				type          = "aws"
			}
  		}

		resource "harness_platform_workspace" "test_2" {
			identifier = "%[1]s-2"
			name = "%[2]s-2"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "description"
			provisioner_type        = "terraform"
			provisioner_version     = "1.5.6"
			repository              = "https://github.com/org/repo"
			repository_branch       = "main"
			repository_path         = "tf/aws/basic"
			cost_estimation_enabled = true
			provider_connector      = "account.${harness_platform_connector_aws.test.id}"
			repository_connector    = "account.${harness_platform_connector_github.test.id}"
			connector {
				connector_ref = "account.${harness_platform_connector_aws.test.id}"
				type          = "aws"
			}
  		}

		data "harness_platform_workspaces" "all" {
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			depends_on = [harness_platform_workspace.test_1, harness_platform_workspace.test_2]
		}

		data "harness_platform_workspaces" "filtered" {
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			search_term = "%[2]s-2"
			depends_on = [harness_platform_workspace.test_1, harness_platform_workspace.test_2]
		}

`, id, name)
}
