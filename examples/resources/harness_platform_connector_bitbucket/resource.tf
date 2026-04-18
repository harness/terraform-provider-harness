# Credentials http (with username + personal access token - UsernameToken)
resource "harness_platform_connector_bitbucket" "username_token" {
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

  # Defaults to auth_type = "UsernameToken" when omitted (backward compatible).
  api_authentication {
    auth_type = "UsernameToken"
    username  = "username"
    token_ref = "account.secret_id"
  }
}

# Credentials http with Bitbucket Cloud Workspace API Token (email + API token)
# Use this when migrating off Bitbucket app passwords (EOL 2026-06-09).
resource "harness_platform_connector_bitbucket" "email_api_token" {
  identifier  = "identifier_email_api_token"
  name        = "name_email_api_token"
  description = "Bitbucket Cloud with Workspace API Token"
  tags        = ["foo:bar"]

  url                = "https://bitbucket.org/my-workspace"
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
    auth_type = "EmailAndApiToken"
    email     = "user@example.com"        # or use email_ref to reference a Harness secret
    token_ref = "account.api_token_secret"
  }
}

# Credentials http with Bitbucket repo/project Access Token
resource "harness_platform_connector_bitbucket" "access_token" {
  identifier  = "identifier_access_token"
  name        = "name_access_token"
  description = "Bitbucket with Access Token"
  tags        = ["foo:bar"]

  url                = "https://bitbucket.org/my-workspace"
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
    auth_type = "AccessToken"
    token_ref = "account.access_token_secret"
  }
}

# Credentials ssh
resource "harness_platform_connector_bitbucket" "ssh" {
  identifier  = "identifier_ssh"
  name        = "name_ssh"
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
