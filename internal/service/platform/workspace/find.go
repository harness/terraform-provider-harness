package workspace

import (
	"context"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
)

const (
	defaultLimit = int32(100)
	maxPages     = int32(1000)
)

// List all workspaces across pages.
func findWorkspaces(ctx context.Context, orgID, projectID, accountID string, workspace *nextgen.WorkspacesApiService, searchTerm string, limit int32) ([]nextgen.IacmWorkspaceResourceSummary, *http.Response, error) {
	if limit <= 0 {
		limit = defaultLimit
	}

	page := int32(1)
	var (
		all      []nextgen.IacmWorkspaceResourceSummary
		httpResp *http.Response
	)

	for {
		opts := &nextgen.WorkspacesApiWorkspacesListWorkspacesOpts{
			Limit: optional.NewInt32(limit),
			Page:  optional.NewInt32(page),
		}
		if searchTerm != "" {
			opts.SearchTerm = optional.NewString(searchTerm)
		}

		workspaces, resp, err := workspace.WorkspacesListWorkspaces(ctx, orgID, projectID, accountID, opts)
		httpResp = resp
		if err != nil {
			return nil, httpResp, fmt.Errorf("list workspaces page %d: %w", page, err)
		}
		all = append(all, workspaces...)

		if len(workspaces) < int(limit) {
			break
		}

		page++

		if page > maxPages {
			return nil, httpResp, fmt.Errorf("list workspaces exceeded max pages (%d)", maxPages)
		}
	}
	return all, httpResp, nil
}
