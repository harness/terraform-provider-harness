package authentication_settings

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/harness/harness-go-sdk/harness/ldap"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testSSOLDAPAccount = "account-123"

// fakeSSOLDAPServer emulates the next-generation /v1/ldap-settings account
// singleton API. Writes mutate in-memory state so the read-back at the end of
// create/update observes them, like the real API.
type fakeSSOLDAPServer struct {
	mu sync.Mutex

	settings    *ldap.LdapSettingsDTO
	hasSettings bool

	// recorded requests
	createReqs []ldap.LdapSettingsDTO
	updateIDs  []string
	updateReqs []ldap.LdapSettingsDTO
	deleteIDs  []string
	getCount   int

	// failure knobs: when set, the endpoint responds with this HTTP status and
	// an io.harness.eraro.ResponseMessage body
	failGet    int
	failCreate int
	failUpdate int
	failDelete int
	errMessage string
}

func (f *fakeSSOLDAPServer) fail(w http.ResponseWriter, status int) {
	message := f.errMessage
	if message == "" {
		message = "fake failure"
	}
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ldap.ResponseMessage{Message: message})
}

func (f *fakeSSOLDAPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/v1/ldap-settings":
		var req ldap.LdapSettingsRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		if req.LdapSettings != nil {
			f.createReqs = append(f.createReqs, *req.LdapSettings)
		}
		if f.failCreate != 0 {
			f.fail(w, f.failCreate)
			return
		}
		f.settings = req.LdapSettings
		f.hasSettings = true
		_ = json.NewEncoder(w).Encode(ldap.LdapSettingsResponse{LdapSettings: f.settings})

	case r.Method == http.MethodGet && r.URL.Path == "/v1/ldap-settings":
		f.getCount++
		if f.failGet != 0 {
			f.fail(w, f.failGet)
			return
		}
		if !f.hasSettings {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(ldap.ResponseMessage{Message: "no LDAP settings"})
			return
		}
		_ = json.NewEncoder(w).Encode(ldap.LdapSettingsResponse{LdapSettings: f.settings})

	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/v1/ldap-settings/"):
		id := strings.TrimPrefix(r.URL.Path, "/v1/ldap-settings/")
		var req ldap.LdapSettingsRequest
		_ = json.NewDecoder(r.Body).Decode(&req)
		f.updateIDs = append(f.updateIDs, id)
		if req.LdapSettings != nil {
			f.updateReqs = append(f.updateReqs, *req.LdapSettings)
		}
		if f.failUpdate != 0 {
			f.fail(w, f.failUpdate)
			return
		}
		f.settings = req.LdapSettings
		f.hasSettings = true
		_ = json.NewEncoder(w).Encode(ldap.LdapSettingsResponse{LdapSettings: f.settings})

	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/v1/ldap-settings/"):
		id := strings.TrimPrefix(r.URL.Path, "/v1/ldap-settings/")
		f.deleteIDs = append(f.deleteIDs, id)
		if f.failDelete != 0 {
			f.fail(w, f.failDelete)
			return
		}
		f.settings = nil
		f.hasSettings = false
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func newFakeSSOLDAPSession(t *testing.T) (*fakeSSOLDAPServer, *internal.Session) {
	t.Helper()
	fake := &fakeSSOLDAPServer{hasSettings: true, settings: fullLdapSettingsDTO()}
	server := httptest.NewServer(fake)
	t.Cleanup(server.Close)

	cfg := ldap.NewConfiguration()
	cfg.BasePath = server.URL
	cfg.AccountId = testSSOLDAPAccount
	cfg.ApiKey = "key"
	cfg.HTTPClient.RetryMax = 0
	return fake, &internal.Session{AccountId: testSSOLDAPAccount, LDAPClient: ldap.NewAPIClient(cfg)}
}

func fullLdapSettingsDTO() *ldap.LdapSettingsDTO {
	referrals := true
	maxHops := int32(3)
	recursive := true
	return &ldap.LdapSettingsDTO{
		AccountIdentifier: testSSOLDAPAccount,
		Identifier:        "Corp_LDAP",
		Name:              "Corp LDAP",
		Url:               "ldaps://ldap.example.com:636",
		Disabled:          false,
		CronExpression:    "0 0 */2 ? * * *",
		LdapConnectionSettings: &ldap.LdapConnectionSettingsDTO{
			Host:                              "ldap.example.com",
			Port:                              636,
			SslEnabled:                        true,
			ReferralsEnabled:                  &referrals,
			MaxReferralHops:                   &maxHops,
			ConnectionTimeout:                 3000,
			ResponseTimeout:                   4000,
			UseRecursiveGroupMembershipSearch: &recursive,
			BindDN:                            "cn=admin,dc=example,dc=com",
			SecretRefPath:                     "account.ldap-bind",
			DelegateSelectors:                 []string{"sel-1"},
		},
		LdapUserSettings: []ldap.LdapUserSettingsDTO{
			{
				BaseDN:              "ou=users,dc=example,dc=com",
				SearchFilter:        "(objectClass=person)",
				UidAttr:             "uid",
				SamAccountNameAttr:  "sAMAccountName",
				EmailAttr:           "mail",
				DisplayNameAttr:     "cn",
				GroupMembershipAttr: "memberOf",
			},
		},
		LdapGroupSettings: []ldap.LdapGroupSettingsDTO{
			{
				BaseDN:             "ou=groups,dc=example,dc=com",
				SearchFilter:       "(objectClass=group)",
				NameAttr:           "cn",
				DescriptionAttr:    "description",
				UserMembershipAttr: "member",
				ReferencedUserAttr: "dn",
			},
		},
	}
}

func fullSSOLDAPRaw() map[string]interface{} {
	return map[string]interface{}{
		"name":            "Corp LDAP",
		"disabled":        false,
		"cron_expression": "0 0 */2 ? * * *",
		"url":             "ldaps://ldap.example.com:636",
		"connection_settings": []interface{}{map[string]interface{}{
			"host":                                  "ldap.example.com",
			"port":                                  636,
			"ssl_enabled":                           true,
			"referrals_enabled":                     true,
			"max_referral_hops":                     3,
			"connection_timeout":                    3000,
			"response_timeout":                      4000,
			"use_recursive_group_membership_search": true,
			"bind_dn":                               "cn=admin,dc=example,dc=com",
			"secret_ref_path":                       "account.ldap-bind",
			"delegate_selectors":                    []interface{}{"sel-1"},
		}},
		"user_settings": []interface{}{map[string]interface{}{
			"base_dn":               "ou=users,dc=example,dc=com",
			"search_filter":         "(objectClass=person)",
			"uid_attr":              "uid",
			"sam_account_name_attr": "sAMAccountName",
			"email_attr":            "mail",
			"display_name_attr":     "cn",
			"group_membership_attr": "memberOf",
		}},
		"group_settings": []interface{}{map[string]interface{}{
			"base_dn":              "ou=groups,dc=example,dc=com",
			"search_filter":        "(objectClass=group)",
			"name_attr":            "cn",
			"description_attr":     "description",
			"user_membership_attr": "member",
			"referenced_user_attr": "dn",
		}},
	}
}

func TestUnitSSOLDAPCreateDerivesIdentifierFromName(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)

	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, fullSSOLDAPRaw())
	diags := resourceSSOLDAPCreate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.createReqs, 1)
	sent := fake.createReqs[0]
	assert.Equal(t, testSSOLDAPAccount, sent.AccountIdentifier)
	assert.Equal(t, "Corp_LDAP", sent.Identifier)
	assert.Equal(t, "Corp LDAP", sent.Name)
	assert.Equal(t, "ldaps://ldap.example.com:636", sent.Url)
	assert.Equal(t, "0 0 */2 ? * * *", sent.CronExpression)
	assert.False(t, sent.Disabled)

	require.NotNil(t, sent.LdapConnectionSettings)
	cs := sent.LdapConnectionSettings
	assert.Equal(t, "ldap.example.com", cs.Host)
	assert.Equal(t, int32(636), cs.Port)
	assert.True(t, cs.SslEnabled)
	require.NotNil(t, cs.ReferralsEnabled)
	assert.True(t, *cs.ReferralsEnabled)
	require.NotNil(t, cs.MaxReferralHops)
	assert.Equal(t, int32(3), *cs.MaxReferralHops)
	assert.Equal(t, int32(3000), cs.ConnectionTimeout)
	assert.Equal(t, int32(4000), cs.ResponseTimeout)
	require.NotNil(t, cs.UseRecursiveGroupMembershipSearch)
	assert.True(t, *cs.UseRecursiveGroupMembershipSearch)
	assert.Equal(t, "cn=admin,dc=example,dc=com", cs.BindDN)
	assert.Equal(t, "account.ldap-bind", cs.SecretRefPath)
	assert.Equal(t, []string{"sel-1"}, cs.DelegateSelectors)

	require.Len(t, sent.LdapUserSettings, 1)
	assert.Equal(t, "ou=users,dc=example,dc=com", sent.LdapUserSettings[0].BaseDN)
	assert.Equal(t, "memberOf", sent.LdapUserSettings[0].GroupMembershipAttr)

	require.Len(t, sent.LdapGroupSettings, 1)
	assert.Equal(t, "ou=groups,dc=example,dc=com", sent.LdapGroupSettings[0].BaseDN)
	assert.Equal(t, "dn", sent.LdapGroupSettings[0].ReferencedUserAttr)

	assert.Equal(t, testSSOLDAPAccount, d.Id())
	assert.Equal(t, "Corp_LDAP", d.Get("identifier"))
	assert.Equal(t, "Corp LDAP", d.Get("name"))
}

