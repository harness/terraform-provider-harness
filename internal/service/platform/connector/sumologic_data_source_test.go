package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorSumologic(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_sumologic.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorSumologic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://api.us2.sumologic.com/"),
					resource.TestCheckResourceAttr(resourceName, "access_id_ref", "account.acctest_sumo_access_id"),
					resource.TestCheckResourceAttr(resourceName, "access_key_ref", "account.acctest_sumo_access_key"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorSumologic(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_sumologic" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://api.us2.sumologic.com/"
			delegate_selectors = ["harness-delegate"]
			access_id_ref = "account.acctest_sumo_access_id"
			access_key_ref = "account.acctest_sumo_access_key"
		}

		data "harness_platform_connector_sumologic" "test" {
			identifier = harness_platform_connector_sumologic.test.identifier
		}
	`, name)
}
