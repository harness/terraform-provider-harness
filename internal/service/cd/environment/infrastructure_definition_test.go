package environment_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceInfrastructureDefinition_RenameCloudProvider(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	updatedName := fmt.Sprintf("%s_updated", expectedName)
	resourceName := "harness_infrastructure_definition.test"
	resourceCPName := "harness_cloudprovider_kubernetes.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefK8s_Rename_CloudProvider(expectedName, expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccInfraDefCreation(t, resourceName, expectedName),
				),
			},
			{
				Config: testAccInfraDefK8s_Rename_CloudProvider(expectedName, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					resource.TestCheckResourceAttr(resourceCPName, "name", updatedName),
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

func testAccInfraDefCreation(t *testing.T, resourceName string, name string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		infraDef, err := testAccGetInfraDef(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, infraDef)
		require.Equal(t, name, infraDef.Name)

		return nil
	}
}

func testAccGetInfraDef(resourceName string, state *terraform.State) (*cac.InfrastructureDefinition, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	if r == nil {
		return nil, fmt.Errorf("could not find resource %s", resourceName)
	}

	c := acctest.TestAccGetApiClientFromProvider()
	id := r.Primary.ID
	appId := r.Primary.Attributes["app_id"]
	envId := r.Primary.Attributes["env_id"]

	return c.CDClient.ConfigAsCodeClient.GetInfraDefinitionById(appId, envId, id)
}

func testAccInfraDefDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		infraDef, _ := testAccGetInfraDef(resourceName, state)
		if infraDef != nil {
			return fmt.Errorf("Found infrasturcture definition: %s", infraDef.Id)
		}

		return nil
	}
}

func testAccInfraDefK8s_Rename_CloudProvider(name string, cloudproviderName string) string {
	return fmt.Sprintf(`
		resource "harness_cloudprovider_kubernetes" "test" {
			name = "%[2]s"

			authentication {
				delegate_selectors = ["primary"]
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

			lifecycle {
				create_before_destroy = true
			}
		}
`, name, cloudproviderName)
}

func infraDefImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		id := primary.ID
		app_id := primary.Attributes["app_id"]
		env_id := primary.Attributes["env_id"]

		return fmt.Sprintf("%s/%s/%s", app_id, env_id, id), nil
	}
}
