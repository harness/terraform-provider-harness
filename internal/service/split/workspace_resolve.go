package split

import (
	"context"
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/split"
)

const (
	workspaceResolveMaxAttempts = 5
	workspaceResolveBaseDelay   = 200 * time.Millisecond
	workspaceResolveMaxDelay    = 8 * time.Second

	// Flag set list (FindByName) can lag after Create in some Split workspaces; poll with a bounded total wait.
	flagSetFindMaxAttempts  = 60
	flagSetFindPollInterval = 400 * time.Millisecond
)

// isWorkspaceListRetriable returns true for Split workspace list throttling and transient duplicate-key errors
// observed when concurrent requests race on lazy Harness→Split project mapping.
func isWorkspaceListRetriable(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	if strings.Contains(msg, "workspaces list: 429") || strings.Contains(msg, "429 Too Many Requests") {
		return true
	}
	if strings.Contains(msg, "workspaces list: 500") &&
		(strings.Contains(msg, "E11000") || strings.Contains(msg, "duplicate key")) {
		return true
	}
	return false
}

func sleepWorkspaceResolveBackoff(ctx context.Context, attempt int) error {
	shift := attempt
	if shift > 6 {
		shift = 6
	}
	base := workspaceResolveBaseDelay * time.Duration(uint(1)<<uint(shift))
	if base > workspaceResolveMaxDelay {
		base = workspaceResolveMaxDelay
	}
	jitterN := int64(base / 5)
	if jitterN < 1 {
		jitterN = 1
	}
	d := base + time.Duration(rand.Int64N(jitterN))
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

func findByOrganizationAndProjectWithRetry(ctx context.Context, client *split.APIClient, orgID, projectID string) ([]split.Workspace, error) {
	var lastErr error
	for attempt := 0; attempt < workspaceResolveMaxAttempts; attempt++ {
		if attempt > 0 {
			if err := sleepWorkspaceResolveBackoff(ctx, attempt-1); err != nil {
				return nil, err
			}
		}
		ws, err := client.Workspaces.FindByOrganizationAndProject(orgID, projectID)
		if err == nil {
			return ws, nil
		}
		lastErr = err
		if !isWorkspaceListRetriable(err) || attempt == workspaceResolveMaxAttempts-1 {
			return nil, err
		}
	}
	return nil, lastErr
}

func findByNameWithRetry(ctx context.Context, client *split.APIClient, name, nameOp string) ([]split.Workspace, error) {
	var lastErr error
	for attempt := 0; attempt < workspaceResolveMaxAttempts; attempt++ {
		if attempt > 0 {
			if err := sleepWorkspaceResolveBackoff(ctx, attempt-1); err != nil {
				return nil, err
			}
		}
		ws, err := client.Workspaces.FindByName(name, nameOp)
		if err == nil {
			return ws, nil
		}
		lastErr = err
		if !isWorkspaceListRetriable(err) || attempt == workspaceResolveMaxAttempts-1 {
			return nil, err
		}
	}
	return nil, lastErr
}

func sleepFlagSetFindPoll(ctx context.Context) error {
	t := time.NewTimer(flagSetFindPollInterval)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

// findFlagSetByNameWithRetry calls the SDK until the flag set appears or attempts are exhausted (list can lag briefly after create).
func findFlagSetByNameWithRetry(ctx context.Context, client *split.APIClient, workspaceID, orgID, projectID, name string) (*split.FlagSet, error) {
	for attempt := 0; attempt < flagSetFindMaxAttempts; attempt++ {
		if attempt > 0 {
			if err := sleepFlagSetFindPoll(ctx); err != nil {
				return nil, err
			}
		}
		fs, err := client.FlagSets.FindByName(workspaceID, name)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			return fs, nil
		}
	}
	return nil, fmt.Errorf("flag set %q not found in workspace for org_id %q project_id %q after %d list attempts", name, orgID, projectID, flagSetFindMaxAttempts)
}

// isLargeSegmentGetRetriable is true when Get returns 404 shortly after create (eventual consistency).
func isLargeSegmentGetRetriable(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	return strings.Contains(s, "404") || strings.Contains(s, "not found") || strings.Contains(s, "Not Found")
}

// findLargeSegmentByNameWithRetry calls LargeSegments.Get until the segment appears or attempts are exhausted.
func findLargeSegmentByNameWithRetry(ctx context.Context, client *split.APIClient, workspaceID, orgID, projectID, name string) (*split.LargeSegmentMetadata, error) {
	var lastErr error
	for attempt := 0; attempt < flagSetFindMaxAttempts; attempt++ {
		if attempt > 0 {
			if err := sleepFlagSetFindPoll(ctx); err != nil {
				return nil, err
			}
		}
		seg, err := client.LargeSegments.Get(workspaceID, name)
		if err == nil {
			return seg, nil
		}
		lastErr = err
		if !isLargeSegmentGetRetriable(err) || attempt == flagSetFindMaxAttempts-1 {
			return nil, err
		}
	}
	return nil, lastErr
}
