package triggers_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTriggers(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(6))
	name := id
	resourceName := "data.harness_platform_triggers.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTriggers(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "org_id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "project_id", id),
					resource.TestCheckResourceAttr(resourceName, "target_id", id),
				),
			},
		},
	})
}

func testAccDataSourceTriggers(id string, name string) string {
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

resource "harness_platform_pipeline" "pipeline" {
	identifier = "%[1]s"
	org_id = harness_platform_project.test.org_id
	project_id = harness_platform_project.test.id
	name = "%[2]s"
  yaml = <<-EOT
  pipeline:
    name: %[2]s
    identifier: %[1]s
    projectIdentifier:  ${harness_platform_project.test.id}
    orgIdentifier:  ${harness_platform_project.test.org_id}
    tags: {}
    stages:
        - stage:
              name: Approval
              identifier: Approval
              description: ""
              type: Approval
              spec:
                  execution:
                      steps:
                          - step:
                                name: Approval Step
                                identifier: Approval_Step
                                type: HarnessApproval
                                timeout: 1d
                                spec:
                                    approvalMessage: Please review the following information and approve the pipeline progression
                                    includePipelineExecutionHistory: true
                                    approvers:
                                        minimumCount: 1
                                        disallowPipelineExecutor: false
                                        userGroups:
                                            - account.testmv
                                    approverInputs: []
              tags: {}
EOT
}

	resource "harness_platform_triggers" "test" {
		identifier = "%[1]s"
		org_id = harness_platform_project.test.org_id
		project_id = harness_platform_project.test.id
		name = "%[2]s"
		target_id = harness_platform_pipeline.pipeline.id
		yaml = <<-EOT
    ---
    trigger:
      name: "%[2]s"
      identifier: "%[1]s"
      enabled: true
      description: ""
      tags: {}
      projectIdentifier: "${harness_platform_project.test.id}"
      orgIdentifier: "${harness_platform_project.test.org_id}"
      pipelineIdentifier: "${harness_platform_pipeline.pipeline.id}"
      source:
        type: "Webhook"
        pollInterval: "0"
        spec:
          type: "Github"
          spec:
            type: "Push"
            spec:
              connectorRef: "account.Jajoo"
              autoAbortPreviousExecutions: false
              payloadConditions:
              - key: "changedFiles"
                operator: "Equals"
                value: "fjjfjfjf"
              - key: "targetBranch"
                operator: "Equals"
                value: "fhfhfh"
              headerConditions: []
              repoName: "gfgfgf"
              actions: []
      inputYaml: "pipeline: {}\n"
      EOT
	}

	data "harness_platform_triggers" "test" {
    identifier = harness_platform_triggers.test.id
		org_id = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		target_id = harness_platform_pipeline.pipeline.id
	}
	`, id, name)
}
