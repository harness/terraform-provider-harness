resource "harness_platform_iacm_workspace_template" "example" {
  org_id       = harness_platform_organization.test.id
  project_id   = harness_platform_project.test.id
  workspace_id = harness_platform_workspace.example.identifier
  template_id  = "my_template"
  version      = "v1.0.0"
}
