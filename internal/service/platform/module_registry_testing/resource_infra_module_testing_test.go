package module_registry_testing_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceModuleTesting(t *testing.T) {
	resourceName := "harness_platform_infra_module_testing.test"
	name := "modulename"
	system := "terraform"
	org := "default"
	project := "ManualTest"
	provisionerType := "terraform"
	provisionerVersion := "1.5.0"
	updatedProvisionerVersion := "1.6.0"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceModuleTestingDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceModuleTesting(name, system, org, project, provisionerType, provisionerVersion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "testing_metadata.0.org", org),
					resource.TestCheckResourceAttr(resourceName, "testing_metadata.0.project", project),
					resource.TestCheckResourceAttr(resourceName, "testing_metadata.0.provisioner_type", provisionerType),
					resource.TestCheckResourceAttr(resourceName, "testing_metadata.0.provisioner_version", provisionerVersion),
					resource.TestCheckResourceAttr(resourceName, "testing_enabled", "true"),
				),
			},
			{
				Config: testAccResourceModuleTesting(name, system, org, project, provisionerType, updatedProvisionerVersion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "testing_metadata.0.provisioner_version", updatedProvisionerVersion),
					resource.TestCheckResourceAttr(resourceName, "testing_enabled", "true"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.AccountLevelResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

func TestAccResourceModuleTestingWithReleasePipeline(t *testing.T) {
	resourceName := "harness_platform_infra_module_testing.test"
	name := strings.ToLower(t.Name())
	system := "terraform"
	org := "default"
	project := "ManualTest"
	provisionerType := "terraform"
	provisionerVersion := "1.5.0"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceModuleTestingDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceModuleTestingWithReleasePipeline(name, system, org, project, provisionerType, provisionerVersion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "testing_metadata.0.org", org),
					resource.TestCheckResourceAttr(resourceName, "testing_metadata.0.project", project),
					resource.TestCheckResourceAttr(resourceName, "testing_metadata.0.provisioner_type", provisionerType),
					resource.TestCheckResourceAttr(resourceName, "testing_metadata.0.provisioner_version", provisionerVersion),
					resource.TestCheckResourceAttr(resourceName, "testing_enabled", "true"),
				),
			},
		},
	})
}

func testAccResourceModuleTestingDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		mod, _ := testAccGetPlatformModuleTesting(resourceName, state)
		if mod != nil && mod.TestingEnabled {
			return fmt.Errorf("module testing still enabled for module: %s", mod.Name)
		}
		return nil
	}
}

func testAccGetPlatformModuleTesting(resourceName string, state *terraform.State) (*nextgen.ModuleResource, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	module, resp, err := c.ModuleRegistryApi.ModuleRegistryListModulesById(
		ctx,
		id,
		c.AccountId,
	)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return &module, nil
}

func testAccResourceModuleTesting(name, system, org, project, provisionerType, provisionerVersion string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "secret%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "terraformtesttest" {
			identifier = "github%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://github.com/account"
			connection_type = "Account"
			validation_repo = "some_repo"
			delegate_selectors = ["harness-delegate"]
			credentials {
				http {
					username = "admin"
					token_ref = "account.${harness_platform_secret_text.test.id}"
				}
			}
		}

		resource "harness_platform_connector_aws" "terraformtestaws" {
			identifier = "aws%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			manual {
				access_key_ref = "account.${harness_platform_secret_text.test.id}"
				secret_key_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		resource "harness_platform_infra_module" "terraformtestmodule" {
			name = "module%[1]s"
			system = "system%[2]s"
			description = "description"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "tf/aws/basic"
			repository_connector = "account.${harness_platform_connector_github.terraformtesttest.id}"
		}

		resource "harness_platform_infra_module_testing" "test" {
			module_id = harness_platform_infra_module.terraformtestmodule.id
			org = "%[3]s"
			project = "%[4]s"
			provider_connector = "account.${harness_platform_connector_aws.terraformtestaws.id}"
			provisioner_type = "%[5]s"
			provisioner_version = "%[6]s"
			pipelines = ["iacm_auto_generated_tofu_testing", "iacm_auto_generated_integration_testing"]
		}
	`, name, system, org, project, provisionerType, provisionerVersion)
}

func testAccResourceModuleTestingWithReleasePipeline(name, system, org, project, provisionerType, provisionerVersion string) string {
	return fmt.Sprintf(`
		resource "harness_platform_secret_text" "test" {
			identifier = "secret%[1]s"
			name = "secret%[1]s"
			description = "test"
			tags = ["foo:bar"]

			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier = "github%[1]s"
			name = "github%[1]s"
			description = "test"
			tags = ["foo:bar"]

			url = "https://github.com/account"
			connection_type = "Account"
			validation_repo = "some_repo"
			delegate_selectors = ["harness-delegate"]
			credentials {
				http {
					username = "admin"
					token_ref = "account.${harness_platform_secret_text.test.id}"
				}
			}
		}

		resource "harness_platform_connector_aws" "test" {
			identifier = "aws%[1]s"
			name = "aws%[1]s"
			description = "test"
			tags = ["foo:bar"]

			manual {
				access_key_ref = "account.${harness_platform_secret_text.test.id}"
				secret_key_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		resource "harness_platform_infra_module" "test" {
			name = "module%[1]s"
			system = "system%[2]s"
			description = "description"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "tf/aws/basic"
			repository_connector = "account.${harness_platform_connector_github.test.id}"
		}

		resource "harness_platform_infra_module_testing" "test" {
			module_id = harness_platform_infra_module.test.id
			org = "%[3]s"
			project = "%[4]s"
			provider_connector = "account.${harness_platform_connector_aws.test.id}"
			provisioner_type = "%[5]s"
			provisioner_version = "%[6]s"
			pipelines = ["iacm_auto_generated_tofu_testing", "iacm_auto_generated_integration_testing"]
			release_pipeline = "iacm_auto_generated_sync_pipeline"
		}
	`, name, system, org, project, provisionerType, provisionerVersion)
}
