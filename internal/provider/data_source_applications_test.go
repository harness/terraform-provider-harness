package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceApplication(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceApplication("GEHhvKUCTiiY_MWsUfbRLA"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.harness_application.foo", "name", regexp.MustCompile("changed")),
				),
			},
		},
	})
}

func testAccDataSourceApplication(appId string) string {
	f := testProvider()
	if f != "" {
		fmt.Println("here")
	}
	return fmt.Sprintf(`
		data "harness_application" "foo" {
			id = "%s"
		}
	`, appId)
}
