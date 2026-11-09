# Manage the login settings of a Harness account.
# This is a singleton resource: one instance per Harness account.
resource "harness_platform_login_settings" "example" {
  password_strength_policy {
    enabled                            = true
    min_number_of_characters           = 12
    min_number_of_uppercase_characters = 1
    min_number_of_lowercase_characters = 1
    min_number_of_special_characters   = 1
    min_number_of_digits               = 1
  }

  user_lockout_policy {
    enable_lockout_policy                    = true
    number_of_failed_attempts_before_lockout = 5
    lockout_period                           = 24
    notify_user                              = true
  }

  password_expiration_policy {
    enabled                                        = true
    days_before_password_expires                   = 90
    days_before_user_notified_of_password_expiration = 5
  }
}
