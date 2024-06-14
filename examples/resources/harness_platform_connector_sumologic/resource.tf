resource "harness_platform_connector_sumologic" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://api.us2.sumologic.com/"
  delegate_selectors = ["harness-delegate"]
  access_id_ref      = "account.secret_id"
  access_key_ref     = "account.secret_id"
}
