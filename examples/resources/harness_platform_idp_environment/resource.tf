resource "harness_platform_idp_environment" "test" {
  identifier           = "identifier"
  org_id               = "org_id"
  project_id           = "project_id"
  name                 = "test"
  owner                = "group:account/_account_all_users"
  blueprint_identifier = "blueprint_identifier"
  blueprint_version    = "v1.0.0"
  target_state         = "running"
  based_on             = "org_id.project_id/environment_identifier"
  overrides            = <<-EOT
    config: {}
    entities: {}
    EOT
  inputs               = <<-EOT
    ttl: 1h30m
    EOT
}
