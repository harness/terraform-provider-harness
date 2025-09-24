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

		fmt.Printf("DEBUG: Making request to: %s\n", url)

		var result FlagSetListResult
		err := f.client.get(url, &result)
		if err != nil {
			fmt.Printf("DEBUG: API request error: %v\n", err)
			return nil, err
		}

		fmt.Printf("DEBUG: API response successful, got %d flag sets in this batch\n", len(result.Data))
		for i, fs := range result.Data {
			fmt.Printf("DEBUG: Batch flag set %d: ID=%s, Name=%s\n", i, getStringValue(fs.ID), getStringValue(fs.Name))
		}

		allFlagSets = append(allFlagSets, result.Data...)

		if result.NextMarker == nil {
			fmt.Printf("DEBUG: No more pages, total flag sets: %d\n", len(allFlagSets))
			break
		}
		nextMarker = result.NextMarker
		fmt.Printf("DEBUG: Next marker: %s\n", *nextMarker)
	}

	return allFlagSets, nil
}

// FindByName retrieves a flag set by workspace ID and flag set name
func (f *FlagSetsService) FindByName(workspaceID, flagSetName string) (*FlagSet, error) {
	flagSets, err := f.List(workspaceID)
	if err != nil {
		return nil, fmt.Errorf("failed to list flag sets: %v", err)
	}

	// Debug logging - print what we actually got back
	fmt.Printf("DEBUG: Found %d flag sets in workspace %s\n", len(flagSets), workspaceID)
	for i, flagSet := range flagSets {
		var name string
		if flagSet.Name != nil {
			name = *flagSet.Name
		} else {
			name = "<nil>"
		}
		fmt.Printf("DEBUG: Flag set %d: ID=%s, Name=%s\n", i, getStringValue(flagSet.ID), name)
	}
	fmt.Printf("DEBUG: Looking for flag set with name: '%s'\n", flagSetName)

	for _, flagSet := range flagSets {
		if flagSet.Name != nil && *flagSet.Name == flagSetName {
			fmt.Printf("DEBUG: Found matching flag set: %s\n", *flagSet.Name)
			return flagSet, nil
		}
	}

	return nil, fmt.Errorf("flag set with name '%s' not found in workspace '%s'", flagSetName, workspaceID)
}

// Helper function to safely get string value
func getStringValue(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}