package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorAzure_InheritFromDelegate_SA(t *testing.T) {

	id := fmt.Sprintf("ConnectorAzure_InheritFromDelegate_SA"+"_%s", utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_azure_cloud_provider.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAzure_inheritFromDelegate_sa(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceConnectorAzure_inheritFromDelegate_sa(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
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

func TestAccResourceConnectorAzure_InheritFromDelegate_UA(t *testing.T) {

	id := fmt.Sprintf("ConnectorAzure_InheritFromDelegate_UA"+"_%s", utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_azure_cloud_provider.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAzure_inheritFromDelegate_ua(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceConnectorAzure_inheritFromDelegate_ua(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
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

func TestAccResourceConnectorAzure_ManualDetails_Certificate(t *testing.T) {

	id := fmt.Sprintf("ConnectorAzure_ManualDetails_Certificate"+"_%s", utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_azure_cloud_provider.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAzure_manualDetails_certificate(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceConnectorAzure_manualDetails_certificate(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
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

func TestAccResourceConnectorAzure_ManualDetails_Secret(t *testing.T) {

	id := fmt.Sprintf("ConnectorAzure_ManualDetails_Secret"+"_%s", utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_azure_cloud_provider.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAzure_manualDetails_secret(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceConnectorAzure_manualDetails_secret(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
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

func testAccResourceConnectorAzure_manualDetails_secret(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "azureSecretManager"
		value_type = "Reference"
		value = "secret"
	}

	resource "harness_platform_connector_azure_cloud_provider" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		credentials {
			type = "ManualConfig"
			azure_manual_details {
				application_id = "application_id"
				tenant_id = "tenant_id"
				auth {
					type = "Secret"
					azure_client_secret_key {
						secret_ref = "account.${harness_platform_secret_text.test.id}"
					}
				}
			}
		}

		azure_environment_type = "AZURE"
		delegate_selectors = ["harness-delegate"]
	}
`, id, name)
}

func testAccResourceConnectorAzure_manualDetails_certificate(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "azureSecretManager"
		value_type = "Reference"
		value = "secret"
	}

	resource "harness_platform_connector_azure_cloud_provider" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		credentials {
			type = "ManualConfig"
			azure_manual_details {
				application_id = "application_id"
				tenant_id = "tenant_id"
				auth {
					type = "Certificate"
					azure_client_key_cert {
						certificate_ref = "account.${harness_platform_secret_text.test.id}"
					}
				}
			}
		}

		azure_environment_type = "AZURE"
		delegate_selectors = ["harness-delegate"]
	}
`, id, name)
}

func testAccResourceConnectorAzure_inheritFromDelegate_ua(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_azure_cloud_provider" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			credentials {
				type = "InheritFromDelegate"
				azure_inherit_from_delegate_details {
					auth {
						azure_msi_auth_ua {
							client_id = "client_id"
						}
						type = "UserAssignedManagedIdentity"
					}
				}
			}

			azure_environment_type = "AZURE"
			delegate_selectors = ["harness-delegate"]
		}
`, id, name)
}

func testAccResourceConnectorAzure_inheritFromDelegate_sa(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_azure_cloud_provider" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			credentials {
				type = "InheritFromDelegate"
				azure_inherit_from_delegate_details {
					auth {
						type = "SystemAssignedManagedIdentity"
					}
				}
			}

			azure_environment_type = "AZURE"
			delegate_selectors = ["harness-delegate"]
		}
`, id, name)
}
