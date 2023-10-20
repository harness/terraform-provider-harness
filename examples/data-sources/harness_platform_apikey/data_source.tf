data "harness_platform_apikey" "test" {
  identifier  = "test_apikey"
  name        = "test_apikey"
  parent_id   = "parent_id"
  apikey_type = "USER"
  account_id  = "account_id"
  org_id      = "org_id"
  project_id  = "project_id"
}
