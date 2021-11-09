package environment_test

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceInfrastructureDefinition_Azure_WebApp(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefAzure_webapp(expectedName),
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

func testAccInfraDefAzure_webapp(name string) string {
	return fmt.Sprintf(`
		data "harness_secret_manager" "default" {
			default = true
		}

		resource "harness_encrypted_text" "test" {
			name = "%[1]s"
			value = "%[2]s"
			secret_manager_id = data.harness_secret_manager.default.id
		}

		resource "harness_cloudprovider_azure" "test" {
			name = "%[1]s"
			client_id = "%[3]s"
			tenant_id = "%[4]s"
			key = harness_encrypted_text.test.name
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
			cloud_provider_type = "AZURE"
			deployment_type = "AZURE_WEBAPP"

			azure_webapp {
				cloud_provider_name = harness_cloudprovider_azure.test.name
				resource_group = "test-rg"
				subscription_id = "test-sub"
			}
		}
`, name, helpers.TestEnvVars.AzureClientSecret.Get(), helpers.TestEnvVars.AzureClientId.Get(), helpers.TestEnvVars.AzureTenantId.Get())
}
