package security_governance_v3

// Round-trip (expand -> flatten) coverage for the security governance V3
// condition specs and the rule time-window flattener. These are the most
// deeply nested chaos schemas and thus the most drift-prone. Deterministic;
// no live API / TF_ACC.

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func toStrs(v interface{}) []string {
	switch t := v.(type) {
	case []string:
		return t
	case []interface{}:
		out := make([]string, len(t))
		for i, e := range t {
			out[i], _ = e.(string)
		}
		return out
	}
	return nil
}

func TestFaultSpecV3_RoundTrip(t *testing.T) {
	in := []interface{}{map[string]interface{}{
		"operator": "IN",
		"faults": []interface{}{
			map[string]interface{}{"fault_type": "FAULT", "name": "pod-delete"},
			map[string]interface{}{"fault_type": "FAULT_GROUP", "name": "network"},
		},
	}}

	out := flattenFaultSpecV3(expandFaultSpecV3(in))
	if len(out) != 1 {
		t.Fatalf("expected 1 fault_spec, got %d", len(out))
	}
	m := out[0].(map[string]interface{})
	if m["operator"] != "IN" {
		t.Errorf("operator drift: got %v want IN", m["operator"])
	}
	faults := m["faults"].([]map[string]interface{})
	if len(faults) != 2 {
		t.Fatalf("expected 2 faults, got %d", len(faults))
	}
	if faults[0]["fault_type"] != "FAULT" || faults[0]["name"] != "pod-delete" {
		t.Errorf("fault[0] drift: %#v", faults[0])
	}
	if faults[1]["fault_type"] != "FAULT_GROUP" || faults[1]["name"] != "network" {
		t.Errorf("fault[1] drift: %#v", faults[1])
	}
}

// TestFaultTypeMapping documents the intentional FAULT_NAME -> FAULT
// normalization while confirming FAULT and FAULT_GROUP round-trip exactly.
func TestFaultTypeMapping(t *testing.T) {
	cases := map[string]string{
		"FAULT":       "FAULT",
		"FAULT_GROUP": "FAULT_GROUP",
		"FAULT_NAME":  "FAULT", // normalized (legacy input)
	}
	for in, want := range cases {
		if got := faultTypeFromREST(faultTypeToREST(in)); got != want {
			t.Errorf("faultType(%q) round-trip = %q, want %q", in, got, want)
		}
	}
}

func TestK8sSpecV3_RoundTrip(t *testing.T) {
	in := map[string]interface{}{
		"infra_spec": []interface{}{map[string]interface{}{
			"operator":  "IN",
			"infra_ids": []interface{}{"infra-a", "infra-b"},
		}},
		"application_spec": []interface{}{map[string]interface{}{
			"operator": "NOT_IN",
			"workloads": []interface{}{map[string]interface{}{
				"namespace":          "prod",
				"kind":               "Deployment",
				"label":              "app=web",
				"services":           []interface{}{"svc-1"},
				"application_map_id":  "map-1",
				"namespace_labels":   map[string]interface{}{"team": "core"},
			}},
		}},
		"chaos_service_account_spec": []interface{}{map[string]interface{}{
			"operator":         "IN",
			"service_accounts": []interface{}{"sa-1", "sa-2"},
		}},
	}

	out := flattenK8sSpecV3(expandK8sSpecV3(in))
	if len(out) != 1 {
		t.Fatalf("expected flattened k8s_spec, got %d", len(out))
	}
	m := out[0].(map[string]interface{})

	infra := m["infra_spec"].([]interface{})[0].(map[string]interface{})
	if infra["operator"] != "IN" {
		t.Errorf("infra_spec operator drift: %v", infra["operator"])
	}
	if got := toStrs(infra["infra_ids"]); len(got) != 2 || got[0] != "infra-a" || got[1] != "infra-b" {
		t.Errorf("infra_ids drift: %#v", infra["infra_ids"])
	}

	app := m["application_spec"].([]interface{})[0].(map[string]interface{})
	if app["operator"] != "NOT_IN" {
		t.Errorf("application_spec operator drift: %v", app["operator"])
	}
	wl := app["workloads"].([]map[string]interface{})[0]
	if wl["namespace"] != "prod" || wl["kind"] != "Deployment" || wl["label"] != "app=web" || wl["application_map_id"] != "map-1" {
		t.Errorf("workload scalar drift: %#v", wl)
	}
	if got := toStrs(wl["services"]); len(got) != 1 || got[0] != "svc-1" {
		t.Errorf("workload services drift: %#v", wl["services"])
	}
	nl := wl["namespace_labels"].(map[string]interface{})
	if nl["team"] != "core" {
		t.Errorf("namespace_labels drift: %#v", nl)
	}

	sa := m["chaos_service_account_spec"].([]interface{})[0].(map[string]interface{})
	if got := toStrs(sa["service_accounts"]); len(got) != 2 {
		t.Errorf("service_accounts drift: %#v", sa["service_accounts"])
	}
}

