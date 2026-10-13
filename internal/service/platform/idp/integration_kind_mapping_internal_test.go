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

func TestMappingFieldsUnchanged(t *testing.T) {
	kinds := func(uid, name string) []any {
		return []any{map[string]any{"kind": "incident", "mapping": []any{map[string]any{
			"unique_identifier_field": uid, "display_name_field": name}}}}
	}
	old := kinds("id", "title")
	if err := mappingFieldsUnchanged(old, kinds("", ""), true); err != nil {
		t.Errorf("omitted fields: %v", err)
	}
	if err := mappingFieldsUnchanged(old, kinds("id", "name"), false); err != nil {
		t.Errorf("display_name_field on a platform type: %v", err)
	}
	if err := mappingFieldsUnchanged(old, kinds("id", "name"), true); err == nil ||
		!strings.Contains(err.Error(), "display_name_field of kind incident") {
		t.Errorf("display_name_field on Airbyte: %v", err)
	}
	if err := mappingFieldsUnchanged(old, kinds("key", "title"), false); err == nil ||
		!strings.Contains(err.Error(), "unique_identifier_field of kind incident") {
		t.Errorf("unique_identifier_field: %v", err)
	}
}

func TestSwitchingToMerge(t *testing.T) {
	kinds := func(a, b string) []any {
		return []any{map[string]any{"kind": "a", "action": a}, map[string]any{"kind": "b", "action": b}}
	}
	got := switchingToMerge(kinds("merge", "register"), kinds("merge", "merge"))
	if !reflect.DeepEqual(got, map[string]bool{"b": true}) {
		t.Errorf("switchingToMerge = %v", got)
	}
}
