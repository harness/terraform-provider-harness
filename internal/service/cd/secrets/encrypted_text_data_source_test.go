package secrets_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceEncryptedTextByName(t *testing.T) {

	// Setup
	var (
		expectedName = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		resourceName = "data.harness_encrypted_text.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEncryptedTextByName(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					resource.TestCheckResourceAttrSet(resourceName, "secret_manager_id"),
					resource.TestCheckResourceAttr(resourceName, "usage_scope.0.environment_filter_type", "NON_PRODUCTION_ENVIRONMENTS"),
				),
			},
		},
	})
}

func testAccDataSourceEncryptedTextByName(name string) string {
	return fmt.Sprintf(`
		data "harness_secret_manager" "test" {
			default = true
		}

		resource "harness_encrypted_text" "test" {
			name              = "%[1]s"
			value             = "foo"
			secret_manager_id = data.harness_secret_manager.test.id
			
			usage_scope {
				environment_filter_type = "NON_PRODUCTION_ENVIRONMENTS"
			}
		}

		data "harness_encrypted_text" "test" {
			name = harness_encrypted_text.test.name
		}
	`, name)
}
