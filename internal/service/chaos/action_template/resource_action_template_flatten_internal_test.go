package action_template

// White-box read-side (flatten) tests for the harness_chaos_action_template
// resource. These reproduce the "field read incorrectly -> perpetual diff /
// apply error" class of customer issues deterministically (no live API).

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ptrIface(v interface{}) *interface{} { return &v }

// TestSetVariablesData_ValueRoundTrip verifies that template variable values
// returned by the API (typed as *interface{}) are flattened to the string the
// Terraform schema expects. The API echoes back the value the provider sent
// (always a string), so the state must contain that string verbatim.
func TestSetVariablesData_ValueRoundTrip(t *testing.T) {
	vars := []chaos.TemplateVariable{
		{Name: "STR_VAR", Value: ptrIface("hello"), Required: true},
		{Name: "NUM_VAR", Value: ptrIface(float64(240))},
		{Name: "RUNTIME_VAR", Value: ptrIface("<+input>")},
	}

	d := schema.TestResourceDataRaw(t, ResourceActionTemplate().Schema, map[string]interface{}{})
	if err := setVariablesData(d, vars); err != nil {
		t.Fatalf("setVariablesData returned error: %v", err)
	}

	want := map[string]string{
		"STR_VAR":     "hello",
		"NUM_VAR":     "240",
		"RUNTIME_VAR": "<+input>",
	}

	list := d.Get("variables").([]interface{})
	if len(list) != len(vars) {
		t.Fatalf("expected %d variables in state, got %d", len(vars), len(list))
	}
	for _, item := range list {
		m := item.(map[string]interface{})
		name := m["name"].(string)
		if got := m["value"]; got != want[name] {
			t.Errorf("variable %q value drift: got %#v (%T), want %q", name, got, got, want[name])
		}
	}
}

// TestSetRunPropertiesData_RoundTrip verifies run_properties string/bool/int
// fields flatten back into state without drift.
func TestSetRunPropertiesData_RoundTrip(t *testing.T) {
	retries := interface{}(float64(3))
	props := &chaos.ActionActionTemplateRunProperties{
		InitialDelay:  "5s",
		Interval:      "10s",
		Timeout:       "60s",
		Verbosity:     "info",
		StopOnFailure: true,
		MaxRetries:    &retries,
	}

	d := schema.TestResourceDataRaw(t, ResourceActionTemplate().Schema, map[string]interface{}{})
	if err := setRunPropertiesData(d, props); err != nil {
		t.Fatalf("setRunPropertiesData returned error: %v", err)
	}

	checks := map[string]interface{}{
		"run_properties.0.initial_delay":   "5s",
		"run_properties.0.interval":        "10s",
		"run_properties.0.timeout":         "60s",
		"run_properties.0.verbosity":       "info",
		"run_properties.0.stop_on_failure": true,
		"run_properties.0.max_retries":     3,
	}
	for path, want := range checks {
		if got := d.Get(path); got != want {
			t.Errorf("%s: got %#v (%T), want %#v", path, got, got, want)
		}
	}
}
