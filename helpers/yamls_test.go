package helpers_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/stretchr/testify/require"
)

func TestYamlDiffSuppressFunction(t *testing.T) {
	require.True(t, helpers.YamlDiffSuppressFunction("", "", "", nil))

	require.True(t, helpers.YamlDiffSuppressFunction("",
		`---
field1: value1
field2: value2
`,
		`---
field2: value2
field1: value1
`, nil))

	require.True(t, helpers.YamlDiffSuppressFunction("",
		`---
"field1": "value1"
"field2": "value2"
`,
		`---
field2: value2
field1: value1
`, nil))

	require.False(t, helpers.YamlDiffSuppressFunction("",
		`---
field1: value1
`,
		`---
field2: value2
`, nil))

	require.False(t, helpers.YamlDiffSuppressFunction("",
		`---
field1: value1
`,
		`---
field1: value1
field2: value2
`, nil))

	require.False(t, helpers.YamlDiffSuppressFunction("",
		`---
field1: value1
field2: value2
`,
		`---
field1: value1
`, nil))

}

// TestYamlDiffSuppressFunction_RuntimeInputDefault covers PIPE-38416: a pipeline created
// from a template with templateInputs left as a bare "<+input>" runtime input (as in the
// customer's main.tf) reports a reconcile diff because the backend echoes the field back
// decorated with the template's declared default, e.g. "<+input>.default(python)".
func TestYamlDiffSuppressFunction_RuntimeInputDefault(t *testing.T) {
	require.True(t, helpers.YamlDiffSuppressFunction("",
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
`,
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
`, nil))

	// A field that was never a runtime input must still report a real diff.
	require.False(t, helpers.YamlDiffSuppressFunction("",
		`---
field1: <+input>
`,
		`---
field1: concrete-value
`, nil))

	// A real content change on a non-runtime-input field inside the same structure
	// must still be caught - this isn't blanket string-mismatch tolerance.
	require.False(t, helpers.YamlDiffSuppressFunction("",
		`---
variables:
  - name: A
    value: <+input>
  - name: B
    value: concrete-value
`,
		`---
variables:
  - name: A
    value: <+input>.default(a)
  - name: B
    value: different-value
`, nil))
}
