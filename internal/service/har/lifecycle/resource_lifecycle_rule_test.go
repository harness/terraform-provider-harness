package lifecycle_test

import (
	"fmt"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/harness/harness-go-sdk/harness/har"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const resourceType = "harness_platform_har_lifecycle_rule"
const resourceName = "harness_platform_har_lifecycle_rule.test"

func testAccLifecycleRuleCheckDestroy(s *terraform.State) error {
	c, ctx := acctest.TestAccGetHarClientWithContext()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourceType {
			continue
		}
		accountId := rs.Primary.Attributes["account_id"]
		ruleId := rs.Primary.ID
		resp, _, err := c.LifecycleApi.GetLifecycleRule(ctx, accountId, "", "", ruleId)
		if err == nil && resp.Id != "" {
			return fmt.Errorf("lifecycle rule still exists: %s", ruleId)
		}
	}
	return nil
}

func importStateIdFunc(sc ruleScope) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		ruleId := rs.Primary.ID
		switch {
		case sc.projectId != "":
			return fmt.Sprintf("%s/%s/%s/%s", sc.accountId, sc.orgId, sc.projectId, ruleId), nil
		case sc.orgId != "":
			return fmt.Sprintf("%s/%s/%s", sc.accountId, sc.orgId, ruleId), nil
		default:
			return fmt.Sprintf("%s/%s", sc.accountId, ruleId), nil
		}
	}
}

func runAtAllScopes(t *testing.T, fn func(t *testing.T, sc ruleScope)) {
	t.Helper()
	for _, tc := range allScopeCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			fn(t, tc.scope)
		})
	}
}

// TestAccLifecycleRuleDeleteAction tests creating a basic DELETE rule at every scope.
func TestAccLifecycleRuleDeleteAction(t *testing.T) {
	runAtAllScopes(t, func(t *testing.T, sc ruleScope) {
		name := fmt.Sprintf("tf_lc_%s", randAlphanumeric(6))
		resource.UnitTest(t, resource.TestCase{
			PreCheck:          func() { acctest.TestAccPreCheck(t) },
			ProviderFactories: acctest.ProviderFactories,
			CheckDestroy:      testAccLifecycleRuleCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccLifecycleRuleBasic(name, sc),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "name", name),
						resource.TestCheckResourceAttr(resourceName, "action", "DELETE"),
						resource.TestCheckResourceAttr(resourceName, "apply_to.0.mode", "ALL_IN_SCOPE"),
						resource.TestCheckResourceAttr(resourceName, "account_id", sc.accountId),
						resource.TestCheckResourceAttrSet(resourceName, "rule_id"),
					),
				},
				{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateIdFunc: importStateIdFunc(sc),
				},
			},
		})
	})
}

// TestAccLifecycleRuleProtectAction tests creating a PROTECT rule at every scope.
func TestAccLifecycleRuleProtectAction(t *testing.T) {
	runAtAllScopes(t, func(t *testing.T, sc ruleScope) {
		name := fmt.Sprintf("tf_lc_%s", randAlphanumeric(6))
		resource.UnitTest(t, resource.TestCase{
			PreCheck:          func() { acctest.TestAccPreCheck(t) },
			ProviderFactories: acctest.ProviderFactories,
			CheckDestroy:      testAccLifecycleRuleCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccLifecycleRuleProtect(name, sc),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "name", name),
						resource.TestCheckResourceAttr(resourceName, "action", "PROTECT"),
						resource.TestCheckResourceAttr(resourceName, "apply_to.0.mode", "ALL_IN_SCOPE"),
						resource.TestCheckResourceAttrSet(resourceName, "rule_id"),
					),
				},
			},
		})
	})
}

// TestAccLifecycleRuleKeepLastN tests KEEP_LAST_N criteria at every scope.
func TestAccLifecycleRuleKeepLastN(t *testing.T) {
	runAtAllScopes(t, func(t *testing.T, sc ruleScope) {
		name := fmt.Sprintf("tf_lc_%s", randAlphanumeric(6))
		resource.UnitTest(t, resource.TestCase{
			PreCheck:          func() { acctest.TestAccPreCheck(t) },
			ProviderFactories: acctest.ProviderFactories,
			CheckDestroy:      testAccLifecycleRuleCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccLifecycleRuleWithKeepLastN(name, sc, 10),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "name", name),
						resource.TestCheckResourceAttr(resourceName, "action", "DELETE"),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.match", "ALL"),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.rules.0.type", "KEEP_LAST_N"),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.rules.0.value", "10"),
						resource.TestCheckResourceAttr(resourceName, "description", "Keep last 10 versions"),
						resource.TestCheckResourceAttrSet(resourceName, "rule_id"),
					),
				},
				{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateIdFunc: importStateIdFunc(sc),
				},
			},
		})
	})
}

// TestAccLifecycleRuleAgeBased tests AGE_BASED criteria at every scope.
func TestAccLifecycleRuleAgeBased(t *testing.T) {
	runAtAllScopes(t, func(t *testing.T, sc ruleScope) {
		name := fmt.Sprintf("tf_lc_%s", randAlphanumeric(6))
		resource.UnitTest(t, resource.TestCase{
			PreCheck:          func() { acctest.TestAccPreCheck(t) },
			ProviderFactories: acctest.ProviderFactories,
			CheckDestroy:      testAccLifecycleRuleCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccLifecycleRuleWithAgeBased(name, sc),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "criteria.0.match", "ALL"),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.rules.0.type", "AGE_BASED"),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.rules.0.value", "30"),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.rules.0.unit", "DAYS"),
						resource.TestCheckResourceAttrSet(resourceName, "rule_id"),
					),
				},
			},
		})
	})
}

