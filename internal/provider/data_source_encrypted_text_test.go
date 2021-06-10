package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/micahlmartin/terraform-provider-harness/internal/envvar"
	"github.com/micahlmartin/terraform-provider-harness/internal/utils"
)

func TestAccDataSourceEncryptedTextByName(t *testing.T) {

	// Setup
	var (
		expectedName = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		resourceName = "data.harness_encrypted_text.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEncryptedTextByName(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					resource.TestCheckResourceAttr(resourceName, "secret_manager_id", os.Getenv(envvar.HarnessAccountId)),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.application_filter_type", "ALL"),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", "NON_PRODUCTION_ENVIRONMENTS"),
				),
			},
		},
	})
}

func testAccDataSourceEncryptedTextByName(name string) string {
	return fmt.Sprintf(`
		resource "harness_encrypted_text" "test" {
			name              = "%[1]s"
			value             = "foo"
		
			usage_scope {
				application_filter_type = "ALL"
				environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
			}
		}

		data "harness_encrypted_text" "test" {
			name = harness_encrypted_text.test.name
		}
	`, name)
}
