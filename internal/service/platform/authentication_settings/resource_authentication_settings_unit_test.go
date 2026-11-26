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

const testAuthSettingsAccount = "account-123"

// fakeAuthSettingsServer emulates the account-level /ng/api/authentication-settings
// endpoints used by the resource: the v2 read plus the two-factor, whitelisted
// domains, session timeout and public access writes. Writes mutate in-memory state
// so the read-back at the end of create/update observes them, like the real API.
type fakeAuthSettingsServer struct {
	mu sync.Mutex

	// state returned by the v2 read
	twoFactor       bool
	domains         []string
	publicAccess    bool
	sessionTimeout  int32
	absoluteTimeout int32
	mechanism       string

	// recorded request bodies
	twoFactorReqs     []nextgen.TwoFactorAdminOverrideSettings
	domainReqs        [][]string
	sessionReqs       []nextgen.SessionTimeoutSettings
	publicAccessReqs  []bool
	getCount          int
	whitelistedCalled bool

	// failure knobs: when set, the endpoint responds with this HTTP status
	failGet          int
	failTwoFactor    int
	failDomains      int
	failSession      int
	failPublicAccess int

	// result overrides for the RestResponseBoolean.resource field. Two-factor
	// defaults to the post-write state (the backend only returns true when 2FA
	// is enabled); the others default to true.
	twoFactorResult    *bool
	domainsResult      *bool
	sessionResult      *bool
	publicAccessResult *bool
	nilReadResource    bool
}

func (f *fakeAuthSettingsServer) fail(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(nextgen.Failure{Status: "ERROR", Code: "INVALID_REQUEST", Message: "fake failure"})
}

func (f *fakeAuthSettingsServer) boolResponse(w http.ResponseWriter, override *bool, fallback bool) {
	result := fallback
	if override != nil {
		result = *override
	}
	_ = json.NewEncoder(w).Encode(nextgen.RestResponseBoolean{Resource: result})
}

func (f *fakeAuthSettingsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/ng/api/authentication-settings/v2":
		f.getCount++
		if f.failGet != 0 {
			f.fail(w, f.failGet)
			return
		}
		resp := nextgen.RestResponseAuthenticationSettingsV2Response{}
		if !f.nilReadResource {
			resp.Resource = &nextgen.AuthenticationSettingsV2Response{
				TwoFactorEnabled:                f.twoFactor,
				WhitelistedDomains:              f.domains,
				AuthenticationMechanism:         f.mechanism,
				SessionTimeoutInMinutes:         f.sessionTimeout,
				AbsoluteSessionTimeoutInMinutes: f.absoluteTimeout,
				PublicAccessEnabled:             f.publicAccess,
			}
		}
		_ = json.NewEncoder(w).Encode(resp)

	case r.Method == http.MethodPut && r.URL.Path == "/ng/api/authentication-settings/two-factor-admin-override-settings":
		var body nextgen.TwoFactorAdminOverrideSettings
		_ = json.NewDecoder(r.Body).Decode(&body)
		f.twoFactorReqs = append(f.twoFactorReqs, body)
		if f.failTwoFactor != 0 {
			f.fail(w, f.failTwoFactor)
			return
		}
		f.twoFactor = body.AdminOverrideTwoFactorEnabled
		f.boolResponse(w, f.twoFactorResult, f.twoFactor)

	case r.Method == http.MethodPut && r.URL.Path == "/ng/api/authentication-settings/whitelisted-domains":
		var body []string
		_ = json.NewDecoder(r.Body).Decode(&body)
		f.domainReqs = append(f.domainReqs, body)
		f.whitelistedCalled = true
		if f.failDomains != 0 {
			f.fail(w, f.failDomains)
			return
		}
		f.domains = body
		f.boolResponse(w, f.domainsResult, true)

	case r.Method == http.MethodPut && r.URL.Path == "/ng/api/authentication-settings/session-timeout-account-level":
		var body nextgen.SessionTimeoutSettings
		_ = json.NewDecoder(r.Body).Decode(&body)
		f.sessionReqs = append(f.sessionReqs, body)
		if f.failSession != 0 {
			f.fail(w, f.failSession)
			return
		}
		f.sessionTimeout = body.SessionTimeOutInMinutes
		f.absoluteTimeout = body.AbsoluteSessionTimeOutInMinutes
		f.boolResponse(w, f.sessionResult, true)

	case r.Method == http.MethodPut && r.URL.Path == "/ng/api/authentication-settings/public-access":
		var body bool
		_ = json.NewDecoder(r.Body).Decode(&body)
		f.publicAccessReqs = append(f.publicAccessReqs, body)
		if f.failPublicAccess != 0 {
			f.fail(w, f.failPublicAccess)
			return
		}
		f.publicAccess = body
		f.boolResponse(w, f.publicAccessResult, true)

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func newFakeAuthSettingsSession(t *testing.T) (*fakeAuthSettingsServer, *internal.Session) {
	t.Helper()
	fake := &fakeAuthSettingsServer{mechanism: "USER_PASSWORD"}
	server := httptest.NewServer(fake)
	t.Cleanup(server.Close)

	cfg := nextgen.NewConfiguration()
	cfg.BasePath = server.URL
	cfg.AccountId = testAuthSettingsAccount
	cfg.ApiKey = "key"
	cfg.HTTPClient.RetryMax = 0
	return fake, &internal.Session{AccountId: testAuthSettingsAccount, PLClient: nextgen.NewAPIClient(cfg)}
}

func boolPtr(b bool) *bool { return &b }

func TestUnitAuthSettingsReadPopulatesState(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)
	fake.twoFactor = true
	fake.domains = []string{"example.com"}
	fake.publicAccess = true
	fake.sessionTimeout = 60
	fake.absoluteTimeout = 120

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{})
	diags := resourceAuthenticationSettingsRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, testAuthSettingsAccount, d.Id())
	assert.Equal(t, true, d.Get("two_factor_enabled"))
	assert.Equal(t, []interface{}{"example.com"}, d.Get("whitelisted_domains").(*schema.Set).List())
	assert.Equal(t, "USER_PASSWORD", d.Get("authentication_mechanism"))
	assert.Equal(t, 60, d.Get("session_timeout_minutes"))
	assert.Equal(t, 120, d.Get("absolute_session_timeout_minutes"))
	assert.Equal(t, true, d.Get("public_access_enabled"))
}

