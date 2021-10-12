package provider

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceK8sCloudProviderConnector_delegate(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		updatedName  = fmt.Sprintf("%s_updated", name)
		resourceName = "harness_cloudprovider_kubernetes.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceK8sCloudProvider_delegate(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckK8sCloudProviderExists(t, resourceName, name),
				),
			},
			{
				Config: testAccResourceK8sCloudProvider_delegate(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					testAccCheckK8sCloudProviderExists(t, resourceName, updatedName),
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

func TestAccResourceK8sCloudProviderConnector_usagescope_application(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "harness_cloudprovider_kubernetes.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceK8sCloudProvider_usagescope_application(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckK8sCloudProviderExists(t, resourceName, name),
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

func TestAccResourceK8sCloudProviderConnector_username_password(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "harness_cloudprovider_kubernetes.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceK8sCloudProvider_username_password(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckK8sCloudProviderExists(t, resourceName, name),
				),
			},
		},
	})
}

func TestAccResourceK8sCloudProviderConnector_oidc(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "harness_cloudprovider_kubernetes.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceK8sCloudProvider_oidc(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckK8sCloudProviderExists(t, resourceName, name),
				),
			},
		},
	})
}

func TestAccResourceK8sCloudProviderConnector_service_account(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "harness_cloudprovider_kubernetes.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCloudProviderDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceK8sCloudProvider_service_account(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckK8sCloudProviderExists(t, resourceName, name),
				),
			},
		},
	})
}

