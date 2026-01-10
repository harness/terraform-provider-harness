package workspace_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceWorkspace(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "data.harness_platform_workspace.test"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceWorkspace(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "tag1"),
					resource.TestCheckResourceAttr(resourceName, "tags.1", "tag2"),
				),
			},
		},
	})
}

func testAccDataSourceWorkspace(id string, name string) string {
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
		}		

		resource "harness_platform_workspace" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "description"
			provisioner_type        = "terraform"
			provisioner_version     = "1.5.6"
			repository              = "https://github.com/org/repo"
			repository_branch       = "main"
			repository_path         = "tf/aws/basic"
			cost_estimation_enabled = true
			repository_connector    = "account.${harness_platform_connector_github.test.id}"
			variable_sets = [harness_platform_infra_variable_set.test.id]
			tags = ["tag1", "tag2"]
  		}

		resource "harness_platform_workspace" "test1" {
			identifier = "test1"
			name = "test1"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "description"
			provisioner_type        = "terraform"
			provisioner_version     = "1.5.6"
			repository              = "https://github.com/org/repo"
			repository_branch       = "main"
			repository_path         = "tf/aws/basic"
			cost_estimation_enabled = true
			repository_connector    = "account.${harness_platform_connector_github.test.id}"
			variable_sets = [harness_platform_infra_variable_set.test.id]
			tags = ["tag1", "tag2"]
  		}

		resource "harness_platform_workspace" "test2" {
			identifier = "test2"
			name = "test2"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "description"
			provisioner_type        = "terraform"
			provisioner_version     = "1.5.6"
			repository              = "https://github.com/org/repo"
			repository_branch       = "main"
			repository_path         = "tf/aws/basic"
			cost_estimation_enabled = true
			repository_connector    = "account.${harness_platform_connector_github.test.id}"
			variable_sets = [harness_platform_infra_variable_set.test.id]
			tags = ["tag1", "tag2"]
  		}

		data "harness_platform_workspace" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			depends_on = [harness_platform_workspace.test]
		}

		data "harness_platform_workspace" "test2" {
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			name_prefix = "test"
			depends_on = [harness_platform_workspace.test]
		}

		output "workspace_ids" {
			value = data.harness_platform_workspace.test2
		}

`, id, name)
}
