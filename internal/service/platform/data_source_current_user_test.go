package platform_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCurrentUser(t *testing.T) {
	t.Skip("Only works with personal access tokens.")

	resourceName := "data.harness_platform_current_user.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCurrentUser(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "micah.martin@harness.io"),
				),
			},
		},
	})
}

func testAccDataSourceCurrentUser() string {
	return `
		data "harness_platform_current_user" "test" {}
	`
}
