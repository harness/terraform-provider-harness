package api

import (
	"fmt"
	"net/url"
)

// TrafficTypeAttributesService handles operations related to Split traffic type attributes
type TrafficTypeAttributesService service

// TrafficTypeAttribute represents a traffic type attribute
type TrafficTypeAttribute struct {
	ID                  *string `json:"id"`
	TrafficTypeID       *string `json:"trafficTypeId"`
	DisplayName         *string `json:"displayName"`
	Description         *string `json:"description"`
	DataType            *string `json:"dataType"`
	IsSearchable        *bool   `json:"isSearchable"`
	SuggestedValuesFor  *string `json:"suggestedValuesFor"`
	SuggestedValuesJSON *string `json:"suggestedValuesJSON"`
	OrganizationID      *string `json:"organizationId"`
}

// TrafficTypeAttributeCreateRequest represents the request body for creating a traffic type attribute
type TrafficTypeAttributeCreateRequest struct {
	ID              string   `json:"id"`
	DisplayName     string   `json:"displayName"`
	Description     *string  `json:"description,omitempty"`
	DataType        string   `json:"dataType"`
	IsSearchable    *bool    `json:"isSearchable,omitempty"`
	SuggestedValues []string `json:"suggestedValues,omitempty"`
}

// TrafficTypeAttributeUpdateRequest represents the request body for updating a traffic type attribute
type TrafficTypeAttributeUpdateRequest struct {
	DisplayName         string  `json:"displayName"`
	Description         *string `json:"description,omitempty"`
	DataType            string  `json:"dataType"`
	IsSearchable        *bool   `json:"isSearchable,omitempty"`
	SuggestedValuesFor  *string `json:"suggestedValuesFor,omitempty"`
	SuggestedValuesJSON *string `json:"suggestedValuesJSON,omitempty"`
}

// Create creates a new traffic type attribute in the specified workspace and traffic type
func (t *TrafficTypeAttributesService) Create(workspaceID, trafficTypeID string, req *TrafficTypeAttributeCreateRequest) (*TrafficTypeAttribute, error) {
	var result TrafficTypeAttribute
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/schema/ws/%s/trafficTypes/%s", workspaceID, trafficTypeID)
	err := t.client.post(url, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a traffic type attribute by workspace ID, traffic type ID and attribute ID
func (t *TrafficTypeAttributesService) Get(workspaceID, trafficTypeID, attributeID string) (*TrafficTypeAttribute, error) {
	attributes, err := t.List(workspaceID, trafficTypeID)
	if err != nil {
		return nil, err
	}

	for _, attribute := range attributes {
		if attribute.ID != nil && *attribute.ID == attributeID {
			return attribute, nil
		}
	}

	return nil, fmt.Errorf("traffic type attribute with ID '%s' not found in workspace '%s' traffic type '%s'", attributeID, workspaceID, trafficTypeID)
}

// Update modifies an existing traffic type attribute
func (t *TrafficTypeAttributesService) Update(workspaceID, attributeID string, req *TrafficTypeAttributeUpdateRequest) (*TrafficTypeAttribute, error) {
	var result TrafficTypeAttribute
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/trafficTypeAttributes/%s", attributeID)
	err := t.client.patch(url, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a traffic type attribute by workspace ID, traffic type ID and attribute ID
func (t *TrafficTypeAttributesService) Delete(workspaceID, trafficTypeID, attributeID string) error {
	// URL-encode the attributeID to handle special characters
	escapedAttributeID := url.QueryEscape(attributeID)
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/schema/ws/%s/trafficTypes/%s/%s", workspaceID, trafficTypeID, escapedAttributeID)
	err := t.client.delete(url)
	if err != nil {
		return err
	}
	return nil
}

// List retrieves all traffic type attributes for a workspace and traffic type
func (t *TrafficTypeAttributesService) List(workspaceID, trafficTypeID string) ([]*TrafficTypeAttribute, error) {
	var result []*TrafficTypeAttribute
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/schema/ws/%s/trafficTypes/%s", workspaceID, trafficTypeID)
	err := t.client.get(url, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}