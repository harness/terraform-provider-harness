package authentication_settings

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testSSOSAMLAccount = "account-123"

// fakeSSOSAMLServer emulates the SAML SSO endpoints: the v2 authentication
// settings read (which carries the SAML entries in ngAuthSettings), the
// multipart metadata upload/update, the delete and the enable/disable login
// endpoint. Like the real backend, a created configuration is persisted with
// authenticationEnabled=false (only the dedicated endpoint flips it), and an
// update rebuilds the configuration from the form fields.
type fakeSSOSAMLServer struct {
	mu sync.Mutex

	configs map[string]*nextgen.SamlSettings
	nextID  int

	// recorded requests
	createForms  []map[string][]string
	createFiles  [][]byte
	updateIDs    []string
	updateForms  []map[string][]string
	updateFiles  [][]byte
	enableReqs   map[string][]bool
	deleteIDs    []string
	getCount     int
	createCalled int

	// failure knobs
	failGet       int
	failGetAfter  int // when >= 0, GET fails with failGetStatus after this many successful GETs
	failGetStatus int
	failCreate    int
	failUpdate    int
	failDelete    int
	failEnable    int
	deleteMessage string
	createNoUUID  bool
	createHidden  bool
	enableResult  *bool
}

type samlFormCapture struct {
	values map[string][]string
	file   []byte
}

func parseSAMLMultipart(r *http.Request) samlFormCapture {
	capture := samlFormCapture{}
	if err := r.ParseMultipartForm(1 << 20); err != nil {
		return capture
	}
	capture.values = r.MultipartForm.Value
	if f, _, err := r.FormFile("file"); err == nil {
		defer f.Close()
		capture.file, _ = io.ReadAll(f)
	}
	return capture
}

func (f *fakeSSOSAMLServer) fail(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(nextgen.Failure{Status: "ERROR", Code: "INVALID_REQUEST", Message: message})
}

func samlFormString(values map[string][]string, key string) string {
	if v := values[key]; len(v) > 0 {
		return v[0]
	}
	return ""
}

func (f *fakeSSOSAMLServer) applyForm(config *nextgen.SamlSettings, values map[string][]string) {
	config.DisplayName = samlFormString(values, "displayName")
	config.GroupMembershipAttr = samlFormString(values, "groupMembershipAttr")
	config.AuthorizationEnabled = samlFormString(values, "authorizationEnabled") == "true"
	config.LogoutUrl = samlFormString(values, "logoutUrl")
	config.EntityIdentifier = samlFormString(values, "entityIdentifier")
	config.SamlProviderType = samlFormString(values, "samlProviderType")
	config.ClientId = samlFormString(values, "clientId")
	config.FriendlySamlName = samlFormString(values, "friendlySamlName")
	config.JitEnabled = samlFormString(values, "jitEnabled") == "true"
	config.JitValidationKey = samlFormString(values, "jitValidationKey")
	config.JitValidationValue = samlFormString(values, "jitValidationValue")
}

