resource "harness_chaos_image_registry" "example" {
  org_id         = "<org_id>"
  project_id     = "<project_id>"
  
  registry_server        = "<registry_server>"
  registry_account       = "<registry_account>"
  is_private             = true
  secret_name            = "<secret_name>"
  is_default             = false
  is_override_allowed    = true
  use_custom_images      = true

  # Custom images configuration
  custom_images {
    log_watcher = "<log_watcher_image>"
    ddcr        = "<ddcr_image>"
    ddcr_lib    = "<ddcr_lib_image>"
    ddcr_fault  = "<ddcr_fault_image>"
  }
}