func TestUnitSSOLDAPCreateWithExplicitIdentifier(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)

	raw := fullSSOLDAPRaw()
	raw["identifier"] = "custom-ldap"
	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, raw)
	diags := resourceSSOLDAPCreate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.createReqs, 1)
	assert.Equal(t, "custom-ldap", fake.createReqs[0].Identifier)
	assert.Equal(t, "custom-ldap", d.Get("identifier"))
}

func TestUnitSSOLDAPCreateApiError(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)
	fake.failCreate = http.StatusBadRequest
	fake.errMessage = "identifier must not be empty"

	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, fullSSOLDAPRaw())
	diags := resourceSSOLDAPCreate(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "Bad Request: identifier must not be empty")
}

func TestUnitSSOLDAPReadPopulatesState(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)

	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, map[string]interface{}{})
	d.SetId(testSSOLDAPAccount)
	diags := resourceSSOLDAPRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, testSSOLDAPAccount, d.Id())
	assert.Equal(t, "Corp_LDAP", d.Get("identifier"))
	assert.Equal(t, "Corp LDAP", d.Get("name"))
	assert.Equal(t, false, d.Get("disabled"))
	assert.Equal(t, "0 0 */2 ? * * *", d.Get("cron_expression"))
	assert.Equal(t, "ldaps://ldap.example.com:636", d.Get("url"))

	cs := d.Get("connection_settings").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "ldap.example.com", cs["host"])
	assert.Equal(t, 636, cs["port"])
	assert.Equal(t, true, cs["ssl_enabled"])
	assert.Equal(t, true, cs["referrals_enabled"])
	assert.Equal(t, 3, cs["max_referral_hops"])
	assert.Equal(t, 3000, cs["connection_timeout"])
	assert.Equal(t, 4000, cs["response_timeout"])
	assert.Equal(t, true, cs["use_recursive_group_membership_search"])
	assert.Equal(t, "cn=admin,dc=example,dc=com", cs["bind_dn"])
	assert.Equal(t, "account.ldap-bind", cs["secret_ref_path"])
	assert.Equal(t, []interface{}{"sel-1"}, cs["delegate_selectors"].(*schema.Set).List())

	us := d.Get("user_settings").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "ou=users,dc=example,dc=com", us["base_dn"])
	assert.Equal(t, "memberOf", us["group_membership_attr"])

	gs := d.Get("group_settings").([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "ou=groups,dc=example,dc=com", gs["base_dn"])
	assert.Equal(t, "member", gs["user_membership_attr"])
	assert.Positive(t, fake.getCount)
}

