resource "harness_platform_pipeline_filters" "test" {
  identifier = "identifier"
  name       = "name"
  org_id     = "org_id"
  project_id = "project_id"
  type       = "PipelineExecution"
  filter_properties {
    tags        = ["foo:bar"]
    filter_type = "PipelineExecution"
  }
  filter_visibility = "EveryOne"
}
