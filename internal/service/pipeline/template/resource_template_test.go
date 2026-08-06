package template_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/utils"
	openapi_client_nextgen "github.com/harness/harness-openapi-go-client/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceTemplateProjectScope(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	updatedName := fmt.Sprintf("%s_updated", id)

	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
			"null": {},
		},
		CheckDestroy: testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateProjectScope(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				Config: testAccResourceTemplateProjectScope(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type", "comments", "description"},
			},
		},
	})
}

func TestAccResourceTemplateProjectScopeInline(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	updatedName := fmt.Sprintf("%s_updated", id)

	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateProjectScopeInline(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				Config: testAccResourceTemplateProjectScopeInline(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type", "comments"},
			},
		},
	})
}

func TestAccResourceTemplate_OrgScope(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	updatedName := fmt.Sprintf("%s_updated", id)

	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
			"null": {},
		},
		CheckDestroy: testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateOrgScope(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				Config: testAccResourceTemplateOrgScope(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type", "comments"},
			},
		},
	})
}

func TestAccResourceTemplate_OrgScopeInline(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	updatedName := fmt.Sprintf("%s_updated", id)

	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateOrgScopeInline(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				Config: testAccResourceTemplateOrgScopeInline(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type", "comments"},
			},
		},
	})
}

func TestAccResourceTemplate_OrgScopeImportFromGit(t *testing.T) {

	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
			"null": {},
		},
		CheckDestroy: testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateOrgScopeImportFromGit(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "orgtemplate2"),
					resource.TestCheckResourceAttr(resourceName, "name", "orgtemplate2"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type", "comments", "git_details.0.branch_name", "git_details.0.file_path", "git_details.0.last_commit_id", "git_details.0.repo_name", "git_import_details.#", "git_import_details.0.%", "git_import_details.0.branch_name", "git_import_details.0.connector_ref", "git_import_details.0.file_path", "git_import_details.0.is_force_import", "git_import_details.0.repo_name", "import_from_git", "is_stable", "template_import_request.#", "template_import_request.0.%", "template_import_request.0.template_description", "template_import_request.0.template_name", "template_import_request.0.template_version", "template_yaml", "version", "git_details.0.last_object_id"},
			},
		},
	})
}

func TestAccResourceTemplate_ProjectScopeImportFromGit(t *testing.T) {
	id := "projecttemplate"
	name := id
	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
			"null": {},
		},
		CheckDestroy: testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateProjectScopeImportFromGit(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "projecttemplate"),
					resource.TestCheckResourceAttr(resourceName, "name", "projecttemplate"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type", "comments", "git_details.0.branch_name", "git_details.0.file_path", "git_details.0.last_commit_id", "git_details.0.repo_name", "git_import_details.#", "git_import_details.0.%", "git_import_details.0.branch_name", "git_import_details.0.connector_ref", "git_import_details.0.file_path", "git_import_details.0.is_force_import", "git_import_details.0.repo_name", "import_from_git", "is_stable", "template_import_request.#", "template_import_request.0.%", "template_import_request.0.template_description", "template_import_request.0.template_name", "template_import_request.0.template_version", "template_yaml", "version", "git_details.0.last_object_id"},
			},
		},
	})
}

func TestAccResourceTemplate_ProjectScopeImportFromGitNonDefaultBranch(t *testing.T) {
	id := "projecttemplate"
	name := id
	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
			"null": {},
		},
		CheckDestroy: testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateProjectScopeImportFromGitNonDefaultBranch(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "projecttemplate"),
					resource.TestCheckResourceAttr(resourceName, "name", "projecttemplate"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type", "comments", "git_details.0.branch_name", "git_details.0.file_path", "git_details.0.last_commit_id", "git_details.0.repo_name", "git_import_details.#", "git_import_details.0.%", "git_import_details.0.branch_name", "git_import_details.0.connector_ref", "git_import_details.0.file_path", "git_import_details.0.is_force_import", "git_import_details.0.repo_name", "import_from_git", "is_stable", "template_import_request.#", "template_import_request.0.%", "template_import_request.0.template_description", "template_import_request.0.template_name", "template_import_request.0.template_version", "template_yaml", "version", "git_details.0.last_object_id"},
			},
		},
	})
}

