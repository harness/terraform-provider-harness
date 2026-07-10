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
	disabledName := fmt.Sprintf("%s_disabled", updatedName)
	dryRunDisabledName := fmt.Sprintf("%s_dryrun_disabled", disabledName)
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
				Config: testAccResourceRuleEnforcementWithEnabled(disabledName, awsAccountId, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", disabledName),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "execution_schedule", "0 0 * * * *"),
					resource.TestCheckResourceAttr(resourceName, "execution_timezone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "target_accounts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target_regions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_dry_run", "true"),
					resource.TestCheckResourceAttr(resourceName, "description", "Dummy"),
				),
			},
			{
				Config: testAccResourceRuleEnforcementWithBooleans(dryRunDisabledName, awsAccountId, false, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", dryRunDisabledName),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "execution_schedule", "0 0 * * * *"),
					resource.TestCheckResourceAttr(resourceName, "execution_timezone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "target_accounts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target_regions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_dry_run", "false"),
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

func TestAccResourceRuleEnforcementWithFalseBooleans(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_governance_rule_enforcement.test"
	awsAccountId := os.Getenv("AWS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccRuleEnforcementDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRuleEnforcementWithFalseBooleans(name, awsAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "cloud_provider", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "execution_schedule", "0 0 * * * *"),
					resource.TestCheckResourceAttr(resourceName, "execution_timezone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "target_accounts.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "target_regions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_dry_run", "false"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test false booleans"),
				),
			},
		},
	})
}

func testAccResourceRuleEnforcement(name, awsAccountId string) string {
	return testAccResourceRuleEnforcementWithBooleans(name, awsAccountId, true, true)
}

func testAccResourceRuleEnforcementWithEnabled(name, awsAccountId string, isEnabled bool) string {
	return testAccResourceRuleEnforcementWithBooleans(name, awsAccountId, isEnabled, true)
}

func testAccResourceRuleEnforcementWithBooleans(name, awsAccountId string, isEnabled bool, isDryRun bool) string {
	return fmt.Sprintf(`
		resource "harness_governance_rule" "rule" {
			name           = "%[1]s_rule"
			cloud_provider = "AWS"
			description    = "Dummy"
			rules_yaml     = "policies:\n  - name: aws-list-ec2\n    resource: aws.ec2"
		}

		resource "harness_governance_rule_enforcement" "test" {
			name               = "%[1]s"
			cloud_provider     = "AWS"
			rule_ids           = [harness_governance_rule.rule.rule_id]
			rule_set_ids       = []
			execution_schedule = "0 0 * * * *"
			execution_timezone = "UTC"
			is_enabled         = %[3]t
			target_accounts    = ["%[2]s"]
			target_regions     = ["us-east-1"]
			is_dry_run         = %[4]t
			description        = "Dummy"
		}
	`, name, awsAccountId, isEnabled, isDryRun)
}

func testAccResourceRuleEnforcementWithFalseBooleans(name, awsAccountId string) string {
	return fmt.Sprintf(`
		resource "harness_governance_rule" "rule" {
			name           = "%[1]s_rule"
			cloud_provider = "AWS"
			description    = "Test false booleans"
			rules_yaml     = "policies:\n  - name: aws-list-ec2\n    resource: aws.ec2"
		}

		resource "harness_governance_rule_enforcement" "test" {
			name               = "%[1]s"
			cloud_provider     = "AWS"
			rule_ids           = [harness_governance_rule.rule.rule_id]
			rule_set_ids       = []
			execution_schedule = "0 0 * * * *"
			execution_timezone = "UTC"
			is_enabled         = false
			target_accounts    = ["%[2]s"]
			target_regions     = ["us-east-1"]
			is_dry_run         = false
			description        = "Test false booleans"
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
