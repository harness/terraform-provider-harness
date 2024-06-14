# Credentials http
resource "harness_platform_connector_gitlab" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://gitlab.com/account"
  connection_type    = "Account"
  validation_repo    = "some_repo"
  delegate_selectors = ["harness-delegate"]
  credentials {
    http {
      username  = "username"
      token_ref = "account.secret_id"
    }
  }
}

resource "harness_platform_connector_gitlab" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://gitlab.com/account"
  connection_type    = "Account"
  validation_repo    = "some_repo"
  delegate_selectors = ["harness-delegate"]
  credentials {
    http {
      username  = "username"
      token_ref = "account.secret_id"
    }
  }
  api_authentication {
    token_ref = "account.secret_id"
  }
}

# Credentials ssh
resource "harness_platform_connector_gitlab" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  url                = "https://gitlab.com/account"
  connection_type    = "Account"
  validation_repo    = "some_repo"
  delegate_selectors = ["harness-delegate"]
  credentials {
    ssh {
      ssh_key_ref = "account.test"
    }
  }
}
