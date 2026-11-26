package authentication_settings

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testLoginSettingsAccount = "account-123"

// fakeLoginSettingsServer emulates the endpoints behind GetLoginSettings and
// UpdateLoginSettings: the v2 authentication settings read (which carries the
// login settings in the USER_PASSWORD entry of ngAuthSettings) and the login
// settings update. Updates mutate in-memory state so the read-back at the end
// of create/update observes them, like the real API.
type fakeLoginSettingsServer struct {
	mu sync.Mutex

	settings    nextgen.LoginSettings
	hasSettings bool

	// recorded requests
	updateReqs []nextgen.LoginSettings
	getCount   int

	// failure knobs: when set, the endpoint responds with this HTTP status
	failGet    int
	failUpdate int
}

func (f *fakeLoginSettingsServer) fail(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(nextgen.Failure{Status: "ERROR", Code: "INVALID_REQUEST", Message: "fake failure"})
}

func (f *fakeLoginSettingsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		resp := nextgen.RestResponseAuthenticationSettingsV2Response{
			Resource: &nextgen.AuthenticationSettingsV2Response{},
		}
		if f.hasSettings {
			settings := f.settings
			resp.Resource.NgAuthSettings = []nextgen.SamlSettings{
				{SettingsType: "SAML", Identifier: "saml-config"},
				{SettingsType: "USER_PASSWORD", LoginSettings: &settings},
			}
		}
		_ = json.NewEncoder(w).Encode(resp)

	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/ng/api/authentication-settings/login-settings/"):
		var body nextgen.LoginSettings
		_ = json.NewDecoder(r.Body).Decode(&body)
		f.updateReqs = append(f.updateReqs, body)
		if f.failUpdate != 0 {
			f.fail(w, f.failUpdate)
			return
		}
		f.settings = body
		f.hasSettings = true
		_ = json.NewEncoder(w).Encode(nextgen.RestResponseLoginSettings{Resource: &body})

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func newFakeLoginSettingsSession(t *testing.T) (*fakeLoginSettingsServer, *internal.Session) {
	t.Helper()
	fake := &fakeLoginSettingsServer{
		hasSettings: true,
		settings: nextgen.LoginSettings{
			Uuid:      "ls-uuid-1",
			AccountId: testLoginSettingsAccount,
			PasswordStrengthPolicy: &nextgen.PasswordStrengthPolicy{
				Enabled:               true,
				MinNumberOfCharacters: 10,
			},
			UserLockoutPolicy: &nextgen.UserLockoutPolicy{
				EnableLockoutPolicy:                 true,
				NumberOfFailedAttemptsBeforeLockout: 4,
				LockOutPeriod:                       48,
				NotifyUser:                          false,
				UserGroupsToNotify:                  []map[string]interface{}{{"uuid": "ug-1", "name": "admins"}},
			},
			PasswordExpirationPolicy: &nextgen.PasswordExpirationPolicy{
				Enabled:                   true,
				DaysBeforePasswordExpires: 60,
				DaysBeforeUserNotifiedOfPasswordExpiration: 7,
			},
			LastUpdatedBy: &nextgen.EmbeddedUser{Uuid: "u-1", Name: "Admin", Email: "admin@harness.io"},
			LastUpdatedAt: 1700000000000,
		},
	}
	server := httptest.NewServer(fake)
	t.Cleanup(server.Close)

	cfg := nextgen.NewConfiguration()
	cfg.BasePath = server.URL
	cfg.AccountId = testLoginSettingsAccount
	cfg.ApiKey = "key"
	cfg.HTTPClient.RetryMax = 0
	return fake, &internal.Session{AccountId: testLoginSettingsAccount, PLClient: nextgen.NewAPIClient(cfg)}
}

func fullLoginSettingsRaw() map[string]interface{} {
	return map[string]interface{}{
		"password_strength_policy": []interface{}{map[string]interface{}{
			"enabled":                            true,
			"min_number_of_characters":           12,
			"min_number_of_uppercase_characters": 2,
			"min_number_of_lowercase_characters": 2,
			"min_number_of_special_characters":   1,
			"min_number_of_digits":               1,
		}},
		"user_lockout_policy": []interface{}{map[string]interface{}{
			"enable_lockout_policy":                    true,
			"number_of_failed_attempts_before_lockout": 3,
			"lockout_period":                           12,
			"notify_user":                              false,
		}},
		"password_expiration_policy": []interface{}{map[string]interface{}{
			"enabled":                      true,
			"days_before_password_expires": 30,
			"days_before_user_notified_of_password_expiration": 3,
		}},
	}
}

func TestUnitLoginSettingsReadPopulatesState(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, map[string]interface{}{})
	diags := resourceLoginSettingsRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, testLoginSettingsAccount, d.Id())
	strength := d.Get("password_strength_policy").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, strength["enabled"])
	assert.Equal(t, 10, strength["min_number_of_characters"])
	lockout := d.Get("user_lockout_policy").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, true, lockout["enable_lockout_policy"])
	assert.Equal(t, 4, lockout["number_of_failed_attempts_before_lockout"])
	assert.Equal(t, 48, lockout["lockout_period"])
	assert.Equal(t, false, lockout["notify_user"])
	expiration := d.Get("password_expiration_policy").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, 60, expiration["days_before_password_expires"])
	updatedBy := d.Get("last_updated_by").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "u-1", updatedBy["uuid"])
	assert.Equal(t, "admin@harness.io", updatedBy["email"])
	assert.Equal(t, 1700000000000, d.Get("last_updated_at"))
	// Sanity check the fake served a request.
	assert.Positive(t, fake.getCount)
}

