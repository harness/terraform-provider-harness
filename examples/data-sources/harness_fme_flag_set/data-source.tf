# Name must match an existing flag set in the Split workspace for the Harness org/project.
data "harness_fme_flag_set" "example" {
  org_id     = "organization_id"
  project_id = "project_id"
  name       = "my-flag-set"
}
