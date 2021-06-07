package provider

import (
	"fmt"
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
				Config: testAccResourceApplication(t.Name()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("harness_application.foo", "name", regexp.MustCompile(fmt.Sprintf("^%s", t.Name()))),
					resource.TestMatchResourceAttr("harness_application.foo", "description", regexp.MustCompile("^some")),
				),
			},
			{
				Config: testAccResourceApplication(t.Name() + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("harness_application.foo", "name", regexp.MustCompile(fmt.Sprintf("^%s-updated", t.Name()))),
					resource.TestMatchResourceAttr("harness_application.foo", "description", regexp.MustCompile("^some")),
				),
			},
		},
	})
}

func testAccResourceApplication(name string) string {
	return fmt.Sprintf(`
		resource "harness_application" "foo" {
			name = "%s"
			description = "some app description here"
		}
`, name)
}
