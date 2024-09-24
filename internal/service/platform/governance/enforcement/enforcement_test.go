package governance_enforcement_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceRuleEnforcement(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_governance_rule_enforcement.test"
	awsAccountId := os.Getenv("AWS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRuleEnforcementDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRuleEnforcement(name, awsAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "execution_schedule", "0 0 * * * *"),
					resource.TestCheckResourceAttr(resourceName, "execution_timezone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "target_accounts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target_regions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_dry_run", "true"),
					resource.TestCheckResourceAttr(resourceName, "description", "Dummy"),
				),
			},
			{
				Config: testAccResourceRuleEnforcement(updatedName, awsAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "execution_schedule", "0 0 * * * *"),
					resource.TestCheckResourceAttr(resourceName, "execution_timezone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "target_accounts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target_regions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_dry_run", "true"),
					resource.TestCheckResourceAttr(resourceName, "description", "Dummy"),
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

func testAccResourceRuleEnforcement(name, awsAccountId string) string {
	return fmt.Sprintf(`
		resource "harness_governance_rule_enforcement" "test" {
			name               = "%[1]s"
			cloud_provider     = "AWS"
			rule_ids          = ["YW-qYiJRSaO3Fqei2EqqRQ"]
			rule_set_ids      = []
			execution_schedule = "0 0 * * * *"
			execution_timezone = "UTC"
			is_enabled         = true
			target_accounts    = ["%[2]s"]
			target_regions     = ["us-east-1"]
			is_dry_run         = true
			description        = "Dummy"
		}
	`, name, awsAccountId)
}

func testAccRuleEnforcementDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		ruleEnforcement, _ := testGetRuleEnforcement(resourceName, state)
		if ruleEnforcement != nil {
			return fmt.Errorf("Found rule enforcement: %s", ruleEnforcement.EnforcementName)
		}
		return nil
	}
}

func testGetRuleEnforcement(resourceName string, state *terraform.State) (*nextgen.EnforcementDetails, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	enforcementId := r.Primary.ID

	resp, _, err := c.RuleEnforcementApi.EnforcementDetails(ctx, c.AccountId, readRuleEnforcementRequest(enforcementId))

	if err != nil {
		return nil, err
	}

	if resp.Data != nil {
		return resp.Data, nil
	}
	return nil, nil
}

func readRuleEnforcementRequest(id string) *nextgen.RuleEnforcementApiEnforcementDetailsOpts {
	return &nextgen.RuleEnforcementApiEnforcementDetailsOpts{
		EnforcementId: optional.NewString(id),
	}
}
