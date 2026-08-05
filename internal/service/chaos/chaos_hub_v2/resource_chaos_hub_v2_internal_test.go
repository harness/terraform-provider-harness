package chaos_hub_v2

// White-box unit tests for the pure import-ID parser of the
// harness_chaos_hub_v2 resource. Deterministic; no live API / TF_ACC.

import "testing"

func TestParseChaosHubV2ImportID(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		wantOrg     string
		wantProject string
		wantHub     string
		wantErr     bool
	}{
		{name: "account scope", id: "my-hub", wantHub: "my-hub"},
		{name: "org scope", id: "my-org/my-hub", wantOrg: "my-org", wantHub: "my-hub"},
		{name: "project scope", id: "my-org/my-project/my-hub", wantOrg: "my-org", wantProject: "my-project", wantHub: "my-hub"},
		{name: "invalid too many segments", id: "a/b/c/d", wantErr: true},
		{name: "invalid empty", id: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			org, project, hub, err := parseChaosHubV2ImportID(tt.id)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for id %q, got none", tt.id)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for id %q: %v", tt.id, err)
			}
			if org != tt.wantOrg || project != tt.wantProject || hub != tt.wantHub {
				t.Errorf("parse(%q) = org=%q project=%q hub=%q; want org=%q project=%q hub=%q",
					tt.id, org, project, hub, tt.wantOrg, tt.wantProject, tt.wantHub)
			}
		})
	}
}