func TestAccResourceTemplate_AccountScopeImportFromGit(t *testing.T) {
	id := "accounttemplate"
	name := id

	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
			"null": {},
		},
		CheckDestroy: testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateAccountScopeImportFromGit(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "accounttemplate"),
					resource.TestCheckResourceAttr(resourceName, "name", "accounttemplate"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type", "comments", "git_details.0.branch_name", "git_details.0.file_path", "git_details.0.last_commit_id", "git_details.0.repo_name", "git_import_details.#", "git_import_details.0.%", "git_import_details.0.branch_name", "git_import_details.0.connector_ref", "git_import_details.0.file_path", "git_import_details.0.is_force_import", "git_import_details.0.repo_name", "import_from_git", "is_stable", "template_import_request.#", "template_import_request.0.%", "template_import_request.0.template_description", "template_import_request.0.template_name", "template_import_request.0.template_version", "template_yaml", "version", "git_details.0.last_object_id"},
			},
		},
	})
}

func TestAccResourceTemplate_OrgScopeInline_UpdateStable(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id

	resourceName := "harness_platform_template.test2"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTemplateDestroy(resourceName),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateOrgScopeInlineMultipleVersion(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				Config: testAccResourceTemplateOrgScopeInlineUpdateStable(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
		},
	})
}

func TestAccResourceTemplate_AccountScope(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	updatedName := fmt.Sprintf("%s_updated", id)

	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
			"null": {},
		},
		CheckDestroy: testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateAccScope(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				Config: testAccResourceTemplateAccScope(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type", "comments"},
			},
		},
	})
}

func testAccTemplateDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		template, _ := testAccGetTemplate(resourceName, state)
		var template_identifier string
		if template != nil {
			if template.Template.Identifier != "" {
				template_identifier = template.Template.Identifier
			} else {
				template_identifier = template.Template.Slug
			}
			return fmt.Errorf("Found template: %s", template_identifier)
		}

		return nil
	}
}

func TestAccResourceTemplate_AccountScopeInline(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	updatedName := fmt.Sprintf("%s_updated", id)

	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
			"null": {},
		},
		CheckDestroy: testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateAccScopeInline(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				Config: testAccResourceTemplateAccScopeInline(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "comments", "comments"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       acctest.OrgResourceImportStateIdFunc(resourceName),
				ImportStateVerifyIgnore: []string{"git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type", "comments"},
			},
		},
	})
}

func TestAccResourceTemplate_DeleteUnderlyingResource(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	project_id := id + "project"
	org_id := id + "org"
	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateProjectScopeInline(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					c, ctx := acctest.TestAccGetClientWithContext()
					_, err := c.ProjectTemplateApi.DeleteTemplateProject(ctx, project_id, id, org_id, "ab", &openapi_client_nextgen.ProjectTemplateApiDeleteTemplateProjectOpts{})
					require.NoError(t, err)
				},
				Config:             testAccResourceTemplateProjectScopeInline(id, name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccGetTemplate(resourceName string, state *terraform.State) (*openapi_client_nextgen.TemplateWithInputsResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetClientWithContext()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]
	version := r.Primary.Attributes["version"]
	branch_name := r.Primary.Attributes["branch_name"]
	var err error
	var resp openapi_client_nextgen.TemplateWithInputsResponse

	if projId != "" {
		if version == "" {
			resp, _, err = c.ProjectTemplateApi.GetTemplateStableProject(ctx, orgId, projId, id, &openapi_client_nextgen.ProjectTemplateApiGetTemplateStableProjectOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     buildField(r, branch_name)})
		} else {
			resp, _, err = c.ProjectTemplateApi.GetTemplateProject(ctx, orgId, projId, id, version, &openapi_client_nextgen.ProjectTemplateApiGetTemplateProjectOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     buildField(r, branch_name)})
		}
	} else if orgId != "" && projId == "" {
		if version == "" {
			resp, _, err = c.OrgTemplateApi.GetTemplateStableOrg(ctx, orgId, id, &openapi_client_nextgen.OrgTemplateApiGetTemplateStableOrgOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     buildField(r, branch_name),
			})
		} else {
			resp, _, err = c.OrgTemplateApi.GetTemplateOrg(ctx, orgId, id, version, &openapi_client_nextgen.OrgTemplateApiGetTemplateOrgOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     buildField(r, branch_name),
			})
		}
	} else {
		if version == "" {
			resp, _, err = c.AccountTemplateApi.GetTemplateStableAcc(ctx, id, &openapi_client_nextgen.AccountTemplateApiGetTemplateStableAccOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     buildField(r, branch_name),
			})
		} else {
			resp, _, err = c.AccountTemplateApi.GetTemplateAcc(ctx, id, version, &openapi_client_nextgen.AccountTemplateApiGetTemplateAccOpts{
				HarnessAccount: optional.NewString(c.AccountId),
				BranchName:     buildField(r, branch_name),
			})
		}
	}

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccResourceTemplateAccScopeInline(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_template" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			comments = "comments"
			version = "ab"
			is_stable = true
			git_details {
				branch_name = "main"
				commit_message = "Commit"
				file_path = ".harness/GitEnabledPipeline%[1]s.yaml"
				connector_ref = "account.TF_Jajoo_github_connector"
				store_type = "REMOTE"
				repo_name = "jajoo_git"
		}
			template_yaml = <<-EOT
			template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: ab
      type: Pipeline
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
    
      EOT
	}
	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_template.test]
		destroy_duration = "4s"
	}
	
	resource "null_resource" "next" {
  			depends_on = [time_sleep.wait_4_seconds]
	}
	`, id, name)
}

func testAccResourceTemplateAccScope(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_template" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			comments = "comments"
			version = "ab"
			is_stable = true
			git_details {
				branch_name = "main"
				commit_message = "Commit"
				file_path = ".harness/GitEnabledPipeline%[1]s.yaml"
				connector_ref = "account.TF_Jajoo_github_connector"
				store_type = "REMOTE"
				repo_name = "jajoo_git"
		}
			template_yaml = <<-EOT
			template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: ab
      type: Pipeline
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
    
      EOT
	}
	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_template.test]
		destroy_duration = "4s"
	}

	resource "null_resource" "next" {
  			depends_on = [time_sleep.wait_4_seconds]
	}
	`, id, name)
}

