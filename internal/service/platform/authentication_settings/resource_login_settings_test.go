package authentication_settings_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceLoginSettings(t *testing.T) {
	resourceName := "harness_platform_login_settings.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccLoginSettingsDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccLoginSettingsFull(true, 12, 3, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "password_strength_policy.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "password_strength_policy.0.min_number_of_characters", "12"),
					resource.TestCheckResourceAttr(resourceName, "password_strength_policy.0.min_number_of_uppercase_characters", "1"),
					resource.TestCheckResourceAttr(resourceName, "user_lockout_policy.0.enable_lockout_policy", "true"),
					resource.TestCheckResourceAttr(resourceName, "user_lockout_policy.0.number_of_failed_attempts_before_lockout", "3"),
					resource.TestCheckResourceAttr(resourceName, "user_lockout_policy.0.lockout_period", "24"),
					resource.TestCheckResourceAttr(resourceName, "user_lockout_policy.0.notify_user", "true"),
					resource.TestCheckResourceAttr(resourceName, "password_expiration_policy.0.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "password_expiration_policy.0.days_before_password_expires", "90"),
				),
			},
			{
				// Manage only the password strength policy; the lockout and expiration
				// policies must keep the values applied in the previous step (the
				// resource does read-modify-write, same behavior as the UI).
				Config: testAccLoginSettingsStrengthOnly(false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "password_strength_policy.0.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "user_lockout_policy.0.enable_lockout_policy", "true"),
					resource.TestCheckResourceAttr(resourceName, "user_lockout_policy.0.number_of_failed_attempts_before_lockout", "3"),
					resource.TestCheckResourceAttr(resourceName, "password_expiration_policy.0.enabled", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// testAccLoginSettingsDestroy verifies that destroying the resource resets the
// account login settings to the Harness defaults: password strength and
// expiration policies disabled, lockout policy disabled with the platform
// default thresholds (5 attempts / 24 hours / notify user).
func testAccLoginSettingsDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		c, ctx := acctest.TestAccGetPlatformClientWithContext()

		resp, _, err := c.AuthenticationSettingsApi.GetLoginSettings(ctx, c.AccountId)
		if err != nil {
			return err
		}
		if resp.Resource == nil {
			return nil
		}

		settings := resp.Resource
		if settings.PasswordStrengthPolicy != nil && settings.PasswordStrengthPolicy.Enabled {
			return fmt.Errorf("password strength policy is still enabled for account %s", c.AccountId)
		}
		if settings.PasswordExpirationPolicy != nil && settings.PasswordExpirationPolicy.Enabled {
			return fmt.Errorf("password expiration policy is still enabled for account %s", c.AccountId)
		}
		if lockout := settings.UserLockoutPolicy; lockout != nil {
			if lockout.EnableLockoutPolicy {
				return fmt.Errorf("user lockout policy is still enabled for account %s", c.AccountId)
			}
			if lockout.NumberOfFailedAttemptsBeforeLockout != 5 || lockout.LockOutPeriod != 24 || !lockout.NotifyUser {
				return fmt.Errorf("user lockout policy was not reset to defaults for account %s: %+v", c.AccountId, lockout)
			}
		}
		return nil
	}
}

func testAccLoginSettingsFull(strengthEnabled bool, minCharacters, lockoutAttempts int, expirationEnabled bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_login_settings" "test" {
			password_strength_policy {
				enabled                            = %[1]t
				min_number_of_characters           = %[2]d
				min_number_of_uppercase_characters = 1
				min_number_of_lowercase_characters = 1
				min_number_of_special_characters   = 1
				min_number_of_digits               = 1
			}

			user_lockout_policy {
				enable_lockout_policy                    = true
				number_of_failed_attempts_before_lockout = %[3]d
				lockout_period                           = 24
				notify_user                              = true
			}

			password_expiration_policy {
				enabled                                       = %[4]t
				days_before_password_expires                  = 90
				days_before_user_notified_of_password_expiration = 5
			}
		}
	`, strengthEnabled, minCharacters, lockoutAttempts, expirationEnabled)
}

func testAccLoginSettingsStrengthOnly(enabled bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_login_settings" "test" {
			password_strength_policy {
				enabled                            = %[1]t
				min_number_of_characters           = 8
				min_number_of_uppercase_characters = 1
				min_number_of_lowercase_characters = 1
				min_number_of_special_characters   = 0
				min_number_of_digits               = 1
			}
		}
	`, enabled)
}
