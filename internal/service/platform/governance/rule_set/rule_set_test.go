package governance_rule_set_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceRuleSet(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_governance_rule_set.test"

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
					resource.TestCheckResourceAttr(resourceName, "rule_ids.#", "1"),
				),
			},
			{
				Config: testAccResourceRule(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "description", "Dummy"),
					resource.TestCheckResourceAttr(resourceName, "rule_ids.#", "1"),
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

func testAccResourceRule(name string) string {
	return fmt.Sprintf(`
		resource "harness_governance_rule_set" "test" {
			name               = "%[1]s"
			cloud_provider     = "AWS"
			description        = "Dummy"
			rule_ids 		   = ["1NBh8oKjQ4KmzvgkUqN5sQ"]
		}
	`, name)
}

func testAccRuleDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rule, _ := testGetRuleSet(resourceName, state)
		if rule != nil {
			return fmt.Errorf("Found rule set: %s", rule.Name)
		}
		return nil
	}
}

func testGetRuleSet(resourceName string, state *terraform.State) (*nextgen.RuleSet, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	ruleId := r.Primary.ID

	resp, _, err := c.RuleSetsApi.ListRuleSets(ctx, readRuleSetRequest(ruleId), c.AccountId)

	if err != nil {
		return nil, err
	}

	if len(resp.Data.RuleSet) > 0 {
		return &resp.Data.RuleSet[0], nil
	}
	return nil, nil
}

func readRuleSetRequest(id string) nextgen.CreateRuleSetFilterDto {
	return nextgen.CreateRuleSetFilterDto{
		RuleSet: &nextgen.RuleSetRequest{
			RuleSetIds: []string{id},
		},
	}
}
