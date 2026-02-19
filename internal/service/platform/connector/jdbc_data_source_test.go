package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorJDBC(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_jdbc.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorJDBC(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "jdbc:sqlserver://1.2.3;trustServerCertificate=true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.username_password.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.username_password.0.password_ref", fmt.Sprintf("account.%s", name)),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.auth_type", "UsernamePassword"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorJDBCDefaultAuth(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_jdbc.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorJDBCDefaultAuth(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "jdbc:sqlserver://1.2.3;trustServerCertificate=true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.username_password.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.username_password.0.password_ref", fmt.Sprintf("account.%s", name)),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.auth_type", "UsernamePassword"),
				),
			},
		},
	})
}

func TestAccDataSourceConnectorJDBCServiceAccountAuth(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_jdbc.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorJDBCServiceAccountAuth(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "jdbc:sqlserver://1.2.3;trustServerCertificate=true"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.service_account.0.token_ref", fmt.Sprintf("account.%s", name)),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.auth_type", "ServiceAccount"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorJDBC(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_jdbc" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		url = "jdbc:sqlserver://1.2.3;trustServerCertificate=true"
		delegate_selectors = ["harness-delegate"]
		credentials {
			auth_type = "UsernamePassword"
			username_password {
				username = "admin"
				password_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}
		depends_on = [harness_platform_secret_text.test]
	}

	data "harness_platform_connector_jdbc" "test" {
		identifier = harness_platform_connector_jdbc.test.identifier
	}
	`, name)
}

func testAccDataSourceConnectorJDBCDefaultAuth(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_jdbc" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		url = "jdbc:sqlserver://1.2.3;trustServerCertificate=true"
		delegate_selectors = ["harness-delegate"]
		credentials {
			username_password {
				username = "admin"
				password_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}
		depends_on = [harness_platform_secret_text.test]
	}

	data "harness_platform_connector_jdbc" "test" {
		identifier = harness_platform_connector_jdbc.test.identifier
	}
	`, name)
}

func testAccDataSourceConnectorJDBCServiceAccountAuth(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_connector_jdbc" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		url = "jdbc:sqlserver://1.2.3;trustServerCertificate=true"
		delegate_selectors = ["harness-delegate"]
		credentials {
			auth_type = "ServiceAccount"
			service_account {
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}
		depends_on = [harness_platform_secret_text.test]
	}

	data "harness_platform_connector_jdbc" "test" {
		identifier = harness_platform_connector_jdbc.test.identifier
	}
	`, name)
}

func TestAccDataSourceConnectorJDBCKeyPairAuth(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_jdbc.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorJDBCKeyPairAuth(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "jdbc:snowflake://account.snowflakecomputing.com?warehouse=warehouse_name&db=db_name&schema=schema_name"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.key_pair.0.username", "snowflake_user"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.key_pair.0.private_key_file_ref", fmt.Sprintf("account.%s", name)),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.key_pair.0.private_key_passphrase_ref", fmt.Sprintf("account.%s_passphrase", name)),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.auth_type", "KeyPair"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorJDBCKeyPairAuth(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}

	resource "harness_platform_secret_text" "passphrase" {
		identifier = "%[1]s_passphrase"
		name = "%[1]s_passphrase"
		description = "test passphrase"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "passphrase"
	}

	resource "harness_platform_connector_jdbc" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		url = "jdbc:snowflake://account.snowflakecomputing.com?warehouse=warehouse_name&db=db_name&schema=schema_name"
		delegate_selectors = ["harness-delegate"]
		credentials {
			auth_type = "KeyPair"
			key_pair {
				username = "snowflake_user"
				private_key_file_ref = "account.${harness_platform_secret_text.test.id}"
				private_key_passphrase_ref = "account.${harness_platform_secret_text.passphrase.id}"
			}
		}
		depends_on = [harness_platform_secret_text.test, harness_platform_secret_text.passphrase]
	}

	data "harness_platform_connector_jdbc" "test" {
		identifier = harness_platform_connector_jdbc.test.identifier
	}
	`, name)
}
