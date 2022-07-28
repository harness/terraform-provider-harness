package input_set_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceInputSet(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_input_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInputSetDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceInputSet(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					//TODO: add more tests here.
				),
			},
			{
				Config: testAccResourceInputSet(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.PipelineResourceImportStateIdFunc(resourceName),
			},
		},
	})

}

func testAccInputSetDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		inputSet, _ := testAccGetInputSet(resourceName, state)
		if inputSet != nil {
			return fmt.Errorf("Found input set: %s", inputSet.Identifier)
		}
		return nil
	}
}

func testAccGetInputSet(resourceName string, state *terraform.State) (*nextgen.InputSetResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	orgIdentifier := buildField(r, "org_id").Value()
	projectIdentifier := buildField(r, "project_id").Value()
	pipelineIdentifier := buildField(r, "pipeline_id").Value()
	resp, _, err := c.InputSetsApi.GetInputSet(ctx, id, c.AccountId, orgIdentifier, projectIdentifier, pipelineIdentifier,
		&nextgen.PipelineInputSetApiGetInputSetOpts{
			Branch:                  buildField(r, "branch"),
			RepoIdentifier:          buildField(r, "repo_identifier"),
			GetDefaultFromOtherRepo: buildBoolField(r, "get_default_from_other_repo", optional.EmptyBool()),
		})

	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, nil
	}

	return resp.Data, nil
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func buildBoolField(r *terraform.ResourceState, field string, def optional.Bool) optional.Bool {
	if attr, ok := r.Primary.Attributes[field]; ok {
		val, err := strconv.ParseBool(attr)
		if err == nil {
			return optional.NewBool(val)
		}
	}
	return def
}

func testAccResourceInputSet(id string, name string) string {
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

        resource "harness_platform_pipeline" "test" {
						identifier = "%[1]s"
						org_id = harness_platform_project.test.org_id
						project_id = harness_platform_project.test.id
						name = "%[2]s"
            yaml = <<-EOT
                pipeline:
                    identifier: %[1]s
                    name: %[2]s
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
                    variables:
                        - name: JDK_IMAGE
                          type: String
                          default: us.gcr.io/platform-205701/cie-agent-harness-core-jdk11:latest
                          value: <+input>.allowedValues(us.gcr.io/platform-205701/cie-agent-harness-core-jdk11:latest)
            EOT
        }

				resource "harness_platform_input_set" "test" {
						identifier = "%[1]s"
						name = "%[2]s"
						org_id = harness_platform_organization.test.id
						project_id = harness_platform_project.test.id
						pipeline_id = harness_platform_pipeline.test.id
						yaml = <<-EOT
                inputSet:
                  identifier: "%[1]s"
                  name: "%[2]s"
                  orgIdentifier: "${harness_platform_organization.test.id}"
                  projectIdentifier: "${harness_platform_project.test.id}"
                  pipeline:
                    identifier: "${harness_platform_pipeline.test.id}"
                    variables:
                    - name: "JDK_IMAGE"
                      type: "String"
                      value: "us.gcr.io/platform-205701/cie-agent-harness-core-jdk11:latest"
            EOT
				}
        `, id, name)
}
