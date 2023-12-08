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