func testAccResourceTemplateOrgScopeInline(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]sorg"
		name = "%[2]s"
	}

	resource "harness_platform_template" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			name = "%[2]s"
			comments = "comments"
			version = "ab"
			is_stable = true
			template_yaml = <<-EOT
			template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: ab
      type: Pipeline
      orgIdentifier: ${harness_platform_organization.test.id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
    
      EOT
	}
	`, id, name)
}

func testAccResourceTemplateOrgScopeInlineMultipleVersion(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]sorg"
		name = "%[2]s"
	}

	resource "harness_platform_template" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			name = "%[2]s"
			comments = "comments"
			version = "ab"
			is_stable = true
			force_delete = true
			template_yaml = <<-EOT
			template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: ab
      type: Pipeline
      orgIdentifier: ${harness_platform_organization.test.id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
    
      EOT
	}

	resource "harness_platform_template" "test2" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			name = "%[2]s"
			comments = "comments"
			version = "abc"
			is_stable = false
			force_delete = true
			template_yaml = <<-EOT
			template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: abc
      type: Pipeline
      orgIdentifier: ${harness_platform_organization.test.id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
      EOT

	  depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_template.test]
		destroy_duration = "4s"
	}
	`, id, name)
}

func testAccResourceTemplateOrgScopeInlineUpdateStable(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]sorg"
		name = "%[2]s"
	}

	resource "harness_platform_template" "test2" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			name = "%[2]s"
			comments = "comments"
			version = "abc"
			is_stable = true
			force_delete = true
			template_yaml = <<-EOT
			template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: abc
      type: Pipeline
      orgIdentifier: ${harness_platform_organization.test.id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
      EOT
	}

	resource "harness_platform_template" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			name = "%[2]s"
			comments = "comments"
			version = "ab"
			is_stable = false
			force_delete = true
			template_yaml = <<-EOT
			template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: ab
      type: Pipeline
      orgIdentifier: ${harness_platform_organization.test.id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
    
      EOT

	  depends_on = [time_sleep.wait_10_seconds]
	}

	resource "time_sleep" "wait_10_seconds" {
		depends_on = [harness_platform_template.test2]
		destroy_duration = "10s"
	}
	`, id, name)
}

func testAccResourceTemplateOrgScopeInlineUpdateStable2(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_template" "test2" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			name = "%[2]s"
			comments = "comments"
			version = "abc"
			force_delete = true
			template_yaml = <<-EOT
			template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: abc
      type: Pipeline
      orgIdentifier: ${harness_platform_organization.test.id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
    
      EOT

	  depends_on = [harness_platform_template.test]

	}

	resource "harness_platform_template" "test" {
		identifier = "%[1]s"
		org_id = "%[1]s"
		name = "%[2]s"
		comments = "comments"
		force_delete = true
		version = "ab"
}

	`, id, name)
}

