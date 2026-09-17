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

	policyIdentifier := fmt.Sprintf("policy%s", utils.RandStringBytes(6))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePolicyset(id, policyType, action, policyIdentifier),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", id),
					resource.TestCheckResourceAttr(resourceName, "type", policyType),
					resource.TestCheckResourceAttr(resourceName, "action", action),
					resource.TestCheckResourceAttr(resourceName, "policy_references.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "policy_references.*", map[string]string{
						"identifier": policyIdentifier,
						"severity":   "warning",
					}),
				),
			},
		},
	})
}

func testAccDataSourcePolicyset(id, policyType, action, policyIdentifier string) string {
	return fmt.Sprintf(`
		resource "harness_platform_policy" "test" {
			identifier = "%[4]s"
			name       = "%[4]s"
			rego       = "some text"
		}

		resource "harness_platform_policyset" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			type = "%[2]s"
			action = "%[3]s"

			policy_references {
				identifier = harness_platform_policy.test.identifier
				severity   = "warning"
			}
		}

		data "harness_platform_policyset" "test" {
			identifier = harness_platform_policyset.test.identifier
			name = harness_platform_policyset.test.name
			type = "pipeline"
			action = "onrun"
		}
	`, id, policyType, action, policyIdentifier)
}
