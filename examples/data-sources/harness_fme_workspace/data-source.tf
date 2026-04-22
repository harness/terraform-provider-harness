# By Harness organization and project identifiers (matches harness_platform_* conventions).
data "harness_fme_workspace" "by_project" {
  org_id     = "organization_id"
  project_id = "project_id"
}

# By exact Split workspace name.
data "harness_fme_workspace" "by_name" {
  name = "my-workspace-name"
}
