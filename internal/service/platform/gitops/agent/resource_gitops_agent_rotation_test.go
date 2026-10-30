package agent

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

type recordedRequest struct {
	Method string
	Path   string
	Query  map[string][]string
	APIKey string
}

type fakeGitops struct {
	mu           sync.Mutex
	requests     []recordedRequest
	regenStatus  int
	regenBody    string
	updateStatus int
	getStatus    int
	getBody      string
	server       *httptest.Server
}

func newFakeGitops(t *testing.T) *fakeGitops {
	t.Helper()
	f := &fakeGitops{regenStatus: http.StatusOK, regenBody: `{"credentials":{"privateKey":"NEWKEY"}}`, updateStatus: http.StatusOK, getStatus: http.StatusOK, getBody: `{"identifier":"a1","credentials":{"privateKey":"GETKEY"}}`}
	f.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.mu.Lock()
		f.requests = append(f.requests, recordedRequest{r.Method, r.URL.Path, r.URL.Query(), r.Header.Get("x-api-key")})
		f.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/gitops/api/v1/agents/a1/credentials":
			w.WriteHeader(f.regenStatus)
			_, _ = w.Write([]byte(f.regenBody))
		case r.Method == http.MethodPut && r.URL.Path == "/gitops/api/v1/agents/a1":
			w.WriteHeader(f.updateStatus)
			_, _ = w.Write([]byte(`{"identifier":"a1","name":"a1","type":"MANAGED_ARGO_PROVIDER","metadata":{"namespace":"argocd"}}`))
		case r.Method == http.MethodGet && r.URL.Path == "/gitops/api/v1/agents/a1":
			w.WriteHeader(f.getStatus)
			_, _ = w.Write([]byte(f.getBody))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(f.server.Close)
	return f
}

func (f *fakeGitops) count(method, path string) int {
	f.mu.Lock()
	defer f.mu.Unlock()
	n := 0
	for _, r := range f.requests {
		if r.Method == method && r.Path == path {
			n++
		}
	}
	return n
}

func (f *fakeGitops) session() *internal.Session {
	hc := retryablehttp.NewClient()
	hc.RetryMax = 0
	return &internal.Session{
		AccountId: "acct",
		Endpoint:  f.server.URL,
		PLClient: nextgen.NewAPIClient(&nextgen.Configuration{
			AccountId:  "acct",
			BasePath:   f.server.URL,
			ApiKey:     "test-api-key",
			HTTPClient: hc,
		}),
	}
}

func baseState(rotation, generation string, skip bool) *terraform.InstanceState {
	skipStr := "false"
	if skip {
		skipStr = "true"
	}
	attrs := map[string]string{
		"id":                        "a1",
		"identifier":                "a1",
		"name":                      "a1",
		"type":                      "MANAGED_ARGO_PROVIDER",
		"org_id":                    "default",
		"project_id":                "proj",
		"skip_agent_token_in_state": skipStr,
		"token_generation":          generation,
		"metadata.#":                "1",
		"metadata.0.namespace":      "argocd",
	}
	if rotation != "" {
		attrs["token_rotation"] = rotation
	}
	return &terraform.InstanceState{ID: "a1", Attributes: attrs}
}

func rotationConfig(rotation int, skip bool) *terraform.ResourceConfig {
	return terraform.NewResourceConfigRaw(map[string]interface{}{
		"identifier":                "a1",
		"name":                      "a1",
		"type":                      "MANAGED_ARGO_PROVIDER",
		"org_id":                    "default",
		"project_id":                "proj",
		"skip_agent_token_in_state": skip,
		"token_rotation":            rotation,
		"metadata": []interface{}{
			map[string]interface{}{"namespace": "argocd"},
		},
	})
}

func TestCustomizeDiff_RotationChangeMarksGenerationUnknown(t *testing.T) {
	diff, err := ResourceGitopsAgent().Diff(context.Background(), baseState("0", "0", true), rotationConfig(1, true), nil)
	require.NoError(t, err)
	require.NotNil(t, diff)
	attr, ok := diff.Attributes["token_generation"]
	require.True(t, ok, "token_generation should be in the diff")
	require.True(t, attr.NewComputed, "token_generation should be unknown after apply")
}

