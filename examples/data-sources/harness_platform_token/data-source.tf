data "harness_platform_token" "test" {
  identifier  = "test_token"
  parent_id   = "apikey_parent_id"
  org_id      = "org_id"
  project_id  = "project_id"
  apikey_id   = "apikey_id"
  apikey_type = "USER"
}