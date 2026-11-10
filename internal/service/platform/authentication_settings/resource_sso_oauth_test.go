package authentication_settings_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceSSOOAuth(t *testing.T) {
	resourceName := "harness_platform_sso_oauth.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSSOOAuthDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccSSOOAuth(`["GOOGLE", "GITHUB"]`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "allowed_providers.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "allowed_providers.*", "GOOGLE"),
					resource.TestCheckTypeSetElemAttr(resourceName, "allowed_providers.*", "GITHUB"),
				),
			},
			{
				// Swap providers: set semantics — removal of GOOGLE/GITHUB and
				// addition of AZURE in one apply.
				Config: testAccSSOOAuth(`["AZURE"]`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "allowed_providers.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "allowed_providers.*", "AZURE"),
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

// TestAccResourceSSOOAuthDestroyWithMechanismOAuth mirrors the UI disable flow:
// the account's authentication mechanism is switched to OAUTH (as the UI toggle
// does) after the providers are configured, and the resource is then destroyed.
// The backend rejects removing the OAuth mechanism while it is the active
// mechanism, so destroy must first switch the mechanism back to USER_PASSWORD —
// exactly like the UI. The framework destroy after the last step exercises this.
func TestAccResourceSSOOAuthDestroyWithMechanismOAuth(t *testing.T) {
	resourceName := "harness_platform_sso_oauth.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSSOOAuthDestroyWithMechanismCheck(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccSSOOAuth(`["GOOGLE"]`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "allowed_providers.#", "1"),
					// Simulate the UI toggle: enable OAuth as the account's
					// authentication mechanism out of band.
					testAccSSOOAuthSetAuthMechanism("OAUTH"),
				),
			},
		},
	})
}

// TestAccResourceSSOOAuthEmptyProvidersValidation verifies that an empty
// provider list fails at plan time, matching the UI guard and the API's
// "At least one OAuth provider must be selected." rejection.
func TestAccResourceSSOOAuthEmptyProvidersValidation(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccSSOOAuth(`[]`),
				ExpectError: regexp.MustCompile(`allowed_providers`),
			},
		},
	})
}

// testAccSSOOAuthSetAuthMechanism switches the account's authentication
// mechanism via the platform API, simulating an out-of-band UI change.
func testAccSSOOAuthSetAuthMechanism(mechanism string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		c, ctx := acctest.TestAccGetPlatformClientWithContext()

		resp, _, err := c.AuthenticationSettingsApi.UpdateAuthMechanism(ctx, c.AccountId,
			&nextgen.AuthenticationSettingsApiUpdateAuthMechanismOpts{
				AuthenticationMechanism: optional.NewString(mechanism),
			})
		if err != nil {
			return err
		}
		if !resp.Resource {
			return fmt.Errorf("failed to set authentication mechanism to %s for account %s", mechanism, c.AccountId)
		}
		return nil
	}
}

// testAccSSOOAuthDestroyWithMechanismCheck verifies the full UI-parity destroy
// outcome: no OAUTH entry remains and the account authentication mechanism was
// switched back to USER_PASSWORD.
func testAccSSOOAuthDestroyWithMechanismCheck(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		if err := testAccSSOOAuthDestroy(resourceName)(state); err != nil {
			return err
		}

		c, ctx := acctest.TestAccGetPlatformClientWithContext()
		resp, _, err := c.AuthenticationSettingsApi.GetAuthenticationSettingsV2(ctx, c.AccountId)
		if err != nil {
			return err
		}
		if resp.Resource != nil && resp.Resource.AuthenticationMechanism != "USER_PASSWORD" {
			return fmt.Errorf("expected authentication mechanism USER_PASSWORD after OAuth destroy, got %s for account %s",
				resp.Resource.AuthenticationMechanism, c.AccountId)
		}
		return nil
	}
}

// testAccSSOOAuthDestroy verifies that destroying the resource removes the OAuth
// mechanism from the account (no OAUTH entry remains in the v2 authentication
// settings response).
func testAccSSOOAuthDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		c, ctx := acctest.TestAccGetPlatformClientWithContext()

		resp, _, err := c.AuthenticationSettingsApi.GetAuthenticationSettingsV2(ctx, c.AccountId)
		if err != nil {
			return err
		}
		if resp.Resource == nil {
			return nil
		}
		for _, settings := range resp.Resource.NgAuthSettings {
			if settings.SettingsType == "OAUTH" {
				return fmt.Errorf("OAuth SSO configuration still present for account %s: %+v", c.AccountId, settings)
			}
		}
		return nil
	}
}

func testAccSSOOAuth(providers string) string {
	return fmt.Sprintf(`
		resource "harness_platform_sso_oauth" "test" {
			allowed_providers = %[1]s
		}
	`, providers)
}
