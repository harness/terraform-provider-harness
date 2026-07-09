package fault_template

// White-box test for flattenTemplateVariables: API variable values (*interface{})
// of any scalar type must flatten to the schema's string value without being
// dropped (which would cause a perpetual diff).

import (
	"testing"

	"github.com/harness/harness-go-sdk/harness/chaos"
)

func ptrIface(v interface{}) *interface{} { return &v }

func TestFlattenTemplateVariables_ValueCoercion(t *testing.T) {
	vars := []chaos.TemplateVariable{
		{Name: "STR_VAR", Value: ptrIface("hello"), Required: true},
		{Name: "NUM_VAR", Value: ptrIface(float64(240))},
		{Name: "BOOL_VAR", Value: ptrIface(true)},
		{Name: "RUNTIME_VAR", Value: ptrIface("<+input>")},
	}

	out := flattenTemplateVariables(vars)
	if len(out) != len(vars) {
		t.Fatalf("expected %d flattened variables, got %d", len(vars), len(out))
	}

	want := map[string]string{
		"STR_VAR":     "hello",
		"NUM_VAR":     "240",
		"BOOL_VAR":    "true",
		"RUNTIME_VAR": "<+input>",
	}
	for _, m := range out {
		name := m["name"].(string)
		got, ok := m["value"]
		if !ok {
			t.Errorf("variable %q: value key missing (would drop to empty and drift)", name)
			continue
		}
		if got != want[name] {
			t.Errorf("variable %q value: got %#v (%T), want %q", name, got, got, want[name])
		}
	}
}
