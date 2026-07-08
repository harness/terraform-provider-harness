package workspace_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccResourceWorkspace_VaultConnector tests basic Vault connector functionality
func TestAccResourceWorkspace_VaultConnector(t *testing.T) {
	resourceName := "harness_platform_workspace.test_vault"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))
	updatedName := fmt.Sprintf("%s_updated", name)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspaceWithVault(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "connector.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "connector.*", map[string]string{
						"type": "vault",
					}),
				),
			},
			{
				Config: testAccResourceWorkspaceWithVault(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "connector.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "connector.*", map[string]string{
						"type": "vault",
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

// TestAccResourceWorkspace_VaultConnectorUpdate tests updating Vault connector reference
func TestAccResourceWorkspace_VaultConnectorUpdate(t *testing.T) {
	resourceName := "harness_platform_workspace.test_vault_update"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspaceWithVaultConnector(id, name, "vault_connector_1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "connector.*", map[string]string{
						"type": "vault",
					}),
				),
			},
			{
				Config: testAccResourceWorkspaceWithVaultConnector(id, name, "vault_connector_2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "connector.*", map[string]string{
						"type": "vault",
					}),
				),
			},
		},
	})
}

// TestAccResourceWorkspace_VaultAndAWSConnectors tests Vault with AWS connector
func TestAccResourceWorkspace_VaultAndAWSConnectors(t *testing.T) {
	resourceName := "harness_platform_workspace.test_vault_aws"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspaceWithVaultAndAWS(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "connector.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "connector.*", map[string]string{
						"type": "vault",
					}),
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

// TestAccResourceWorkspace_VaultMultipleConnectorTypes tests Vault with all connector types
func TestAccResourceWorkspace_VaultMultipleConnectorTypes(t *testing.T) {
	resourceName := "harness_platform_workspace.test_vault_multi"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspaceWithAllConnectorTypes(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "connector.#", "4"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "connector.*", map[string]string{
						"type": "vault",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "connector.*", map[string]string{
						"type": "aws",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "connector.*", map[string]string{
						"type": "azure",
					}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "connector.*", map[string]string{
						"type": "gcp",
					}),
				),
			},
		},
	})
}

// TestAccResourceWorkspace_VaultConnectorRemoval tests removing Vault connector
func TestAccResourceWorkspace_VaultConnectorRemoval(t *testing.T) {
	resourceName := "harness_platform_workspace.test_vault_removal"
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccResourceWorkspaceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceWorkspaceWithVault(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connector.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "connector.*", map[string]string{
						"type": "vault",
					}),
				),
			},
			{
				Config: testAccResourceWorkspaceWithoutConnectors(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					// Connector block should be absent or empty
					resource.TestCheckNoResourceAttr(resourceName, "connector.#"),
				),
			},
		},
	})
}

// TestAccResourceWorkspace_DuplicateVaultConnectors tests duplicate Vault connectors (should fail)
func TestAccResourceWorkspace_DuplicateVaultConnectors(t *testing.T) {
	name := t.Name()
	id := fmt.Sprintf("%s_%s", name, utils.RandStringBytes(5))

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceWorkspaceWithDuplicateVault(id, name),
				ExpectError: regexp.MustCompile("vault types must be unique for connectors"),
			},
		},
	})
}

// Helper function: Basic workspace with Vault connector
func testAccResourceWorkspaceWithVault(id string, name string) string {
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

		resource "harness_platform_connector_vault" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test vault connector"
			tags = ["vault:test"]
			vault_url = "https://vault.example.com"
			base_path = "secret"
			default = false
			delegate_selectors = ["harness-delegate"]
			auth_token {
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		resource "harness_platform_workspace" "test_vault" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "test workspace with vault connector"
			provisioner_type = "terraform"
			provisioner_version = "1.5.7"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "/"
			cost_estimation_enabled = false
			repository_connector = "account.${harness_platform_connector_github.test.id}"

			connector {
				connector_ref = "account.${harness_platform_connector_vault.test.id}"
				type = "vault"
			}

			tags = ["vault-test"]
		}
	`, id, name)
}

// Helper function: Workspace with specific Vault connector reference
func testAccResourceWorkspaceWithVaultConnector(id string, name string, vaultConnectorId string) string {
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
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
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

		resource "harness_platform_connector_vault" "test_1" {
			identifier = "vault_1_%[1]s"
			name = "vault_1_%[2]s"
			description = "test vault connector 1"
			vault_url = "https://vault1.example.com"
			base_path = "secret"
			default = false
			delegate_selectors = ["harness-delegate"]
			auth_token {
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		resource "harness_platform_connector_vault" "test_2" {
			identifier = "vault_2_%[1]s"
			name = "vault_2_%[2]s"
			description = "test vault connector 2"
			vault_url = "https://vault2.example.com"
			base_path = "secret"
			default = false
			delegate_selectors = ["harness-delegate"]
			auth_token {
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		resource "harness_platform_workspace" "test_vault_update" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "test workspace with vault connector update"
			provisioner_type = "terraform"
			provisioner_version = "1.5.7"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "/"
			cost_estimation_enabled = false
			repository_connector = "account.${harness_platform_connector_github.test.id}"

			connector {
				connector_ref = "account.${harness_platform_connector_vault.%[3]s.id}"
				type = "vault"
			}

			tags = ["vault-test", "update-test"]
		}
	`, id, name, vaultConnectorId)
}