func TestMachineSpecV3_RoundTrip(t *testing.T) {
	in := map[string]interface{}{
		"infra_spec": []interface{}{map[string]interface{}{
			"operator":  "IN",
			"infra_ids": []interface{}{"m-1"},
		}},
	}
	out := flattenMachineSpecV3(expandMachineSpecV3(in))
	if len(out) != 1 {
		t.Fatalf("expected flattened machine_spec, got %d", len(out))
	}
	infra := out[0].(map[string]interface{})["infra_spec"].([]interface{})[0].(map[string]interface{})
	if infra["operator"] != "IN" {
		t.Errorf("machine infra operator drift: %v", infra["operator"])
	}
	if got := toStrs(infra["infra_ids"]); len(got) != 1 || got[0] != "m-1" {
		t.Errorf("machine infra_ids drift: %#v", infra["infra_ids"])
	}
}

func TestSetRuleV3Data_TimeWindowsFlatten(t *testing.T) {
	resp := &chaos.ChaosguardrulesGetRuleResponse{
		Name:         "rule-1",
		Description:  "desc",
		IsEnabled:    true,
		UserGroupIds: []string{"ug-1"},
		Tags:         []string{"t1"},
		ConditionIds: []string{"c-1", "c-2"},
		TimeWindows: []chaos.SecurityGovernanceTimeWindow{
			{
				TimeZone:  "UTC",
				StartTime: 100,
				EndTime:   200,
				Duration:  "1h",
				Recurrence: &chaos.SecurityGovernanceRecurrence{
					Type_: "Weekly",
					Spec: &chaos.SecurityGovernanceRecurrenceSpec{
						Until: 500,
						Value: 2,
					},
				},
			},
		},
	}

	d := schema.TestResourceDataRaw(t, ResourceChaosSecurityGovernanceRuleV3().Schema, map[string]interface{}{})
	if diags := setRuleV3Data(d, resp, "org", "proj"); diags.HasError() {
		t.Fatalf("setRuleV3Data returned diagnostics: %v", diags)
	}

	checks := map[string]interface{}{
		"name":                                    "rule-1",
		"is_enabled":                              true,
		"time_windows.0.time_zone":                "UTC",
		"time_windows.0.start_time":               100,
		"time_windows.0.end_time":                 200,
		"time_windows.0.duration":                 "1h",
		"time_windows.0.recurrence.0.type":        "Weekly",
		"time_windows.0.recurrence.0.until":       500,
		"time_windows.0.recurrence.0.value":       2,
	}
	for path, want := range checks {
		if got := d.Get(path); got != want {
			t.Errorf("%s: got %#v (%T), want %#v", path, got, got, want)
		}
	}
	if got := d.Get("condition_ids").([]interface{}); len(got) != 2 {
		t.Errorf("condition_ids drift: %#v", got)
	}
}
