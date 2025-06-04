package svcdiscovery_test

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
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test configuration
var (
	accountID          = flag.String("account", "", "Harness account ID")
	apiKey             = flag.String("api-key", "", "Harness API key")
	orgIdentifier      = "<your-org-identifier>"
	projectID          = "<your-project-identifier>"
	envID              = "<your-env-identifier>"
	connectorRef       = "<your-connector-identifier>"
	namespace          = "<your-namespace>"
	observedNamespaces = []string{"<your-namespace-to-observe>"} // List of namespaces to observe
	testAgentPrefix    = "test-sdk"
	cron               = svcdiscovery.DatabaseCronConfig{
		Expression: "15 * * * *",
	}
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
func setupClient() (*svcdiscovery.APIClient, context.Context) {
	cfg := svcdiscovery.NewConfiguration()
	cfg.BasePath = "https://app.harness.io/gateway/servicediscovery"
	client := svcdiscovery.NewAPIClient(cfg)
	client.AccountId = *accountID
	client.ApiKey = *apiKey

	ctx := context.WithValue(context.Background(), svcdiscovery.ContextAPIKey, svcdiscovery.APIKey{
		Key: client.ApiKey,
	})

	return client, ctx
}

// TestAgentCRUD tests the complete CRUD lifecycle of an agent
func TestAgentCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client, ctx := setupClient()

	// Create platform infrastructure first
	infraName := generateTestName("test")
	infraID := generateTestName("test")
	platformInfraID := createPlatformInfrastructure(t, infraName, infraID)
	if platformInfraID == "" {
		t.Fatal("Failed to create platform infrastructure")
	}
	t.Logf("Created platform infrastructure with ID: %s", platformInfraID)

	// Run tests sequentially
	t.Run("1. Create and Get Agent", func(t *testing.T) {
		// Create and get the agent
		createdAgent := testCreateAndGetAgent(t, client, ctx, platformInfraID)

		// Test listing the agent
		t.Run("2. List Agents", func(t *testing.T) {
			testListAgents(t, client, ctx, createdAgent)
		})

		// Test updating the agent
		t.Run("3. Update Agent", func(t *testing.T) {
			testUpdateAgent(t, client, ctx, createdAgent)
		})

		// Test application map operations
		t.Run("4. Application Map Operations", func(t *testing.T) {
			testCreateAndDeleteApplicationMap(t, client, ctx, createdAgent)
		})

		// Test deleting the agent
		t.Run("5. Delete Agent", func(t *testing.T) {
			testDeleteAgent(t, client, ctx, createdAgent)
		})
	})
}

