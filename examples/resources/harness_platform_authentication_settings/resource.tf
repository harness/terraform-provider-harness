# Manage account level authentication settings.
# This is a singleton resource: one instance per Harness account.
resource "harness_platform_authentication_settings" "example" {
  two_factor_enabled               = true
  whitelisted_domains              = ["harness.io", "example.com"]
  session_timeout_minutes          = 60
  absolute_session_timeout_minutes = 480
  public_access_enabled            = false
}
