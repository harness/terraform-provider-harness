package infrastructure_v2_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// sanitizeK8sResourceName converts a string to be compatible with Kubernetes resource name requirements
func sanitizeK8sResourceName(name string) string {
	// Convert to lowercase
	name = strings.ToLower(name)

	// Replace invalid characters with '-'
	re := regexp.MustCompile(`[^a-z0-9-]`)
	name = re.ReplaceAllString(name, "-")

	// Remove leading/trailing dashes
	name = strings.Trim(name, "-")

	// Ensure the name is not empty
	if name == "" {
		name = "infra"
	}

	// Truncate to 63 characters if needed
	if len(name) > 63 {
		name = name[:63]
	}

	return name
}

// TestAccDataSourceChaosInfrastructureV2_basic verifies the basic data source functionality for Chaos Infrastructure V2.
func TestAccDataSourceChaosInfrastructureV2_basic(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "data.harness_chaos_infrastructure_v2.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosInfrastructureV2Config(id, rName, "KUBERNETESV2", "NAMESPACE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", sanitizeK8sResourceName(rName)),
					// resource.TestCheckResourceAttr(resourceName, "infra_type", "KUBERNETESV2"),
					resource.TestCheckResourceAttr(resourceName, "infra_scope", "NAMESPACE"),
					// resource.TestCheckResourceAttrSet(resourceName, "status"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
				),
			},
		},
	})
}

// TestAccDataSourceChaosInfrastructureV2_WithAllOptions verifies the data source with all possible options set.
func TestAccDataSourceChaosInfrastructureV2_WithAllOptions(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "data.harness_chaos_infrastructure_v2.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosInfrastructureV2Config_WithAllOptions(id, rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", sanitizeK8sResourceName(rName)),
					// resource.TestCheckResourceAttr(resourceName, "infra_type", "KUBERNETESV2"),
					resource.TestCheckResourceAttr(resourceName, "ai_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "insecure_skip_verify", "true"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "chaos-namespace"),
					resource.TestCheckResourceAttr(resourceName, "service_account", "litmus-admin"),
					// resource.TestCheckResourceAttrSet(resourceName, "status"),
				),
			},
		},
	})
}

// TestAccDataSourceChaosInfrastructureV2_KubernetesType verifies the data source for Kubernetes infra type.
func TestAccDataSourceChaosInfrastructureV2_KubernetesType(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "data.harness_chaos_infrastructure_v2.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceChaosInfrastructureV2Config(id, rName, "KUBERNETES", "NAMESPACE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "infra_type", "KUBERNETESV2"),
					resource.TestCheckResourceAttr(resourceName, "infra_scope", "NAMESPACE"),
				),
			},
		},
	})
}

// TestAccDataSourceChaosInfrastructureV2_NotFound verifies the behavior when the infra does not exist.
func TestAccDataSourceChaosInfrastructureV2_NotFound(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceChaosInfrastructureV2Config_NonExistent(id, rName),
				ExpectError: regexp.MustCompile("not found"),
			},
		},
	})
}

// Terraform Configurations

func testAccDataSourceChaosInfrastructureV2Config(id, name, infraType, infraScope string) string {
	return fmt.Sprintf(`
		// 1. Create Organization
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		// 2. Create Project
		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
		}

		// 3. Create Kubernetes Connector
		resource "harness_platform_connector_kubernetes" "test" {
            identifier  = "%[1]s"
            name        = "%[2]s"
            org_id     = harness_platform_organization.test.id
            project_id = harness_platform_project.test.id

            inherit_from_delegate {
                delegate_selectors = ["kubernetes-delegate"]
            }

            tags = []
        }

		// 4. Create Environment
		resource "harness_platform_environment" "test" {
            identifier = "%[1]s"
            name       = "%[2]s"
            org_id     = harness_platform_organization.test.id
            project_id = harness_platform_project.test.id
            type       = "PreProduction"
        }

		// 5. Create Harness Infrastructure Definition
		resource "harness_platform_infrastructure" "test" {
		    identifier  = "%[1]s"
		    name        = "%[2]s"
	        org_id     = harness_platform_organization.test.id
		    project_id = harness_platform_project.test.id
		    env_id     = harness_platform_environment.test.id
		    deployment_type = "Kubernetes"
		    type       = "KubernetesDirect"

            yaml = <<-EOT
            infrastructureDefinition:
                name: "%[2]s"
                identifier: "%[1]s"
                orgIdentifier: ${harness_platform_organization.test.id}
                projectIdentifier: ${harness_platform_project.test.id}
                environmentRef: ${harness_platform_environment.test.id}
                type: KubernetesDirect
                deploymentType: Kubernetes
                allowSimultaneousDeployments: false
                spec:
                    connectorRef: ${harness_platform_connector_kubernetes.test.id}
                    namespace: "chaos"
                    releaseName: "release-%[1]s"
            EOT
            tags = []
        }

		// 6. Create Chaos Infrastructure
		resource "harness_chaos_infrastructure_v2" "test" {
			org_id        = harness_platform_organization.test.id
			project_id    = harness_platform_project.test.id
			environment_id = harness_platform_environment.test.id
			name          = "%[2]s"
			infra_id      = harness_platform_infrastructure.test.id
			description   = "Test Infrastructure"
			infra_type    = "%[3]s"
			infra_scope   = "%[4]s"
			namespace     = "chaos"
			service_account = "litmus"
		}

		// 7. Create Data Source
		data "harness_chaos_infrastructure_v2" "test" {
			org_id        = harness_platform_organization.test.id
			project_id    = harness_platform_project.test.id
			environment_id = harness_platform_environment.test.id
			infra_id      = harness_chaos_infrastructure_v2.test.infra_id
		}
	`, id, name, infraType, infraScope)
}

