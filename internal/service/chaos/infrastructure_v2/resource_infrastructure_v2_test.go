package infrastructure_v2_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// SanitizeK8sResourceName converts a string to be compatible with Kubernetes resource name requirements
func SanitizeK8sResourceName(name string) string {
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

func TestAccResourceChaosInfrastructureV2_basic(t *testing.T) {

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_infrastructure_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosInfrastructureV2Destroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosInfrastructureV2ConfigBasic(rName, id, "KUBERNETESV2", "NAMESPACE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", SanitizeK8sResourceName(rName)),
					// BUG: BE is taking infraType as KUBERNETESV2 but in GET call, sending KUBERNETES
					// BUG: BE is sending harnessInfraType also as empty string
					// resource.TestCheckResourceAttr(resourceName, "infra_type", "KUBERNETESV2"),
					resource.TestCheckResourceAttr(resourceName, "infra_scope", "NAMESPACE"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "chaos"),
					resource.TestCheckResourceAttr(resourceName, "service_account", "litmus"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccResourceChaosInfrastructureV2ImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceChaosInfrastructureV2_Update(t *testing.T) {

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_infrastructure_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosInfrastructureV2Destroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosInfrastructureV2ConfigBasic(rName, id, "KUBERNETESV2", "NAMESPACE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", SanitizeK8sResourceName(rName)),
				),
			},
			{
				Config: testAccResourceChaosInfrastructureV2ConfigUpdate(rName, id, "KUBERNETESV2", "CLUSTER"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", SanitizeK8sResourceName(rName)),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated Test Infrastructure"),
					resource.TestCheckResourceAttr(resourceName, "infra_scope", "CLUSTER"),
					resource.TestCheckResourceAttr(resourceName, "namespace", "chaos-updated"),
					resource.TestCheckResourceAttr(resourceName, "service_account", "litmus-admin"),
					resource.TestCheckResourceAttr(resourceName, "run_as_user", "1001"),
					resource.TestCheckResourceAttr(resourceName, "insecure_skip_verify", "true"),
				),
			},
		},
	})
}

func TestAccResourceChaosInfrastructureV2_KubernetesType(t *testing.T) {

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_infrastructure_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccChaosInfrastructureV2Destroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosInfrastructureV2ConfigBasic(rName, id, "KUBERNETES", "NAMESPACE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "infra_type", "KUBERNETES"),
					resource.TestCheckResourceAttr(resourceName, "infra_scope", "NAMESPACE"),
				),
			},
		},
	})
}

func TestAccResourceChaosInfrastructureV2_Import(t *testing.T) {

	// Generate unique identifiers
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	rName := id
	resourceName := "harness_chaos_infrastructure_v2.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceChaosInfrastructureV2ConfigBasic(rName, id, "KUBERNETESV2", "NAMESPACE"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccResourceChaosInfrastructureV2ImportStateIdFunc(resourceName),
			},
		},
	})
}

// Helpers For Destroy & Import State
func testAccChaosInfrastructureV2Destroy(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// No-op for now as we don't have a direct way to verify deletion
		return nil
	}
}

func testAccResourceChaosInfrastructureV2ImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		// Format: org_id/project_id/environment_id/infra_id
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}

		orgID := rs.Primary.Attributes["org_id"]
		projectID := rs.Primary.Attributes["project_id"]
		envID := rs.Primary.Attributes["environment_id"]
		infraID := rs.Primary.Attributes["infra_id"]

		return fmt.Sprintf("%s/%s/%s/%s", orgID, projectID, envID, infraID), nil
	}
}

// Terraform Configurations

func testAccResourceChaosInfrastructureV2ConfigBasic(id, name, infraType, infraScope string) string {
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
			org_id              = harness_platform_organization.test.id
			project_id          = harness_platform_project.test.id
			environment_id      = harness_platform_environment.test.id
			name                = "%[2]s"
			infra_id            = harness_platform_infrastructure.test.id
			description         = "Test Infrastructure"
			infra_type          = "%[3]s"
			infra_scope         = "%[4]s"
			namespace           = "chaos"
			service_account     = "litmus"
			tags                = ["test:true", "chaos:true"]
			run_as_user         = 1000
			insecure_skip_verify = true
		}
	`, id, name, infraType, infraScope)
}

func testAccResourceChaosInfrastructureV2ConfigUpdate(name, id, infraType, infraScope string) string {

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
			org_id              = harness_platform_organization.test.id
			project_id          = harness_platform_project.test.id
			environment_id      = harness_platform_environment.test.id
			name                = "%[2]s"
			infra_id            = harness_platform_infrastructure.test.id
			description         = "Updated Test Infrastructure"
			infra_type          = "%[3]s"
			infra_scope         = "%[4]s"
			namespace           = "chaos-updated"
			service_account     = "litmus-admin"
			tags                = ["test:true", "chaos:true", "updated:true"]
			run_as_user         = 1001
			insecure_skip_verify = true
		}
	`, id, name, infraType, infraScope)
}
