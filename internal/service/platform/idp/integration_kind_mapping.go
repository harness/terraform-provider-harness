package idp

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// kindMappingSchema is the mapping block on a kind: sent with the create only. It's kept in state as configured and
// never read back; harness_platform_idp_integration_kind_mapping manages a mapping after the create.
func kindMappingSchema() *schema.Schema {
	s := mappingWritableSchema()
	s["payload_schema"] = &schema.Schema{
		Description: "Custom types, custom kind: JSON Schema of the payload.",
		Type:        schema.TypeString,
		Optional:    true,
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
	notComputed(s)
	return block("The kind's mapping, sent with the create so the first sync uses it. Omitted fields keep the "+
		"type's defaults. Create only: it can't be added or changed afterwards, and removing it leaves the mapping "+
		"as it is. Use harness_platform_idp_integration_kind_mapping to manage the mapping after the create, and "+
		"don't use both for the same kind.", false, s)
}

// notComputed makes a schema plain Optional, as a create-only block holds only what the config set.
func notComputed(m map[string]*schema.Schema) {
	for _, s := range m {
		s.Computed = false
		if r, ok := s.Elem.(*schema.Resource); ok {
			notComputed(r.Schema)
		}
	}
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
		out[kind] = mappingBody(m)
	}
	return out
}

func mappingBody(values map[string]any) map[string]any {
	body := expandObject(mappingWritableSchema(), values)
	if raw := str(values["payload_schema"]); raw != "" {
		var schema map[string]any
		if json.Unmarshal([]byte(raw), &schema) == nil {
			body["payload_schema"] = schema
		}
	}
	return body
}

// mappingsUnchanged rejects a mapping block added or changed on an existing integration; removing one is allowed.
func mappingsUnchanged(oldKinds, newKinds []any) error {
	prior := kindMappings(oldKinds)
	for kind, m := range kindMappings(newKinds) {
		if p, ok := prior[kind]; ok && jsonEquivalent(jsonString(mappingBody(p)), jsonString(mappingBody(m))) {
			continue
		}
		return fmt.Errorf("the mapping block of kind %s is only sent at create; to change the mapping of an "+
			"existing integration, remove the block and use harness_platform_idp_integration_kind_mapping", kind)
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
