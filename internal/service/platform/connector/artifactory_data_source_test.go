package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorArtifactory(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_artifactory.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnector_artifactory(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://artifactory.example.com"),
					resource.TestCheckResourceAttr(resourceName, "delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.username", "admin"),
				),
			},
		},
	})
}

func testAccDataSourceConnector_artifactory(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_secret_text" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		secret_manager_identifier = "harnessSecretManager"
		value_type = "Inline"
		value = "secret"
	}
		resource "harness_platform_connector_artifactory" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://artifactory.example.com"
			delegate_selectors = ["harness-delegate"]
			credentials {	
				username = "admin"
				password_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		data "harness_platform_connector_artifactory" "test" {
			identifier = harness_platform_connector_artifactory.test.identifier
		}
	`, name)
}
