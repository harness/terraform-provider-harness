package helpers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v3"
)

// runtimeInputToken is the bare Harness runtime-input expression. The backend may echo it
// back decorated with a suffix such as .default(...)/.allowedValues(...)/.executionInput()
// (e.g. when the referenced template declares a default for that field), which is not a
// real configuration change.
const runtimeInputToken = "<+input>"

// YamlDiffSuppressFunction returns true if two content of yaml strings are identical.
// That helps to avoid unnecessary changes in plan if the yaml format was changed only, but not the data.
func YamlDiffSuppressFunction(k, old, new string, d *schema.ResourceData) bool {
	var oldYaml, newYaml interface{}
	if err := yaml.Unmarshal([]byte(old), &oldYaml); err != nil {
		return false
	}
	if err := yaml.Unmarshal([]byte(new), &newYaml); err != nil {
		return false
	}
	return yamlValuesEqual(oldYaml, newYaml)
}

// yamlValuesEqual recursively compares two values decoded from YAML, treating a bare
// "<+input>" runtime-input expression as equal to the same expression decorated by the
// backend with a suffix (e.g. "<+input>.default(foo)"). Everything else falls back to
// strict equality, so this never masks real value or ordering changes.
func yamlValuesEqual(a, b interface{}) bool {
	switch av := a.(type) {
	case string:
		bv, ok := b.(string)
		if !ok {
			return false
		}
		if av == bv {
			return true
		}
		return isRuntimeInputVariant(av, bv) || isRuntimeInputVariant(bv, av)
	case map[string]interface{}:
		bv, ok := b.(map[string]interface{})
		if !ok || len(av) != len(bv) {
			return false
		}
		for key, aval := range av {
			bval, ok := bv[key]
			if !ok || !yamlValuesEqual(aval, bval) {
				return false
			}
		}
		return true
	case []interface{}:
		bv, ok := b.([]interface{})
		if !ok || len(av) != len(bv) {
			return false
		}
		for i := range av {
			if !yamlValuesEqual(av[i], bv[i]) {
				return false
			}
		}
		return true
	default:
		return reflect.DeepEqual(a, b)
	}
}

// isRuntimeInputVariant returns true if bare is the plain "<+input>" token and decorated is
// that same token with a non-empty suffix expression, e.g. "<+input>.default(foo)".
func isRuntimeInputVariant(bare, decorated string) bool {
	return bare == runtimeInputToken &&
		strings.HasPrefix(decorated, runtimeInputToken) &&
		decorated != runtimeInputToken
}

// YamlResourceFields contains common fields extracted from a resource's YAML definition.
type YamlResourceFields struct {
	Identifier  string
	Name        string
	Description string
}

// ExtractYamlResourceFields extracts common resource fields (identifier, name, description)
// from a YAML string. The rootKey parameter specifies the top-level key in the YAML
// (e.g., "infrastructureDefinition", "environment").
func ExtractYamlResourceFields(yamlString string, rootKey string) (YamlResourceFields, error) {
	var yamlData map[string]any
	if err := yaml.Unmarshal([]byte(yamlString), &yamlData); err != nil {
		return YamlResourceFields{}, fmt.Errorf("failed to parse YAML: %w", err)
	}

	resourceRaw, ok := yamlData[rootKey]
	if !ok {
		return YamlResourceFields{}, fmt.Errorf("YAML must contain a '%s' key", rootKey)
	}

	resourceMap, ok := resourceRaw.(map[string]any)
	if !ok {
		return YamlResourceFields{}, fmt.Errorf("'%s' key must be a map", rootKey)
	}

	identifier, _ := resourceMap["identifier"].(string)
	name, _ := resourceMap["name"].(string)
	description, _ := resourceMap["description"].(string)

	if identifier == "" {
		return YamlResourceFields{}, fmt.Errorf("YAML must contain '%s.identifier' field", rootKey)
	}
	if name == "" {
		return YamlResourceFields{}, fmt.Errorf("YAML must contain '%s.name' field", rootKey)
	}

	return YamlResourceFields{
		Identifier:  identifier,
		Name:        name,
		Description: description,
	}, nil
}

// ValidateYamlFieldsMatch validates that identifier and name in the schema match those in the YAML,
// and sets all fields (identifier, name, description) from the YAML when not explicitly provided.
// The resourceType is used for error messages (e.g., "infrastructure", "environment").
func ValidateYamlFieldsMatch(d *schema.ResourceDiff, yamlFields YamlResourceFields, resourceType string) error {
	// Validate and set identifier
	if v, ok := d.GetOk("identifier"); ok && v.(string) != "" {
		if v.(string) != yamlFields.Identifier {
			return fmt.Errorf("identifier in schema (%s) does not match identifier in %s YAML (%s)",
				v.(string), resourceType, yamlFields.Identifier)
		}
	} else {
		if err := d.SetNew("identifier", yamlFields.Identifier); err != nil {
			return err
		}
	}

	// Validate and set name
	if v, ok := d.GetOk("name"); ok && v.(string) != "" {
		if v.(string) != yamlFields.Name {
			return fmt.Errorf("name in schema (%s) does not match name in %s YAML (%s)",
				v.(string), resourceType, yamlFields.Name)
		}
	} else {
		if err := d.SetNew("name", yamlFields.Name); err != nil {
			return err
		}
	}

	// Always set description from YAML since it's the source of truth.
	// Unlike identifier and name which may be explicitly set in HCL,
	// description is typically only defined in the YAML.
	if err := d.SetNew("description", yamlFields.Description); err != nil {
		return err
	}

	return nil
}
