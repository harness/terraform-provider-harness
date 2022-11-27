resource "harness_platform_connector_pagerduty" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  delegate_selectors = ["harness-delegate"]
  api_token_ref      = "account.secret_id"
}
