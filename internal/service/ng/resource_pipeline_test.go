package ng_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourcePipeline(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	updatedId := fmt.Sprintf("%s_updated", id)
	orgId := "testOrg"
	projId := "testProj"
	resourceName := "harness_pipeline.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPipelineDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipeline(id, orgId, projId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", id),
				),
			},
			{
				Config: testAccResourcePipeline(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", updatedId),
					resource.TestCheckResourceAttr(resourceName, "name", updatedId),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					id := primary.ID
					org_id := primary.Attributes["org_id"]
					return fmt.Sprintf("%s/%s", org_id, id), nil
				},
			},
		},
	})
}

func testAccGetPipeline(resourceName string, state *terraform.State) (*nextgen.Pipeline, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c := acctest.TestAccGetApiClientFromProvider()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]

	resp, _, err := c.NGClient.PipelineApi.GetPipeline(context.Background(), c.AccountId, orgId, projId, id, &nextgen.PipelineApiGetPipelineOpts{})
	if err != nil {
		return nil, err
	}

	return resp.Data.Pipeline, nil
}

func testAccPipelineDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		pipeline, _ := testAccGetPipeline(resourceName, state)
		if pipeline != nil {
			return fmt.Errorf("Found pipeline: %s", pipeline.Identifier)
		}

		return nil
	}
}

func testAccResourcePipeline(id string, orgId string, projId string) string {
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