func TestUnitAuthSettingsReadNilResourceClearsId(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)
	fake.nilReadResource = true

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{})
	d.SetId(testAuthSettingsAccount)
	diags := resourceAuthenticationSettingsRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, "", d.Id())
}

func TestUnitAuthSettingsReadApiError(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)
	fake.failGet = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{})
	diags := resourceAuthenticationSettingsRead(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitDataSourceAuthenticationSettingsRead(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)
	fake.twoFactor = true
	fake.domains = []string{"example.com"}
	fake.sessionTimeout = 90

	d := schema.TestResourceDataRaw(t, DataSourceAuthenticationSettings().Schema, map[string]interface{}{})
	diags := resourceAuthenticationSettingsRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, testAuthSettingsAccount, d.Id())
	assert.Equal(t, true, d.Get("two_factor_enabled"))
	assert.Equal(t, 90, d.Get("session_timeout_minutes"))
	assert.Equal(t, "USER_PASSWORD", d.Get("authentication_mechanism"))
}

func TestUnitAuthSettingsCreateMinimal(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
		"two_factor_enabled":               true,
		"whitelisted_domains":              nil,
		"session_timeout_minutes":          nil,
		"absolute_session_timeout_minutes": nil,
		"public_access_enabled":            nil,
	})
	d.MarkNewResource()
	diags := resourceAuthenticationSettingsCreateOrUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	// Only two-factor is written; undeclared settings are left untouched.
	require.Len(t, fake.twoFactorReqs, 1)
	assert.True(t, fake.twoFactorReqs[0].AdminOverrideTwoFactorEnabled)
	assert.False(t, fake.whitelistedCalled)
	assert.Empty(t, fake.sessionReqs)
	assert.Empty(t, fake.publicAccessReqs)

	assert.Equal(t, testAuthSettingsAccount, d.Id())
	assert.Equal(t, true, d.Get("two_factor_enabled"))
}

func TestUnitAuthSettingsCreateWithDomainsAndPublicAccess(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
		"two_factor_enabled":               true,
		"whitelisted_domains":              []interface{}{"example.com", "www.test.com"},
		"public_access_enabled":            true,
		"session_timeout_minutes":          nil,
		"absolute_session_timeout_minutes": nil,
	})
	d.MarkNewResource()
	diags := resourceAuthenticationSettingsCreateOrUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.domainReqs, 1)
	assert.ElementsMatch(t, []string{"example.com", "www.test.com"}, fake.domainReqs[0])
	require.Len(t, fake.publicAccessReqs, 1)
	assert.True(t, fake.publicAccessReqs[0])
	// Neither timeout field is declared, so no session timeout write happens.
	assert.Empty(t, fake.sessionReqs)
}

