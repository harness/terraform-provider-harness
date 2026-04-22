package split

import (
	"context"
	"fmt"
)

// AccWorkspaceID returns the Split workspace ID for the given Harness org and project identifiers.
// Used by acceptance tests for Split API verification (e.g. tail-step checks after FME resources are destroyed).
func AccWorkspaceID(ctx context.Context, meta interface{}, orgID, projectID string) (string, error) {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return "", fmt.Errorf("split client: %v", diags)
	}
	w, err := WorkspaceByOrganizationAndProject(ctx, sessionFromMeta(meta), client, orgID, projectID)
	if err != nil {
		return "", err
	}
	return w.ID, nil
}

// AccSegmentEnvAssociationActive reports whether a classic segment is active in a Split environment
// (same semantics as the harness_fme_segment_environment_association resource Read).
func AccSegmentEnvAssociationActive(ctx context.Context, meta interface{}, orgID, projectID, envID, segName string) (bool, error) {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return false, fmt.Errorf("split client: %v", diags)
	}
	wsID, err := resolveWorkspaceID(ctx, sessionFromMeta(meta), client, orgID, projectID)
	if err != nil {
		return false, err
	}
	return segmentActivatedInEnvironment(ctx, client, wsID, envID, segName)
}

// AccLargeSegmentDefinedInEnvironment reports whether a large segment has a definition in the environment.
func AccLargeSegmentDefinedInEnvironment(ctx context.Context, meta interface{}, orgID, projectID, envID, segName string) (bool, error) {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return false, fmt.Errorf("split client: %v", diags)
	}
	wsID, err := resolveWorkspaceID(ctx, sessionFromMeta(meta), client, orgID, projectID)
	if err != nil {
		return false, err
	}
	return largeSegmentDefinedInEnvironment(ctx, client, wsID, envID, segName)
}

// AccRuleBasedSegmentInEnvironment reports whether a rule-based segment entry exists in the environment list.
func AccRuleBasedSegmentInEnvironment(ctx context.Context, meta interface{}, orgID, projectID, envID, segName string) (bool, error) {
	client, diags := SplitClientFromMeta(ctx, meta)
	if diags.HasError() {
		return false, fmt.Errorf("split client: %v", diags)
	}
	wsID, err := resolveWorkspaceID(ctx, sessionFromMeta(meta), client, orgID, projectID)
	if err != nil {
		return false, err
	}
	entry, err := ruleBasedSegmentPresentInEnvironment(ctx, client, wsID, envID, segName)
	if err != nil {
		return false, err
	}
	return entry != nil, nil
}
