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
	if got := mappingBody(kindMappings(kinds)["custom"], false); got["payload_schema"] != nil {
		t.Errorf("update body sends payload_schema: %v", got)
	}
}

func TestPayloadSchemasUnchanged(t *testing.T) {
	kinds := func(schema string) []any {
		return []any{map[string]any{"kind": "custom", "mapping": []any{map[string]any{"payload_schema": schema}}}}
	}
	old := kinds(`{"type":"object","properties":{"id":{}}}`)
	if err := payloadSchemasUnchanged(old, kinds(`{"properties": {"id": {}}, "type": "object"}`)); err != nil {
		t.Errorf("reformatted schema: %v", err)
	}
	if err := payloadSchemasUnchanged(old, kinds("")); err != nil {
		t.Errorf("omitted schema: %v", err)
	}
	err := payloadSchemasUnchanged(old, kinds(`{"type":"object","properties":{"name":{}}}`))
	if err == nil || !strings.Contains(err.Error(), "payload_schema of kind custom can't change") {
		t.Errorf("changed schema: %v", err)
	}
}
