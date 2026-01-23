package probe_template_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceProbeTemplate_byIdentity(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "data.harness_chaos_probe_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceProbeTemplate_byIdentity(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "type", "httpProbe"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
				),
			},
		},
	})
}

func TestAccDataSourceProbeTemplate_byName(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "data.harness_chaos_probe_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceProbeTemplate_byName(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "type", "httpProbe"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
				),
			},
		},
	})
}

func TestAccDataSourceProbeTemplate_cmdProbe(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "data.harness_chaos_probe_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceProbeTemplate_cmdProbe(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "type", "cmdProbe"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
				),
			},
		},
	})
}

func TestAccDataSourceProbeTemplate_k8sProbe(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "data.harness_chaos_probe_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceProbeTemplate_k8sProbe(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "type", "k8sProbe"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
				),
			},
		},
	})
}

func testAccDataSourceProbeTemplate_byIdentity(name string) string {
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

		resource "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test HTTP probe template"
			type         = "httpProbe"

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
				timeout  = "10s"
				interval = "5s"
			}
		}

		data "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = harness_chaos_probe_template.test.identity
		}
	`, name)
}

func testAccDataSourceProbeTemplate_byName(name string) string {
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

		resource "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test HTTP probe template"
			type         = "httpProbe"

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
				timeout  = "10s"
				interval = "5s"
			}
		}

		data "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			name         = harness_chaos_probe_template.test.name
		}
	`, name)
}

func testAccDataSourceProbeTemplate_cmdProbe(name string) string {
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

		resource "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test CMD probe template"
			type         = "cmdProbe"

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

		data "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = harness_chaos_probe_template.test.identity
		}
	`, name)
}

func testAccDataSourceProbeTemplate_k8sProbe(name string) string {
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

		resource "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test K8s probe template"
			type         = "k8sProbe"

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

		data "harness_chaos_probe_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = harness_chaos_probe_template.test.identity
		}
	`, name)
}
