package setting

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

const (
	testSettingId       = "default_repo_for_git_experience"
	testSettingCategory = "GIT_EXPERIENCE"
	testSettingGroup    = "git_experience"
)

type storedSetting struct {
	value          string
	allowOverrides bool
}

// fakeSettingsServer emulates /ng/api/settings for a single setting. Values are stored per scope;
// a scope without a stored value inherits the default.
type fakeSettingsServer struct {
	mu            sync.Mutex
	values        map[string]storedSetting
	requests      []nextgen.SettingRequestDto
	listedCats    []string
	failUpdateMsg string
}

func scopeKey(r *http.Request) string {
	q := r.URL.Query()
	return q.Get("orgIdentifier") + "/" + q.Get("projectIdentifier")
}

func sourceForScope(r *http.Request) string {
	q := r.URL.Query()
	switch {
	case q.Get("projectIdentifier") != "":
		return settingSourceProject
	case q.Get("orgIdentifier") != "":
		return settingSourceOrg
	default:
		return settingSourceAccount
	}
}

func (f *fakeSettingsServer) settingAt(r *http.Request) nextgen.SettingDto {
	s := nextgen.SettingDto{
		Identifier:      testSettingId,
		Name:            "Default Repository for Git Experience",
		Category:        testSettingCategory,
		GroupIdentifier: testSettingGroup,
		ValueType:       "String",
		AllowOverrides:  true,
		SettingSource:   settingSourceDefault,
	}
	if stored, ok := f.values[scopeKey(r)]; ok {
		s.Value = stored.value
		s.AllowOverrides = stored.allowOverrides
		s.SettingSource = sourceForScope(r)
	}
	return s
}

func (f *fakeSettingsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		category := r.URL.Query().Get("category")
		f.listedCats = append(f.listedCats, category)
		resp := nextgen.ResponseDtoListSettingResponseDto{Status: "SUCCESS"}
		if category == testSettingCategory {
			s := f.settingAt(r)
			resp.Data = []nextgen.SettingResponseDto{{Setting: &s, LastModifiedAt: 1000}}
		}
		_ = json.NewEncoder(w).Encode(resp)
	case http.MethodPut:
		var body []nextgen.SettingRequestDto
		_ = json.NewDecoder(r.Body).Decode(&body)
		f.requests = append(f.requests, body...)
		resp := nextgen.ResponseDtoListSettingUpdateResponseDto{Status: "SUCCESS"}
		for _, req := range body {
			if f.failUpdateMsg != "" {
				resp.Data = append(resp.Data, nextgen.SettingUpdateResponseDto{Identifier: req.Identifier, ErrorMessage: f.failUpdateMsg})
				continue
			}
			if req.UpdateType == settingUpdateTypeRestore {
				delete(f.values, scopeKey(r))
			} else {
				f.values[scopeKey(r)] = storedSetting{value: req.Value, allowOverrides: req.AllowOverrides}
			}
			s := f.settingAt(r)
			resp.Data = append(resp.Data, nextgen.SettingUpdateResponseDto{Identifier: req.Identifier, Setting: &s, LastModifiedAt: 2000, UpdateStatus: true})
		}
		_ = json.NewEncoder(w).Encode(resp)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func newTestSession(t *testing.T) (*fakeSettingsServer, *internal.Session) {
	t.Helper()
	fake := &fakeSettingsServer{values: map[string]storedSetting{}}
	server := httptest.NewServer(fake)
	t.Cleanup(server.Close)

	cfg := nextgen.NewConfiguration()
	cfg.BasePath = server.URL
	cfg.AccountId = "account-123"
	cfg.ApiKey = "key"
	cfg.HTTPClient.RetryMax = 0
	return fake, &internal.Session{AccountId: "account-123", PLClient: nextgen.NewAPIClient(cfg)}
}

func projectSettingData(t *testing.T, value string, allowOverrides bool) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, ResourceSetting().Schema, map[string]interface{}{
		"identifier":      testSettingId,
		"org_id":          "org",
		"project_id":      "proj",
		"value":           value,
		"allow_overrides": allowOverrides,
	})
}

