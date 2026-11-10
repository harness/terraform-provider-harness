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

// TestAccResourceSSOSAMLFields exercises the optional-field flows of the
// resource against a live account:
//   - entity_identifier set and then cleared (clearing reverts to the Harness
//     default, mirroring the API and the UI);
//   - Azure client_id/client_secret set and rotated, and the post-import update
//     path that sends the client secret mask to preserve the stored secret;
//   - group_membership_attr driving the derived authorization_enabled flag;
//   - Just-in-time provisioning fields set and cleared;
//   - authentication_enabled toggled off through the dedicated endpoint.
func TestAccResourceSSOSAMLFields(t *testing.T) {
	resourceName := "harness_platform_sso_saml.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSSOSAMLDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				// All optional fields set (client credentials pair with the AZURE
				// provider type, as in the UI).
				Config: testAccSSOSAMLAllFields(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "TF SAML Fields"),
					resource.TestCheckResourceAttr(resourceName, "entity_identifier", "tf-saml-entity-identifier"),
					resource.TestCheckResourceAttr(resourceName, "logout_url", "https://idp-tf-acc-test.example.com/logout"),
					resource.TestCheckResourceAttr(resourceName, "group_membership_attr", "groups"),
					resource.TestCheckResourceAttr(resourceName, "authorization_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "authentication_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "saml_provider_type", "AZURE"),
					resource.TestCheckResourceAttr(resourceName, "client_id", "tf-azure-client-id"),
					resource.TestCheckResourceAttr(resourceName, "client_secret", "tf-azure-client-secret"),
					resource.TestCheckResourceAttr(resourceName, "jit_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "jit_validation_key", "department"),
					resource.TestCheckResourceAttr(resourceName, "jit_validation_value", "engineering"),
				),
			},
			{
				// Secret rotation and JIT value change: the update re-sends the real
				// client secret from the configuration.
				Config: testAccSSOSAMLAllFieldsRotated(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "client_id", "tf-azure-client-id-2"),
					resource.TestCheckResourceAttr(resourceName, "client_secret", "tf-azure-client-secret-2"),
					resource.TestCheckResourceAttr(resourceName, "jit_validation_key", "team"),
					resource.TestCheckResourceAttr(resourceName, "jit_validation_value", "platform"),
				),
			},
			{
				// Removing the optional fields clears them on the backend, mirroring
				// the API and the UI: entity_identifier reverts to the Harness default
				// (empty), authorization_enabled is derived off without a group
				// attribute, and the JIT matching fields are cleared.
				Config: testAccSSOSAMLOptionalsCleared(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "entity_identifier", ""),
					resource.TestCheckResourceAttr(resourceName, "logout_url", ""),
					resource.TestCheckResourceAttr(resourceName, "group_membership_attr", ""),
					resource.TestCheckResourceAttr(resourceName, "authorization_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "jit_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "jit_validation_key", ""),
					resource.TestCheckResourceAttr(resourceName, "jit_validation_value", ""),
					// Client credentials are untouched by clearing the other fields.
					resource.TestCheckResourceAttr(resourceName, "client_id", "tf-azure-client-id-2"),
					resource.TestCheckResourceAttr(resourceName, "client_secret", "tf-azure-client-secret-2"),
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
				// Post-import update with neither metadata_xml nor client_secret in the
				// configuration: the stored metadata is kept, and the update sends the
				// client secret mask so the stored Azure secret is preserved (the same
				// request the UI sends when editing without re-entering the secret).
				Config: testAccSSOSAMLFieldsNoMetadata("TF SAML Fields Imported", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", "TF SAML Fields Imported"),
					resource.TestCheckResourceAttrSet(resourceName, "origin"),
					resource.TestCheckResourceAttr(resourceName, "client_id", "tf-azure-client-id-2"),
				),
			},
			{
				// Toggling login off goes through the dedicated enable/disable endpoint
				// rather than the settings update.
				Config: testAccSSOSAMLFieldsNoMetadata("TF SAML Fields Imported", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "authentication_enabled", "false"),
				),
			},
		},
	})
}

func testAccSSOSAMLAllFields() string {
	return `
		resource "harness_platform_sso_saml" "test" {
			display_name           = "TF SAML Fields"
			metadata_xml           = file("testdata/idp-metadata.xml")
			entity_identifier      = "tf-saml-entity-identifier"
			logout_url             = "https://idp-tf-acc-test.example.com/logout"
			group_membership_attr  = "groups"
			saml_provider_type     = "AZURE"
			client_id              = "tf-azure-client-id"
			client_secret          = "tf-azure-client-secret"
			jit_enabled            = true
			jit_validation_key     = "department"
			jit_validation_value   = "engineering"
			authentication_enabled = true
		}
	`
}

func testAccSSOSAMLAllFieldsRotated() string {
	return `
		resource "harness_platform_sso_saml" "test" {
			display_name           = "TF SAML Fields"
			metadata_xml           = file("testdata/idp-metadata.xml")
			entity_identifier      = "tf-saml-entity-identifier"
			logout_url             = "https://idp-tf-acc-test.example.com/logout"
			group_membership_attr  = "groups"
			saml_provider_type     = "AZURE"
			client_id              = "tf-azure-client-id-2"
			client_secret          = "tf-azure-client-secret-2"
			jit_enabled            = true
			jit_validation_key     = "team"
			jit_validation_value   = "platform"
			authentication_enabled = true
		}
	`
}

func testAccSSOSAMLOptionalsCleared() string {
	return `
		resource "harness_platform_sso_saml" "test" {
			display_name           = "TF SAML Fields"
			metadata_xml           = file("testdata/idp-metadata.xml")
			saml_provider_type     = "AZURE"
			client_id              = "tf-azure-client-id-2"
			client_secret          = "tf-azure-client-secret-2"
			authentication_enabled = true
		}
	`
}

func testAccSSOSAMLFieldsNoMetadata(displayName string, authenticationEnabled bool) string {
	return fmt.Sprintf(`
		resource "harness_platform_sso_saml" "test" {
			display_name           = %[1]q
			authentication_enabled = %[2]t
		}
	`, displayName, authenticationEnabled)
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
