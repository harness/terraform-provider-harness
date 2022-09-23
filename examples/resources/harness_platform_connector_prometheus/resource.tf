resource "harness_platform_connector_prometheus" "example" {
  identifier  = "idntifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://prometheus.com/"
  delegate_selectors = ["harness-delegate"]
  user_name          = "user_name"
  password_ref       = "account.secret_identifier"
  headers {
    encrypted_value_ref = "account.secret_identifier"
    value_encrypted     = true
    key                 = "key"
    value               = "value"
  }
}
