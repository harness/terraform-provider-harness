package environment_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceInfrastructureDefinition_Custom(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefCustom(expectedName),
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

func testAccInfraDefCustom(name string) string {
	return fmt.Sprintf(`
		resource "harness_application" "test" {
			name = "%[1]s"
		}

		resource "harness_environment" "test" {
			name = "%[1]s"
			app_id = harness_application.test.id
			type = "NON_PROD"
		}

		resource "harness_yaml_config" "test" {
			path = "Setup/Template Library/AutomationOneNow/custom_test.yaml"
			content = <<EOF
harnessApiVersion: '1.0'
type: CUSTOM_DEPLOYMENT_TYPE
fetchInstanceScript: curl http://localhost:8081/instances.json >> $${INSTANCE_OUTPUT_PATH}
hostAttributes:
  hostname: host
hostObjectArrayPath: hosts
variables:
- name: url
- name: file_name
EOF
		}

		resource "harness_infrastructure_definition" "test" {
			name = "%[1]s"
			app_id = harness_application.test.id
			env_id = harness_environment.test.id
			cloud_provider_type = "CUSTOM"
			deployment_type = "CUSTOM"
			deployment_template_uri = "AutomationOneNow/${harness_yaml_config.test.name}"

			custom {
				deployment_type_template_version = "1"
				variable {
					name = "url"
					value = "localhost:8081"
				}
	
				variable {
					name = "file_name"
					value = "instances.json"
				}
			}
		}
`, name)
}
