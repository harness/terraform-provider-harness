resource "harness_platform_connector_dynatrace" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://dynatrace.com/"
  delegate_selectors = ["harness-delegate"]
  api_token_ref      = "account.secret_id"
}
