package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorGcp(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorGcp(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
				),
			},
		},
	})
}

func TestOidcDataSourceConnectorGcp(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_gcp.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testOidcDataSourceConnectorGcp(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.delegate_selectors.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.workload_pool_id", "harness-pool-test"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.provider_id", "harness"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.gcp_project_id", "1234567"),
					resource.TestCheckResourceAttr(resourceName, "oidc_authentication.0.service_account_email", "harness.sample.iam.gserviceaccount.com"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorGcp(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_gcp" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			inherit_from_delegate {
				delegate_selectors = ["harness-delegate"]
			}
		}

		data "harness_platform_connector_gcp" "test" {
			identifier = harness_platform_connector_gcp.test.identifier
		}
	`, name)
}

func testOidcDataSourceConnectorGcp(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_gcp" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			oidc_authentication {
				workload_pool_id = "harness-pool-test"
				provider_id = "harness"
				gcp_project_id = "1234567"
				service_account_email = "harness.sample.iam.gserviceaccount.com"
				delegate_selectors = ["harness-delegate"]
			}
		}

		data "harness_platform_connector_gcp" "test" {
			identifier = harness_platform_connector_gcp.test.identifier
		}
	`, name)
}
