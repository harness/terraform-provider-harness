package input_set_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceInputSet(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_input_set.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceInputSet(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "pipeline_id", id),
				),
			},
		},
	})
}

func testAccDataSourceInputSet(id string, name string) string {
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
            - name: key
              type: String
              default: value
              value: <+input>.allowedValues(value)
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
                    - name: "key"
                      type: "String"
                      value: "value"
            EOT
    }

            data "harness_platform_input_set" "test" {
                org_id = harness_platform_organization.test.id
                project_id = harness_platform_project.test.id
                identifier = harness_platform_input_set.test.identifier
                pipeline_id = harness_platform_pipeline.test.id
            }
    `, id, name)

}
