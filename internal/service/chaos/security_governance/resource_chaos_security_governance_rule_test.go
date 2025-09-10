package security_governance_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceChaosSecurityGovernanceRule(t *testing.T) {
	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_security_governance_rule.test"
	conditionResourceName := "harness_chaos_security_governance_condition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecurityGovernanceRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosSecurityGovernanceRuleConfig(rName, id, "Kubernetes"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttrPair(resourceName, "condition_ids.0", conditionResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "time_windows.0.time_zone", "UTC"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceChaosSecurityGovernanceRule_Update(t *testing.T) {
	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_security_governance_rule.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecurityGovernanceRuleDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosSecurityGovernanceRuleConfig(rName, id, "Kubernetes"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				),
			},
			{
				Config: testAccResourceChaosSecurityGovernanceRuleConfigUpdate(rName, id, "Kubernetes"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "Updated test rule"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "updated"),
					resource.TestCheckResourceAttr(resourceName, "user_group_ids.0", "account.Account_Admin"),
				),
			},
		},
	})
}

// Helpers for Destroy & Import State

func testAccSecurityGovernanceRuleDestroy(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Implement actual destroy check if needed
		return nil
	}
}

// Terraform Configurations

func testAccResourceChaosSecurityGovernanceRuleConfig(name, id, infraType string) string {

	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[2]s"
		name       = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		org_id      = harness_platform_organization.test.id
		color       = "#0063F7"
		description = "Test project for Chaos Security Governance"
		tags        = ["foo:bar", "baz:qux"]
	}

	resource "harness_chaos_security_governance_condition" "test" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "%[1]s-condition"
		description = "Test condition"
		infra_type = "%[3]s"
		tags       = ["test"]

		fault_spec {
			operator = "NOT_EQUAL_TO"
			faults {
				fault_type = "FAULT"
				name       = "container-kill"
			}
		}

		k8s_spec {
			infra_spec {
				operator = "EQUAL_TO"
				infra_ids = ["infra1", "infra2"]
			}
			application_spec {
				operator = "EQUAL_TO"
				workloads {
					label = "sdsdsd"
					namespace = "sdsd"
				}
			}
			chaos_service_account_spec {
				operator = "EQUAL_TO"
				service_accounts = ["service_account1", "service_account2"]
			}
		}
	}

	resource "harness_chaos_security_governance_rule" "test" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "%[1]s"
		description = "Test rule"
		is_enabled = true
		tags       = ["test"]

		condition_ids = [harness_chaos_security_governance_condition.test.id]

		user_group_ids = ["account.Account_Admin"]

		time_windows {
			time_zone = "UTC"
			start_time = 1756616940000
			duration = "30m"
			recurrence {
				type = "None"
				until = -1
			}
		}
	}
	`, name, id, infraType)
}

func testAccResourceChaosSecurityGovernanceRuleConfigUpdate(name, id, infraType string) string {
	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[2]s"
		name       = "%[1]s"
	}

	resource "harness_platform_project" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		org_id      = harness_platform_organization.test.id
		color       = "#0063F7"
		description = "Test project for Chaos Security Governance"
		tags        = ["foo:bar", "baz:qux"]
	}

	resource "harness_chaos_security_governance_condition" "test" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "%[1]s-condition"
		description = "Test condition"
		infra_type = "%[3]s"
		tags       = ["test"]

		fault_spec {
			operator = "NOT_EQUAL_TO"
			faults {
				fault_type = "FAULT"
				name       = "container-kill"
			}
		}

		k8s_spec {
			infra_spec {
				operator = "EQUAL_TO"
				infra_ids = ["infra1", "infra2"]
			}
			application_spec {
				operator = "EQUAL_TO"
				workloads {
					label = "sdsdsd"
					namespace = "sdsd"
				}
			}
			chaos_service_account_spec {
				operator = "EQUAL_TO"
				service_accounts = ["service_account1", "service_account2"]
			}
		}
	}

	resource "harness_chaos_security_governance_rule" "test" {
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
		name       = "%[1]s"
		description = "Updated test rule"
		is_enabled = true
		tags       = ["updated"]
		user_group_ids = ["account.Account_Admin"]

		condition_ids = [harness_chaos_security_governance_condition.test.id]

		time_windows {
			time_zone = "UTC"
			start_time = 1756616940000
			end_time = 1756616940020
			recurrence {
				type = "None"
				until = -1
			}
		}
	}
	`, name, id, infraType)
}