func TestAccResourceK8sCloudProviderConnector_DeleteUnderlyingResource(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "harness_cloudprovider_kubernetes.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceK8sCloudProvider_service_account(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckK8sCloudProviderExists(t, resourceName, name),
				),
			},
			{
				PreConfig: func() {
					testAccConfigureProvider()
					c := testAccProvider.Meta().(*api.Client)
					cp, err := c.CloudProviders().GetKubernetesCloudProviderByName(name)
					require.NoError(t, err)
					require.NotNil(t, cp)

					err = c.CloudProviders().DeleteCloudProvider(cp.Id)
					require.NoError(t, err)
				},
				Config:             testAccResourceGcpCloudProvider(name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// func TestAccResourceK8sCloudProviderConnector_custom(t *testing.T) {

// 	var (
// 		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
// 		resourceName = "harness_cloudprovider_kubernetes.test"
// 	)

// 	resource.UnitTest(t, resource.TestCase{
// 		PreCheck:          func() { testAccPreCheck(t) },
// 		ProviderFactories: providerFactories,
// 		CheckDestroy:      testAccCloudProviderDestroy(resourceName),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccResourceK8sCloudProvider_custom(name),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr(resourceName, "name", name),
// 					testAccCheckK8sCloudProviderExists(t, resourceName, name),
// 				),
// 			},
// 		},
// 	})
// }

func testAccResourceK8sCloudProvider_delegate(name string) string {
	return fmt.Sprintf(`
		resource "harness_cloudprovider_kubernetes" "test" {
			name = "%[1]s"
			skip_validation = true

			authentication {
				delegate_selectors = ["test"]
			}

			usage_scope {
				environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
			}
			
			usage_scope {
				environment_filter_type = "PRODUCTION_ENVIRONMENTS"
			}
		}
`, name)
}

func testAccResourceK8sCloudProvider_usagescope_application(name string) string {
	return fmt.Sprintf(`
		resource "harness_application" "test" {
			name = "%[1]s"
		}

		resource "harness_environment" "test" {
			name = "%[1]s"
			app_id = harness_application.test.id
			type = "NON_PROD"
		}

		resource "harness_cloudprovider_kubernetes" "test" {
			name = "%[1]s"
			skip_validation = true

			authentication {
				delegate_selectors = ["test"]
			}

			usage_scope {
				application_id = harness_application.test.id
				environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
			}

			usage_scope {
				environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
			}

			usage_scope {
				application_id = harness_application.test.id
				environment_id = harness_environment.test.id
			}

		}
`, name)
}

func testAccResourceK8sCloudProvider_username_password(name string) string {
	return fmt.Sprintf(`
		data "harness_secret_manager" "default" {
			default = true
		}

		resource "harness_encrypted_text" "username" {
			name = "%[1]s_username"
			value = "username"
			secret_manager_id = data.harness_secret_manager.default.id
		}

		resource "harness_encrypted_text" "password" {
			name = "%[1]s_password"
			value = "password"
			secret_manager_id = data.harness_secret_manager.default.id
		}

		resource "harness_cloudprovider_kubernetes" "test" {
			name = "%[1]s"
			skip_validation = true

			authentication {
				username_password {
					master_url = "https://localhost.com"
					username_secret_name = harness_encrypted_text.username.name
					password_secret_name = harness_encrypted_text.password.name
				}
			}
		}
`, name)
}

func testAccResourceK8sCloudProvider_oidc(name string) string {
	return fmt.Sprintf(`
		data "harness_secret_manager" "default" {
			default = true
		}

		resource "harness_encrypted_text" "username" {
			name = "%[1]s_username"
			value = "username"
			secret_manager_id = data.harness_secret_manager.default.id
		}

		resource "harness_encrypted_text" "password" {
			name = "%[1]s_password"
			value = "password"
			secret_manager_id = data.harness_secret_manager.default.id
		}

		resource "harness_encrypted_text" "client_id" {
			name = "%[1]s_client_id"
			value = "client_id"
			secret_manager_id = data.harness_secret_manager.default.id
		}

		resource "harness_encrypted_text" "client_secret" {
			name = "%[1]s_client_secret"
			value = "client_secret"
			secret_manager_id = data.harness_secret_manager.default.id
		}

		resource "harness_cloudprovider_kubernetes" "test" {
			name = "%[1]s"
			skip_validation = true

			authentication {
				oidc {
					identity_provider_url = "https://identity.com"
					username = harness_encrypted_text.username.name
					password_secret_name = harness_encrypted_text.password.name
					client_id_secret_name = harness_encrypted_text.client_id.name
					client_secret_secret_name = harness_encrypted_text.client_secret.name
					scopes = ["openid", "profile", "email"]
					master_url = "https://localhost.com"
				}
			}
		}
`, name)
}

func testAccResourceK8sCloudProvider_service_account(name string) string {
	return fmt.Sprintf(`
		data "harness_secret_manager" "default" {
			default = true
		}

		resource "harness_encrypted_text" "token" {
			name = "%[1]s_token"
			value = "token"
			secret_manager_id = data.harness_secret_manager.default.id
		}

		resource "harness_encrypted_text" "ca_cert" {
			name = "%[1]s_password"
			value = "ca_cert"
			secret_manager_id = data.harness_secret_manager.default.id
		}

		resource "harness_cloudprovider_kubernetes" "test" {
			name = "%[1]s"
			skip_validation = true

			authentication {
				service_account {
					master_url = "https://localhost.com"
					service_account_token_secret_name = harness_encrypted_text.token.name
					ca_certificate_secret_name = harness_encrypted_text.ca_cert.name
				}
			}
		}
`, name)
}

// func testAccResourceK8sCloudProvider_custom(name string) string {
// 	return fmt.Sprintf(`
// 		data "harness_secret_manager" "default" {
// 			default = true
// 		}

// 		resource "harness_encrypted_text" "username" {
// 			name = "%[1]s_username"
// 			value = "username"
// 			secret_manager_id = data.harness_secret_manager.default.id
// 		}

// 		resource "harness_encrypted_text" "password" {
// 			name = "%[1]s_password"
// 			value = "password"
// 			secret_manager_id = data.harness_secret_manager.default.id
// 		}

// 		resource "harness_encrypted_text" "ca_cert" {
// 			name = "%[1]s_ca_cert"
// 			value = "cacert"
// 			secret_manager_id = data.harness_secret_manager.default.id
// 		}

// 		resource "harness_encrypted_text" "client_cert" {
// 			name = "%[1]s_client_cert"
// 			value = "clientcert"
// 			secret_manager_id = data.harness_secret_manager.default.id
// 		}

// 		resource "harness_encrypted_text" "client_key" {
// 			name = "%[1]s_client_key"
// 			value = "clientcert"
// 			secret_manager_id = data.harness_secret_manager.default.id
// 		}

// 		resource "harness_encrypted_text" "client_key_passphrase" {
// 			name = "%[1]s_client_key_passphrase"
// 			value = "clientcert"
// 			secret_manager_id = data.harness_secret_manager.default.id
// 		}

// 		resource "harness_encrypted_text" "service_account_token" {
// 			name = "%[1]s_service_account_token"
// 			value = "clientcert"
// 			secret_manager_id = data.harness_secret_manager.default.id
// 		}

// 		resource "harness_cloudprovider_kubernetes" "test" {
// 			name = "%[1]s"
// 			skip_validation = true

// 			authentication {
// 				custom {
// 					master_url = "https://localhost.com"
// 					username_secret_name = harness_encrypted_text.username.name
// 					password_secret_name = harness_encrypted_text.password.name
// 					ca_certificate_secret_name = harness_encrypted_text.ca_cert.name
// 					client_certificate_secret_name = harness_encrypted_text.client_cert.name
// 					client_key_secret_name = harness_encrypted_text.client_key.name
// 					client_key_passphrase_secret_name = harness_encrypted_text.client_key_passphrase.name
// 					client_key_algorithm = "rsa-sa"
// 					service_account_token_secret_name = harness_encrypted_text.service_account_token.name
// 				}
// 			}
// 		}
// `, name)
// }

func testAccCheckK8sCloudProviderExists(t *testing.T, resourceName, cloudProviderName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		cp := &cac.KubernetesCloudProvider{}
		err := testAccGetCloudProvider(resourceName, state, cp)
		if err != nil {
			return err
		}
		return nil
	}
}
