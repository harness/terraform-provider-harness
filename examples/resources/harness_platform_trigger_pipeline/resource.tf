 terraform {
   required_providers {
     harness = {
       source = "registry.terraform.io/harness/harness"
     }
   }
 }

resource "harness_platform_trigger_pipeline" "scs_ecs_timing_test" {
  org_id      = "default"
  project_id  = "rssnyder"
  pipeline_id = "SCS_ECS_timing_test_local"

  # Fills the pipeline's runtime inputs (codebase build, connector, stage variables).
  input_set_yaml = <<-EOT
    pipeline:
      identifier: SCS_ECS_timing_test_local
      properties:
        ci:
          codebase:
            build:
              type: branch
              spec:
                branch: master
      stages:
        - stage:
            identifier: Build_n_Scan
            type: CI
            spec:
              execution:
                steps:
                  - step:
                      identifier: BuildAndPushECR
                      type: BuildAndPushECR
                      spec:
                        connectorRef: account.harness_impeng_play
            variables:
              - name: image_name
                type: String
                value: harness-private
              - name: tag
                type: String
                value: latest
              - name: ecr_account
                type: String
                value: "664418987337"
              - name: region
                type: String
                value: us-west-2
  EOT

  wait_for_completion   = true
  poll_interval_seconds = 15
}

# Status of the run: Success, Failed, Aborted, etc.
output "pipeline_status" {
  value = harness_platform_trigger_pipeline.scs_ecs_timing_test.status
}

# Identifier of the execution, useful for linking back to the Harness UI.
output "pipeline_execution_id" {
  value = harness_platform_trigger_pipeline.scs_ecs_timing_test.plan_execution_id
}

# JSON-encoded map of output variables/outcomes exported by the pipeline execution.
output "pipeline_context" {
  value = jsondecode(harness_platform_trigger_pipeline.scs_ecs_timing_test.outputs)
}
