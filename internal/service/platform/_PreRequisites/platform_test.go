package _PreRequisites

import (
	"testing"

	_ "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	_ "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceSecretText_Create(t *testing.T) {
	//resourceName := "harness_platform_secret_text.test"
	resourceData := schema.TestResourceDataRaw(t, resourceSecretText().Schema, map[string]interface{}{
		"identifier":                "test_secret",
		"name":                      "Test Secret",
		"description":               "A test secret",
		"secret_manager_identifier": "harnessSecretManager",
		"value_type":                "Inline",
		"value":                     "super-secret-value",
	})

	// Call create method
	err := resourceSecretTextCreate(resourceData, nil)
	if err != nil {
		t.Fatalf("Error creating secret text: %s", err)
	}

	// Verify created resource
	if id := resourceData.Id(); id != "test_secret" {
		t.Errorf("Expected ID to be 'test_secret', got: %s", id)
	}

	if err := resourceSecretTextRead(resourceData, nil); err != nil {
		t.Fatalf("Error reading secret text: %s", err)
	}
}

func TestAccResourceProject_Create(t *testing.T) {
	//resourceName := "harness_platform_project.test"
	resourceData := schema.TestResourceDataRaw(t, resourceProject().Schema, map[string]interface{}{
		"identifier": "test_project",
		"name":       "Test Project",
		"org_id":     "default",
		"color":      "#0063F7",
	})

	// Call create method
	err := resourceProjectCreate(resourceData, nil)
	if err != nil {
		t.Fatalf("Error creating project: %s", err)
	}

	// Verify created resource
	if id := resourceData.Id(); id != "test_project" {
		t.Errorf("Expected ID to be 'test_project', got: %s", id)
	}

	if err := resourceProjectRead(resourceData, nil); err != nil {
		t.Fatalf("Error reading project: %s", err)
	}
}

func TestAccResourceAzureKeyVaultConnector_Create(t *testing.T) {
	//resourceName := "harness_platform_connector_azure_key_vault.test"
	resourceData := schema.TestResourceDataRaw(t, resourceAzureKeyVaultConnector().Schema, map[string]interface{}{
		"identifier":             "test_azure_connector",
		"name":                   "Test Azure Connector",
		"description":            "A test Azure Key Vault Connector",
		"client_id":              "test-client-id",
		"secret_key":             "test_secret_key",
		"tenant_id":              "test-tenant-id",
		"vault_name":             "test-vault",
		"subscription":           "test-subscription",
		"is_default":             false,
		"azure_environment_type": "AZURE",
	})

	// Call create method
	err := resourceAzureKeyVaultConnectorCreate(resourceData, nil)
	if err != nil {
		t.Fatalf("Error creating Azure Key Vault connector: %s", err)
	}

	// Verify created resource
	if id := resourceData.Id(); id != "test_azure_connector" {
		t.Errorf("Expected ID to be 'test_azure_connector', got: %s", id)
	}

	if err := resourceAzureKeyVaultConnectorRead(resourceData, nil); err != nil {
		t.Fatalf("Error reading Azure Key Vault connector: %s", err)
	}
}

func TestAccResourceAwsSecretManagerConnector_Create(t *testing.T) {
	//resourceName := "harness_platform_connector_aws_secret_manager.test"
	resourceData := schema.TestResourceDataRaw(t, resourceAwsSecretManagerConnector().Schema, map[string]interface{}{
		"identifier":     "test_aws_connector",
		"name":           "Test AWS Connector",
		"description":    "A test AWS Secret Manager Connector",
		"region":         "us-east-1",
		"secret_key_ref": "account.test_secret_key",
		"access_key_ref": "account.test_access_key",
	})

	// Call create method
	err := resourceAwsSecretManagerConnectorCreate(resourceData, nil)
	if err != nil {
		t.Fatalf("Error creating AWS Secret Manager connector: %s", err)
	}

	// Verify created resource
	if id := resourceData.Id(); id != "test_aws_connector" {
		t.Errorf("Expected ID to be 'test_aws_connector', got: %s", id)
	}

	if err := resourceAwsSecretManagerConnectorRead(resourceData, nil); err != nil {
		t.Fatalf("Error reading AWS Secret Manager connector: %s", err)
	}
}
