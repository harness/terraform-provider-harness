package split

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/split"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// SplitClientFromMeta returns the configured Split API client from provider meta or diagnostics.
func SplitClientFromMeta(ctx context.Context, meta interface{}) (*split.APIClient, diag.Diagnostics) {
	session, ok := meta.(*internal.Session)
	if !ok || session == nil {
		return nil, diag.FromErr(fmt.Errorf("invalid session: expected *internal.Session"))
	}
	client, _ := session.GetSplitClientWithContext(ctx)
	if client == nil {
		return nil, diag.FromErr(fmt.Errorf("Split client is not configured; ensure platform_api_key and account_id are set"))
	}
	return client, nil
}

func sessionFromMeta(meta interface{}) *internal.Session {
	s, ok := meta.(*internal.Session)
	if !ok {
		return nil
	}
	return s
}

// WorkspaceByOrganizationAndProject returns the unique workspace for Harness org and project identifiers.
// It uses the session workspace cache (when session is non-nil) and retries transient Split list errors (429, duplicate-key 500).
func WorkspaceByOrganizationAndProject(ctx context.Context, session *internal.Session, client *split.APIClient, orgID, projectID string) (split.Workspace, error) {
	if orgID == "" || projectID == "" {
		return split.Workspace{}, fmt.Errorf("organization and project identifiers are required")
	}
	if session != nil {
		if w, ok := session.GetFMEWorkspace(orgID, projectID); ok {
			return w, nil
		}
	}
	ws, err := findByOrganizationAndProjectWithRetry(ctx, client, orgID, projectID)
	if err != nil {
		return split.Workspace{}, err
	}
	if len(ws) == 0 {
		return split.Workspace{}, fmt.Errorf("no workspace found for organization %q project %q", orgID, projectID)
	}
	if len(ws) > 1 {
		return split.Workspace{}, fmt.Errorf("multiple workspaces (%d) found for organization %q project %q", len(ws), orgID, projectID)
	}
	out := ws[0]
	if session != nil {
		session.SetFMEWorkspace(orgID, projectID, out)
	}
	return out, nil
}

// WorkspaceByName returns the unique workspace whose name matches exactly (Split nameOp IS).
// When session is non-nil and the API returns organization and project identifiers, the workspace
// is stored in the session cache for later org_id/project_id resolutions in the same provider run.
func WorkspaceByName(ctx context.Context, session *internal.Session, client *split.APIClient, name string) (split.Workspace, error) {
	nameOp := "IS"
	ws, err := findByNameWithRetry(ctx, client, name, nameOp)
	if err != nil {
		return split.Workspace{}, err
	}
	if len(ws) == 0 {
		return split.Workspace{}, fmt.Errorf("no workspace found with name %q", name)
	}
	if len(ws) > 1 {
		return split.Workspace{}, fmt.Errorf("multiple workspaces (%d) found with name %q", len(ws), name)
	}
	out := ws[0]
	if session != nil && out.OrganizationIdentifier != "" && out.ProjectIdentifier != "" {
		session.SetFMEWorkspace(out.OrganizationIdentifier, out.ProjectIdentifier, out)
	}
	return out, nil
}

func resolveWorkspaceID(ctx context.Context, session *internal.Session, client *split.APIClient, orgID, projectID string) (string, error) {
	w, err := WorkspaceByOrganizationAndProject(ctx, session, client, orgID, projectID)
	if err != nil {
		return "", err
	}
	return w.ID, nil
}

// ResolveWorkspaceIDFromOrgProject returns the Split workspace ID for Harness org and project identifiers,
// using session workspace caching and list retries. Intended for import paths and internal helpers.
func ResolveWorkspaceIDFromOrgProject(ctx context.Context, meta interface{}, client *split.APIClient, orgID, projectID string) (string, error) {
	return resolveWorkspaceID(ctx, sessionFromMeta(meta), client, orgID, projectID)
}