func TestUnitSSOLDAPReadNotFoundClearsId(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)
	fake.hasSettings = false

	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, map[string]interface{}{})
	d.SetId(testSSOLDAPAccount)
	diags := resourceSSOLDAPRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, "", d.Id())
	assert.True(t, d.IsNewResource())
}

func TestUnitSSOLDAPReadApiError(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)
	fake.failGet = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, map[string]interface{}{})
	d.SetId(testSSOLDAPAccount)
	diags := resourceSSOLDAPRead(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitSSOLDAPUpdate(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)

	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, fullSSOLDAPRaw())
	d.SetId(testSSOLDAPAccount)
	diags := resourceSSOLDAPUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.updateReqs, 1)
	// The state has no identifier (fresh data in the test), so the path falls
	// back to the account ID.
	assert.Equal(t, []string{testSSOLDAPAccount}, fake.updateIDs)
	assert.Equal(t, "Corp_LDAP", fake.updateReqs[0].Identifier)
	assert.Equal(t, "Corp LDAP", d.Get("name"))
}

func TestUnitSSOLDAPUpdateUsesStateIdentifierInPath(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)

	raw := fullSSOLDAPRaw()
	raw["identifier"] = "custom-ldap"
	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, raw)
	d.SetId(testSSOLDAPAccount)
	diags := resourceSSOLDAPUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, []string{"custom-ldap"}, fake.updateIDs)
}

