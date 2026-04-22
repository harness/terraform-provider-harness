# Default Split workspaces include a "user" traffic type.
data "harness_fme_traffic_type" "user" {
  org_id     = "organization_id"
  project_id = "project_id"
  name       = "user"
}
