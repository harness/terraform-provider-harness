package api

import (
	"fmt"
)

// SplitsService handles communication with the splits (feature flags) related
// methods of the Split.io APIv2.
type SplitsService service

// Split represents a feature flag in Split.io
type Split struct {
	ID                     *string             `json:"id"`
	Name                   *string             `json:"name"`
	Description            *string             `json:"description"`
	Tags                   []SplitTag          `json:"tags,omitempty"`
	CreationTime           *int64              `json:"creationTime"`
	RolloutStatusTimestamp *int64              `json:"rolloutStatusTimestamp"`
	TrafficType            *TrafficType        `json:"trafficType"`
	RolloutStatus          *SplitRolloutStatus `json:"rolloutStatus"`
}

// SplitRolloutStatus represents the rollout status of a split
type SplitRolloutStatus struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

// SplitTag represents a tag on a split
type SplitTag struct {
	Name string `json:"name"`
}

// SplitCreateRequest represents a request to create a split
type SplitCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// SplitUpdateRequest represents a request to update a split
type SplitUpdateRequest struct {
	Description string `json:"description"`
}

// TrafficType represents a traffic type
type TrafficType struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

// SplitListResult represents the response from the splits list endpoint with pagination
type SplitListResult struct {
	GenericListResult
	Objects []*Split `json:"objects"`
}

// User represents a user
type User struct {
	ID    *string `json:"id"`
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

// Create creates a new split
func (s *SplitsService) Create(workspaceID, trafficTypeID string, opts *SplitCreateRequest) (*Split, error) {
	var result Split
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/splits/ws/%s/trafficTypes/%s", workspaceID, trafficTypeID)
	err := s.client.post(url, opts, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get fetches a split by ID
func (s *SplitsService) Get(workspaceID, splitID string) (*Split, error) {
	var result Split
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/splits/ws/%s/%s", workspaceID, splitID)
	err := s.client.get(url, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies an existing split description
func (s *SplitsService) Update(workspaceID, splitName string, opts *SplitUpdateRequest) (*Split, error) {
	var result Split
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/splits/ws/%s/%s/updateDescription", workspaceID, splitName)
	err := s.client.put(url, opts, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a split
func (s *SplitsService) Delete(workspaceID, splitName string) error {
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/splits/ws/%s/%s", workspaceID, splitName)
	err := s.client.delete(url)
	if err != nil {
		return err
	}
	return nil
}

// List returns splits for a workspace with optional pagination
func (s *SplitsService) List(workspaceID string, opts *GenericListQueryParams) (*SplitListResult, error) {
	var result SplitListResult
	baseURL := fmt.Sprintf("https://api.split.io/internal/api/v2/splits/ws/%s", workspaceID)
	finalURL := s.client.buildURL(baseURL, opts)
	err := s.client.get(finalURL, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListAll returns all splits for a workspace by handling pagination automatically
func (s *SplitsService) ListAll(workspaceID string) ([]*Split, error) {
	var allSplits []*Split
	offset := 0
	limit := 50 // Use a reasonable page size

	for {
		opts := &GenericListQueryParams{
			Offset: offset,
			Limit:  limit,
		}

		result, err := s.List(workspaceID, opts)
		if err != nil {
			return nil, err
		}

		allSplits = append(allSplits, result.Objects...)

		// If we got fewer results than the limit, we've reached the end
		if len(result.Objects) < limit {
			break
		}

		// If we have totalCount and we've retrieved all items, break
		if result.TotalCount != nil && len(allSplits) >= *result.TotalCount {
			break
		}

		offset += limit
	}

	return allSplits, nil
}

// ActivateInEnvironment activates a split in an environment
func (s *SplitsService) ActivateInEnvironment(workspaceID, splitName, environmentID string) error {
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/splits/%s/%s", environmentID, splitName)
	err := s.client.post(url, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
