package helpers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// newCodeTestResourceData returns ResourceData with the given ID set, for exercising
// the Code API error handlers.
func newCodeTestResourceData(t *testing.T, id string) *schema.ResourceData {
	t.Helper()
	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"name": {Type: schema.TypeString, Optional: true},
	}, map[string]interface{}{})
	d.SetId(id)
	return d
}

// TestHandleCodeReadApiError_404ClearsState is the core regression test for CODE-6342:
// a 404 from the Code API on read means the object was deleted out-of-band, so the
// resource must be dropped from state without an error.
func TestHandleCodeReadApiError_404ClearsState(t *testing.T) {
	d := newCodeTestResourceData(t, "my-repo")

	diags := HandleCodeReadApiError(code.GenericSwaggerError{}, d, &http.Response{StatusCode: http.StatusNotFound})

	if diags.HasError() {
		t.Fatalf("expected no error diagnostics on 404 read, got: %v", diags)
	}
	if d.Id() != "" {
		t.Errorf("expected resource ID to be cleared on 404 read, got: %q", d.Id())
	}
}

// TestHandleCodeApiError_404NonReadReturnsError verifies that a 404 outside of a read
// is still surfaced as an error and does not clear state.
func TestHandleCodeApiError_404NonReadReturnsError(t *testing.T) {
	d := newCodeTestResourceData(t, "my-repo")

	diags := HandleCodeApiError(code.GenericSwaggerError{}, d, &http.Response{StatusCode: http.StatusNotFound})

	if !diags.HasError() {
		t.Fatal("expected error diagnostics on non-read 404, got none")
	}
	if d.Id() != "my-repo" {
		t.Errorf("expected resource ID to be preserved, got: %q", d.Id())
	}
}

// TestHandleCodeReadApiError_NonNotFoundKeepsState verifies that auth, validation and
// server errors can never wipe a resource from state.
func TestHandleCodeReadApiError_NonNotFoundKeepsState(t *testing.T) {
	for _, status := range []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusInternalServerError,
	} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			d := newCodeTestResourceData(t, "my-repo")

			diags := HandleCodeReadApiError(code.GenericSwaggerError{}, d, &http.Response{StatusCode: status})

			if !diags.HasError() {
				t.Fatalf("expected error diagnostics on %d read, got none", status)
			}
			if d.Id() != "my-repo" {
				t.Errorf("expected resource ID to be preserved on %d read, got: %q", status, d.Id())
			}
		})
	}
}

// TestHandleCodeReadApiError_NilResponse verifies that an error without an HTTP
// response (e.g. network failure) is returned and does not clear state.
func TestHandleCodeReadApiError_NilResponse(t *testing.T) {
	d := newCodeTestResourceData(t, "my-repo")

	diags := HandleCodeReadApiError(code.GenericSwaggerError{}, d, nil)

	if !diags.HasError() {
		t.Fatal("expected error diagnostics when response is nil, got none")
	}
	if d.Id() != "my-repo" {
		t.Errorf("expected resource ID to be preserved, got: %q", d.Id())
	}
}

// TestHandleCodeDeleteApiError_404IsSuccess verifies that deleting an object that is
// already gone is not an error.
func TestHandleCodeDeleteApiError_404IsSuccess(t *testing.T) {
	d := newCodeTestResourceData(t, "my-repo")

	diags := HandleCodeDeleteApiError(code.GenericSwaggerError{}, d, &http.Response{StatusCode: http.StatusNotFound})

	if diags.HasError() {
		t.Fatalf("expected no error diagnostics on 404 delete, got: %v", diags)
	}
}

// TestHandleCodeDeleteApiError_500ReturnsError verifies that real delete failures are
// still surfaced.
func TestHandleCodeDeleteApiError_500ReturnsError(t *testing.T) {
	d := newCodeTestResourceData(t, "my-repo")

	diags := HandleCodeDeleteApiError(code.GenericSwaggerError{}, d, &http.Response{StatusCode: http.StatusInternalServerError})

	if !diags.HasError() {
		t.Fatal("expected error diagnostics on 500 delete, got none")
	}
}

// TestHandleReadApiError_CodeError404DoesNotClearState documents the original bug:
// the nextgen read handler does not recognize code.GenericSwaggerError, so a 404 is
// returned as a hard error. Code resources must use HandleCodeReadApiError instead.
func TestHandleReadApiError_CodeError404DoesNotClearState(t *testing.T) {
	d := newCodeTestResourceData(t, "my-repo")

	diags := HandleReadApiError(code.GenericSwaggerError{}, d, &http.Response{StatusCode: http.StatusNotFound})

	if !diags.HasError() {
		t.Fatal("expected the nextgen handler to return an error for a Code SDK 404")
	}
	if d.Id() != "my-repo" {
		t.Errorf("expected resource ID to be preserved by the nextgen handler, got: %q", d.Id())
	}
}

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

