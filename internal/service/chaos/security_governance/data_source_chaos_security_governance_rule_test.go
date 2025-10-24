package security_governance_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccDataSourceChaosSecurityGovernanceRule verifies the basic data source functionality for Chaos Security Governance Rule.
func TestAccDataSourceChaosSecurityGovernanceRule(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	dataSourceName := "data.harness_chaos_security_governance_rule.test"
	resourceName := "harness_chaos_security_governance_rule.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosSecurityGovernanceRuleConfig(rName, id, "Kubernetes"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "is_enabled", resourceName, "is_enabled"),
					resource.TestCheckResourceAttrPair(dataSourceName, "tags", resourceName, "tags"),
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
				),
			},
		},
	})
}

// TestAccDataSourceChaosSecurityGovernanceRule_ByName verifies the data source lookup by name for Chaos Security Governance Rule.
func TestAccDataSourceChaosSecurityGovernanceRule_ByName(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	dataSourceName := "data.harness_chaos_security_governance_rule.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosSecurityGovernanceRuleConfigByName(rName, id, "Kubernetes"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", rName),
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
				),
			},
		},
	})
}

// Terraform Configurations

func testAccDataSourceChaosSecurityGovernanceRuleConfig(name, id, infraType string) string {

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
			operator = "EQUAL_TO"
			faults {
				fault_type = "FAULT"
				name       = "pod-delete"
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

	data "harness_chaos_security_governance_rule" "test" {
		id         = harness_chaos_security_governance_rule.test.id
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
	`, name, id, infraType)
}

func testAccDataSourceChaosSecurityGovernanceRuleConfigByName(name, id, infraType string) string {

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
			operator = "EQUAL_TO"
			faults {
				fault_type = "FAULT"
				name       = "pod-delete"
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

	data "harness_chaos_security_governance_rule" "test" {
		id         = harness_chaos_security_governance_rule.test.id
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
	`, name, id, infraType)
}
