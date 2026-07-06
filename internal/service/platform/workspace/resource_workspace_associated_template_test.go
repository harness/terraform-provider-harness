package workspace_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccResourceWorkspaceAssociatedTemplate verifies the associated_template block on
// the workspace resource: association at create time and in-place version updates. The
// underlying Workspace template (versions v1 and v2) is provisioned as a
// harness_platform_template dependency.
func TestAccResourceWorkspaceAssociatedTemplate(t *testing.T) {
	resourceName := "harness_platform_workspace.test_tpl"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspaceAssociatedTemplate(id, name, "v1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "associated_template.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "associated_template.0.template_id", id),
					resource.TestCheckResourceAttr(resourceName, "associated_template.0.version", "v1"),
				),
			},
			{
				// version can be updated in place; template_id is ForceNew.
				Config: testAccResourceWorkspaceAssociatedTemplate(id, name, "v2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "associated_template.0.template_id", id),
					resource.TestCheckResourceAttr(resourceName, "associated_template.0.version", "v2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
				// associated_template is only reflected on read when already declared on the
				// workspace in state; a freshly imported workspace has no block yet, so it is
				// not populated on import.
				ImportStateVerifyIgnore: []string{"associated_template"},
			},
		},
	})
}

// TestAccResourceWorkspaceAssociatedTemplate_AddAndRemove verifies that a template can be
// associated with, and removed from, a workspace across applies. Because template_id is
// ForceNew, both transitions recreate the workspace (embedding the association on create,
// dropping it when the old workspace is destroyed) rather than going through update.
func TestAccResourceWorkspaceAssociatedTemplate_AddAndRemove(t *testing.T) {
	resourceName := "harness_platform_workspace.test_tpl"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// Create the workspace with no associated template.
				Config: testAccResourceWorkspaceAssociatedTemplateNoTemplate(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "associated_template.#", "0"),
				),
			},
			{
				// Add a template (ForceNew on template_id recreates the workspace).
				Config: testAccResourceWorkspaceAssociatedTemplate(id, name, "v1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "associated_template.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "associated_template.0.template_id", id),
					resource.TestCheckResourceAttr(resourceName, "associated_template.0.version", "v1"),
				),
			},
			{
				// Remove the template (again recreates the workspace, dropping the association).
				Config: testAccResourceWorkspaceAssociatedTemplateNoTemplate(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "associated_template.#", "0"),
				),
			},
		},
	})
}

// testAccResourceWorkspaceAssociatedTemplate builds a workspace with an associated_template
// block (version v1/v2) pointing at a Workspace-type harness_platform_template.
func testAccResourceWorkspaceAssociatedTemplate(id string, name string, version string) string {
	block := fmt.Sprintf(`
			associated_template {
				template_id = harness_platform_template.test_v1.identifier
				version     = "%s"
			}
`, version)
	return testAccResourceWorkspaceAssociatedTemplateConfig(id, name, block)
}

// testAccResourceWorkspaceAssociatedTemplateNoTemplate builds the same workspace with no
// associated_template block, so the association can be added later via update (POST path).
func testAccResourceWorkspaceAssociatedTemplateNoTemplate(id string, name string) string {
	return testAccResourceWorkspaceAssociatedTemplateConfig(id, name, "")
}

// testAccResourceWorkspaceAssociatedTemplateConfig renders the org/project/connector/two
// template versions and a workspace with the supplied associated_template block (which may
// be empty).
func testAccResourceWorkspaceAssociatedTemplateConfig(id string, name string, associatedTemplateBlock string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_platform_secret_text" "test" {
			identifier                = "%[1]s"
			name                      = "%[2]s"
			description               = "test"
			tags                      = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type                = "Inline"
			value                     = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier         = "%[1]s"
			name               = "%[2]s"
			description        = "test"
			tags               = ["foo:bar"]
			url                = "https://github.com/account"
			connection_type    = "Account"
			validation_repo    = "some_repo"
			delegate_selectors = ["harness-delegate"]
			credentials {
				http {
					username  = "admin"
					token_ref = "account.${harness_platform_secret_text.test.id}"
				}
			}
		}

		resource "harness_platform_template" "test_v1" {
			identifier   = "%[1]s"
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			name         = "%[2]s"
			version      = "v1"
			is_stable    = true
			force_delete = true
			template_yaml = <<-EOT
%[3]s
			EOT
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on       = [harness_platform_template.test_v1]
			destroy_duration = "4s"
		}

		resource "harness_platform_template" "test_v2" {
			identifier   = "%[1]s"
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			name         = "%[2]s"
			version      = "v2"
			is_stable    = false
			force_delete = true
			template_yaml = <<-EOT
%[4]s
			EOT

			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "harness_platform_workspace" "test_tpl" {
			identifier              = "%[1]s"
			name                    = "%[2]s"
			org_id                  = harness_platform_organization.test.id
			project_id              = harness_platform_project.test.id
			provisioner_type        = "terraform"
			provisioner_version     = "1.5.6"
			repository              = "https://github.com/org/repo"
			repository_branch       = "main"
			repository_path         = "tf/aws/basic"
			cost_estimation_enabled = true
			provider_connector      = "account.${harness_platform_connector_github.test.id}"
			repository_connector    = "account.${harness_platform_connector_github.test.id}"
%[5]s
			depends_on = [harness_platform_template.test_v2]
		}
`, id, name,
		testAccWorkspaceAssociatedTemplateYaml(id, name, "v1"),
		testAccWorkspaceAssociatedTemplateYaml(id, name, "v2"),
		associatedTemplateBlock)
}

// testAccWorkspaceAssociatedTemplateYaml renders the YAML body of a Workspace-type template.
func testAccWorkspaceAssociatedTemplateYaml(id string, name string, version string) string {
	return fmt.Sprintf(`template:
  name: "%[2]s"
  identifier: "%[1]s"
  versionLabel: %[3]s
  type: Workspace
  projectIdentifier: ${harness_platform_project.test.id}
  orgIdentifier: ${harness_platform_organization.test.id}
  tags: {}
  spec:
    provider:
      type: others
      title: thirdPartyGitProvider
      info: thirdPartyGitProviderInfo
      icon: service-github
      size: 20
      disabled: false
    tags: []
    description: ""
    cost_estimation:
      enabled: false
      locked: false
    default_pipelines:
      plan: ""
      apply: ""
      destroy: ""
      drift: ""
      locked: false
    provisioner:
      type: opentofu
      version: 1.12.3
      locked: false
    repository:
      isHarnessCode: false
      connector: account.${harness_platform_connector_github.test.id}
      name: spring5-oauth2-resource-server
      branch: master
      gitFetchType: branch
      locked: false
    variables: []
    terraform_variables: []
    terraform_variable_files: []`, id, name, version)
}