func TestDecodeAPIErrorBody_NumericAndStringCode(t *testing.T) {
	tests := []struct {
		name        string
		body        string
		wantCode    string
		wantMessage string
		wantErr     bool
	}{
		{
			name:        "numeric code",
			body:        `{"status":"ERROR","code":400,"message":"User email is already invited","correlationId":"cid-1"}`,
			wantCode:    "400",
			wantMessage: "User email is already invited",
		},
		{
			name:        "string code",
			body:        `{"status":"ERROR","code":"INVALID_REQUEST","message":"Role binding is invalid"}`,
			wantCode:    "INVALID_REQUEST",
			wantMessage: "Role binding is invalid",
		},
		{
			name:        "string entity not found",
			body:        `{"code":"ENTITY_NOT_FOUND","message":"User not found"}`,
			wantCode:    "ENTITY_NOT_FOUND",
			wantMessage: "User not found",
		},
		{
			name:        "null code keeps message",
			body:        `{"code":null,"message":"missing code"}`,
			wantCode:    "",
			wantMessage: "missing code",
		},
		{
			name:        "numeric code only",
			body:        `{"code":400}`,
			wantCode:    "400",
			wantMessage: "400",
		},
		{
			name:        "string code only",
			body:        `{"code":"RESOURCE_NOT_FOUND_EXCEPTION"}`,
			wantCode:    "RESOURCE_NOT_FOUND_EXCEPTION",
			wantMessage: "RESOURCE_NOT_FOUND_EXCEPTION",
		},
		{
			name:        "numeric response message code",
			body:        `{"responseMessages":[{"code":400,"level":"ERROR","message":"identifier may not be empty"}]}`,
			wantCode:    "",
			wantMessage: "identifier may not be empty",
		},
		{
			name:        "string response message code",
			body:        `{"responseMessages":[{"code":"INVALID_ARGUMENT","level":"ERROR","message":"identifier may not be empty"}]}`,
			wantCode:    "",
			wantMessage: "identifier may not be empty",
		},
		{
			name:        "field error when message is absent",
			body:        `{"code":"INVALID_REQUEST","errors":[{"fieldId":"email","error":"must not be empty"}]}`,
			wantCode:    "INVALID_REQUEST",
			wantMessage: "email must not be empty",
		},
		{
			name:    "not json",
			body:    `not-json`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := decodeAPIErrorBody([]byte(tt.body))
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected decode error")
				}
				if apiErrorMessage([]byte(tt.body)) != "" {
					t.Fatal("expected no message from an undecodable body")
				}
				return
			}
			if err != nil {
				t.Fatalf("decodeAPIErrorBody() error = %v", err)
			}
			if got := string(parsed.Code); got != tt.wantCode {
				t.Errorf("code = %q, want %q", got, tt.wantCode)
			}
			if got := apiErrorMessage([]byte(tt.body)); got != tt.wantMessage {
				t.Errorf("message = %q, want %q", got, tt.wantMessage)
			}
		})
	}
}

