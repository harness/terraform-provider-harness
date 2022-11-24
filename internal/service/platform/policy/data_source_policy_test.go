package policy_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePolicy(t *testing.T) {
	id := t.Name() + utils.RandStringBytes(6)
	resourceName := "data.harness_platform_policy.test"
	rego := "package test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePolicy(id, rego),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", id),
					resource.TestCheckResourceAttr(resourceName, "rego", rego),
				),
			},
		},
	})
}

func testAccDataSourcePolicy(id, rego string) string {
	return fmt.Sprintf(`
		resource "harness_platform_policy" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			rego = "%[2]s"
		}

		data "harness_platform_policy" "test" {
			identifier = harness_platform_policy.test.identifier
			name = harness_platform_policy.test.name
			rego = "package test"
		}
	`, id, rego)
}
