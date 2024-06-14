package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorServiceNow_UsernamePassword(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_service_now.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorServiceNow_UsernamePassword(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_now_url", "https://test.atlassian.com"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.username_password.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnectorServiceNow_UsernamePassword(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_now_url", "https://test.atlassian.com"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.username_password.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceConnectorServiceNow_Adfs(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_service_now.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorServiceNow_Adfs(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_now_url", "https://test.atlassian.com"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.adfs.0.adfs_url", "https://adfs_url.com"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnectorServiceNow_Adfs(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_now_url", "https://test.atlassian.com"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.adfs.0.adfs_url", "https://adfs_url.com"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceConnectorServiceNow_RefreshToken(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_service_now.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorServiceNow_RefreshToken(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_now_url", "https://test.service-now.com"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.refresh_token.0.token_url", "https://test.service-now.com/oauth_token.do"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.refresh_token.0.scope", "email openid profile"),					
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnectorServiceNow_RefreshToken(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_now_url", "https://test.service-now.com"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.refresh_token.0.token_url", "https://test.service-now.com/oauth_token.do"),
					resource.TestCheckResourceAttr(resourceName, "auth.0.refresh_token.0.scope", "email openid profile"),					
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceConnectorServiceNow_UsernamePassword(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

		resource "harness_platform_connector_service_now" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			service_now_url = "https://test.atlassian.com"
			delegate_selectors = ["harness-delegate"]
			auth {
				auth_type = "UsernamePassword"
				username_password {
					username = "admin"
					password_ref = "account.${harness_platform_secret_text.test.id}"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}
`, id, name)
}

func testAccResourceConnectorServiceNow_Adfs(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

		resource "harness_platform_connector_service_now" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			service_now_url = "https://test.atlassian.com"
			delegate_selectors = ["harness-delegate"]
			auth {
				auth_type = "AdfsClientCredentialsWithCertificate"
				adfs {
					certificate_ref = "account.${harness_platform_secret_text.test.id}"
					private_key_ref = "account.${harness_platform_secret_text.test.id}"
					client_id_ref = "account.${harness_platform_secret_text.test.id}"
					resource_id_ref = "account.${harness_platform_secret_text.test.id}"
					adfs_url = "https://adfs_url.com"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}
`, id, name)
}

func testAccResourceConnectorServiceNow_RefreshToken(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

		resource "harness_platform_connector_service_now" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			service_now_url = "https://test.service-now.com"
			delegate_selectors = ["harness-delegate"]
			auth {
				auth_type = "RefreshTokenGrantType"
				refresh_token {
					token_url = "https://test.service-now.com/oauth_token.do"
					refresh_token_ref = "account.${harness_platform_secret_text.test.id}"
					client_id_ref = "account.${harness_platform_secret_text.test.id}"
					client_secret_ref = "account.${harness_platform_secret_text.test.id}"
					scope = "email openid profile"
				}
			}
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}
`, id, name)
}