func TestUnitAuthSettingsCreateSessionTimeoutBackfillsFromAccount(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)
	fake.sessionTimeout = 45
	fake.absoluteTimeout = 240

	// On create there is no state to carry undeclared timeout fields from, so
	// the live account values are read back and sent to the full-replace API.
	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
		"two_factor_enabled":               true,
		"session_timeout_minutes":          60,
		"absolute_session_timeout_minutes": nil,
		"whitelisted_domains":              nil,
		"public_access_enabled":            nil,
	})
	d.MarkNewResource()
	diags := resourceAuthenticationSettingsCreateOrUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.sessionReqs, 1)
	assert.Equal(t, int32(45), fake.sessionReqs[0].SessionTimeOutInMinutes)
	assert.Equal(t, int32(240), fake.sessionReqs[0].AbsoluteSessionTimeOutInMinutes)
	// One GET for the backfill, one for the final read.
	assert.Equal(t, 2, fake.getCount)
}

func TestUnitAuthSettingsUpdateAllFields(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
		"two_factor_enabled":               true,
		"whitelisted_domains":              []interface{}{"example.com"},
		"session_timeout_minutes":          60,
		"absolute_session_timeout_minutes": 120,
		"public_access_enabled":            true,
	})
	d.SetId(testAuthSettingsAccount)
	diags := resourceAuthenticationSettingsCreateOrUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.twoFactorReqs, 1)
	require.Len(t, fake.domainReqs, 1)
	assert.Equal(t, []string{"example.com"}, fake.domainReqs[0])
	require.Len(t, fake.sessionReqs, 1)
	assert.Equal(t, int32(60), fake.sessionReqs[0].SessionTimeOutInMinutes)
	assert.Equal(t, int32(120), fake.sessionReqs[0].AbsoluteSessionTimeOutInMinutes)
	require.Len(t, fake.publicAccessReqs, 1)
	assert.True(t, fake.publicAccessReqs[0])
	// No backfill read on update: only the final read.
	assert.Equal(t, 1, fake.getCount)

	assert.Equal(t, 60, d.Get("session_timeout_minutes"))
	assert.Equal(t, 120, d.Get("absolute_session_timeout_minutes"))
}

func TestUnitAuthSettingsUpdatePropagatesDomainsError(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)
	fake.failDomains = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
		"two_factor_enabled":  true,
		"whitelisted_domains": []interface{}{"example.com"},
	})
	d.SetId(testAuthSettingsAccount)
	diags := resourceAuthenticationSettingsCreateOrUpdate(context.Background(), d, session)
	assert.True(t, diags.HasError())
	assert.Empty(t, fake.sessionReqs)
	assert.Empty(t, fake.publicAccessReqs)
}

func TestUnitAuthSettingsUpdatePropagatesSessionTimeoutError(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)
	fake.failSession = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
		"two_factor_enabled":      true,
		"session_timeout_minutes": 60,
	})
	d.SetId(testAuthSettingsAccount)
	diags := resourceAuthenticationSettingsCreateOrUpdate(context.Background(), d, session)
	assert.True(t, diags.HasError())
	assert.Empty(t, fake.publicAccessReqs)
}

func TestUnitAuthSettingsUpdatePropagatesPublicAccessError(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)
	fake.failPublicAccess = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
		"two_factor_enabled":    true,
		"public_access_enabled": true,
	})
	d.SetId(testAuthSettingsAccount)
	diags := resourceAuthenticationSettingsCreateOrUpdate(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitAuthSettingsCreatePropagatesBackfillReadError(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
		"two_factor_enabled":      true,
		"session_timeout_minutes": 60,
	})
	d.MarkNewResource()
	// Fail only the backfill read: the first GET errors, later ones succeed.
	fake.failGet = http.StatusBadRequest
	diags := resourceAuthenticationSettingsCreateOrUpdate(context.Background(), d, session)
	assert.True(t, diags.HasError())
	assert.Empty(t, fake.sessionReqs)
}

func TestUnitAuthSettingsCreatePropagatesTwoFactorError(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)
	fake.failTwoFactor = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
		"two_factor_enabled":               true,
		"whitelisted_domains":              nil,
		"session_timeout_minutes":          nil,
		"absolute_session_timeout_minutes": nil,
		"public_access_enabled":            nil,
	})
	diags := resourceAuthenticationSettingsCreateOrUpdate(context.Background(), d, session)
	assert.True(t, diags.HasError())
	assert.False(t, fake.whitelistedCalled)
	assert.Empty(t, fake.sessionReqs)
	assert.Empty(t, fake.publicAccessReqs)
}

