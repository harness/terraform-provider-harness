package security_governance_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceChaosSecurityGovernanceRule(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
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

func TestAccDataSourceChaosSecurityGovernanceRule_ByName(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
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

func testAccDataSourceChaosSecurityGovernanceRuleConfig(name, id, infraType string) string {
	// Use the account ID from environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		accountId = "test" // Default for test cases when not set
	}

	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[2]s"
		name       = "%[1]s"
		account_id = "%[4]s"
	}

	resource "harness_platform_project" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		org_id      = harness_platform_organization.test.id
		account_id  = "%[4]s"
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
			faults   = ["pod-delete"]
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

		time_windows {
			time_zone = "UTC"
			start_time = "00:00"
			end_time   = "23:59"
			days       = ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY"]
		}
	}

	data "harness_chaos_security_governance_rule" "test" {
		id         = harness_chaos_security_governance_rule.test.id
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
	`, name, id, infraType, accountId)
}

func testAccDataSourceChaosSecurityGovernanceRuleConfigByName(name, id, infraType string) string {
	// Use the account ID from environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		accountId = "test" // Default for test cases when not set
	}

	return fmt.Sprintf(`
	resource "harness_platform_organization" "test" {
		identifier = "%[2]s"
		name       = "%[1]s"
		account_id = "%[4]s"
	}

	resource "harness_platform_project" "test" {
		identifier  = "%[2]s"
		name        = "%[1]s"
		org_id      = harness_platform_organization.test.id
		account_id  = "%[4]s"
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
			faults   = ["pod-delete"]
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

		time_windows {
			time_zone = "UTC"
			start_time = "00:00"
			end_time   = "23:59"
			days       = ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY"]
		}
	}

	data "harness_chaos_security_governance_rule" "test" {
		name       = harness_chaos_security_governance_rule.test.name
		org_id     = harness_platform_organization.test.id
		project_id = harness_platform_project.test.id
	}
	`, name, id, infraType, accountId)
}