func testAccResourceTemplateOrgScope(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_template" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_organization.test.id
			name = "%[2]s"
			comments = "comments"
			version = "ab"
			is_stable = true
			git_details {
				branch_name = "main"
				commit_message = "Commit"
				file_path = ".harness/GitEnabledPipeline%[1]s.yaml"
				connector_ref = "account.TF_Jajoo_github_connector"
				store_type = "REMOTE"
				repo_name = "jajoo_git"
		}
			template_yaml = <<-EOT
			template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: ab
      type: Pipeline
      orgIdentifier: ${harness_platform_organization.test.id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
    
      EOT
	}
	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_template.test]
		destroy_duration = "4s"
	}

	resource "null_resource" "next" {
  			depends_on = [time_sleep.wait_4_seconds]
	}
	`, id, name)
}

func testAccResourceTemplateProjectScope(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]sorg"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]sproject"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_template" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			name = "%[2]s"
			comments = "comments"
			version = "ab"
			description = "description"
			is_stable = true
			git_details {
				branch_name = "main"
				commit_message = "Commit"
				file_path = ".harness/GitEnabledPipeline%[1]s.yaml"
				connector_ref = "account.TF_Jajoo_github_connector"
				store_type = "REMOTE"
				repo_name = "jajoo_git"
		}
			template_yaml = <<-EOT
			template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: ab
      type: Pipeline
      projectIdentifier: ${harness_platform_project.test.id}
      orgIdentifier: ${harness_platform_project.test.org_id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
    
      EOT
	}
	
	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_template.test]
		destroy_duration = "4s"
	}
		
	resource "null_resource" "next" {
  			depends_on = [time_sleep.wait_4_seconds]
	}
	`, id, name)
}

