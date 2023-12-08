package helpers

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v3"
)

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
	return reflect.DeepEqual(oldYaml, newYaml)

}
