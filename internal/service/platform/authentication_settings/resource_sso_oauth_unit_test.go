package authentication_settings

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testSSOOAuthAccount = "account-123"

// fakeSSOOAuthServer emulates the endpoints behind the OAuth SSO resource: the
// v2 authentication settings read (which carries the OAUTH entry in
// ngAuthSettings), the OAuth providers update, the auth mechanism switch and
// the OAuth mechanism removal. Writes mutate in-memory state so the read-back
// at the end of create/update observes them, like the real API.
type fakeSSOOAuthServer struct {
	mu sync.Mutex

	// state; hasOAuth reflects whether the account has an OAuth configuration
	mechanism    string
	providers    []string
	filter       string
	hasOAuth     bool
	oauthRemoved bool

	// recorded requests
	updateProviderReqs []nextgen.OAuthSettings
	mechanismReqs      []string
	removeCalled       int
	getCount           int

	// failure knobs: when set, the endpoint responds with this HTTP status
	failGet         int
	failUpdate      int
	failMechanism   int
	removeStatus    int
	removeMessage   string
	updateResult    *bool
	mechanismResult *bool
	removeResult    *bool
}

func (f *fakeSSOOAuthServer) fail(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(nextgen.Failure{Status: "ERROR", Code: "INVALID_REQUEST", Message: message})
}

func (f *fakeSSOOAuthServer) boolResponse(w http.ResponseWriter, override *bool, fallback bool) {
	result := fallback
	if override != nil {
		result = *override
	}
	_ = json.NewEncoder(w).Encode(nextgen.RestResponseBoolean{Resource: result})
}

func (f *fakeSSOOAuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/ng/api/authentication-settings/v2":
		f.getCount++
		if f.failGet != 0 {
			f.fail(w, f.failGet, "fake failure")
			return
		}
		resp := nextgen.RestResponseAuthenticationSettingsV2Response{
			Resource: &nextgen.AuthenticationSettingsV2Response{
				AuthenticationMechanism: f.mechanism,
			},
		}
		if f.hasOAuth {
			resp.Resource.NgAuthSettings = []nextgen.SamlSettings{
				{SettingsType: "SAML", Identifier: "saml-config"},
				{SettingsType: "OAUTH", AllowedProviders: f.providers, Filter: f.filter},
			}
		}
		_ = json.NewEncoder(w).Encode(resp)

	case r.Method == http.MethodPut && r.URL.Path == "/ng/api/authentication-settings/oauth/update-providers":
		var body nextgen.OAuthSettings
		_ = json.NewDecoder(r.Body).Decode(&body)
		f.updateProviderReqs = append(f.updateProviderReqs, body)
		if f.failUpdate != 0 {
			f.fail(w, f.failUpdate, "fake failure")
			return
		}
		f.providers = body.AllowedProviders
		f.filter = body.Filter
		f.hasOAuth = true
		f.boolResponse(w, f.updateResult, true)

	case r.Method == http.MethodPut && r.URL.Path == "/ng/api/authentication-settings/update-auth-mechanism":
		mechanism := r.URL.Query().Get("authenticationMechanism")
		f.mechanismReqs = append(f.mechanismReqs, mechanism)
		if f.failMechanism != 0 {
			f.fail(w, f.failMechanism, "fake failure")
			return
		}
		f.mechanism = mechanism
		f.boolResponse(w, f.mechanismResult, true)

	case r.Method == http.MethodDelete && r.URL.Path == "/ng/api/authentication-settings/oauth/remove-mechanism":
		f.removeCalled++
		if f.removeStatus != 0 {
			message := f.removeMessage
			if message == "" {
				message = "fake failure"
			}
			f.fail(w, f.removeStatus, message)
			return
		}
		f.hasOAuth = false
		f.oauthRemoved = true
		f.boolResponse(w, f.removeResult, true)

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func newFakeSSOOAuthSession(t *testing.T) (*fakeSSOOAuthServer, *internal.Session) {
	t.Helper()
	fake := &fakeSSOOAuthServer{
		mechanism: "USER_PASSWORD",
		hasOAuth:  true,
		providers: []string{"GITHUB", "GOOGLE"},
		filter:    "",
	}
	server := httptest.NewServer(fake)
	t.Cleanup(server.Close)

	cfg := nextgen.NewConfiguration()
	cfg.BasePath = server.URL
	cfg.AccountId = testSSOOAuthAccount
	cfg.ApiKey = "key"
	cfg.HTTPClient.RetryMax = 0
	return fake, &internal.Session{AccountId: testSSOOAuthAccount, PLClient: nextgen.NewAPIClient(cfg)}
}

func boolPtrOAuth(b bool) *bool { return &b }

func TestUnitSSOOAuthReadPopulatesState(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.providers = []string{"AZURE"}
	fake.filter = "corp"

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
	diags := resourceSSOOAuthRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, testSSOOAuthAccount, d.Id())
	assert.Equal(t, []interface{}{"AZURE"}, d.Get("allowed_providers").(*schema.Set).List())
	assert.Equal(t, "corp", d.Get("filter"))
	assert.Positive(t, fake.getCount)
}