func testAccResourceTemplateProjectScopeInline(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]sorg"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]sproject"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_template" "test" {
			identifier = "%[1]s"
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			name = "%[2]s"
			comments = "comments"
			version = "ab"
			is_stable = true
			template_yaml = <<-EOT
			template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: ab
      type: Pipeline
      projectIdentifier: ${harness_platform_project.test.id}
      orgIdentifier: ${harness_platform_project.test.org_id}
      tags: {}
      spec:
        stages:
          - stage:
              name: dvvdvd
              identifier: dvvdvd
              description: ""
              type: Deployment
              spec:
                deploymentType: Kubernetes
                service:
                  serviceRef: <+input>
                  serviceInputs: <+input>
                environment:
                  environmentRef: <+input>
                  deployToAll: false
                  environmentInputs: <+input>
                  serviceOverrideInputs: <+input>
                  infrastructureDefinitions: <+input>
                execution:
                  steps:
                    - step:
                        name: Rollout Deployment
                        identifier: rolloutDeployment
                        type: K8sRollingDeploy
                        timeout: 10m
                        spec:
                          skipDryRun: false
                          pruningEnabled: false
                  rollbackSteps:
                    - step:
                        name: Rollback Rollout Deployment
                        identifier: rollbackRolloutDeployment
                        type: K8sRollingRollback
                        timeout: 10m
                        spec:
                          pruningEnabled: false
              tags: {}
              failureStrategies:
                - onFailure:
                    errors:
                      - AllErrors
                    action:
                      type: StageRollback
    
      EOT
	}
	`, id, name)
}

func testAccResourceTemplateOrgScopeImportFromGit() string {
	return fmt.Sprintf(`
        resource "harness_platform_template" "test" {
                        identifier = "orgtemplate2"
                        org_id = "default"
                        name = "orgtemplate2"
						version = "v2"
                        import_from_git = true
                        git_import_details {
                            branch_name = "main"
                            file_path = ".harness/orgtemplate2.yaml"
                            connector_ref = "account.TF_open_repo_github_connector"
                            repo_name = "open-repo"
                        }
                        template_import_request {
                            template_name = "orgtemplate2"
							template_version = "v2"
                            template_description = ""
                        }
                }
		
		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_template.test]
			destroy_duration = "4s"
		}

		resource "null_resource" "next" {
  			depends_on = [time_sleep.wait_4_seconds]
		}
        `)
}

func testAccResourceTemplateProjectScopeImportFromGit(id string, name string) string {
	// This has project and org id static due to its config in the git.
	return fmt.Sprintf(`
		resource "harness_platform_project" "Project_Test" {
				identifier = "TF_Project_Pipeline_Test"
				name = "TF_Project_Pipeline_Test"
				color = "#0063F7"
				org_id = "default"
		}
        resource "harness_platform_template" "test" {
                        identifier = "%[1]s"
                        org_id = "default"
						project_id = harness_platform_project.Project_Test.identifier
                        name = "%[2]s"
						version = "v2"
                        import_from_git = true
                        git_import_details {
                            branch_name = "main"
                            file_path = ".harness/projecttemplate.yaml"
                            connector_ref = "account.TF_open_repo_github_connector"
                            repo_name = "open-repo"
                        }
                        template_import_request {
                            template_name = "%[2]s"
							template_version = "v2"
                            template_description = ""
                        }
                }

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_template.test]
			destroy_duration = "4s"
		}
		
		resource "null_resource" "next" {
  			depends_on = [time_sleep.wait_4_seconds]
		}
        `, id, name)
}

func testAccResourceTemplateProjectScopeImportFromGitNonDefaultBranch(id string, name string) string {
	// This has project and org id static due to its config in the git.
	return fmt.Sprintf(`
        resource "harness_platform_template" "test" {
                        identifier = "%[1]s"
                        org_id = "default"
						project_id = "TF_Pipeline_Test"
                        name = "%[2]s"
						version = "v2"
                        import_from_git = true
                        git_import_details {
                            branch_name = "main-patch"
                            file_path = ".harness/projecttemplate.yaml"
                            connector_ref = "account.TF_open_repo_github_connector"
                            repo_name = "open-repo"
                        }
                        template_import_request {
                            template_name = "%[2]s"
							template_version = "v2"
                            template_description = ""
                        }
                }

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_template.test]
			destroy_duration = "4s"
		}
		
		resource "null_resource" "next" {
  			depends_on = [time_sleep.wait_4_seconds]
		}
        `, id, name)
}

func testAccResourceTemplateAccountScopeImportFromGit(id string, name string) string {
	return fmt.Sprintf(`
        resource "harness_platform_template" "test" {
                        identifier = "%[1]s"
                        name = "%[2]s"
						version = "v2"

                        import_from_git = true
                        git_import_details {
                            branch_name = "main"
                            file_path = ".harness/accounttemplate.yaml"
                            connector_ref = "account.TF_open_repo_github_connector"
                            repo_name = "open-repo"
                        }
                        template_import_request {
                            template_name = "accounttemplate"
							template_version = "v2"
                            template_description = ""
                        }
                }

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_template.test]
			destroy_duration = "4s"
		}

		resource "null_resource" "next" {
  			depends_on = [time_sleep.wait_4_seconds]
		}
        `, id, name)
}

/*
# Needs different API keys for org and account scope tests.
# set HARNESS_PLATFORM_API_KEY, TF_VAR_harness_api_key
# project level import test
  HARNESS_TEST_ORG_ID=default \
  HARNESS_TEST_PROJECT_ID=senthilproj01 \
  TF_ACC=1 go test -v ./internal/service/pipeline/template/... \
      -run TestAccTemplateImport_ProjectScope \
      -timeout 20m

# org level import test
  export HARNESS_PLATFORM_API_KEY=<org-scoped-token>
  HARNESS_TEST_ORG_ID=senthilorg01 \
  TF_ACC=1 go test -v ./internal/service/pipeline/template/... \
      -run TestAccTemplateImport_OrgScope \
      -timeout 20m

# account level import test
  TF_ACC=1 go test -v ./internal/service/pipeline/template/... \
      -run TestAccTemplateImport_AccountScope \
      -timeout 20m
*/

// TestAccTemplateImport_ProjectScope_StableVersion verifies that importing with the stable
// version label fetches exactly that version, not just whatever Harness marks stable.
// Requires env vars: HARNESS_TEST_ORG_ID, HARNESS_TEST_PROJECT_ID
func TestAccTemplateImport_ProjectScope_StableVersion(t *testing.T) {
	orgID := os.Getenv("HARNESS_TEST_ORG_ID")
	projectID := os.Getenv("HARNESS_TEST_PROJECT_ID")
	if orgID == "" || projectID == "" {
		t.Skip("HARNESS_TEST_ORG_ID and HARNESS_TEST_PROJECT_ID must be set")
	}
	id := fmt.Sprintf("hsf86proj_%s", utils.RandStringBytes(6))
	name := id
	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// Create only v1 (stable) — single resource avoids ImportStateVerify confusion
				Config: testAccTemplateImportProjectSingleVersion(id, name, orgID, projectID, "v1", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "version", "v1"),
					resource.TestCheckResourceAttr(resourceName, "is_stable", "true"),
				),
			},
			{
				// Import stable version (v1) explicitly — must get v1
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       templateProjectImportStateIdWithVersion(resourceName, "v1"),
				ImportStateVerifyIgnore: []string{"comments", "force_delete", "git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type"},
			},
		},
	})
}

// TestAccTemplateImport_ProjectScope_NonStableVersion verifies that importing a non-stable
// version fetches that specific version, not the stable one.
// Requires env vars: HARNESS_TEST_ORG_ID, HARNESS_TEST_PROJECT_ID
func TestAccTemplateImport_ProjectScope_NonStableVersion(t *testing.T) {
	orgID := os.Getenv("HARNESS_TEST_ORG_ID")
	projectID := os.Getenv("HARNESS_TEST_PROJECT_ID")
	if orgID == "" || projectID == "" {
		t.Skip("HARNESS_TEST_ORG_ID and HARNESS_TEST_PROJECT_ID must be set")
	}
	id := fmt.Sprintf("hsf86projns_%s", utils.RandStringBytes(6))
	name := id
	resourceName := "harness_platform_template.test_v2"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// Create v1 (stable) and v2 (non-stable)
				Config: testAccTemplateImportProjectTwoVersions(id, name, orgID, projectID, "v1", "v2", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "version", "v2"),
					resource.TestCheckResourceAttr(resourceName, "is_stable", "false"),
				),
			},
			{
				// Import v2 (non-stable) — must get v2, not v1 (stable)
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateIdFunc: templateProjectImportStateIdWithVersion(resourceName, "v2"),
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) == 0 {
						return fmt.Errorf("no state after import")
					}
					attrs := s[0].Attributes
					if attrs["version"] != "v2" {
						return fmt.Errorf("expected version=v2 after import, got %q", attrs["version"])
					}
					if attrs["is_stable"] != "false" {
						return fmt.Errorf("expected is_stable=false after import, got %q", attrs["is_stable"])
					}
					return nil
				},
			},
		},
	})
}

// TestAccTemplateImport_OrgScope_StableVersion verifies import at org scope with stable version.
// Requires env var: HARNESS_TEST_ORG_ID
func TestAccTemplateImport_OrgScope_StableVersion(t *testing.T) {
	orgID := os.Getenv("HARNESS_TEST_ORG_ID")
	if orgID == "" {
		t.Skip("HARNESS_TEST_ORG_ID must be set")
	}
	id := fmt.Sprintf("hsf86org_%s", utils.RandStringBytes(6))
	name := id
	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccTemplateImportOrgSingleVersion(id, name, orgID, "v1", true),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "version", "v1"),
					resource.TestCheckResourceAttr(resourceName, "is_stable", "true"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       templateOrgImportStateIdWithVersion(resourceName, "v1"),
				ImportStateVerifyIgnore: []string{"force_delete", "comments", "git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type"},
			},
		},
	})
}

// TestAccTemplateImport_OrgScope_NonStableVersion verifies import of a non-stable version at org scope.
// Requires env var: HARNESS_TEST_ORG_ID
func TestAccTemplateImport_OrgScope_NonStableVersion(t *testing.T) {
	orgID := os.Getenv("HARNESS_TEST_ORG_ID")
	if orgID == "" {
		t.Skip("HARNESS_TEST_ORG_ID must be set")
	}
	id := fmt.Sprintf("hsf86orgns_%s", utils.RandStringBytes(6))
	name := id
	resourceName := "harness_platform_template.test_v2"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccTemplateImportOrgTwoVersions(id, name, orgID, "v1", "v2", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "version", "v2"),
					resource.TestCheckResourceAttr(resourceName, "is_stable", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateIdFunc: templateOrgImportStateIdWithVersion(resourceName, "v2"),
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) == 0 {
						return fmt.Errorf("no state after import")
					}
					attrs := s[0].Attributes
					if attrs["version"] != "v2" {
						return fmt.Errorf("expected version=v2 after import, got %q", attrs["version"])
					}
					if attrs["is_stable"] != "false" {
						return fmt.Errorf("expected is_stable=false after import, got %q", attrs["is_stable"])
					}
					return nil
				},
			},
		},
	})
}

// TestAccTemplateImport_AccountScope_StableVersion verifies import at account scope with stable version.
func TestAccTemplateImport_AccountScope_StableVersion(t *testing.T) {
	id := fmt.Sprintf("hsf86acc_%s", utils.RandStringBytes(6))
	name := id
	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccTemplateImportAccountSingleVersion(id, name, "v1", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "version", "v1"),
					resource.TestCheckResourceAttr(resourceName, "is_stable", "true"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       templateAccountImportStateIdWithVersion(resourceName, "v1"),
				ImportStateVerifyIgnore: []string{"force_delete", "comments", "git_details.0.commit_message", "git_details.0.connector_ref", "git_details.0.store_type"},
			},
		},
	})
}

// TestAccTemplateImport_AccountScope_NonStableVersion verifies import of a non-stable version at account scope.
func TestAccTemplateImport_AccountScope_NonStableVersion(t *testing.T) {
	id := fmt.Sprintf("hsf86accns_%s", utils.RandStringBytes(6))
	name := id
	resourceName := "harness_platform_template.test_v2"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccTemplateImportAccountTwoVersions(id, name, "v1", "v2", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "version", "v2"),
					resource.TestCheckResourceAttr(resourceName, "is_stable", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateIdFunc: templateAccountImportStateIdWithVersion(resourceName, "v2"),
				ImportStateCheck: func(s []*terraform.InstanceState) error {
					if len(s) == 0 {
						return fmt.Errorf("no state after import")
					}
					attrs := s[0].Attributes
					if attrs["version"] != "v2" {
						return fmt.Errorf("expected version=v2 after import, got %q", attrs["version"])
					}
					if attrs["is_stable"] != "false" {
						return fmt.Errorf("expected is_stable=false after import, got %q", attrs["is_stable"])
					}
					return nil
				},
			},
		},
	})
}

func templateProjectImportStateIdWithVersion(resourceName, version string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		id := primary.ID
		orgId := primary.Attributes["org_id"]
		projId := primary.Attributes["project_id"]
		return fmt.Sprintf("%s/%s/%s/%s", orgId, projId, id, version), nil
	}
}

func templateOrgImportStateIdWithVersion(resourceName, version string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		id := primary.ID
		orgId := primary.Attributes["org_id"]
		return fmt.Sprintf("%s/%s/%s", orgId, id, version), nil
	}
}

func templateAccountImportStateIdWithVersion(resourceName, version string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		id := primary.ID
		return fmt.Sprintf("%s/%s", id, version), nil
	}
}

func testAccTemplateImportProjectSingleVersion(id, name, orgID, projectID, version string, isStable bool) string {
	return fmt.Sprintf(`
