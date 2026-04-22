package internal

import (
	"context"
	"sync"

	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/harness-go-sdk/harness/dbops"
	"github.com/harness/harness-go-sdk/harness/har"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/po"
	"github.com/harness/harness-go-sdk/harness/policymgmt"
	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/harness/harness-go-sdk/harness/svcdiscovery"
	openapi_client_nextgen "github.com/harness/harness-openapi-go-client/nextgen"
)

type Session struct {
	AccountId   string
	Endpoint    string
	CDClient    *cd.ApiClient
	PLClient    *nextgen.APIClient
	DBOpsClient *dbops.APIClient
	Client      *openapi_client_nextgen.APIClient
	CodeClient  *code.APIClient
	ChaosClient *chaos.APIClient
	SDClient    *svcdiscovery.APIClient
	SplitClient *split.APIClient
	// fmeWorkspaces caches Split workspace lookups by Harness org_id + project_id (see split.WorkspaceByOrganizationAndProject).
	fmeWorkspaceMu sync.Mutex
	fmeWorkspaces  map[string]split.Workspace
	HARClient      *har.APIClient
	IDPClient      *idp.APIClient
	POClient       *po.APIClient
}

func (s *Session) GetPlatformClient() (*nextgen.APIClient, context.Context) {
	return s.GetPlatformClientWithContext(nil)
}

func (s *Session) GetPlatformClientWithContext(ctx context.Context) (*nextgen.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	return s.PLClient.WithAuthContext(ctx)
}

func (s *Session) GetClientWithContext(ctx context.Context) (*openapi_client_nextgen.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	return s.Client.WithAuthContext(ctx)
}

func (s *Session) GetPolicyManagementClient() *policymgmt.APIClient {
	c := policymgmt.NewAPIClient(
		policymgmt.NewConfiguration(),
	)
	c.ChangeBasePath(s.Endpoint + "/pm")
	return c
}

func (s *Session) GetDBOpsClientWithContext(ctx context.Context) (*dbops.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	return s.DBOpsClient.WithAuthContext(ctx)
}

func (s *Session) GetCodeClientWithContext(ctx context.Context) (*code.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	return s.CodeClient.WithAuthContext(ctx)
}

func (s *Session) GetChaosClientWithContext(ctx context.Context) (*chaos.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	return s.ChaosClient.WithAuthContext(ctx)
}

func (s *Session) GetServiceDiscoveryClientWithContext(ctx context.Context) (*svcdiscovery.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	return s.SDClient.WithAuthContext(ctx)
}

func (s *Session) GetSplitClientWithContext(ctx context.Context) (*split.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	return s.SplitClient, ctx
}

func (s *Session) GetHarClientWithContext(ctx context.Context) (*har.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	return s.HARClient.WithAuthContext(ctx)
}

func (s *Session) GetIDPClientWithContext(ctx context.Context) (*idp.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	return s.IDPClient, ctx
}

func (s *Session) GetPOClientWithContext(ctx context.Context) (*po.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	return s.POClient, ctx
}

func fmeWorkspaceCacheKey(orgID, projectID string) string {
	return orgID + "\x00" + projectID
}

// GetFMEWorkspace returns a cached Split workspace for the Harness org/project pair, if present.
func (s *Session) GetFMEWorkspace(orgID, projectID string) (split.Workspace, bool) {
	if s == nil {
		return split.Workspace{}, false
	}
	s.fmeWorkspaceMu.Lock()
	defer s.fmeWorkspaceMu.Unlock()
	if s.fmeWorkspaces == nil {
		return split.Workspace{}, false
	}
	w, ok := s.fmeWorkspaces[fmeWorkspaceCacheKey(orgID, projectID)]
	return w, ok
}

// SetFMEWorkspace stores a Split workspace in the session cache for the Harness org/project pair.
func (s *Session) SetFMEWorkspace(orgID, projectID string, w split.Workspace) {
	if s == nil {
		return
	}
	s.fmeWorkspaceMu.Lock()
	defer s.fmeWorkspaceMu.Unlock()
	if s.fmeWorkspaces == nil {
		s.fmeWorkspaces = make(map[string]split.Workspace)
	}
	s.fmeWorkspaces[fmeWorkspaceCacheKey(orgID, projectID)] = w
}
