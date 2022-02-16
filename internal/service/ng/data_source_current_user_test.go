package ng_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCurrentUser(t *testing.T) {
	t.Skip("Skipping test until we can figure why the endpoint is erroring out")
	resourceName := "data.harness_current_user.test"

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
		data "harness_current_user" "test" {}
	`
}
