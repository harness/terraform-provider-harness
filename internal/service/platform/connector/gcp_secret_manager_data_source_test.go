package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorGcpSm(t *testing.T) {
	var (
		name = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpSM(name, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
		},
	})
}
func TestAccDataSourceConnectorGcpSmProjectLevel(t *testing.T) {
	t.Skip()
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		gcpName      = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpSMProjectLevel(name, name, gcpName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", gcpName),
					resource.TestCheckResourceAttr(resourceName, "name", gcpName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
		},
	})
}
func TestAccDataSourceConnectorGcpSmOrgLevel(t *testing.T) {
	t.Skip()
	var (
		name          = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		connectorName = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
		resourceName  = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpSMOrgLevel(name, connectorName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", connectorName),
					resource.TestCheckResourceAttr(resourceName, "name", connectorName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "false"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorGcpSmDefault(t *testing.T) {
	t.Skip()
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpSMDefault(name, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				),
			},
		},
	})
}
func TestAccDataSourceConnectorGcpSmDefaultProjectLevel(t *testing.T) {
	t.Skip()
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		gcpName      = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpSMDefaultProjectLevel(name, gcpName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", gcpName),
					resource.TestCheckResourceAttr(resourceName, "name", gcpName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				),
			},
		},
	})
}
func TestAccDataSourceConnectorGcpSmDefaultOrgLevel(t *testing.T) {
	t.Skip()
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		gcpname      = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
		resourceName = "data.harness_platform_connector_gcp_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcpSMDefaultOrgLevel(name, gcpname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", gcpname),
					resource.TestCheckResourceAttr(resourceName, "name", gcpname),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_default", "true"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorGcpSM(id string, name string) string {
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
			credentials_ref = "account.${harness_platform_secret_text.test.id}"
			delegate_selectors = ["harness-delegate"]
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}
		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			
		}
`, id, name)
}
func testAccDataSourceConnectorGcpSMProjectLevel(id string, name string, gcpName string) string {
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
		identifier = "%[1]s"
		name = "%[2]s"
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
			secret_manager_identifier = "%[1]s"
			value_type = "Reference"
			value = "secret"
			depends_on = [time_sleep.wait_5_seconds]
		}

		resource "time_sleep" "wait_5_seconds" {
			depends_on = [harness_platform_connector_azure_key_vault.test]
			create_duration = "5s"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[3]s"
			name = "%[3]s"
			description = "test"
			org_id= harness_platform_organization.test.id
			project_id=harness_platform_project.test.id
			tags = ["foo:bar"]
			credentials_ref = "${harness_platform_secret_text.test.id}"
			delegate_selectors = ["harness-delegate"]
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			create_duration = "4s"
		}
		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
			project_id = harness_platform_connector_gcp_secret_manager.test.project_id
		}
`, id, name, gcpName)
}

func testAccDataSourceConnectorGcpSMOrgLevel(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_connector_azure_key_vault" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
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
			secret_manager_identifier = "%[1]s"
			value_type = "Reference"
			value = "secret"
			depends_on = [time_sleep.wait_5_seconds]
		}

		resource "time_sleep" "wait_5_seconds" {
			depends_on = [harness_platform_connector_azure_key_vault.test]
			create_duration = "5s"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[2]s"
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
		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
			
		}
`, id, name)
}

func testAccDataSourceConnectorGcpSMDefault(id string, name string) string {
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
			credentials_ref = "account.${harness_platform_secret_text.test.id}"
			delegate_selectors = ["harness-delegate"]
			depends_on = [time_sleep.wait_4_seconds]
		}

		resource "time_sleep" "wait_4_seconds" {
			depends_on = [harness_platform_secret_text.test]
			destroy_duration = "4s"
		}
		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			
		}
`, id, name)
}
func testAccDataSourceConnectorGcpSMDefaultProjectLevel(id string, gcpname string) string {

	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}
	
	resource "harness_platform_project" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		org_id = harness_platform_organization.test.id
		color = "#472848"
	}

	resource "harness_platform_connector_azure_key_vault" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
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
			secret_manager_identifier = "%[1]s"
			value_type = "Reference"
			value = "secret"
			depends_on = [time_sleep.wait_5_seconds]
	}

	resource "time_sleep" "wait_5_seconds" {
		depends_on = [harness_platform_connector_azure_key_vault.test]
		create_duration = "5s"
	}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[2]s"
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

		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
			project_id =harness_platform_connector_gcp_secret_manager.test.project_id
			
		}
`, id, gcpname)
}
func testAccDataSourceConnectorGcpSMDefaultOrgLevel(id string, gcpname string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
	}

	resource "harness_platform_connector_azure_key_vault" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
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
			secret_manager_identifier = "%[1]s"
			value_type = "Reference"
			value = "secret"
			depends_on = [time_sleep.wait_5_seconds]
		}

		resource "time_sleep" "wait_5_seconds" {
			depends_on = [harness_platform_connector_azure_key_vault.test]
			create_duration = "5s"
		}

		resource "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = "%[2]s"
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
		data "harness_platform_connector_gcp_secret_manager" "test" {
			identifier = harness_platform_connector_gcp_secret_manager.test.identifier
			org_id = harness_platform_connector_gcp_secret_manager.test.org_id
		}
`, id, gcpname)
}
