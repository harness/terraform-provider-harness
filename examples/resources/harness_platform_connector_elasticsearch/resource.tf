# Authentication mechanism as api token
resource "harness_platform_connector_elasticsearch" "token" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "http://elk6.dev.harness.io:9200/"
  delegate_selectors = ["harness-delegate"]
  api_token {
    client_id         = "client_id"
    client_secret_ref = "account.secret_id"
  }
}

# Authentication mechanism as username and password
resource "harness_platform_connector_elasticsearch" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "http://elk6.dev.harness.io:9200/"
  delegate_selectors = ["harness-delegate"]
  username_password {
    username     = "username"
    password_ref = "account.secret_id"
  }
}

# Authentication mechanism without authentication
resource "harness_platform_connector_elasticsearch" "no_authentication" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "http://elk6.dev.harness.io:9200/"
  delegate_selectors = ["harness-delegate"]
}