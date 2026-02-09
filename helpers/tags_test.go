package helpers

import (
	"reflect"
	"testing"
)

func TestExpandPipelineTags(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected map[string]string
	}{
		{
			name:  "simple key:value tags",
			input: []interface{}{"env:prod", "version:1.0"},
			expected: map[string]string{
				"env":     "prod",
				"version": "1.0",
			},
		},
		{
			name:  "tag with no colon",
			input: []interface{}{"standalone"},
			expected: map[string]string{
				"standalone": "",
			},
		},
		{
			name:  "tag with empty value",
			input: []interface{}{"key:"},
			expected: map[string]string{
				"key": "",
			},
		},
		{
			name:  "PIPE-30810: Harness expression with ternary operator and colons",
			input: []interface{}{"ImagePush:<+<+pipeline.variables.variableA>==\"yes\"?<+pipeline.variables.tag>:<+pipeline.name>>"},
			expected: map[string]string{
				"ImagePush": "<+<+pipeline.variables.variableA>==\"yes\"?<+pipeline.variables.tag>:<+pipeline.name>>",
			},
		},
		{
			name:  "tag with URL containing port",
			input: []interface{}{"registry:https://registry.example.com:5000/repo"},
			expected: map[string]string{
				"registry": "https://registry.example.com:5000/repo",
			},
		},
		{
			name:  "tag with timestamp",
			input: []interface{}{"created:2024:01:15:10:30:00"},
			expected: map[string]string{
				"created": "2024:01:15:10:30:00",
			},
		},
		{
			name:  "multiple colons in value",
			input: []interface{}{"complex:a:b:c:d:e"},
			expected: map[string]string{
				"complex": "a:b:c:d:e",
			},
		},
		{
			name: "mixed tags - real world NAB scenario",
			input: []interface{}{
				"env:prod",
				"version:1.0.0",
				"ImagePush:<+<+pipeline.variables.variableA>==\"yes\"?<+pipeline.variables.tag>:<+pipeline.name>>",
				"service:api-gateway",
			},
			expected: map[string]string{
				"env":       "prod",
				"version":   "1.0.0",
				"ImagePush": "<+<+pipeline.variables.variableA>==\"yes\"?<+pipeline.variables.tag>:<+pipeline.name>>",
				"service":   "api-gateway",
			},
		},
		{
			name:     "empty tag list",
			input:    []interface{}{},
			expected: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExpandPipelineTags(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ExpandPipelineTags() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestExpandPipelineTags_EdgeCases tests specific edge cases
func TestExpandPipelineTags_EdgeCases(t *testing.T) {
	t.Run("preserves nested Harness expressions", func(t *testing.T) {
		input := []interface{}{"tag:<+<+var1>:<+var2>:<+var3>>"}
		result := ExpandPipelineTags(input)

		expected := "<+<+var1>:<+var2>:<+var3>>"
		if result["tag"] != expected {
			t.Errorf("Failed to preserve nested expression: got %q, want %q", result["tag"], expected)
		}
	})

	t.Run("handles consecutive colons", func(t *testing.T) {
		input := []interface{}{"key:::value"}
		result := ExpandPipelineTags(input)

		if result["key"] != "::value" {
			t.Errorf("Failed to handle consecutive colons: got %q, want %q", result["key"], "::value")
		}
	})
}

// TestExpandPipelineTagsThenFlattenTags verifies round-trip consistency
// This is critical for PIPE-30810 - tags must survive expand/flatten cycle
func TestExpandPipelineTagsThenFlattenTags(t *testing.T) {
	tests := []struct {
		name  string
		input []interface{}
	}{
		{
			name:  "simple tags",
			input: []interface{}{"env:prod", "version:1.0"},
		},
		{
			name:  "tags with colons in values",
			input: []interface{}{"registry:https://example.com:5000", "time:2024:01:15"},
		},
		{
			name:  "Harness expressions - PIPE-30810",
			input: []interface{}{"ImagePush:<+<+pipeline.variables.var>?<+tag>:<+name>>"},
		},
		{
			name:  "mixed tags",
			input: []interface{}{"simple:value", "complex:a:b:c", "standalone"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Expand then flatten - should get back the same tags
			expanded := ExpandPipelineTags(tt.input)
			flattened := FlattenTags(expanded)

			// Convert to maps for comparison (order doesn't matter)
			inputMap := make(map[string]bool)
			for _, v := range tt.input {
				inputMap[v.(string)] = true
			}

			flattenedMap := make(map[string]bool)
			for _, v := range flattened {
				flattenedMap[v] = true
			}

			// Verify same number of tags
			if len(inputMap) != len(flattenedMap) {
				t.Errorf("Round-trip failed: input had %d tags, output has %d", len(inputMap), len(flattenedMap))
				t.Errorf("Input: %v", tt.input)
				t.Errorf("Flattened: %v", flattened)
				return
			}

			// Verify all input tags are in output
			for tag := range inputMap {
				if !flattenedMap[tag] {
					t.Errorf("Round-trip failed: tag %q was lost", tag)
					t.Errorf("Input: %v", tt.input)
					t.Errorf("Flattened: %v", flattened)
				}
			}
		})
	}
}

// TestFlattenTagsWithColons specifically tests FlattenTags with values containing colons
func TestFlattenTagsWithColons(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		validate func([]string) bool
	}{
		{
			name: "tag value with URL and port",
			input: map[string]string{
				"registry": "https://registry.example.com:5000/repo",
			},
			validate: func(result []string) bool {
				return len(result) == 1 && result[0] == "registry:https://registry.example.com:5000/repo"
			},
		},
		{
			name: "tag value with Harness expression containing colons",
			input: map[string]string{
				"ImagePush": "<+<+pipeline.variables.variableA>==\"yes\"?<+pipeline.variables.tag>:<+pipeline.name>>",
			},
			validate: func(result []string) bool {
				return len(result) == 1 && result[0] == "ImagePush:<+<+pipeline.variables.variableA>==\"yes\"?<+pipeline.variables.tag>:<+pipeline.name>>"
			},
		},
		{
			name: "tag value with timestamp",
			input: map[string]string{
				"created": "2024:01:15:10:30:00",
			},
			validate: func(result []string) bool {
				return len(result) == 1 && result[0] == "created:2024:01:15:10:30:00"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FlattenTags(tt.input)
			if !tt.validate(result) {
				t.Errorf("FlattenTags() = %v, validation failed for input %v", result, tt.input)
			}
		})
	}
}
