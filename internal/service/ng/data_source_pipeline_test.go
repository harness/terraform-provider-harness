package ng_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePipeline(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	orgId := "testOrg"
	projId := "testProj"
	resourceName := "data.harness_pipeline.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePipeline(id, orgId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", orgId),
					resource.TestCheckResourceAttr(resourceName, "project_id", projId),
				),
			},
		},
	})
}

func testAccDataSourcePipeline(id string, orgId string, projId string) string {
	return fmt.Sprintf(`
		resource "harness_pipeline" "test" {
			pipeline_yaml = <<-EOT
				pipeline:
					name: %[1]s
					identifier: %[1]s
					allowStageExecutions: false
					projectIdentifier: %[3]s
					orgIdentifier: %[2]s
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
	`, id, orgId, projId)
}