func TestUnitLoginSettingsReadMissingEntryClearsId(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)
	fake.hasSettings = false

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, map[string]interface{}{})
	d.SetId(testLoginSettingsAccount)
	diags := resourceLoginSettingsRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, "", d.Id())
}

func TestUnitLoginSettingsReadApiError(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)
	fake.failGet = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, map[string]interface{}{})
	diags := resourceLoginSettingsRead(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitDataSourceLoginSettingsRead(t *testing.T) {
	_, session := newFakeLoginSettingsSession(t)

	d := schema.TestResourceDataRaw(t, DataSourceLoginSettings().Schema, map[string]interface{}{})
	diags := resourceLoginSettingsRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, testLoginSettingsAccount, d.Id())
	strength := d.Get("password_strength_policy").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, 10, strength["min_number_of_characters"])
}

func TestUnitLoginSettingsCreateMergesAllPolicies(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, fullLoginSettingsRaw())
	diags := resourceLoginSettingsCreateOrUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.updateReqs, 1)
	sent := fake.updateReqs[0]
	assert.Equal(t, "ls-uuid-1", sent.Uuid)

	require.NotNil(t, sent.PasswordStrengthPolicy)
	assert.True(t, sent.PasswordStrengthPolicy.Enabled)
	assert.Equal(t, int32(12), sent.PasswordStrengthPolicy.MinNumberOfCharacters)
	assert.Equal(t, int32(2), sent.PasswordStrengthPolicy.MinNumberOfUppercaseCharacters)
	assert.Equal(t, int32(2), sent.PasswordStrengthPolicy.MinNumberOfLowercaseCharacters)
	assert.Equal(t, int32(1), sent.PasswordStrengthPolicy.MinNumberOfSpecialCharacters)
	assert.Equal(t, int32(1), sent.PasswordStrengthPolicy.MinNumberOfDigits)

	require.NotNil(t, sent.UserLockoutPolicy)
	assert.True(t, sent.UserLockoutPolicy.EnableLockoutPolicy)
	assert.Equal(t, int32(3), sent.UserLockoutPolicy.NumberOfFailedAttemptsBeforeLockout)
	assert.Equal(t, int32(12), sent.UserLockoutPolicy.LockOutPeriod)
	assert.False(t, sent.UserLockoutPolicy.NotifyUser)
	// userGroupsToNotify is carried over from the current settings.
	assert.Equal(t, []map[string]interface{}{{"uuid": "ug-1", "name": "admins"}}, sent.UserLockoutPolicy.UserGroupsToNotify)

	require.NotNil(t, sent.PasswordExpirationPolicy)
	assert.True(t, sent.PasswordExpirationPolicy.Enabled)
	assert.Equal(t, int32(30), sent.PasswordExpirationPolicy.DaysBeforePasswordExpires)
	assert.Equal(t, int32(3), sent.PasswordExpirationPolicy.DaysBeforeUserNotifiedOfPasswordExpiration)

	assert.Equal(t, testLoginSettingsAccount, d.Id())
	strength := d.Get("password_strength_policy").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, 12, strength["min_number_of_characters"])
}

