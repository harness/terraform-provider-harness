package convert

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// getString gets a string value from a map and validates it
func getString(m map[string]interface{}, key string, required bool) (string, error) {
	v, ok := m[key]
	if !ok || v == nil {
		if required {
			return "", &ErrRequiredField{Field: key}
		}
		return "", nil
	}

	str, ok := v.(string)
	if !ok {
		return "", &ErrInvalidType{Expected: "string", Actual: v}
	}

	return strings.TrimSpace(str), nil
}

// getBool gets a bool value from a map with a default
func getBool(m map[string]interface{}, key string, def bool) (bool, error) {
	v, ok := m[key]
	if !ok || v == nil {
		return def, nil
	}

	b, ok := v.(bool)
	if !ok {
		return false, &ErrInvalidType{Expected: "bool", Actual: v}
	}

	return b, nil
}

// getInt gets an int value from a map with a default
func getInt(m map[string]interface{}, key string, def int) (int, error) {
	v, ok := m[key]
	if !ok || v == nil {
		return def, nil
	}

	switch val := v.(type) {
	case int:
		return val, nil
	case float64:
		return int(val), nil
	case string:
		i, err := strconv.Atoi(val)
		if err != nil {
			return 0, &ErrInvalidValue{Field: key, Value: val, Msg: "not a valid integer"}
		}
		return i, nil
	default:
		return 0, &ErrInvalidType{Expected: "int", Actual: v}
	}
}

// getStringSlice gets a string slice from a map
func getStringSlice(m map[string]interface{}, key string) ([]string, error) {
	v, ok := m[key]
	if !ok || v == nil {
		return nil, nil
	}

	slice, ok := v.([]interface{})
	if !ok {
		return nil, &ErrInvalidType{Expected: "[]interface{}", Actual: v}
	}

	result := make([]string, 0, len(slice))
	for i, item := range slice {
		str, ok := item.(string)
		if !ok {
			return nil, &ErrInvalidType{
				Expected: "string",
				Actual:   item,
				Msg:      fmt.Sprintf("at index %d", i),
			}
		}
		result = append(result, str)
	}

	return result, nil
}

// toInterfaceSlice converts a slice of any type to []interface{}
func toInterfaceSlice(slice interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected a slice, got %T", slice)
	}

	result := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = v.Index(i).Interface()
	}

	return result, nil
}

// setValue sets a value in the resource data, handling errors
func setValue(d *schema.ResourceData, key string, value interface{}) diag.Diagnostics {
	if err := d.Set(key, value); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set %s: %w", key, err))
	}
	return nil
}