func (f *fakeSSOSAMLServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		if f.failGetAfter >= 0 && f.getCount > f.failGetAfter {
			f.fail(w, f.failGetStatus, "fake failure")
			return
		}
		resp := nextgen.RestResponseAuthenticationSettingsV2Response{
			Resource: &nextgen.AuthenticationSettingsV2Response{},
		}
		for _, config := range f.configs {
			entry := *config
			resp.Resource.NgAuthSettings = append(resp.Resource.NgAuthSettings, entry)
		}
		_ = json.NewEncoder(w).Encode(resp)

	case r.Method == http.MethodPost && r.URL.Path == "/ng/api/authentication-settings/saml-metadata":
		capture := parseSAMLMultipart(r)
		f.createCalled++
		f.createForms = append(f.createForms, capture.values)
		f.createFiles = append(f.createFiles, capture.file)
		if f.failCreate != 0 {
			f.fail(w, f.failCreate, "fake failure")
			return
		}
		f.nextID++
		id := "saml-sso-" + strconv.Itoa(f.nextID)
		if !f.createHidden {
			config := &nextgen.SamlSettings{SettingsType: "SAML", Identifier: id}
			f.applyForm(config, capture.values)
			// The backend always persists a new configuration with login disabled.
			config.AuthenticationEnabled = false
			f.configs[id] = config
		}
		uuid := id
		if f.createNoUUID {
			uuid = ""
		}
		_ = json.NewEncoder(w).Encode(nextgen.RestResponseSamlSettingsDto{
			Resource: &nextgen.SamlSettingsDto{Uuid: uuid},
		})

	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/ng/api/authentication-settings/saml-metadata-upload/") && strings.HasSuffix(r.URL.Path, "/authentication"):
		id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/ng/api/authentication-settings/saml-metadata-upload/"), "/authentication")
		enable := r.URL.Query().Get("enable") == "true"
		if f.enableReqs == nil {
			f.enableReqs = map[string][]bool{}
		}
		f.enableReqs[id] = append(f.enableReqs[id], enable)
		if f.failEnable != 0 {
			f.fail(w, f.failEnable, "fake failure")
			return
		}
		if config, ok := f.configs[id]; ok {
			config.AuthenticationEnabled = enable
		}
		result := true
		if f.enableResult != nil {
			result = *f.enableResult
		}
		_ = json.NewEncoder(w).Encode(nextgen.RestResponseBoolean{Resource: result})

	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/ng/api/authentication-settings/saml-metadata-upload/"):
		id := strings.TrimPrefix(r.URL.Path, "/ng/api/authentication-settings/saml-metadata-upload/")
		capture := parseSAMLMultipart(r)
		f.updateIDs = append(f.updateIDs, id)
		f.updateForms = append(f.updateForms, capture.values)
		f.updateFiles = append(f.updateFiles, capture.file)
		if f.failUpdate != 0 {
			f.fail(w, f.failUpdate, "fake failure")
			return
		}
		if config, ok := f.configs[id]; ok {
			f.applyForm(config, capture.values)
		}
		_ = json.NewEncoder(w).Encode(nextgen.RestResponseSsoConfig{})

	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/ng/api/authentication-settings/saml-metadata/"):
		id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/ng/api/authentication-settings/saml-metadata/"), "/delete")
		f.deleteIDs = append(f.deleteIDs, id)
		if f.failDelete != 0 {
			message := f.deleteMessage
			if message == "" {
				message = "fake failure"
			}
			f.fail(w, f.failDelete, message)
			return
		}
		delete(f.configs, id)
		_ = json.NewEncoder(w).Encode(nextgen.RestResponseSsoConfig{})

	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func newFakeSSOSAMLSession(t *testing.T) (*fakeSSOSAMLServer, *internal.Session) {
	t.Helper()
	fake := &fakeSSOSAMLServer{configs: map[string]*nextgen.SamlSettings{}, failGetAfter: -1}
	server := httptest.NewServer(fake)
	t.Cleanup(server.Close)

	cfg := nextgen.NewConfiguration()
	cfg.BasePath = server.URL
	cfg.AccountId = testSSOSAMLAccount
	cfg.ApiKey = "key"
	cfg.HTTPClient.RetryMax = 0
	return fake, &internal.Session{AccountId: testSSOSAMLAccount, PLClient: nextgen.NewAPIClient(cfg)}
}

func boolPtrSAML(b bool) *bool { return &b }

func fullSSOSAMLRaw() map[string]interface{} {
	return map[string]interface{}{
		"display_name":           "Corp SSO",
		"friendly_saml_name":     "corpsaml",
		"metadata_xml":           "<xml>metadata</xml>",
		"entity_identifier":      "https://entity.example.com",
		"logout_url":             "https://logout.example.com",
		"group_membership_attr":  "groups",
		"authentication_enabled": true,
		"saml_provider_type":     "AZURE",
		"client_id":              "client-1",
		"client_secret":          "secret-1",
		"jit_enabled":            true,
		"jit_validation_key":     "key-1",
		"jit_validation_value":   "value-1",
	}
}

func TestUnitSSOSAMLCreateWithLoginEnabled(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, fullSSOSAMLRaw())
	diags := resourceSSOSAMLCreate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.createForms, 1)
	form := fake.createForms[0]
	assert.Equal(t, "Corp SSO", samlFormString(form, "displayName"))
	assert.Equal(t, "groups", samlFormString(form, "groupMembershipAttr"))
	assert.Equal(t, "true", samlFormString(form, "authorizationEnabled"))
	assert.Equal(t, "corpsaml", samlFormString(form, "friendlySamlName"))
	assert.Equal(t, "https://entity.example.com", samlFormString(form, "entityIdentifier"))
	assert.Equal(t, "https://logout.example.com", samlFormString(form, "logoutUrl"))
	assert.Equal(t, "AZURE", samlFormString(form, "samlProviderType"))
	assert.Equal(t, "client-1", samlFormString(form, "clientId"))
	assert.Equal(t, "secret-1", samlFormString(form, "clientSecret"))
	assert.Equal(t, "true", samlFormString(form, "jitEnabled"))
	assert.Equal(t, "key-1", samlFormString(form, "jitValidationKey"))
	assert.Equal(t, "value-1", samlFormString(form, "jitValidationValue"))
	require.Len(t, fake.createFiles, 1)
	assert.Equal(t, "<xml>metadata</xml>", string(fake.createFiles[0]))

	id := d.Id()
	assert.NotEmpty(t, id)
	assert.Equal(t, id, d.Get("identifier"))

	// The configuration was persisted with login disabled, so the resource
	// reconciles authentication_enabled=true through the dedicated endpoint.
	require.Contains(t, fake.enableReqs, id)
	assert.Equal(t, []bool{true}, fake.enableReqs[id])

	// Final state reflects the read-back.
	assert.Equal(t, "Corp SSO", d.Get("display_name"))
	assert.Equal(t, true, d.Get("authentication_enabled"))
	assert.Equal(t, true, d.Get("authorization_enabled"))
	assert.Equal(t, "client-1", d.Get("client_id"))
	// metadata_xml and client_secret are not returned by the API; the
	// configured values stay in state.
	assert.Equal(t, "<xml>metadata</xml>", d.Get("metadata_xml"))
	assert.Equal(t, "secret-1", d.Get("client_secret"))
}

