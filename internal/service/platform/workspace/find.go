package workspace

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
)

const defaultLimit = int32(100)

// List all workspaces across pages.
func findWorkspaces(ctx context.Context, orgID, projectID, accountID string, workspace *nextgen.WorkspacesApiService, searchTerm string) ([]nextgen.IacmWorkspaceResourceSummary, *http.Response, error) {
	page := int32(1)
	var (
		all      []nextgen.IacmWorkspaceResourceSummary
		httpResp *http.Response
	)

	for {
		opts := &nextgen.WorkspacesApiWorkspacesListWorkspacesOpts{
			Limit:      optional.NewInt32(defaultLimit),
			Page:       optional.NewInt32(page),
			SearchTerm: optional.NewString(searchTerm),
		}

		workspaces, resp, err := workspace.WorkspacesListWorkspaces(ctx, orgID, projectID, accountID, opts)
		httpResp = resp
		if err != nil {
			return nil, httpResp, err
		}
		all = append(all, workspaces...)
		if int32(len(workspaces)) < defaultLimit {
			break
		}
		page++
	}
	return all, httpResp, nil
}
