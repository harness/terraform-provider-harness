package fault_template_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceFaultTemplate_basic(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFaultTemplate_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
					resource.TestCheckResourceAttrSet(resourceName, "revision"),
					resource.TestCheckResourceAttr(resourceName, "category.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "category.0", "Kubernetes"),
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

func TestAccResourceFaultTemplate_update(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFaultTemplate_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Test fault template"),
				),
			},
			{
				Config: testAccResourceFaultTemplate_updated(name, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated fault template"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "3"),
				),
			},
		},
	})
}

func TestAccResourceFaultTemplate_withKubernetesSpec(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFaultTemplate_withKubernetesSpec(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.chaos.0.kubernetes.0.image", "chaosnative/chaos-go-runner:ci"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.chaos.0.kubernetes.0.image_pull_policy", "Always"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.chaos.0.kubernetes.0.command.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.chaos.0.kubernetes.0.args.#", "2"),
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

func TestAccResourceFaultTemplate_withVolumes(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFaultTemplate_withVolumes(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.chaos.0.kubernetes.0.config_maps.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.chaos.0.kubernetes.0.config_maps.0.name", "test-config"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.chaos.0.kubernetes.0.secrets.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.chaos.0.kubernetes.0.secrets.0.secret_name", "test-secret"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.chaos.0.kubernetes.0.host_file_volumes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.chaos.0.kubernetes.0.host_file_volumes.0.name", "test-hostpath"),
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

func TestAccResourceFaultTemplate_withTargets(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFaultTemplate_withTargets(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.target.0.kubernetes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.target.0.kubernetes.0.kind", "deployment"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.target.0.kubernetes.0.namespace", "default"),
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

// Variables are returned by GetFaultTemplate in the `fault` payload (parsed
// from the stored template YAML), so they now survive Read/import.
func TestAccResourceFaultTemplate_withVariables(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFaultTemplate_withVariables(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "variables.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "variables.0.name", "TARGET_NAMESPACE"),
					resource.TestCheckResourceAttr(resourceName, "variables.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "variables.1.name", "CHAOS_DURATION"),
					resource.TestCheckResourceAttr(resourceName, "variables.1.type", "string"),
				),
			},
			{
				// Drift check: variables must round-trip with no perpetual diff.
				Config:   testAccResourceFaultTemplate_withVariables(name),
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

func TestAccResourceFaultTemplate_withLinks(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFaultTemplate_withLinks(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "links.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "links.0.name", "Documentation"),
					resource.TestCheckResourceAttr(resourceName, "links.0.url", "https://docs.example.com"),
					resource.TestCheckResourceAttr(resourceName, "links.1.name", "Source"),
					resource.TestCheckResourceAttr(resourceName, "links.1.url", "https://github.com/example/fault"),
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

// Helper functions to generate test configurations

func testAccResourceFaultTemplate_basic(name string) string {
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
			description = "Test chaos hub for fault template"
		}

		resource "harness_chaos_fault_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test fault template"
			category     = ["Kubernetes"]
			infrastructures = ["KubernetesV2"]
			type         = "Custom"
			tags         = ["test", "terraform"]

			spec {
				chaos {
					fault_name = "byoc-injector"

					params {
						name  = "TOTAL_CHAOS_DURATION"
						value = "60"
					}
				}
			}
		}
	`, name)
}

func testAccResourceFaultTemplate_updated(identity, name string) string {
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
			description = "Test chaos hub for fault template"
		}

		resource "harness_chaos_fault_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[2]s"
			description  = "Updated fault template"
			category     = ["Kubernetes"]
			infrastructures = ["KubernetesV2"]
			type         = "Custom"
			tags         = ["test", "terraform", "updated"]

			spec {
				chaos {
					fault_name = "byoc-injector"

					params {
						name  = "TOTAL_CHAOS_DURATION"
						value = "120"
					}
				}
			}
		}
	`, identity, name)
}

func testAccResourceFaultTemplate_withKubernetesSpec(name string) string {
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
			description = "Test chaos hub for fault template"
		}

		resource "harness_chaos_fault_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test fault template with kubernetes spec"
			category     = ["Kubernetes"]
			infrastructures = ["KubernetesV2"]
			type         = "Custom"
			tags         = ["test", "terraform"]

			spec {
				chaos {
					fault_name = "byoc-injector"

					params {
						name  = "TOTAL_CHAOS_DURATION"
						value = "60"
					}

					kubernetes {
						image             = "chaosnative/chaos-go-runner:ci"
						image_pull_policy = "Always"
						command           = ["/bin/sh"]
						args              = ["-c", "echo 'Running chaos'"]

						labels = {
							"app" = "chaos"
						}

						annotations = {
							"chaos.io/type" = "custom"
						}

						resources {
							limits = {
								"cpu"    = "500m"
								"memory" = "512Mi"
							}
							requests = {
								"cpu"    = "250m"
								"memory" = "256Mi"
							}
						}

						env {
							name  = "CHAOS_NAMESPACE"
							value = "default"
						}
					}
				}
			}
		}
	`, name)
}

func testAccResourceFaultTemplate_withVolumes(name string) string {
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
			description = "Test chaos hub for fault template"
		}

		resource "harness_chaos_fault_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test fault template with volumes"
			category     = ["Kubernetes"]
			infrastructures = ["KubernetesV2"]
			type         = "Custom"
			tags         = ["test", "terraform"]

			spec {
				chaos {
					fault_name = "byoc-injector"

					params {
						name  = "TOTAL_CHAOS_DURATION"
						value = "60"
					}

					kubernetes {
						image             = "chaosnative/chaos-go-runner:ci"
						image_pull_policy = "Always"

						config_maps {
							name       = "test-config"
							mount_path = "/etc/config"
						}

						secrets {
							secret_name = "test-secret"
							mount_path  = "/etc/secret"
						}

						host_file_volumes {
							name       = "test-hostpath"
							mount_path = "/host/data"
							host_path  = "/data"
							type       = "Directory"
						}
					}
				}
			}
		}
	`, name)
}

func testAccResourceFaultTemplate_withTargets(name string) string {
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
			description = "Test chaos hub for fault template"
		}

		resource "harness_chaos_fault_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test fault template with targets"
			category     = ["Kubernetes"]
			infrastructures = ["KubernetesV2"]
			type         = "Custom"
			tags         = ["test", "terraform"]

			spec {
				chaos {
					fault_name = "pod-delete"

					params {
						name  = "TOTAL_CHAOS_DURATION"
						value = "60"
					}
				}

				target {
					kubernetes {
						kind      = "deployment"
						namespace = "default"
						names     = ["app-deployment"]
						labels    = { app = "myapp" }
					}
				}
			}
		}
	`, name)
}

func testAccResourceFaultTemplate_withVariables(name string) string {
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
			description = "Test chaos hub for fault template"
		}

		resource "harness_chaos_fault_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test fault template with variables"
			category     = ["Kubernetes"]
			infrastructures = ["KubernetesV2"]
			type         = "Custom"
			tags         = ["test", "terraform"]

			variables {
				name        = "TARGET_NAMESPACE"
				description = "Namespace to target"
				type        = "string"
				value       = "<+input>"
			}

			variables {
				name        = "CHAOS_DURATION"
				description = "Duration of chaos in seconds"
				type        = "string"
				value       = "<+input>.default('60')"
			}

			spec {
				chaos {
					fault_name = "pod-delete"

					params {
						name  = "TARGET_NAMESPACE"
						value = "<+variables.TARGET_NAMESPACE>"
					}

					params {
						name  = "TOTAL_CHAOS_DURATION"
						value = "<+variables.CHAOS_DURATION>"
					}
				}
			}
		}
	`, name)
}

