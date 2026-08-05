package governance_rule_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceRule(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_governance_rule.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRule(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "description", "Dummy"),
					resource.TestCheckResourceAttr(resourceName, "rules_yaml", "policies:\n  - name: aws-list-ec2\n    resource: aws.ec2"),
				),
			},
			{
				Config: testAccResourceRule(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "description", "Dummy"),
					resource.TestCheckResourceAttr(resourceName, "rules_yaml", "policies:\n  - name: aws-list-ec2\n    resource: aws.ec2"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       false,
				ImportStateVerifyIgnore: []string{"identifier"},
				ImportStateIdFunc:       acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// TestAccResourceRule_CCM32336_OutOfBandDeleteRecreates verifies that when a
// governance rule (cloud asset policy) is deleted out-of-band (UI / direct API),
// the next terraform refresh treats the GET as "not found" and re-plans a create
// instead of erroring out with "giving up after 11 attempt(s)".
//
// Regression test for CCM-32336. The governance rule Read() routes through
// helpers.HandleReadApiError which clears state on 404 + ENTITY_NOT_FOUND.
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

func testAccResourceRule(name string) string {
	return fmt.Sprintf(`
		resource "harness_governance_rule" "test" {
			name               = "%[1]s"
			cloud_provider     = "AWS"
			description        = "Dummy"
			rules_yaml 		   = "policies:\n  - name: aws-list-ec2\n    resource: aws.ec2"
		}
	`, name)
}

func testAccRuleDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rule, _ := testGetRule(resourceName, state)
		if rule != nil {
			return fmt.Errorf("Found rule: %s", rule.Name)
		}
		return nil
	}
}

func testGetRule(resourceName string, state *terraform.State) (*nextgen.CcmRule, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	ruleId := r.Primary.ID

	resp, _, err := c.RuleApi.GetPolicies(ctx, readRuleRequest(ruleId), c.AccountId, nil)

	if err != nil {
		return nil, err
	}

	if len(resp.Data.Rules) > 0 {
		return &resp.Data.Rules[0], nil
	}
	return nil, nil
}

func readRuleRequest(id string) nextgen.ListDto {
	return nextgen.ListDto{
		Query: &nextgen.RuleRequest{
			PolicyIds: []string{id},
		},
	}
}