func TestUnitSSOSAMLCreateWithLoginDisabled(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)

	raw := fullSSOSAMLRaw()
	raw["authentication_enabled"] = false
	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, raw)
	diags := resourceSSOSAMLCreate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	// Created disabled and desired disabled: no enable call.
	assert.Empty(t, fake.enableReqs)
	assert.Equal(t, false, d.Get("authentication_enabled"))
}

func TestUnitSSOSAMLCreateRequiresMetadata(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)

	raw := fullSSOSAMLRaw()
	delete(raw, "metadata_xml")
	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, raw)
	diags := resourceSSOSAMLCreate(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "metadata_xml is required")
	assert.Equal(t, 0, fake.createCalled)
}

func TestUnitSSOSAMLCreateApiError(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.failCreate = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, fullSSOSAMLRaw())
	diags := resourceSSOSAMLCreate(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitSSOSAMLCreateNoIdentifierReturned(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.createNoUUID = true

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, fullSSOSAMLRaw())
	diags := resourceSSOSAMLCreate(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "no identifier")
}

func TestUnitSSOSAMLCreateNotFoundAfterCreate(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.createHidden = true

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, fullSSOSAMLRaw())
	diags := resourceSSOSAMLCreate(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "not found after create")
}

func TestUnitSSOSAMLCreateEnableError(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.failEnable = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, fullSSOSAMLRaw())
	diags := resourceSSOSAMLCreate(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitSSOSAMLCreateEnableResultFalse(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.enableResult = boolPtrSAML(false)

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, fullSSOSAMLRaw())
	diags := resourceSSOSAMLCreate(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "authentication enabled")
}

func TestUnitSSOSAMLCreateReadBackError(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	// Fail the read-back right after the create call.
	fake.failGetAfter = 0
	fake.failGetStatus = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, fullSSOSAMLRaw())
	diags := resourceSSOSAMLCreate(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitSSOSAMLUpdateWithMetadata(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.configs["saml-1"] = &nextgen.SamlSettings{
		SettingsType:          "SAML",
		Identifier:            "saml-1",
		DisplayName:           "Old Name",
		AuthenticationEnabled: true,
	}

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, fullSSOSAMLRaw())
	d.SetId("saml-1")
	diags := resourceSSOSAMLUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.updateIDs, 1)
	assert.Equal(t, "saml-1", fake.updateIDs[0])
	form := fake.updateForms[0]
	assert.Equal(t, "Corp SSO", samlFormString(form, "displayName"))
	require.Len(t, fake.updateFiles, 1)
	assert.Equal(t, "<xml>metadata</xml>", string(fake.updateFiles[0]))

	// authentication_enabled changed (state had none): reconciled via endpoint.
	require.Contains(t, fake.enableReqs, "saml-1")
	assert.Equal(t, []bool{true}, fake.enableReqs["saml-1"])

	assert.Equal(t, "Corp SSO", d.Get("display_name"))
	assert.Equal(t, "saml-1", d.Get("identifier"))
}