// Helper function: Workspace with Vault and AWS connectors
func testAccResourceWorkspaceWithVaultAndAWS(id string, name string) string {
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
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
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

		resource "harness_platform_connector_vault" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test vault connector"
			vault_url = "https://vault.example.com"
			base_path = "secret"
			default = false
			delegate_selectors = ["harness-delegate"]
			auth_token {
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		resource "harness_platform_connector_awscc" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test aws connector"
			tags = ["aws:test"]
			account_id = "123456789012"
			features_enabled = ["OPTIMIZATION"]
			cross_account_access {
				role_arn = "arn:aws:iam::123456789012:role/HarnessCERole"
				external_id = "harness:123456789:abc"
			}
		}

		resource "harness_platform_workspace" "test_vault_aws" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "test workspace with vault and aws connectors"
			provisioner_type = "terraform"
			provisioner_version = "1.5.7"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "/"
			cost_estimation_enabled = false
			repository_connector = "account.${harness_platform_connector_github.test.id}"

			connector {
				connector_ref = "account.${harness_platform_connector_vault.test.id}"
				type = "vault"
			}

			connector {
				connector_ref = "account.${harness_platform_connector_awscc.test.id}"
				type = "aws"
			}

			tags = ["vault-test", "aws-test", "multi-connector"]
		}
	`, id, name)
}

// Helper function: Workspace with all connector types
func testAccResourceWorkspaceWithAllConnectorTypes(id string, name string) string {
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
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
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

		resource "harness_platform_connector_vault" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test vault connector"
			vault_url = "https://vault.example.com"
			base_path = "secret"
			default = false
			delegate_selectors = ["harness-delegate"]
			auth_token {
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		resource "harness_platform_connector_awscc" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test aws connector"
			account_id = "123456789012"
			features_enabled = ["OPTIMIZATION"]
			cross_account_access {
				role_arn = "arn:aws:iam::123456789012:role/HarnessCERole"
				external_id = "harness:123456789:abc"
			}
		}

		resource "harness_platform_connector_azure_cloud_cost" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test azure connector"
			features_enabled = ["OPTIMIZATION"]
			tenant_id = "tenant-id"
			subscription_id = "subscription-id"
			billing_export_spec {
				storage_account_name = "storage-account"
				container_name = "container"
				directory_name = "directory"
				report_name = "report"
				subscription_id = "subscription-id"
			}
		}

		resource "harness_platform_connector_gcp_cloud_cost" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test gcp connector"
			features_enabled = ["OPTIMIZATION"]
			gcp_project_id = "gcp-project-id"
			service_account_email = "service-account@gcp-project.iam.gserviceaccount.com"
			billing_export_spec {
				data_set_id = "dataset"
				table_id = "table"
			}
		}

		resource "harness_platform_workspace" "test_vault_multi" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "test workspace with all connector types"
			provisioner_type = "terraform"
			provisioner_version = "1.5.7"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "/"
			cost_estimation_enabled = false
			repository_connector = "account.${harness_platform_connector_github.test.id}"

			connector {
				connector_ref = "account.${harness_platform_connector_vault.test.id}"
				type = "vault"
			}

			connector {
				connector_ref = "account.${harness_platform_connector_awscc.test.id}"
				type = "aws"
			}

			connector {
				connector_ref = "account.${harness_platform_connector_azure_cloud_cost.test.id}"
				type = "azure"
			}

			connector {
				connector_ref = "account.${harness_platform_connector_gcp_cloud_cost.test.id}"
				type = "gcp"
			}

			tags = ["vault-test", "multi-cloud", "all-connectors"]
		}
	`, id, name)
}

// Helper function: Workspace without any connectors
func testAccResourceWorkspaceWithoutConnectors(id string, name string) string {
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
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
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

		resource "harness_platform_workspace" "test_vault_removal" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "test workspace without connectors"
			provisioner_type = "terraform"
			provisioner_version = "1.5.7"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "/"
			cost_estimation_enabled = false
			repository_connector = "account.${harness_platform_connector_github.test.id}"

			tags = ["vault-test", "removal-test"]
		}
	`, id, name)
}

// Helper function: Workspace with duplicate Vault connectors (should fail)
func testAccResourceWorkspaceWithDuplicateVault(id string, name string) string {
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
			secret_manager_identifier = "harnessSecretManager"
			value_type = "Inline"
			value = "secret"
		}

		resource "harness_platform_connector_github" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
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

		resource "harness_platform_connector_vault" "test_1" {
			identifier = "vault_1_%[1]s"
			name = "vault_1_%[2]s"
			description = "test vault connector 1"
			vault_url = "https://vault1.example.com"
			base_path = "secret"
			default = false
			delegate_selectors = ["harness-delegate"]
			auth_token {
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		resource "harness_platform_connector_vault" "test_2" {
			identifier = "vault_2_%[1]s"
			name = "vault_2_%[2]s"
			description = "test vault connector 2"
			vault_url = "https://vault2.example.com"
			base_path = "secret"
			default = false
			delegate_selectors = ["harness-delegate"]
			auth_token {
				token_ref = "account.${harness_platform_secret_text.test.id}"
			}
		}

		resource "harness_platform_workspace" "test_vault_duplicate" {
			identifier = "%[1]s"
			name = "%[2]s"
			org_id = harness_platform_organization.test.id
			project_id = harness_platform_project.test.id
			description = "test workspace with duplicate vault connectors"
			provisioner_type = "terraform"
			provisioner_version = "1.5.7"
			repository = "https://github.com/org/repo"
			repository_branch = "main"
			repository_path = "/"
			cost_estimation_enabled = false
			repository_connector = "account.${harness_platform_connector_github.test.id}"

			# First Vault connector
			connector {
				connector_ref = "account.${harness_platform_connector_vault.test_1.id}"
				type = "vault"
			}

			# Second Vault connector - THIS SHOULD FAIL
			connector {
				connector_ref = "account.${harness_platform_connector_vault.test_2.id}"
				type = "vault"
			}

			tags = ["vault-test", "duplicate-error-test"]
		}
	`, id, name)
}
