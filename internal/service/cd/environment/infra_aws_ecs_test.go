package environment_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceInfrastructureDefinition_AwsEcs_Fargate(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefAwsAmi_ECS_Fargate(expectedName),
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

func TestAccResourceInfrastructureDefinition_AwsEcs_EC2(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefAwsAmi_ECS_EC2(expectedName),
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

func testAccInfraDefAwsAmi_ECS_Fargate(name string) string {
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
			deployment_type = "ECS"

			aws_ecs {
				assign_public_ip = true
				cloud_provider_name = harness_cloudprovider_aws.test.name
				cluster_name = "test-cluster"
				execution_role = "arn::some::accountId:role/testrole"
				launch_type = "FARGATE"
				region = "us-west-2"
				security_group_ids = ["sg-12345678"]
				subnet_ids = ["subnet-e13232"]
				vpc_id = "vcp-xyz123"
			}
		}
`, aws, env, name)
}

func testAccInfraDefAwsAmi_ECS_EC2(name string) string {
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
			deployment_type = "ECS"

			aws_ecs {
				assign_public_ip = true
				cloud_provider_name = harness_cloudprovider_aws.test.name
				cluster_name = "test-cluster"
				launch_type = "EC2"
				region = "us-west-2"
				security_group_ids = ["sg-12345678"]
				subnet_ids = ["subnet-e13232"]
				vpc_id = "vcp-xyz123"
			}
		}
`, aws, env, name)
}