func TestUnitSSOLDAPUpdateApiError(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)
	fake.failUpdate = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, fullSSOLDAPRaw())
	d.SetId(testSSOLDAPAccount)
	diags := resourceSSOLDAPUpdate(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitSSOLDAPDelete(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)

	raw := fullSSOLDAPRaw()
	raw["identifier"] = "custom-ldap"
	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, raw)
	d.SetId(testSSOLDAPAccount)
	diags := resourceSSOLDAPDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, []string{"custom-ldap"}, fake.deleteIDs)
	assert.False(t, fake.hasSettings)
	assert.Equal(t, "", d.Id())
}

func TestUnitSSOLDAPDeleteNotFoundIsIdempotent(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)
	fake.failDelete = http.StatusNotFound

	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, map[string]interface{}{})
	d.SetId(testSSOLDAPAccount)
	diags := resourceSSOLDAPDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, "", d.Id())
}

func TestUnitSSOLDAPDeleteApiError(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)
	fake.failDelete = http.StatusBadRequest
	fake.errMessage = "LDAP settings cannot be deleted while user groups are linked."

	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, map[string]interface{}{})
	d.SetId(testSSOLDAPAccount)
	diags := resourceSSOLDAPDelete(context.Background(), d, session)
	assert.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "user groups are linked")
}

func TestUnitSSOLDAPImport(t *testing.T) {
	_, session := newFakeSSOLDAPSession(t)

	t.Run("foreign account id is rejected", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, map[string]interface{}{})
		d.SetId("some-other-account")
		_, err := resourceSSOLDAPImport(context.Background(), d, session)
		require.Error(t, err)
		assert.Contains(t, err.Error(), testSSOLDAPAccount)
	})

	t.Run("provider account id imports", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, map[string]interface{}{})
		d.SetId(testSSOLDAPAccount)
		imported, err := resourceSSOLDAPImport(context.Background(), d, session)
		require.NoError(t, err)
		require.Len(t, imported, 1)
		assert.Equal(t, testSSOLDAPAccount, imported[0].Id())
	})
}

func TestUnitDataSourceSSOLDAPRead(t *testing.T) {
	_, session := newFakeSSOLDAPSession(t)

	d := schema.TestResourceDataRaw(t, DataSourceSSOLDAP().Schema, map[string]interface{}{})
	diags := dataSourceSSOLDAPRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, testSSOLDAPAccount, d.Id())
	assert.Equal(t, "Corp_LDAP", d.Get("identifier"))
	assert.Equal(t, "Corp LDAP", d.Get("name"))
}

func TestUnitDataSourceSSOLDAPNotFound(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)
	fake.hasSettings = false

	d := schema.TestResourceDataRaw(t, DataSourceSSOLDAP().Schema, map[string]interface{}{})
	diags := dataSourceSSOLDAPRead(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "no LDAP SSO configuration found")
}

