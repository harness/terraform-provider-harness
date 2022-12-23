resource "harness_platform_user" "example" {
  org_id     = "org_id"
  project_id = "project_id"
  emails     = ["john.doe@harness.io"]
  role_bindings {
    resource_group_identifier = "_all_project_level_resources"
    role_identifier           = "_project_viewer"
    role_name                 = "Project Viewer"
    resource_group_name       = "All Project Level Resources"
    managed_role              = true
  }
}
