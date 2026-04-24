package helpers

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestIsUndefinedResponseTypeError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "exact match",
			err:      errors.New("undefined response type"),
			expected: true,
		},
		{
			name:     "wrapped message",
			err:      errors.New("error decoding response: undefined response type"),
			expected: true,
		},
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "different error",
			err:      errors.New("connection refused"),
			expected: false,
		},
		{
			name:     "empty error",
			err:      errors.New(""),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isUndefinedResponseTypeError(tt.err)
			if result != tt.expected {
				t.Errorf("isUndefinedResponseTypeError() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHandleUndefinedResponseTypeError_WithHTTPResponse(t *testing.T) {
	httpResp := &http.Response{StatusCode: 502}
	diags := handleUndefinedResponseTypeError(httpResp)

	if len(diags) == 0 {
		t.Fatal("expected diagnostics, got none")
	}

	msg := diags[0].Summary
	if !strings.Contains(msg, "HTTP 502") {
		t.Errorf("expected message to contain 'HTTP 502', got: %s", msg)
	}
	if !strings.Contains(msg, "HARNESS_ENDPOINT") {
		t.Errorf("expected message to mention HARNESS_ENDPOINT, got: %s", msg)
	}
	if !strings.Contains(msg, "HARNESS_ACCOUNT_ID") {
		t.Errorf("expected message to mention HARNESS_ACCOUNT_ID, got: %s", msg)
	}
	if !strings.Contains(msg, "HARNESS_PLATFORM_API_KEY") {
		t.Errorf("expected message to mention HARNESS_PLATFORM_API_KEY, got: %s", msg)
	}
}

func TestHandleUndefinedResponseTypeError_WithoutHTTPResponse(t *testing.T) {
	diags := handleUndefinedResponseTypeError(nil)

	if len(diags) == 0 {
		t.Fatal("expected diagnostics, got none")
	}

	msg := diags[0].Summary
	if strings.Contains(msg, "HTTP") {
		t.Errorf("expected message to NOT contain HTTP status when response is nil, got: %s", msg)
	}
	if !strings.Contains(msg, "HARNESS_ENDPOINT") {
		t.Errorf("expected message to mention HARNESS_ENDPOINT, got: %s", msg)
	}
}
