resource "harness_platform_feature_flag_target_group" "target" {
  org_id     = "test"
  project_id    = "test"

  identifier  = "MY_FEATURE"
  environment = "MY_ENVIRONMENT"
  name        = "MY_FEATURE"
  account_id  = "MY_ACCOUNT_ID"
  included    = ["target_id_1"]
  excluded    = ["target_id_2"]
  rule        =  {
    attribute = "MY_ATTRIBUTE"
    operator  = "EQUALS"
    value     = "MY_VALUE"
  }              
}