func testAccDataSourceChaosInfrastructureV2Config_WithAllOptions(id, name string) string {
	return fmt.Sprintf(`
		// 1. Create Organization
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		// 2. Create Project
		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
		}

		// 3. Create Kubernetes Connector
		resource "harness_platform_connector_kubernetes" "test" {
            identifier  = "%[1]s"
            name        = "%[2]s"
            org_id     = harness_platform_organization.test.id
            project_id = harness_platform_project.test.id

            inherit_from_delegate {
                delegate_selectors = ["kubernetes-delegate"]
            }

            tags = []
        }

		// 4. Create Environment
		resource "harness_platform_environment" "test" {
            identifier = "%[1]s"
            name       = "%[2]s"
            org_id     = harness_platform_organization.test.id
            project_id = harness_platform_project.test.id
            type       = "PreProduction"
        }

		// 5. Create Harness Infrastructure Definition
		resource "harness_platform_infrastructure" "test" {
		    identifier  = "%[1]s"
		    name        = "%[2]s"
	        org_id     = harness_platform_organization.test.id
		    project_id = harness_platform_project.test.id
		    env_id     = harness_platform_environment.test.id
		    deployment_type = "Kubernetes"
		    type       = "KubernetesDirect"

            yaml = <<-EOT
            infrastructureDefinition:
                name: "%[2]s"
                identifier: "%[1]s"
                orgIdentifier: ${harness_platform_organization.test.id}
                projectIdentifier: ${harness_platform_project.test.id}
                environmentRef: ${harness_platform_environment.test.id}
                type: KubernetesDirect
                deploymentType: Kubernetes
                allowSimultaneousDeployments: false
                spec:
                    connectorRef: ${harness_platform_connector_kubernetes.test.id}
                    namespace: "chaos"
                    releaseName: "release-%[1]s"
            EOT
            tags = []
        }

		resource "harness_chaos_infrastructure_v2" "test" {
			org_id             = harness_platform_organization.test.id
			project_id         = harness_platform_project.test.id
			environment_id     = harness_platform_environment.test.id
			name               = "%[2]s"
			infra_id           = harness_platform_infrastructure.test.id
			description        = "Test Infrastructure with all options"
			infra_type    	   = "KUBERNETESV2"
			infra_scope        = "CLUSTER"
			namespace          = "chaos-namespace"
			service_account    = "litmus-admin"
			ai_enabled         = true
			insecure_skip_verify = true
		}

		data "harness_chaos_infrastructure_v2" "test" {
			org_id        = harness_platform_organization.test.id
			project_id    = harness_platform_project.test.id
			environment_id = harness_platform_environment.test.id
			infra_id      = harness_chaos_infrastructure_v2.test.infra_id
		}
	`, id, name)
}

func testAccDataSourceChaosInfrastructureV2Config_NonExistent(id, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
		}

		resource "harness_platform_environment" "test" {
			identifier = "%[1]s"
			name       = "%[2]s"
			org_id     = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			type       = "PreProduction"
		}

		data "harness_chaos_infrastructure_v2" "test" {
			org_id        = harness_platform_organization.test.id
			project_id    = harness_platform_project.test.id
			environment_id = harness_platform_environment.test.id
			infra_id      = "%[1]s"
		}
	`, id, name)
}
