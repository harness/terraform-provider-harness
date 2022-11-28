resource "harness_platform_connector_splunk" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://splunk.com/"
  delegate_selectors = ["harness-delegate"]
  account_id         = "splunk_account_id"
  username           = "username"
  password_ref       = "account.secret_id"
}
