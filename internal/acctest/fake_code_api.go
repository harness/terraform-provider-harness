package acctest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

// FakeCodeAPI is a minimal in-memory stand-in for the Harness Code REST API
// (/code/api/v1) covering repos, repo webhooks and repo branch rules.
//
// It lets tests exercise the real provider code and a real Terraform CLI without a
// Harness account, including "deleted out-of-band" scenarios. Like the real service,
// it answers 404 for objects that do not exist, and for any webhook or rule whose
// parent repo does not exist.
//
// It stores whatever the provider sends and echoes it back, so it is intentionally
// not a faithful model of server-side validation. All objects are account-level
// (org/project query parameters are ignored).
type FakeCodeAPI struct {
	Server *httptest.Server

	mu       sync.Mutex
	objects  map[string]map[string]interface{}
	requests []string
}

const fakeCodeAPIBasePath = "/code/api/v1/"

// NewFakeCodeAPI starts a fake Code API that is shut down when the test ends.
func NewFakeCodeAPI(t testing.TB) *FakeCodeAPI {
	t.Helper()
	f := &FakeCodeAPI{objects: map[string]map[string]interface{}{}}
	f.Server = httptest.NewServer(http.HandlerFunc(f.serve))
	t.Cleanup(f.Server.Close)
	return f
}

// SetProviderEnv points the provider (configured from the environment) at the fake.
func (f *FakeCodeAPI) SetProviderEnv(t *testing.T) {
	t.Helper()
	t.Setenv("HARNESS_ENDPOINT", f.Server.URL)
	t.Setenv("HARNESS_ACCOUNT_ID", "fake_account")
	t.Setenv("HARNESS_PLATFORM_API_KEY", "fake_key")
}

// Requests returns the log of requests served so far, e.g. "GET repos/x -> 404".
func (f *FakeCodeAPI) Requests() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string(nil), f.requests...)
}

// Has reports whether an object exists. path is relative to the API base path,
// e.g. "repos/my_repo" or "repos/my_repo/webhooks/my_hook".
func (f *FakeCodeAPI) Has(path string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	_, ok := f.objects[path]
	return ok
}

// DeleteOutOfBand removes an object the way a user or another tool would, bypassing
// Terraform. Deleting a repo also removes its webhooks and rules.
func (f *FakeCodeAPI) DeleteOutOfBand(path string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.deleteLocked(path)
	f.requests = append(f.requests, fmt.Sprintf("OUT-OF-BAND DELETE %s", path))
}

func (f *FakeCodeAPI) deleteLocked(path string) {
	delete(f.objects, path)
	for key := range f.objects {
		if strings.HasPrefix(key, path+"/") {
			delete(f.objects, key)
		}
	}
}

func (f *FakeCodeAPI) serve(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()

	path := strings.Trim(strings.TrimPrefix(r.URL.Path, fakeCodeAPIBasePath), "/")
	status := f.handleLocked(w, r, path)
	f.requests = append(f.requests, fmt.Sprintf("%s %s -> %d", r.Method, path, status))
}

func (f *FakeCodeAPI) handleLocked(w http.ResponseWriter, r *http.Request, path string) int {
	segs := strings.Split(path, "/")

	// Children (webhooks, rules) only exist while their parent repo exists.
	if len(segs) > 2 {
		if _, ok := f.objects["repos/"+segs[1]]; !ok {
			return writeNotFound(w)
		}
	}

	switch r.Method {
	case http.MethodPost:
		// POST on a collection: repos, repos/{id}/webhooks, repos/{id}/rules.
		if len(segs) != 1 && len(segs) != 3 {
			return writeNotFound(w)
		}
		body := map[string]interface{}{}
		_ = json.NewDecoder(r.Body).Decode(&body)
		id, _ := body["identifier"].(string)
		if id == "" {
			return writeJSON(w, http.StatusBadRequest, map[string]interface{}{"message": "identifier is required"})
		}
		fillDefaults(segs[len(segs)-1], id, body)
		f.objects[path+"/"+id] = body
		return writeJSON(w, http.StatusCreated, body)

	case http.MethodGet:
		obj, ok := f.objects[path]
		if !ok {
			return writeNotFound(w)
		}
		return writeJSON(w, http.StatusOK, obj)

	case http.MethodPatch, http.MethodPut:
		obj, ok := f.objects[path]
		if !ok {
			return writeNotFound(w)
		}
		patch := map[string]interface{}{}
		_ = json.NewDecoder(r.Body).Decode(&patch)
		for k, v := range patch {
			obj[k] = v
		}
		return writeJSON(w, http.StatusOK, obj)

	case http.MethodDelete:
		if _, ok := f.objects[path]; !ok {
			return writeNotFound(w)
		}
		f.deleteLocked(path)
		w.WriteHeader(http.StatusNoContent)
		return http.StatusNoContent
	}

	return writeNotFound(w)
}

func fillDefaults(collection, id string, body map[string]interface{}) {
	switch collection {
	case "repos":
		body["path"] = id
		body["git_url"] = "https://fake.harness.io/git/" + id
		body["importing"] = false
		body["created"] = 1
		body["updated"] = 1
	case "rules":
		body["created_by"] = map[string]interface{}{"id": 1, "display_name": "fake"}
		body["created"] = 1
		body["updated"] = 1
	case "webhooks":
		body["created_by"] = 1
		body["created"] = 1
		body["updated"] = 1
		body["has_secret"] = false
	}
}

func writeNotFound(w http.ResponseWriter) int {
	return writeJSON(w, http.StatusNotFound, map[string]interface{}{"message": "resource not found"})
}

func writeJSON(w http.ResponseWriter, status int, body interface{}) int {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
	return status
}
