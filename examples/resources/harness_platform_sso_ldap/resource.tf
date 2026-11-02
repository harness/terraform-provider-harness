# Manage the LDAP SSO configuration of a Harness account.

variable "ldap_bind_password" {
  description = "Password of the LDAP bind user, used when password_type is INLINE."
  type        = string
  sensitive   = true
}

variable "ldap_password_type" {
  description = "How the bind password is provided: INLINE or ENCRYPTED_SECRET."
  type        = string
  default     = "INLINE"
}

resource "harness_platform_sso_ldap" "example" {
  identifier      = "example_ldap"
  display_name    = "Example LDAP"
  cron_expression = "0 0/15 * * * ?"
  disabled        = false

  connection_settings {
    host            = "ldap.example.com"
    port            = 636
    ssl_enabled     = true
    bind_dn         = "cn=admin,dc=example,dc=com"
    bind_password   = var.ldap_bind_password
    password_type   = var.ldap_password_type
    connect_timeout = 5000
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

# The bind password can also reference a Harness secret instead of an inline value.
# For this mode, set the ldap_password_type variable to ENCRYPTED_SECRET.
resource "harness_platform_sso_ldap" "example_with_secret" {
  identifier   = "example_ldap_secret"
  display_name = "Example LDAP with Secret"

  connection_settings {
    host          = "ldap.example.com"
    port          = 636
    ssl_enabled   = true
    bind_dn       = "cn=admin,dc=example,dc=com"
    bind_password = harness_platform_secret_text.ldap_bind.id
    password_type = var.ldap_password_type
  }

  user_settings {
    base_dn = "ou=users,dc=example,dc=com"
  }

  group_settings {
    base_dn = "ou=groups,dc=example,dc=com"
  }
}
