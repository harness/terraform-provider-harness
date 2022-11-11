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
)

func TestAccResourceTemplate(t *testing.T) {
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
				Config: testAccResourceTemplate(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAccResourceTemplate(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
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
				),
			},
			{
				Config: testAccResourceTemplateOrgScope(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
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
		if template != nil {
			return fmt.Errorf("Found template: %s", template.Template.Slug)
		}

		return nil
	}
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
				connector_ref = "account.RichaGithub"
				store_type = "REMOTE"
				repo_name = "rjajoo"
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

func testAccResourceTemplate(id string, name string) string {
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
			org_id = harness_platform_project.test.org_id
			project_id = harness_platform_project.test.id
			name = "%[2]s"
			comments = "comments"
			version = "ab"
			is_stable = true
			git_details {
				branch_name = "main"
				commit_message = "Commit"
				file_path = ".harness/GitEnabledPipeline%[1]s.yaml"
				connector_ref = "account.RichaGithub"
				store_type = "REMOTE"
				repo_name = "rjajoo"
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
