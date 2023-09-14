package connector_test

import (
	"fmt"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceConnector_customhealthsource(t *testing.T) {

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_customhealthsource.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnector_customhealthsource(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://prometheus.com/"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "method", "GET"),
					resource.TestCheckResourceAttr(resourceName, "headers.0.value_encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "headers.0.key", "key"),
					resource.TestCheckResourceAttr(resourceName, "headers.0.value", "value"),
				),
			},
			{
				Config: testAccResourceConnector_customhealthsource(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://prometheus.com/"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "method", "GET"),
					resource.TestCheckResourceAttr(resourceName, "headers.0.value_encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "headers.0.key", "key"),
					resource.TestCheckResourceAttr(resourceName, "headers.0.value", "value"),
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

func testAccResourceConnector_customhealthsource(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_customhealthsource" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://prometheus.com/"
			delegate_selectors = ["harness-delegate"]
			method = "GET"
			validation_path = "loki/api/v1/labels"
			headers {
				value_encrypted = false
				key = "key"
				value = "value"
			}
            params {
				encrypted_value_ref = "account.doNotDeleteHSM"
				value_encrypted = true
				key = "param1"
			}
		}`, id, name)
}
