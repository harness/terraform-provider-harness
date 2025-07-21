package chaos

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
)

// ChaosHubClient provides methods to interact with the Chaos Hub API
type ChaosHubClient struct {
	client *APIClient
}

// NewChaosHubClient creates a new Chaos Hub client
func NewChaosHubClient(client *APIClient) *ChaosHubClient {
	return &ChaosHubClient{client: client}
}

// Get retrieves a Chaos Hub configuration
func (c *ChaosHubClient) Get(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	hubID string,
) (*model.ChaosHubStatus, error) {
	query := `
        query GetChaosHub($identifiers: IdentifiersRequest!, $chaosHubID: String!) {
            getChaosHub(identifiers: $identifiers, chaosHubID: $chaosHubID) {
                id
                repoName
                repoURL
                repoBranch
                connectorId
                connectorScope
                AuthType
                isAvailable
                totalFaults
                totalExperiments
                name
                lastSyncedAt
                isDefault
                tags
                createdBy {
                    userID
                    username
                    email
                }
                updatedBy {
                    userID
                    username
                    email
                }
                createdAt
                updatedAt
                description
            }
        }`

	variables := map[string]interface{}{
		"identifiers": identifiers,
		"chaosHubID":  hubID,
	}

	var response struct {
		GetChaosHub *model.ChaosHubStatus `json:"getChaosHub"`
	}

	if err := c.client.ExecuteGraphQL(ctx, query, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to get chaos hub: %w", err)
	}

	return response.GetChaosHub, nil
}

// Context key for storing identifiers
const ContextKeyIdentifiers = "identifiers"

// Create creates a new Chaos Hub configuration
func (c *ChaosHubClient) Create(
	ctx context.Context,
	hubName string,
	repoBranch string,
	connectorID string,
	identifiers model.IdentifiersRequest,
	opts ...func(*model.ChaosHubRequest) *model.ChaosHubRequest,
) (*model.ChaosHub, error) {
	req := &model.ChaosHubRequest{
		HubName:        hubName,
		RepoBranch:     repoBranch,
		ConnectorID:    connectorID,
		ConnectorScope: model.ConnectorScopeProject, // Default to project scope
	}

	// Apply any additional options
	for _, opt := range opts {
		req = opt(req)
	}

	mutation := `
        mutation AddChaosHub($request: ChaosHubRequest!, $identifiers: IdentifiersRequest!) {
            addChaosHub(request: $request, identifiers: $identifiers) {
                id
                identifiers {
                    accountIdentifier
                    orgIdentifier
                    projectIdentifier
                }
                repoName
                repoURL
                repoBranch
                AuthType
                connectorId
                connectorScope
                name
                createdAt
                updatedAt
                lastSyncedAt
                isDefault
                tags
                createdBy {
                    userID
                    username
                    email
                }
                updatedBy {
                    userID
                    username
                    email
                }
                description
            }
        }`

	variables := map[string]interface{}{
		"request":     req,
		"identifiers": identifiers,
	}

	var response struct {
		AddChaosHub *model.ChaosHub `json:"addChaosHub"`
	}

	if err := c.client.ExecuteGraphQL(ctx, mutation, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to create chaos hub: %w", err)
	}

	return response.AddChaosHub, nil
}

// Update updates an existing Chaos Hub configuration
func (c *ChaosHubClient) Update(
	ctx context.Context,
	hubID string,
	hubName string,
	repoBranch string,
	connectorID string,
	identifiers model.IdentifiersRequest,
	opts ...func(*model.ChaosHubRequest) *model.ChaosHubRequest,
) (*model.ChaosHub, error) {
	req := &model.ChaosHubRequest{
		HubName:       hubName,
		RepoBranch:    repoBranch,
		ConnectorID:   connectorID,
		ConnectorScope: model.ConnectorScopeProject, // Default to project scope
	}

	// Apply any additional options
	for _, opt := range opts {
		req = opt(req)
	}

	mutation := `
        mutation UpdateChaosHub($id: ID!, $request: ChaosHubRequest!, $identifiers: IdentifiersRequest!) {
            updateChaosHub(id: $id, request: $request, identifiers: $identifiers) {
                id
                repoName
                repoURL
                repoBranch
                connectorId
                connectorScope
                AuthType
                name
                lastSyncedAt
                isDefault
                tags
                createdBy {
                    userID
                    username
                    email
                }
                updatedBy {
                    userID
                    username
                    email
                }
                createdAt
                updatedAt
                description
            }
        }`

	variables := map[string]interface{}{
		"id":         hubID,
		"request":    req,
		"identifiers": identifiers,
	}

	var response struct {
		UpdateChaosHub *model.ChaosHub `json:"updateChaosHub"`
	}

	if err := c.client.ExecuteGraphQL(ctx, mutation, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to update chaos hub: %w", err)
	}

	return response.UpdateChaosHub, nil
}

