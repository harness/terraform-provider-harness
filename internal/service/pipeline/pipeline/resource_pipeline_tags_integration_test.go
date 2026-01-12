package pipeline_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/harness/terraform-provider-harness/internal/acctest"
)

// TestAccResourcePipeline_TagsWithColons tests the fix for PIPE-30810
// This ensures that pipeline tags containing colons (like Harness expressions)
// are properly handled by the Terraform provider
func TestAccResourcePipeline_TagsWithColons(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id

	resourceName := "harness_platform_pipeline.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPipelineDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineWithComplexTags(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Verify that tags with colons are preserved
					resource.TestCheckResourceAttr(resourceName, "tags.#", "5"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// TestAccResourcePipeline_TagsWithHarnessExpressions tests Harness dynamic expressions
func TestAccResourcePipeline_TagsWithHarnessExpressions(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id

	resourceName := "harness_platform_pipeline.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccPipelineDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipelineWithHarnessExpressions(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
		},
	})
}

// testAccResourcePipelineWithComplexTags creates a pipeline with various tag formats
// This tests the fix for PIPE-30810 where tags with colons were being truncated
func testAccResourcePipelineWithComplexTags(id string, name string) string {
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
            
            # Various tag formats to test PIPE-30810 fix
            tags = [
                "env:prod",
                "version:1.0.0",
                "registry:https://registry.example.com:5000/repo",
                "timestamp:2024:01:15:10:30:00",
                "service:api-gateway"
            ]
            
            yaml = <<-EOT
                pipeline:
                    name: %[2]s
                    identifier: %[1]s
                    allowStageExecutions: false
                    projectIdentifier: ${harness_platform_project.test.id}
                    orgIdentifier: ${harness_platform_project.test.org_id}
                    tags:
                        env: prod
                        version: "1.0.0"
                        registry: "https://registry.example.com:5000/repo"
                        timestamp: "2024:01:15:10:30:00"
                        service: api-gateway
                    stages:
                        - stage:
                            name: S1
                            identifier: S1
                            description: ""
                            type: Custom
                            spec:
                              execution:
                                steps:
                                  - step:
                                      type: ShellScript
                                      name: ShellScript_1
                                      identifier: ShellScript_1
                                      spec:
                                        shell: Bash
                                        executionTarget: {}
                                        source:
                                          type: Inline
                                          spec:
                                            script: echo "test"
                                        environmentVariables: []
                                        outputVariables: []
                                      timeout: 10m
                            tags: {}
            EOT
        }
    `, id, name)
}

// testAccResourcePipelineWithHarnessExpressions creates a pipeline with Harness dynamic expressions
// This is the exact scenario from PIPE-30810 (National Australia Bank)
func testAccResourcePipelineWithHarnessExpressions(id string, name string) string {
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
            
            # PIPE-30810: Harness expressions with ternary operators containing colons
            tags = [
                "env:prod",
                "ImagePush:<+<+pipeline.variables.variableA>==\"yes\"?<+pipeline.variables.tag>:<+pipeline.name>>",
                "BuildNumber:<+pipeline.sequenceId>",
                "GitBranch:<+trigger.branch>:<+trigger.commitSha>"
            ]
            
            yaml = <<-EOT
                pipeline:
                    name: %[2]s
                    identifier: %[1]s
                    allowStageExecutions: false
                    projectIdentifier: ${harness_platform_project.test.id}
                    orgIdentifier: ${harness_platform_project.test.org_id}
                    tags:
                        env: prod
                        ImagePush: "<+<+pipeline.variables.variableA>==\"yes\"?<+pipeline.variables.tag>:<+pipeline.name>>"
                        BuildNumber: "<+pipeline.sequenceId>"
                        GitBranch: "<+trigger.branch>:<+trigger.commitSha>"
                    variables:
                        - name: variableA
                          type: String
                          description: "Test variable for ternary expression"
                          value: "yes"
                        - name: tag
                          type: String
                          description: "Docker tag"
                          value: "latest"
                    stages:
                        - stage:
                            name: S1
                            identifier: S1
                            description: ""
                            type: Custom
                            spec:
                              execution:
                                steps:
                                  - step:
                                      type: ShellScript
                                      name: ShellScript_1
                                      identifier: ShellScript_1
                                      spec:
                                        shell: Bash
                                        executionTarget: {}
                                        source:
                                          type: Inline
                                          spec:
                                            script: echo "test"
                                        environmentVariables: []
                                        outputVariables: []
                                      timeout: 10m
                            tags: {}
            EOT
        }
    `, id, name)
}
