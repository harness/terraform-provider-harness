package governance_enforcement_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccResourceRuleEnforcement_CCM32336_OutOfBandDeleteRecreates verifies that
// when a governance rule enforcement is deleted out-of-band (UI / direct API),
// the next terraform refresh treats the GET as "not found" and re-plans a
// create instead of erroring out with "giving up after 11 attempt(s)".
//
// Regression test for CCM-32336. The enforcement Read() routes through
// helpers.HandleReadApiError which clears state on 404 + ENTITY_NOT_FOUND.
func TestAccResourceRuleEnforcement_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_governance_rule_enforcement.test"
	awsAccountId := os.Getenv("AWS_ACCOUNT_ID")

	var enforcementIDBefore string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRuleEnforcementDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRuleEnforcement(name, awsAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					func(s *terraform.State) error {
						r := acctest.TestAccGetResource(resourceName, s)
						enforcementIDBefore = r.Primary.ID
						return nil
					},
				),
			},
			{
				PreConfig: func() {
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					if _, _, err := c.RuleEnforcementApi.DeleteRuleEnforcement(
						ctx, c.AccountId, enforcementIDBefore,
					); err != nil {
						t.Fatalf("CCM-32336: out-of-band delete failed: %v", err)
					}
				},
				Config:             testAccResourceRuleEnforcement(name, awsAccountId),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccResourceRuleEnforcement(name, awsAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrWith(resourceName, "id", func(value string) error {
						if value == "" || value == enforcementIDBefore {
							return fmt.Errorf("expected new enforcement id after recreate, got %q (before %q)", value, enforcementIDBefore)
						}
						return nil
					}),
				),
			},
		},
	})
}
