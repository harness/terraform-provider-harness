package chaos

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
)

// SecurityGovernanceConditionClient provides methods to interact with the Security Governance Condition API
type SecurityGovernanceConditionClient struct {
	client *APIClient
}

// NewSecurityGovernanceConditionClient creates a new Security Governance Condition client
func NewSecurityGovernanceConditionClient(client *APIClient) *SecurityGovernanceConditionClient {
	return &SecurityGovernanceConditionClient{client: client}
}

// Get retrieves a security governance condition by ID
func (c *SecurityGovernanceConditionClient) Get(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	conditionID string,
) (*model.ConditionResponse, error) {
	query := `
        query GetCondition($identifiers: IdentifiersRequest!, $conditionId: String!) {
            getCondition(identifiers: $identifiers, conditionId: $conditionId) {
                identifiers {
                    accountIdentifier
                    orgIdentifier
                    projectIdentifier
                }
                condition {
                    name
                    description
                    tags
                    conditionId
                    infraType
                    faultSpec {
                        operator
                        faults {
                            faultType
                            name
                        }
                    }
                    k8sSpec {
                        infraSpec {
                            operator
                            infraIds
                        }
                        applicationSpec {
                            operator
                            workloads {
                                label
                                namespace
                                kind
                                services
                                applicationMapId
                                env {
                                    name
                                    value
                                }
                            }
                        }
                        chaosServiceAccountSpec {
                            operator
                            serviceAccounts
                        }
                    }
                    machineSpec {
                        infraSpec {
                            operator
                            infraIds
                        }
                    }
                }
                createdAt
                updatedAt
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
            }
        }`

	variables := map[string]interface{}{
		"identifiers": identifiers,
		"conditionId":  conditionID,
	}

	var response struct {
		GetCondition *model.ConditionResponse `json:"getCondition"`
	}

	if err := c.client.ExecuteGraphQL(ctx, query, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to get security governance condition: %w", err)
	}

	return response.GetCondition, nil
}

// List retrieves a list of security governance conditions with optional filtering and pagination
func (c *SecurityGovernanceConditionClient) List(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	request model.ListConditionRequest,
) (*model.ListConditionResponse, error) {
	query := `
        query ListCondition($identifiers: IdentifiersRequest!, $request: ListConditionRequest!) {
            listCondition(identifiers: $identifiers, request: $request) {
                totalConditions
                conditions {
                    identifiers {
                        accountIdentifier
                        orgIdentifier
                        projectIdentifier
                    }
                    condition {
                        name
                        description
                        tags
                        conditionId
                        infraType
                    }
                    createdAt
                    updatedAt
                }
            }
        }`

	variables := map[string]interface{}{
		"identifiers": identifiers,
		"request":     request,
	}

	var response struct {
		ListCondition *model.ListConditionResponse `json:"listCondition"`
	}

	if err := c.client.ExecuteGraphQL(ctx, query, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to list security governance conditions: %w", err)
	}

	return response.ListCondition, nil
}

// Create creates a new security governance condition
func (c *SecurityGovernanceConditionClient) Create(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	request model.ConditionRequest,
) (*model.StandardResponse, error) {
	mutation := `
        mutation AddCondition($identifiers: IdentifiersRequest!, $request: ConditionRequest!) {
            addCondition(identifiers: $identifiers, request: $request) {
                message
                correlationId
                response
            }
        }`

	variables := map[string]interface{}{
		"identifiers": identifiers,
		"request":     request,
	}

	var response struct {
		AddCondition *model.StandardResponse `json:"addCondition"`
	}

	if err := c.client.ExecuteGraphQL(ctx, mutation, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to create security governance condition: %w", err)
	}

	return response.AddCondition, nil
}

// Update updates an existing security governance condition
func (c *SecurityGovernanceConditionClient) Update(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	request model.ConditionRequest,
) (*model.StandardResponse, error) {
	mutation := `
        mutation UpdateCondition($identifiers: IdentifiersRequest!, $request: ConditionRequest!) {
            updateCondition(identifiers: $identifiers, request: $request) {
                message
                correlationId
                response
            }
        }`

	variables := map[string]interface{}{
		"identifiers": identifiers,
		"request":     request,
	}

	var response struct {
		UpdateCondition *model.StandardResponse `json:"updateCondition"`
	}

	if err := c.client.ExecuteGraphQL(ctx, mutation, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to update security governance condition: %w", err)
	}

	return response.UpdateCondition, nil
}

// Delete deletes a security governance condition by ID
func (c *SecurityGovernanceConditionClient) Delete(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	conditionID string,
) (*model.StandardResponse, error) {
	mutation := `
        mutation DeleteCondition($identifiers: IdentifiersRequest!, $conditionId: String!) {
            deleteCondition(identifiers: $identifiers, conditionId: $conditionId) {
                message
                correlationId
                response
            }
        }`

	variables := map[string]interface{}{
		"identifiers": identifiers,
		"conditionId":  conditionID,
	}

	var response struct {
		DeleteCondition *model.StandardResponse `json:"deleteCondition"`
	}

	if err := c.client.ExecuteGraphQL(ctx, mutation, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to delete security governance condition: %w", err)
	}

	return response.DeleteCondition, nil
}
