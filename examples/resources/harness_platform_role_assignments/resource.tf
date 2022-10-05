resource "harness_platform_role_assignments" "example" {
  identifier                = "identifier"
  org_id                    = "org_id"
  project_id                = "project_id"
  resource_group_identifier = "_all_project_level_resources"
  role_identifier           = "_project_viewer"
  principal {
    identifier = harness_platform_service_account.test.id
    type       = "SERVICE_ACCOUNT"
  }
  disabled = false
  managed  = false
}
