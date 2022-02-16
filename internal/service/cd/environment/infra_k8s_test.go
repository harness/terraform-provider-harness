package environment_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceInfrastructureDefinition_K8sDirect(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	updatedName := fmt.Sprintf("%s_updated", expectedName)
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefK8s(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccInfraDefCreation(t, resourceName, expectedName),
				),
			},
			{
				Config: testAccInfraDefK8s(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					testAccInfraDefCreation(t, resourceName, updatedName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: infraDefImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func testAccInfraDefK8s(name string) string {
	return fmt.Sprintf(`
		resource "harness_cloudprovider_kubernetes" "test" {
			name = "%[1]s"
			skip_validation = true
			authentication {
				delegate_selectors = ["k8s"]
			}
		}

		resource "harness_application" "test" {
			name = "%[1]s"
		}

		resource "harness_environment" "test" {
			name = "%[1]s"
			app_id = harness_application.test.id
			type = "NON_PROD"
		}

		resource "harness_infrastructure_definition" "test" {
			name = "%[1]s"
			app_id = harness_application.test.id
			env_id = harness_environment.test.id
			cloud_provider_type = "KUBERNETES_CLUSTER"
			deployment_type = "KUBERNETES"

			kubernetes {
				cloud_provider_name = harness_cloudprovider_kubernetes.test.name
				namespace = "testing"
				release_name = "release-$${infra.kubernetes.infraId}"
			}
		}
`, name)
}
