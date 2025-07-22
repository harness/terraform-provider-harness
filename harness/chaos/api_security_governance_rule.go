package chaos

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
)

// SecurityGovernanceRuleClient provides methods to interact with the Security Governance Rule API
type SecurityGovernanceRuleClient struct {
	client *APIClient
}

// NewSecurityGovernanceRuleClient creates a new Security Governance Rule client
func NewSecurityGovernanceRuleClient(client *APIClient) *SecurityGovernanceRuleClient {
	return &SecurityGovernanceRuleClient{client: client}
}

// Get retrieves a security governance rule by ID
func (c *SecurityGovernanceRuleClient) Get(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	ruleID string,
) (*model.RuleResponse, error) {
	query := `
        query GetRule($identifiers: IdentifiersRequest!, $ruleId: String!) {
            getRule(identifiers: $identifiers, ruleId: $ruleId) {
                identifiers {
                    accountIdentifier
                    orgIdentifier
                    projectIdentifier
                }
                rule {
                    name
                    description
                    tags
                    ruleId
                    isEnabled
                    userGroupIds
                    timeWindows {
                        duration
                        endTime
                        startTime
                        timeZone
                        recurrence {
                            type
                            spec {
                                until
                                value
                            }
                        }
                    }
                    conditions {
                        name
                        description
                        conditionId
                        infraType
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
		"ruleId":      ruleID,
	}

	var response struct {
		GetRule *model.RuleResponse `json:"getRule"`
	}

	if err := c.client.ExecuteGraphQL(ctx, query, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to get security governance rule: %w", err)
	}

	return response.GetRule, nil
}

// List retrieves a list of security governance rules with optional filtering and pagination
func (c *SecurityGovernanceRuleClient) List(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	request model.ListRuleRequest,
) ([]*model.RuleResponse, error) {
	query := `
        query ListRule($identifiers: IdentifiersRequest!, $request: ListRuleRequest!) {
            listRule(identifiers: $identifiers, request: $request) {
                identifiers {
                    accountIdentifier
                    orgIdentifier
                    projectIdentifier
                }
                rule {
                    name
                    description
                    tags
                    ruleId
                    isEnabled
                    userGroupIds
                    conditions {
                        name
                        conditionId
                    }
                }
                createdAt
                updatedAt
            }
        }`

	variables := map[string]interface{}{
		"identifiers": identifiers,
		"request":     request,
	}

	var response struct {
		ListRule []*model.RuleResponse `json:"listRule"`
	}

	if err := c.client.ExecuteGraphQL(ctx, query, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to list security governance rules: %w", err)
	}

	return response.ListRule, nil
}

// Create creates a new security governance rule
func (c *SecurityGovernanceRuleClient) Create(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	request model.RuleInput,
) (*model.StandardResponse, error) {
	mutation := `
        mutation AddRule($identifiers: IdentifiersRequest!, $request: RuleInput!) {
            addRule(identifiers: $identifiers, request: $request) {
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
		AddRule *model.StandardResponse `json:"addRule"`
	}

	if err := c.client.ExecuteGraphQL(ctx, mutation, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to create security governance rule: %w", err)
	}

	return response.AddRule, nil
}

// Update updates an existing security governance rule
func (c *SecurityGovernanceRuleClient) Update(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	request model.RuleInput,
) (*model.StandardResponse, error) {
	mutation := `
        mutation UpdateRule($identifiers: IdentifiersRequest!, $request: RuleInput!) {
            updateRule(identifiers: $identifiers, request: $request) {
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
		UpdateRule *model.StandardResponse `json:"updateRule"`
	}

	if err := c.client.ExecuteGraphQL(ctx, mutation, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to update security governance rule: %w", err)
	}

	return response.UpdateRule, nil
}

// Delete deletes a security governance rule by ID
func (c *SecurityGovernanceRuleClient) Delete(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	ruleID string,
) (*model.StandardResponse, error) {
	mutation := `
        mutation DeleteRule($identifiers: IdentifiersRequest!, $ruleId: String!) {
            deleteRule(identifiers: $identifiers, ruleId: $ruleId) {
                message
                correlationId
                response
            }
        }`

	variables := map[string]interface{}{
		"identifiers": identifiers,
		"ruleId":      ruleID,
	}

	var response struct {
		DeleteRule *model.StandardResponse `json:"deleteRule"`
	}

	if err := c.client.ExecuteGraphQL(ctx, mutation, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to delete security governance rule: %w", err)
	}

	return response.DeleteRule, nil
}

// Tune enables or disables a security governance rule
func (c *SecurityGovernanceRuleClient) Tune(
	ctx context.Context,
	identifiers model.IdentifiersRequest,
	ruleID string,
	enable bool,
) (*model.StandardResponse, error) {
	mutation := `
        mutation TuneRule($identifiers: IdentifiersRequest!, $ruleId: String!, $enable: Boolean!) {
            tuneRule(identifiers: $identifiers, ruleId: $ruleId, enable: $enable) {
                message
                correlationId
                response
            }
        }`

	variables := map[string]interface{}{
		"identifiers": identifiers,
		"ruleId":      ruleID,
		"enable":      enable,
	}

	var response struct {
		TuneRule *model.StandardResponse `json:"tuneRule"`
	}

	if err := c.client.ExecuteGraphQL(ctx, mutation, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to tune security governance rule: %w", err)
	}

	return response.TuneRule, nil
}