func TestUnitSSOOAuthReadRemovedOutOfBand(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.hasOAuth = false

	// Resource read (existing ID): the resource disappears from state.
	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
	d.SetId(testSSOOAuthAccount)
	diags := resourceSSOOAuthRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, "", d.Id())
}

func TestUnitSSOOAuthReadNoConfigDataSource(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.hasOAuth = false

	// Data source read (no prior ID): reports the account with empty providers.
	d := schema.TestResourceDataRaw(t, DataSourceSSOOAuth().Schema, map[string]interface{}{})
	diags := resourceSSOOAuthRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, testSSOOAuthAccount, d.Id())
	assert.Empty(t, d.Get("allowed_providers").(*schema.Set).List())
}

func TestUnitSSOOAuthReadApiError(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.failGet = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
	diags := resourceSSOOAuthRead(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitDataSourceSSOOAuthRead(t *testing.T) {
	_, session := newFakeSSOOAuthSession(t)

	d := schema.TestResourceDataRaw(t, DataSourceSSOOAuth().Schema, map[string]interface{}{})
	diags := resourceSSOOAuthRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, testSSOOAuthAccount, d.Id())
	assert.ElementsMatch(t, []interface{}{"GITHUB", "GOOGLE"}, d.Get("allowed_providers").(*schema.Set).List())
}

func TestUnitSSOOAuthCreate(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{
		"allowed_providers": []interface{}{"GITHUB", "AZURE"},
		"filter":            "corp",
	})
	diags := resourceSSOOAuthCreateOrUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.updateProviderReqs, 1)
	sent := fake.updateProviderReqs[0]
	assert.Equal(t, "OAUTH", sent.SettingsType)
	assert.ElementsMatch(t, []string{"GITHUB", "AZURE"}, sent.AllowedProviders)
	assert.Equal(t, "corp", sent.Filter)

	assert.Equal(t, testSSOOAuthAccount, d.Id())
	assert.ElementsMatch(t, []interface{}{"GITHUB", "AZURE"}, d.Get("allowed_providers").(*schema.Set).List())
}

func TestUnitSSOOAuthCreateUpdateError(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.failUpdate = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{
		"allowed_providers": []interface{}{"GITHUB"},
	})
	diags := resourceSSOOAuthCreateOrUpdate(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitSSOOAuthCreateResultFalse(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.updateResult = boolPtrOAuth(false)

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{
		"allowed_providers": []interface{}{"GITHUB"},
	})
	diags := resourceSSOOAuthCreateOrUpdate(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "failed to update OAuth providers")
}

func TestUnitSSOOAuthDelete(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
	d.SetId(testSSOOAuthAccount)
	diags := resourceSSOOAuthDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	// The mechanism is already USER_PASSWORD, so no switch is needed.
	assert.Empty(t, fake.mechanismReqs)
	assert.Equal(t, 1, fake.removeCalled)
	assert.True(t, fake.oauthRemoved)
	assert.Equal(t, "", d.Id())
}

func TestUnitSSOOAuthDeleteSwitchesMechanismFirst(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.mechanism = "OAUTH"

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
	d.SetId(testSSOOAuthAccount)
	diags := resourceSSOOAuthDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	// The account mechanism is switched to USER_PASSWORD before removal.
	require.Len(t, fake.mechanismReqs, 1)
	assert.Equal(t, "USER_PASSWORD", fake.mechanismReqs[0])
	assert.Equal(t, 1, fake.removeCalled)
	assert.Equal(t, "", d.Id())
}

