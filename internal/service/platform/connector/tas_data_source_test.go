package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorTas(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_tas.test"
		id           = name
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnector_tas(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.type", "ManualConfig"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.endpoint_url", "https://tas.example.com"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.tas_manual_details.0.password_ref", "account."+id),
				),
			},
		},
	})
}

func testAccDataSourceConnector_tas(id, name string) string {
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

	resource "harness_platform_connector_tas" "test" {
		identifier = "%[1]s"
		name = "%[2]s"
		description = "test"
		tags = ["foo:bar"]

		credentials {
			type = "ManualConfig"
			tas_manual_details {
				endpoint_url = "https://tas.example.com"
				username = "admin"
				password_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		delegate_selectors = ["harness-delegate"]
	}

	data "harness_platform_connector_tas" "test" {
		identifier = harness_platform_connector_tas.test.identifier
	}
`, id, name)
}
