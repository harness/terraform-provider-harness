resource "harness_platform_connector_jira" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://jira.com"
  delegate_selectors = ["harness-delegate"]
  username           = "admin"
  password_ref       = "account.secret_id"
}