func TestUnitSSOOAuthDeleteMechanismSwitchError(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.mechanism = "OAUTH"
	fake.failMechanism = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
	d.SetId(testSSOOAuthAccount)
	diags := resourceSSOOAuthDelete(context.Background(), d, session)
	assert.True(t, diags.HasError())
	assert.Equal(t, 0, fake.removeCalled)
}

func TestUnitSSOOAuthDeleteMechanismSwitchResultFalse(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.mechanism = "OAUTH"
	fake.mechanismResult = boolPtrOAuth(false)

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
	d.SetId(testSSOOAuthAccount)
	diags := resourceSSOOAuthDelete(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "USER_PASSWORD")
	assert.Equal(t, 0, fake.removeCalled)
}

func TestUnitSSOOAuthDeleteAlreadyRemovedIsIdempotent(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.removeStatus = http.StatusBadRequest
	fake.removeMessage = "No Oauth settings found for this account."

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
	d.SetId(testSSOOAuthAccount)
	diags := resourceSSOOAuthDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, "", d.Id())
}

func TestUnitSSOOAuthDeleteRemoveError(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.removeStatus = http.StatusInternalServerError

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
	d.SetId(testSSOOAuthAccount)
	diags := resourceSSOOAuthDelete(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitSSOOAuthDeleteRemoveResultFalse(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.removeResult = boolPtrOAuth(false)

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
	d.SetId(testSSOOAuthAccount)
	diags := resourceSSOOAuthDelete(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "failed to remove the OAuth mechanism")
}

func TestUnitSSOOAuthDeleteGetError(t *testing.T) {
	fake, session := newFakeSSOOAuthSession(t)
	fake.failGet = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
	d.SetId(testSSOOAuthAccount)
	diags := resourceSSOOAuthDelete(context.Background(), d, session)
	assert.True(t, diags.HasError())
	assert.Equal(t, 0, fake.removeCalled)
}

func TestUnitSSOOAuthImport(t *testing.T) {
	_, session := newFakeSSOOAuthSession(t)

	t.Run("foreign account id is rejected", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
		d.SetId("some-other-account")
		_, err := resourceSSOOAuthImport(context.Background(), d, session)
		require.Error(t, err)
		assert.Contains(t, err.Error(), testSSOOAuthAccount)
	})

	t.Run("provider account id imports", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceSSOOAuth().Schema, map[string]interface{}{})
		d.SetId(testSSOOAuthAccount)
		imported, err := resourceSSOOAuthImport(context.Background(), d, session)
		require.NoError(t, err)
		require.Len(t, imported, 1)
		assert.Equal(t, testSSOOAuthAccount, imported[0].Id())
	})
}

func TestUnitExpandOAuthProviders(t *testing.T) {
	assert.Nil(t, expandOAuthProviders(nil))
	assert.Nil(t, expandOAuthProviders(schema.NewSet(schema.HashString, nil)))
	assert.ElementsMatch(t, []string{"GITHUB", "GOOGLE"},
		expandOAuthProviders(schema.NewSet(schema.HashString, []interface{}{"GITHUB", "GOOGLE"})))
}

func TestUnitFindOAuthSettings(t *testing.T) {
	assert.Nil(t, findOAuthSettings(nil))
	assert.Nil(t, findOAuthSettings(&nextgen.AuthenticationSettingsV2Response{}))
	assert.Nil(t, findOAuthSettings(&nextgen.AuthenticationSettingsV2Response{
		NgAuthSettings: []nextgen.SamlSettings{{SettingsType: "SAML"}},
	}))

	found := findOAuthSettings(&nextgen.AuthenticationSettingsV2Response{
		NgAuthSettings: []nextgen.SamlSettings{
			{SettingsType: "SAML"},
			{SettingsType: "OAUTH", AllowedProviders: []string{"GITHUB"}},
		},
	})
	require.NotNil(t, found)
	assert.Equal(t, []string{"GITHUB"}, found.AllowedProviders)
}
