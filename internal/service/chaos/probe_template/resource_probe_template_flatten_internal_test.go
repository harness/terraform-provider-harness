package probe_template

// White-box read-side (flatten) test for probe template variables. The API
// returns variable values as *interface{}; if a non-string scalar is not
// coerced to string, d.Set fails and (because the error is ignored) the
// variables block is silently dropped, producing a perpetual diff.

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ptrIface(v interface{}) *interface{} { return &v }

func TestProbeTemplate_SetVariables_ValueRoundTrip(t *testing.T) {
	tmpl := &chaos.GithubComHarnessHceSaasGraphqlServerPkgDatabaseMongodbChaosprobetemplateChaosProbeTemplate{
		Identity: "p",
		Name:     "p",
		Variables: []chaos.TemplateVariable{
			{Name: "STR_VAR", Value: ptrIface("hello")},
			{Name: "NUM_VAR", Value: ptrIface(float64(240))},
			{Name: "RUNTIME_VAR", Value: ptrIface("<+input>")},
		},
	}

	d := schema.TestResourceDataRaw(t, ResourceProbeTemplate().Schema, map[string]interface{}{})
	if diags := setProbeTemplateDataSimplified(d, tmpl, "acc", "", "", "hub"); diags.HasError() {
		t.Fatalf("setProbeTemplateDataSimplified returned diagnostics: %v", diags)
	}

	want := map[string]string{
		"STR_VAR":     "hello",
		"NUM_VAR":     "240",
		"RUNTIME_VAR": "<+input>",
	}
	list := d.Get("variables").([]interface{})
	if len(list) != len(tmpl.Variables) {
		t.Fatalf("expected %d variables in state, got %d (non-string values likely dropped)", len(tmpl.Variables), len(list))
	}
	for _, item := range list {
		m := item.(map[string]interface{})
		name := m["name"].(string)
		if got := m["value"]; got != want[name] {
			t.Errorf("variable %q value drift: got %#v (%T), want %q", name, got, got, want[name])
		}
	}
}
