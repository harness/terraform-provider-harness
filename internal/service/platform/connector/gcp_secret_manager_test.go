package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorGcpSM(t *testing.T) {
	t.Skip()

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_secret_manager.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorGcpSM(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
			{
				Config: testAccResourceConnectorGcpSM(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
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

func TestProjectResourceConnectorGcpSM(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	connectorName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(10))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_secret_manager.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceConnectorGcpSM(id, name, connectorName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
			{
				Config: testProjectResourceConnectorGcpSM(id, updatedName, connectorName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestOrgResourceConnectorGcpSM(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	connectorName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(10))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_secret_manager.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceConnectorGcpSM(id, name, connectorName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
			{
				Config: testOrgResourceConnectorGcpSM(id, updatedName, connectorName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceConnectorGcpSMDefault(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_secret_manager.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorGcpSMDefault(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				),
			},
			{
				Config: testAccResourceConnectorGcpSMDefault(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
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
func TestProjectResourceConnectorGcpSMDefault(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	connectorName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(10))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_secret_manager.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testProjectResourceConnectorGcpSMDefault(id, name, connectorName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				),
			},
			{
				Config: testProjectResourceConnectorGcpSMDefault(id, updatedName, connectorName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}
func TestOrgResourceConnectorGcpSMDefault(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	connectorName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(10))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_secret_manager.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testOrgResourceConnectorGcpSMDefault(id, name, connectorName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				),
			},
			{
				Config: testOrgResourceConnectorGcpSMDefault(id, updatedName, connectorName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: acctest.OrgResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func testAccResourceConnectorGcpSM(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			secret_manager_identifier = "azureSecretManager"
			value_type = "Reference"
			value = "secret"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
		
			delegate_selectors = ["harness-delegate"]
			credentials_ref = "account.${harness_platform_secret_text.test.id}"
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}
`, id, name)
}

func testProjectResourceConnectorGcpSM(id string, name string, connectorName string) string {
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

	resource "harness_platform_connector_azure_key_vault" "test" {
		identifier = "%[3]s"
		name = "%[3]s"
		description = "test"
		tags = ["foo:bar"]
		org_id= harness_platform_organization.test.id
		project_id=harness_platform_project.test.id
		client_id = "38fca8d7-4dda-41d5-b106-e5d8712b733a"
		secret_key = "account.azuretest"
		tenant_id = "b229b2bb-5f33-4d22-bce0-730f6474e906"
		vault_name = "Aman-test"
		subscription = "20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0"
		is_default = false

		azure_environment_type = "AZURE"
		depends_on = [time_sleep.wait_3_seconds]
	}

	resource "time_sleep" "wait_3_seconds" {
		depends_on = [harness_platform_project.test]
		create_duration = "3s"
	}
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			org_id= harness_platform_organization.test.id
			project_id=harness_platform_project.test.id
			tags = ["foo:bar"]
			secret_manager_identifier = "%[3]s"
			value_type = "Reference"
			value = "secret"
			depends_on = [time_sleep.wait_5_seconds]
	}

	resource "time_sleep" "wait_5_seconds" {
		depends_on = [harness_platform_connector_azure_key_vault.test]
		create_duration = "5s"
	}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			delegate_selectors = ["harness-delegate"]
			credentials_ref = "${harness_platform_secret_text.test.id}"
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			create_duration = "4s"
		}
`, id, name, connectorName)
}

func testOrgResourceConnectorGcpSM(id string, name string, connectorName string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_connector_azure_key_vault" "test" {
		identifier = "%[3]s"
		name = "%[3]s"
		description = "test"
		tags = ["foo:bar"]
		org_id= harness_platform_organization.test.id
		client_id = "38fca8d7-4dda-41d5-b106-e5d8712b733a"
		secret_key = "account.azuretest"
		tenant_id = "b229b2bb-5f33-4d22-bce0-730f6474e906"
		vault_name = "Aman-test"
		subscription = "20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0"
		is_default = false

		azure_environment_type = "AZURE"
		depends_on = [time_sleep.wait_3_seconds]
	}

	resource "time_sleep" "wait_3_seconds" {
		depends_on = [harness_platform_organization.test]
		create_duration = "3s"
	}
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			org_id= harness_platform_organization.test.id
			tags = ["foo:bar"]
			secret_manager_identifier = "%[3]s"
			value_type = "Reference"
			value = "secret"
			depends_on = [time_sleep.wait_5_seconds]
		}

		resource "time_sleep" "wait_5_seconds" {
			depends_on = [harness_platform_connector_azure_key_vault.test]
			create_duration = "5s"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			delegate_selectors = ["harness-delegate"]
			credentials_ref = "org.${harness_platform_secret_text.test.id}"
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			create_duration = "4s"
		}
`, id, name, connectorName)
}

func testAccResourceConnectorGcpSMDefault(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]
			secret_manager_identifier = "azureSecretManager"
			value_type = "Reference"
			value = "secret"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			is_default = true
		
			delegate_selectors = ["harness-delegate"]
			credentials_ref = "account.${harness_platform_secret_text.test.id}"
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}
`, id, name)
}

func testProjectResourceConnectorGcpSMDefault(id string, name string, connectorName string) string {
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

	resource "harness_platform_connector_azure_key_vault" "test" {
		identifier = "%[3]s"
		name = "%[3]s"
		description = "test"
		tags = ["foo:bar"]
		org_id= harness_platform_organization.test.id
		project_id=harness_platform_project.test.id
		client_id = "38fca8d7-4dda-41d5-b106-e5d8712b733a"
		secret_key = "account.azuretest"
		tenant_id = "b229b2bb-5f33-4d22-bce0-730f6474e906"
		vault_name = "Aman-test"
		subscription = "20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0"

		azure_environment_type = "AZURE"
		depends_on = [time_sleep.wait_3_seconds]
	}

	resource "time_sleep" "wait_3_seconds" {
		depends_on = [harness_platform_project.test]
		create_duration = "3s"
	}
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			org_id= harness_platform_organization.test.id
			project_id=harness_platform_project.test.id
			tags = ["foo:bar"]
			secret_manager_identifier = "%[3]s"
			value_type = "Reference"
			value = "secret"
			depends_on = [time_sleep.wait_5_seconds]
	}

	resource "time_sleep" "wait_5_seconds" {
		depends_on = [harness_platform_connector_azure_key_vault.test]
		create_duration = "5s"
	}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			delegate_selectors = ["harness-delegate"]
			credentials_ref = "${harness_platform_secret_text.test.id}"
			is_default = true
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			create_duration = "4s"
		}
`, id, name, connectorName)
}

func testOrgResourceConnectorGcpSMDefault(id string, name string, connectorName string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
	}

	resource "harness_platform_connector_azure_key_vault" "test" {
		identifier = "%[3]s"
		name = "%[3]s"
		description = "test"
		tags = ["foo:bar"]
		org_id= harness_platform_organization.test.id
		client_id = "38fca8d7-4dda-41d5-b106-e5d8712b733a"
		secret_key = "account.azuretest"
		tenant_id = "b229b2bb-5f33-4d22-bce0-730f6474e906"
		vault_name = "Aman-test"
		subscription = "20d6a917-99fa-4b1b-9b2e-a3d624e9dcf0"
		azure_environment_type = "AZURE"
		depends_on = [time_sleep.wait_3_seconds]
	}

	resource "time_sleep" "wait_3_seconds" {
		depends_on = [harness_platform_organization.test]
		create_duration = "3s"
	}
		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			org_id= harness_platform_organization.test.id
			tags = ["foo:bar"]
			secret_manager_identifier = "%[3]s"
			value_type = "Reference"
			value = "secret"
			depends_on = [time_sleep.wait_5_seconds]
		}

		resource "time_sleep" "wait_5_seconds" {
			depends_on = [harness_platform_connector_azure_key_vault.test]
			create_duration = "5s"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			org_id = harness_platform_organization.test.id
			delegate_selectors = ["harness-delegate"]
			is_default = true
			credentials_ref = "org.${harness_platform_secret_text.test.id}"
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			create_duration = "4s"
		}
`, id, name, connectorName)
}