func TestCustomizeDiff_NoRotationChangeKeepsGenerationStable(t *testing.T) {
	diff, err := ResourceGitopsAgent().Diff(context.Background(), baseState("1", "1", true), rotationConfig(1, true), nil)
	require.NoError(t, err)
	if diff != nil {
		_, ok := diff.Attributes["token_generation"]
		require.False(t, ok, "token_generation must not change when token_rotation is unchanged")
	}
}

func TestUpdate_RotationCallsRegenerateAndBumpsGeneration(t *testing.T) {
	f := newFakeGitops(t)
	r := ResourceGitopsAgent()
	cfg := rotationConfig(1, true)
	state := baseState("0", "0", true)
	diff, err := r.Diff(context.Background(), state, cfg, f.session())
	require.NoError(t, err)

	out, diags := r.Apply(context.Background(), state, diff, f.session())
	require.False(t, diags.HasError(), "%v", diags)

	require.Equal(t, 1, f.count(http.MethodPut, "/gitops/api/v1/agents/a1"))
	require.Equal(t, 1, f.count(http.MethodPost, "/gitops/api/v1/agents/a1/credentials"))
	require.Equal(t, "1", out.Attributes["token_generation"])
	require.Equal(t, "1", out.Attributes["token_rotation"])
	require.Empty(t, out.Attributes["agent_token"], "key must not be stored when skip_agent_token_in_state is true")

	for _, req := range f.requests {
		if req.Path == "/gitops/api/v1/agents/a1/credentials" {
			require.Equal(t, []string{"acct"}, req.Query["accountIdentifier"])
			require.Equal(t, []string{"default"}, req.Query["orgIdentifier"])
			require.Equal(t, []string{"proj"}, req.Query["projectIdentifier"])
			require.Equal(t, "test-api-key", req.APIKey)
		}
	}
}

func TestUpdate_RotationStoresKeyWhenSkipIsOff(t *testing.T) {
	f := newFakeGitops(t)
	r := ResourceGitopsAgent()
	state := baseState("0", "0", false)
	diff, err := r.Diff(context.Background(), state, rotationConfig(1, false), f.session())
	require.NoError(t, err)

	out, diags := r.Apply(context.Background(), state, diff, f.session())
	require.False(t, diags.HasError(), "%v", diags)
	require.Equal(t, "NEWKEY", out.Attributes["agent_token"])
	require.Equal(t, "1", out.Attributes["token_generation"])
}

func TestUpdate_NoRotationChangeDoesNotRegenerate(t *testing.T) {
	f := newFakeGitops(t)
	r := ResourceGitopsAgent()
	state := baseState("1", "1", true)
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"identifier":                "a1",
		"name":                      "a1-renamed",
		"type":                      "MANAGED_ARGO_PROVIDER",
		"org_id":                    "default",
		"project_id":                "proj",
		"skip_agent_token_in_state": true,
		"token_rotation":            1,
		"metadata":                  []interface{}{map[string]interface{}{"namespace": "argocd"}},
	})
	diff, err := r.Diff(context.Background(), state, cfg, f.session())
	require.NoError(t, err)
	require.NotNil(t, diff)

	_, diags := r.Apply(context.Background(), state, diff, f.session())
	require.False(t, diags.HasError(), "%v", diags)
	require.Equal(t, 1, f.count(http.MethodPut, "/gitops/api/v1/agents/a1"))
	require.Equal(t, 0, f.count(http.MethodPost, "/gitops/api/v1/agents/a1/credentials"))
}

func TestUpdate_RegenerateFailureSurfacesError(t *testing.T) {
	f := newFakeGitops(t)
	f.regenStatus = http.StatusInternalServerError
	f.regenBody = `{"message":"boom"}`
	r := ResourceGitopsAgent()
	state := baseState("0", "0", true)
	diff, err := r.Diff(context.Background(), state, rotationConfig(1, true), f.session())
	require.NoError(t, err)

	_, diags := r.Apply(context.Background(), state, diff, f.session())
	require.True(t, diags.HasError())
	require.Contains(t, diags[0].Summary, "regenerating agent credentials")
}