func TestUnitAuthSettingsDelete(t *testing.T) {
	fake, session := newFakeAuthSettingsSession(t)
	fake.twoFactor = true
	fake.domains = []string{"example.com"}
	fake.publicAccess = true

	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
		"two_factor_enabled": true,
	})
	d.SetId(testAuthSettingsAccount)
	diags := resourceAuthenticationSettingsDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.twoFactorReqs, 1)
	assert.False(t, fake.twoFactorReqs[0].AdminOverrideTwoFactorEnabled)
	require.Len(t, fake.domainReqs, 1)
	assert.NotNil(t, fake.domainReqs[0])
	assert.Empty(t, fake.domainReqs[0])
	require.Len(t, fake.publicAccessReqs, 1)
	assert.False(t, fake.publicAccessReqs[0])

	assert.False(t, fake.twoFactor)
	assert.Empty(t, fake.domains)
	assert.False(t, fake.publicAccess)
}

func TestUnitAuthSettingsDeletePropagatesErrors(t *testing.T) {
	t.Run("two factor failure stops delete", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.failTwoFactor = http.StatusBadRequest

		d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
			"two_factor_enabled": true,
		})
		d.SetId(testAuthSettingsAccount)
		diags := resourceAuthenticationSettingsDelete(context.Background(), d, session)
		assert.True(t, diags.HasError())
		assert.False(t, fake.whitelistedCalled)
		assert.Empty(t, fake.publicAccessReqs)
	})

	t.Run("whitelisted domains failure stops delete", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.failDomains = http.StatusBadRequest

		d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
			"two_factor_enabled": true,
		})
		d.SetId(testAuthSettingsAccount)
		diags := resourceAuthenticationSettingsDelete(context.Background(), d, session)
		assert.True(t, diags.HasError())
		assert.Empty(t, fake.publicAccessReqs)
	})

	t.Run("public access failure is returned", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.failPublicAccess = http.StatusBadRequest

		d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{
			"two_factor_enabled": true,
		})
		d.SetId(testAuthSettingsAccount)
		diags := resourceAuthenticationSettingsDelete(context.Background(), d, session)
		assert.True(t, diags.HasError())
	})
}

func TestUnitAuthSettingsImporter(t *testing.T) {
	_, session := newFakeAuthSettingsSession(t)
	resource := ResourceAuthenticationSettings()

	t.Run("empty id is rejected", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})
		_, err := resource.Importer.State(d, session)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid identifier")
	})

	t.Run("foreign account id is rejected", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})
		d.SetId("some-other-account")
		_, err := resource.Importer.State(d, session)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "cannot import")
		assert.Contains(t, err.Error(), testAuthSettingsAccount)
	})

	t.Run("provider account id imports", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{})
		d.SetId(testAuthSettingsAccount)
		imported, err := resource.Importer.State(d, session)
		require.NoError(t, err)
		require.Len(t, imported, 1)
		assert.Equal(t, testAuthSettingsAccount, imported[0].Id())
	})
}

func TestUnitExpandStringSet(t *testing.T) {
	assert.Nil(t, expandStringSet(nil))
	assert.Nil(t, expandStringSet(schema.NewSet(hashNormalizedDomain, nil)))
	assert.ElementsMatch(t, []string{"a.com", "b.com"},
		expandStringSet(schema.NewSet(hashNormalizedDomain, []interface{}{"a.com", "b.com"})))
}

func TestUnitReadAuthenticationSettings(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceAuthenticationSettings().Schema, map[string]interface{}{})
	readAuthenticationSettings(d, testAuthSettingsAccount, &nextgen.AuthenticationSettingsV2Response{
		TwoFactorEnabled:                true,
		WhitelistedDomains:              []string{"example.com"},
		AuthenticationMechanism:         "SAML",
		PublicAccessEnabled:             true,
		SessionTimeoutInMinutes:         0,
		AbsoluteSessionTimeoutInMinutes: 0,
	})

	assert.Equal(t, testAuthSettingsAccount, d.Id())
	assert.Equal(t, true, d.Get("two_factor_enabled"))
	assert.Equal(t, "SAML", d.Get("authentication_mechanism"))
	assert.Equal(t, true, d.Get("public_access_enabled"))
	// Unset timeouts come back as 0 and are written through as-is.
	assert.Equal(t, 0, d.Get("session_timeout_minutes"))
	assert.Equal(t, 0, d.Get("absolute_session_timeout_minutes"))
}

