package chaos_test

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test configuration
var (
	accountID       = flag.String("account", "", "Harness account ID")
	apiKey          = flag.String("api-key", "", "Harness API key")
	orgIdentifier   = "<replace-with-your-org-identifier>"
	projectID       = "<replace-with-your-project-identifier>"
	envID           = "<replace-with-your-env-identifier>"
	connectorRef    = "<replace-with-your-connector-ref>"
	infraNamespace  = "<replace-with-your-infra-namespace>"
	serviceAccount  = "<replace-with-your-service-account>"
	testInfraPrefix = "test-sdk"
)

// TestMain handles setup and teardown
func TestMain(m *testing.M) {
	flag.Parse()
	if *accountID == "" || *apiKey == "" {
		fmt.Println("Skipping tests: account ID or API key not provided")
		os.Exit(0)
	}
	os.Exit(m.Run())
}

// setupClient initializes and returns a new API client with context
func setupClient() (*chaos.APIClient, context.Context) {
	cfg := chaos.NewConfiguration()
	cfg.BasePath = "https://app.harness.io/gateway/chaos/manager/api"
	client := chaos.NewAPIClient(cfg)
	client.AccountId = *accountID
	client.ApiKey = *apiKey

	ctx := context.WithValue(context.Background(), chaos.ContextAPIKey, chaos.APIKey{
		Key: client.ApiKey,
	})

	return client, ctx
}

// TestInfrastructureCRUD tests the complete CRUD lifecycle of an infrastructure
func TestInfrastructureCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client, ctx := setupClient()

	t.Run("Create and Get Infrastructure", func(t *testing.T) {
		t.Parallel()
		testCreateAndGetInfrastructure(t, client, ctx)
	})

	t.Run("List Infrastructures", func(t *testing.T) {
		t.Parallel()
		testListInfrastructures(t, client, ctx)
	})

	t.Run("Delete Infrastructure", func(t *testing.T) {
		t.Parallel()
		testDeleteInfrastructure(t, client, ctx)
	})
}

// testCreateAndGetInfrastructure tests infrastructure creation and retrieval
func testCreateAndGetInfrastructure(t *testing.T, client *chaos.APIClient, ctx context.Context) {
	infraName := generateTestName("infra")
	infraID := generateTestName("infra")

	t.Logf("Creating test infrastructure: %s", infraName)

	// Create platform infrastructure first
	platformInfraID := createPlatformInfrastructure(t, client, ctx, infraName, infraID)
	if platformInfraID == "" {
		t.Fatal("Failed to create platform infrastructure")
	}

	t.Logf("Created platform infrastructure with ID: %s", platformInfraID)

	// Create chaos infrastructure with reference to the platform infrastructure
	createReq := buildCreateRequest(infraName, platformInfraID)
	t.Logf("Creating chaos infrastructure with ID: %s", createReq.InfraID)
	createdInfra, httpResp, err := client.ChaosSdkApi.RegisterInfraV2(
		ctx,
		createReq,
		*accountID,
		orgIdentifier,
		projectID,
		nil,
	)
	require.NoError(t, err, "Failed to create infrastructure")
	require.NotNil(t, httpResp, "Response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected status code 200")
	assert.Equal(t, platformInfraID, createdInfra.Identity, "Infrastructure ID should match")

	// Test getting the created infrastructure
	infra, _, err := client.ChaosSdkApi.GetInfraV2(
		ctx,
		createdInfra.Identity,
		*accountID,
		orgIdentifier,
		projectID,
		envID,
	)
	require.NoError(t, err, "Failed to get infrastructure")
	assert.Equal(t, platformInfraID, infra.InfraID, "Infrastructure ID should match")
	assert.Equal(t, infraName, infra.Name, "Infrastructure name should match")
}

// testListInfrastructures tests listing infrastructures
func testListInfrastructures(t *testing.T, client *chaos.APIClient, ctx context.Context) {
	// Create a test infrastructure first
	infraName := generateTestName("list-test")
	infraID := generateTestName("list-test")

	// Create platform infrastructure first
	platformInfraID := createPlatformInfrastructure(t, client, ctx, infraName, infraID)
	if platformInfraID == "" {
		t.Fatal("Failed to create platform infrastructure")
	}

	t.Logf("Created platform infrastructure with ID: %s", platformInfraID)

	// Create chaos infrastructure with reference to the platform infrastructure
	createReq := buildCreateRequest(infraName, platformInfraID)
	t.Logf("Creating chaos infrastructure with ID: %s", createReq.InfraID)
	_, _, err := client.ChaosSdkApi.RegisterInfraV2(
		ctx,
		createReq,
		*accountID,
		orgIdentifier,
		projectID,
		nil,
	)
	require.NoError(t, err, "Failed to create test infrastructure")

	// List infrastructures
	listResp, resp, err := client.ChaosSdkApi.ListInfraV2(
		ctx,
		chaos.InfraV2ListKubernetesInfraV2Request{},
		*accountID,
		orgIdentifier,
		projectID,
		0,
		10,
		&chaos.ChaosSdkApiListInfraV2Opts{
			EnvironmentIdentifier: optional.NewString(envID),
		},
	)
	require.NoError(t, err, "Failed to list infrastructures")
	assert.NotNil(t, resp, "Response should not be nil")
	assert.Greater(t, len(listResp.Infras), 0, "Should return at least one infrastructure")
}

