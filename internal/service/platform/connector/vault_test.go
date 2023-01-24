package connector_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorVault_Token(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_vault.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorVault_token(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "base_path", "harness"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "renewal_interval_minutes", "0"),
					resource.TestCheckResourceAttr(resourceName, "secret_engine_manually_configured", "true"),
					resource.TestCheckResourceAttr(resourceName, "access_type", "TOKEN"),
				),
			},
			{
				Config: testAccResourceConnectorVault_token(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "base_path", "harness"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "renewal_interval_minutes", "0"),
					resource.TestCheckResourceAttr(resourceName, "secret_engine_manually_configured", "true"),
					resource.TestCheckResourceAttr(resourceName, "access_type", "TOKEN"),
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

func TestAccResourceConnectorVault_VaultAgent(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_vault.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorVault_vault_agent(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "base_path", "base_path"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "renewal_interval_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "secret_engine_manually_configured", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_vault_agent", "true"),
					resource.TestCheckResourceAttr(resourceName, "sink_path", "sink_path"),
					resource.TestCheckResourceAttr(resourceName, "access_type", "VAULT_AGENT"),
				),
			},
			{
				Config: testAccResourceConnectorVault_vault_agent(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "base_path", "base_path"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "renewal_interval_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "secret_engine_manually_configured", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_vault_agent", "true"),
					resource.TestCheckResourceAttr(resourceName, "sink_path", "sink_path"),
					resource.TestCheckResourceAttr(resourceName, "access_type", "VAULT_AGENT"),
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

func TestAccResourceConnectorVault_K8sAuth(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_vault.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorVault_k8s_auth(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "base_path", "base_path"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "renewal_interval_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "secret_engine_manually_configured", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_vault_agent", "false"),
					resource.TestCheckResourceAttr(resourceName, "access_type", "K8s_AUTH"),
				),
			},
			{
				Config: testAccResourceConnectorVault_k8s_auth(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "base_path", "base_path"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "renewal_interval_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "secret_engine_manually_configured", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_vault_agent", "false"),
					resource.TestCheckResourceAttr(resourceName, "access_type", "K8s_AUTH"),
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

func TestAccResourceConnectorVault_AppRole(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_vault.test"
	vault_sercet := os.Getenv("HARNESS_TEST_VAULT_SECRET")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorVault_app_role(id, name, vault_sercet),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "base_path", "vikas-test/"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "renewal_interval_minutes", "60"),
					resource.TestCheckResourceAttr(resourceName, "secret_engine_manually_configured", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_vault_agent", "false"),
					resource.TestCheckResourceAttr(resourceName, "access_type", "APP_ROLE"),
				),
			},
			{
				Config: testAccResourceConnectorVault_app_role(id, updatedName, vault_sercet),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "base_path", "vikas-test/"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "renewal_interval_minutes", "60"),
					resource.TestCheckResourceAttr(resourceName, "secret_engine_manually_configured", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_vault_agent", "false"),
					resource.TestCheckResourceAttr(resourceName, "access_type", "APP_ROLE"),
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

func TestAccResourceConnectorVault_AwsAuth(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_vault.test"
	vault_sercet := os.Getenv("HARNESS_TEST_VAULT_SECRET")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorVault_aws_auth(id, name, vault_sercet),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "base_path", "vikas-test/"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "renewal_interval_minutes", "60"),
					resource.TestCheckResourceAttr(resourceName, "secret_engine_manually_configured", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_vault_agent", "false"),
					resource.TestCheckResourceAttr(resourceName, "access_type", "AWS_IAM"),
				),
			},
			{
				Config: testAccResourceConnectorVault_aws_auth(id, updatedName, vault_sercet),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "base_path", "vikas-test/"),
					resource.TestCheckResourceAttr(resourceName, "is_read_only", "false"),
					resource.TestCheckResourceAttr(resourceName, "renewal_interval_minutes", "60"),
					resource.TestCheckResourceAttr(resourceName, "secret_engine_manually_configured", "true"),
					resource.TestCheckResourceAttr(resourceName, "use_vault_agent", "false"),
					resource.TestCheckResourceAttr(resourceName, "access_type", "AWS_IAM"),
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

