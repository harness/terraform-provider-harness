package idp_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/harness/terraform-provider-harness/internal/provider"
	"github.com/harness/terraform-provider-harness/internal/service/platform/idp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// fakeIntegrationsAPI is an in-memory integrations API: revisions, If-Match, defaults and mappings.
type fakeIntegrationsAPI struct {
	mu           sync.Mutex
	integrations map[string]map[string]any
	mappings     map[string]map[string]any
	requests     []string
	nextUID      int
}

func newFakeIntegrationsAPI(t *testing.T) *httptest.Server {
	f := &fakeIntegrationsAPI{integrations: map[string]map[string]any{}, mappings: map[string]map[string]any{}}
	srv := httptest.NewServer(f)
	t.Cleanup(srv.Close)
	t.Setenv("HARNESS_ENDPOINT", srv.URL)
	t.Setenv("HARNESS_ACCOUNT_ID", "acc")
	t.Setenv("HARNESS_PLATFORM_API_KEY", "pat.acc.test.key")
	return srv
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func apiError(w http.ResponseWriter, status int, code, msg string) {
	writeJSON(w, status, map[string]any{"code": code, "message": msg, "details": []any{}})
}

func (f *fakeIntegrationsAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.requests = append(f.requests, r.Method+" "+r.URL.Path)
	if r.Header.Get("x-api-key") == "" {
		apiError(w, http.StatusUnauthorized, "UNAUTHORIZED", "no key")
		return
	}
	const prefix = "/integration/api/v1/accounts/acc"
	path := strings.TrimPrefix(r.URL.Path, prefix)
	if path == r.URL.Path {
		apiError(w, http.StatusNotFound, "NOT_FOUND", "bad prefix "+r.URL.Path)
		return
	}
	if strings.HasPrefix(path, "/integration-types/") {
		f.serveType(w, strings.TrimPrefix(path, "/integration-types/"))
		return
	}
	scope, rest, ok := strings.Cut(path, "/integrations")
	if !ok {
		apiError(w, http.StatusNotFound, "NOT_FOUND", path)
		return
	}
	org, project := "", ""
	if parts := strings.Split(strings.Trim(scope, "/"), "/"); len(parts) >= 2 {
		org = parts[1]
		if len(parts) == 4 {
			project = parts[3]
		}
	}
	parts := strings.Split(strings.Trim(rest, "/"), "/")
	switch {
	case rest == "" && r.Method == http.MethodGet:
		f.list(w, scope)
	case rest == "" && r.Method == http.MethodPost:
		f.create(w, r, scope, org, project)
	case len(parts) == 1:
		f.item(w, r, scope+"/"+parts[0])
	case len(parts) == 4 && parts[1] == "kinds" && parts[3] == "mapping":
		f.mapping(w, r, scope+"/"+parts[0], parts[2])
	case len(parts) == 2 && parts[1] == "catalog":
		writeJSON(w, http.StatusOK, map[string]any{"version": "v1", "streams": []any{map[string]any{
			"name": "cmdb_ci_service", "supported_sync_modes": []string{"full_refresh", "incremental"},
			"sync_mode": "incremental", "primary_key": [][]string{{"sys_id"}},
			"json_schema": map[string]any{"type": "object"},
		}}})
	default:
		apiError(w, http.StatusNotFound, "NOT_FOUND", path)
	}
}

func (f *fakeIntegrationsAPI) serveType(w http.ResponseWriter, typ string) {
	writeJSON(w, http.StatusOK, map[string]any{
		"type": typ, "display_name": typ, "category": "itsm", "mode": "airbyte", "connector": "required",
		"allowed_scopes": []string{"account", "org", "project"}, "singleton": false,
		"kinds": []any{map[string]any{
			"kind": "cmdb_ci_service", "secondary": false, "sync_modes": []string{"incremental"},
			"default_action": "register", "mapping_fields": []any{map[string]any{
				"field_name": "owner", "display_name": "Owner", "description": "", "target_path": "spec.owner",
				"required": false,
			}},
		}},
	})
}

func (f *fakeIntegrationsAPI) list(w http.ResponseWriter, scope string) {
	items := []any{}
	for key, in := range f.integrations {
		if strings.HasPrefix(key, scope+"/") && !strings.Contains(strings.TrimPrefix(key, scope+"/"), "/") {
			items = append(items, map[string]any{
				"identifier": in["identifier"], "name": in["name"], "type": in["type"], "mode": "airbyte",
				"enabled": in["enabled"], "org": in["org"], "project": in["project"], "kinds": []string{},
				"health": "healthy",
			})
		}
	}
	writeJSON(w, http.StatusOK, map[string]any{"items": items, "page": map[string]any{"total": len(items)}})
}

func checkIfMatch(w http.ResponseWriter, r *http.Request, revision float64) bool {
	want := strconv.Quote(strconv.Itoa(int(revision)))
	if got := r.Header.Get("If-Match"); got != "" && got != want {
		apiError(w, http.StatusPreconditionFailed, "REVISION_MISMATCH", "revision is "+want)
		return false
	}
	return true
}

// applyDefaults mimics the server filling in omitted values.
func applyDefaults(in map[string]any) {
	spec, _ := in["spec"].(map[string]any)
	if _, ok := spec["update_frequency"]; !ok && in["type"] == "ServiceNow" {
		spec["update_frequency"] = "24h"
	}
	kinds, _ := spec["kinds"].([]any)
	// The server returns kinds sorted by name.
	sort.Slice(kinds, func(i, j int) bool {
		return kinds[i].(map[string]any)["kind"].(string) < kinds[j].(map[string]any)["kind"].(string)
	})
	for _, k := range kinds {
		kind := k.(map[string]any)
		if _, ok := kind["action"]; !ok {
			kind["action"] = "register"
		}
		if _, ok := kind["filter"]; !ok {
			kind["filter"] = ""
		}
		if _, ok := kind["auto_import"]; !ok {
			kind["auto_import"] = false
		}
		for field, def := range map[string]string{"idp_kind": "component", "idp_type": "service"} {
			if _, ok := kind[field]; !ok && in["type"] == "ServiceNow" {
				kind[field] = def
			}
		}
	}
	if gh, ok := spec["github"].(map[string]any); ok {
		if _, ok := gh["agent_file_name"]; !ok {
			gh["agent_file_name"] = "AGENTS.md"
		}
	}
	in["status"] = map[string]any{"mode": "airbyte", "health": "healthy", "effective_kinds": []string{},
		"endpoints": []any{map[string]any{"kind": "x", "method": "POST", "url": "/integrations/x/data/x"}}}
}

func (f *fakeIntegrationsAPI) create(w http.ResponseWriter, r *http.Request, scope, org, project string) {
	var in map[string]any
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		apiError(w, http.StatusBadRequest, "INVALID_JSON", err.Error())
		return
	}
	key := scope + "/" + in["identifier"].(string)
	if _, ok := f.integrations[key]; ok {
		apiError(w, http.StatusConflict, "ALREADY_EXISTS", "exists")
		return
	}
	// Create mappings are stored per kind, as the API does, and aren't part of the resource.
	if mappings, ok := in["mappings"].(map[string]any); ok {
		for kind, m := range mappings {
			stored := defaultMapping(kind)
			for k, v := range m.(map[string]any) {
				stored[k] = v
			}
			trimCorrelation(stored)
			stored["is_default"] = false
			f.mappings[key+"/"+kind] = stored
		}
		delete(in, "mappings")
	}
	f.nextUID++
	in["uid"] = fmt.Sprintf("%s-x%d", strings.ToLower(in["identifier"].(string)), f.nextUID)
	in["revision"], in["org"], in["project"], in["created"], in["updated"] = float64(1), org, project, 1, 1
	applyDefaults(in)
	f.integrations[key] = in
	writeJSON(w, http.StatusCreated, in)
}

