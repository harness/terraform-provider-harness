package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorSplunk(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_splunk.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorSplunk(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://splunk.com/"),
					resource.TestCheckResourceAttr(resourceName, "username", "admin"),
					resource.TestCheckResourceAttr(resourceName, "account_id", "splunk_account_id"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorSplunk(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_splunk" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://splunk.com/"
			delegate_selectors = ["harness-delegate"]
			account_id = "splunk_account_id"
			username = "admin"
			password_ref = "account.doNotDeleteHSM"
		}

		data "harness_platform_connector_splunk" "test" {
			identifier = harness_platform_connector_splunk.test.identifier
		}
	`, name)
}
