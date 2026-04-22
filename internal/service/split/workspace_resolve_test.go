package split

import (
	"errors"
	"testing"
)

func TestIsWorkspaceListRetriable(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"other", errors.New("not found"), false},
		{"429 list", errors.New(`workspaces list: 429 429 Too Many Requests: {"code":429}`), true},
		{"429 phrase", errors.New("something 429 Too Many Requests"), true},
		{"500 dup", errors.New(`workspaces list: 500 500 Internal Server Error: E11000 duplicate key`), true},
		{"500 E11000", errors.New(`workspaces list: 500 : {"message":"E11000 duplicate key`), true},
		{"500 no dup", errors.New(`workspaces list: 500 500 Internal Server Error: other`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isWorkspaceListRetriable(tt.err); got != tt.want {
				t.Fatalf("isWorkspaceListRetriable(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

func TestIsLargeSegmentGetRetriable(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"other", errors.New("split: 500"), false},
		{"404 status", errors.New("large segment get: 404 404 Not Found: {}"), true},
		{"not found phrase", errors.New("large segment get: 400 not found here"), true},
		{"Not Found caps", errors.New("Not Found"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isLargeSegmentGetRetriable(tt.err); got != tt.want {
				t.Fatalf("isLargeSegmentGetRetriable(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}