resource "harness_platform_template" "test" {
  identifier   = "%[1]s"
  name         = "%[2]s"
  org_id       = "%[3]s"
  project_id   = "%[4]s"
  version      = "%[5]s"
  is_stable    = %[6]t
  force_delete = true
  template_yaml = <<-EOT
    template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: %[5]s
      type: Step
      projectIdentifier: %[4]s
      orgIdentifier: %[3]s
      tags: {}
      spec:
        type: Run
        spec:
          shell: Sh
          command: echo %[5]s
  EOT
}
`, id, name, orgID, projectID, version, isStable)
}

func testAccTemplateImportProjectTwoVersions(id, name, orgID, projectID, stableVersion, secondVersion string, firstIsStable bool) string {
	return fmt.Sprintf(`
resource "harness_platform_template" "test" {
  identifier   = "%[1]s"
  name         = "%[2]s"
  org_id       = "%[3]s"
  project_id   = "%[4]s"
  version      = "%[5]s"
  is_stable    = %[7]t
  force_delete = true
  template_yaml = <<-EOT
    template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: %[5]s
      type: Step
      projectIdentifier: %[4]s
      orgIdentifier: %[3]s
      tags: {}
      spec:
        type: Run
        spec:
          shell: Sh
          command: echo %[5]s
  EOT
}

