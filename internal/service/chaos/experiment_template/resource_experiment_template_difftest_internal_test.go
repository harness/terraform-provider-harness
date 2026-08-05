package experiment_template

// Comprehensive drift (perpetual-diff) regression tests for the
// harness_chaos_experiment_template resource.
//
// These are deterministic white-box unit tests (no live API / TF_ACC). They
// model the verified backend behaviour (verbatim store + echo of the submitted
// manifest) by:
//
//	build(config) -> JSON manifest -> json.Unmarshal into SDK response models
//	              -> read() into Terraform state
//
// and assert the state matches the config. Any field that the provider reads
// but does not write (or writes but does not read) breaks this invariant and
// surfaces to customers as a permanent "update in-place" plan.
//
// The suite intentionally exercises the full matrix of nested blocks
// (faults / actions / probes / values / vertices) and scalar spec fields, plus
// ordering preservation and the <+input> runtime-input DiffSuppress behaviour.

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// echoManifest models the backend verbatim store+echo for a full experiment
// template manifest: it decodes the JSON manifest the provider builds into the
// SDK GET response model, exactly as the backend would return it.
func echoManifest(t *testing.T, manifest string) *chaos.ChaosexperimenttemplateGetExperimentTemplateResponse {
	t.Helper()
	var resp chaos.ChaosexperimenttemplateGetExperimentTemplateResponse
	if err := json.Unmarshal([]byte(manifest), &resp); err != nil {
		t.Fatalf("failed to decode manifest as backend echo: %v\nmanifest: %s", err, manifest)
	}
	return &resp
}

func strPtr(s string) *string { return &s }

// ---------------------------------------------------------------------------
// Full-manifest round-trip: the strongest reproduction of the customer's
// "perpetual diff". Builds the manifest from a ResourceData, echoes it, reads
// it back into a fresh ResourceData, and asserts every nested field is stable.
// ---------------------------------------------------------------------------

