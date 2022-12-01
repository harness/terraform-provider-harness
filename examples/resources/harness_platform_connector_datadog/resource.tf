resource "harness_platform_connector_datadog" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                 = "https://datadog.com"
  delegate_selectors  = ["harness-delegate"]
  application_key_ref = "account.secret_id"
  api_key_ref         = "account.secret_id"
}
