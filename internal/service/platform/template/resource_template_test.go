package template_test

import (
	"fmt"
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
		CheckDestroy:      testAccTemplateDestroy(resourceName),
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
		CheckDestroy:      testAccTemplateDestroy(resourceName),
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
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id

	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTemplateDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTemplateOrgScopeImportFromGit(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", "orgtemplate"),
					resource.TestCheckResourceAttr(resourceName, "name", "orgtemplate"),
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
		CheckDestroy:      testAccTemplateDestroy(resourceName),
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

func TestAccResourceTemplate_AccountScopeImportFromGit(t *testing.T) {
	id := "accounttemplate"
	name := id

	resourceName := "harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTemplateDestroy(resourceName),
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
		CheckDestroy:      testAccTemplateDestroy(resourceName),
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
		CheckDestroy:      testAccTemplateDestroy(resourceName),
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
				connector_ref = "account.Jajoo"
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
				connector_ref = "account.Jajoo"
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
				connector_ref = "account.Jajoo"
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
				connector_ref = "account.Jajoo"
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

func testAccResourceTemplateOrgScopeImportFromGit(id string, name string) string {
	return fmt.Sprintf(`
        resource "harness_platform_organization" "test" {
					identifier = "%[1]sorg"
					name = "%[2]s"
				}
        resource "harness_platform_template" "test" {
                        identifier = "orgtemplate"
                        org_id = "default"
                        name = "orgtemplate"
						version = "v2"
                        import_from_git = true
                        git_import_details {
                            branch_name = "main"
                            file_path = ".harness/orgtemplate.yaml"
                            connector_ref = "account.DoNotDeleteGithub"
                            repo_name = "open-repo"
                        }
                        template_import_request {
                            template_name = "orgtemplate"
							template_version = "v2"
                            template_description = ""
                        }
                }
        `, id, name)
}

func testAccResourceTemplateProjectScopeImportFromGit(id string, name string) string {
	return fmt.Sprintf(`
        resource "harness_platform_organization" "test" {
					identifier = "%[1]s"
					name = "%[2]s"
				}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			color = "#472848"
		}
        resource "harness_platform_template" "test" {
                        identifier = "%[1]s"
                        org_id = harness_platform_project.test.id
						project_id = harness_platform_organization.test.id
                        name = "%[2]s"
						version = "v2"
                        import_from_git = true
                        git_import_details {
                            branch_name = "main"
                            file_path = ".harness/projecttemplate.yaml"
                            connector_ref = "account.DoNotDeleteGithub"
                            repo_name = "open-repo"
                        }
                        template_import_request {
                            template_name = "%[2]s"
							template_version = "v2"
                            template_description = ""
                        }
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
                            connector_ref = "account.DoNotDeleteGithub"
                            repo_name = "open-repo"
                        }
                        template_import_request {
                            template_name = "accounttemplate"
							template_version = "v2"
                            template_description = ""
                        }
                }
        `, id, name)
}