func TestFullManifest_RoundTrip_NoDrift(t *testing.T) {
	raw := map[string]interface{}{
		"identity":     "customer-tmpl",
		"name":         "customer-tmpl",
		"hub_identity": "enterprise-hub",
		"tags":         []interface{}{"kubernetes", "network"},
		"spec": []interface{}{
			map[string]interface{}{
				"infra_type":     "KubernetesV2",
				"infra_id":       "prod-infra", // concrete (not <+input>) so it round-trips
				"cleanup_policy": "delete",
				"faults": []interface{}{
					map[string]interface{}{
						"identity":      "time-chaos",
						"name":          "time-chaos",
						"revision":      "v1",
						"is_enterprise": true,
						"auth_enabled":  false,
						"conditions_v2": []interface{}{
							map[string]interface{}{"operator": "AND", "values": []interface{}{"true", "<+input>"}},
						},
						"values": []interface{}{
							map[string]interface{}{"name": "TOTAL_CHAOS_DURATION", "value": "240"},
							map[string]interface{}{"name": "OFFSET", "value": "<+input>"},
							map[string]interface{}{"name": "POD_AFFECTED_PERCENTAGE", "value": "100"},
						},
					},
				},
				"probes": []interface{}{
					map[string]interface{}{
						"identity":               "latency-probe",
						"name":                   "latency-probe",
						"is_enterprise":          false,
						"duration":               "30",
						"weightage":              10,
						"enable_data_collection": true,
						"conditions_v2": []interface{}{
							map[string]interface{}{"operator": "OR", "values": []interface{}{"false"}},
						},
						"values": []interface{}{
							map[string]interface{}{"name": "COMPARATOR_VALUE", "value": "<+input>"},
						},
					},
				},
				"actions": []interface{}{
					map[string]interface{}{
						"identity":               "k8s-delay",
						"name":                   "k8s-delay",
						"is_enterprise":          false,
						"continue_on_completion": true,
						"conditions_v2": []interface{}{
							map[string]interface{}{"operator": "AND", "values": []interface{}{"<+input>"}},
						},
					},
				},
				"status_check_timeouts": []interface{}{
					map[string]interface{}{"delay": 5, "timeout": 180},
				},
			},
		},
	}

	d := schema.TestResourceDataRaw(t, ResourceExperimentTemplateSchema(), raw)

	manifest, err := buildExperimentTemplateManifest(d)
	if err != nil {
		t.Fatalf("buildExperimentTemplateManifest: %v", err)
	}

	resp := echoManifest(t, manifest)

	d2 := schema.TestResourceDataRaw(t, ResourceExperimentTemplateSchema(), map[string]interface{}{})
	if err := setExperimentTemplateData(d2, resp, "", "", "enterprise-hub"); err != nil {
		t.Fatalf("setExperimentTemplateData: %v", err)
	}

	checks := []struct {
		path string
		want interface{}
	}{
		{"name", "customer-tmpl"},
		{"identity", "customer-tmpl"},
		{"spec.0.infra_type", "KubernetesV2"},
		{"spec.0.infra_id", "prod-infra"},
		{"spec.0.cleanup_policy", "delete"},
		// fault
		{"spec.0.faults.0.identity", "time-chaos"},
		{"spec.0.faults.0.revision", "v1"},
		{"spec.0.faults.0.is_enterprise", true},
		{"spec.0.faults.0.auth_enabled", false},
		{"spec.0.faults.0.values.0.name", "TOTAL_CHAOS_DURATION"},
		{"spec.0.faults.0.values.0.value", "240"},
		{"spec.0.faults.0.values.1.value", "<+input>"},
		{"spec.0.faults.0.values.2.value", "100"},
		// probe
		{"spec.0.probes.0.name", "latency-probe"},
		{"spec.0.probes.0.duration", "30"},
		{"spec.0.probes.0.weightage", 10},
		{"spec.0.probes.0.enable_data_collection", true},
		{"spec.0.probes.0.values.0.name", "COMPARATOR_VALUE"},
		{"spec.0.probes.0.values.0.value", "<+input>"},
		// action
		{"spec.0.actions.0.name", "k8s-delay"},
		{"spec.0.actions.0.continue_on_completion", true},
		{"spec.0.actions.0.is_enterprise", false},
		// conditions_v2 round-trip (fault OR->AND, probe, action) through full manifest
		{"spec.0.faults.0.conditions_v2.0.operator", "AND"},
		{"spec.0.faults.0.conditions_v2.0.values.0", "true"},
		{"spec.0.faults.0.conditions_v2.0.values.1", "<+input>"},
		{"spec.0.probes.0.conditions_v2.0.operator", "OR"},
		{"spec.0.probes.0.conditions_v2.0.values.0", "false"},
		{"spec.0.actions.0.conditions_v2.0.operator", "AND"},
		{"spec.0.actions.0.conditions_v2.0.values.0", "<+input>"},
		// status check timeouts
		{"spec.0.status_check_timeouts.0.delay", 5},
		{"spec.0.status_check_timeouts.0.timeout", 180},
	}

	for _, c := range checks {
		got := d2.Get(c.path)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("drift at %q: got %#v (%T), want %#v (%T)", c.path, got, got, c.want, c.want)
		}
	}

	// tags (TypeList) round-trip and order preservation
	gotTags := d2.Get("tags").([]interface{})
	wantTags := []interface{}{"kubernetes", "network"}
	if !reflect.DeepEqual(gotTags, wantTags) {
		t.Errorf("tags drift: got %#v want %#v", gotTags, wantTags)
	}
}

// TestFullManifest_RoundTrip_InfraIDRuntimeInput asserts that a spec.infra_id of
// "<+input>" (which the manifest intentionally omits) does not produce a real
// diff because the schema's DiffSuppressFunc suppresses old="" vs new="<+input>".
func TestFullManifest_RoundTrip_InfraIDRuntimeInput(t *testing.T) {
	suppress := ResourceExperimentTemplateSchema()["spec"].Elem.(*schema.Resource).Schema["infra_id"].DiffSuppressFunc
	if suppress == nil {
		t.Fatal("expected DiffSuppressFunc on spec.infra_id")
	}
	if !suppress("spec.0.infra_id", "", "<+input>", nil) {
		t.Error("expected infra_id diff old=\"\" new=\"<+input>\" to be suppressed")
	}
	if suppress("spec.0.infra_id", "old-infra", "new-infra", nil) {
		t.Error("expected genuine infra_id change to NOT be suppressed")
	}
}

// ---------------------------------------------------------------------------
// Probe matrix
// ---------------------------------------------------------------------------

