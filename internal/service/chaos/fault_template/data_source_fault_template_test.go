package fault_template_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceFaultTemplate_byIdentity(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "data.harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFaultTemplate_byIdentity(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
					resource.TestCheckResourceAttr(resourceName, "category.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "category.0", "Kubernetes"),
				),
			},
		},
	})
}

func TestAccDataSourceFaultTemplate_byName(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "data.harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFaultTemplate_byName(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
				),
			},
		},
	})
}

func TestAccDataSourceFaultTemplate_withKubernetesSpec(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "data.harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFaultTemplate_withKubernetesSpec(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.chaos.0.kubernetes.0.image", "chaosnative/chaos-go-runner:ci"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
				),
			},
		},
	})
}

func TestAccDataSourceFaultTemplate_withTargets(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "data.harness_chaos_fault_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFaultTemplate_withTargets(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "spec.0.target.0.kubernetes.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
				),
			},
		},
	})
}

// Helper functions to generate test configurations

func testAccDataSourceFaultTemplate_byIdentity(name string) string {
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
						key   = "TOTAL_CHAOS_DURATION"
						value = "60"
					}
				}
			}
		}

		data "harness_chaos_fault_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = harness_chaos_fault_template.test.identity
		}
	`, name)
}

func testAccDataSourceFaultTemplate_byName(name string) string {
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
						key   = "TOTAL_CHAOS_DURATION"
						value = "60"
					}
				}
			}
		}

		data "harness_chaos_fault_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			name         = harness_chaos_fault_template.test.name
		}
	`, name)
}

func testAccDataSourceFaultTemplate_withKubernetesSpec(name string) string {
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
						key   = "TOTAL_CHAOS_DURATION"
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
					}
				}
			}
		}

		data "harness_chaos_fault_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = harness_chaos_fault_template.test.identity
		}
	`, name)
}

func testAccDataSourceFaultTemplate_withTargets(name string) string {
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
						key   = "TOTAL_CHAOS_DURATION"
						value = "60"
					}
				}

				target {
					kubernetes {
						kind      = "deployment"
						namespace = "default"
						names     = "app-deployment"
						labels    = "app=myapp"
					}
				}
			}
		}

		data "harness_chaos_fault_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = harness_chaos_fault_template.test.identity
		}
	`, name)
}
