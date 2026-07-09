package action_template_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceActionTemplate_basic(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_action_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceActionTemplate_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
					resource.TestCheckResourceAttrSet(resourceName, "revision"),
				),
			},
			{
				// Drift check: re-planning the identical config must be a no-op.
				Config:   testAccResourceActionTemplate_basic(name),
				PlanOnly: true,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					// Strip account_id from the ID for import
					// ID format: account_id/org_id/project_id/hub_identity/identity
					// Import format: org_id/project_id/hub_identity/identity
					resourceState := s.RootModule().Resources[resourceName]
					parts := strings.Split(resourceState.Primary.ID, "/")
					if len(parts) == 5 {
						return strings.Join(parts[1:], "/"), nil
					}
					return resourceState.Primary.ID, nil
				},
			},
		},
	})
}

func TestAccResourceActionTemplate_update(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_chaos_action_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceActionTemplate_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Test action template"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
				),
			},
			{
				Config: testAccResourceActionTemplate_updated(name, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated action template"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "3"),
				),
			},
		},
	})
}

func testAccResourceActionTemplate_basic(name string) string {
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
			description = "Test chaos hub for action template"
		}

		resource "harness_chaos_action_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Test action template"
			type         = "delay"
			tags         = ["test", "terraform"]
			
			delay_action {
				duration = "30s"
			}
			
			depends_on = [harness_chaos_hub_v2.test]
		}
	`, name)
}

func testAccResourceActionTemplate_updated(identity, name string) string {
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
			description = "Test chaos hub for action template"
		}

		resource "harness_chaos_action_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[2]s"
			description  = "Updated action template"
			type         = "delay"
			tags         = ["test", "terraform", "updated"]
			
			delay_action {
				duration = "60s"
			}
		}
	`, identity, name)
}

// TestAccResourceActionTemplate_customScript covers the customScript action type
// (previously only `delay` was integration-tested), including a drift (PlanOnly)
// re-plan and import.
func TestAccResourceActionTemplate_customScript(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_action_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceActionTemplate_customScript(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "type", "customScript"),
					resource.TestCheckResourceAttr(resourceName, "custom_script_action.0.command", "bash"),
					resource.TestCheckResourceAttr(resourceName, "custom_script_action.0.args.#", "2"),
				),
			},
			{
				Config:   testAccResourceActionTemplate_customScript(name),
				PlanOnly: true,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: actionTemplateImportIDFunc(resourceName),
			},
		},
	})
}

// TestAccResourceActionTemplate_container covers the container action type.
func TestAccResourceActionTemplate_container(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_action_template.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceActionTemplate_container(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identity", name),
					resource.TestCheckResourceAttr(resourceName, "type", "container"),
					resource.TestCheckResourceAttr(resourceName, "container_action.0.image", "busybox:latest"),
					resource.TestCheckResourceAttr(resourceName, "container_action.0.command.#", "1"),
				),
			},
			{
				Config:   testAccResourceActionTemplate_container(name),
				PlanOnly: true,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: actionTemplateImportIDFunc(resourceName),
			},
		},
	})
}

// actionTemplateImportIDFunc strips the leading account_id from the canonical
// resource ID (account_id/org_id/project_id/hub_identity/identity) to build the
// org/project-scoped import form (org_id/project_id/hub_identity/identity).
func actionTemplateImportIDFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		parts := strings.Split(rs.Primary.ID, "/")
		if len(parts) == 5 {
			return strings.Join(parts[1:], "/"), nil
		}
		return rs.Primary.ID, nil
	}
}

func testAccResourceActionTemplate_customScript(name string) string {
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
			description = "Test chaos hub for action template"
		}

		resource "harness_chaos_action_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Custom script action template"
			type         = "customScript"
			tags         = ["test", "terraform"]

			custom_script_action {
				command = "bash"
				args    = ["-c", "echo 'running chaos script'"]

				env {
					name  = "TARGET_NAMESPACE"
					value = "default"
				}
			}
		}
	`, name)
}

func testAccResourceActionTemplate_container(name string) string {
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
			description = "Test chaos hub for action template"
		}

		resource "harness_chaos_action_template" "test" {
			org_id       = harness_platform_organization.test.id
			project_id   = harness_platform_project.test.id
			hub_identity = harness_chaos_hub_v2.test.identity
			identity     = "%[1]s"
			name         = "%[1]s"
			description  = "Container action template"
			type         = "container"
			tags         = ["test", "terraform"]

			container_action {
				image     = "busybox:latest"
				command   = ["sh"]
				args      = "echo 'running container action'; sleep 5"
				namespace = "default"
			}
		}
	`, name)
}