func TestUnitDataSourceSSOLDAPReadApiError(t *testing.T) {
	fake, session := newFakeSSOLDAPSession(t)
	fake.failGet = http.StatusInternalServerError

	d := schema.TestResourceDataRaw(t, DataSourceSSOLDAP().Schema, map[string]interface{}{})
	diags := dataSourceSSOLDAPRead(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitLdapIdentifierFromName(t *testing.T) {
	assert.Equal(t, "Corp_LDAP", ldapIdentifierFromName("Corp LDAP"))
	assert.Equal(t, "my_ldap_config", ldapIdentifierFromName("  my ldap config  "))
	assert.Equal(t, "ldap", ldapIdentifierFromName("ldap"))
}

func TestUnitLdapSettingsIdentifierFallback(t *testing.T) {
	_, session := newFakeSSOLDAPSession(t)

	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, map[string]interface{}{})
	assert.Equal(t, testSSOLDAPAccount, ldapSettingsIdentifier(d, session))

	d.Set("identifier", "custom-ldap")
	assert.Equal(t, "custom-ldap", ldapSettingsIdentifier(d, session))
}

func TestUnitExpandLdapSettingsNil(t *testing.T) {
	assert.Nil(t, expandLdapConnectionSettings(nil))
	assert.Nil(t, expandLdapConnectionSettings([]interface{}{nil}))
	assert.Nil(t, expandLdapGroupSettings(nil))
	assert.Nil(t, expandLdapGroupSettings([]interface{}{nil}))
	assert.Empty(t, expandLdapUserSettings(nil))
	assert.Empty(t, expandLdapUserSettings([]interface{}{nil}))
}

func TestUnitFlattenLdapSettingsNil(t *testing.T) {
	assert.Nil(t, flattenLdapConnectionSettings(nil))
	assert.Empty(t, flattenLdapUserSettings(nil))
	assert.Empty(t, flattenLdapGroupSettings(nil))

	// Nil optional pointers flatten to zero values.
	out := flattenLdapConnectionSettings(&ldap.LdapConnectionSettingsDTO{Host: "h"})
	m := out[0].(map[string]interface{})
	assert.Equal(t, false, m["referrals_enabled"])
	assert.Equal(t, int32(0), m["max_referral_hops"])
	assert.Equal(t, false, m["use_recursive_group_membership_search"])
}

func TestUnitLdapBoolIntHelpers(t *testing.T) {
	assert.False(t, boolValue(nil))
	b := true
	assert.True(t, boolValue(&b))
	assert.Equal(t, int32(0), intValue(nil))
	i := int32(7)
	assert.Equal(t, int32(7), intValue(&i))

	assert.Nil(t, expandLdapStringSet(nil))
	assert.Empty(t, expandLdapStringSet(schema.NewSet(schema.HashString, nil)))
	assert.ElementsMatch(t, []string{"a", "b"},
		expandLdapStringSet(schema.NewSet(schema.HashString, []interface{}{"a", "b"})))

	assert.Empty(t, flattenStringSlice(nil))
	assert.Equal(t, []interface{}{"a"}, flattenStringSlice([]string{"a"}))
}

func TestUnitReadLdapSettingsNilResource(t *testing.T) {
	d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, map[string]interface{}{})
	d.SetId(testSSOLDAPAccount)
	// A 200 response without ldap_settings leaves the state untouched.
	readLdapSettings(d, ldap.LdapSettingsResponse{})
	assert.Equal(t, testSSOLDAPAccount, d.Id())
}

func TestUnitHandleLdapApiError(t *testing.T) {
	ctx := context.Background()

	cases := []struct {
		name     string
		status   int
		contains string
	}{
		{"bad request", http.StatusBadRequest, "Bad Request: fake failure"},
		{"unauthorized", http.StatusUnauthorized, "401"},
		{"forbidden", http.StatusForbidden, "403"},
		{"not found", http.StatusNotFound, "not found"},
		{"server error falls through", http.StatusInternalServerError, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fake, session := newFakeSSOLDAPSession(t)
			fake.failCreate = tc.status
			d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, fullSSOLDAPRaw())
			diags := resourceSSOLDAPCreate(ctx, d, session)
			require.True(t, diags.HasError())
			if tc.contains != "" {
				assert.Contains(t, diags[0].Summary, tc.contains)
			}
		})
	}

	t.Run("non-ldap error falls back to the generic handler", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceSSOLDAP().Schema, map[string]interface{}{})
		diags := handleLdapApiError(errors.New("connection refused"), d, nil)
		require.True(t, diags.HasError())
		assert.Contains(t, diags[0].Summary, "connection refused")
	})
}