func (f *fakeIntegrationsAPI) item(w http.ResponseWriter, r *http.Request, key string) {
	cur, ok := f.integrations[key]
	if !ok {
		apiError(w, http.StatusNotFound, "NOT_FOUND", "integration not found")
		return
	}
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, cur)
	case http.MethodPut:
		if !checkIfMatch(w, r, cur["revision"].(float64)) {
			return
		}
		var in map[string]any
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			apiError(w, http.StatusBadRequest, "INVALID_JSON", err.Error())
			return
		}
		if in["identifier"] != cur["identifier"] {
			apiError(w, http.StatusUnprocessableEntity, "IMMUTABLE_FIELD", "identifier")
			return
		}
		for _, k := range []string{"uid", "org", "project", "created"} {
			in[k] = cur[k]
		}
		in["revision"], in["updated"] = cur["revision"].(float64)+1, 2
		applyDefaults(in)
		f.integrations[key] = in
		writeJSON(w, http.StatusOK, in)
	case http.MethodDelete:
		if !checkIfMatch(w, r, cur["revision"].(float64)) {
			return
		}
		delete(f.integrations, key)
		w.WriteHeader(http.StatusNoContent)
	}
}

func (f *fakeIntegrationsAPI) mapping(w http.ResponseWriter, r *http.Request, key, kind string) {
	cur, ok := f.integrations[key]
	if !ok {
		apiError(w, http.StatusNotFound, "NOT_FOUND", "integration not found")
		return
	}
	m, ok := f.mappings[key+"/"+kind]
	if !ok {
		m = defaultMapping(kind)
	}
	m["revision"], m["updated"] = cur["revision"], 1
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, m)
	case http.MethodPut:
		if !checkIfMatch(w, r, cur["revision"].(float64)) {
			return
		}
		var in map[string]any
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			apiError(w, http.StatusBadRequest, "INVALID_JSON", err.Error())
			return
		}
		for k, v := range in {
			m[k] = v
		}
		if _, ok := in["correlation"]; !ok {
			delete(m, "correlation")
		}
		trimCorrelation(m)
		cur["revision"] = cur["revision"].(float64) + 1
		m["revision"], m["is_default"] = cur["revision"], false
		f.mappings[key+"/"+kind] = m
		writeJSON(w, http.StatusOK, m)
	}
}

