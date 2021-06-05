package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceApplication(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplication,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"harness_application.foo", "sample_attribute", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccResourceApplication = `
resource "harness_application" "foo" {
  name = "my_app"
	description = "some app description here"
}
`
