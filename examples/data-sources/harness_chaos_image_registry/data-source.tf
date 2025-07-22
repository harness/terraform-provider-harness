# Data source to verify the registry
data "harness_chaos_image_registry" "example" {
  org_id     = "<org_id>"
  project_id = "<project_id>"
}

# Example of checking override status
data "harness_chaos_image_registry" "override_check" {
  org_id     = "<org_id>"
  project_id = "<project_id>"
  check_override = true
}
