package probe_template_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProbeTemplate_httpProbe_basic(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_probe_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProbeTemplate_httpProbe_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "type", "httpProbe"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
					resource.TestCheckResourceAttrSet(resourceName, "revision"),
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

func TestAccResourceProbeTemplate_httpProbe_update(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_chaos_probe_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProbeTemplate_httpProbe_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Test HTTP probe template"),
				),
			},
			{
				Config: testAccResourceProbeTemplate_httpProbe_updated(name, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated HTTP probe template"),
				),
			},
		},
	})
}

func TestAccResourceProbeTemplate_cmdProbe_basic(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_probe_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProbeTemplate_cmdProbe_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "type", "cmdProbe"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
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

func TestAccResourceProbeTemplate_k8sProbe_basic(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_probe_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProbeTemplate_k8sProbe_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "type", "k8sProbe"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
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

func TestAccResourceProbeTemplate_apmProbe_prometheus(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_probe_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProbeTemplate_apmProbe_prometheus(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "type", "apmProbe"),
					resource.TestCheckResourceAttr(resourceName, "apm_probe.0.apm_type", "Prometheus"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
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

func TestAccResourceProbeTemplate_apmProbe_datadog(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_probe_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProbeTemplate_apmProbe_datadog(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "type", "apmProbe"),
					resource.TestCheckResourceAttr(resourceName, "apm_probe.0.apm_type", "Datadog"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
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

func testAccResourceProbeTemplate_httpProbe_basic(name string) string {
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
			description = "Test chaos hub for probe template"
		}

		resource "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test HTTP probe template"
			type         = "httpProbe"
			tags         = ["test", "terraform"]

			http_probe {
				url = "https://example.com/health"
				method {
					get {
						criteria   = "=="
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
	`, name)
}

func testAccResourceProbeTemplate_httpProbe_updated(identity, name string) string {
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
			description = "Test chaos hub for probe template"
		}

		resource "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[2]s"
			description  = "Updated HTTP probe template"
			type         = "httpProbe"
			tags         = ["test", "terraform", "updated"]

			http_probe {
				url = "https://example.com/health"
				method {
					get {
						criteria   = "=="
						response_code = "200"
					}
				}
			}

			run_properties {
				timeout          = "15s"
				interval         = "10s"
				polling_interval = "2s"
				stop_on_failure  = true
			}
		}
	`, identity, name)
}

func testAccResourceProbeTemplate_cmdProbe_basic(name string) string {
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
			description = "Test chaos hub for probe template"
		}

		resource "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test CMD probe template"
			type         = "cmdProbe"
			tags         = ["test", "terraform"]

			cmd_probe {
				command = "echo 'test'"
				source {
					inline {
						command = "echo 'test'"
					}
				}
			}

			run_properties {
				timeout  = "10s"
				interval = "5s"
			}
		}
	`, name)
}

func testAccResourceProbeTemplate_k8sProbe_basic(name string) string {
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
			description = "Test chaos hub for probe template"
		}

		resource "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test K8s probe template"
			type         = "k8sProbe"
			tags         = ["test", "terraform"]

			k8s_probe {
				version   = "v1"
				resource  = "pods"
				namespace = "default"
				operation = "present"
			}

			run_properties {
				timeout  = "10s"
				interval = "5s"
			}
		}
	`, name)
}

func testAccResourceProbeTemplate_apmProbe_prometheus(name string) string {
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
			description = "Test chaos hub for probe template"
		}

		resource "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test APM Prometheus probe template"
			type         = "apmProbe"
			tags         = ["test", "terraform", "apm", "prometheus"]

			apm_probe {
				apm_type = "Prometheus"

				comparator {
					type     = "float"
					criteria = "<="
					value    = "100"
				}

				prometheus_inputs {
					connector_id = "<+input>"
					query        = "up{job='api-server'}"
				}
			}

			run_properties {
				timeout          = "30s"
				interval         = "10s"
				polling_interval = "2s"
				stop_on_failure  = false
			}
		}
	`, name)
}

func testAccResourceProbeTemplate_apmProbe_datadog(name string) string {
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
			description = "Test chaos hub for probe template"
		}

		resource "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test APM Datadog probe template"
			type         = "apmProbe"
			tags         = ["test", "terraform", "apm", "datadog"]

			apm_probe {
				apm_type = "Datadog"

				comparator {
					type     = "float"
					criteria = "<="
					value    = "95"
				}

				datadog_inputs {
					connector_id = "<+input>"
					query        = "avg:system.cpu.user{*}"
				}
			}

			run_properties {
				timeout          = "30s"
				interval         = "10s"
				polling_interval = "2s"
				stop_on_failure  = true
			}
		}
	`, name)
}
