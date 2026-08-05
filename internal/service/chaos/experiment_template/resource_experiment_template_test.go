package experiment_template_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceExperimentTemplate_basic(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_experiment_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceExperimentTemplate_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.infra_type", "KubernetesV2"),
				),
			},
			{
				// Drift check: re-planning the identical config must be a no-op.
				// This guards against the perpetual "update in-place" diff class.
				Config:   testAccResourceExperimentTemplate_basic(name),
				PlanOnly: true,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceExperimentTemplate_update(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_experiment_template.test"
	updatedName := name + "_updated"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceExperimentTemplate_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.cleanup_policy", "delete"),
				),
			},
			{
				Config: testAccResourceExperimentTemplate_updated(name, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "spec.0.cleanup_policy", "retain"),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
				),
			},
		},
	})
}

func TestAccResourceExperimentTemplate_cleanupPolicyRetain(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_experiment_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceExperimentTemplate_cleanupPolicyRetain(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.cleanup_policy", "retain"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.status_check_timeouts.0.delay", "10"),
				),
			},
			{
				Config:   testAccResourceExperimentTemplate_cleanupPolicyRetain(name),
				PlanOnly: true,
			},
		},
	})
}

func testAccResourceExperimentTemplate_basic(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_chaos_hub_v2" "test" {
			org_id      = harness_platform_organization.test.id
			project_id  = harness_platform_project.test.id
			identity    = "%[1]s"
			name        = "%[1]s"
			description = "Test chaos hub for experiment template"
		}

		resource "harness_chaos_experiment_template" "test" {
			identity     = "%[1]s"
			name         = "%[1]s"
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			description  = "Test experiment template"

			spec {
				infra_type = "KubernetesV2"
				infra_id   = "<+input>"

				cleanup_policy = "delete"

				status_check_timeouts {
					delay   = 5
					timeout = 180
				}
			}

			tags = ["test", "experiment"]
		}
	`, name)
}

func testAccResourceExperimentTemplate_updated(identity, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_chaos_hub_v2" "test" {
			org_id      = harness_platform_organization.test.id
			project_id  = harness_platform_project.test.id
			identity    = "%[1]s"
			name        = "%[1]s"
			description = "Test chaos hub for experiment template"
		}

		resource "harness_chaos_experiment_template" "test" {
			identity     = "%[1]s"
			name         = "%[2]s"
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			description  = "Updated description"

			spec {
				infra_type = "KubernetesV2"
				infra_id   = "<+input>"

				cleanup_policy = "retain"

				status_check_timeouts {
					delay   = 10
					timeout = 300
				}
			}

			tags = ["test", "experiment", "updated"]
		}
	`, identity, name)
}

