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
	case rest == "" && r.Method == http.MethodPost:
		f.create(w, r, scope, org, project)
	case len(parts) == 1:
		f.item(w, r, scope+"/"+parts[0])
	case len(parts) == 4 && parts[1] == "kinds" && parts[3] == "mapping":
		f.mapping(w, r, scope+"/"+parts[0], parts[2])
	default:
		apiError(w, http.StatusNotFound, "NOT_FOUND", path)
	}
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
	// The server drops kindless entries, which is how a type without kinds is sent.
	all, _ := spec["kinds"].([]any)
	kinds := []any{}
	for _, k := range all {
		if k.(map[string]any)["kind"] != nil {
			kinds = append(kinds, k)
		}
	}
	spec["kinds"] = kinds
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
	if in["type"] == "Jira" {
		jira, _ := spec["jira"].(map[string]any)
		if jira == nil {
			jira = map[string]any{}
			spec["jira"] = jira
		}
		if start, _ := jira["start_date"].(string); start == "" {
			jira["start_date"] = "2025-10-15T00:00:00Z"
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
	// The UI sends no identifier, so the server uses its generated uid, which contains -.
	generated := in["identifier"] == nil || in["identifier"] == ""
	if generated {
		in["identifier"] = fmt.Sprintf("%s-x%d", strings.ToLower(strings.ReplaceAll(in["name"].(string), " ", "-")),
			f.nextUID+1)
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
			f.mappings[key+"/"+kind] = stored
		}
		delete(in, "mappings")
	}
	f.nextUID++
	in["uid"] = fmt.Sprintf("%s-x%d", strings.ToLower(in["identifier"].(string)), f.nextUID)
	if generated {
		in["uid"] = in["identifier"]
	}
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
		// Omitted field mappings are the type's defaults.
		if _, ok := in["field_mappings"]; !ok {
			m["field_mappings"] = defaultMapping(kind)["field_mappings"]
		}
		trimCorrelation(m)
		cur["revision"] = cur["revision"].(float64) + 1
		m["revision"] = cur["revision"]
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
		"idp_kind": "component", "idp_type": "service",
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
`

const svcMappingConfig = `
resource "harness_platform_idp_integration_kind_mapping" "svc" {
  integration_id = harness_platform_idp_integration.snow.identifier
  org_id         = "default"
  kind           = "cmdb_ci_service"

  field_mapping {
    source_field = "managed_by"
    target_field = "spec.owner"
  }
}
`

func TestUnitIntegrationLifecycle(t *testing.T) {
	srv := newFakeIntegrationsAPI(t)
	const res = "harness_platform_idp_integration.snow"
	const mapping = "data.harness_platform_idp_integration_kind_mapping.svc"
	const mappingRes = "harness_platform_idp_integration_kind_mapping.svc"
	const svc = "kind.0.mapping.0."
	created := fmt.Sprintf(snowConfig, "ServiceNow", ".active", "service", "owned_by")
	renamed := fmt.Sprintf(snowConfig, "ServiceNow prod", ".active == true", "service", "owned_by")
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: unitProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: created,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "id", "Snow_Prod"),
					resource.TestCheckResourceAttr(res, "identifier", "Snow_Prod"),
					resource.TestCheckResourceAttr(res, "uid", "snow_prod-x1"),
					// The mapping went with the create, so there was no second write.
					resource.TestCheckResourceAttr(res, "revision", "1"),
					// The block is kept as configured, not read back.
					resource.TestCheckResourceAttr(res, svc+"unique_identifier_field", ""),
					resource.TestCheckResourceAttr(res, svc+"correlation.0.source_path", ".service"),
					resource.TestCheckResourceAttr(res, svc+"field_mapping.0.source_field", "owned_by"),
					resource.TestCheckResourceAttr(res, "kind.1.mapping.#", "0"),
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
					resource.TestCheckResourceAttr(mapping, "field_mapping.0.source_field", "owned_by"),
					resource.TestCheckResourceAttr("data.harness_platform_idp_integration.snow", "revision", "1"),
					resource.TestCheckResourceAttr("data.harness_platform_idp_integration.snow", "kind.#", "2"),
				),
			},
			{
				Config: renamed,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "name", "ServiceNow prod"),
					resource.TestCheckResourceAttr(res, "kind.0.filter", ".active == true"),
					resource.TestCheckResourceAttr(res, "uid", "snow_prod-x1"),
					resource.TestCheckResourceAttr(res, svc+"field_mapping.0.source_field", "owned_by"),
					resource.TestCheckResourceAttr(res, "revision", "2"),
				),
			},
			{
				Config:      fmt.Sprintf(snowConfig, "ServiceNow prod", ".active == true", "service", "managed_by"),
				ExpectError: regexp.MustCompile(`the mapping block of kind cmdb_ci_service is only sent at create`),
			},
			{
				// Removing the block only changes the state.
				Config: withoutMappings(renamed),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "revision", "2"),
					resource.TestCheckResourceAttr(res, "kind.0.mapping.#", "0"),
					resource.TestCheckResourceAttr(mapping, "field_mapping.0.source_field", "owned_by"),
				),
			},
			{
				Config: withoutMappings(renamed) + svcMappingConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(mappingRes, "id", "default/Snow_Prod/cmdb_ci_service"),
					resource.TestCheckResourceAttr(mappingRes, "field_mapping.0.source_field", "managed_by"),
					// The resource owns the whole mapping, so the correlation the create sent is gone.
					resource.TestCheckResourceAttr(mappingRes, "correlation.#", "0"),
					resource.TestCheckResourceAttr(mappingRes, "unique_identifier_field", "sys_id"),
					resource.TestCheckResourceAttr(mappingRes, "revision", "3"),
				),
			},
			{
				Config: withoutMappings(fmt.Sprintf(snowConfig, "ServiceNow prod", ".active == true", "database",
					"managed_by")) + svcMappingConfig,
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
			{
				ResourceName:      mappingRes,
				ImportState:       true,
				ImportStateId:     "default/Snow_Prod/cmdb_ci_service",
				ImportStateVerify: true,
			},
			{
				// Destroying the mapping resource restores the default field mappings of an Airbyte type.
				Config: withoutMappings(renamed),
				Check: func(*terraform.State) error {
					m := getFake(t, srv.URL+"/integration/api/v1/accounts/acc/orgs/default/integrations/Snow_Prod/"+
						"kinds/cmdb_ci_service/mapping")
					fms, _ := m["field_mappings"].([]any)
					if len(fms) != 1 || fms[0].(map[string]any)["source_field"] != "owner" {
						return fmt.Errorf("field_mappings = %v, want the default", m["field_mappings"])
					}
					if m["unique_identifier_field"] != "sys_id" {
						return fmt.Errorf("unique_identifier_field = %v, want it kept", m["unique_identifier_field"])
					}
					return nil
				},
			},
		},
		CheckDestroy: func(s *terraform.State) error { return nil },
	})
}

func getFake(t *testing.T, url string) map[string]any {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("x-api-key", "test")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var out map[string]any
	_ = json.NewDecoder(resp.Body).Decode(&out)
	return out
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

func TestUnitIntegrationJiraSettings(t *testing.T) {
	newFakeIntegrationsAPI(t)
	const res = "harness_platform_idp_integration.jira"
	config := `
resource "harness_platform_idp_integration" "jira" {
  identifier = "jira"
  name       = "Jira"
  type       = "Jira"
  connector {
    inline { credentials_ref = "account.jira" }
  }
  kind { kind = "projects" }
  %s
}
`
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: unitProviderFactories(),
		Steps: []resource.TestStep{
			{
				// The server's default start date is read back without a diff.
				Config: fmt.Sprintf(config, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "jira.0.start_date", "2025-10-15T00:00:00Z"),
					resource.TestCheckResourceAttr(res, "jira.0.custom_fields.#", "0"),
				),
			},
			{
				// So is a block that sets only custom_fields.
				Config: fmt.Sprintf(config, `jira {
    custom_fields = ["customfield_10456"]
  }`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "jira.0.start_date", "2025-10-15T00:00:00Z"),
					resource.TestCheckResourceAttr(res, "jira.0.custom_fields.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(config, `jira {
    start_date    = "2024-01-01T00:00:00Z"
    custom_fields = ["customfield_10456", "customfield_10001"]
  }`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "jira.0.start_date", "2024-01-01T00:00:00Z"),
					resource.TestCheckTypeSetElemAttr(res, "jira.0.custom_fields.*", "customfield_10456"),
					resource.TestCheckResourceAttr(res, "jira.0.custom_fields.#", "2"),
				),
			},
			{
				Config: fmt.Sprintf(config, `jira {
    custom_fields = ["customfield_1"]
  }`) + "\n" + `resource "harness_platform_idp_integration" "gh" {
  identifier = "gh2"
  name       = "GitHub"
  type       = "GitHub"
  github { orgs = ["acme"] }
  jira { custom_fields = ["customfield_1"] }
  kind { kind = "repository" }
}
`,
				ExpectError: regexpMust("jira is only allowed when type is Jira"),
			},
		},
	})
}

var mappingBlock = regexp.MustCompile(`(?s)\n    mapping \{.*?\n    \}`)

func withoutMappings(config string) string { return mappingBlock.ReplaceAllString(config, "") }

func TestUnitIntegrationCreateConflict(t *testing.T) {
	newFakeIntegrationsAPI(t)
	const gh = `
resource "harness_platform_idp_integration" "%s" {
  identifier = "gh"
  name       = "GitHub"
  type       = "GitHub"
  github {
    orgs = ["acme"]
  }
  kind { kind = "repository" }
  %s
}
`
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: unitProviderFactories(),
		Steps: []resource.TestStep{{
			Config: fmt.Sprintf(gh, "a", "") +
				fmt.Sprintf(gh, "b", "depends_on = [harness_platform_idp_integration.a]"),
			ExpectError: regexpMust("import it with terraform import"),
		}},
	})
}

func regexpMust(s string) *regexp.Regexp { return regexp.MustCompile(regexp.QuoteMeta(s)) }

func TestUnitIntegrationIdentifierValidation(t *testing.T) {
	validate := idp.ResourceIntegration().Schema["identifier"].ValidateFunc
	for id, ok := range map[string]bool{"Snow_Prod": true, "_a$1": true, "my-hook": true, "1gh": true,
		"170-sanity-ojqobi": true, "-a": false, "$a": false, "my hook": false, "my.hook": false, "": false,
		"true": false, "account": false, "projects": false} {
		if _, errs := validate(id, "identifier"); (len(errs) == 0) != ok {
			t.Errorf("identifier %q: valid = %v, want %v", id, len(errs) == 0, ok)
		}
	}
}

// TestUnitIntegrationImportServerIdentifier imports an integration created in the UI, whose identifier is the
// server's uid and contains -.
func TestUnitIntegrationImportServerIdentifier(t *testing.T) {
	srv := newFakeIntegrationsAPI(t)
	const config = `
resource "harness_platform_idp_integration" "ui" {
  identifier = %q
  name       = "UI GitHub"
  type       = "GitHub"
  connector {
    inline {
      credentials_ref = "account.gh"
    }
  }
  github {
    orgs = ["acme"]
  }
  kind { kind = "repository" }
}
`
	createInUI := func() {
		body := `{"name":"UI GitHub","type":"GitHub","enabled":true,"spec":{"type":"GitHub",` +
			`"connector":{"inline":{"credentials_ref":"account.gh"}},"github":{"orgs":["acme"]},` +
			`"kinds":[{"kind":"repository"}]}}`
		req, _ := http.NewRequest(http.MethodPost, srv.URL+"/integration/api/v1/accounts/acc/integrations",
			strings.NewReader(body))
		req.Header.Set("x-api-key", "pat.acc.test.key")
		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode != http.StatusCreated {
			t.Fatalf("creating the UI integration: %v %v", err, resp)
		}
		resp.Body.Close()
	}
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: unitProviderFactories(),
		Steps: []resource.TestStep{
			{
				PreConfig:          createInUI,
				Config:             fmt.Sprintf(config, "ui-github-x1"),
				ResourceName:       "harness_platform_idp_integration.ui",
				ImportState:        true,
				ImportStateId:      "ui-github-x1",
				ImportStatePersist: true,
			},
			{Config: fmt.Sprintf(config, "ui-github-x1"), PlanOnly: true},
		},
	})
}

func TestUnitIntegrationImportReorderedKinds(t *testing.T) {
	srv := newFakeIntegrationsAPI(t)
	const res = "harness_platform_idp_integration.ui"
	// The kinds aren't sorted by name, as the server returns them.
	const config = `
resource "harness_platform_idp_integration" "ui" {
  identifier = "ui-github-x2"
  name       = "UI GitHub"
  type       = "GitHub"
  connector {
    inline {
      credentials_ref = "account.gh"
    }
  }
  github {
    orgs = ["acme"]
  }
  kind { kind = "repository" }
  kind { kind = "pull_requests" }
}
`
	createInUI := func() {
		body := `{"identifier":"ui-github-x2","name":"UI GitHub","type":"GitHub","enabled":true,` +
			`"spec":{"type":"GitHub","connector":{"inline":{"credentials_ref":"account.gh"}},` +
			`"github":{"orgs":["acme"]},"kinds":[{"kind":"repository","idp_kind":"component","action":"merge"},` +
			`{"kind":"pull_requests","idp_kind":"system","action":"register"}]}}`
		req, _ := http.NewRequest(http.MethodPost, srv.URL+"/integration/api/v1/accounts/acc/integrations",
			strings.NewReader(body))
		req.Header.Set("x-api-key", "pat.acc.test.key")
		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode != http.StatusCreated {
			t.Fatalf("creating the UI integration: %v %v", err, resp)
		}
		resp.Body.Close()
	}
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: unitProviderFactories(),
		Steps: []resource.TestStep{
			{
				PreConfig:          createInUI,
				Config:             config,
				ResourceName:       res,
				ImportState:        true,
				ImportStateId:      "ui-github-x2",
				ImportStatePersist: true,
				ImportStateCheck: func(states []*terraform.InstanceState) error {
					if got := states[0].Attributes["kind.0.kind"]; got != "pull_requests" {
						return fmt.Errorf("kind.0.kind is %q after the import, want the server's order", got)
					}
					return nil
				},
			},
			{Config: config, PlanOnly: true, ExpectNonEmptyPlan: true},
			{
				// Applying the reorder only changes the state, so the revision stays.
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "revision", "1"),
					resource.TestCheckResourceAttr(res, "kind.0.kind", "repository"),
					resource.TestCheckResourceAttr(res, "kind.1.kind", "pull_requests"),
					// Each kind keeps its own server values, not those of the kind at its position before.
					resource.TestCheckResourceAttr(res, "kind.0.idp_kind", "component"),
					resource.TestCheckResourceAttr(res, "kind.0.action", "merge"),
					resource.TestCheckResourceAttr(res, "kind.1.idp_kind", "system"),
					resource.TestCheckResourceAttr(res, "kind.1.action", "register"),
				),
			},
		},
	})
}

func TestUnitIntegrationWithoutKinds(t *testing.T) {
	newFakeIntegrationsAPI(t)
	const res = "harness_platform_idp_integration.traceable"
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: unitProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: `
resource "harness_platform_idp_integration" "traceable" {
  identifier = "traceable"
  name       = "Traceable"
  type       = "HarnessTraceable"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "kind.#", "0"),
					resource.TestCheckResourceAttr(res, "revision", "1"),
				),
			},
			{
				ResourceName:      res,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
