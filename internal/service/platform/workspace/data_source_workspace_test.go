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

func TestAccDataSourceWorkspace_list(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "data.harness_platform_workspace.test_list"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceWorkspace_list(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "tags.0", "tag1"),
					resource.TestCheckResourceAttr(resourceName, "tags.1", "tag2"),
				),
			},
		},
	})
}

func testAccDataSourceWorkspace(id, name string) string {
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

		data "harness_platform_workspace" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			depends_on = [harness_platform_workspace.test]
		}

`, id, name)
}

func testAccDataSourceWorkspace_list(id, name string) string {
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

		resource "harness_platform_workspace" "list1" {
			identifier = "list1"
			name = "list1"
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

		resource "harness_platform_workspace" "list2" {
			identifier = "list2"
			name = "list-2"
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

		data "harness_platform_workspace" "test_list" {
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			name_prefix = "list"
			depends_on = [harness_platform_workspace.list1, harness_platform_workspace.list2]
		}

		data "harness_platform_workspace" "test_list_all" {
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			depends_on = [harness_platform_workspace.list1, harness_platform_workspace.list2]
		}

`, id, name)
}
