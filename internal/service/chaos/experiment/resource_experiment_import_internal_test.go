package experiment

// White-box unit test for the experiment import-ID parser. Deterministic; no
// live API / TF_ACC. Locks in the required 3-part "org_id/project_id/experiment_identity"
// form and rejection of malformed IDs.

import "testing"

func TestParseExperimentImportID(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		wantOrg     string
		wantProject string
		wantExp     string
		wantErr     bool
	}{
		{
			name:        "valid project-scoped id",
			id:          "my-org/my-project/my-experiment",
			wantOrg:     "my-org",
			wantProject: "my-project",
			wantExp:     "my-experiment",
		},
		{
			name:    "too few segments",
			id:      "my-org/my-project",
			wantErr: true,
		},
		{
			name:    "too many segments",
			id:      "my-org/my-project/my-experiment/extra",
			wantErr: true,
		},
		{
			name:    "empty segment",
			id:      "my-org//my-experiment",
			wantErr: true,
		},
		{
			name:    "empty id",
			id:      "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			org, project, exp, err := parseExperimentImportID(tt.id)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for id %q, got none", tt.id)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for id %q: %v", tt.id, err)
			}
			if org != tt.wantOrg || project != tt.wantProject || exp != tt.wantExp {
				t.Errorf("parse(%q) = org=%q project=%q exp=%q; want org=%q project=%q exp=%q",
					tt.id, org, project, exp, tt.wantOrg, tt.wantProject, tt.wantExp)
			}
		})
	}
}
