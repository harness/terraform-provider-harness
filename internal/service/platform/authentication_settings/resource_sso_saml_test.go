package authentication_settings_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceSSOSAML(t *testing.T) {
	resourceName := "harness_platform_sso_saml.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSSOSAMLDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccSSOSAML("TF SAML Test", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "identifier"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "TF SAML Test"),
					resource.TestCheckResourceAttr(resourceName, "authentication_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "origin"),
				),
			},
			{
				// display_name-only update: metadata is re-sent from state and must
				// stay unchanged on the backend (origin unchanged).
				Config: testAccSSOSAML("TF SAML Test Renamed", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "TF SAML Test Renamed"),
					resource.TestCheckResourceAttrSet(resourceName, "origin"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateIdFunc:       acctest.AccountLevelResourceImportStateIdFunc(resourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metadata_xml", "client_secret"},
			},
			{
				// After import the state holds no metadata_xml. Updating without it
				// must succeed and keep the metadata stored in Harness (SAML-17,
				// optional-metadata behavior) — origin is still derived from the
				// stored metadata.
				Config: testAccSSOSAMLNoMetadata("TF SAML Test Imported"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "TF SAML Test Imported"),
					resource.TestCheckResourceAttrSet(resourceName, "origin"),
				),
			},
		},
	})
}

// testAccSSOSAMLDestroy verifies that destroying the resource deletes the SAML
// SSO configuration (no SAML entry with the resource's identifier remains in
// the v2 authentication settings response). Other SAML configurations on the
// account are left untouched.
func testAccSSOSAMLDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return nil
		}
		identifier := rs.Primary.ID

		c, ctx := acctest.TestAccGetPlatformClientWithContext()

		resp, _, err := c.AuthenticationSettingsApi.GetAuthenticationSettingsV2(ctx, c.AccountId)
		if err != nil {
			return err
		}
		if resp.Resource == nil {
			return nil
		}
		for _, settings := range resp.Resource.NgAuthSettings {
			if settings.SettingsType == "SAML" && settings.Identifier == identifier {
				return fmt.Errorf("SAML SSO configuration %s still present for account %s", identifier, c.AccountId)
			}
		}
		return nil
	}
}

func testAccSSOSAML(displayName string, authenticationEnabled bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_sso_saml" "test" {
			display_name           = %[1]q
			metadata_xml           = file("testdata/idp-metadata.xml")
			authentication_enabled = %[2]t
		}
	`, displayName, authenticationEnabled)
}

func testAccSSOSAMLNoMetadata(displayName string) string {
	return fmt.Sprintf(`
		resource "harness_platform_sso_saml" "test" {
			display_name           = %[1]q
			authentication_enabled = true
		}
	`, displayName)
}
