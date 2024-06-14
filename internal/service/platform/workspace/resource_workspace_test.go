package workspace_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceWorkspace(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	resourceName := "harness_platform_workspace.test"
	updatedName := fmt.Sprintf("%s_updated", name)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspace(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceWorkspace(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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

func testAccResourceWorkspaceDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		ws, _ := testAccGetPlatformWorkspace(resourceName, state)
		if ws != nil {
			return fmt.Errorf("Workspace found: %s", ws.Identifier)
		}
		return nil
	}
}

func testAccGetPlatformWorkspace(resourceName string, state *terraform.State) (*nextgen.IacmShowWorkspaceResponseBody, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	org := r.Primary.Attributes["org_id"]
	project := r.Primary.Attributes["project_id"]

	workspace, resp, err := c.WorkspaceApi.WorkspacesShowWorkspace(ctx, org, project, id, c.AccountId)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return &workspace, nil
}

func testAccResourceWorkspace(id string, name string) string {
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
			provider_connector      = "account.${harness_platform_connector_github.test.id}"
			repository_connector    = "account.${harness_platform_connector_github.test.id}"
			environment_variable {
				key = "key1"
				value = "value1"
				value_type = "string"
			}
			environment_variable {
				key = "key2"
				value = "account.${harness_platform_secret_text.test.id}"
				value_type = "secret"
			}
			terraform_variable {
				key = "key1"
				value = "1111"
				value_type = "string"
			}
			terraform_variable {
				key = "key2"
				value = "value2"
				value_type = "string"
			}
			terraform_variable_file {
				repository              = "https://github.com/org/repo"
				repository_branch       = "main"
				repository_path         = "tf/aws/basic"
				repository_connector    = "account.${harness_platform_connector_github.test.id}"
			}			
  		}
`, id, name)
}
