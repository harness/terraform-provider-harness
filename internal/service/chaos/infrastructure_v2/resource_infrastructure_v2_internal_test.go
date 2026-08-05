package infrastructure_v2

// White-box unit tests for the pure ID parser of the
// harness_chaos_infrastructure_v2 resource. Deterministic; no live API / TF_ACC.
//
// Chaos infrastructure v2 is always project-scoped, so the ID always has the
// 4-part form org_id/project_id/environment_id/infra_id.

import "testing"

func TestParseInfrastructureV2ID(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		wantOrg  string
		wantProj string
		wantEnv  string
		wantInfr string
		wantErr  bool
	}{
		{
			name:     "project scope full form",
			id:       "my-org/my-project/my-env/my-infra",
			wantOrg:  "my-org",
			wantProj: "my-project",
			wantEnv:  "my-env",
			wantInfr: "my-infra",
		},
		{name: "invalid too few segments", id: "my-org/my-project/my-infra", wantErr: true},
		{name: "invalid too many segments", id: "a/b/c/d/e", wantErr: true},
		{name: "invalid empty", id: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			org, proj, env, infra, err := parseInfrastructureV2ID(tt.id)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for id %q, got none", tt.id)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for id %q: %v", tt.id, err)
			}
			if org != tt.wantOrg || proj != tt.wantProj || env != tt.wantEnv || infra != tt.wantInfr {
				t.Errorf("parse(%q) = org=%q proj=%q env=%q infra=%q; want org=%q proj=%q env=%q infra=%q",
					tt.id, org, proj, env, infra, tt.wantOrg, tt.wantProj, tt.wantEnv, tt.wantInfr)
			}
		})
	}
}
