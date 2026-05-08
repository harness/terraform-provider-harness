package governance_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Blocked on CCM-32499: provider readRuleResponse panics with
// "index out of range [0] with length 0" because POST /governance/rule/list
// returns HTTP 200 with {"Rules": []} for a deleted rule and the provider does
// not bounds-check before accessing Rules[0].
//
// TestAccResourceRule_CCM32336_OutOfBandDeleteRecreates verifies that when a
// governance rule (cloud asset policy) is deleted out-of-band (UI / direct API),
// the next terraform refresh treats the GET as "not found" and re-plans a create
// instead of erroring out with "giving up after 11 attempt(s)".
//
// Regression test for CCM-32336 (the bug class is: a NextGen GET endpoint
// returning HTTP 500 for a deleted entity causes terraform plan to fail).
// The governance rule resource Read() routes through helpers.HandleReadApiError
// which clears state on 404 + ENTITY_NOT_FOUND.
func TestAccResourceRule_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_governance_rule.test"

	var ruleIDBefore string

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
						ruleIDBefore = r.Primary.ID
						return nil
					},
				),
			},
			{
				PreConfig: func() {
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					if _, _, err := c.RuleApi.DeleteRule(ctx, c.AccountId, ruleIDBefore); err != nil {
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
						if value == "" || value == ruleIDBefore {
							return fmt.Errorf("expected new rule id after recreate, got %q (before %q)", value, ruleIDBefore)
						}
						return nil
					}),
				),
			},
		},
	})
}
