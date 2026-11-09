package authentication_settings_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceLoginSettings(t *testing.T) {
	resourceName := "data.harness_platform_login_settings.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceLoginSettings(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "user_lockout_policy.0.enable_lockout_policy"),
					resource.TestCheckResourceAttrSet(resourceName, "user_lockout_policy.0.number_of_failed_attempts_before_lockout"),
				),
			},
		},
	})
}

func testAccDataSourceLoginSettings() string {
	return `
		data "harness_platform_login_settings" "test" {
		}
	`
}
