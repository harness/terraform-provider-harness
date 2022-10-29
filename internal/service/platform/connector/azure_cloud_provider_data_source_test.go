package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorAzureCloudProvider(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_azure_cloud_provider.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnector_azureCloudPrvider(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
				),
			},
		},
	})
}

func testAccDataSourceConnector_azureCloudPrvider(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_azure_cloud_provider" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		credentials {
			type = "InheritFromDelegate"
			azure_inherit_from_delegate_details {
				auth {
					azure_msi_auth_ua {
						client_id = "client_id"
					}
					type = "UserAssignedManagedIdentity"
				}
			}
		}

		azure_environment_type = "AZURE"
		delegate_selectors = ["harness-delegate"]
	}

	data "harness_platform_connector_azure_cloud_provider" "test" {
		identifier = harness_platform_connector_azure_cloud_provider.test.identifier
	}
`, name)
}
