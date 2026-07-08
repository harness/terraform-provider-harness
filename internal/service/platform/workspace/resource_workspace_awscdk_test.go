package workspace_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccResourceWorkspace_AwsCdkPython tests basic AWS CDK workspace with Python
func TestAccResourceWorkspace_AwsCdkPython(t *testing.T) {
	resourceName := "harness_platform_workspace.test_cdk"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspaceAwsCdkPython(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "provisioner_type", "awscdk"),
					resource.TestCheckResourceAttr(resourceName, "provisioner_config.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "provisioner_config.*", map[string]string{
						"language":                "python",
						"language_version":        "3.12",
						"package_manager":         "pip",
						"package_manager_version": "25.3",
					}),
					resource.TestCheckResourceAttr(resourceName, "tags.0", "cdk-test"),
				),
			},
			{
				Config: testAccResourceWorkspaceAwsCdkPython(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "provisioner_type", "awscdk"),
					resource.TestCheckResourceAttr(resourceName, "provisioner_config.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// TestAccResourceWorkspace_AwsCdkTypescript tests AWS CDK workspace with TypeScript
func TestAccResourceWorkspace_AwsCdkTypescript(t *testing.T) {
	resourceName := "harness_platform_workspace.test_cdk_ts"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspaceAwsCdkTypescript(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "provisioner_type", "awscdk"),
					resource.TestCheckResourceAttr(resourceName, "provisioner_config.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "provisioner_config.*", map[string]string{
						"language":                "typescript",
						"language_version":        "24.13.0",
						"package_manager":         "npm",
						"package_manager_version": "11.10.0",
					}),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// TestAccResourceWorkspace_AwsCdkUpdate tests updating provisioner_config fields
func TestAccResourceWorkspace_AwsCdkUpdate(t *testing.T) {
	resourceName := "harness_platform_workspace.test_cdk_update"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspaceAwsCdkWithVersion(id, name, "3.13", "25.3"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "provisioner_type", "awscdk"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "provisioner_config.*", map[string]string{
						"language":                "python",
						"language_version":        "3.13",
						"package_manager":         "pip",
						"package_manager_version": "25.3",
					}),
				),
			},
			{
				Config: testAccResourceWorkspaceAwsCdkWithVersion(id, name, "3.13", "25.3"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "provisioner_type", "awscdk"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "provisioner_config.*", map[string]string{
						"language":                "python",
						"language_version":        "3.13",
						"package_manager":         "pip",
						"package_manager_version": "25.3",
					}),
				),
			},
		},
	})
}

// TestAccResourceWorkspace_AwsCdkWithConnectors tests AWS CDK workspace with provider connectors
func TestAccResourceWorkspace_AwsCdkWithConnectors(t *testing.T) {
	resourceName := "harness_platform_workspace.test_cdk_connectors"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspaceAwsCdkWithConnectors(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "provisioner_type", "awscdk"),
					resource.TestCheckResourceAttr(resourceName, "provisioner_config.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connector.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "connector.*", map[string]string{
						"type": "aws",
					}),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateIdFunc: acctest.ProjectResourceImportStateIdFunc(resourceName),
			},
		},
	})
}

// Helper function: AWS CDK workspace with Python
func testAccResourceWorkspaceAwsCdkPython(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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

		resource "harness_platform_workspace" "test_cdk" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "test workspace with aws cdk python"
			provisioner_type = "awscdk"
			provisioner_version = "2.1108.0"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "cdk/python"
			cost_estimation_enabled = false
			repository_connector = "account.${harness_platform_connector_github.test.id}"

			provisioner_config {
				language = "python"
				language_version = "3.12"
				package_manager = "pip"
				package_manager_version = "25.3"
			}

			tags = ["cdk-test"]
		}
	`, id, name)
}

// Helper function: AWS CDK workspace with TypeScript
func testAccResourceWorkspaceAwsCdkTypescript(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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

		resource "harness_platform_workspace" "test_cdk_ts" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "test workspace with aws cdk typescript"
			provisioner_type = "awscdk"
			provisioner_version = "2.1108.0"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "cdk/typescript"
			cost_estimation_enabled = false
			repository_connector = "account.${harness_platform_connector_github.test.id}"

			provisioner_config {
				language = "typescript"
				language_version = "24.13.0"
				package_manager = "npm"
				package_manager_version = "11.10.0"
			}

			tags = ["cdk-test", "typescript"]
		}
	`, id, name)
}

// Helper function: AWS CDK workspace with parameterized versions
func testAccResourceWorkspaceAwsCdkWithVersion(id string, name string, langVersion string, pkgVersion string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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

		resource "harness_platform_workspace" "test_cdk_update" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "test workspace with aws cdk version update"
			provisioner_type = "awscdk"
			provisioner_version = "2.1108.0"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "cdk/python"
			cost_estimation_enabled = false
			repository_connector = "account.${harness_platform_connector_github.test.id}"

			provisioner_config {
				language = "python"
				language_version = "%[3]s"
				package_manager = "pip"
				package_manager_version = "%[4]s"
			}

			tags = ["cdk-test", "update-test"]
		}
	`, id, name, langVersion, pkgVersion)
}

// Helper function: AWS CDK workspace with provider connectors
func testAccResourceWorkspaceAwsCdkWithConnectors(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_organization" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
		}

		resource "harness_platform_project" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
		}

		resource "harness_platform_secret_text" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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

		resource "harness_platform_workspace" "test_cdk_connectors" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "test workspace with aws cdk and connectors"
			provisioner_type = "awscdk"
			provisioner_version = "2.1108.0"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "cdk/python"
			cost_estimation_enabled = false
			repository_connector = "account.${harness_platform_connector_github.test.id}"

			provisioner_config {
				language = "python"
				language_version = "3.12"
				package_manager = "pip"
				package_manager_version = "25.3"
			}

			connector {
				connector_ref = "aws_connector_ref"
				type = "aws"
			}

			tags = ["cdk-test", "connector-test"]
		}
	`, id, name)
}
