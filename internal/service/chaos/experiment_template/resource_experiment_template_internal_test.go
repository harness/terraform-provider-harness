package experiment_template

// White-box (package-internal) unit tests for the pure build/read/import helper
// functions of the harness_chaos_experiment_template resource.
//
// These tests are deterministic and DO NOT require a live Harness API or
// TF_ACC. They model the (verified) backend behaviour for experiment
// templates: the backend stores the submitted manifest verbatim (base64 YAML)
// and echoes it back unchanged on read. We reproduce that echo here by
// marshalling the manifest the provider builds and unmarshalling it into the
// SDK response models, then flattening it back into Terraform state.
//
// The core invariant under test is round-trip symmetry:
//
//	read(echo(build(config))) == config
//
// Any field the provider writes to state on read but never sends in the build
// manifest (or vice-versa) breaks this invariant and manifests as a perpetual
// "update in-place" diff for customers.

import (
	"encoding/json"
	"testing"

	"github.com/harness/harness-go-sdk/harness/chaos"
)

// echoProbes simulates the backend verbatim store+echo for probes.
func echoProbes(built []map[string]interface{}) []chaos.ExperimenttemplateProbe {
	b, err := json.Marshal(built)
	if err != nil {
		panic(err)
	}
	var out []chaos.ExperimenttemplateProbe
	if err := json.Unmarshal(b, &out); err != nil {
		panic(err)
	}
	return out
}

func echoFaults(built []map[string]interface{}) []chaos.ExperimenttemplateFault {
	b, err := json.Marshal(built)
	if err != nil {
		panic(err)
	}
	var out []chaos.ExperimenttemplateFault
	if err := json.Unmarshal(b, &out); err != nil {
		panic(err)
	}
	return out
}

func echoActions(built []map[string]interface{}) []chaos.ExperimenttemplateAction {
	b, err := json.Marshal(built)
	if err != nil {
		panic(err)
	}
	var out []chaos.ExperimenttemplateAction
	if err := json.Unmarshal(b, &out); err != nil {
		panic(err)
	}
	return out
}

// ---------------------------------------------------------------------------
// Issue 2: perpetual no-op diff caused by build/read asymmetry
// ---------------------------------------------------------------------------

// TestProbes_RoundTrip_EnableDataCollection reproduces the customer-reported
// perpetual diff: a probe with enable_data_collection=true loses that value on
// the round-trip because buildProbes never writes it to the manifest, while
// readProbes always reads it back from the API response.
func TestProbes_RoundTrip_EnableDataCollection(t *testing.T) {
	cfg := []interface{}{
		map[string]interface{}{
			"identity":               "probe-tmpl-1",
			"name":                   "latency-probe",
			"is_enterprise":          false,
			"duration":               "30",
			"weightage":              10,
			"enable_data_collection": true,
			"values": []interface{}{
				map[string]interface{}{"name": "COMPARATOR_VALUE", "value": "<+input>"},
			},
		},
	}

	got := readProbes(echoProbes(buildProbes(cfg)))
	if len(got) != 1 {
		t.Fatalf("expected 1 probe, got %d", len(got))
	}

	if got[0]["enable_data_collection"] != true {
		t.Errorf("enable_data_collection round-trip drift: got %v, want true "+
			"(buildProbes must send enableDataCollection in the manifest)", got[0]["enable_data_collection"])
	}
}

// TestProbes_RoundTrip_Values ensures probe variable values (including runtime
// <+input>) survive the round-trip unchanged.
func TestProbes_RoundTrip_Values(t *testing.T) {
	cfg := []interface{}{
		map[string]interface{}{
			"identity":      "probe-tmpl-1",
			"name":          "latency-probe",
			"is_enterprise": false,
			"values": []interface{}{
				map[string]interface{}{"name": "COMPARATOR_VALUE", "value": "<+input>"},
				map[string]interface{}{"name": "THRESHOLD", "value": "240"},
			},
		},
	}

	got := readProbes(echoProbes(buildProbes(cfg)))
	vals, _ := got[0]["values"].([]map[string]interface{})
	if len(vals) != 2 {
		t.Fatalf("expected 2 values, got %d (%#v)", len(vals), got[0]["values"])
	}
	if vals[0]["name"] != "COMPARATOR_VALUE" || vals[0]["value"] != "<+input>" {
		t.Errorf("value[0] drift: got %#v", vals[0])
	}
	if vals[1]["name"] != "THRESHOLD" || vals[1]["value"] != "240" {
		t.Errorf("value[1] drift: got %#v", vals[1])
	}
}

// TestFaults_RoundTrip ensures fault fields (revision, auth_enabled, values)
// survive the round-trip.
func TestFaults_RoundTrip(t *testing.T) {
	cfg := []interface{}{
		map[string]interface{}{
			"identity":      "time-chaos",
			"name":          "time-chaos",
			"revision":      "v1",
			"is_enterprise": true,
			"auth_enabled":  false,
			"values": []interface{}{
				map[string]interface{}{"name": "TOTAL_CHAOS_DURATION", "value": "240"},
				map[string]interface{}{"name": "OFFSET", "value": "<+input>"},
			},
		},
	}

	got := readFaults(echoFaults(buildFaults(cfg)))
	if got[0]["revision"] != "v1" {
		t.Errorf("fault revision drift: got %v want v1", got[0]["revision"])
	}
	if got[0]["is_enterprise"] != true {
		t.Errorf("fault is_enterprise drift: got %v want true", got[0]["is_enterprise"])
	}
	if got[0]["auth_enabled"] != false {
		t.Errorf("fault auth_enabled drift: got %v want false", got[0]["auth_enabled"])
	}
	vals, _ := got[0]["values"].([]map[string]interface{})
	if len(vals) != 2 {
		t.Fatalf("expected 2 fault values, got %d", len(vals))
	}
}

// TestActions_RoundTrip ensures action fields survive the round-trip.
func TestActions_RoundTrip(t *testing.T) {
	cfg := []interface{}{
		map[string]interface{}{
			"identity":               "k8s-delay",
			"name":                   "k8s-delay",
			"is_enterprise":          false,
			"continue_on_completion": true,
		},
	}

	got := readActions(echoActions(buildActions(cfg)))
	if got[0]["continue_on_completion"] != true {
		t.Errorf("action continue_on_completion drift: got %v want true", got[0]["continue_on_completion"])
	}
	if got[0]["is_enterprise"] != false {
		t.Errorf("action is_enterprise drift: got %v want false", got[0]["is_enterprise"])
	}
}

// ---------------------------------------------------------------------------
// Issue 1: import ID format mismatch
// ---------------------------------------------------------------------------

func TestParseExperimentTemplateImportID(t *testing.T) {
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
		{
			name:    "invalid empty",
			id:      "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			org, project, hub, id, canon, err := parseExperimentTemplateImportID(tt.id)
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