func TestUnitSetTwoFactorAuth(t *testing.T) {
	ctx := context.Background()

	t.Run("enable succeeds", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		diags := setTwoFactorAuth(ctx, session.PLClient, true)
		require.False(t, diags.HasError(), "%v", diags)
		require.Len(t, fake.twoFactorReqs, 1)
		assert.True(t, fake.twoFactorReqs[0].AdminOverrideTwoFactorEnabled)
	})

	t.Run("enable without confirmation is an error", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.twoFactorResult = boolPtr(false)
		diags := setTwoFactorAuth(ctx, session.PLClient, true)
		require.True(t, diags.HasError())
		assert.Contains(t, diags[0].Summary, "two-factor")
	})

	t.Run("disable succeeds with resource=false", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.twoFactor = true
		diags := setTwoFactorAuth(ctx, session.PLClient, false)
		require.False(t, diags.HasError(), "%v", diags)
	})

	t.Run("api error is propagated", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.failTwoFactor = http.StatusBadRequest
		diags := setTwoFactorAuth(ctx, session.PLClient, true)
		assert.True(t, diags.HasError())
	})
}

func TestUnitUpdateWhitelistedDomains(t *testing.T) {
	ctx := context.Background()

	t.Run("nil clears with an empty array", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		diags := updateWhitelistedDomains(ctx, session.PLClient, nil)
		require.False(t, diags.HasError(), "%v", diags)
		require.Len(t, fake.domainReqs, 1)
		assert.NotNil(t, fake.domainReqs[0])
		assert.Empty(t, fake.domainReqs[0])
	})

	t.Run("values are sent through", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		diags := updateWhitelistedDomains(ctx, session.PLClient, []string{"example.com"})
		require.False(t, diags.HasError(), "%v", diags)
		require.Len(t, fake.domainReqs, 1)
		assert.Equal(t, []string{"example.com"}, fake.domainReqs[0])
	})

	t.Run("resource=false is an error", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.domainsResult = boolPtr(false)
		diags := updateWhitelistedDomains(ctx, session.PLClient, []string{"example.com"})
		require.True(t, diags.HasError())
		assert.Contains(t, diags[0].Summary, "whitelisted domains")
	})

	t.Run("api error is propagated", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.failDomains = http.StatusBadRequest
		diags := updateWhitelistedDomains(ctx, session.PLClient, []string{"example.com"})
		assert.True(t, diags.HasError())
	})
}

func TestUnitSetSessionTimeout(t *testing.T) {
	ctx := context.Background()
	settings := nextgen.SessionTimeoutSettings{SessionTimeOutInMinutes: 60, AbsoluteSessionTimeOutInMinutes: 120}

	t.Run("success", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		diags := setSessionTimeout(ctx, session.PLClient, settings)
		require.False(t, diags.HasError(), "%v", diags)
		require.Len(t, fake.sessionReqs, 1)
		assert.Equal(t, settings, fake.sessionReqs[0])
	})

	t.Run("resource=false is an error", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.sessionResult = boolPtr(false)
		diags := setSessionTimeout(ctx, session.PLClient, settings)
		require.True(t, diags.HasError())
		assert.Contains(t, diags[0].Summary, "session timeout")
	})

	t.Run("api error is propagated", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.failSession = http.StatusBadRequest
		diags := setSessionTimeout(ctx, session.PLClient, settings)
		assert.True(t, diags.HasError())
	})
}

func TestUnitSetPublicAccess(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		diags := setPublicAccess(ctx, session.PLClient, true)
		require.False(t, diags.HasError(), "%v", diags)
		require.Len(t, fake.publicAccessReqs, 1)
		assert.True(t, fake.publicAccessReqs[0])
	})

	t.Run("resource=false is an error", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.publicAccessResult = boolPtr(false)
		diags := setPublicAccess(ctx, session.PLClient, true)
		require.True(t, diags.HasError())
		assert.Contains(t, diags[0].Summary, "public access")
	})

	t.Run("api error is propagated", func(t *testing.T) {
		fake, session := newFakeAuthSettingsSession(t)
		fake.failPublicAccess = http.StatusBadRequest
		diags := setPublicAccess(ctx, session.PLClient, true)
		assert.True(t, diags.HasError())
	})
}
