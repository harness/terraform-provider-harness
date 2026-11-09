package authentication_settings_test

import (
	"fmt"
	"testing"

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
			},
		},
	})
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