// testUpdateInfrastructure tests updating an infrastructure
func testUpdateInfrastructure(t *testing.T, client *chaos.APIClient, ctx context.Context) {
	// Create a test infrastructure first
	infraName := generateTestName("update-test")
	infraID := generateTestName("update-test")

	// Create platform infrastructure first
	platformInfraID := createPlatformInfrastructure(t, client, ctx, infraName, infraID)
	if platformInfraID == "" {
		t.Fatal("Failed to create platform infrastructure")
	}

	t.Logf("Created platform infrastructure with ID: %s", platformInfraID)

	// Create chaos infrastructure with reference to the platform infrastructure
	createReq := buildCreateRequest(infraName, platformInfraID)
	t.Logf("Creating chaos infrastructure with ID: %s", createReq.InfraID)
	createdInfra, _, err := client.ChaosSdkApi.RegisterInfraV2(
		ctx,
		createReq,
		*accountID,
		orgIdentifier,
		projectID,
		nil,
	)
	require.NoError(t, err, "Failed to create test infrastructure")

	// Update the infrastructure
	updatedName := infraName + "-updated"
	updateReq := chaos.InfraV2UpdateKubernetesInfrastructureV2Request{
		Name:          updatedName,
		Description:   "Updated description",
		Tags:          []string{"updated-tag"},
		Identity:      createdInfra.Identity,
		EnvironmentID: envID,
	}

	t.Logf("Updating infrastructure with request: %+v", updateReq)

	updateResp, httpResp, err := client.ChaosSdkApi.UpdateInfraV2(
		ctx,
		updateReq,
		*accountID,
		orgIdentifier,
		projectID,
		nil,
	)

	if err != nil {
		t.Fatalf("Update failed with error: %v", err)
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode >= 400 {
		body, readErr := io.ReadAll(httpResp.Body)
		if readErr != nil {
			t.Logf("Failed to read response body: %v", readErr)
		} else {
			t.Logf("Response body: %s", string(body))
		}
		t.Logf("Response status: %d %s", httpResp.StatusCode, httpResp.Status)
		t.Logf("Response headers: %+v", httpResp.Header)
		t.Fatalf("Update infrastructure failed with status code: %d", httpResp.StatusCode)
	}

	require.NoError(t, err, "Failed to update infrastructure")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected status code 200")
	assert.NotNil(t, updateResp, "Update response should not be nil")

	// Fetch the updated infrastructure to verify changes
	updatedInfra, _, err := client.ChaosSdkApi.GetInfraV2(
		ctx,
		createdInfra.Identity,
		*accountID,
		orgIdentifier,
		projectID,
		envID,
	)
	require.NoError(t, err, "Failed to fetch updated infrastructure")
	assert.Equal(t, updatedName, updatedInfra.Name, "Name should be updated")
}

// testDeleteInfrastructure tests deleting an infrastructure
func testDeleteInfrastructure(t *testing.T, client *chaos.APIClient, ctx context.Context) {
	// Create a test infrastructure first
	infraName := generateTestName("delete-test")
	infraID := generateTestName("delete-test")

	// Create platform infrastructure first
	platformInfraID := createPlatformInfrastructure(t, client, ctx, infraName, infraID)
	if platformInfraID == "" {
		t.Fatal("Failed to create platform infrastructure")
	}

	t.Logf("Created platform infrastructure with ID: %s", platformInfraID)

	// Create chaos infrastructure with reference to the platform infrastructure
	createReq := buildCreateRequest(infraName, platformInfraID)
	t.Logf("Creating chaos infrastructure with ID: %s", createReq.InfraID)
	createdInfra, _, err := client.ChaosSdkApi.RegisterInfraV2(
		ctx,
		createReq,
		*accountID,
		orgIdentifier,
		projectID,
		nil,
	)
	require.NoError(t, err, "Failed to create test infrastructure")

	// Delete the infrastructure
	deleteResp, httpResp, err := client.ChaosSdkApi.DeleteInfraV2(
		ctx,
		createdInfra.Identity,
		envID,
		*accountID,
		orgIdentifier,
		projectID,
	)
	require.NoError(t, err, "Failed to delete infrastructure")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected status code 200")
	assert.NotNil(t, deleteResp, "Delete response should not be nil")

	// Verify the infrastructure is deleted
	_, _, err = client.ChaosSdkApi.GetInfraV2(
		ctx,
		createdInfra.Identity,
		*accountID,
		orgIdentifier,
		projectID,
		envID,
	)
	assert.Error(t, err, "Should return error for deleted infrastructure")
}

// createPlatformInfrastructure creates a platform infrastructure using the NextGen API and returns its ID
func createPlatformInfrastructure(t *testing.T, client *chaos.APIClient, ctx context.Context, name, id string) string {
	t.Logf("Creating platform infrastructure: %s", name)

	// Create a valid identifier (only alphanumeric and underscores, starting with a letter)
	validID := "infra_" + strings.ReplaceAll(id, "-", "_")
	platformID := validID + "_platform"
	platformName := name + " platform"

	t.Logf("Account ID: %s, Org: %s, Project: %s, Env: %s",
		*accountID, orgIdentifier, projectID, envID)

	// Create the YAML configuration for the platform infrastructure
	yamlConfig := fmt.Sprintf(`infrastructureDefinition:
  name: "%s"
  identifier: "%s"
  orgIdentifier: %s
  projectIdentifier: %s
  environmentRef: %s
  deploymentType: Kubernetes
  type: KubernetesDirect
  spec:
    connectorRef: %s
    namespace: %s
    releaseName: release-<+INFRA_KEY>
  allowSimultaneousDeployments: false`,
		platformName,
		platformID,
		orgIdentifier,
		projectID,
		envID,
		connectorRef,
		infraNamespace,
	)

	t.Logf("YAML Configuration:\n%s", yamlConfig)

	// Create the NextGen client
	nextgenCfg := nextgen.NewConfiguration()
	nextgenCfg.BasePath = "https://app.harness.io/gateway"
	nextgenClient := nextgen.NewAPIClient(nextgenCfg)

	// Create context with API key
	nextgenCtx := context.WithValue(context.Background(), nextgen.ContextAPIKey, nextgen.APIKey{
		Key: *apiKey,
	})

	// Create the infrastructure request
	req := nextgen.InfrastructureRequest{
		Identifier:        platformID,
		OrgIdentifier:     orgIdentifier,
		ProjectIdentifier: projectID,
		EnvironmentRef:    envID,
		Name:              platformName,
		Description:       "Platform infrastructure for testing",
		Tags: map[string]string{
			"platform": "true",
		},
		Type_: "KubernetesDirect",
		Yaml:  yamlConfig,
	}

	t.Logf("Sending request to create infrastructure: %+v", req)

	// Call the NextGen API to create the infrastructure
	resp, httpResp, err := nextgenClient.InfrastructuresApi.CreateInfrastructure(
		nextgenCtx,
		*accountID,
		&nextgen.InfrastructuresApiCreateInfrastructureOpts{
			Body: optional.NewInterface(req),
		},
	)

	if err != nil {
		t.Logf("Error creating platform infrastructure: %v", err)
		if httpResp != nil {
			body, readErr := io.ReadAll(httpResp.Body)
			if readErr == nil {
				t.Logf("Response body: %s", string(body))
			}
			headers := ""
			for k, v := range httpResp.Header {
				headers += fmt.Sprintf("  %s: %v\n", k, v)
			}
			t.Logf("Response status: %d %s", httpResp.StatusCode, httpResp.Status)
			t.Logf("Response headers:\n%s", headers)
		}
		t.Fatal("Skipping test as platform infrastructure creation failed")
		return ""
	}

	if resp.Data == nil || resp.Data.Infrastructure == nil {
		t.Fatal("Failed to get infrastructure ID from response")
		return ""
	}

	t.Logf("Successfully created platform infrastructure with ID: %s", resp.Data.Infrastructure.Identifier)
	return resp.Data.Infrastructure.Identifier
}

// buildCreateRequest creates a standard infrastructure create request
// id is the unique identifier for the infrastructure
func buildCreateRequest(name, id string) chaos.InfraV2RegisterInfrastructureV2Request {
	infraType := chaos.KUBERNETES_InfraV2InfraType
	scope := chaos.CLUSTER_InfraV2InfraScope

	req := chaos.InfraV2RegisterInfrastructureV2Request{
		Identifier: &chaos.InfraV2Identifiers{
			AccountIdentifier: *accountID,
			OrgIdentifier:     orgIdentifier,
			ProjectIdentifier: projectID,
		},
		Identity:       id,
		InfraID:        id,
		Name:           name,
		Description:    "Test infrastructure created by SDK test",
		Tags:           []string{"test", "sdk-test"},
		InfraType:      &infraType,
		InfraScope:     &scope,
		InfraNamespace: infraNamespace,
		ServiceAccount: serviceAccount,
		EnvironmentID:  envID,
	}

	return req
}

// generateTestName generates a unique test name with timestamp
func generateTestName(prefix string) string {
	return fmt.Sprintf("%s%s-%d", testInfraPrefix, prefix, time.Now().UnixNano())
}
