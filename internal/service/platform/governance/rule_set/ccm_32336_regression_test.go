package governance_rule_set_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccResourceRuleSet_CCM32336_OutOfBandDeleteRecreates verifies that when a
// governance rule set is deleted out-of-band (UI / direct API), the next
// terraform refresh treats the GET as "not found" and re-plans a create
// instead of erroring out with "giving up after 11 attempt(s)".
//
// Regression test for CCM-32336. The governance rule set Read() routes through
// helpers.HandleReadApiError which clears state on 404 + ENTITY_NOT_FOUND.
func TestAccResourceRuleSet_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_governance_rule_set.test"

	var ruleSetIDBefore string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRule(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					func(s *terraform.State) error {
						r := acctest.TestAccGetResource(resourceName, s)
						ruleSetIDBefore = r.Primary.ID
						return nil
					},
				),
			},
			{
				PreConfig: func() {
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					if _, _, err := c.RuleSetsApi.DeleteRuleSet(ctx, c.AccountId, ruleSetIDBefore); err != nil {
						t.Fatalf("CCM-32336: out-of-band delete failed: %v", err)
					}
				},
				Config:             testAccResourceRule(name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccResourceRule(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrWith(resourceName, "id", func(value string) error {
						if value == "" || value == ruleSetIDBefore {
							return fmt.Errorf("expected new rule set id after recreate, got %q (before %q)", value, ruleSetIDBefore)
						}
						return nil
					}),
				),
			},
		},
	})
}
