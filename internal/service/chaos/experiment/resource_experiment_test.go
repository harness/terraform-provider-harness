package experiment_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceChaosExperiment_basic(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_chaos_experiment.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosExperiment_basic(name),
				Check: resource.ComposeTestCheckFunc(
					// Basic fields
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Test experiment from template"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),

					// Computed IDs
					resource.TestCheckResourceAttrSet(resourceName, "experiment_id"),
					resource.TestCheckResourceAttrSet(resourceName, "identity"),
					resource.TestCheckResourceAttrSet(resourceName, "infra_id"),

					// CHAOS-SPECIFIC: Import type
					resource.TestCheckResourceAttr(resourceName, "import_type", "REFERENCE"),

					// CHAOS-SPECIFIC: Hub identity (project-level, no prefix)
					resource.TestCheckResourceAttr(resourceName, "hub_identity", name),

					// CHAOS-SPECIFIC: Template identity
					resource.TestCheckResourceAttr(resourceName, "template_identity", name),

					// CHAOS-SPECIFIC: Infrastructure type
					resource.TestCheckResourceAttr(resourceName, "infra_type", "KubernetesV2"),

					// CRITICAL: Template details populated
					resource.TestCheckResourceAttr(resourceName, "template_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "template_details.0.identity"),
					resource.TestCheckResourceAttrSet(resourceName, "template_details.0.hub_reference"),
					resource.TestCheckResourceAttrSet(resourceName, "template_details.0.reference"),
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

func testAccResourceChaosExperiment_basic(name string) string {
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
	org_id       = harness_platform_organization.test.id
	project_id   = harness_platform_project.test.id
	hub_identity = harness_chaos_hub_v2.test.identity
	identity     = "%[1]s"
	name         = "%[1]s"
	description  = "Test experiment template"

	spec {
		infra_type = "KubernetesV2"


		cleanup_policy = "delete"
	}
}

resource "harness_platform_environment" "test" {
	identifier = "%[1]s"
	name       = "%[1]s"
	org_id     = harness_platform_organization.test.id
	project_id = harness_platform_project.test.id
	type       = "PreProduction"
}

resource "harness_platform_connector_kubernetes" "test" {
	identifier = "%[1]s"
	name       = "%[1]s"
	org_id     = harness_platform_organization.test.id
	project_id = harness_platform_project.test.id

	inherit_from_delegate {
		delegate_selectors = ["kubernetes-delegate"]
	}
}

resource "harness_platform_infrastructure" "test" {
	identifier      = "%[1]s"
	name            = "%[1]s"
	org_id          = harness_platform_organization.test.id
	project_id      = harness_platform_project.test.id
	env_id          = harness_platform_environment.test.id
	deployment_type = "Kubernetes"
	type            = "KubernetesDirect"

	yaml = <<-EOT
infrastructureDefinition:
  name: %[1]s
  identifier: %[1]s
  orgIdentifier: ${harness_platform_organization.test.id}
  projectIdentifier: ${harness_platform_project.test.id}
  environmentRef: ${harness_platform_environment.test.id}
  type: KubernetesDirect
  deploymentType: Kubernetes
  spec:
    connectorRef: ${harness_platform_connector_kubernetes.test.id}
    namespace: "chaos-test"
    releaseName: "release-%[1]s"
	EOT
}

resource "harness_chaos_infrastructure_v2" "test" {
	org_id         = harness_platform_organization.test.id
	project_id     = harness_platform_project.test.id
	environment_id = harness_platform_environment.test.id
	name           = "%[1]s"
	infra_id       = harness_platform_infrastructure.test.id
	description    = "Test infrastructure"
	infra_type     = "KUBERNETESV2"
	infra_scope    = "NAMESPACE"
	namespace      = "chaos-test"
	service_account = "litmus"
}

resource "harness_chaos_experiment" "test" {
	org_id            = harness_platform_organization.test.id
	project_id        = harness_platform_project.test.id
	template_identity = harness_chaos_experiment_template.test.identity
	
	# Hub scope (where template lives) - project-level
	hub_org_id        = harness_platform_organization.test.id
	hub_project_id    = harness_platform_project.test.id
	hub_identity      = harness_chaos_hub_v2.test.identity
	
	name              = "%[1]s"
	# Correct format: environment_id/infra_id
	infra_ref         = "${harness_platform_environment.test.id}/${harness_chaos_infrastructure_v2.test.infra_id}"
	description       = "Test experiment from template"
	import_type       = "REFERENCE"

	tags = ["test", "terraform"]
}
`, name)
}

func TestAccResourceChaosExperiment_localImport(t *testing.T) {
	// Use shorter name to stay under 47 char identity limit
	name := fmt.Sprintf("local_%s", utils.RandStringBytes(5))
	resourceName := "harness_chaos_experiment.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosExperiment_localImport(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),

					// CHAOS-SPECIFIC: LOCAL import type
					resource.TestCheckResourceAttr(resourceName, "import_type", "LOCAL"),

					// CRITICAL: Manifest should be populated for LOCAL imports
					resource.TestCheckResourceAttrSet(resourceName, "manifest"),

					// CRITICAL: fault_ids still set even for LOCAL imports
					resource.TestCheckResourceAttrSet(resourceName, "fault_ids.#"),

					// NOTE: template_details is empty for LOCAL imports since content is copied
					// Only REFERENCE imports maintain template references
				),
			},
		},
	})
}

// TestAccResourceChaosExperiment_hubScopeFields tests that hub_org_id and hub_project_id
// are correctly preserved from config and that hub_identity is properly read from API.
// This is a regression test to ensure hub scope fields work correctly.
func TestAccResourceChaosExperiment_hubScopeFields(t *testing.T) {
	// Use shorter name to stay under 47 char identity limit
	name := fmt.Sprintf("hub_scope_%s", utils.RandStringBytes(5))
	resourceName := "harness_chaos_experiment.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosExperiment_basic(name),
				Check: resource.ComposeTestCheckFunc(
					// Basic fields
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "experiment_id"),

					// REGRESSION TEST: Hub scope fields from config should be preserved
					// These are input fields provided by user, not derived
					resource.TestCheckResourceAttrSet(resourceName, "hub_org_id"),
					resource.TestCheckResourceAttrSet(resourceName, "hub_project_id"),

					// Hub identity is read from API's hub_reference field
					resource.TestCheckResourceAttrSet(resourceName, "hub_identity"),

					// Template details should have hub_reference from API
					resource.TestCheckResourceAttr(resourceName, "template_details.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "template_details.0.hub_reference"),
				),
			},
			{
				// Import test - hub_org_id and hub_project_id are derived from hub_reference prefix
				// This provides better UX - users get complete state after import
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceChaosExperiment_localImport(name string) string {
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
	org_id       = harness_platform_organization.test.id
	project_id   = harness_platform_project.test.id
	hub_identity = harness_chaos_hub_v2.test.identity
	identity     = "%[1]s"
	name         = "%[1]s"
	description  = "Test experiment template with fault"

	spec {
		infra_type = "KubernetesV2"
		
		# Add a fault with full definition for LOCAL import
		faults {
			identity      = "pod-delete"
			name          = "pod-delete-test"
			revision      = "v1"
			is_enterprise = true
			auth_enabled  = false
			
			# Add runtime input variables
			values {
				name  = "TARGET_WORKLOAD_KIND"
				value = "<+input>"
			}
			values {
				name  = "TARGET_WORKLOAD_NAMESPACE"
				value = "<+input>"
			}
		}

		# Add vertices to define workflow execution order
		vertices {
			name = "v-start"
			start {
				faults {
					name = "pod-delete-test"
				}
			}
		}
		
		vertices {
			name = "v-end"
			end {
				faults {
					name = "pod-delete-test"
				}
			}
		}

		cleanup_policy = "delete"
	}
}

resource "harness_platform_environment" "test" {
	identifier = "%[1]s"
	name       = "%[1]s"
	org_id     = harness_platform_organization.test.id
	project_id = harness_platform_project.test.id
	type       = "PreProduction"
}

resource "harness_platform_connector_kubernetes" "test" {
	identifier = "%[1]s"
	name       = "%[1]s"
	org_id     = harness_platform_organization.test.id
	project_id = harness_platform_project.test.id

	inherit_from_delegate {
		delegate_selectors = ["kubernetes-delegate"]
	}
}

resource "harness_platform_infrastructure" "test" {
	identifier      = "%[1]s"
	name            = "%[1]s"
	org_id          = harness_platform_organization.test.id
	project_id      = harness_platform_project.test.id
	env_id          = harness_platform_environment.test.id
	deployment_type = "Kubernetes"
	type            = "KubernetesDirect"

	yaml = <<-EOT
infrastructureDefinition:
  name: %[1]s
  identifier: %[1]s
  orgIdentifier: ${harness_platform_organization.test.id}
  projectIdentifier: ${harness_platform_project.test.id}
  environmentRef: ${harness_platform_environment.test.id}
  type: KubernetesDirect
  deploymentType: Kubernetes
  spec:
    connectorRef: ${harness_platform_connector_kubernetes.test.id}
    namespace: "chaos-test"
    releaseName: "release-%[1]s"
	EOT
}

resource "harness_chaos_infrastructure_v2" "test" {
	org_id         = harness_platform_organization.test.id
	project_id     = harness_platform_project.test.id
	environment_id = harness_platform_environment.test.id
	name           = "%[1]s"
	infra_id       = harness_platform_infrastructure.test.id
	description    = "Test infrastructure"
	infra_type     = "KUBERNETESV2"
	infra_scope    = "NAMESPACE"
	namespace      = "chaos-test"
	service_account = "litmus"
}

resource "harness_chaos_experiment" "test" {
	org_id            = harness_platform_organization.test.id
	project_id        = harness_platform_project.test.id
	template_identity = harness_chaos_experiment_template.test.identity

	# Hub scope (where template lives) - project-level
	hub_org_id        = harness_platform_organization.test.id
	hub_project_id    = harness_platform_project.test.id
	hub_identity      = harness_chaos_hub_v2.test.identity

	name              = "%[1]s"
	# LOCAL import requires format: environment_id/infra_id
	infra_ref         = "${harness_platform_environment.test.id}/${harness_chaos_infrastructure_v2.test.infra_id}"
	description       = "Test experiment with LOCAL import"
	import_type       = "LOCAL"

	tags = ["test", "local"]
}
`, name)
}
