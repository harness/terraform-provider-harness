package security_governance_v3

// White-box unit test for the shared project-scoped import-ID parser used by
// both the condition and rule resources. Deterministic; no live API / TF_ACC.
// Locks in the required 3-part "org_id/project_id/<id>" form, non-empty
// segments, and the labeled error message.

import (
	"strings"
	"testing"
)

func TestParseScopedImportIDV3(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		label       string
		wantOrg     string
		wantProject string
		wantID      string
		wantErr     bool
	}{
		{
			name:        "valid rule id",
			id:          "my-org/my-project/my-rule",
			label:       "rule-id",
			wantOrg:     "my-org",
			wantProject: "my-project",
			wantID:      "my-rule",
		},
		{
			name:        "valid condition id",
			id:          "my-org/my-project/my-condition",
			label:       "condition-id",
			wantOrg:     "my-org",
			wantProject: "my-project",
			wantID:      "my-condition",
		},
		{
			name:    "too few segments",
			id:      "my-org/my-project",
			label:   "rule-id",
			wantErr: true,
		},
		{
			name:    "too many segments",
			id:      "my-org/my-project/my-rule/extra",
			label:   "rule-id",
			wantErr: true,
		},
		{
			name:    "empty middle segment",
			id:      "my-org//my-rule",
			label:   "rule-id",
			wantErr: true,
		},
		{
			name:    "empty id",
			id:      "",
			label:   "condition-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			org, project, id, err := parseScopedImportIDV3(tt.id, tt.label)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for id %q, got none", tt.id)
				}
				// error message should reference the labeled segment
				if !strings.Contains(err.Error(), strings.Split(tt.label, "-")[0]) {
					t.Errorf("error %q does not reference label %q", err.Error(), tt.label)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for id %q: %v", tt.id, err)
			}
			if org != tt.wantOrg || project != tt.wantProject || id != tt.wantID {
				t.Errorf("parse(%q) = org=%q project=%q id=%q; want org=%q project=%q id=%q",
					tt.id, org, project, id, tt.wantOrg, tt.wantProject, tt.wantID)
			}
		})
	}
}