func TestUnitSSOSAMLUpdateWithoutMetadata(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.configs["saml-1"] = &nextgen.SamlSettings{
		SettingsType:          "SAML",
		Identifier:            "saml-1",
		DisplayName:           "Old Name",
		AuthenticationEnabled: false,
	}

	// No metadata_xml (e.g. after import) and no authentication_enabled change:
	// the update sends form fields only, and the enable endpoint is not called.
	raw := fullSSOSAMLRaw()
	delete(raw, "metadata_xml")
	raw["authentication_enabled"] = false
	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, raw)
	d.SetId("saml-1")
	diags := resourceSSOSAMLUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.updateIDs, 1)
	assert.Empty(t, fake.updateFiles[0])
	assert.Empty(t, fake.enableReqs)
	assert.Equal(t, "Corp SSO", d.Get("display_name"))
}

func TestUnitSSOSAMLUpdateEnableError(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.configs["saml-1"] = &nextgen.SamlSettings{SettingsType: "SAML", Identifier: "saml-1"}
	fake.failEnable = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, fullSSOSAMLRaw())
	d.SetId("saml-1")
	diags := resourceSSOSAMLUpdate(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitSSOSAMLUpdateApiError(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.configs["saml-1"] = &nextgen.SamlSettings{SettingsType: "SAML", Identifier: "saml-1"}
	fake.failUpdate = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, fullSSOSAMLRaw())
	d.SetId("saml-1")
	diags := resourceSSOSAMLUpdate(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitSSOSAMLDelete(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.configs["saml-1"] = &nextgen.SamlSettings{SettingsType: "SAML", Identifier: "saml-1"}

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, map[string]interface{}{})
	d.SetId("saml-1")
	diags := resourceSSOSAMLDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.deleteIDs, 1)
	assert.Equal(t, "saml-1", fake.deleteIDs[0])
	assert.NotContains(t, fake.configs, "saml-1")
	assert.Equal(t, "", d.Id())
}

func TestUnitSSOSAMLDeleteAlreadyRemovedIsIdempotent(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.failDelete = http.StatusBadRequest
	fake.deleteMessage = "No Saml Metadata found for the given samlSSOId."

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, map[string]interface{}{})
	d.SetId("saml-1")
	diags := resourceSSOSAMLDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, "", d.Id())
}

func TestUnitSSOSAMLDeleteApiError(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.failDelete = http.StatusBadRequest
	fake.deleteMessage = "Cannot delete the SAML SSO configuration while user groups are linked."

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, map[string]interface{}{})
	d.SetId("saml-1")
	diags := resourceSSOSAMLDelete(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitSSOSAMLReadPopulatesState(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.configs["saml-1"] = &nextgen.SamlSettings{
		SettingsType:          "SAML",
		Identifier:            "saml-1",
		DisplayName:           "Corp SSO",
		FriendlySamlName:      "corpsaml",
		EntityIdentifier:      "https://entity.example.com",
		Origin:                "https://origin.example.com",
		LogoutUrl:             "https://logout.example.com",
		GroupMembershipAttr:   "groups",
		AuthorizationEnabled:  true,
		AuthenticationEnabled: true,
		SamlProviderType:      "AZURE",
		ClientId:              "client-1",
		JitEnabled:            true,
		JitValidationKey:      "key-1",
		JitValidationValue:    "value-1",
	}

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, map[string]interface{}{
		"metadata_xml":  "<xml>keep-me</xml>",
		"client_secret": "keep-secret",
	})
	d.SetId("saml-1")
	diags := resourceSSOSAMLRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, "saml-1", d.Id())
	assert.Equal(t, "saml-1", d.Get("identifier"))
	assert.Equal(t, "Corp SSO", d.Get("display_name"))
	assert.Equal(t, "corpsaml", d.Get("friendly_saml_name"))
	assert.Equal(t, "https://entity.example.com", d.Get("entity_identifier"))
	assert.Equal(t, "https://origin.example.com", d.Get("origin"))
	assert.Equal(t, "https://logout.example.com", d.Get("logout_url"))
	assert.Equal(t, "groups", d.Get("group_membership_attr"))
	assert.Equal(t, true, d.Get("authorization_enabled"))
	assert.Equal(t, true, d.Get("authentication_enabled"))
	assert.Equal(t, "AZURE", d.Get("saml_provider_type"))
	assert.Equal(t, "client-1", d.Get("client_id"))
	assert.Equal(t, true, d.Get("jit_enabled"))
	assert.Equal(t, "key-1", d.Get("jit_validation_key"))
	assert.Equal(t, "value-1", d.Get("jit_validation_value"))
	// Not returned by the API: configured values are kept.
	assert.Equal(t, "<xml>keep-me</xml>", d.Get("metadata_xml"))
	assert.Equal(t, "keep-secret", d.Get("client_secret"))
	assert.Positive(t, fake.getCount)
}

