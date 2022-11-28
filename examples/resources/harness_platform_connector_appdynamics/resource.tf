# Authentication mechanism as api token
resource "harness_platform_connector_appdynamics" "token" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://appdynamics.com/"
  account_name       = "myaccount"
  delegate_selectors = ["harness-delegate"]
  api_token {
    client_id         = "client_id"
    client_secret_ref = "account.secret_id"
  }
}

# Authentication mechanism as username and password
resource "harness_platform_connector_appdynamics" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://appdynamics.com/"
  account_name       = "myaccount"
  delegate_selectors = ["harness-delegate"]
  username_password {
    username     = "username"
    password_ref = "account.secret_id"
  }
}
