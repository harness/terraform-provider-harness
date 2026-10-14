package idp

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func testClient(t *testing.T, h http.HandlerFunc) *integrationsClient {
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	old := retryBackoff
	retryBackoff = time.Millisecond
	t.Cleanup(func() { retryBackoff = old })
	return &integrationsClient{base: srv.URL, accountID: "acc", apiKey: "k", http: srv.Client()}
}

func TestClientRetriesTransientGets(t *testing.T) {
	calls := 0
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"message":"Internal error occurred"}`))
			return
		}
		_, _ = w.Write([]byte(`{"identifier":"a","revision":4}`))
	})
	got, err := c.getIntegration(context.Background(), "o", "", "a")
	if err != nil || got["identifier"] != "a" || calls != 3 {
		t.Fatalf("got %v, %v after %d calls", got, err, calls)
	}
}

func TestClientDoesNotRetryWrites(t *testing.T) {
	calls := 0
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusInternalServerError)
	})
	_, err := c.createIntegration(context.Background(), "", "", map[string]any{"identifier": "a"})
	if !isIntegrationsStatus(err, http.StatusInternalServerError) || calls != 1 {
		t.Fatalf("got %v after %d calls", err, calls)
	}
}

func TestClientGivesUpOnPersistentErrors(t *testing.T) {
	calls := 0
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusServiceUnavailable)
	})
	_, err := c.getIntegration(context.Background(), "", "", "a")
	if !isIntegrationsStatus(err, http.StatusServiceUnavailable) || calls != getRetries+1 {
		t.Fatalf("got %v after %d calls", err, calls)
	}
}

func TestClientRequest(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		want := map[string]string{"x-api-key": "k", "Harness-Account": "acc", "If-Match": `"7"`,
			"Content-Type": "application/json"}
		for k, v := range want {
			if r.Header.Get(k) != v {
				t.Errorf("header %s = %q, want %q", k, r.Header.Get(k), v)
			}
		}
		if r.Method != http.MethodPut || r.URL.Path != "/accounts/acc/orgs/o/projects/p/integrations/a$b" {
			t.Errorf("request %s %s", r.Method, r.URL.EscapedPath())
		}
		w.WriteHeader(http.StatusPreconditionFailed)
		_, _ = w.Write([]byte(`{"code":"REVISION_MISMATCH","message":"stale","details":[{"field":"revision",` +
			`"code":"STALE","message":"is 8"}],"request_id":"r1"}`))
	})
	_, err := c.replaceIntegration(context.Background(), "o", "p", "a$b", 7, map[string]any{})
	if !isIntegrationsStatus(err, http.StatusPreconditionFailed) {
		t.Fatalf("got %v", err)
	}
	if want := "integrations API: 412 REVISION_MISMATCH: stale; revision: is 8 (request r1)"; err.Error() != want {
		t.Errorf("error %q, want %q", err.Error(), want)
	}
}

func TestTags(t *testing.T) {
	m := expandTags([]any{"team:idp", "tier", "url:http://x"})
	if m["team"] != "idp" || m["tier"] != "" || m["url"] != "http://x" {
		t.Fatalf("expand %v", m)
	}
	got := flattenTags(m)
	if len(got) != 3 || got[0] != "team:idp" || got[1] != "tier" || got[2] != "url:http://x" {
		t.Fatalf("flatten %v", got)
	}
}

func TestClientEncodesSpacesInQuery(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "search_term=my%20app" {
			t.Errorf("query %q", r.URL.RawQuery)
		}
		_, _ = w.Write([]byte(`{}`))
	})
	if _, err := c.do(context.Background(), http.MethodGet, "/x", url.Values{"search_term": {"my app"}}, nil, nil,
		nil); err != nil {
		t.Fatal(err)
	}
}

func TestWriteMappingOmitsFieldMappings(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_ = json.NewEncoder(w).Encode(map[string]any{
				"revision": 2, "field_mappings": []any{map[string]any{"source_field": "id", "target_field": "id"}},
			})
			return
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if _, ok := body["field_mappings"]; ok {
			t.Errorf("omitted field_mappings are sent: %v", body["field_mappings"])
		}
		_, _ = w.Write([]byte(`{}`))
	})
	body := map[string]any{"correlation": map[string]any{}}
	if _, err := writeMapping(context.Background(), c, "o", "", "a", "k", body); err != nil {
		t.Fatal(err)
	}
}
