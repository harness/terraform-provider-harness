package helpers

import "testing"

func TestTruncateString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxLen   int
		expected string
	}{
		{
			name:     "shorter than max returned as-is",
			input:    "hello",
			maxLen:   10,
			expected: "hello",
		},
		{
			name:     "equal to max returned as-is",
			input:    "hello",
			maxLen:   5,
			expected: "hello",
		},
		{
			name:     "longer than max is truncated with ellipsis",
			input:    "hello world",
			maxLen:   8,
			expected: "hello...",
		},
		{
			name:     "max not larger than ellipsis is hard cut",
			input:    "hello",
			maxLen:   2,
			expected: "he",
		},
		{
			name:     "non-positive max returns empty",
			input:    "hello",
			maxLen:   0,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TruncateString(tt.input, tt.maxLen); got != tt.expected {
				t.Errorf("TruncateString(%q, %d) = %q, want %q", tt.input, tt.maxLen, got, tt.expected)
			}
		})
	}
}

func TestNormalizeIdentifier(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "trims and lowercases",
			input:    "  MyConnector  ",
			expected: "myconnector",
		},
		{
			name:     "already normalized",
			input:    "abc",
			expected: "abc",
		},
		{
			name:     "whitespace only returns empty",
			input:    "   ",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeIdentifier(tt.input); got != tt.expected {
				t.Errorf("NormalizeIdentifier(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
