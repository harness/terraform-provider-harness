# Credentials http
resource "harness_platform_connector_bitbucket" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://bitbucket.com/account"
  connection_type    = "Account"
  validation_repo    = "some_repo"
  delegate_selectors = ["harness-delegate"]
  credentials {
    http {
      username     = "username"
      password_ref = "account.secret_id"
    }
  }

  api_authentication {
    username  = "username"
    token_ref = "account.secret_id"
  }
}

# Credentials ssh
resource "harness_platform_connector_bitbucket" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://bitbucket.com/account"
  connection_type    = "Account"
  validation_repo    = "some_repo"
  delegate_selectors = ["harness-delegate"]
  credentials {
    ssh {
      ssh_key_ref = "account.secret_id"
    }
  }
}
