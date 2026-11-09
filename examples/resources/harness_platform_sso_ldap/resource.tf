# Manage the account-level LDAP SSO configuration of a Harness account via the
# NextGen LDAP APIs. Requires the PL_ENABLE_NG_LDAP_SETTINGS feature flag to be
# enabled on the account. Exactly one LDAP configuration exists per account.

# The bind password is referenced via an account-level Harness secret.
resource "harness_platform_secret_text" "ldap_bind_password" {
  identifier  = "ldap_bind_password"
  name        = "ldap_bind_password"
  description = "Password of the LDAP bind user"
  tags        = []

  secret_manager_identifier = "harnessSecretManager"
  value_type                = "Inline"
  value                     = var.ldap_bind_password
}

variable "ldap_bind_password" {
  description = "Password of the LDAP bind user."
  type        = string
  sensitive   = true
}

resource "harness_platform_sso_ldap" "example" {
  name            = "Example LDAP"
  cron_expression = "0 0 */1 ? * * *"
  disabled        = false

  connection_settings {
    host            = "ldap.example.com"
    port            = 636
    ssl_enabled     = true
    bind_dn         = "cn=admin,dc=example,dc=com"
    secret_ref_path = "account.${harness_platform_secret_text.ldap_bind_password.identifier}"
  }

  user_settings {
    base_dn       = "ou=users,dc=example,dc=com"
    search_filter = "(objectClass=person)"
  }

  group_settings {
    base_dn       = "ou=groups,dc=example,dc=com"
    search_filter = "(objectClass=group)"
  }
}