func testAccResourceExperimentTemplate_cleanupPolicyRetain(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_chaos_hub_v2" "test" {
			org_id      = harness_platform_organization.test.id
			project_id  = harness_platform_project.test.id
			identity    = "%[1]s"
			name        = "%[1]s"
			description = "Test chaos hub"
		}

		resource "harness_chaos_experiment_template" "test" {
			identity     = "%[1]s"
			name         = "%[1]s"
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			description  = "Template with retain cleanup policy"

			spec {
				infra_type = "KubernetesV2"
				infra_id   = "<+input>"

				cleanup_policy = "retain"

				status_check_timeouts {
					delay   = 10
					timeout = 300
				}
			}

			tags = ["test", "cleanup-retain"]
		}
	`, name)
}

// TestAccResourceExperimentTemplate_nestedNoDrift integration-tests the
// perpetual-diff fixes: it creates a probe (referencing a probe template in the
// same hub) with enable_data_collection=true, then re-plans the identical
// config (PlanOnly) to assert the plan is empty. Before the fix, the omitted
// enableDataCollection field produced a permanent "update in-place" diff.
func TestAccResourceExperimentTemplate_nestedNoDrift(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_experiment_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceExperimentTemplate_nested(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.probes.0.name", "latency-check"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.probes.0.enable_data_collection", "true"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.probes.0.conditions_v2.0.operator", "AND"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.probes.0.conditions_v2.0.values.#", "2"),
				),
			},
			{
				// Drift check: identical config must produce an empty plan.
				Config:   testAccResourceExperimentTemplate_nested(name),
				PlanOnly: true,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccResourceExperimentTemplate_accountScopeImportShortForm reproduces the
// customer-reported import failure: an account-scoped experiment template
// imported via the short "hub_identity/identity" form. Before the fix the
// import handler rejected anything other than 4 slash-separated segments.
func TestAccResourceExperimentTemplate_accountScopeImportShortForm(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_experiment_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceExperimentTemplate_accountScope(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "org_id", ""),
					resource.TestCheckResourceAttr(resourceName, "project_id", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: experimentTemplateShortFormImportIdFunc(resourceName),
			},
		},
	})
}

// experimentTemplateShortFormImportIdFunc builds the account-scope short-form
// import ID (hub_identity/identity) from resource state.
func experimentTemplateShortFormImportIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s",
			rs.Primary.Attributes["hub_identity"],
			rs.Primary.Attributes["identity"]), nil
	}
}

func testAccResourceExperimentTemplate_nested(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_chaos_hub_v2" "test" {
			org_id      = harness_platform_organization.test.id
			project_id  = harness_platform_project.test.id
			identity    = "%[1]s"
			name        = "%[1]s"
			description = "Test chaos hub for nested experiment template"
		}

		resource "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Probe for nested experiment template"
			type         = "httpProbe"

			http_probe {
				url = "https://example.com/health"
				method {
					get {
						criteria      = "=="
						response_code = "200"
					}
				}
			}

			run_properties {
				timeout          = "10s"
				interval         = "5s"
				polling_interval = "1s"
				stop_on_failure  = false
			}
		}

		resource "harness_chaos_experiment_template" "test" {
			identity     = "%[1]s"
			name         = "%[1]s"
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			description  = "Nested experiment template with probe"

			spec {
				infra_type     = "KubernetesV2"
				infra_id       = "<+input>"
				cleanup_policy = "delete"

				probes {
					identity               = harness_chaos_probe_template.test.identity
					name                   = "latency-check"
					is_enterprise          = false
					enable_data_collection = true

					conditions_v2 {
						operator = "AND"
						values   = ["true", "<+input>"]
					}
				}

				status_check_timeouts {
					delay   = 5
					timeout = 180
				}
			}

			tags = ["nested", "drift"]
		}
	`, name)
}

func testAccResourceExperimentTemplate_accountScope(name string) string {
	return fmt.Sprintf(`
		resource "harness_chaos_hub_v2" "test" {
			identity    = "%[1]s"
			name        = "%[1]s"
			description = "Account-scope chaos hub"
		}

		resource "harness_chaos_experiment_template" "test" {
			identity     = "%[1]s"
			name         = "%[1]s"
			hub_identity = harness_chaos_hub_v2.test.identity
			description  = "Account-scope experiment template"

			spec {
				infra_type     = "KubernetesV2"
				infra_id       = "<+input>"
				cleanup_policy = "delete"
			}
		}
	`, name)
}

// TestAccResourceExperimentTemplate_orgScopeImport exercises the org-scoped
// import path via the "org_id/hub_identity/identity" (3-part) form, which
// parseExperimentTemplateImportID supports (case 3) but was not integration-tested.
func TestAccResourceExperimentTemplate_orgScopeImport(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_experiment_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceExperimentTemplate_orgScope(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttrSet(resourceName, "org_id"),
					resource.TestCheckResourceAttr(resourceName, "project_id", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: experimentTemplateOrgScopeImportIdFunc(resourceName),
			},
		},
	})
}

