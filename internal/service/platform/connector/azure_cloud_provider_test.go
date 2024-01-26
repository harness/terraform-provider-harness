//go:build connectors || cd
// +build connectors cd

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
				Config: testAccResourceConnector_Azure(id, name, testCredentials_Azure_inheritFromDelegate_sa(), false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceConnector_Azure(id, updatedName, testCredentials_Azure_inheritFromDelegate_sa(), false),
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

func TestOrgResourceConnectorAzure_InheritFromDelegate_UA(t *testing.T) {
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
				Config: testOrgResoureConnector_Azure(id, name, testCredentials_Azure_inheritFromDelegate_ua(), false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testOrgResoureConnector_Azure(id, updatedName, testCredentials_Azure_inheritFromDelegate_ua(), false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
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
				Config: testAccResourceConnector_Azure(id, name, testCredentials_Azure_inheritFromDelegate_ua(), false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceConnector_Azure(id, updatedName, testCredentials_Azure_inheritFromDelegate_ua(), false),
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
				Config: testAccResourceConnector_Azure(id, name, testCredentials_Azure_manualDetails_certificate(), false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceConnector_Azure(id, updatedName, testCredentials_Azure_manualDetails_certificate(), false),
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

func TestProjectResourceConnectorAzure_ManualDetails_Certificate(t *testing.T) {

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
				Config: testProjectResourceConnector_Azure(id, name, testCredentials_Azure_manualDetails_certificate(), false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testProjectResourceConnector_Azure(id, updatedName, testCredentials_Azure_manualDetails_certificate(), false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
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
				Config: testAccResourceConnector_Azure(id, name, testCredentials_Azure_manualDetails_secret(), false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
			{
				Config: testAccResourceConnector_Azure(id, updatedName, testCredentials_Azure_manualDetails_secret(), false),
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

func TestAccResourceConnectorAzure_ForceDelete(t *testing.T) {

	id := fmt.Sprintf("ConnectorAzure_ForceDelete"+"_%s", utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_azure_cloud_provider.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_Azure(id, name, testCredentials_Azure_inheritFromDelegate_sa(), true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "force_delete", "true"),
				),
			},
			{
				Config: testAccResourceConnector_Azure(id, name, testCredentials_Azure_inheritFromDelegate_sa(), true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "force_delete", "true"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_delete"},
			},
		},
	})
}

func testAccResourceConnector_Azure(id string, name string, credentials string, force_delete bool) string {
	return testConnector_Azure(id, name, "", "", credentials, force_delete)
}

func testOrgResoureConnector_Azure(id string, name string, credentials string, force_delete bool) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	%s
	`, id, name, testConnector_Azure(id, name, "harness_platform_organization.test.id", "", credentials, force_delete))
}

func testProjectResourceConnector_Azure(id string, name string, credentials string, force_delete bool) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	%s
	`, id, name, testConnector_Azure(id, name, "harness_platform_organization.test.id", "harness_platform_project.test.id", credentials, force_delete))
}

func testConnector_Azure(id string, name string, org string, project string, credentials string, force_delete bool) string {
	org_template := ""
	project_template := ""
	force_delete_template := ""
	if len(org) > 0 {
		org_template = fmt.Sprintf("org_id = %s", org)
	}

	if len(project) > 0 {
		project_template = fmt.Sprintf("project_id = %s", project)
	}

	if force_delete {
		force_delete_template = fmt.Sprintf("force_delete = %t", force_delete)
	}

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

	resource "harness_platform_connector_azure_cloud_provider" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]
		%[3]s
		%[4]s

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
		%[5]s
	}
	`, id, name, org_template, project_template, force_delete_template)
}

func testCredentials_Azure_inheritFromDelegate_sa() string {
	return fmt.Sprint(`
	credentials {
		type = "InheritFromDelegate"
		azure_inherit_from_delegate_details {
			auth {
				type = "SystemAssignedManagedIdentity"
			}
		}
	}
	`)
}

func testCredentials_Azure_inheritFromDelegate_ua() string {
	return fmt.Sprint(`
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
	`)
}

func testCredentials_Azure_manualDetails_certificate() string {
	return fmt.Sprint(`
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
	`)
}

func testCredentials_Azure_manualDetails_secret() string {
	return fmt.Sprint(`
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
	`)
}