func TestRegenerateAgentCredentials(t *testing.T) {
	t.Run("returns the private key and sends scope and api key", func(t *testing.T) {
		f := newFakeGitops(t)
		key, err := regenerateAgentCredentials(context.Background(), f.session(), "a1", "org", "proj")
		require.NoError(t, err)
		require.Equal(t, "NEWKEY", key)
		req := f.requests[0]
		require.Equal(t, http.MethodPost, req.Method)
		require.Equal(t, []string{"org"}, req.Query["orgIdentifier"])
		require.Equal(t, "test-api-key", req.APIKey)
	})

	t.Run("account level agent sends no org or project", func(t *testing.T) {
		f := newFakeGitops(t)
		_, err := regenerateAgentCredentials(context.Background(), f.session(), "a1", "", "")
		require.NoError(t, err)
		require.NotContains(t, f.requests[0].Query, "orgIdentifier")
		require.NotContains(t, f.requests[0].Query, "projectIdentifier")
		require.Equal(t, []string{"acct"}, f.requests[0].Query["accountIdentifier"])
	})

	t.Run("trailing slash on endpoint is tolerated", func(t *testing.T) {
		f := newFakeGitops(t)
		s := f.session()
		s.Endpoint += "/"
		_, err := regenerateAgentCredentials(context.Background(), s, "a1", "", "")
		require.NoError(t, err)
		require.Equal(t, "/gitops/api/v1/agents/a1/credentials", f.requests[0].Path)
	})

	t.Run("non-2xx is an error and includes the status", func(t *testing.T) {
		f := newFakeGitops(t)
		f.regenStatus = http.StatusForbidden
		_, err := regenerateAgentCredentials(context.Background(), f.session(), "a1", "", "")
		require.Error(t, err)
		require.Contains(t, err.Error(), "403")
	})

	t.Run("malformed body is an error", func(t *testing.T) {
		f := newFakeGitops(t)
		f.regenBody = `not json`
		_, err := regenerateAgentCredentials(context.Background(), f.session(), "a1", "", "")
		require.Error(t, err)
	})

	t.Run("empty key is returned as empty, not an error", func(t *testing.T) {
		f := newFakeGitops(t)
		f.regenBody = `{}`
		key, err := regenerateAgentCredentials(context.Background(), f.session(), "a1", "", "")
		require.NoError(t, err)
		require.Equal(t, "", key)
	})
}

func TestNewAttributesAreOptInAndAdditive(t *testing.T) {
	s := ResourceGitopsAgent().Schema
	require.True(t, s["skip_agent_token_in_state"].Optional)
	require.False(t, s["skip_agent_token_in_state"].Required)
	require.True(t, s["token_rotation"].Optional)
	require.True(t, s["token_generation"].Computed)
	require.False(t, s["token_generation"].Optional)
	// None of the new attributes may force replacement of an existing agent.
	for _, k := range []string{"skip_agent_token_in_state", "token_rotation", "token_generation"} {
		require.False(t, s[k].ForceNew, k)
	}
}

// A state written by a provider version that predates the new attributes (no skip_agent_token_in_state,
// token_rotation or token_generation) must plan clean against an unchanged configuration.
func TestDiff_StateFromOlderProviderVersionPlansClean(t *testing.T) {
	state := &terraform.InstanceState{ID: "a1", Attributes: map[string]string{
		"id":                   "a1",
		"identifier":           "a1",
		"name":                 "a1",
		"type":                 "MANAGED_ARGO_PROVIDER",
		"org_id":               "default",
		"project_id":           "proj",
		"agent_token":          "OLD-KEY-FROM-STATE",
		"metadata.#":           "1",
		"metadata.0.namespace": "argocd",
	}}
	cfg := terraform.NewResourceConfigRaw(map[string]interface{}{
		"identifier": "a1",
		"name":       "a1",
		"type":       "MANAGED_ARGO_PROVIDER",
		"org_id":     "default",
		"project_id": "proj",
		"metadata":   []interface{}{map[string]interface{}{"namespace": "argocd"}},
	})
	diff, err := ResourceGitopsAgent().Diff(context.Background(), state, cfg, nil)
	require.NoError(t, err)
	if diff != nil {
		require.True(t, diff.Empty(), "unexpected diff for unchanged config: %v", diff)
	}
}