// EnvironmentByOrganizationProjectAndName resolves the FME workspace, then returns the Split environment with the given name, or an error if not found.
func EnvironmentByOrganizationProjectAndName(ctx context.Context, session *internal.Session, client *split.APIClient, orgID, projectID, envName string) (*split.Environment, error) {
	if envName == "" {
		return nil, fmt.Errorf("environment name is required")
	}
	workspaceID, err := resolveWorkspaceID(ctx, session, client, orgID, projectID)
	if err != nil {
		return nil, fmt.Errorf("resolve workspace for org_id %q project_id %q: %w", orgID, projectID, err)
	}
	env, err := client.Environments.FindByName(workspaceID, envName)
	if err != nil {
		return nil, err
	}
	if env == nil {
		return nil, fmt.Errorf("environment %q not found in workspace for org_id %q project_id %q", envName, orgID, projectID)
	}
	return env, nil
}

// TrafficTypeByOrganizationProjectAndName resolves the workspace and returns the traffic type with the given name.
func TrafficTypeByOrganizationProjectAndName(ctx context.Context, session *internal.Session, client *split.APIClient, orgID, projectID, name string) (*split.TrafficType, error) {
	if name == "" {
		return nil, fmt.Errorf("traffic type name is required")
	}
	workspaceID, err := resolveWorkspaceID(ctx, session, client, orgID, projectID)
	if err != nil {
		return nil, fmt.Errorf("resolve workspace for org_id %q project_id %q: %w", orgID, projectID, err)
	}
	tt, err := client.TrafficTypes.FindByName(workspaceID, name)
	if err != nil {
		return nil, err
	}
	if tt == nil {
		return nil, fmt.Errorf("traffic type %q not found in workspace for org_id %q project_id %q", name, orgID, projectID)
	}
	return tt, nil
}

// FlagSetByOrganizationProjectAndName resolves the workspace and returns the flag set with the given name.
func FlagSetByOrganizationProjectAndName(ctx context.Context, session *internal.Session, client *split.APIClient, orgID, projectID, name string) (*split.FlagSet, error) {
	if name == "" {
		return nil, fmt.Errorf("flag set name is required")
	}
	workspaceID, err := resolveWorkspaceID(ctx, session, client, orgID, projectID)
	if err != nil {
		return nil, fmt.Errorf("resolve workspace for org_id %q project_id %q: %w", orgID, projectID, err)
	}
	return findFlagSetByNameWithRetry(ctx, client, workspaceID, orgID, projectID, name)
}

// LargeSegmentByOrganizationProjectAndName resolves the workspace and returns the large segment with the given name.
func LargeSegmentByOrganizationProjectAndName(ctx context.Context, session *internal.Session, client *split.APIClient, orgID, projectID, name string) (*split.LargeSegmentMetadata, error) {
	if name == "" {
		return nil, fmt.Errorf("large segment name is required")
	}
	workspaceID, err := resolveWorkspaceID(ctx, session, client, orgID, projectID)
	if err != nil {
		return nil, fmt.Errorf("resolve workspace for org_id %q project_id %q: %w", orgID, projectID, err)
	}
	return findLargeSegmentByNameWithRetry(ctx, client, workspaceID, orgID, projectID, name)
}

// ResolveWorkspaceIDFromSchema reads org_id and project_id from the resource or data source schema,
// resolves the workspace ID with caching and retries, and returns the workspace ID or diagnostics on failure.
func ResolveWorkspaceIDFromSchema(ctx context.Context, d *schema.ResourceData, meta interface{}) (string, diag.Diagnostics) {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags != nil && diags.HasError() {
		return "", diags
	}
	orgID, ok := d.Get("org_id").(string)
	if !ok || orgID == "" {
		return "", diag.FromErr(fmt.Errorf("org_id is required for workspace resolution"))
	}
	projectID, ok := d.Get("project_id").(string)
	if !ok || projectID == "" {
		return "", diag.FromErr(fmt.Errorf("project_id is required for workspace resolution"))
	}
	w, err := WorkspaceByOrganizationAndProject(ctx, sessionFromMeta(meta), client, orgID, projectID)
	if err != nil {
		return "", diag.Errorf("workspace not found for org_id %q and project_id %q: %v", orgID, projectID, err)
	}
	return w.ID, nil
}
