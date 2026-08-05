package action_template

// White-box unit tests for the pure import-ID parser of the
// harness_chaos_action_template resource. Deterministic; no live API / TF_ACC.

import "testing"

func TestParseActionTemplateImportID(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		wantOrg     string
		wantProject string
		wantHub     string
		wantID      string
		wantErr     bool
	}{
		{name: "account scope", id: "my-hub/my-template", wantHub: "my-hub", wantID: "my-template"},
		{name: "org scope", id: "my-org/my-hub/my-template", wantOrg: "my-org", wantHub: "my-hub", wantID: "my-template"},
		{name: "project scope", id: "my-org/my-project/my-hub/my-template", wantOrg: "my-org", wantProject: "my-project", wantHub: "my-hub", wantID: "my-template"},
		{name: "account canonical form", id: "//my-hub/my-template", wantHub: "my-hub", wantID: "my-template"},
		{name: "invalid single segment", id: "my-template", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			org, project, hub, id, err := parseActionTemplateImportID(tt.id)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for id %q, got none", tt.id)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for id %q: %v", tt.id, err)
			}
			if org != tt.wantOrg || project != tt.wantProject || hub != tt.wantHub || id != tt.wantID {
				t.Errorf("parse(%q) = org=%q project=%q hub=%q id=%q; want org=%q project=%q hub=%q id=%q",
					tt.id, org, project, hub, id, tt.wantOrg, tt.wantProject, tt.wantHub, tt.wantID)
			}
		})
	}
}