func TestUnitLoginSettingsUpdateKeepsUnmanagedPolicies(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)

	// Only the password strength policy is configured; the other policies must
	// be sent back with their current account values.
	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, map[string]interface{}{
		"password_strength_policy": []interface{}{map[string]interface{}{
			"enabled":                  true,
			"min_number_of_characters": 14,
		}},
	})
	d.SetId(testLoginSettingsAccount)
	diags := resourceLoginSettingsCreateOrUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.updateReqs, 1)
	sent := fake.updateReqs[0]
	require.NotNil(t, sent.PasswordStrengthPolicy)
	assert.Equal(t, int32(14), sent.PasswordStrengthPolicy.MinNumberOfCharacters)

	require.NotNil(t, sent.UserLockoutPolicy)
	assert.Equal(t, int32(4), sent.UserLockoutPolicy.NumberOfFailedAttemptsBeforeLockout)
	assert.Equal(t, int32(48), sent.UserLockoutPolicy.LockOutPeriod)

	require.NotNil(t, sent.PasswordExpirationPolicy)
	assert.Equal(t, int32(60), sent.PasswordExpirationPolicy.DaysBeforePasswordExpires)
}

func TestUnitLoginSettingsCreateGetError(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)
	fake.failGet = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, fullLoginSettingsRaw())
	diags := resourceLoginSettingsCreateOrUpdate(context.Background(), d, session)
	assert.True(t, diags.HasError())
	assert.Empty(t, fake.updateReqs)
}

func TestUnitLoginSettingsCreateMissingSettings(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)
	fake.hasSettings = false

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, fullLoginSettingsRaw())
	diags := resourceLoginSettingsCreateOrUpdate(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "no login settings found")
	assert.Empty(t, fake.updateReqs)
}

func TestUnitLoginSettingsUpdateApiError(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)
	fake.failUpdate = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, fullLoginSettingsRaw())
	diags := resourceLoginSettingsCreateOrUpdate(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitLoginSettingsDeleteResetsToDefaults(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, map[string]interface{}{})
	d.SetId(testLoginSettingsAccount)
	diags := resourceLoginSettingsDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.updateReqs, 1)
	sent := fake.updateReqs[0]
	assert.Equal(t, "ls-uuid-1", sent.Uuid)

	require.NotNil(t, sent.PasswordStrengthPolicy)
	assert.False(t, sent.PasswordStrengthPolicy.Enabled)

	require.NotNil(t, sent.UserLockoutPolicy)
	assert.False(t, sent.UserLockoutPolicy.EnableLockoutPolicy)
	assert.Equal(t, int32(5), sent.UserLockoutPolicy.NumberOfFailedAttemptsBeforeLockout)
	assert.Equal(t, int32(24), sent.UserLockoutPolicy.LockOutPeriod)
	assert.True(t, sent.UserLockoutPolicy.NotifyUser)
	// userGroupsToNotify is preserved across the reset, like the UI does.
	assert.Equal(t, []map[string]interface{}{{"uuid": "ug-1", "name": "admins"}}, sent.UserLockoutPolicy.UserGroupsToNotify)

	require.NotNil(t, sent.PasswordExpirationPolicy)
	assert.False(t, sent.PasswordExpirationPolicy.Enabled)
	assert.Equal(t, int32(90), sent.PasswordExpirationPolicy.DaysBeforePasswordExpires)
	assert.Equal(t, int32(5), sent.PasswordExpirationPolicy.DaysBeforeUserNotifiedOfPasswordExpiration)
}

func TestUnitLoginSettingsDeleteWithoutLockoutPolicy(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)
	fake.settings.UserLockoutPolicy = nil

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, map[string]interface{}{})
	d.SetId(testLoginSettingsAccount)
	diags := resourceLoginSettingsDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.updateReqs, 1)
	require.NotNil(t, fake.updateReqs[0].UserLockoutPolicy)
	assert.Empty(t, fake.updateReqs[0].UserLockoutPolicy.UserGroupsToNotify)
}

func TestUnitLoginSettingsDeleteMissingSettings(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)
	fake.hasSettings = false

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, map[string]interface{}{})
	d.SetId(testLoginSettingsAccount)
	diags := resourceLoginSettingsDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Empty(t, fake.updateReqs)
}

func TestUnitLoginSettingsDeleteGetError(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)
	fake.failGet = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, map[string]interface{}{})
	d.SetId(testLoginSettingsAccount)
	diags := resourceLoginSettingsDelete(context.Background(), d, session)
	assert.True(t, diags.HasError())
	assert.Empty(t, fake.updateReqs)
}