// experimentTemplateOrgScopeImportIdFunc builds the org-scope import ID
// (org_id/hub_identity/identity) from resource state.
func experimentTemplateOrgScopeImportIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		return fmt.Sprintf("%s/%s/%s",
			rs.Primary.Attributes["org_id"],
			rs.Primary.Attributes["hub_identity"],
			rs.Primary.Attributes["identity"]), nil
	}
}

func testAccResourceExperimentTemplate_orgScope(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_chaos_hub_v2" "test" {
			org_id      = harness_platform_organization.test.id
			identity    = "%[1]s"
			name        = "%[1]s"
			description = "Org-scope chaos hub"
		}

		resource "harness_chaos_experiment_template" "test" {
			org_id       = harness_platform_organization.test.id
			identity     = "%[1]s"
			name         = "%[1]s"
			hub_identity = harness_chaos_hub_v2.test.identity
			description  = "Org-scope experiment template"

			spec {
				infra_type     = "KubernetesV2"
				infra_id       = "<+input>"
				cleanup_policy = "delete"
			}
		}
	`, name)
}

// TestAccResourceExperimentTemplate_conditionsV2FaultAction proves conditions_v2
// round-trips (no drift) on FAULTS and ACTIONS, not just probes. It references a
// fault template and an action template in the same hub, attaches conditions_v2
// (AND / OR + values incl. "<+input>") to each, and re-plans (PlanOnly) to assert
// an empty plan.
func TestAccResourceExperimentTemplate_conditionsV2FaultAction(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_experiment_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceExperimentTemplate_conditionsV2FaultAction(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.faults.0.conditions_v2.0.operator", "AND"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.faults.0.conditions_v2.0.values.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.actions.0.conditions_v2.0.operator", "OR"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.actions.0.conditions_v2.0.values.#", "2"),
				),
			},
			{
				// Drift check: conditions_v2 on fault + action must round-trip.
				Config:   testAccResourceExperimentTemplate_conditionsV2FaultAction(name),
				PlanOnly: true,
			},
		},
	})
}

func testAccResourceExperimentTemplate_conditionsV2FaultAction(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_chaos_hub_v2" "test" {
			org_id      = harness_platform_organization.test.id
			project_id  = harness_platform_project.test.id
			identity    = "%[1]s"
			name        = "%[1]s"
			description = "Hub for conditions_v2 fault/action test"
		}

		resource "harness_chaos_fault_template" "test" {
			org_id          = harness_platform_organization.test.id
			project_id      = harness_platform_project.test.id
			hub_identity    = harness_chaos_hub_v2.test.identity
			identity        = "%[1]s-fault"
			name            = "%[1]s-fault"
			category        = ["Kubernetes"]
			infrastructures = ["KubernetesV2"]
			type            = "Custom"

			spec {
				chaos {
					fault_name = "byoc-injector"
					params {
						name  = "TOTAL_CHAOS_DURATION"
						value = "30"
					}
				}
			}
		}

		resource "harness_chaos_action_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s-action"
			name         = "%[1]s-action"
			type         = "delay"

			delay_action {
				duration = "30s"
			}
		}

		resource "harness_chaos_experiment_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "conditions_v2 on fault + action"

			spec {
				infra_type     = "KubernetesV2"
				infra_id       = "<+input>"
				cleanup_policy = "delete"

				faults {
					identity      = harness_chaos_fault_template.test.identity
					name          = "%[1]s-fault"
					revision      = "v1"
					is_enterprise = false
					auth_enabled  = false

					conditions_v2 {
						operator = "AND"
						values   = ["true"]
					}
				}

				actions {
					identity               = harness_chaos_action_template.test.identity
					name                   = "%[1]s-action"
					is_enterprise          = false
					continue_on_completion = true

					conditions_v2 {
						operator = "OR"
						values   = ["true", "<+input>"]
					}
				}

				status_check_timeouts {
					delay   = 5
					timeout = 180
				}
			}
		}
	`, name)
}
