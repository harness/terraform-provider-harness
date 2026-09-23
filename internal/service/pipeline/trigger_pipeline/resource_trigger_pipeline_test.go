package trigger_pipeline_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccResourceTriggerPipeline_WaitAndCaptureOutputs triggers a real pipeline run (a Custom
// stage with a ShellScript step that exports an output variable), then flips
// wait_for_completion on via an in-place update and asserts the resource waits for the
// already-triggered execution (same id) instead of starting a new one, and that the
// captured status/outputs reflect the completed run.
func TestAccResourceTriggerPipeline_WaitAndCaptureOutputs(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_platform_trigger_pipeline.test"

	var firstId string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Fire-and-forget: trigger the run but don't wait.
				Config: testAccResourceTriggerPipelineConfig(id, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "plan_execution_id"),
					resource.TestCheckResourceAttr(resourceName, "wait_for_completion", "false"),
					resource.TestCheckResourceAttrWith(resourceName, "id", func(v string) error {
						firstId = v
						return nil
					}),
				),
			},
			{
				// Flip wait_for_completion on; this is an in-place update, not a replace, so it
				// must observe the execution already started in the previous step.
				Config: testAccResourceTriggerPipelineConfig(id, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith(resourceName, "id", func(v string) error {
						if v != firstId {
							return fmt.Errorf("expected wait_for_completion update to keep id %q, got %q", firstId, v)
						}
						return nil
					}),
					resource.TestCheckResourceAttr(resourceName, "status", "Success"),
					resource.TestMatchResourceAttr(resourceName, "outputs", regexp.MustCompile(`hello world`)),
				),
			},
		},
	})
}

func testAccResourceTriggerPipelineConfig(id string, waitForCompletion bool) string {
	return fmt.Sprintf(`
resource "harness_platform_organization" "test" {
  identifier = "%[1]s"
  name       = "%[1]s"
}

resource "harness_platform_project" "test" {
  identifier = "%[1]s"
  name       = "%[1]s"
  org_id     = harness_platform_organization.test.id
  color      = "#472848"
}

resource "harness_platform_pipeline" "test" {
  identifier = "%[1]s"
  name       = "%[1]s"
  org_id     = harness_platform_project.test.org_id
  project_id = harness_platform_project.test.id

  yaml = <<-EOT
    pipeline:
      name: %[1]s
      identifier: %[1]s
      orgIdentifier: ${harness_platform_project.test.org_id}
      projectIdentifier: ${harness_platform_project.test.id}
      tags: {}
      stages:
        - stage:
            name: custom
            identifier: custom
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
                            script: export greeting="hello world"
                        environmentVariables: []
                        outputVariables:
                          - name: greeting
                            type: String
                            value: greeting
                        timeout: 2m
  EOT
}

resource "harness_platform_trigger_pipeline" "test" {
  org_id      = harness_platform_project.test.org_id
  project_id  = harness_platform_project.test.id
  pipeline_id = harness_platform_pipeline.test.id

  wait_for_completion   = %[2]t
  poll_interval_seconds = 2
}
`, id, waitForCompletion)
}
