package authentication_settings_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceSSOLDAP(t *testing.T) {
	resourceName := "harness_platform_sso_ldap.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSSOLDAPDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccSSOLDAP("TF LDAP Test", false, "0 0 */1 ? * * *"),
				Check: resource.ComposeTestCheckFunc(
					// identifier is derived from the name: trimmed, spaces -> underscores.
					resource.TestCheckResourceAttr(resourceName, "identifier", "TF_LDAP_Test"),
					resource.TestCheckResourceAttr(resourceName, "name", "TF LDAP Test"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "cron_expression", "0 0 */1 ? * * *"),
					resource.TestCheckResourceAttr(resourceName, "connection_settings.0.host", "ldap.example.com"),
					resource.TestCheckResourceAttr(resourceName, "connection_settings.0.port", "636"),
					resource.TestCheckResourceAttr(resourceName, "connection_settings.0.ssl_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "connection_settings.0.secret_ref_path", "account.tf_acc_ldap_bind"),
					resource.TestCheckResourceAttr(resourceName, "user_settings.0.base_dn", "ou=users,dc=example,dc=com"),
					resource.TestCheckResourceAttr(resourceName, "user_settings.0.search_filter", "(objectClass=person)"),
					resource.TestCheckResourceAttr(resourceName, "group_settings.0.base_dn", "ou=groups,dc=example,dc=com"),
				),
			},
			{
				// Update name/disabled/cron; the identifier was fixed at create and must not change.
				Config: testAccSSOLDAP("TF LDAP Test Updated", true, "0 0 */2 ? * * *"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", "TF_LDAP_Test"),
					resource.TestCheckResourceAttr(resourceName, "name", "TF LDAP Test Updated"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "cron_expression", "0 0 */2 ? * * *"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
				// The GET API returns the full configuration (including secret_ref_path,
				// which is a reference, not secret material), so all attributes verify.
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDataSourceSSOLDAP(t *testing.T) {
	resourceName := "harness_platform_sso_ldap.test"
	dataSourceName := "data.harness_platform_sso_ldap.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSSOLDAPDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccSSOLDAPDataSource("TF LDAP DS Test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "identifier"),
					resource.TestCheckResourceAttr(dataSourceName, "name", "TF LDAP DS Test"),
					resource.TestCheckResourceAttr(dataSourceName, "connection_settings.0.host", "ldap.example.com"),
					resource.TestCheckResourceAttr(dataSourceName, "connection_settings.0.secret_ref_path", "account.tf_acc_ldap_bind"),
					resource.TestCheckResourceAttr(dataSourceName, "user_settings.0.base_dn", "ou=users,dc=example,dc=com"),
				),
			},
		},
	})
}

// testAccSSOLDAPDestroy verifies that destroying the resource deletes the account's
// LDAP configuration: the NextGen LDAP settings API must return 404 afterwards.
func testAccSSOLDAPDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		if _, ok := state.RootModule().Resources[resourceName]; !ok {
			return nil
		}

		session := acctest.TestAccGetApiClientFromProvider()
		c, ctx := session.GetLDAPClientWithContext(nil)

		_, httpResp, err := c.LdapSettingsApi.GetLdapSettings(ctx)
		if err != nil {
			if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
				return nil
			}
			return err
		}
		return fmt.Errorf("LDAP SSO configuration still present for account %s", session.AccountId)
	}
}

func testAccSSOLDAPSecret() string {
	return `
		resource "harness_platform_secret_text" "test" {
			identifier  = "tf_acc_ldap_bind"
			name        = "tf_acc_ldap_bind"
			description = "Bind password for the LDAP acceptance test"
			tags        = []

			secret_manager_identifier = "harnessSecretManager"
			value_type                = "Inline"
			value                     = "dummy-bind-password"
		}
	`
}

func testAccSSOLDAP(name string, disabled bool, cronExpression string) string {
	return fmt.Sprintf(`
		%[1]s

		resource "harness_platform_sso_ldap" "test" {
			name            = %[2]q
			disabled        = %[3]t
			cron_expression = %[4]q

			connection_settings {
				host            = "ldap.example.com"
				bind_dn         = "cn=admin,dc=example,dc=com"
				secret_ref_path = "account.${harness_platform_secret_text.test.identifier}"
			}

			user_settings {
				base_dn = "ou=users,dc=example,dc=com"
			}

			group_settings {
				base_dn = "ou=groups,dc=example,dc=com"
			}
		}
	`, testAccSSOLDAPSecret(), name, disabled, cronExpression)
}

func testAccSSOLDAPDataSource(name string) string {
	return fmt.Sprintf(`
		%[1]s

		resource "harness_platform_sso_ldap" "test" {
			name = %[2]q

			connection_settings {
				host            = "ldap.example.com"
				bind_dn         = "cn=admin,dc=example,dc=com"
				secret_ref_path = "account.${harness_platform_secret_text.test.identifier}"
			}

			user_settings {
				base_dn = "ou=users,dc=example,dc=com"
			}
		}

		data "harness_platform_sso_ldap" "test" {
			depends_on = [harness_platform_sso_ldap.test]
		}
	`, testAccSSOLDAPSecret(), name)
}
