package idp

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// kindMappingSchema is the mapping block on a kind: sent with the create, then written through the kind's mapping.
func kindMappingSchema() *schema.Schema {
	s := mappingWritableSchema()
	s["payload_schema"] = &schema.Schema{
		Description: "Custom types, custom kind: JSON Schema of the payload. Set at create only.",
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		DiffSuppressFunc: func(_, old, new string, _ *schema.ResourceData) bool {
			return jsonEquivalent(old, new)
		},
		ValidateFunc: func(v any, k string) ([]string, []error) {
			var m map[string]any
			if err := json.Unmarshal([]byte(v.(string)), &m); err != nil {
				return nil, []error{fmt.Errorf("%s must be a JSON object: %w", k, err)}
			}
			return nil, nil
		},
	}
	return block("The kind's mapping, sent with the create so the first sync uses it. Omitted fields keep the "+
		"type's defaults. Without the block the mapping isn't managed; removing the block resets it to the type's "+
		"default, except for Custom and HarnessCI kinds, which have none and keep theirs.", false, s)
}

// kindMappings returns each kind's mapping block values, for kinds that have one.
func kindMappings(kinds []any) map[string]map[string]any {
	out := map[string]map[string]any{}
	for _, k := range kinds {
		kind, _ := k.(map[string]any)
		if l, _ := kind["mapping"].([]any); len(l) > 0 {
			m, _ := l[0].(map[string]any)
			if m == nil {
				m = map[string]any{}
			}
			out[str(kind["kind"])] = m
		}
	}
	return out
}

func createMappings(kinds []any) map[string]any {
	out := map[string]any{}
	for kind, m := range kindMappings(kinds) {
		out[kind] = mappingBody(m, true)
	}
	return out
}

// mappingBody is a mapping write; payload_schema is only sent on create, as it can't change.
func mappingBody(values map[string]any, create bool) map[string]any {
	body := expandObject(mappingWritableSchema(), values)
	if raw := str(values["payload_schema"]); create && raw != "" {
		var schema map[string]any
		if json.Unmarshal([]byte(raw), &schema) == nil {
			body["payload_schema"] = schema
		}
	}
	return body
}

// kindMappingState leaves default field mappings out unless prior lists them, as omitting them means the defaults.
func kindMappingState(m, prior map[string]any) map[string]any {
	out := flattenObject(mappingWritableSchema(), m)
	if fms, _ := prior["field_mapping"].([]any); m["is_default_field_mappings"] == true && len(fms) == 0 {
		out["field_mapping"] = []any{}
	}
	out["payload_schema"] = jsonString(m["payload_schema"])
	return out
}

// mappingFieldsUnchanged rejects changes to create-only mapping fields; display_name_field is create-only on Airbyte.
func mappingFieldsUnchanged(oldKinds, newKinds []any, airbyte bool) error {
	fields := []string{"unique_identifier_field"}
	if airbyte {
		fields = append(fields, "display_name_field")
	}
	prior := kindMappings(oldKinds)
	for kind, m := range kindMappings(newKinds) {
		for _, field := range fields {
			if old, cur := str(prior[kind][field]), str(m[field]); old != "" && cur != "" && old != cur {
				return fmt.Errorf("%s of kind %s can't change from %q to %q; recreate the integration instead",
					field, kind, old, cur)
			}
		}
	}
	return nil
}

// payloadSchemasUnchanged rejects a changed payload_schema, which is fixed at create.
func payloadSchemasUnchanged(oldKinds, newKinds []any) error {
	prior := kindMappings(oldKinds)
	for kind, m := range kindMappings(newKinds) {
		p, ok := prior[kind]
		if !ok {
			continue
		}
		if old, cur := str(p["payload_schema"]), str(m["payload_schema"]); old != "" && cur != "" &&
			!jsonEquivalent(old, cur) {
			return fmt.Errorf("payload_schema of kind %s can't change; recreate the integration instead", kind)
		}
	}
	return nil
}

func jsonEquivalent(a, b string) bool {
	if a == b {
		return true
	}
	var x, y any
	if json.Unmarshal([]byte(a), &x) != nil || json.Unmarshal([]byte(b), &y) != nil {
		return false
	}
	return jsonString(x) == jsonString(y)
}
