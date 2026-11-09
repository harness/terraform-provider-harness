package authentication_settings_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSSOOAuth(t *testing.T) {
	resourceName := "data.harness_platform_sso_oauth.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				// The data source reads the account live; depends_on makes the
				// read observe the providers configured by the resource.
				Config: testAccDataSourceSSOOAuth(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "allowed_providers.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "allowed_providers.*", "GOOGLE"),
				),
			},
		},
	})
}

func testAccDataSourceSSOOAuth() string {
	return `
		resource "harness_platform_sso_oauth" "test" {
			allowed_providers = ["GOOGLE"]
		}

		data "harness_platform_sso_oauth" "test" {
			depends_on = [harness_platform_sso_oauth.test]
		}
	`
}
