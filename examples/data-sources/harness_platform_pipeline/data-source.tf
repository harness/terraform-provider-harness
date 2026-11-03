data "harness_platform_pipeline" "example" {
  identifier = "identifier"
  org_id     = "org_id"
  project_id = "project_id"

  # Read a remote pipeline from the branch where it is stored when that branch
  # is not the repository default branch. Defaults to true.
  load_from_fallback_branch = true
}
