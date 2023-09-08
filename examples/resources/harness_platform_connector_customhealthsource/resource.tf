# Authentication mechanism as username and password
resource "harness_platform_connector_customhealthsource" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://prometheus.com/"
  delegate_selectors = ["harness-delegate"]
  method             = "GET"
  validation_path    = "loki/api/v1/labels"
  headers {
    encrypted_value_ref = "account.doNotDeleteHSM"
    value_encrypted     = true
    key                 = "key"
    value               = "value"
  }
}