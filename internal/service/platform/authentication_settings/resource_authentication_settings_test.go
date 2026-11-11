package authentication_settings_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceAuthenticationSettings(t *testing.T) {
	resourceName := "harness_platform_authentication_settings.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccAuthenticationSettingsDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// Create with only session_timeout_minutes declared must not wipe an
				// absolute session timeout set out-of-band (seeded by PreConfig): the
				// provider reads the live value and carries it into the full-replace
				// session-timeout API, the same read-then-write behavior as the UI.
				PreConfig: func() {
					testAccSetAccountSessionTimeout(t, 60, 600)
				},
				Config: testAccAuthenticationSettingsSessionOnly(true, 45),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "two_factor_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "session_timeout_minutes", "45"),
					resource.TestCheckResourceAttr(resourceName, "absolute_session_timeout_minutes", "600"),
				),
			},
			{
				Config: testAccAuthenticationSettings(true, `["harness.io"]`, 60, 480, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "two_factor_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "whitelisted_domains.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "whitelisted_domains.*", "harness.io"),
					resource.TestCheckResourceAttr(resourceName, "session_timeout_minutes", "60"),
					resource.TestCheckResourceAttr(resourceName, "absolute_session_timeout_minutes", "480"),
					resource.TestCheckResourceAttr(resourceName, "public_access_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "authentication_mechanism"),
				),
			},
			{
				Config: testAccAuthenticationSettings(false, `["harness.io", "example.com"]`, 30, 120, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "two_factor_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "whitelisted_domains.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "whitelisted_domains.*", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "session_timeout_minutes", "30"),
					resource.TestCheckResourceAttr(resourceName, "absolute_session_timeout_minutes", "120"),
				),
			},
			{
				// Config that declares only two_factor_enabled must leave the values
				// managed in the previous steps untouched: undeclared Optional+Computed
				// fields carry the last-read API values and are sent back on update
				// (the same read-then-write behavior as the Harness UI).
				Config: testAccAuthenticationSettings2FAOnly(false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "two_factor_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "whitelisted_domains.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "session_timeout_minutes", "30"),
					resource.TestCheckResourceAttr(resourceName, "absolute_session_timeout_minutes", "120"),
					resource.TestCheckResourceAttr(resourceName, "public_access_enabled", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
		},
	})
}

// testAccAuthenticationSettingsDestroy verifies that destroying the resource resets the
// account authentication settings to the Harness defaults (two-factor authentication
// disabled and whitelisted domains cleared).
func testAccAuthenticationSettingsDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		c, ctx := acctest.TestAccGetPlatformClientWithContext()

		resp, _, err := c.AuthenticationSettingsApi.GetAuthenticationSettingsV2(ctx, c.AccountId)
		if err != nil {
			return err
		}
		if resp.Resource == nil {
			return nil
		}
		if resp.Resource.TwoFactorEnabled {
			return fmt.Errorf("two-factor authentication is still enabled for account %s", c.AccountId)
		}
		if len(resp.Resource.WhitelistedDomains) > 0 {
			return fmt.Errorf("whitelisted domains are still set for account %s: %v", c.AccountId, resp.Resource.WhitelistedDomains)
		}
		if resp.Resource.PublicAccessEnabled {
			return fmt.Errorf("public access is still enabled for account %s", c.AccountId)
		}
		return nil
	}
}

// testAccSetAccountSessionTimeout seeds the account session timeout out-of-band,
// simulating a value an admin set through the Harness UI.
func testAccSetAccountSessionTimeout(t *testing.T, sessionTimeout, absoluteSessionTimeout int32) {
	c, ctx := acctest.TestAccGetPlatformClientWithContext()

	_, _, err := c.AuthenticationSettingsApi.SetSessionTimeoutAtAccountLevel(ctx, nextgen.SessionTimeoutSettings{
		SessionTimeOutInMinutes:         sessionTimeout,
		AbsoluteSessionTimeOutInMinutes: absoluteSessionTimeout,
	}, c.AccountId)
	if err != nil {
		t.Fatalf("failed to seed account session timeout: %v", err)
	}
}

func testAccAuthenticationSettingsSessionOnly(twoFactorEnabled bool, sessionTimeout int) string {
	return fmt.Sprintf(`
		resource "harness_platform_authentication_settings" "test" {
			two_factor_enabled      = %[1]t
			session_timeout_minutes = %[2]d
		}
	`, twoFactorEnabled, sessionTimeout)
}

func testAccAuthenticationSettings2FAOnly(twoFactorEnabled bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_authentication_settings" "test" {
			two_factor_enabled = %[1]t
		}
	`, twoFactorEnabled)
}

func testAccAuthenticationSettings(twoFactorEnabled bool, whitelistedDomains string, sessionTimeout, absoluteSessionTimeout int, publicAccess bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_authentication_settings" "test" {
			two_factor_enabled                = %[1]t
			whitelisted_domains               = %[2]s
			session_timeout_minutes           = %[3]d
			absolute_session_timeout_minutes  = %[4]d
			public_access_enabled             = %[5]t
		}
	`, twoFactorEnabled, whitelistedDomains, sessionTimeout, absoluteSessionTimeout, publicAccess)
}