func TestHandleApiError_UserAPINumericAndStringFailureCode(t *testing.T) {
	tests := []struct {
		name         string
		status       int
		body         string
		read         bool
		wantMessage  string
		wantClear    bool
		sdkDecodeErr string
	}{
		{
			name:         "numeric code on create",
			status:       http.StatusBadRequest,
			body:         `{"status":"ERROR","code":400,"message":"User email is already invited","correlationId":"cid-123"}`,
			wantMessage:  "User email is already invited",
			sdkDecodeErr: "Failure.code",
		},
		{
			name:        "string code on create",
			status:      http.StatusBadRequest,
			body:        `{"status":"ERROR","code":"INVALID_REQUEST","message":"Role binding is invalid"}`,
			wantMessage: "Role binding is invalid",
		},
		{
			name:         "numeric code on read",
			status:       http.StatusBadRequest,
			body:         `{"status":"ERROR","code":404,"message":"No user found for email"}`,
			read:         true,
			wantMessage:  "No user found for email",
			sdkDecodeErr: "Failure.code",
		},
		{
			name:      "string entity not found on read clears state",
			status:    http.StatusNotFound,
			body:      `{"code":"ENTITY_NOT_FOUND","message":"User not found"}`,
			read:      true,
			wantClear: true,
		},
		{
			name:        "numeric code on read not-found does not panic",
			status:      http.StatusNotFound,
			body:        `{"code":404,"message":"User does not exist"}`,
			read:        true,
			wantMessage: "User does not exist",
		},
		{
			name:         "numeric model error code",
			status:       http.StatusInternalServerError,
			body:         `{"code":500,"message":"Failed to update user groups"}`,
			wantMessage:  "Failed to update user groups",
			sdkDecodeErr: "ModelError.code",
		},
	}

	ops := []string{"AddUsers", "GetUsers", "AddUserToUserGroups", "RemoveUser"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, op := range ops {
				t.Run(op, func(t *testing.T) {
					srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(tt.status)
						_, _ = w.Write([]byte(tt.body))
					}))
					t.Cleanup(srv.Close)

					err, httpResp := callUserAPI(t, srv.URL, op)
					if err == nil {
						t.Fatal("expected API error")
					}
					swaggerErr, ok := err.(nextgen.GenericSwaggerError)
					if !ok {
						t.Fatalf("expected nextgen.GenericSwaggerError, got %T", err)
					}
					if tt.sdkDecodeErr != "" {
						if swaggerErr.Model() != nil {
							t.Fatalf("SDK decoded a model; numeric code should fail inside harness-go-sdk, model=%#v", swaggerErr.Model())
						}
						if !strings.Contains(swaggerErr.Error(), tt.sdkDecodeErr) {
							t.Fatalf("SDK error = %q, want it to mention %s", swaggerErr.Error(), tt.sdkDecodeErr)
						}
					} else if tt.status == http.StatusBadRequest && swaggerErr.Model() == nil {
						t.Fatal("SDK should decode a string Failure.code")
					}

					parsed, parseErr := ParseErrorBody(swaggerErr)
					if parseErr != nil {
						t.Fatalf("ParseErrorBody() error = %v", parseErr)
					}
					if tt.status == http.StatusBadRequest && tt.name == "string code on create" && parsed.Code != "INVALID_REQUEST" {
						t.Errorf("ParseErrorBody code = %q, want INVALID_REQUEST", parsed.Code)
					}
					if tt.name == "numeric code on create" && parsed.Code != "400" {
						t.Errorf("ParseErrorBody code = %q, want 400", parsed.Code)
					}

					d := newCodeTestResourceData(t, "user-1")
					var diags diag.Diagnostics
					if tt.read {
						diags = HandleReadApiError(err, d, httpResp)
					} else {
						diags = HandleApiError(err, d, httpResp)
					}

					if tt.wantClear {
						if diags.HasError() {
							t.Fatalf("expected state to be cleared without error, got %v", diags)
						}
						if d.Id() != "" {
							t.Fatalf("expected id to be cleared, got %q", d.Id())
						}
						return
					}
					if d.Id() != "user-1" {
						t.Fatalf("expected id to be preserved, got %q", d.Id())
					}
					if !diags.HasError() {
						t.Fatal("expected error diagnostic")
					}
					msg := diags[0].Summary
					if !strings.Contains(msg, tt.wantMessage) {
						t.Errorf("diagnostic = %q, want it to contain %q", msg, tt.wantMessage)
					}
					if strings.Contains(msg, "cannot unmarshal") {
						t.Errorf("diagnostic still hides the API error: %s", msg)
					}
				})
			}
		})
	}
}

func callUserAPI(t *testing.T, baseURL, op string) (error, *http.Response) {
	t.Helper()

	cfg := nextgen.NewConfiguration()
	cfg.BasePath = baseURL
	httpClient := retryablehttp.NewClient()
	httpClient.RetryMax = 0
	httpClient.Logger = nil
	// Return the response to the SDK. The default policy treats 500 as
	// retryable and drops the body, which would hide ModelError decoding.
	httpClient.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		return false, nil
	}
	cfg.HTTPClient = httpClient
	client := nextgen.NewAPIClient(cfg)
	ctx := context.Background()

	switch op {
	case "AddUsers":
		_, resp, err := client.UserApi.AddUsers(ctx, nextgen.AddUsersDto{
			Emails:       []string{"user@example.com"},
			RoleBindings: []nextgen.RoleBinding{},
		}, "account", nil)
		return err, resp
	case "GetUsers":
		_, resp, err := client.UserApi.GetUsers(ctx, "account", &nextgen.UserApiGetUsersOpts{})
		return err, resp
	case "AddUserToUserGroups":
		_, resp, err := client.UserApi.AddUserToUserGroups(ctx, "account", "user-1", &nextgen.UserApiAddUserToUserGroupsOpts{})
		return err, resp
	case "RemoveUser":
		_, resp, err := client.UserApi.RemoveUser(ctx, "user-1", "account", nil)
		return err, resp
	default:
		t.Fatalf("unknown user API operation %s", op)
		return nil, nil
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
