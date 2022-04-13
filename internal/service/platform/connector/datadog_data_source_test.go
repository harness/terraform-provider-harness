package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorDatadog(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_datadog.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorDatadog(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://datadog.com"),
					resource.TestCheckResourceAttr(resourceName, "application_key_ref", "account.acctest_datadog_app_key"),
					resource.TestCheckResourceAttr(resourceName, "api_key_ref", "account.acctest_datadog_api_key"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorDatadog(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_datadog" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://datadog.com"
			delegate_selectors = ["harness-delegate"]
			application_key_ref = "account.acctest_datadog_app_key"
			api_key_ref = "account.acctest_datadog_api_key"
		}

		data "harness_platform_connector_datadog" "test" {
			identifier = harness_platform_connector_datadog.test.identifier
		}
	`, name)
}