func testAccResourceFaultTemplate_withLinks(name string) string {
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
			description = "Test chaos hub for fault template"
		}

		resource "harness_chaos_fault_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test fault template with links"
			category     = ["Kubernetes"]
			infrastructures = ["KubernetesV2"]
			type         = "Custom"
			tags         = ["test", "terraform"]

			links {
				name = "Documentation"
				url  = "https://docs.example.com"
			}

			links {
				name = "Source"
				url  = "https://github.com/example/fault"
			}

			spec {
				chaos {
					fault_name = "byoc-injector"

					params {
						name  = "TOTAL_CHAOS_DURATION"
						value = "60"
					}
				}
			}
		}
	`, name)
}

// TestAccResourceFaultTemplate_accountScopeImport reproduces the customer-reported
// import failure for an ACCOUNT-scoped fault template imported via the short
// "hub_identity/identity" (2-part) form. Before the fix the import handler only
// accepted the 4-part org_id/project_id/hub_identity/identity form.
func TestAccResourceFaultTemplate_accountScopeImport(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFaultTemplate_accountScope(name),
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
				ImportStateIdFunc: faultTemplateImportIDFunc(resourceName, 2),
			},
		},
	})
}

// TestAccResourceFaultTemplate_orgScopeImport reproduces the import failure for an
// ORG-scoped fault template imported via the "org_id/hub_identity/identity"
// (3-part) form.
func TestAccResourceFaultTemplate_orgScopeImport(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFaultTemplate_orgScope(name),
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
				ImportStateIdFunc: faultTemplateImportIDFunc(resourceName, 3),
			},
		},
	})
}

// faultTemplateImportIDFunc builds a scoped import ID from resource state:
//   - parts == 2: hub_identity/identity            (account scope)
//   - parts == 3: org_id/hub_identity/identity     (org scope)
func faultTemplateImportIDFunc(resourceName string, parts int) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		hub := rs.Primary.Attributes["hub_identity"]
		identity := rs.Primary.Attributes["identity"]
		switch parts {
		case 3:
			return fmt.Sprintf("%s/%s/%s", rs.Primary.Attributes["org_id"], hub, identity), nil
		default:
			return fmt.Sprintf("%s/%s", hub, identity), nil
		}
	}
}

func testAccResourceFaultTemplate_accountScope(name string) string {
	return fmt.Sprintf(`
		resource "harness_chaos_hub_v2" "test" {
			identity    = "%[1]s"
			name        = "%[1]s"
			description = "Account-scope chaos hub for fault template"
		}

		resource "harness_chaos_fault_template" "test" {
			hub_identity    = harness_chaos_hub_v2.test.identity
			identity        = "%[1]s"
			name            = "%[1]s"
			description     = "Account-scope fault template"
			category        = ["Kubernetes"]
			infrastructures = ["KubernetesV2"]
			type            = "Custom"

			spec {
				chaos {
					fault_name = "byoc-injector"

					params {
						name  = "TOTAL_CHAOS_DURATION"
						value = "60"
					}
				}
			}
		}
	`, name)
}

func testAccResourceFaultTemplate_orgScope(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[1]s"
		}

		resource "harness_chaos_hub_v2" "test" {
			org_id      = harness_platform_organization.test.id
			identity    = "%[1]s"
			name        = "%[1]s"
			description = "Org-scope chaos hub for fault template"
		}

		resource "harness_chaos_fault_template" "test" {
			org_id          = harness_platform_organization.test.id
			hub_identity    = harness_chaos_hub_v2.test.identity
			identity        = "%[1]s"
			name            = "%[1]s"
			description     = "Org-scope fault template"
			category        = ["Kubernetes"]
			infrastructures = ["KubernetesV2"]
			type            = "Custom"

			spec {
				chaos {
					fault_name = "byoc-injector"

					params {
						name  = "TOTAL_CHAOS_DURATION"
						value = "60"
					}
				}
			}
		}
	`, name)
}
