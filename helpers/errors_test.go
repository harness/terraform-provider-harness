package helpers

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestHandleIacmReadApiError_404ClearsState verifies that a 404 during a read on an
// IaCM endpoint (Ansible inventory/playbook, workspace, etc.) clears the resource from
// state instead of erroring. IaCM endpoints return an IacmError body that has no "code"
// field, so the generic code-based not-found detection never fires for them. A resource
// deleted out-of-band must be detected so Terraform can recreate it.
func TestHandleIacmReadApiError_404ClearsState(t *testing.T) {
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"name": {Type: schema.TypeString, Optional: true},
	}, map[string]interface{}{})
	d.SetId("my-inventory")

	httpResp := &http.Response{StatusCode: 404}
	diags := HandleIacmReadApiError(errors.New("not found"), d, httpResp)

	if diags.HasError() {
		t.Fatalf("expected no error diagnostics on 404 read, got: %v", diags)
	}
	if d.Id() != "" {
		t.Errorf("expected resource ID to be cleared on 404 read, got: %q", d.Id())
	}
}

// TestHandleIacmReadApiError_NonNotFoundReturnsError verifies that non-404 errors still
// surface as diagnostics and do NOT clear state.
func TestHandleIacmReadApiError_NonNotFoundReturnsError(t *testing.T) {
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"name": {Type: schema.TypeString, Optional: true},
	}, map[string]interface{}{})
	d.SetId("my-inventory")

	httpResp := &http.Response{StatusCode: 500}
	diags := HandleIacmReadApiError(errors.New("boom"), d, httpResp)

	if !diags.HasError() {
		t.Fatal("expected error diagnostics on 500 read, got none")
	}
	if d.Id() == "" {
		t.Error("expected resource ID to be preserved on non-404 read")
	}
}

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
