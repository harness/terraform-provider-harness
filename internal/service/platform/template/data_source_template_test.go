package template_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTemplate(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTemplate(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

func TestAccDataSourceTemplateInline(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTemplateInline(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "store_type", "INLINE"),
				),
			},
		},
	})
}

func TestAccDataSourceTemplateOrgScope(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTemplateOrgScope(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
				),
			},
		},
	})
}

func TestAccDataSourceTemplateProjectScope(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_template.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTemplateProjectScope(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
				),
			},
		},
	})
}

func testAccDataSourceTemplateInline(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_template" "test" {
			identifier = "%[1]s"
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

  data "harness_platform_template" "test" {
		identifier = harness_platform_template.test.id
	}
	`, id, name)
}

func testAccDataSourceTemplateProjectScope(id string, name string) string {
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

  data "harness_platform_template" "test" {
		identifier = harness_platform_template.test.id
    org_id = harness_platform_project.test.org_id
    project_id = harness_platform_project.test.id
	}
	`, id, name)
}

func testAccDataSourceTemplate(id string, name string) string {
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

	data "harness_platform_template" "test" {
		identifier = harness_platform_template.test.id
	}
	`, id, name)
}

func testAccDataSourceTemplateOrgScope(id string, name string) string {
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

  data "harness_platform_template" "test" {
		identifier = harness_platform_template.test.id
    org_id = harness_platform_organization.test.id
	}
	`, id, name)
}