// TestAccLifecycleRuleMultipleCriteria tests multiple criteria with ANY match at every scope.
func TestAccLifecycleRuleMultipleCriteria(t *testing.T) {
	runAtAllScopes(t, func(t *testing.T, sc ruleScope) {
		name := fmt.Sprintf("tf_lc_%s", randAlphanumeric(6))
		resource.UnitTest(t, resource.TestCase{
			PreCheck:          func() { acctest.TestAccPreCheck(t) },
			ProviderFactories: acctest.ProviderFactories,
			CheckDestroy:      testAccLifecycleRuleCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccLifecycleRuleWithMultipleCriteria(name, sc),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "criteria.0.match", "ANY"),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.rules.#", "2"),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.rules.0.type", "KEEP_LAST_N"),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.rules.1.type", "AGE_BASED"),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.rules.1.unit", "DAYS"),
					),
				},
			},
		})
	})
}

// TestAccLifecycleRuleWithSchedule tests a cron schedule at every scope.
func TestAccLifecycleRuleWithSchedule(t *testing.T) {
	runAtAllScopes(t, func(t *testing.T, sc ruleScope) {
		name := fmt.Sprintf("tf_lc_%s", randAlphanumeric(6))
		resource.UnitTest(t, resource.TestCase{
			PreCheck:          func() { acctest.TestAccPreCheck(t) },
			ProviderFactories: acctest.ProviderFactories,
			CheckDestroy:      testAccLifecycleRuleCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccLifecycleRuleWithSchedule(name, sc),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "schedule.0.expression", "0 2 * * *"),
						resource.TestCheckResourceAttr(resourceName, "schedule.0.timezone", "UTC"),
						resource.TestCheckResourceAttrSet(resourceName, "rule_id"),
					),
				},
				{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateIdFunc: importStateIdFunc(sc),
				},
			},
		})
	})
}

// TestAccLifecycleRuleDockerFilter tests Docker filter config at every scope.
func TestAccLifecycleRuleDockerFilter(t *testing.T) {
	runAtAllScopes(t, func(t *testing.T, sc ruleScope) {
		name := fmt.Sprintf("tf_lc_%s", randAlphanumeric(6))
		resource.UnitTest(t, resource.TestCase{
			PreCheck:          func() { acctest.TestAccPreCheck(t) },
			ProviderFactories: acctest.ProviderFactories,
			CheckDestroy:      testAccLifecycleRuleCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccLifecycleRuleWithDockerFilter(name, sc),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "package_type", "DOCKER"),
						resource.TestCheckResourceAttr(resourceName, "filter_config.0.package_type", "DOCKER"),
						resource.TestCheckResourceAttr(resourceName, "filter_config.0.package_name_allowed_pattern.#", "2"),
						resource.TestCheckResourceAttr(resourceName, "filter_config.0.tag_name_allowed_pattern.#", "2"),
						resource.TestCheckResourceAttrSet(resourceName, "rule_id"),
					),
				},
			},
		})
	})
}

// TestAccLifecycleRuleUpdate tests in-place update at every scope.
func TestAccLifecycleRuleUpdate(t *testing.T) {
	runAtAllScopes(t, func(t *testing.T, sc ruleScope) {
		name := fmt.Sprintf("tf_lc_%s", randAlphanumeric(6))
		resource.UnitTest(t, resource.TestCase{
			PreCheck:          func() { acctest.TestAccPreCheck(t) },
			ProviderFactories: acctest.ProviderFactories,
			CheckDestroy:      testAccLifecycleRuleCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccLifecycleRuleWithKeepLastN(name, sc, 10),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "name", name),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.rules.0.value", "10"),
					),
				},
				{
					Config: testAccLifecycleRuleUpdated(name, sc, 20),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
						resource.TestCheckResourceAttr(resourceName, "criteria.0.rules.0.value", "20"),
						resource.TestCheckResourceAttr(resourceName, "description", "Updated rule"),
						resource.TestCheckResourceAttr(resourceName, "schedule.0.expression", "0 3 * * 0"),
					),
				},
			},
		})
	})
}

// TestAccLifecycleRuleDryRun creates a rule and triggers a dry-run at every scope.
func TestAccLifecycleRuleDryRun(t *testing.T) {
	runAtAllScopes(t, func(t *testing.T, sc ruleScope) {
		name := fmt.Sprintf("tf_lc_%s", randAlphanumeric(6))
		resource.UnitTest(t, resource.TestCase{
			PreCheck:          func() { acctest.TestAccPreCheck(t) },
			ProviderFactories: acctest.ProviderFactories,
			CheckDestroy:      testAccLifecycleRuleCheckDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccLifecycleRuleWithKeepLastN(name, sc, 5),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet(resourceName, "rule_id"),
						func(s *terraform.State) error {
							rs, ok := s.RootModule().Resources[resourceName]
							if !ok {
								return fmt.Errorf("resource not found: %s", resourceName)
							}
							ruleId := rs.Primary.ID
							c, ctx := acctest.TestAccGetHarClientWithContext()
							exec, _, err := c.LifecycleApi.TriggerLifecycleDryRun(ctx, sc.accountId, sc.orgId, sc.projectId, har.LifecycleDryRunRequest{
								PolicyId: ruleId,
							})
							if err != nil {
								return fmt.Errorf("dry-run trigger failed: %v", err)
							}
							if exec.Id == "" {
								return fmt.Errorf("dry-run returned empty execution ID")
							}
							return nil
						},
					),
				},
			},
		})
	})
}
