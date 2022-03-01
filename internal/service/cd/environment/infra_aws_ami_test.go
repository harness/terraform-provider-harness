package environment_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceInfrastructureDefinition_AwsAmI_ASG(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefAwsAmi_ASG(expectedName),
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

func TestAccResourceInfrastructureDefinition_AwsAmI_Spot(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	resourceName := "harness_infrastructure_definition.test"
	config := testAccInfraDefAwsAmi_SpotInst(expectedName)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: config,
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

func testAccInfraDefAwsAmi_ASG(name string) string {
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
			deployment_type = "AMI"

			aws_ami {
				ami_deployment_type = "AWS_ASG"
				asg_identifies_workload = true
				autoscaling_group_name = "EC2ContainerService-wfmx-ECSDelegate-EcsInstanceAsg-122ZBWFY19B8E"
				classic_loadbalancers = ["af16610fdd8d011e99b5c0eeaa137c8d"]
				cloud_provider_name = harness_cloudprovider_aws.test.name
				region = "us-west-2"
				stage_classic_loadbalancers = ["af16610fdd8d011e99b5c0eeaa137c8d"]
				stage_target_group_arns = ["arn:aws:elasticloadbalancing:us-west-2:759984737373:targetgroup/Prod/a37c86dbe0700bfd"]
				target_group_arns = ["arn:aws:elasticloadbalancing:us-west-2:759984737373:targetgroup/Prod/a37c86dbe0700bfd"]
				use_traffic_shift = false
			}
		}
`, aws, env, name)
}

func testAccInfraDefAwsAmi_SpotInst(name string) string {
	aws := acctest.TestAccResourceAwsCloudProvider(name)
	env := acctest.TestAccResourceInfraDefEnvironment(name)

	return fmt.Sprintf(`
		%[1]s

		%[2]s

    resource "harness_encrypted_text" "test" {
			name = "%s"
			secret_manager_id = data.harness_secret_manager.default.id
			value = "%[5]s"
		}

		resource "harness_cloudprovider_spot" "test" {
			name = "%[3]s_spot"
			account_id = "%[4]s"
			token_secret_name = harness_encrypted_text.test.name
		}
		
		resource "harness_infrastructure_definition" "test" {
			name = "%[3]s"
			app_id = harness_application.test.id
			env_id = harness_environment.test.id
			cloud_provider_type = "AWS"
			deployment_type = "AMI"

			aws_ami {
				ami_deployment_type = "SPOTINST"
				asg_identifies_workload = false
				cloud_provider_name = harness_cloudprovider_aws.test.name
				region = "us-west-2"
				spotinst_cloud_provider_name = harness_cloudprovider_spot.test.name
				spotinst_config_json = <<EOF
					{
						"test": "test"
					}
				EOF

				use_traffic_shift = true
			}
		}
`, aws, env, name, helpers.TestEnvVars.SpotInstAccountId.Get(), helpers.TestEnvVars.SpotInstToken.Get())
}
