# Manage the OAuth SSO configuration of a Harness account.
# This is a singleton resource: one instance per Harness account.
resource "harness_platform_sso_oauth" "example" {
  allowed_providers = ["GOOGLE", "GITHUB", "AZURE"]
}
