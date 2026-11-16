resource "harness_platform_gitops_agent_dr_promotion" "example" {
  org_id               = "org_id"
  project_id           = "project_id"
  agent_id             = "agent_id"
  dr_identifier        = "dr_identifier"
  disable_sync_windows = false

  triggers = {
    promoted_at = "2024-01-01T00:00:00Z"
  }
}
