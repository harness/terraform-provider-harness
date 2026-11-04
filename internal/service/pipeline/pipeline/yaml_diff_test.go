package pipeline

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func resourceDataWithIgnoreRuntimeInputDefaults(t *testing.T, enabled bool) *schema.ResourceData {
	t.Helper()
	return schema.TestResourceDataRaw(t, ResourcePipeline().Schema, map[string]interface{}{
		ignoreRuntimeInputDefaultsAttr: enabled,
	})
}

// TestPipelineYamlDiffSuppressFunction_Baseline mirrors helpers.YamlDiffSuppressFunction's
// own test cases - map key reordering is ignored, real structural changes are not - to make
// sure scoping this to the pipeline resource didn't lose any existing behavior.
// These cases must hold with the opt-in flag off (default) and on.
func TestPipelineYamlDiffSuppressFunction_Baseline(t *testing.T) {
	for _, enabled := range []bool{false, true} {
		d := resourceDataWithIgnoreRuntimeInputDefaults(t, enabled)

		require.True(t, pipelineYamlDiffSuppressFunction("", "", "", d))

		require.True(t, pipelineYamlDiffSuppressFunction("",
			`---
field1: value1
field2: value2
`,
			`---
field2: value2
field1: value1
`, d))

		require.False(t, pipelineYamlDiffSuppressFunction("",
			`---
field1: value1
`,
			`---
field2: value2
`, d))

		require.False(t, pipelineYamlDiffSuppressFunction("",
			`---
field1: value1
field2: value2
`,
			`---
field1: value1
`, d))
	}
}

// TestPipelineYamlDiffSuppressFunction_RuntimeInputDefault_FlagOff is the default /
// regression-safe behavior: bare "<+input>" vs backend decoration is a real plan diff.
func TestPipelineYamlDiffSuppressFunction_RuntimeInputDefault_FlagOff(t *testing.T) {
	d := resourceDataWithIgnoreRuntimeInputDefaults(t, false)

	require.False(t, pipelineYamlDiffSuppressFunction("",
		`---
pipeline:
  template:
    templateInputs:
      variables:
        - name: TECH_STACK
          type: String
          value: <+input>.default(python)
`,
		`---
pipeline:
  template:
    templateInputs:
      variables:
        - name: TECH_STACK
          type: String
          value: <+input>
`, d))
}

// TestPipelineYamlDiffSuppressFunction_RuntimeInputDefault_FlagOn covers PIPE-38416 when
// the customer opts in: state carries the backend decoration while config keeps the bare
// token, and the plan must stay clean.
func TestPipelineYamlDiffSuppressFunction_RuntimeInputDefault_FlagOn(t *testing.T) {
	d := resourceDataWithIgnoreRuntimeInputDefaults(t, true)

	require.True(t, pipelineYamlDiffSuppressFunction("",
		// old: state, as returned by the backend
		`---
pipeline:
  template:
    templateInputs:
      properties:
        ci:
          codebase:
            repoName: <+input>.allowedValues(epy-bassam)
            build: <+input>.default(ci)
      variables:
        - name: TECH_STACK
          type: String
          value: <+input>.default(python)
`,
		// new: the customer's config, left as bare runtime inputs
		`---
pipeline:
  template:
    templateInputs:
      properties:
        ci:
          codebase:
            repoName: <+input>
            build: <+input>
      variables:
        - name: TECH_STACK
          type: String
          value: <+input>
`, d))

	// Reverse direction must NOT be suppressed: config explicitly overriding with a
	// decorated expression while state has the bare token is a real, intentional change.
	require.False(t, pipelineYamlDiffSuppressFunction("",
		`---
field1: <+input>
`,
		`---
field1: <+input>.default(explicit-override)
`, d))

	// A field that was never a runtime input must still report a real diff.
	require.False(t, pipelineYamlDiffSuppressFunction("",
		`---
field1: <+input>
`,
		`---
field1: concrete-value
`, d))

	// A real content change on a non-runtime-input field inside the same structure
	// must still be caught - this isn't blanket string-mismatch tolerance.
	require.False(t, pipelineYamlDiffSuppressFunction("",
		`---
variables:
  - name: A
    value: <+input>.default(a)
  - name: B
    value: different-value
`,
		`---
variables:
  - name: A
    value: <+input>
  - name: B
    value: concrete-value
`, d))

	// A string that only shares the "<+input>" prefix, with no "." suffix, is not a
	// backend decoration and must still show up as a diff.
	require.False(t, pipelineYamlDiffSuppressFunction("",
		`---
field1: <+input>foo
`,
		`---
field1: <+input>
`, d))

	// Two different decorations are a real edit, not the bare-token echo.
	require.False(t, pipelineYamlDiffSuppressFunction("",
		`---
field1: <+input>.default(python)
`,
		`---
field1: <+input>.default(java)
`, d))

	// executionInput() is another backend suffix and follows the same rule.
	require.True(t, pipelineYamlDiffSuppressFunction("",
		`---
field1: <+input>.executionInput()
`,
		`---
field1: <+input>
`, d))
}
