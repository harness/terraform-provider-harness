
resource "harness_pipeline" "terraform_pipeline_example" {
  identifier  = "terraform_pipeline_example"
  org_id      = "default"
  project_id  = "First"
  pipeline_yaml = file("./example-pipeline.yaml")
}
