package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorAwsSm(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_aws_secret_manager.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsSm(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "secret_name_prefix", "test"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.inherit_from_delegate", "true"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorAwsSm(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_aws_secret_manager" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_name_prefix = "test"
			region = "us-east-1"
			delegate_selectors = ["harness-delegate"]
			credentials {
				inherit_from_delegate = true
			}
		}

		data "harness_platform_connector_aws_secret_manager" "test" {
			identifier = harness_platform_connector_aws_secret_manager.test.identifier
			name = harness_platform_connector_aws_secret_manager.test.name
		}
	`, name)
}
