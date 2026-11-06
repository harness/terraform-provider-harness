package pipeline

import (
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v3"
)

// runtimeInputToken is the bare Harness runtime-input expression. When a pipeline is
// created from a template reference whose templateInputs are left as a bare "<+input>",
// the backend can echo that field back in state decorated with the referenced template's
// own declared default/allowed-values expression, e.g. "<+input>.default(python)". See
// PIPE-38416. That decoration is not a real configuration change for customers who opt in
// via ignore_runtime_input_defaults.
//
// This is intentionally a pipeline-local DiffSuppressFunc rather than a change to the
// shared helpers.YamlDiffSuppressFunction, which ~14 other resources also rely on - a
// mistake there would have a much wider blast radius than this one field.
const (
	runtimeInputToken              = "<+input>"
	ignoreRuntimeInputDefaultsAttr = "ignore_runtime_input_defaults"
)

// pipelineYamlDiffSuppressFunction compares pipeline "yaml" state vs config. Map-key
// reordering is always ignored (same as helpers.YamlDiffSuppressFunction). Bare
// "<+input>" vs backend-decorated "<+input>.default(...)" / ".allowedValues(...)" is
// ignored only when ignore_runtime_input_defaults is true. Default is false so existing
// pipelines keep strict drift detection.
func pipelineYamlDiffSuppressFunction(k, old, new string, d *schema.ResourceData) bool {
	var oldYaml, newYaml interface{}
	if err := yaml.Unmarshal([]byte(old), &oldYaml); err != nil {
		return false
	}
	if err := yaml.Unmarshal([]byte(new), &newYaml); err != nil {
		return false
	}
	return pipelineYamlValuesEqual(oldYaml, newYaml, ignoreRuntimeInputDefaultsEnabled(d))
}

func ignoreRuntimeInputDefaultsEnabled(d *schema.ResourceData) bool {
	if d == nil {
		return false
	}
	enabled, ok := d.Get(ignoreRuntimeInputDefaultsAttr).(bool)
	return ok && enabled
}

// pipelineYamlValuesEqual recursively compares a value decoded from the "old" (state) YAML
// against the corresponding value from the "new" (config) YAML. old and new keep their
// state/config identity through the recursion, which is what makes the runtime-input
// direction check below safe: the reverse (config decorated, state bare) is NOT treated as
// equal, because a user explicitly writing a decorated expression in config (e.g.
// overriding the template's default) is a real, intentional change.
func pipelineYamlValuesEqual(old, new interface{}, ignoreRuntimeInputDefaults bool) bool {
	switch oldVal := old.(type) {
	case string:
		newVal, ok := new.(string)
		if !ok {
			return false
		}
		if oldVal == newVal {
			return true
		}
		if !ignoreRuntimeInputDefaults {
			return false
		}
		// newVal is config, oldVal is state. Suppress only when config is still the bare
		// token and state is that token plus a backend suffix. Those suffixes always
		// start with "." (`<+input>.default(...)`, `.allowedValues(...)`, `.executionInput()`).
		// A value that merely starts with "<+input>" (for example "<+input>foo") is a real change.
		return newVal == runtimeInputToken && strings.HasPrefix(oldVal, runtimeInputToken+".")
	case map[string]interface{}:
		newVal, ok := new.(map[string]interface{})
		if !ok || len(oldVal) != len(newVal) {
			return false
		}
		for key, oldChild := range oldVal {
			newChild, ok := newVal[key]
			if !ok || !pipelineYamlValuesEqual(oldChild, newChild, ignoreRuntimeInputDefaults) {
				return false
			}
		}
		return true
	case []interface{}:
		newVal, ok := new.([]interface{})
		if !ok || len(oldVal) != len(newVal) {
			return false
		}
		for i := range oldVal {
			if !pipelineYamlValuesEqual(oldVal[i], newVal[i], ignoreRuntimeInputDefaults) {
				return false
			}
		}
		return true
	default:
		return reflect.DeepEqual(old, new)
	}
}
