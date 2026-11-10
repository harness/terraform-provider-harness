# Manage a SAML SSO configuration of a Harness account.
# Multiple SAML configurations are supported per account.
resource "harness_platform_sso_saml" "example" {
  display_name   = "Example SAML"
  metadata_xml   = file("${path.module}/idp-metadata.xml")
  logout_url     = "https://idp.example.com/logout"
  jit_enabled    = false
}

# SAML configuration with group-based authorization and Azure client credentials.
# client_id/client_secret are only used with Azure AD: Harness calls Microsoft Graph
# to resolve the user's groups at login time. client_secret is the raw secret value
# of the Azure AD app registration (not a reference to a Harness secret); Harness
# stores it encrypted in the account's default secret manager.
variable "azure_saml_client_secret" {
  type      = string
  sensitive = true
}

resource "harness_platform_sso_saml" "example_with_groups" {
  display_name          = "Example SAML with Groups"
  metadata_xml          = file("${path.module}/idp-metadata.xml")
  group_membership_attr = "memberOf"
  saml_provider_type    = "AZURE"
  client_id             = "00000000-0000-0000-0000-000000000000"
  client_secret         = var.azure_saml_client_secret
  friendly_saml_name    = "Harness"
  jit_enabled           = true
  jit_validation_key    = "email"
  jit_validation_value  = ".*@example\\.com"
}