func TestUnitSSOSAMLReadMissingClearsId(t *testing.T) {
	_, session := newFakeSSOSAMLSession(t)

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, map[string]interface{}{})
	d.SetId("saml-unknown")
	diags := resourceSSOSAMLRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, "", d.Id())
}

func TestUnitSSOSAMLReadApiError(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.failGet = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, map[string]interface{}{})
	d.SetId("saml-1")
	diags := resourceSSOSAMLRead(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitDataSourceSSOSAMLRead(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.configs["saml-1"] = &nextgen.SamlSettings{
		SettingsType:         "SAML",
		Identifier:           "saml-1",
		DisplayName:          "Corp SSO",
		SamlProviderType:     "OKTA",
		AuthorizationEnabled: true,
	}

	d := schema.TestResourceDataRaw(t, DataSourceSSOSAML().Schema, map[string]interface{}{
		"identifier": "saml-1",
	})
	diags := dataSourceSSOSAMLRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	assert.Equal(t, "saml-1", d.Id())
	assert.Equal(t, "Corp SSO", d.Get("display_name"))
	assert.Equal(t, "OKTA", d.Get("saml_provider_type"))
}

func TestUnitDataSourceSSOSAMLNotFound(t *testing.T) {
	_, session := newFakeSSOSAMLSession(t)

	d := schema.TestResourceDataRaw(t, DataSourceSSOSAML().Schema, map[string]interface{}{
		"identifier": "saml-unknown",
	})
	diags := dataSourceSSOSAMLRead(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "no SAML SSO configuration found")
}

func TestUnitDataSourceSSOSAMLReadApiError(t *testing.T) {
	fake, session := newFakeSSOSAMLSession(t)
	fake.failGet = http.StatusBadRequest

	d := schema.TestResourceDataRaw(t, DataSourceSSOSAML().Schema, map[string]interface{}{
		"identifier": "saml-1",
	})
	diags := dataSourceSSOSAMLRead(context.Background(), d, session)
	assert.True(t, diags.HasError())
}

func TestUnitFindSAMLSettings(t *testing.T) {
	assert.Nil(t, findSAMLSettings(nil, "saml-1"))
	assert.Nil(t, findSAMLSettings(&nextgen.AuthenticationSettingsV2Response{}, "saml-1"))
	assert.Nil(t, findSAMLSettings(&nextgen.AuthenticationSettingsV2Response{
		NgAuthSettings: []nextgen.SamlSettings{
			{SettingsType: "SAML", Identifier: "saml-other"},
			{SettingsType: "OAUTH"},
		},
	}, "saml-1"))

	found := findSAMLSettings(&nextgen.AuthenticationSettingsV2Response{
		NgAuthSettings: []nextgen.SamlSettings{
			{SettingsType: "SAML", Identifier: "saml-other"},
			{SettingsType: "SAML", Identifier: "saml-1", DisplayName: "Corp SSO"},
		},
	}, "saml-1")
	require.NotNil(t, found)
	assert.Equal(t, "Corp SSO", found.DisplayName)
}

func TestUnitBuildSAMLSettingsRequest(t *testing.T) {
	t.Run("group membership attr enables authorization", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, map[string]interface{}{
			"display_name":          "Corp SSO",
			"group_membership_attr": "groups",
		})
		req := buildSAMLSettingsRequest(d)
		require.NotNil(t, req.AuthorizationEnabled)
		assert.True(t, *req.AuthorizationEnabled)
		assert.Equal(t, "Corp SSO", req.DisplayName)
		assert.Empty(t, req.ClientSecret)
	})

	t.Run("client id without secret sends the keep mask", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, map[string]interface{}{
			"display_name": "Corp SSO",
			"client_id":    "client-1",
		})
		req := buildSAMLSettingsRequest(d)
		assert.Equal(t, "client-1", req.ClientId)
		assert.Equal(t, samlClientSecretMask, req.ClientSecret)
		require.NotNil(t, req.JitEnabled)
		assert.False(t, *req.JitEnabled)
	})

	t.Run("explicit secret is sent as-is", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceSSOSAML().Schema, map[string]interface{}{
			"display_name":  "Corp SSO",
			"client_id":     "client-1",
			"client_secret": "secret-1",
			"jit_enabled":   true,
		})
		req := buildSAMLSettingsRequest(d)
		assert.Equal(t, "secret-1", req.ClientSecret)
		assert.True(t, *req.JitEnabled)
	})
}
