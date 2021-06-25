package provider

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
					resource.TestCheckResourceAttrSet(resourceName, "secret_manager_id"),
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

			lifecycle {
				ignore_changes = [secret_manager_id]
			}
		}

		data "harness_encrypted_text" "test" {
			name = harness_encrypted_text.test.name
		}
	`, name)
}
