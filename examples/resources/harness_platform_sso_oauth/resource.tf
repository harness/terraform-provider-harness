# Manage the OAuth SSO configuration of a Harness account.
# This is a singleton resource: one instance per Harness account.
# Note: this only configures the allowed OAuth providers; enabling OAuth as the
# account's login mechanism is managed separately via account authentication settings.
resource "harness_platform_sso_oauth" "example" {
  allowed_providers = ["GOOGLE", "GITHUB", "AZURE"]
}
