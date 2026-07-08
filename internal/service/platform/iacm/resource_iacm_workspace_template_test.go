package iacm_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccResourceIacmWorkspaceTemplate exercises create, in-place version update and
// import for the workspace<->template association resource. The underlying Workspace
// template (two versions, so the in-place version switch is exercised for real) is
// provisioned as a harness_platform_template dependency.
func TestAccResourceIacmWorkspaceTemplate(t *testing.T) {
	resourceName := "harness_platform_iacm_workspace_template.test"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccResourceIacmWorkspaceTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceIacmWorkspaceTemplate(id, name, "v1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workspace_id", id),
					resource.TestCheckResourceAttr(resourceName, "template_id", id),
					resource.TestCheckResourceAttr(resourceName, "version", "v1"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
				),
			},
			{
				// version is updatable in place (no ForceNew).
				Config: testAccResourceIacmWorkspaceTemplate(id, name, "v2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workspace_id", id),
					resource.TestCheckResourceAttr(resourceName, "template_id", id),
					resource.TestCheckResourceAttr(resourceName, "version", "v2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccIacmWorkspaceTemplateImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
		},
	})
}

// testAccIacmWorkspaceTemplateImportStateIdFunc builds the 4-part import id
// (<org_id>/<project_id>/<template_id>/<workspace_id>) documented in import.sh.
func testAccIacmWorkspaceTemplateImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		orgId := primary.Attributes["org_id"]
		projId := primary.Attributes["project_id"]
		templateId := primary.Attributes["template_id"]
		workspaceId := primary.Attributes["workspace_id"]
		return fmt.Sprintf("%s/%s/%s/%s", orgId, projId, templateId, workspaceId), nil
	}
}

func testAccResourceIacmWorkspaceTemplateDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		assoc, _ := testAccGetIacmWorkspaceTemplate(resourceName, state)
		if assoc != nil {
			return fmt.Errorf("workspace template association found: %s/%s", assoc.TemplateID, assoc.WorkspaceID)
		}
		return nil
	}
}

func testAccGetIacmWorkspaceTemplate(resourceName string, state *terraform.State) (*nextgen.IacmWorkspaceTemplateResource, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	if r == nil {
		return nil, nil
	}
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	org := r.Primary.Attributes["org_id"]
	project := r.Primary.Attributes["project_id"]
	templateId := r.Primary.Attributes["template_id"]
	workspaceId := r.Primary.Attributes["workspace_id"]

	results, _, err := c.WorkspaceTemplatesApi.WorkspaceTemplatesGetWorkspacesByTemplateID(ctx, c.AccountId, org, project, templateId)
	if err != nil {
		return nil, err
	}

	for i := range results {
		if results[i].WorkspaceID == workspaceId {
			return &results[i], nil
		}
	}

	return nil, nil
}

// testAccResourceIacmWorkspaceTemplate builds a config with all dependencies: an org,
// project, github connector, a workspace, a Workspace-type template with versions v1
// and v2, and the association pointing at the requested version.
func testAccResourceIacmWorkspaceTemplate(id string, name string, version string) string {
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

		resource "harness_platform_workspace" "test" {
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
		}

		%[3]s

		resource "harness_platform_iacm_workspace_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			workspace_id = harness_platform_workspace.test.identifier
			template_id  = harness_platform_template.test_v1.identifier
			version      = "%[4]s"

			depends_on = [harness_platform_template.test_v2]
		}
`, id, name, testAccWorkspaceTemplateVersions(id, name), version)
}

// testAccWorkspaceTemplateVersions returns two versions (v1, v2) of a Workspace-type
// harness_platform_template sharing the same identifier, serialized with a time_sleep
// so the second version is created after the first.
func testAccWorkspaceTemplateVersions(id string, name string) string {
	return fmt.Sprintf(`
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
`, id, name, testAccWorkspaceTemplateYaml(id, name, "v1"), testAccWorkspaceTemplateYaml(id, name, "v2"))
}

// testAccWorkspaceTemplateYaml renders the YAML body of a Workspace-type template.
func testAccWorkspaceTemplateYaml(id string, name string, version string) string {
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
      locked: true
    provisioner:
      type: opentofu
      version: 1.12.3
      locked: true
    repository:
      isHarnessCode: false
      connector: account.${harness_platform_connector_github.test.id}
      name: spring5-oauth2-resource-server
      branch: master
      gitFetchType: branch
      locked: true
    variables: []
    terraform_variables: []
    terraform_variable_files: []`, id, name, version)
}
