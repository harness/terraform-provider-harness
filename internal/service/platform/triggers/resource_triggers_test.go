package triggers_test

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceTriggers(t *testing.T) {

	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)

	resourceName := "harness_platform_triggers.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccTriggersDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTriggers(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "target_id", id),
				),
			},
			{
				Config: testAccResourceTriggers(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "target_id", id),
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

func testAccGetTriggers(resourceName string, state *terraform.State) (*nextgen.NgTriggerResponse, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID
	orgId := r.Primary.Attributes["org_id"]
	projId := r.Primary.Attributes["project_id"]
	targetId := r.Primary.Attributes["target_id"]

	resp, _, err := c.TriggersApi.GetTrigger(ctx, c.AccountId, orgId, projId, targetId, id)

	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func buildField(r *terraform.ResourceState, field string) optional.String {
	if attr, ok := r.Primary.Attributes[field]; ok {
		return optional.NewString(attr)
	}
	return optional.EmptyString()
}

func testAccTriggersDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		trigger, _ := testAccGetTriggers(resourceName, state)
		if trigger != nil {
			return fmt.Errorf("Founr trigger: %s", trigger.Identifier)
		}

		return nil
	}
}

func testAccResourceTriggers(id string, name string) string {
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
	`, id, name)
}
