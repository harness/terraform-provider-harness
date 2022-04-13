package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorK8s_InheritFromDelegate(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_kubernetes.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorK8s_InheritFromDelegate(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.#", "1"),
				),
			},
			{
				Config: testAccResourceConnectorK8s_InheritFromDelegate(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "inherit_from_delegate.0.delegate_selectors.#", "1"),
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

func TestAccResourceConnectorK8s_ClientKeyCert(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_kubernetes.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorK8s_ClientKeyCert(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.master_url", "https://kubernetes.example.com"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.ca_cert_ref", "account.TEST_k8ss_client_stuff"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.client_cert_ref", "account.test_k8s_client_cert"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.client_key_ref", "account.TEST_k8s_client_key"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.client_key_passphrase_ref", "account.TEST_k8s_client_test"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.client_key_algorithm", "RSA"),
				),
			},
			{
				Config: testAccResourceConnectorK8s_ClientKeyCert(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.master_url", "https://kubernetes.example.com"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.ca_cert_ref", "account.TEST_k8ss_client_stuff"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.client_cert_ref", "account.test_k8s_client_cert"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.client_key_ref", "account.TEST_k8s_client_key"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.client_key_passphrase_ref", "account.TEST_k8s_client_test"),
					resource.TestCheckResourceAttr(resourceName, "client_key_cert.0.client_key_algorithm", "RSA"),
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

func TestAccResourceConnectorK8s_UsernamePassword(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_kubernetes.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorK8s_UsernamePassword(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "username_password.0.master_url", "https://kubernetes.example.com"),
					resource.TestCheckResourceAttr(resourceName, "username_password.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "username_password.0.password_ref", "account.TEST_k8s_client_test"),
				),
			},
			{
				Config: testAccResourceConnectorK8s_UsernamePassword(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "username_password.0.master_url", "https://kubernetes.example.com"),
					resource.TestCheckResourceAttr(resourceName, "username_password.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "username_password.0.password_ref", "account.TEST_k8s_client_test"),
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

func TestAccResourceConnectorK8s_ServiceAccount(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_kubernetes.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorK8s_ServiceAccount(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_account.0.master_url", "https://kubernetes.example.com"),
					resource.TestCheckResourceAttr(resourceName, "service_account.0.service_account_token_ref", "account.TEST_k8s_client_test"),
				),
			},
			{
				Config: testAccResourceConnectorK8s_ServiceAccount(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "service_account.0.master_url", "https://kubernetes.example.com"),
					resource.TestCheckResourceAttr(resourceName, "service_account.0.service_account_token_ref", "account.TEST_k8s_client_test"),
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

func TestAccResourceConnectorK8s_OpenIDConnect(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_kubernetes.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorK8s_OpenIDConnect(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.master_url", "https://kubernetes.example.com"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.issuer_url", "https://oidc.example.com"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.username_ref", "account.TEST_k8s_client_test"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.client_id_ref", "account.TEST_k8s_client_test"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.password_ref", "account.TEST_k8s_client_test"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.secret_ref", "account.TEST_k8s_client_test"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.scopes.#", "2"),
				),
			},
			{
				Config: testAccResourceConnectorK8s_OpenIDConnect(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.master_url", "https://kubernetes.example.com"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.issuer_url", "https://oidc.example.com"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.username_ref", "account.TEST_k8s_client_test"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.client_id_ref", "account.TEST_k8s_client_test"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.password_ref", "account.TEST_k8s_client_test"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.secret_ref", "account.TEST_k8s_client_test"),
					resource.TestCheckResourceAttr(resourceName, "openid_connect.0.scopes.#", "2"),
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

func testAccResourceConnectorK8s_ClientKeyCert(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_kubernetes" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			client_key_cert {
				master_url = "https://kubernetes.example.com"
				ca_cert_ref = "account.TEST_k8ss_client_stuff"
				client_cert_ref = "account.test_k8s_client_cert"
				client_key_ref = "account.TEST_k8s_client_key"
				client_key_passphrase_ref = "account.TEST_k8s_client_test"
				client_key_algorithm = "RSA"
			}
		}
`, id, name)
}

func testAccResourceConnectorK8s_UsernamePassword(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_kubernetes" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			username_password {
				master_url = "https://kubernetes.example.com"
				username = "admin"
				password_ref = "account.TEST_k8s_client_test"
			}
		}
`, id, name)
}

func testAccResourceConnectorK8s_ServiceAccount(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_kubernetes" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			service_account {
				master_url = "https://kubernetes.example.com"
				service_account_token_ref = "account.TEST_k8s_client_test"
			}
		}
`, id, name)
}

func testAccResourceConnectorK8s_OpenIDConnect(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_kubernetes" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			openid_connect {
				master_url = "https://kubernetes.example.com"
				issuer_url = "https://oidc.example.com"
				username_ref = "account.TEST_k8s_client_test"
				client_id_ref = "account.TEST_k8s_client_test"
				password_ref = "account.TEST_k8s_client_test"
				secret_ref = "account.TEST_k8s_client_test"
				scopes = [
					"scope1",
					"scope2"
				]
			}
		}
`, id, name)
}

func testAccResourceConnectorK8s_InheritFromDelegate(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_kubernetes" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			inherit_from_delegate {
				delegate_selectors = ["harness-delegate"]
			}
		}
`, id, name)
}
