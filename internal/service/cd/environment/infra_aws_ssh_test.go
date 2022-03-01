package environment_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/harness/terraform-provider-harness/internal/service/cd/environment"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceInfrastructureDefinition_AwsSSH(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefAwsSSH(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					resource.TestCheckResourceAttr(resourceName, "aws_ssh.0.hostname_convention", environment.DefaultHostnameConvention),
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

func testAccInfraDefAwsSSH(name string) string {
	aws := acctest.TestAccResourceAwsCloudProvider(name)
	env := acctest.TestAccResourceInfraDefEnvironment(name)

	return fmt.Sprintf(`
		%[1]s

		%[2]s

		resource "harness_infrastructure_definition" "test" {
			name = "%[3]s"
			app_id = harness_application.test.id
			env_id = harness_environment.test.id
			cloud_provider_type = "AWS"
			deployment_type = "SSH"

			aws_ssh {
				
				tag {
					key = "Name"
					value = "test-instance"
				}
				
				vpc_ids = ["vpc-12345678"]

				cloud_provider_name = harness_cloudprovider_aws.test.name
				region = "us-west-2"
				desired_capacity = 1
				host_connection_type = "PRIVATE_DNS"
				host_connection_attrs_name = "test-ssh-cred"
			}
		}
`, aws, env, name)
}
