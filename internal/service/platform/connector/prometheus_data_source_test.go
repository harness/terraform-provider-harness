package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorPrometheus(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_prometheus.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorPrometheus(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://prometheus.com/"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "user_name", "user_name"),
					resource.TestCheckResourceAttr(resourceName, "password_ref", "account."+name),
					resource.TestCheckResourceAttr(resourceName, "headers.0.value_encrypted", "true"),
					resource.TestCheckResourceAttr(resourceName, "headers.0.key", "key"),
					resource.TestCheckResourceAttr(resourceName, "headers.0.value", "value"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorPrometheus(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]
		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
		lifecycle {
			ignore_changes = [
				value,
			]
		}
	}
		resource "harness_platform_connector_prometheus" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://prometheus.com/"
			delegate_selectors = ["harness-delegate"]
			user_name = "user_name"
			password_ref = "account.${harness_platform_secret_text.test.identifier}"
			headers {
				encrypted_value_ref = "account.${harness_platform_secret_text.test.identifier}"
				value_encrypted = true
				key = "key"
				value = "value"
			}
		}

		data "harness_platform_connector_prometheus" "test" {
			identifier = harness_platform_connector_prometheus.test.identifier
		}
	`, name)
}
