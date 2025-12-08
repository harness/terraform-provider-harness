package api

import (
	"fmt"
)

// WorkspacesService handles communication with the workspaces related
// methods of the Split.io APIv2.
type WorkspacesService service

// Workspace represents a workspace in Split.io
type Workspace struct {
	ID          *string `json:"id"`
	Name        *string `json:"name"`
	Type        *string `json:"type"`
	DisplayName *string `json:"displayName"`
}

// Get fetches a workspace by ID
func (w *WorkspacesService) Get(workspaceID string) (*Workspace, error) {
	var result Workspace
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/workspaces/%s", workspaceID)
	err := w.client.get(url, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindByName finds a workspace by name
func (w *WorkspacesService) FindByName(name string) (*Workspace, error) {
	workspaces, err := w.List()
	if err != nil {
		return nil, err
	}

	for _, workspace := range workspaces {
		if workspace.Name != nil && *workspace.Name == name {
			return workspace, nil
		}
	}

	return nil, fmt.Errorf("workspace with name '%s' not found", name)
}

// WorkspaceListResponse represents the response from the workspaces list endpoint
type WorkspaceListResponse struct {
	Objects []*Workspace `json:"objects"`
}

// WorkspaceListResult represents the response from the workspaces list endpoint with pagination
type WorkspaceListResult struct {
	GenericListResult
	Objects []*Workspace `json:"objects"`
}

// List returns all workspaces (deprecated - use ListAll for consistency)
func (w *WorkspacesService) List() ([]*Workspace, error) {
	return w.ListAll()
}

// ListPaginated returns workspaces with optional pagination
func (w *WorkspacesService) ListPaginated(opts *GenericListQueryParams) (*WorkspaceListResult, error) {
	var result WorkspaceListResult
	baseURL := "https://api.split.io/internal/api/v2/workspaces"

	// Enforce max limit of 100
	if opts != nil && opts.Limit > 100 {
		opts.Limit = 100
	}

	finalURL := w.client.buildURL(baseURL, opts)
	err := w.client.get(finalURL, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListAll returns all workspaces by handling pagination automatically
func (w *WorkspacesService) ListAll() ([]*Workspace, error) {
	var allWorkspaces []*Workspace
	offset := 0
	limit := 100 // Use max limit

	for {
		opts := &GenericListQueryParams{
			Offset: offset,
			Limit:  limit,
		}

		result, err := w.ListPaginated(opts)
		if err != nil {
			return nil, err
		}

		allWorkspaces = append(allWorkspaces, result.Objects...)

		// If we got fewer results than the limit, we've reached the end
		if len(result.Objects) < limit {
			break
		}

		// If we have totalCount and we've retrieved all items, break
		if result.TotalCount != nil && len(allWorkspaces) >= *result.TotalCount {
			break
		}

		offset += limit
	}

	return allWorkspaces, nil
}
