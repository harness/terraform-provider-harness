resource "harness_platform_connector_gcp_secret_manager" "gcp_sm" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  delegate_selectors = ["harness-delegate"]
  credentials_ref    = "account.${harness_platform_secret_text.test.id}"
}
