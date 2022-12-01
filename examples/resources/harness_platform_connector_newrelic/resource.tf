resource "harness_platform_connector_newrelic" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://newrelic.com/"
  delegate_selectors = ["harness-delegate"]
  account_id         = "nr_account_id"
  api_key_ref        = "account.secret_id"
}
