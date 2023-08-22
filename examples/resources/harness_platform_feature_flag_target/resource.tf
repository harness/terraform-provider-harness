// Boolean Flag
resource "harness_platform_feature_flag_target" "target" {
  org_id     = "test"
  project_id = "test"

  identifier  = "MY_FEATURE"
  environment = "MY_ENVIRONMENT"
  name        = "MY_FEATURE"
  account_id  = "MY_ACCOUNT_ID"
  attributes  = { "foo" : "bar" }
}