func TestResourceSettingCreateReadDelete(t *testing.T) {
	fake, session := newTestSession(t)
	d := projectSettingData(t, "my-repo", false)

	diags := resourceSettingCreateOrUpdate(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)

	require.Len(t, fake.requests, 1)
	assert.Equal(t, nextgen.SettingRequestDto{Identifier: testSettingId, Value: "my-repo", AllowOverrides: false, UpdateType: "UPDATE"}, fake.requests[0])
	assert.Contains(t, fake.values, "org/proj")
	assert.Equal(t, testSettingId, d.Id())
	assert.Equal(t, testSettingCategory, d.Get("category"))
	assert.Equal(t, testSettingGroup, d.Get("group_identifier"))

	diags = resourceSettingRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, testSettingId, d.Id())
	assert.Equal(t, "my-repo", d.Get("value"))
	assert.Equal(t, false, d.Get("allow_overrides"))
	// Category is known after create, so only that category is listed.
	assert.Equal(t, []string{testSettingCategory}, fake.listedCats)

	diags = resourceSettingDelete(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	require.Len(t, fake.requests, 2)
	assert.Equal(t, "RESTORE", fake.requests[1].UpdateType)
	assert.NotContains(t, fake.values, "org/proj")
}

func TestResourceSettingReadRemovesRestoredSetting(t *testing.T) {
	fake, session := newTestSession(t)
	d := projectSettingData(t, "my-repo", true)
	require.False(t, resourceSettingCreateOrUpdate(context.Background(), d, session).HasError())

	// Someone restores the setting in the UI: it now inherits the default.
	delete(fake.values, "org/proj")

	diags := resourceSettingRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, "", d.Id())
}

func TestResourceSettingReadDetectsValueDrift(t *testing.T) {
	fake, session := newTestSession(t)
	d := projectSettingData(t, "my-repo", true)
	require.False(t, resourceSettingCreateOrUpdate(context.Background(), d, session).HasError())

	fake.values["org/proj"] = storedSetting{value: "changed-in-ui", allowOverrides: true}

	require.False(t, resourceSettingRead(context.Background(), d, session).HasError())
	assert.Equal(t, testSettingId, d.Id())
	assert.Equal(t, "changed-in-ui", d.Get("value"))
}

func TestResourceSettingUpdateFailureIsReported(t *testing.T) {
	fake, session := newTestSession(t)
	fake.failUpdateMsg = "Setting- default_repo_for_git_experience cannot be overridden at the current scope"
	d := projectSettingData(t, "my-repo", true)

	diags := resourceSettingCreateOrUpdate(context.Background(), d, session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "cannot be overridden")
	assert.Equal(t, "", d.Id())
}

func TestResourceSettingImportSearchesCategories(t *testing.T) {
	fake, session := newTestSession(t)
	fake.values["org/"] = storedSetting{value: "org-repo", allowOverrides: true}

	d := schema.TestResourceDataRaw(t, ResourceSetting().Schema, map[string]interface{}{})
	d.SetId("org/" + testSettingId)
	imported, err := ResourceSetting().Importer.State(d, session)
	require.NoError(t, err)
	d = imported[0]

	diags := resourceSettingRead(context.Background(), d, session)
	require.False(t, diags.HasError(), "%v", diags)
	assert.Equal(t, testSettingId, d.Id())
	assert.Equal(t, "org", d.Get("org_id"))
	assert.Equal(t, "org-repo", d.Get("value"))
	assert.Equal(t, testSettingCategory, d.Get("category"))
	assert.Contains(t, fake.listedCats, testSettingCategory)
}

func TestResourceSettingImportUnknownSetting(t *testing.T) {
	_, session := newTestSession(t)

	d := schema.TestResourceDataRaw(t, ResourceSetting().Schema, map[string]interface{}{})
	d.SetId("does_not_exist")
	imported, err := ResourceSetting().Importer.State(d, session)
	require.NoError(t, err)

	diags := resourceSettingRead(context.Background(), imported[0], session)
	require.True(t, diags.HasError())
	assert.Contains(t, diags[0].Summary, "not found")
}
