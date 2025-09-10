package security_governance_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceChaosSecurityGovernanceCondition(t *testing.T) {

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
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.faults.0.fault_type", "FAULT"),
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.faults.0.name", "pod-delete"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceChaosSecurityGovernanceCondition_Update(t *testing.T) {

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
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.faults.0.fault_type", "FAULT"),
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.faults.0.name", "pod-delete"),
				),
			},
			{
				Config: testAccResourceChaosSecurityGovernanceConditionConfigUpdate(rName, id, "Kubernetes"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.faults.0.fault_type", "FAULT"),
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.faults.0.name", "container-kill"),
					resource.TestCheckResourceAttr(resourceName, "fault_spec.0.operator", "NOT_EQUAL_TO"),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated test condition"),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "updated"),
				),
			},
		},
	})
}

// Helpers for Destroy & Import State

func testAccSecurityGovernanceConditionDestroy(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Implement actual destroy check if needed
		return nil
	}
}

// Terraform Configurations

func testAccResourceChaosSecurityGovernanceConditionConfig(name, id, infraType string) string {

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
		name       = "%[1]s"
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
	`, name, id, infraType)
}

func testAccResourceChaosSecurityGovernanceConditionConfigUpdate(name, id, infraType string) string {

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
		name       = "%[1]s"
		description = "Updated test condition"
		infra_type = "%[3]s"
		tags       = ["updated"]

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
	`, name, id, infraType)
}