func trimCorrelation(m map[string]any) {
	if c, ok := m["correlation"].(map[string]any); ok {
		for _, k := range []string{"source_path", "target_path"} {
			c[k] = strings.TrimPrefix(c[k].(string), ".")
		}
	}
}

func defaultMapping(kind string) map[string]any {
	return map[string]any{"kind": kind, "unique_identifier_field": "sys_id", "display_name_field": "name",
		"idp_kind": "component", "idp_type": "service", "is_default": true,
		"field_mappings": []any{map[string]any{"source_field": "owner", "target_field": "spec.owner"}}}
}

func unitProviderFactories() map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"harness": func() (*schema.Provider, error) { return provider.Provider("dev")(), nil },
	}
}

const snowConfig = `
resource "harness_platform_idp_integration" "snow" {
  identifier = "Snow_Prod"
  name       = %q
  org_id     = "default"
  type       = "ServiceNow"
  tags       = ["team:idp", "tier"]

  connector {
    inline {
      credentials_ref = "account.snow"
    }
  }

  kind {
    kind             = "cmdb_ci_service"
    auto_import      = true
    replication_mode = "incremental"
    filter           = %q
    idp_kind         = "component"
    idp_type         = %q

    mapping {
      field_mapping {
        source_field = %q
        target_field = "spec.owner"
      }
      correlation {
        source_path = ".service"
        target_path = "metadata.identifier"
      }
    }
  }
  kind {
    kind = "cmdb_ci_server"
  }
}

data "harness_platform_idp_integration" "snow" {
  identifier = harness_platform_idp_integration.snow.identifier
  org_id     = "default"
}

data "harness_platform_idp_integration_kind_mapping" "svc" {
  integration_id = harness_platform_idp_integration.snow.identifier
  org_id         = "default"
  kind           = "cmdb_ci_service"
}

data "harness_platform_idp_integrations" "all" {
  org_id     = "default"
  depends_on = [harness_platform_idp_integration.snow]
}

data "harness_platform_idp_integration_type" "snow" {
  type = "ServiceNow"
}

data "harness_platform_idp_integration_catalog" "snow" {
  integration_id = harness_platform_idp_integration.snow.identifier
  org_id         = "default"
}
`

