resource "harness_platform_ccm_perspective" "test" {
  identifier   = "identifier"
  account_id   = "account_id"
  name         = "name"
  clone        = false
  view_version = "v1"
  view_state   = "DRAFT"
  view_type    = "CUSTOMER"
}
