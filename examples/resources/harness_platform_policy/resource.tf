resource "harness_platform_policy" "test" {
  identifier        = harness_platform_policy.test.identifier
  name              = harness_platform_policy.test.name
  description       = harness_platform_policy.test.description
  org_id            = harness_platform_policy.test.org_id
  project_id        = harness_platform_policy.test.project_id
  git_connector_ref = harness_platform_policy.test.git_connector_ref
  git_path          = harness_platform_policy.test.git_path
  git_repo          = harness_platform_policy.test.git_repo
  git_branch        = harness_platform_policy.test.git_branch
  git_base_branch   = harness_platform_policy.test.git_base_branch
  git_is_new_branch = false
  git_import        = false
  git_commit_msg    = harness_platform_policy.test.git_commit_msg
  rego              = <<-REGO
    package pipeline

    # Deny pipelines that don't have an approval step
    # NOTE: Try removing the HarnessApproval step from your input to see the policy fail
    deny[msg] {
        # Find all stages that are Deployments ...
        input.pipeline.stages[i].stage.type == "Approval"

        # ... that are not in the set of stages with HarnessApproval steps
        not stages_with_approval[i]

        # Show a human-friendly error message
        msg := sprintf("Approval stage '%s' does not have a HarnessApproval step", [input.pipeline.stages[i].stage.name])
    }

    # Find the set of stages that contain a HarnessApproval step
    stages_with_approval[i] {
        input.pipeline.stages[i].stage.spec.execution.steps[_].step.type == "HarnessApproval"
    }
REGO
}
