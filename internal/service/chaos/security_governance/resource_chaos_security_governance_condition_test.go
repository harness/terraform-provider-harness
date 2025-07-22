package security_governance_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceChaosSecurityGovernanceCondition(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_security_governance_condition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecurityGovernanceConditionDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosSecurityGovernanceConditionConfig(rName, id, "Kubernetes"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "infra_type", "Kubernetes"),
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.operator", "EQUAL_TO"),
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.faults.0", "pod-delete"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceChaosSecurityGovernanceCondition_Update(t *testing.T) {
	// Check for required environment variables
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")
	if accountId == "" {
		t.Skip("Skipping test because HARNESS_ACCOUNT_ID is not set")
	}

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_security_governance_condition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccSecurityGovernanceConditionDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosSecurityGovernanceConditionConfig(rName, id, "Kubernetes"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.faults.0", "pod-delete"),
				),
			},
			{
				Config: testAccResourceChaosSecurityGovernanceConditionConfigUpdate(rName, id, "Kubernetes"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.faults.0", "container-kill"),
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.operator", "NOT_EQUAL_TO"),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated test condition"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "updated"),
				),
			},
		},
	})
}

func testAccSecurityGovernanceConditionDestroy(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Implement actual destroy check if needed
		return nil
	}
}

func testAccResourceChaosSecurityGovernanceConditionConfig(name, id, infraType string) string {
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
		name       = "%[1]s"
		description = "Test condition"
		infra_type = "%[3]s"
		tags       = ["test"]

		fault_spec {
			operator = "EQUAL_TO"
			faults   = ["pod-delete"]
		}
	}
	`, name, id, infraType, accountId)
}

func testAccResourceChaosSecurityGovernanceConditionConfigUpdate(name, id, infraType string) string {
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
		name       = "%[1]s"
		description = "Updated test condition"
		infra_type = "%[3]s"
		tags       = ["updated"]

		fault_spec {
			operator = "NOT_EQUAL_TO"
			faults   = ["container-kill"]
		}
	}
	`, name, id, infraType, accountId)
}