resource "harness_platform_template" "test_v2" {
  identifier   = "%[1]s"
  name         = "%[2]s"
  org_id       = "%[3]s"
  project_id   = "%[4]s"
  version      = "%[6]s"
  is_stable    = false
  force_delete = true
  depends_on   = [harness_platform_template.test]
  template_yaml = <<-EOT
    template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: %[6]s
      type: Step
      projectIdentifier: %[4]s
      orgIdentifier: %[3]s
      tags: {}
      spec:
        type: Run
        spec:
          shell: Sh
          command: echo %[6]s
  EOT
}
`, id, name, orgID, projectID, stableVersion, secondVersion, firstIsStable)
}

func testAccTemplateImportOrgSingleVersion(id, name, orgID, version string, isStable bool) string {
	return fmt.Sprintf(`
resource "harness_platform_template" "test" {
  identifier   = "%[1]s"
  name         = "%[2]s"
  org_id       = "%[3]s"
  version      = "%[4]s"
  is_stable    = %[5]t
  force_delete = true
  template_yaml = <<-EOT
    template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: %[4]s
      type: Step
      orgIdentifier: %[3]s
      tags: {}
      spec:
        type: Run
        spec:
          shell: Sh
          command: echo %[4]s
  EOT
}
`, id, name, orgID, version, isStable)
}

func testAccTemplateImportOrgTwoVersions(id, name, orgID, stableVersion, secondVersion string, firstIsStable bool) string {
	return fmt.Sprintf(`
