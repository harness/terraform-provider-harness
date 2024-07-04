resource "harness_platform_connector_jdbc" "test" {
  identifier         = "identifer"
  name               = "name"
  description        = "test"
  tags               = ["foo:bar"]
  url                = "jdbc:sqlserver://1.2.3;trustServerCertificate=true"
  delegate_selectors = ["harness-delegate"]
  credentials {
    username     = "admin"
    password_ref = "account.secret_id"
  }
}