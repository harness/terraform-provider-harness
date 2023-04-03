package connector_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorVault(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_vault.test"
		vaultToken   = os.Getenv("HARNESS_TEST_VAULT_SECRET")
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorVault(name, vaultToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
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
		},
	})
}

func testAccDataSourceConnectorVault(name string, vaultToken string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "%[2]s"
	}

	resource "harness_platform_connector_vault" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
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

	data "harness_platform_connector_vault" "test" {
		identifier = harness_platform_connector_vault.test.id
		name = harness_platform_connector_vault.test.name
	}
`, name, vaultToken)
}
