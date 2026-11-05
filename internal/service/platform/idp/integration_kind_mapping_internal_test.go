package idp

import (
	"reflect"
	"strings"
	"testing"
)

func TestKindMappingBodies(t *testing.T) {
	kinds := []any{
		map[string]any{"kind": "custom", "mapping": []any{map[string]any{
			"ingestion":      []any{map[string]any{"identifier": ".id", "name": "", "properties": map[string]any{}}},
			"correlation":    []any{map[string]any{"source_path": "svc", "target_path": "metadata.name", "operator": ""}},
			"payload_schema": `{"type":"object","properties":{"id":{"type":"string"}}}`,
			"field_mapping":  []any{},
		}}},
		map[string]any{"kind": "other", "mapping": []any{}},
	}
	want := map[string]any{"custom": map[string]any{
		"ingestion":      map[string]any{"identifier": ".id"},
		"correlation":    map[string]any{"source_path": "svc", "target_path": "metadata.name"},
		"payload_schema": map[string]any{"type": "object", "properties": map[string]any{"id": map[string]any{"type": "string"}}},
	}}
	if got := createMappings(kinds); !reflect.DeepEqual(got, want) {
		t.Errorf("createMappings = %v, want %v", got, want)
	}
}

func TestMappingsUnchanged(t *testing.T) {
	kinds := func(schema, uid string) []any {
		return []any{map[string]any{"kind": "custom", "mapping": []any{map[string]any{
			"payload_schema": schema, "unique_identifier_field": uid}}}}
	}
	old := kinds(`{"type":"object","properties":{"id":{}}}`, "id")
	if err := mappingsUnchanged(old, kinds(`{"properties": {"id": {}}, "type": "object"}`, "id")); err != nil {
		t.Errorf("reformatted schema: %v", err)
	}
	if err := mappingsUnchanged(old, []any{map[string]any{"kind": "custom"}}); err != nil {
		t.Errorf("removed block: %v", err)
	}
	for name, kinds := range map[string][]any{
		"changed schema": kinds(`{"type":"object","properties":{"name":{}}}`, "id"),
		"changed field":  kinds(`{"type":"object","properties":{"id":{}}}`, "key"),
	} {
		if err := mappingsUnchanged(old, kinds); err == nil ||
			!strings.Contains(err.Error(), "harness_platform_idp_integration_kind_mapping") {
			t.Errorf("%s: %v", name, err)
		}
	}
	if err := mappingsUnchanged([]any{map[string]any{"kind": "custom"}}, old); err == nil {
		t.Errorf("added block: no error")
	}
}
