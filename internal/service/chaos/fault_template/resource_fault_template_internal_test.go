package fault_template

// White-box unit tests for the pure import-ID parser of the
// harness_chaos_fault_template resource. Deterministic; no live API / TF_ACC.
//
// These lock in multi-scope import support (account/org/project), which
// previously failed because the import handler required exactly 4 path
// segments (the same customer-reported bug as experiment_template).

import "testing"

func TestParseFaultTemplateImportID(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		wantOrg     string
		wantProject string
		wantHub     string
		wantID      string
		wantCanon   string
		wantErr     bool
	}{
		{
			name:      "account scope short form",
			id:        "my-hub/my-template",
			wantHub:   "my-hub",
			wantID:    "my-template",
			wantCanon: "//my-hub/my-template",
		},
		{
			name:      "org scope short form",
			id:        "my-org/my-hub/my-template",
			wantOrg:   "my-org",
			wantHub:   "my-hub",
			wantID:    "my-template",
			wantCanon: "my-org//my-hub/my-template",
		},
		{
			name:        "project scope full form",
			id:          "my-org/my-project/my-hub/my-template",
			wantOrg:     "my-org",
			wantProject: "my-project",
			wantHub:     "my-hub",
			wantID:      "my-template",
			wantCanon:   "my-org/my-project/my-hub/my-template",
		},
		{
			name:      "account scope canonical form",
			id:        "//my-hub/my-template",
			wantHub:   "my-hub",
			wantID:    "my-template",
			wantCanon: "//my-hub/my-template",
		},
		{
			name:    "invalid single segment",
			id:      "my-template",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			org, project, hub, id, canon, err := parseFaultTemplateImportID(tt.id)
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
			if canon != tt.wantCanon {
				t.Errorf("parse(%q) canonical = %q; want %q", tt.id, canon, tt.wantCanon)
			}
		})
	}
}

// TestGuardUnsupportedChaosAuthTLS locks in the safety guard: spec.chaos.auth and
// spec.chaos.tls are declared in the schema but not plumbed through to the API, so
// setting them must return a clear error instead of silently dropping the
// (security-sensitive) values. Empty/absent blocks must pass.
func TestGuardUnsupportedChaosAuthTLS(t *testing.T) {
	tests := []struct {
		name    string
		config  map[string]interface{}
		wantErr bool
	}{
		{
			name:    "no auth or tls",
			config:  map[string]interface{}{"fault_name": "pod-delete"},
			wantErr: false,
		},
		{
			name:    "empty auth list",
			config:  map[string]interface{}{"auth": []interface{}{}},
			wantErr: false,
		},
		{
			name:    "auth set",
			config:  map[string]interface{}{"auth": []interface{}{map[string]interface{}{"redis": []interface{}{map[string]interface{}{"password": "x"}}}}},
			wantErr: true,
		},
		{
			name:    "tls set",
			config:  map[string]interface{}{"tls": []interface{}{map[string]interface{}{"ca_certificate": "x"}}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := guardUnsupportedChaosAuthTLS(tt.config)
			if tt.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