func TestUnitIntegrationLifecycle(t *testing.T) {
	newFakeIntegrationsAPI(t)
	const res = "harness_platform_idp_integration.snow"
	const mapping = "data.harness_platform_idp_integration_kind_mapping.svc"
	const svc = "kind.0.mapping.0."
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: unitProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(snowConfig, "ServiceNow", ".active", "service", "owned_by"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "id", "Snow_Prod"),
					resource.TestCheckResourceAttr(res, "identifier", "Snow_Prod"),
					resource.TestCheckResourceAttr(res, "uid", "snow_prod-x1"),
					// The mapping went with the create, so there was no second write.
					resource.TestCheckResourceAttr(res, "revision", "1"),
					resource.TestCheckResourceAttr(res, svc+"unique_identifier_field", "sys_id"),
					resource.TestCheckResourceAttr(res, svc+"correlation.0.source_path", "service"),
					resource.TestCheckResourceAttr(res, svc+"field_mapping.0.source_field", "owned_by"),
					// Read, but not managed, without a block.
					resource.TestCheckResourceAttr(res, "kind.1.mapping.0.unique_identifier_field", "sys_id"),
					resource.TestCheckResourceAttr(res, "update_frequency", "24h"),
					resource.TestCheckResourceAttr(res, "kind.0.action", "register"),
					resource.TestCheckResourceAttr(res, "kind.1.replication_mode", ""),
					resource.TestCheckResourceAttr(res, "kind.0.idp_type", "service"),
					resource.TestCheckResourceAttr(res, "kind.1.idp_kind", "component"),
					resource.TestCheckResourceAttr(res, "tags.#", "2"),
					resource.TestCheckResourceAttr(res, "endpoints.0.method", "POST"),
					resource.TestCheckResourceAttr(mapping, "idp_type", "service"),
					resource.TestCheckResourceAttr(mapping, "idp_kind", "component"),
					resource.TestCheckResourceAttr(mapping, "unique_identifier_field", "sys_id"),
					resource.TestCheckResourceAttr(mapping, "correlation.0.target_path", "metadata.identifier"),
					resource.TestCheckResourceAttr(mapping, "correlation.0.source_path", "service"),
					resource.TestCheckResourceAttr(mapping, "is_default", "false"),
					resource.TestCheckResourceAttr("data.harness_platform_idp_integration.snow", "revision", "1"),
					resource.TestCheckResourceAttr("data.harness_platform_idp_integration.snow", "kind.#", "2"),
					resource.TestCheckResourceAttr("data.harness_platform_idp_integrations.all", "integrations.#", "1"),
					resource.TestCheckResourceAttr("data.harness_platform_idp_integrations.all",
						"integrations.0.identifier", "Snow_Prod"),
					resource.TestCheckResourceAttr("data.harness_platform_idp_integration_type.snow",
						"kinds.0.mapping_fields.0.target_path", "spec.owner"),
					resource.TestCheckResourceAttr("data.harness_platform_idp_integration_catalog.snow",
						"streams.0.primary_key.0", "sys_id"),
				),
			},
			{
				// The mapping write moved the revision to 2, so the update must read it rather than use state.
				Config: fmt.Sprintf(snowConfig, "ServiceNow prod", ".active == true", "service", "managed_by"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", "ServiceNow prod"),
					resource.TestCheckResourceAttr(res, "kind.0.filter", ".active == true"),
					resource.TestCheckResourceAttr(res, "uid", "snow_prod-x1"),
					resource.TestCheckResourceAttr(res, svc+"field_mapping.0.source_field", "managed_by"),
					// The integration write, then the changed mapping.
					resource.TestCheckResourceAttr(res, "revision", "3"),
				),
			},
			{
				// An unchanged mapping isn't rewritten.
				Config: fmt.Sprintf(snowConfig, "ServiceNow", ".active == true", "service", "managed_by"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "revision", "4"),
					resource.TestCheckResourceAttr(res, svc+"field_mapping.0.source_field", "managed_by"),
				),
			},
			{
				Config:      fmt.Sprintf(snowConfig, "ServiceNow prod", ".active == true", "database", "managed_by"),
				ExpectError: regexpMust(`idp_type of kind cmdb_ci_service can't change from "service" to "database"`),
			},
			{
				ResourceName:      res,
				ImportState:       true,
				ImportStateId:     "default/Snow_Prod",
				ImportStateVerify: true,
				// Import has no prior order, so kinds come back sorted by name.
				ImportStateVerifyIgnore: []string{"revision", "updated", "kind."},
			},
		},
		CheckDestroy: func(s *terraform.State) error { return nil },
	})
}

func TestUnitIntegrationPlanChecks(t *testing.T) {
	newFakeIntegrationsAPI(t)
	base := `
resource "harness_platform_idp_integration" "gh" {
  identifier = "gh"
  name       = "GitHub"
  type       = "GitHub"
  github {
    orgs = ["acme"]
  }
  %s
}
`
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: unitProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(base, `kind { kind = "repository" }`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("harness_platform_idp_integration.gh", "github.0.agent_file_name",
						"AGENTS.md"),
				),
			},
			{
				Config:      fmt.Sprintf(base, `kind { kind = "repository" }`+"\n"+`kind { kind = "team" }`),
				ExpectError: regexpMust("kinds can't be added or removed"),
			},
			{
				Config: fmt.Sprintf(base, `kind { kind = "repository" }`+"\n"+`sonarqube {
    token_ref = "x"
    instance {
      cloud { organization = "acme" }
    }
  }`),
				ExpectError: regexpMust("sonarqube is only allowed when type is SonarQube"),
			},
			{
				Config: fmt.Sprintf(base, `kind { kind = "repository" }`+"\n"+`custom {
    mode = "api"
  }`),
				ExpectError: regexpMust("custom is only allowed when type is Custom"),
			},
		},
	})
}

func regexpMust(s string) *regexp.Regexp { return regexp.MustCompile(regexp.QuoteMeta(s)) }

func TestUnitIntegrationIdentifierValidation(t *testing.T) {
	validate := idp.ResourceIntegration().Schema["identifier"].ValidateFunc
	for id, ok := range map[string]bool{"Snow_Prod": true, "_a$1": true, "1gh": false, "$a": false, "my-hook": false, "": false,
		"true": false, "account": false, "projects": false} {
		if _, errs := validate(id, "identifier"); (len(errs) == 0) != ok {
			t.Errorf("identifier %q: valid = %v, want %v", id, len(errs) == 0, ok)
		}
	}
}
