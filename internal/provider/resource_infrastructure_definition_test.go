package provider

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceInfrastructureDefinition_K8sDirect(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	updatedName := fmt.Sprintf("%s_updated", expectedName)
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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

func TestAccResourceInfrastructureDefinition_RenameCloudProvider(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	updatedName := fmt.Sprintf("%s_updated", expectedName)
	resourceName := "harness_infrastructure_definition.test"
	resourceCPName := "harness_cloudprovider_kubernetes.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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

func TestAccResourceInfrastructureDefinition_AwsSSH(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefAwsSSH(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", expectedName),
					resource.TestCheckResourceAttr(resourceName, "aws_ssh.0.hostname_convention", defaultHostnameConvention),
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

func TestAccResourceInfrastructureDefinition_AwsAmI_ASG(t *testing.T) {
	t.Skip("Yaml configuration not peristing properly https://harness.atlassian.net/browse/SWAT-5170")

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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

func TestAccResourceInfrastructureDefinition_AwsEcs_Fargate(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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

func TestAccResourceInfrastructureDefinition_AwsLambda(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefAwsAmi_Lambda(expectedName),
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

func TestAccResourceInfrastructureDefinition_AwsWinrm(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccInfraDefDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccInfraDefAwsAmi_WinRM(expectedName),
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

func TestAccResourceInfrastructureDefinition_Pcf(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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

func TestAccResourceInfrastructureDefinition_Azure_WebApp(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
	resourceName := "harness_infrastructure_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
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
	r := testAccGetResource(resourceName, state)
	if r == nil {
		return nil, fmt.Errorf("could not find resource %s", resourceName)
	}

	c := testAccGetApiClientFromProvider()
	id := r.Primary.ID
	appId := r.Primary.Attributes["app_id"]
	envId := r.Primary.Attributes["env_id"]

	return c.ConfigAsCode().GetInfraDefinitionById(appId, envId, id)
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
				delegate_selectors = ["test-account", "k8s"]
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
`, name, cloudproviderName)
}

func testAccInfraDefK8s(name string) string {
	return fmt.Sprintf(`
		resource "harness_cloudprovider_kubernetes" "test" {
			name = "%[1]s"

			authentication {
				delegate_selectors = ["test-account", "k8s"]
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

func testAccResourceInfraDefEnvironment(name string) string {
	return fmt.Sprintf(`
		resource "harness_application" "test" {
			name = "%[1]s"
		}

		resource "harness_environment" "test" {
			name = "%[1]s"
			app_id = harness_application.test.id
			type = "NON_PROD"
		}
`, name)
}

func testAccInfraDefAwsSSH(name string) string {
	aws := testAccResourceAwsCloudProvider(name)
	env := testAccResourceInfraDefEnvironment(name)

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

func testAccInfraDefAwsAmi_ASG(name string) string {
	aws := testAccResourceAwsCloudProvider(name)
	env := testAccResourceInfraDefEnvironment(name)

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
	aws := testAccResourceAwsCloudProvider(name)
	env := testAccResourceInfraDefEnvironment(name)

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

func testAccInfraDefAwsAmi_ECS_Fargate(name string) string {
	aws := testAccResourceAwsCloudProvider(name)
	env := testAccResourceInfraDefEnvironment(name)

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
	aws := testAccResourceAwsCloudProvider(name)
	env := testAccResourceInfraDefEnvironment(name)

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

func testAccInfraDefAwsAmi_Lambda(name string) string {
	aws := testAccResourceAwsCloudProvider(name)
	env := testAccResourceInfraDefEnvironment(name)

	return fmt.Sprintf(`
		%[1]s

		%[2]s

		resource "harness_infrastructure_definition" "test" {
			name = "%[3]s"
			app_id = harness_application.test.id
			env_id = harness_environment.test.id
			cloud_provider_type = "AWS"
			deployment_type = "AWS_LAMBDA"

			aws_lambda {
				cloud_provider_name = harness_cloudprovider_aws.test.name
				iam_role = "arn:aws:iam::123456789012:role/test-role"
				region = "us-west-2"
				security_group_ids = ["sg-12345678"]
				subnet_ids = ["subnet-e13232"]
				vpc_id = "vcp-xyz123"
			}
		}
`, aws, env, name)
}

func testAccInfraDefAwsAmi_WinRM(name string) string {
	aws := testAccResourceAwsCloudProvider(name)
	env := testAccResourceInfraDefEnvironment(name)

	return fmt.Sprintf(`
		%[1]s

		%[2]s

		resource "harness_infrastructure_definition" "test" {
			name = "%[3]s"
			app_id = harness_application.test.id
			env_id = harness_environment.test.id
			cloud_provider_type = "AWS"
			deployment_type = "WINRM"

			aws_winrm {
				autoscaling_group_name = "test-autoscaling-group"
				cloud_provider_name = harness_cloudprovider_aws.test.name
				desired_capacity = 1 
				host_connection_attrs_name = "winrm-test"
				host_connection_type = "PRIVATE_DNS"
				loadbalancer_name = "lb-test"
				region = "us-west-2"
			}
		}
`, aws, env, name)
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

func infraDefImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		id := primary.ID
		app_id := primary.Attributes["app_id"]
		env_id := primary.Attributes["env_id"]

		return fmt.Sprintf("%s/%s/%s", app_id, env_id, id), nil
	}
}