func testAccResourceConnectorVault_aws_auth(id string, name string, vault_secret string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "%[3]s"
		lifecycle {
			ignore_changes = [
				value,
			]
		}
	}

	resource "harness_platform_connector_vault" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		aws_region = "us-east-2"
		base_path = "vikas-test/"
		access_type = "AWS_IAM"
		default = false
		xvault_aws_iam_server_id = "account.${harness_platform_secret_text.test.id}"
		read_only = true
		renewal_interval_minutes = 60
		secret_engine_manually_configured = true
		secret_engine_name = "harness-test"
		secret_engine_version = 2
		vault_aws_iam_role = "570acf09-ef2a-144b-2fb0-14a42e06ffe3"
		use_aws_iam = true
		use_k8s_auth = false
		use_vault_agent = false
		delegate_selectors = ["harness-delegate"]
		vault_url = "https://vaultqa.harness.io"

		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}
	`, id, name, vault_secret)
}

func testAccResourceConnectorVault_app_role(id string, name string, vault_secret string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "%[3]s"
		lifecycle {
			ignore_changes = [
				value,
			]
		}
	}

	resource "harness_platform_connector_vault" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		app_role_id = "570acf09-ef2a-144b-2fb0-14a42e06ffe3"
		base_path = "vikas-test/"
		access_type = "APP_ROLE"
		default = false
		secret_id = "account.${harness_platform_secret_text.test.id}"
		read_only = true
		renewal_interval_minutes = 60
		secret_engine_manually_configured = true
		secret_engine_name = "harness-test"
		secret_engine_version = 2
		use_aws_iam = false
		use_k8s_auth = false
		use_vault_agent = false
		renew_app_role_token = true
		delegate_selectors = ["harness-delegate"]
		vault_url = "https://vaultqa.harness.io"

		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}
	`, id, name, vault_secret)
}

func testAccResourceConnectorVault_k8s_auth(id string, name string) string {
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

	resource "harness_platform_connector_vault" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		auth_token = "account.${harness_platform_secret_text.test.id}"
		base_path = "base_path"
		access_type = "K8s_AUTH"
		default = false
		k8s_auth_endpoint = "k8s_auth_endpoint"
		namespace = "namespace"
		read_only = true
		renewal_interval_minutes = 10
		secret_engine_manually_configured = true
		secret_engine_name = "secret_engine_name"
		secret_engine_version = 2
		service_account_token_path = "service_account_token_path"
		use_aws_iam = false
		use_k8s_auth = true
		use_vault_agent = false
		vault_k8s_auth_role = "vault_k8s_auth_role"
		vault_aws_iam_role = "vault_aws_iam_role"
		delegate_selectors = ["harness-delegate"]
		vault_url = "https://vault_url.com"

		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}
	`, id, name)
}

func testAccResourceConnectorVault_vault_agent(id string, name string) string {
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

	resource "harness_platform_connector_vault" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		auth_token = "account.${harness_platform_secret_text.test.id}"
		base_path = "base_path"
		access_type = "VAULT_AGENT"
		default = false
		namespace = "namespace"
		read_only = true
		renewal_interval_minutes = 10
		secret_engine_manually_configured = true
		secret_engine_name = "secret_engine_name"
		secret_engine_version = 2
		use_aws_iam = false
		use_k8s_auth = false
		use_vault_agent = true
		sink_path = "sink_path"
		delegate_selectors = ["harness-delegate"]
		vault_url = "https://vault_url.com"

		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}
	`, id, name)
}

func testAccResourceConnectorVault_token(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "%[3]s"
	}

	resource "harness_platform_connector_vault" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		auth_token = "account.${harness_platform_secret_text.test.id}"
		base_path = "harness"
		access_type = "TOKEN"
		default = false
		read_only = true
		renewal_interval_minutes = 0
		secret_engine_manually_configured = true
		secret_engine_name = "QA_Secrets"
		secret_engine_version = 2
		use_aws_iam = false
		use_k8s_auth = false
		vault_url = "https://vaultqa.harness.io"

		depends_on = [time_sleep.wait_4_seconds]
	}

	resource "time_sleep" "wait_4_seconds" {
		depends_on = [harness_platform_secret_text.test]
		destroy_duration = "4s"
	}
	`, id, name, helpers.TestEnvVars.VaultRootToken.Get())
}
