package token

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestResourceTokenForceNewFields pins the set of fields that must replace the
// token instead of updating it in place. A token's identity is the composite of
// account, scope, apiKeyType, parentIdentifier, apiKeyIdentifier and identifier;
// the update API cannot move a token across any of these, and the create-only
// metadata fields are silently ignored by the server on update.
func TestResourceTokenForceNewFields(t *testing.T) {
	forceNewFields := []string{
		"identifier",
		"account_id",
		"org_id",
		"project_id",
		"apikey_type",
		"parent_id",
		"apikey_id",
		"valid",
		"scheduled_expire_time",
	}
	for _, field := range forceNewFields {
		s, ok := ResourceToken().Schema[field]
		require.True(t, ok, "schema field %q missing", field)
		assert.True(t, s.ForceNew, "field %q must be ForceNew", field)
	}

	// email/username/encoded_password are create-only and never returned by the
	// server on read; ForceNew on them would cause a perpetual replacement loop.
	inPlaceFields := []string{"name", "description", "tags", "valid_from", "valid_to", "email", "username", "encoded_password"}
	for _, field := range inPlaceFields {
		s, ok := ResourceToken().Schema[field]
		require.True(t, ok, "schema field %q missing", field)
		assert.False(t, s.ForceNew, "field %q must stay updatable in place", field)
	}
}

// fakeTokenServer emulates DELETE /ng/api/token/{id} with a configurable status.
type fakeTokenServer struct {
	deleteStatus int
	deleteCalls  int
}

func (f *fakeTokenServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		f.deleteCalls++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(f.deleteStatus)
		if f.deleteStatus >= 400 {
			_, _ = w.Write([]byte(`{"status":"ERROR","code":"ENTITY_NOT_FOUND","message":"token does not exist"}`))
			return
		}
		_, _ = w.Write([]byte(`{"status":"SUCCESS","data":true}`))
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func newTokenTestSession(t *testing.T, fake *fakeTokenServer) *internal.Session {
	t.Helper()
	server := httptest.NewServer(fake)
	t.Cleanup(server.Close)

	cfg := nextgen.NewConfiguration()
	cfg.BasePath = server.URL
	cfg.AccountId = "account-123"
	cfg.ApiKey = "key"
	cfg.HTTPClient.RetryMax = 0
	return &internal.Session{AccountId: "account-123", PLClient: nextgen.NewAPIClient(cfg)}
}

func tokenResourceData(t *testing.T) *schema.ResourceData {
	d := schema.TestResourceDataRaw(t, ResourceToken().Schema, map[string]interface{}{
		"identifier":  "token1",
		"account_id":  "account-123",
		"apikey_type": "USER",
		"parent_id":   "user-1",
		"apikey_id":   "apikey1",
	})
	d.SetId("token1")
	return d
}

// TestResourceTokenDeleteNotFoundIsSuccess covers PL-75965: deleting a token that
// is already gone (404) must succeed, since the desired end state is reached.
func TestResourceTokenDeleteNotFoundIsSuccess(t *testing.T) {
	fake := &fakeTokenServer{deleteStatus: http.StatusNotFound}
	session := newTokenTestSession(t, fake)

	diags := resourceTokenDelete(context.Background(), tokenResourceData(t), session)

	assert.False(t, diags.HasError(), "expected no error on 404 delete, got: %v", diags)
	assert.Equal(t, 1, fake.deleteCalls)
}

// TestResourceTokenDeleteServerErrorIsReported ensures real delete failures still
// surface as errors.
func TestResourceTokenDeleteServerErrorIsReported(t *testing.T) {
	fake := &fakeTokenServer{deleteStatus: http.StatusInternalServerError}
	session := newTokenTestSession(t, fake)

	diags := resourceTokenDelete(context.Background(), tokenResourceData(t), session)

	assert.True(t, diags.HasError(), "expected error diagnostics on 500 delete")
	assert.Equal(t, 1, fake.deleteCalls)
}
