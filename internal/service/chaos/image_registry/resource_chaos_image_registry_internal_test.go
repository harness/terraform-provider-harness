package image_registry

// White-box unit tests for the image registry import-ID parser. Deterministic;
// no live API / TF_ACC.
//
// These lock in the supported import forms (project and infrastructure scope)
// and confirm that account-scoped (single-segment) imports remain unsupported -
// an org_id is always required, so a bare/account-level ID must error.

import "testing"

func TestParseImageRegistryImportID(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		wantOrg     string
		wantProject string
		wantInfra   string
		wantErr     bool
	}{
		{
			name:        "project scope",
			id:          "my-org/my-project",
			wantOrg:     "my-org",
			wantProject: "my-project",
		},
		{
			name:        "infrastructure scope",
			id:          "my-org/my-project/my-infra",
			wantOrg:     "my-org",
			wantProject: "my-project",
			wantInfra:   "my-infra",
		},
		{
			name:    "account scope (single segment) is unsupported",
			id:      "my-registry",
			wantErr: true,
		},
		{
			name:    "empty id",
			id:      "",
			wantErr: true,
		},
		{
			name:    "too many segments",
			id:      "a/b/c/d",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			org, project, infra, err := parseImageRegistryImportID(tt.id)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for id %q, got none", tt.id)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for id %q: %v", tt.id, err)
			}
			if org != tt.wantOrg || project != tt.wantProject || infra != tt.wantInfra {
				t.Errorf("parse(%q) = org=%q project=%q infra=%q; want org=%q project=%q infra=%q",
					tt.id, org, project, infra, tt.wantOrg, tt.wantProject, tt.wantInfra)
			}
		})
	}
}