func TestUnitLoginSettingsDeleteUpdateError(t *testing.T) {
	fake, session := newFakeLoginSettingsSession(t)
	fake.failUpdate = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, map[string]interface{}{})
	d.SetId(testLoginSettingsAccount)
	diags := resourceLoginSettingsDelete(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitExpandLoginSettingsPolicies(t *testing.T) {
	t.Run("empty inputs return zero policies", func(t *testing.T) {
		assert.Equal(t, nextgen.PasswordStrengthPolicy{}, expandPasswordStrengthPolicy(nil))
		assert.Equal(t, nextgen.PasswordStrengthPolicy{}, expandPasswordStrengthPolicy([]interface{}{nil}))
		assert.Equal(t, nextgen.UserLockoutPolicy{}, expandUserLockoutPolicy(nil))
		assert.Equal(t, nextgen.UserLockoutPolicy{}, expandUserLockoutPolicy([]interface{}{nil}))
		assert.Equal(t, nextgen.PasswordExpirationPolicy{}, expandPasswordExpirationPolicy(nil))
		assert.Equal(t, nextgen.PasswordExpirationPolicy{}, expandPasswordExpirationPolicy([]interface{}{nil}))
	})

	t.Run("values are expanded", func(t *testing.T) {
		strength := expandPasswordStrengthPolicy([]interface{}{map[string]interface{}{
			"enabled":                            true,
			"min_number_of_characters":           8,
			"min_number_of_uppercase_characters": 1,
			"min_number_of_lowercase_characters": 1,
			"min_number_of_special_characters":   1,
			"min_number_of_digits":               1,
		}})
		assert.Equal(t, int32(8), strength.MinNumberOfCharacters)
		assert.Equal(t, int32(1), strength.MinNumberOfDigits)

		lockout := expandUserLockoutPolicy([]interface{}{map[string]interface{}{
			"enable_lockout_policy":                    true,
			"number_of_failed_attempts_before_lockout": 6,
			"lockout_period":                           2,
			"notify_user":                              true,
		}})
		assert.Equal(t, int32(6), lockout.NumberOfFailedAttemptsBeforeLockout)
		assert.Equal(t, int32(2), lockout.LockOutPeriod)

		expiration := expandPasswordExpirationPolicy([]interface{}{map[string]interface{}{
			"enabled":                      true,
			"days_before_password_expires": 45,
			"days_before_user_notified_of_password_expiration": 4,
		}})
		assert.Equal(t, int32(45), expiration.DaysBeforePasswordExpires)
		assert.Equal(t, int32(4), expiration.DaysBeforeUserNotifiedOfPasswordExpiration)
	})
}

func TestUnitFlattenLoginSettingsPoliciesNil(t *testing.T) {
	assert.Nil(t, flattenPasswordStrengthPolicy(nil))
	assert.Nil(t, flattenUserLockoutPolicy(nil))
	assert.Nil(t, flattenPasswordExpirationPolicy(nil))
}

func TestUnitBuildLoginSettingsWithoutLockoutPolicy(t *testing.T) {
	// When the current settings have no lockout policy, there is nothing to
	// carry userGroupsToNotify over from.
	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, map[string]interface{}{
		"user_lockout_policy": []interface{}{map[string]interface{}{
			"enable_lockout_policy": true,
		}},
	})
	current := &nextgen.LoginSettings{Uuid: "ls-uuid-1"}

	settings := buildLoginSettings(d, current)
	require.NotNil(t, settings.UserLockoutPolicy)
	assert.True(t, settings.UserLockoutPolicy.EnableLockoutPolicy)
	assert.Empty(t, settings.UserLockoutPolicy.UserGroupsToNotify)
}

func TestUnitReadLoginSettingsWithNilPolicies(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceLoginSettings().Schema, map[string]interface{}{})
	readLoginSettings(d, testLoginSettingsAccount, &nextgen.LoginSettings{Uuid: "ls-uuid-1"})

	assert.Equal(t, testLoginSettingsAccount, d.Id())
	assert.Empty(t, d.Get("password_strength_policy").([]interface{}))
	assert.Empty(t, d.Get("user_lockout_policy").([]interface{}))
	assert.Empty(t, d.Get("password_expiration_policy").([]interface{}))
	assert.Empty(t, d.Get("last_updated_by").([]interface{}))
	assert.Equal(t, 0, d.Get("last_updated_at"))
}
