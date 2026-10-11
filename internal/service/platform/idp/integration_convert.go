package idp

import (
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// The integrations API uses snake_case, so attributes carry their JSON names, with these exceptions.
var (
	// blockJSONNames maps repeated blocks, named in the singular for HCL, to their API arrays.
	blockJSONNames = map[string]string{"kind": "kinds", "column": "columns", "field_mapping": "field_mappings"}
	// emptyObjectFlags are bools that stand for an empty API object, e.g. source_repo = true for {}.
	emptyObjectFlags = map[string]bool{"source_repo": true}
	// optionalBools are strings ("true", "false" or unset) for API booleans where unset differs from false.
	optionalBools = map[string]bool{"resolved": true}
)

func jsonName(key string, s *schema.Schema) string {
	if _, ok := s.Elem.(*schema.Resource); ok && s.Type == schema.TypeList && s.MaxItems != 1 {
		if n, ok := blockJSONNames[key]; ok {
			return n
		}
	}
	return key
}

// expandObject converts Terraform values for the attributes in m into an API object, leaving out empty values.
func expandObject(m map[string]*schema.Schema, values map[string]any) map[string]any {
	out := map[string]any{}
	for key, s := range m {
		if s.Computed && !s.Optional {
			continue
		}
		if v, ok := expandValue(key, s, values[key]); ok {
			out[jsonName(key, s)] = v
		}
	}
	return out
}

func expandValue(key string, s *schema.Schema, v any) (any, bool) {
	if v == nil {
		return nil, false
	}
	switch s.Type {
	case schema.TypeString:
		str, _ := v.(string)
		if str == "" {
			return nil, false
		}
		if optionalBools[key] {
			b, err := strconv.ParseBool(str)
			return b, err == nil
		}
		return str, true
	case schema.TypeBool:
		b, _ := v.(bool)
		if emptyObjectFlags[key] {
			return map[string]any{}, b
		}
		return b, b
	case schema.TypeInt:
		i, _ := v.(int)
		return i, i != 0
	case schema.TypeMap:
		m, _ := v.(map[string]any)
		return m, len(m) > 0
	case schema.TypeList, schema.TypeSet:
		items := listItems(v)
		res, isBlock := s.Elem.(*schema.Resource)
		if !isBlock {
			return items, len(items) > 0
		}
		objects := make([]any, 0, len(items))
		for _, item := range items {
			values, _ := item.(map[string]any)
			objects = append(objects, expandObject(res.Schema, values))
		}
		if s.MaxItems == 1 {
			if len(objects) == 0 {
				return nil, false
			}
			return objects[0], true
		}
		return objects, len(objects) > 0
	}
	return nil, false
}

func listItems(v any) []any {
	switch l := v.(type) {
	case []any:
		return l
	case *schema.Set:
		return l.List()
	}
	return nil
}

// flattenObject converts an API object into Terraform values for the attributes in m.
func flattenObject(m map[string]*schema.Schema, obj map[string]any) map[string]any {
	out := make(map[string]any, len(m))
	for key, s := range m {
		if v, ok := flattenValue(key, s, obj[jsonName(key, s)]); ok {
			out[key] = v
		}
	}
	return out
}

func flattenValue(key string, s *schema.Schema, v any) (any, bool) {
	switch s.Type {
	case schema.TypeString:
		switch t := v.(type) {
		case string:
			return t, true
		case bool:
			return strconv.FormatBool(t), true
		case float64:
			return strconv.FormatFloat(t, 'f', -1, 64), true
		}
		return "", v == nil
	case schema.TypeBool:
		if emptyObjectFlags[key] {
			return v != nil, true
		}
		b, _ := v.(bool)
		return b, true
	case schema.TypeInt:
		f, _ := v.(float64)
		return int(f), true
	case schema.TypeMap:
		m, _ := v.(map[string]any)
		return m, true
	case schema.TypeList, schema.TypeSet:
		res, isBlock := s.Elem.(*schema.Resource)
		if isBlock && s.MaxItems == 1 {
			obj, ok := v.(map[string]any)
			if !ok {
				return []any{}, true
			}
			return []any{flattenObject(res.Schema, obj)}, true
		}
		items, _ := v.([]any)
		out := make([]any, 0, len(items))
		for _, item := range items {
			if isBlock {
				obj, _ := item.(map[string]any)
				out = append(out, flattenObject(res.Schema, obj))
			} else {
				out = append(out, item)
			}
		}
		return out, true
	}
	return nil, false
}

// expandTags splits "key:value" on the first colon, as the API stores a map.
func expandTags(tags []any) map[string]any {
	out := make(map[string]any, len(tags))
	for _, t := range tags {
		k, v := splitTag(t.(string))
		out[k] = v
	}
	return out
}

func splitTag(tag string) (string, string) {
	for i := 0; i < len(tag); i++ {
		if tag[i] == ':' {
			return tag[:i], tag[i+1:]
		}
	}
	return tag, ""
}

func flattenTags(tags map[string]any) []string {
	out := make([]string, 0, len(tags))
	for k, v := range tags {
		if s, _ := v.(string); s != "" {
			out = append(out, k+":"+s)
		} else {
			out = append(out, k)
		}
	}
	sort.Strings(out)
	return out
}
