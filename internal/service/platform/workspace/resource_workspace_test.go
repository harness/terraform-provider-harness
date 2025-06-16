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
	resourceName := "harness_platform_workspace.test"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	name2 := t.Name()
	id2 := fmt.Sprintf("%s_%s", name2, utils.RandStringBytes(5))
	updatedName2 := fmt.Sprintf("%s_updated", name2)
	name3 := t.Name()
	id3 := fmt.Sprintf("%s_%s", name3, utils.RandStringBytes(5))
	updatedName3 := fmt.Sprintf("%s_updated", name3)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspace(id, name, "branch"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "default_pipelines.%", "4"),
					resource.TestCheckResourceAttr(resourceName, "default_pipelines.destroy", "destroy_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "default_pipelines.drift", "drift_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "default_pipelines.apply", "apply_pipeline_id"),
					resource.TestCheckResourceAttr(resourceName, "default_pipelines.plan", "plan_pipeline_id"),
				),
			},
			{
				Config: testAccResourceWorkspace(id, updatedName, "branch"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
			{
				Config: testAccResourceWorkspace(id2, name2, "commit"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id2),
				),
			},
			{
				Config: testAccResourceWorkspace(id2, updatedName2, "commit"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id2),
				),
			},
			{
				Config: testAccResourceWorkspace(id3, name3, "sha"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id3),
				),
			},
			{
				Config: testAccResourceWorkspace(id3, updatedName3, "sha"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id3),
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

func testAccResourceWorkspace(id string, name string, repositoryType string) string {

	var repositoryX = ""
	if repositoryType == "branch" {
		repositoryX = `repository_branch        = "main"`
	} else if repositoryType == "commit" {
		repositoryX = `repository_commit        = "tag1"`
	} else {
		repositoryX = `repository_sha        = "abcdef12345"`
	}

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
			%[3]s
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
				%[3]s
				repository_path         = "tf/aws/basic"
				repository_connector    = "account.${harness_platform_connector_github.test.id}"
			}
			default_pipelines = {
				"destroy" = "destroy_pipeline_id"
				"drift"   = "drift_pipeline_id"
				"plan"    = "plan_pipeline_id"
				"apply"   = "apply_pipeline_id"
  			}			
			variable_sets = [harness_platform_infra_variable_set.test.id]
  		}

		resource "harness_platform_workspace" "withMuptipleConnectors" {
			identifier = "w%[1]s"
			name = "w%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "description"
			provisioner_type        = "terraform"
			provisioner_version     = "1.5.6"
			repository              = "https://github.com/org/repo"
			%[3]s
			repository_path         = "tf/aws/basic"
			cost_estimation_enabled = true
			repository_connector    = "account.${harness_platform_connector_github.test.id}"
			connector {
				connector_ref = "con1"
				type = "aws"
			}
			connector {
				connector_ref = "con2"
				type = "azure"
			}
			connector {
				connector_ref = "con3"
				type = "gcp"
			}
  		}
`, id, name, repositoryX)
}
