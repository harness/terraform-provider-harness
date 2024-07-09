resource "harness_platform_iacm_default_pipeline" "example" {
  org_id                  = harness_platform_organization.test.id
  project_id              = harness_platform_project.test.id
  provisioner_type        = "terraform"
  operation               = "plan"
  pipeline                = "pipeline1"
}