resource "harness_platform_template" "test" {
  identifier   = "%[1]s"
  name         = "%[2]s"
  org_id       = "%[3]s"
  version      = "%[4]s"
  is_stable    = %[6]t
  force_delete = true
  template_yaml = <<-EOT
    template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: %[4]s
      type: Step
      orgIdentifier: %[3]s
      tags: {}
      spec:
        type: Run
        spec:
          shell: Sh
          command: echo %[4]s
  EOT
}

resource "harness_platform_template" "test_v2" {
  identifier   = "%[1]s"
  name         = "%[2]s"
  org_id       = "%[3]s"
  version      = "%[5]s"
  is_stable    = false
  force_delete = true
  depends_on   = [harness_platform_template.test]
  template_yaml = <<-EOT
    template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: %[5]s
      type: Step
      orgIdentifier: %[3]s
      tags: {}
      spec:
        type: Run
        spec:
          shell: Sh
          command: echo %[5]s
  EOT
}
`, id, name, orgID, stableVersion, secondVersion, firstIsStable)
}

func testAccTemplateImportAccountSingleVersion(id, name, version string, isStable bool) string {
	return fmt.Sprintf(`
resource "harness_platform_template" "test" {
  identifier   = "%[1]s"
  name         = "%[2]s"
  version      = "%[3]s"
  is_stable    = %[4]t
  force_delete = true
  template_yaml = <<-EOT
    template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: %[3]s
      type: Step
      tags: {}
      spec:
        type: Run
        spec:
          shell: Sh
          command: echo %[3]s
  EOT
}
`, id, name, version, isStable)
}

func testAccTemplateImportAccountTwoVersions(id, name, stableVersion, secondVersion string, firstIsStable bool) string {
	return fmt.Sprintf(`
resource "harness_platform_template" "test" {
  identifier   = "%[1]s"
  name         = "%[2]s"
  version      = "%[3]s"
  is_stable    = %[5]t
  force_delete = true
  template_yaml = <<-EOT
    template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: %[3]s
      type: Step
      tags: {}
      spec:
        type: Run
        spec:
          shell: Sh
          command: echo %[3]s
  EOT
}

resource "harness_platform_template" "test_v2" {
  identifier   = "%[1]s"
  name         = "%[2]s"
  version      = "%[4]s"
  is_stable    = false
  force_delete = true
  depends_on   = [harness_platform_template.test]
  template_yaml = <<-EOT
    template:
      name: "%[2]s"
      identifier: "%[1]s"
      versionLabel: %[4]s
      type: Step
      tags: {}
      spec:
        type: Run
        spec:
          shell: Sh
          command: echo %[4]s
  EOT
}
`, id, name, stableVersion, secondVersion, firstIsStable)
}