// createPlatformInfrastructure creates a platform infrastructure using the NextGen API and returns its ID
func createPlatformInfrastructure(t *testing.T, name, id string) string {
	t.Logf("Creating platform infrastructure: %s", name)

	// Create a valid identifier (only alphanumeric and underscores, starting with a letter)
	validID := "infra_" + strings.ReplaceAll(id, "-", "_")
	platformID := validID + "_platform"
	platformName := name + " platform"

	// Use the configured namespace with a unique suffix for this test
	testNamespace := fmt.Sprintf("%s-%s", namespace, strings.ToLower(strings.ReplaceAll(id, "-", "")))

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
		testNamespace,
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
	infraReq := nextgen.InfrastructureRequest{
		Identifier:        platformID,
		Name:              platformName,
		OrgIdentifier:     orgIdentifier,
		ProjectIdentifier: projectID,
		Type_:             "KubernetesDirect",
		Yaml:              yamlConfig,
		Description:       "Platform infrastructure for testing service discovery",
		Tags: map[string]string{
			"platform": "true",
		},
	}

	t.Logf("Sending request to create infrastructure: %+v", infraReq)

	// Call the NextGen API to create the infrastructure
	resp, httpResp, err := nextgenClient.InfrastructuresApi.CreateInfrastructure(
		nextgenCtx,
		*accountID,
		&nextgen.InfrastructuresApiCreateInfrastructureOpts{
			Body: optional.NewInterface(infraReq),
		},
	)

	if err != nil {
		t.Logf("Error creating platform infrastructure: %v", err)
		if httpResp != nil {
			body, readErr := io.ReadAll(httpResp.Body)
			httpResp.Body.Close()
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

	// Ensure response body is closed if we got a response
	if httpResp != nil {
		httpResp.Body.Close()
	}

	if resp.Data == nil || resp.Data.Infrastructure == nil {
		t.Fatal("Failed to get infrastructure ID from response")
		return ""
	}

	t.Logf("Successfully created platform infrastructure with ID: %s", resp.Data.Infrastructure.Identifier)
	return resp.Data.Infrastructure.Identifier
}

// testCreateAndGetAgent tests agent creation and retrieval
func testCreateAndGetAgent(t *testing.T, client *svcdiscovery.APIClient, ctx context.Context, platformInfraID string) *svcdiscovery.ApiGetAgentResponse {
	agentName := generateTestName("agent")

	t.Logf("Creating test agent: %s with infrastructure: %s", agentName, platformInfraID)

	// Create agent request
	createReq := svcdiscovery.ApiCreateAgentRequest{
		Name:                  agentName,
		EnvironmentIdentifier: envID,
		InfraIdentifier:       platformInfraID, // Use the provided platform infrastructure ID
		Config: &svcdiscovery.DatabaseAgentConfiguration{
			Data: &svcdiscovery.DatabaseDataCollectionConfiguration{
				EnableNodeAgent:       true,
				ObservedNamespaces:    observedNamespaces,
				CollectionWindowInMin: 5,
				Cron:                  &cron,
			},
			Kubernetes: &svcdiscovery.DatabaseKubernetesAgentConfiguration{
				Namespace: fmt.Sprintf("%s", namespace),
				// ServiceAccount:  "harness-agent-sa",
				ImagePullPolicy: "IfNotPresent",
				// RunAsUser:       1000,
				// RunAsGroup:      1000,
				Labels: map[string]string{
					"app": fmt.Sprintf("%s", namespace),
				},
			},
		},
	}

	// Create the agent
	createdAgent, httpResp, err := client.AgentApi.CreateAgent(
		ctx,
		createReq,
		*accountID,
		&svcdiscovery.AgentApiCreateAgentOpts{
			OrganizationIdentifier: optional.NewString(orgIdentifier),
			ProjectIdentifier:      optional.NewString(projectID),
		},
	)

	// Ensure response body is closed if we got a response
	if httpResp != nil {
		defer httpResp.Body.Close()
	}

	if err != nil {
		t.Logf("Error creating agent: %v", err)
		if httpResp != nil {
			body, _ := io.ReadAll(httpResp.Body)
			t.Logf("Response status: %s", httpResp.Status)
			t.Logf("Response body: %s", string(body))
			httpResp.Body.Close()
		}
		t.Fatal("Failed to create agent")
	}

	require.NotNil(t, httpResp, "Response should not be nil")
	if httpResp.StatusCode != 200 && httpResp.StatusCode != 201 {
		body, _ := io.ReadAll(httpResp.Body)
		t.Logf("Unexpected status code: %d", httpResp.StatusCode)
		t.Logf("Response body: %s", string(body))
		httpResp.Body.Close()
		t.Fatalf("Expected status code 200 or 201, got %d", httpResp.StatusCode)
	}

	// Ensure response body is closed
	defer httpResp.Body.Close()
	assert.NotEmpty(t, createdAgent.Identity, "Agent ID should not be empty")

	// Test getting the created agent
	agent, _, err := client.AgentApi.GetAgent(
		ctx,
		createdAgent.Identity,
		*accountID,
		envID,
		&svcdiscovery.AgentApiGetAgentOpts{
			OrganizationIdentifier: optional.NewString(orgIdentifier),
			ProjectIdentifier:      optional.NewString(projectID),
		},
	)

	require.NoError(t, err, "Failed to get agent")
	assert.Equal(t, agentName, agent.Name, "Agent name should match")
	// Note: The agent's description is not directly updatable through the update API

	return &createdAgent
}

// testListAgents tests listing agents
func testListAgents(t *testing.T, client *svcdiscovery.APIClient, ctx context.Context, expectedAgent *svcdiscovery.ApiGetAgentResponse) {
	// List agents
	listResp, httpResp, err := client.AgentApi.ListAgent(
		ctx,
		*accountID,
		envID,
		0,    // page
		100,  // limit
		true, // all
		&svcdiscovery.AgentApiListAgentOpts{
			OrganizationIdentifier: optional.NewString(orgIdentifier),
			ProjectIdentifier:      optional.NewString(projectID),
		},
	)

	// Ensure response body is closed
	if httpResp != nil {
		defer httpResp.Body.Close()
	}

	require.NoError(t, err, "Failed to list agents")
	require.NotNil(t, httpResp, "Response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected status code 200")
	if listResp.Items == nil {
		t.Fatal("Expected non-nil items in list response")
	}

	// Verify the created agent is in the list
	found := false
	for _, agent := range listResp.Items {
		if agent.Identity == expectedAgent.Identity {
			found = true
			assert.Equal(t, expectedAgent.Name, agent.Name, "Agent name should match")
			// Note: The InfraIdentifier might be in the Config field of the agent
			// If you need to verify infrastructure details, you would access them through agent.Config
			break
		}
	}
	assert.True(t, found, "Created agent should be in the list")
}

// testUpdateAgent tests updating an agent
func testUpdateAgent(t *testing.T, client *svcdiscovery.APIClient, ctx context.Context, agent *svcdiscovery.ApiGetAgentResponse) {
	// Update the agent using the provided agent
	updatedName := agent.Name + "-updated"

	// Get the current config to update it
	currentConfig := agent.Config
	if currentConfig == nil {
		currentConfig = &svcdiscovery.DatabaseAgentConfiguration{}
	}

	updateReq := svcdiscovery.ApiUpdateAgentRequest{
		Name:   updatedName,
		Config: currentConfig,
	}

	t.Logf("Updating agent with request: %+v", updateReq)

	updatedAgent, httpResp, err := client.AgentApi.UpdateAgent(
		ctx,
		updateReq,
		*accountID,
		envID,
		agent.Identity,
		&svcdiscovery.AgentApiUpdateAgentOpts{
			OrganizationIdentifier: optional.NewString(orgIdentifier),
			ProjectIdentifier:      optional.NewString(projectID),
		},
	)

	require.NoError(t, err, "Failed to update agent")
	require.NotNil(t, httpResp, "Response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected status code 200")
	assert.Equal(t, updatedName, updatedAgent.Name, "Agent name should be updated")
	// Note: The agent's description is not directly updatable through the update API
}

// testCreateAndDeleteApplicationMap tests creating and deleting an application map
func testCreateAndDeleteApplicationMap(t *testing.T, client *svcdiscovery.APIClient, ctx context.Context, agent *svcdiscovery.ApiGetAgentResponse) {
	appMapName := generateTestName("app-map")
	appMapID := "appmap_" + strings.ReplaceAll(appMapName, "-", "_")

	t.Logf("Waiting for discovered services to be available...")

	// Retry logic to wait for services to be discovered
	var discoveredServices *svcdiscovery.ApiListDiscoveredService
	maxRetries := 12 // 60 seconds total with 5 second intervals
	retryInterval := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		// List the discovered services using K8sresourceApiService
		services, httpResp, err := client.K8sresourceApi.ListDiscoveredService(
			ctx,
			agent.Identity,
			*accountID,
			envID,
			0,    // page
			10,   // limit
			true, // all
			&svcdiscovery.K8sresourceApiListDiscoveredServiceOpts{
				OrganizationIdentifier: optional.NewString(orgIdentifier),
				ProjectIdentifier:      optional.NewString(projectID),
			},
		)

		if httpResp != nil {
			httpResp.Body.Close()
		}

		if err != nil || httpResp == nil || httpResp.StatusCode != 200 {
			t.Logf("Failed to list discovered services (attempt %d/%d): %v", i+1, maxRetries, err)
			time.Sleep(retryInterval)
			continue
		}

		discoveredServices = &services

		// Log the current state of discovered services
		t.Logf("Discovered services: %d (attempt %d/%d)", len(discoveredServices.Items), i+1, maxRetries)
		for j, svc := range discoveredServices.Items {
			t.Logf("  %d. %s (ID: %s, Namespace: %s, Kind: %s)",
				j+1,
				svc.Name,
				svc.Id,
				svc.Spec.Kubernetes.Namespace,
				svc.Spec.Kubernetes.Kind)
		}

		// If we have at least 2 services, we can proceed
		if len(discoveredServices.Items) >= 2 {
			break
		}

		// If this is the last attempt and we still don't have enough services
		if i == maxRetries-1 {
			t.Logf("Timed out waiting for services after %d attempts", maxRetries)
			t.Skip("Need at least 2 discovered services to create an application map")
		}

		time.Sleep(retryInterval)
	}

	if discoveredServices == nil || len(discoveredServices.Items) < 2 {
		t.Fatal("Failed to discover services after multiple attempts")
	}

	t.Logf("Found %d discovered services, using first two for the application map", len(discoveredServices.Items))

	// Use the first two discovered services
	sourceService := discoveredServices.Items[0]
	targetService := discoveredServices.Items[1]

	t.Logf("Creating application map: %s with services %s -> %s", appMapName, sourceService.Name, targetService.Name)

	// Create application map request using discovered services
	createReq := svcdiscovery.ApiCreateNetworkMapRequest{
		Name:        appMapName,
		Identity:    appMapID,
		Description: "Test application map created by SDK test using discovered services",
		Resources: []svcdiscovery.DatabaseNetworkMapEntity{
			{
				Id:   sourceService.Id,
				Name: sourceService.Name,
				Kind: "KUBERNETES",
				Kubernetes: &svcdiscovery.DatabaseNetworkMapEntityKubernetesInfo{
					Namespace: sourceService.Spec.Kubernetes.Namespace,
					Kind:      sourceService.Spec.Kubernetes.Kind,
				},
			},
			{
				Id:   targetService.Id,
				Name: targetService.Name,
				Kind: "KUBERNETES",
				Kubernetes: &svcdiscovery.DatabaseNetworkMapEntityKubernetesInfo{
					Namespace: targetService.Spec.Kubernetes.Namespace,
					Kind:      targetService.Spec.Kubernetes.Kind,
				},
			},
		},
		Connections: []svcdiscovery.DatabaseConnection{
			{
				Type_: "HTTP",
				Port:  "8080",
				From: &svcdiscovery.DatabaseNetworkMapEntity{
					Id:   sourceService.Id,
					Name: sourceService.Name,
					Kind: "KUBERNETES",
				},
				To: &svcdiscovery.DatabaseNetworkMapEntity{
					Id:   targetService.Id,
					Name: targetService.Name,
					Kind: "KUBERNETES",
				},
			},
		},
	}

	// Create the application map
	createdAppMap, httpResp, err := client.ApplicationmapApi.CreateApplicationMap(
		ctx,
		createReq,
		*accountID,
		envID,
		agent.Identity,
		&svcdiscovery.ApplicationmapApiCreateApplicationMapOpts{
			OrganizationIdentifier: optional.NewString(orgIdentifier),
			ProjectIdentifier:      optional.NewString(projectID),
		},
	)

	// Ensure response body is closed
	if httpResp != nil {
		defer httpResp.Body.Close()
	}

	require.NoError(t, err, "Failed to create application map")
	require.NotNil(t, httpResp, "Response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected status code 200")
	assert.Equal(t, appMapName, createdAppMap.Name, "Application map name should match")

	// Verify the application map exists by fetching it
	fetchedAppMap, httpResp, err := client.ApplicationmapApi.GetApplicationMap(
		ctx,
		agent.Identity,
		appMapID,
		*accountID,
		envID,
		&svcdiscovery.ApplicationmapApiGetApplicationMapOpts{
			OrganizationIdentifier: optional.NewString(orgIdentifier),
			ProjectIdentifier:      optional.NewString(projectID),
		},
	)

	// Ensure response body is closed
	if httpResp != nil {
		defer httpResp.Body.Close()
	}

	require.NoError(t, err, "Failed to fetch created application map")
	require.NotNil(t, httpResp, "Response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected status code 200 when fetching application map")
	assert.Equal(t, appMapName, fetchedAppMap.Name, "Fetched application map name should match")
	assert.Equal(t, appMapID, fetchedAppMap.Identity, "Fetched application map ID should match")
	assert.NotEmpty(t, fetchedAppMap.CreatedAt, "CreatedAt should be set")
	assert.NotEmpty(t, fetchedAppMap.UpdatedAt, "UpdatedAt should be set")

	// Verify the resources and connections were created as expected
	assert.Len(t, fetchedAppMap.Resources, 2, "Should have two resources")
	assert.Len(t, fetchedAppMap.Connections, 1, "Should have one connection")
	assert.Equal(t, "HTTP", fetchedAppMap.Connections[0].Type_, "Connection type should be HTTP")

	// Verify the connection is between the expected services
	assert.Equal(t, sourceService.Id, fetchedAppMap.Connections[0].From.Id, "Connection source should match")
	assert.Equal(t, targetService.Id, fetchedAppMap.Connections[0].To.Id, "Connection target should match")

	// Test deleting the application map
	delResp, httpResp, err := client.ApplicationmapApi.DeleteApplicationMap(
		ctx,
		agent.Identity,
		appMapID,
		*accountID,
		envID,
		&svcdiscovery.ApplicationmapApiDeleteApplicationMapOpts{
			OrganizationIdentifier: optional.NewString(orgIdentifier),
			ProjectIdentifier:      optional.NewString(projectID),
		},
	)

	// Ensure response body is closed
	if httpResp != nil {
		defer httpResp.Body.Close()
	}

	require.NoError(t, err, "Failed to delete application map")
	require.NotNil(t, httpResp, "Response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected status code 200")
	assert.NotNil(t, delResp, "Delete response should not be nil")

	// Verify the application map is deleted
	_, httpResp, err = client.ApplicationmapApi.GetApplicationMap(
		ctx,
		agent.Identity,
		appMapID,
		*accountID,
		envID,
		&svcdiscovery.ApplicationmapApiGetApplicationMapOpts{
			OrganizationIdentifier: optional.NewString(orgIdentifier),
			ProjectIdentifier:      optional.NewString(projectID),
		},
	)

	assert.Error(t, err, "Expected error when getting deleted application map")
	if httpResp != nil {
		defer httpResp.Body.Close()
		assert.Equal(t, 500, httpResp.StatusCode, "Expected status code 500 for deleted application map")
	}
}

// testDeleteAgent tests deleting an agent
func testDeleteAgent(t *testing.T, client *svcdiscovery.APIClient, ctx context.Context, agent *svcdiscovery.ApiGetAgentResponse) {

	// Delete the agent
	delResp, httpResp, err := client.AgentApi.DeleteAgent(
		ctx,
		agent.Identity,
		*accountID,
		envID,
		&svcdiscovery.AgentApiDeleteAgentOpts{
			OrganizationIdentifier: optional.NewString(orgIdentifier),
			ProjectIdentifier:      optional.NewString(projectID),
		},
	)

	require.NoError(t, err, "Failed to delete agent")
	require.NotNil(t, httpResp, "Response should not be nil")
	assert.Equal(t, 200, httpResp.StatusCode, "Expected status code 200")
	assert.NotNil(t, delResp, "Delete response should not be nil")

	// Verify the agent is deleted
	_, _, err = client.AgentApi.GetAgent(
		ctx,
		agent.Identity,
		*accountID,
		envID,
		&svcdiscovery.AgentApiGetAgentOpts{
			OrganizationIdentifier: optional.NewString(orgIdentifier),
			ProjectIdentifier:      optional.NewString(projectID),
		},
	)

	// Should return an error for deleted agent
	assert.Error(t, err, "Should return error for deleted agent")
}

// generateTestName generates a unique test name with timestamp
func generateTestName(prefix string) string {
	timestamp := time.Now().UnixNano()
	// Remove special characters and ensure it starts with a letter
	safePrefix := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return '-'
	}, prefix)
	return fmt.Sprintf("%s-%s-%d", testAgentPrefix, safePrefix, timestamp)
}
