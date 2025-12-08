package api

import (
	"fmt"
)

// FlagSetsService handles communication with the flag sets related
// methods of the Split.io APIv2.
type FlagSetsService service

// FlagSet represents a feature flag set in Split.io
type FlagSet struct {
	ID          *string         `json:"id"`
	Name        *string         `json:"name"`
	Description *string         `json:"description"`
	Workspace   *WorkspaceIDRef `json:"workspace"`
	CreatedAt   *string         `json:"createdAt"`
	Type        *string         `json:"type"`
}

// WorkspaceIDRef represents a workspace reference
type WorkspaceIDRef struct {
	Type *string `json:"type"`
	ID   *string `json:"id"`
}

// FlagSetRequest represents a request to create/update a flag set
type FlagSetRequest struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Workspace   *WorkspaceIDRef `json:"workspace,omitempty"`
}

// Create creates a new flag set
func (f *FlagSetsService) Create(opts *FlagSetRequest) (*FlagSet, error) {
	var result FlagSet
	// Flag sets use different API version
	err := f.client.post("https://api.split.io/api/v3/flag-sets", opts, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get fetches a flag set by ID
func (f *FlagSetsService) Get(id string) (*FlagSet, error) {
	var result FlagSet
	url := fmt.Sprintf("https://api.split.io/api/v3/flag-sets/%s", id)
	err := f.client.get(url, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update modifies an existing flag set
func (f *FlagSetsService) Update(id string, opts *FlagSetRequest) (*FlagSet, error) {
	var result FlagSet
	url := fmt.Sprintf("https://api.split.io/api/v3/flag-sets/%s", id)
	err := f.client.patch(url, opts, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a flag set
func (f *FlagSetsService) Delete(id string) error {
	url := fmt.Sprintf("https://api.split.io/api/v3/flag-sets/%s", id)
	err := f.client.delete(url)
	if err != nil {
		return err
	}
	return nil
}

// FlagSetListResult represents a paginated response from the flag sets API
type FlagSetListResult struct {
	Data           []*FlagSet `json:"data"`
	NextMarker     *string    `json:"nextMarker"`
	PreviousMarker *string    `json:"previousMarker"`
	Limit          *int       `json:"limit"`
	Count          *int       `json:"count"`
}

// List returns all flag sets for a workspace
func (f *FlagSetsService) List(workspaceID string) ([]*FlagSet, error) {
	var allFlagSets []*FlagSet
	var nextMarker *string

	for {
		url := fmt.Sprintf("https://api.split.io/api/v3/flag-sets?workspace_id=%s&limit=200", workspaceID)
		if nextMarker != nil {
			url += fmt.Sprintf("&marker=%s", *nextMarker)
		}

		var result FlagSetListResult
		err := f.client.get(url, &result)
		if err != nil {
			return nil, err
		}

		allFlagSets = append(allFlagSets, result.Data...)

		if result.NextMarker == nil {
			break
		}
		nextMarker = result.NextMarker
	}

	return allFlagSets, nil
}

// FindByName retrieves a flag set by workspace ID and flag set name
func (f *FlagSetsService) FindByName(workspaceID, flagSetName string) (*FlagSet, error) {
	flagSets, err := f.List(workspaceID)
	if err != nil {
		return nil, fmt.Errorf("failed to list flag sets: %v", err)
	}

	for _, flagSet := range flagSets {
		if flagSet.Name != nil && *flagSet.Name == flagSetName {
			return flagSet, nil
		}
	}

	return nil, fmt.Errorf("flag set with name '%s' not found in workspace '%s'", flagSetName, workspaceID)
}