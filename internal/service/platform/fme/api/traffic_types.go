package api

import (
	"fmt"
)

// TrafficTypesService handles operations related to Split traffic types
type TrafficTypesService service

// TrafficTypeResponse represents a traffic type response
type TrafficTypeResponse struct {
	ID                  *string `json:"id"`
	Name                *string `json:"name"`
	Type                *string `json:"type"`
	DisplayAttributeID  *string `json:"displayAttributeId"`
}

// TrafficTypeCreateRequest represents the request body for creating a traffic type
type TrafficTypeCreateRequest struct {
	Name string `json:"name"`
}

// Create creates a new traffic type in the specified workspace
func (t *TrafficTypesService) Create(workspaceID string, req *TrafficTypeCreateRequest) (*TrafficTypeResponse, error) {
	var result TrafficTypeResponse
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/trafficTypes/ws/%s", workspaceID)
	err := t.client.post(url, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindByID retrieves a traffic type by workspace ID and traffic type ID
func (t *TrafficTypesService) FindByID(workspaceID, trafficTypeID string) (*TrafficTypeResponse, error) {
	trafficTypes, err := t.List(workspaceID)
	if err != nil {
		return nil, err
	}

	for _, trafficType := range trafficTypes {
		if trafficType.ID != nil && *trafficType.ID == trafficTypeID {
			return trafficType, nil
		}
	}

	return nil, fmt.Errorf("traffic type with ID '%s' not found in workspace '%s'", trafficTypeID, workspaceID)
}

// FindByName retrieves a traffic type by workspace ID and traffic type name
func (t *TrafficTypesService) FindByName(workspaceID, trafficTypeName string) (*TrafficTypeResponse, error) {
	trafficTypes, err := t.List(workspaceID)
	if err != nil {
		return nil, err
	}

	for _, trafficType := range trafficTypes {
		if trafficType.Name != nil && *trafficType.Name == trafficTypeName {
			return trafficType, nil
		}
	}

	return nil, fmt.Errorf("traffic type with name '%s' not found in workspace '%s'", trafficTypeName, workspaceID)
}

// Delete removes a traffic type by workspace ID and traffic type ID
func (t *TrafficTypesService) Delete(workspaceID, trafficTypeID string) error {
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/trafficTypes/%s", trafficTypeID)
	err := t.client.delete(url)
	if err != nil {
		return err
	}
	return nil
}

// List retrieves all traffic types for a workspace
func (t *TrafficTypesService) List(workspaceID string) ([]*TrafficTypeResponse, error) {
	var result []*TrafficTypeResponse
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/trafficTypes/ws/%s", workspaceID)
	err := t.client.get(url, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}