func TestProbes_RoundTrip_Matrix(t *testing.T) {
	tests := []struct {
		name  string
		probe map[string]interface{}
	}{
		{
			name: "data collection enabled",
			probe: map[string]interface{}{
				"identity": "p", "name": "p", "enable_data_collection": true,
			},
		},
		{
			name: "data collection disabled",
			probe: map[string]interface{}{
				"identity": "p", "name": "p", "enable_data_collection": false,
			},
		},
		{
			name: "full scalar fields",
			probe: map[string]interface{}{
				"identity": "p", "name": "p", "is_enterprise": true,
				"duration": "60", "weightage": 25, "revision": 3,
				"enable_data_collection": true,
			},
		},
		{
			name: "runtime and static values",
			probe: map[string]interface{}{
				"identity": "p", "name": "p",
				"values": []interface{}{
					map[string]interface{}{"name": "A", "value": "<+input>"},
					map[string]interface{}{"name": "B", "value": "static"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := []interface{}{tt.probe}
			got := readProbes(echoProbes(buildProbes(cfg)))
			assertProbeRoundTrip(t, tt.probe, got[0])
		})
	}
}

// TestProbes_RoundTrip_ConditionsV2 guards the probe conditions_v2 round-trip
// through the real SDK ExperimentConditions struct. buildProbes must emit
// conditionsV2 {operator, values} so it survives the build -> API echo -> read
// cycle (including "<+input>" runtime values); otherwise a probe with
// conditions_v2 shows a permanent "update in-place" diff. This exercises the
// SDK field added for CHAOS-12144, so it also fails if the SDK shape regresses.
func TestProbes_RoundTrip_ConditionsV2(t *testing.T) {
	cfg := []interface{}{
		map[string]interface{}{
			"identity": "p", "name": "p",
			"conditions_v2": []interface{}{
				map[string]interface{}{
					"operator": "AND",
					"values":   []interface{}{"true", "<+input>"},
				},
			},
		},
	}

	got := readProbes(echoProbes(buildProbes(cfg)))
	cv2, ok := got[0]["conditions_v2"].([]map[string]interface{})
	if !ok || len(cv2) != 1 {
		t.Fatalf("conditions_v2 dropped on round-trip (perpetual diff): %#v", got[0]["conditions_v2"])
	}
	if cv2[0]["operator"] != "AND" {
		t.Errorf("operator drift: got %v want AND", cv2[0]["operator"])
	}
	vals, _ := cv2[0]["values"].([]interface{})
	want := []string{"true", "<+input>"}
	if len(vals) != len(want) {
		t.Fatalf("values length: got %d want %d (%#v)", len(vals), len(want), vals)
	}
	for i, w := range want {
		if vals[i] != w {
			t.Errorf("values[%d]: got %v want %q", i, vals[i], w)
		}
	}
}

// TestFaultsActions_RoundTrip_ConditionsV2 guards conditions_v2 on faults and
// actions (same backend Conditions type).
func TestFaultsActions_RoundTrip_ConditionsV2(t *testing.T) {
	faultCfg := []interface{}{map[string]interface{}{
		"identity": "f", "name": "f",
		"conditions_v2": []interface{}{map[string]interface{}{
			"operator": "OR", "values": []interface{}{"<+input>"},
		}},
	}}
	gotF := readFaults(echoFaults(buildFaults(faultCfg)))
	if cv2, ok := gotF[0]["conditions_v2"].([]map[string]interface{}); !ok || len(cv2) != 1 || cv2[0]["operator"] != "OR" {
		t.Errorf("fault conditions_v2 drift: %#v", gotF[0]["conditions_v2"])
	}

	actionCfg := []interface{}{map[string]interface{}{
		"identity": "a", "name": "a",
		"conditions_v2": []interface{}{map[string]interface{}{
			"operator": "AND", "values": []interface{}{"true"},
		}},
	}}
	gotA := readActions(echoActions(buildActions(actionCfg)))
	if cv2, ok := gotA[0]["conditions_v2"].([]map[string]interface{}); !ok || len(cv2) != 1 || cv2[0]["operator"] != "AND" {
		t.Errorf("action conditions_v2 drift: %#v", gotA[0]["conditions_v2"])
	}
}

func assertProbeRoundTrip(t *testing.T, in, out map[string]interface{}) {
	t.Helper()
	if v, ok := in["enable_data_collection"]; ok {
		if out["enable_data_collection"] != v {
			t.Errorf("enable_data_collection: got %v want %v", out["enable_data_collection"], v)
		}
	}
	if v, ok := in["is_enterprise"]; ok && out["is_enterprise"] != v {
		t.Errorf("is_enterprise: got %v want %v", out["is_enterprise"], v)
	}
	if v, ok := in["duration"].(string); ok && v != "" && out["duration"] != v {
		t.Errorf("duration: got %v want %v", out["duration"], v)
	}
	if v, ok := in["weightage"].(int); ok && v > 0 && out["weightage"] != int32(v) {
		t.Errorf("weightage: got %v want %v", out["weightage"], v)
	}
	if v, ok := in["revision"].(int); ok && v > 0 && out["revision"] != int32(v) {
		t.Errorf("revision: got %v want %v", out["revision"], v)
	}
	if inVals, ok := in["values"].([]interface{}); ok {
		outVals, _ := out["values"].([]map[string]interface{})
		if len(outVals) != len(inVals) {
			t.Fatalf("values length: got %d want %d", len(outVals), len(inVals))
		}
		for i := range inVals {
			want := inVals[i].(map[string]interface{})
			if outVals[i]["name"] != want["name"] || outVals[i]["value"] != want["value"] {
				t.Errorf("values[%d]: got %#v want %#v", i, outVals[i], want)
			}
		}
	}
}

// ---------------------------------------------------------------------------
// Fault matrix
// ---------------------------------------------------------------------------

func TestFaults_RoundTrip_Matrix(t *testing.T) {
	tests := []struct {
		name string
		flt  map[string]interface{}
	}{
		{name: "auth enabled true", flt: map[string]interface{}{"identity": "f", "name": "f", "auth_enabled": true}},
		{name: "auth enabled false", flt: map[string]interface{}{"identity": "f", "name": "f", "auth_enabled": false}},
		{name: "enterprise with revision", flt: map[string]interface{}{"identity": "f", "name": "f", "is_enterprise": true, "revision": "v2"}},
		{
			name: "with values", flt: map[string]interface{}{"identity": "f", "name": "f",
				"values": []interface{}{
					map[string]interface{}{"name": "X", "value": "<+input>"},
					map[string]interface{}{"name": "Y", "value": "5"},
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readFaults(echoFaults(buildFaults([]interface{}{tt.flt})))
			if v, ok := tt.flt["auth_enabled"]; ok && got[0]["auth_enabled"] != v {
				t.Errorf("auth_enabled: got %v want %v", got[0]["auth_enabled"], v)
			}
			if v, ok := tt.flt["is_enterprise"]; ok && got[0]["is_enterprise"] != v {
				t.Errorf("is_enterprise: got %v want %v", got[0]["is_enterprise"], v)
			}
			if v, ok := tt.flt["revision"].(string); ok && v != "" && got[0]["revision"] != v {
				t.Errorf("revision: got %v want %v", got[0]["revision"], v)
			}
			if inVals, ok := tt.flt["values"].([]interface{}); ok {
				outVals, _ := got[0]["values"].([]map[string]interface{})
				if len(outVals) != len(inVals) {
					t.Fatalf("values length: got %d want %d", len(outVals), len(inVals))
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Action matrix
// ---------------------------------------------------------------------------

func TestActions_RoundTrip_Matrix(t *testing.T) {
	tests := []struct {
		name string
		act  map[string]interface{}
	}{
		{name: "continue true", act: map[string]interface{}{"identity": "a", "name": "a", "continue_on_completion": true}},
		{name: "continue false", act: map[string]interface{}{"identity": "a", "name": "a", "continue_on_completion": false}},
		{name: "enterprise with revision", act: map[string]interface{}{"identity": "a", "name": "a", "is_enterprise": true, "revision": 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readActions(echoActions(buildActions([]interface{}{tt.act})))
			if v, ok := tt.act["continue_on_completion"]; ok && got[0]["continue_on_completion"] != v {
				t.Errorf("continue_on_completion: got %v want %v", got[0]["continue_on_completion"], v)
			}
			if v, ok := tt.act["is_enterprise"]; ok && got[0]["is_enterprise"] != v {
				t.Errorf("is_enterprise: got %v want %v", got[0]["is_enterprise"], v)
			}
			if v, ok := tt.act["revision"].(int); ok && v > 0 && got[0]["revision"] != int32(v) {
				t.Errorf("revision: got %v want %v", got[0]["revision"], v)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Ordering preservation across multiple nested items
// ---------------------------------------------------------------------------

func TestOrdering_Preserved(t *testing.T) {
	probes := []interface{}{
		map[string]interface{}{"identity": "p1", "name": "p1"},
		map[string]interface{}{"identity": "p2", "name": "p2"},
		map[string]interface{}{"identity": "p3", "name": "p3"},
	}
	got := readProbes(echoProbes(buildProbes(probes)))
	for i, want := range []string{"p1", "p2", "p3"} {
		if got[i]["name"] != want {
			t.Errorf("probe order[%d]: got %v want %v", i, got[i]["name"], want)
		}
	}

	values := []interface{}{
		map[string]interface{}{"name": "V1", "value": "a"},
		map[string]interface{}{"name": "V2", "value": "b"},
		map[string]interface{}{"name": "V3", "value": "c"},
	}
	gotVals := readValues(echoValues(buildValues(values)))
	for i, want := range []string{"V1", "V2", "V3"} {
		if gotVals[i]["name"] != want {
			t.Errorf("value order[%d]: got %v want %v", i, gotVals[i]["name"], want)
		}
	}
}

func echoValues(built []map[string]interface{}) []chaos.TemplateVariableMinimum {
	b, err := json.Marshal(built)
	if err != nil {
		panic(err)
	}
	var out []chaos.TemplateVariableMinimum
	if err := json.Unmarshal(b, &out); err != nil {
		panic(err)
	}
	return out
}

// ---------------------------------------------------------------------------
// readValues type coercion: the backend may return numeric/bool JSON scalars;
// the provider must coerce them to the schema's string type deterministically
// so they match the (string) config value and do not drift.
// ---------------------------------------------------------------------------

func TestReadValues_TypeCoercion(t *testing.T) {
	mk := func(v interface{}) *interface{} { return &v }
	in := []chaos.TemplateVariableMinimum{
		{Name: "str", Value: mk("hello")},
		{Name: "num_int", Value: mk(float64(240))}, // JSON numbers decode as float64
		{Name: "num_frac", Value: mk(float64(2.5))},
		{Name: "boolean", Value: mk(true)},
		{Name: "runtime", Value: mk("<+input>")},
	}
	out := readValues(in)
	want := map[string]string{
		"str":      "hello",
		"num_int":  "240",
		"num_frac": "2.5",
		"boolean":  "true",
		"runtime":  "<+input>",
	}
	for _, o := range out {
		name := o["name"].(string)
		if o["value"] != want[name] {
			t.Errorf("coercion %q: got %#v want %q", name, o["value"], want[name])
		}
	}
}

// ---------------------------------------------------------------------------
// Vertices round-trip (workflow graph the customer uses)
// ---------------------------------------------------------------------------

func TestVertices_RoundTrip(t *testing.T) {
	cfg := []interface{}{
		map[string]interface{}{
			"name": "vertex-1",
			"start": []interface{}{
				map[string]interface{}{
					"faults": []interface{}{
						map[string]interface{}{"name": "time-chaos"},
					},
					"probes": []interface{}{
						map[string]interface{}{"name": "latency-probe"},
					},
				},
			},
			"end": []interface{}{
				map[string]interface{}{
					"actions": []interface{}{
						map[string]interface{}{"name": "k8s-delay"},
					},
				},
			},
		},
	}

	built := buildVertices(cfg)
	b, _ := json.Marshal(built)
	var echoed []chaos.ExperimenttemplateVertex
	if err := json.Unmarshal(b, &echoed); err != nil {
		t.Fatalf("echo vertices: %v", err)
	}
	got := readVertices(echoed)
	if got[0]["name"] != "vertex-1" {
		t.Fatalf("vertex name drift: got %v", got[0]["name"])
	}
	start, _ := got[0]["start"].([]map[string]interface{})
	if len(start) != 1 {
		t.Fatalf("expected start block, got %#v", got[0]["start"])
	}
	faults, _ := start[0]["faults"].([]map[string]interface{})
	if len(faults) != 1 || faults[0]["name"] != "time-chaos" {
		t.Errorf("start.faults drift: got %#v", start[0]["faults"])
	}
	probes, _ := start[0]["probes"].([]map[string]interface{})
	if len(probes) != 1 || probes[0]["name"] != "latency-probe" {
		t.Errorf("start.probes drift: got %#v", start[0]["probes"])
	}
	end, _ := got[0]["end"].([]map[string]interface{})
	if len(end) != 1 {
		t.Fatalf("expected end block, got %#v", got[0]["end"])
	}
	actions, _ := end[0]["actions"].([]map[string]interface{})
	if len(actions) != 1 || actions[0]["name"] != "k8s-delay" {
		t.Errorf("end.actions drift: got %#v", end[0]["actions"])
	}
}
