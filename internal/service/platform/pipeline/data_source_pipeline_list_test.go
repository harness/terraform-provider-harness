package pipeline_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcepipelineList(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_pipeline_list.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcepipelineList(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "pipelines.0.identifier", id),
					resource.TestCheckResourceAttr(resourceName, "pipelines.0.name", name),
				),
			},
		},
	})
}

func testAccDataSourcepipelineList(id string, name string) string {
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
        resource "harness_platform_pipeline" "test" {
                        identifier = "%[1]s"
                        org_id = harness_platform_project.test.org_id
                        project_id = harness_platform_project.test.id
                        name = "%[2]s"
                        git_details {
                            branch_name = "main"
                            commit_message = "Commit"
                            file_path = ".harness/GitEnabledPipeline%[1]s.yaml"
                            connector_ref = "account.github_Account_level_connector_delegate"
                            store_type = "REMOTE"
                            repo_name = "Gitx-Automation"
                        }
            yaml = <<-EOT
                pipeline:
                    name: %[2]s
                    identifier: %[1]s
                    allowStageExecutions: false
                    projectIdentifier: ${harness_platform_project.test.id}
                    orgIdentifier: ${harness_platform_project.test.org_id}
                    tags: {}
                    stages:
                        - stage:
                            name: dep
                            identifier: dep
                            description: ""
                            type: Deployment
                            spec:
                                serviceConfig:
                                    serviceRef: service
                                    serviceDefinition:
                                        type: Kubernetes
                                        spec:
                                            variables: []
                                infrastructure:
                                    environmentRef: testenv
                                    infrastructureDefinition:
                                        type: KubernetesDirect
                                        spec:
                                            connectorRef: testconf
                                            namespace: test
                                            releaseName: release-<+INFRA_KEY>
                                    allowSimultaneousDeployments: false
                                execution:
                                    steps:
                                        - stepGroup:
                                                name: Canary Deployment
                                                identifier: canaryDepoyment
                                                steps:
                                                    - step:
                                                        name: Canary Deployment
                                                        identifier: canaryDeployment
                                                        type: K8sCanaryDeploy
                                                        timeout: 10m
                                                        spec:
                                                            instanceSelection:
                                                                type: Count
                                                                spec:
                                                                    count: 1
                                                            skipDryRun: false
                                                    - step:
                                                        name: Canary Delete
                                                        identifier: canaryDelete
                                                        type: K8sCanaryDelete
                                                        timeout: 10m
                                                        spec: {}
                                                rollbackSteps:
                                                    - step:
                                                        name: Canary Delete
                                                        identifier: rollbackCanaryDelete
                                                        type: K8sCanaryDelete
                                                        timeout: 10m
                                                        spec: {}
                                        - stepGroup:
                                                name: Primary Deployment
                                                identifier: primaryDepoyment
                                                steps:
                                                    - step:
                                                        name: Rolling Deployment
                                                        identifier: rollingDeployment
                                                        type: K8sRollingDeploy
                                                        timeout: 10m
                                                        spec:
                                                            skipDryRun: false
                                                rollbackSteps:
                                                    - step:
                                                        name: Rolling Rollback
                                                        identifier: rollingRollback
                                                        type: K8sRollingRollback
                                                        timeout: 10m
                                                        spec: {}
                                    rollbackSteps: []
                            tags: {}
                            failureStrategies:
                                - onFailure:
                                        errors:
                                            - AllErrors
                                        action:
                                            type: StageRollback
            EOT
        }
        data "harness_platform_pipeline_list" "test" {
            org_id = harness_platform_pipeline.test.org_id
            project_id = harness_platform_pipeline.test.project_id
        }
    `, id, name)
}
