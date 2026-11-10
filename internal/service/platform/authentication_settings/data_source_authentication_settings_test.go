package authentication_settings_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceAuthenticationSettings(t *testing.T) {
	resourceName := "data.harness_platform_authentication_settings.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAuthenticationSettings(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "two_factor_enabled"),
					resource.TestCheckResourceAttrSet(resourceName, "authentication_mechanism"),
					resource.TestCheckResourceAttrSet(resourceName, "whitelisted_domains.#"),
					resource.TestCheckResourceAttrSet(resourceName, "session_timeout_minutes"),
					resource.TestCheckResourceAttrSet(resourceName, "absolute_session_timeout_minutes"),
					resource.TestCheckResourceAttrSet(resourceName, "public_access_enabled"),
				),
			},
		},
	})
}

func testAccDataSourceAuthenticationSettings() string {
	return `
		data "harness_platform_authentication_settings" "test" {
		}
	`
}
