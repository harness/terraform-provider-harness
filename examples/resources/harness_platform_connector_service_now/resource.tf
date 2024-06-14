# Credential type UsernamePassword
resource "harness_platform_connector_service_now" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  service_now_url    = "https://servicenow.com"
  delegate_selectors = ["harness-delegate"]
  auth {
    auth_type = "UsernamePassword"
    username_password {
      username     = "admin"
      password_ref = "account.password_ref"
    }
  }
}

# Credential type AdfsClientCredentialsWithCertificate
resource "harness_platform_connector_service_now" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  service_now_url    = "https://servicenow.com"
  delegate_selectors = ["harness-delegate"]
  auth {
    auth_type = "AdfsClientCredentialsWithCertificate"
    adfs {
      certificate_ref = "account.certificate_ref"
      private_key_ref = "account.private_key_ref}"
      client_id_ref   = "account.client_id_ref"
      resource_id_ref = "account.resource_id_ref"
      adfs_url        = "https://adfs_url.com"
    }
  }
}

# Credential type RefreshTokenGrantType
resource "harness_platform_connector_service_now" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  service_now_url    = "https://test.service-now.com"
  delegate_selectors = ["harness-delegate"]
  auth {
    auth_type = "RefreshTokenGrantType"
    adfs {
      token_url         = "https://test.service-now.com/oauth_token.do"
      refresh_token_ref = "account.refresh_token_ref"
      client_id_ref     = "account.client_id_ref"
      client_secret_ref = "account.client_secret_ref"
      scope             = "email openid profile"
    }
  }
}
