package environment_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceInfrastructureDefinition_Tanzu(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefTanzu(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					testAccInfraDefCreation(t, resourceName, expectedName),
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

func testAccInfraDefTanzu(name string) string {
	return fmt.Sprintf(`
		data "harness_secret_manager" "default" {
			default = true
		}

		resource "harness_encrypted_text" "test" {
			name = "%[1]s"
			value = "foo"
			secret_manager_id = data.harness_secret_manager.default.id
		}
		
		resource "harness_cloudprovider_tanzu" "test" {
			name = "%[1]s"
			endpoint = "https://endpoint.com"
			skip_validation = true
			username = "username"
			password_secret_name = harness_encrypted_text.test.name
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
			cloud_provider_type = "PCF"
			deployment_type = "PCF"

			tanzu {
				cloud_provider_name = harness_cloudprovider_tanzu.test.name
				organization = "test-org"
				space = "test-space"
			}
		}
`, name)
}
