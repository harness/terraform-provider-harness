package chaos_hub_v2

// White-box read-side test: setChaosHubV2Data must map every created field of a
// GetHub response back into state so there is no create/read asymmetry (drift).
// Deterministic; no live API / TF_ACC.

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestSetChaosHubV2Data_RoundTrip(t *testing.T) {
	hub := &chaos.Chaoshubv2GetHubResponse{
		Identity:    "hub-1",
		Name:        "Hub One",
		HubId:       "hid-1",
		AccountID:   "acc-1",
		Description: "a hub",
		RepoBranch:  "main",
		RepoName:    "chaos-hub",
		RepoUrl:     "https://example.com/chaos-hub.git",
		ConnectorId: "conn-1",
		Tags:        []string{"a", "b"},
		IsDefault:   false,
		IsRemoved:   false,
	}

	d := schema.TestResourceDataRaw(t, ResourceChaosHubV2().Schema, map[string]interface{}{})
	if err := setChaosHubV2Data(d, hub, "acc-1", "org-1", "proj-1"); err != nil {
		t.Fatalf("setChaosHubV2Data error: %v", err)
	}

	checks := map[string]interface{}{
		"identity":     "hub-1",
		"name":         "Hub One",
		"hub_id":       "hid-1",
		"account_id":   "acc-1",
		"org_id":       "org-1",
		"project_id":   "proj-1",
		"description":  "a hub",
		"repo_branch":  "main",
		"repo_name":    "chaos-hub",
		"repo_url":     "https://example.com/chaos-hub.git",
		"connector_id": "conn-1",
		"is_default":   false,
		"is_removed":   false,
	}
	for path, want := range checks {
		if got := d.Get(path); got != want {
			t.Errorf("%s: got %#v (%T), want %#v", path, got, got, want)
		}
	}

	tags := d.Get("tags").([]interface{})
	if len(tags) != 2 {
		t.Errorf("tags drift: got %#v", tags)
	}
}

// TestChaosHubV2_ForceNewContract guards the immutability contract: the update
// API (Chaoshubv2UpdateHubRequest) only accepts name/description/tags, so every
// other user-settable field must be ForceNew. Without this, changing e.g.
// connector_ref or repo_branch would silently no-op and cause a perpetual diff.
func TestChaosHubV2_ForceNewContract(t *testing.T) {
	s := ResourceChaosHubV2().Schema

	// Fields that MUST force recreation (not updatable via the API).
	forceNew := []string{"identity", "org_id", "project_id", "connector_ref", "repo_branch", "repo_name"}
	for _, name := range forceNew {
		f, ok := s[name]
		if !ok {
			t.Errorf("missing schema field %q", name)
			continue
		}
		if !f.ForceNew {
			t.Errorf("field %q must be ForceNew (update API cannot change it)", name)
		}
	}

	// Fields that MUST remain updatable in place (accepted by the update API).
	for _, name := range []string{"name", "description", "tags"} {
		if f, ok := s[name]; ok && f.ForceNew {
			t.Errorf("field %q must be updatable in place, not ForceNew", name)
		}
	}
}
