package api

import (
	"fmt"
)

// EnvironmentSegmentKeysService handles operations related to Split environment segment keys
type EnvironmentSegmentKeysService service

// EnvironmentSegmentKeys represents environment segment keys
type EnvironmentSegmentKeys struct {
	SegmentName    *string   `json:"segmentName"`
	EnvironmentID  *string   `json:"environmentId"`
	Keys           []string  `json:"keys"`
}

// EnvironmentSegmentKeysCreateRequest represents the request for creating environment segment keys
type EnvironmentSegmentKeysCreateRequest struct {
	Keys []string `json:"keys"`
}

// EnvironmentSegmentKeysUpdateRequest represents the request for updating environment segment keys
type EnvironmentSegmentKeysUpdateRequest struct {
	Keys []string `json:"keys"`
}

// EnvironmentSegmentKeysListResult represents the response from the segment keys list endpoint with pagination
type EnvironmentSegmentKeysListResult struct {
	Keys                  []SegmentKeyObject `json:"keys"`
	OpenChangeRequestID   *string            `json:"openChangeRequestId"`
	Count                 int                `json:"count"`
	Offset                int                `json:"offset"`
	Limit                 int                `json:"limit"`
}

// Create creates new environment segment keys
func (e *EnvironmentSegmentKeysService) Create(workspaceID, segmentName, environmentID string, req *EnvironmentSegmentKeysCreateRequest) (*EnvironmentSegmentKeys, error) {
	var result EnvironmentSegmentKeys
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/ws/%s/%s/environments/%s/keys", workspaceID, segmentName, environmentID)
	err := e.client.put(url, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves environment segment keys
func (e *EnvironmentSegmentKeysService) Get(workspaceID, segmentName, environmentID string) (*EnvironmentSegmentKeys, error) {
	var result EnvironmentSegmentKeys
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/ws/%s/%s/environments/%s/keys", workspaceID, segmentName, environmentID)
	err := e.client.get(url, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies existing environment segment keys
func (e *EnvironmentSegmentKeysService) Update(workspaceID, segmentName, environmentID string, req *EnvironmentSegmentKeysUpdateRequest) (*EnvironmentSegmentKeys, error) {
	var result EnvironmentSegmentKeys
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/ws/%s/%s/environments/%s/keys", workspaceID, segmentName, environmentID)
	err := e.client.put(url, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes environment segment keys
func (e *EnvironmentSegmentKeysService) Delete(workspaceID, segmentName, environmentID string) error {
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/ws/%s/%s/environments/%s/keys", workspaceID, segmentName, environmentID)
	err := e.client.delete(url)
	if err != nil {
		return err
	}
	return nil
}

// List returns segment environment keys with optional pagination
func (e *EnvironmentSegmentKeysService) List(workspaceID, segmentName, environmentID string, opts *GenericListQueryParams) (*EnvironmentSegmentKeysListResult, error) {
	var result EnvironmentSegmentKeysListResult
	baseURL := fmt.Sprintf("https://api.split.io/internal/api/v2/segments/ws/%s/%s/environments/%s/keys", workspaceID, segmentName, environmentID)

	// Enforce max limit of 100
	if opts != nil && opts.Limit > 100 {
		opts.Limit = 100
	}

	finalURL := e.client.buildURL(baseURL, opts)
	err := e.client.get(finalURL, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListAll returns all segment environment keys by handling pagination automatically
func (e *EnvironmentSegmentKeysService) ListAll(workspaceID, segmentName, environmentID string) ([]string, error) {
	var allKeys []string
	offset := 0
	limit := 100 // Use max limit

	for {
		opts := &GenericListQueryParams{
			Offset: offset,
			Limit:  limit,
		}

		result, err := e.List(workspaceID, segmentName, environmentID, opts)
		if err != nil {
			return nil, err
		}

		// Convert SegmentKeyObject array to string array
		for _, keyObj := range result.Keys {
			allKeys = append(allKeys, keyObj.Key)
		}

		// If we got fewer results than the limit, we've reached the end
		if len(result.Keys) < limit {
			break
		}

		// If we have count and we've retrieved all items, break
		if len(allKeys) >= result.Count {
			break
		}

		offset += limit
	}

	return allKeys, nil
}