// List retrieves a list of Chaos Hubs
func (c *ChaosHubClient) List(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
) ([]*model.ChaosHubStatus, error) {
	query := `
        query ListChaosHub($identifiers: IdentifiersRequest!, $request: ListChaosHubRequest) {
            listChaosHub(identifiers: $identifiers, request: $request) {
                id
                repoName
                repoURL
                repoBranch
                connectorId
                connectorScope
                AuthType
                isAvailable
                totalFaults
                totalExperiments
                name
                lastSyncedAt
                isDefault
                tags
                createdBy {
                    userID
                    username
                    email
                }
                updatedBy {
                    userID
                    username
                    email
                }
                createdAt
                updatedAt
                description
            }
        }`

	// Create a default request
	request := map[string]interface{}{
		// No filters by default
	}

	variables := map[string]interface{}{
		"identifiers": identifiers,
		"request":     request,
	}

	var response struct {
		ListChaosHub []*model.ChaosHubStatus `json:"listChaosHub"`
	}

	if err := c.client.ExecuteGraphQL(ctx, query, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to list chaos hubs: %w", err)
	}

	return response.ListChaosHub, nil
}

// Delete deletes a Chaos Hub configuration
func (c *ChaosHubClient) Delete(
	ctx context.Context,
	hubID string,
	identifiers model.IdentifiersRequest,
) (bool, error) {
	mutation := `
        mutation DeleteChaosHub($id: ID!, $identifiers: IdentifiersRequest!) {
            deleteChaosHub(id: $id, identifiers: $identifiers)
        }`

	variables := map[string]interface{}{
		"id":         hubID,
		"identifiers": identifiers,
	}

	// The response is a simple boolean, but we need to use a struct to unmarshal it
	var response struct {
		DeleteChaosHub bool `json:"deleteChaosHub"`
	}

	if err := c.client.ExecuteGraphQL(ctx, mutation, variables, &response); err != nil {
		// Provide more helpful error messages for common issues
		if err.Error() == "graphql: chaos hub not found" {
			return false, fmt.Errorf("chaos hub with ID %s not found", hubID)
		}
		return false, fmt.Errorf("failed to delete chaos hub %s: %w", hubID, err)
	}

	if !response.DeleteChaosHub {
		return false, fmt.Errorf("failed to delete chaos hub %s: operation returned false", hubID)
	}

	return true, nil
}

// Sync synchronizes a Chaos Hub with its remote repository
func (c *ChaosHubClient) Sync(
	ctx context.Context,
	hubID string,
	identifiers model.IdentifiersRequest,
) (string, error) {
	mutation := `
        mutation SyncChaosHub($id: ID!, $identifiers: IdentifiersRequest!) {
            syncChaosHub(id: $id, identifiers: $identifiers)
        }`

	variables := map[string]interface{}{
		"id":         hubID,
		"identifiers": identifiers,
	}

	// The response is a simple string, but we need to use a struct to unmarshal it
	var response struct {
		SyncChaosHub string `json:"syncChaosHub"`
	}

	if err := c.client.ExecuteGraphQL(ctx, mutation, variables, &response); err != nil {
		return "", fmt.Errorf("failed to sync chaos hub: %w", err)
	}

	return response.SyncChaosHub, nil
}

// Note: PushWorkflowToChaosHub and PushProbeToChaosHub methods are not implemented at this time
