package registry

import "testing"

func TestNormalizeRemoteURLSuffix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "plain", input: "simple", want: "simple"},
		{name: "leading slash", input: "/simple", want: "simple"},
		{name: "trailing slash", input: "simple/", want: "simple"},
		{name: "both slashes", input: "/simple/", want: "simple"},
		{name: "nested path", input: "/root/pypi/+simple/", want: "root/pypi/+simple"},
		{name: "slashes only", input: "///", want: ""},
		{name: "empty", input: "", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := normalizeRemoteURLSuffix(tt.input); got != tt.want {
				t.Fatalf("normalizeRemoteURLSuffix(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
