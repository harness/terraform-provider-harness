package policyset_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePolicyset(t *testing.T) {
	id := t.Name() + utils.RandStringBytes(6)
	resourceName := "data.harness_platform_policyset.test"
	policyType := "pipeline"
	action := "onrun"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePolicyset(id, policyType, action),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", id),
					resource.TestCheckResourceAttr(resourceName, "type", policyType),
					resource.TestCheckResourceAttr(resourceName, "action", action),
				),
			},
		},
	})
}

func testAccDataSourcePolicyset(id, policyType, action string) string {
	return fmt.Sprintf(`
		resource "harness_platform_policyset" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			type = "%[2]s"
			action = "%[3]s"
		}

		data "harness_platform_policyset" "test" {
			identifier = harness_platform_policyset.test.identifier
			name = harness_platform_policyset.test.name
			type = "pipeline"
			action = "onrun"
		}
	`, id, policyType, action)
}
