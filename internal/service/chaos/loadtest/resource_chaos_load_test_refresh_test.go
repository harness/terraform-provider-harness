package loadtest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"sync"
	"testing"

	sdkloadtest "github.com/harness/harness-go-sdk/harness/loadtest"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestK6RequiresKubernetesTarget(t *testing.T) {
	d := loadTestData(t, map[string]interface{}{
		"identifier":        "checkout",
		"name":              "Checkout",
		"org_id":            "default",
		"project_id":        "demo",
		"tool_type":         "K6",
		"environment_id":    "prod",
		"infrastructure_id": "payments",
		"cleanup_policy":    "delete",
		"managed_infra":     false,
		"tool_config":       `{"mode":"script","script":{"content":"a2s="}}`,
	})
	if _, err := loadTestCreateRequest(d); err == nil || !strings.Contains(err.Error(), "kubernetes") {
		t.Fatalf("error = %v", err)
	}
}

func TestCreateRejectsVariablesInToolConfig(t *testing.T) {
	raw := `{"mode":"script","script":{"content":"a2s="},"variables":[{"name":"HOST","value":"https://example.com","type":"String"}]}`
	if _, errs := validateToolConfig(raw, "tool_config"); len(errs) == 0 || !strings.Contains(errs[0].Error(), "variables") {
		t.Fatalf("validation errors = %v", errs)
	}
	d := loadTestData(t, map[string]interface{}{
		"identifier":        "checkout",
		"name":              "Checkout",
		"org_id":            "default",
		"project_id":        "demo",
		"tool_type":         "K6",
		"target_type":       "kubernetes",
		"environment_id":    "prod",
		"infrastructure_id": "payments",
		"cleanup_policy":    "delete",
		"managed_infra":     false,
		"tool_config":       raw,
	})
	if _, err := loadTestCreateRequest(d); err == nil || !strings.Contains(err.Error(), "variables") {
		t.Fatalf("error = %v", err)
	}
	nested := `{"locust":{"mode":"script","variables":[{"name":"HOST","value":"https://example.com"}]}}`
	if _, errs := validateToolConfig(nested, "tool_config"); len(errs) == 0 {
		t.Fatal("nested locust variables were accepted")
	}
}

func TestVariableValueKeepsWholeNumbers(t *testing.T) {
	got := variableValueString(float64(10000000))
	if got != "10000000" {
		t.Fatalf("value = %q", got)
	}
	if got := variableValueString(float64(1.5)); got != "1.5" {
		t.Fatalf("fraction = %q", got)
	}
	if got := variableValueString(int64(100)); got != "100" {
		t.Fatalf("int = %q", got)
	}
}

func TestNumberVariableIsSentAsJSONNumber(t *testing.T) {
	d := loadTestData(t, map[string]interface{}{
		"identifier": "checkout", "name": "Checkout", "org_id": "default", "project_id": "demo",
		"tool_type": "K6", "target_type": "kubernetes", "cleanup_policy": "delete", "managed_infra": false,
		"environment_id": "prod", "infrastructure_id": "payments",
		"tool_config": `{"mode":"script","script":{"content":"a2s="}}`,
		"variable": []interface{}{
			map[string]interface{}{"name": "USERS", "type": "Number", "value": "100", "description": "", "required": false},
			map[string]interface{}{"name": "RATE", "type": "Number", "value": "1.5", "description": "", "required": false},
			map[string]interface{}{"name": "PEAK", "type": "Number", "value": "<+input>", "description": "", "required": false},
			map[string]interface{}{"name": "FROM", "type": "Number", "value": "<+loadTest.variables.peak>", "description": "", "required": false},
			map[string]interface{}{"name": "HOST", "type": "String", "value": "100", "description": "", "required": false},
		},
	})
	req, err := loadTestCreateRequest(d)
	if err != nil {
		t.Fatal(err)
	}
	if len(req.Variables) != 5 {
		t.Fatalf("variables = %+v", req.Variables)
	}
	if req.Variables[0].Value != int64(100) {
		t.Fatalf("USERS = %#v", req.Variables[0].Value)
	}
	rate, ok := req.Variables[1].Value.(float64)
	if !ok || rate != 1.5 {
		t.Fatalf("RATE = %#v", req.Variables[1].Value)
	}
	if req.Variables[2].Value != "<+input>" || req.Variables[3].Value != "<+loadTest.variables.peak>" || req.Variables[4].Value != "100" {
		t.Fatalf("expressions = %#v %#v %#v", req.Variables[2].Value, req.Variables[3].Value, req.Variables[4].Value)
	}
	body, err := json.Marshal(req.Variables[0])
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `"value":100`) || strings.Contains(string(body), `"value":"100"`) {
		t.Fatalf("wire = %s", body)
	}
	flattened := flattenVariables([]sdkloadtest.LoadTestVariable{{Name: "USERS", Type: "Number", Value: float64(100)}})
	if flattened[0].(map[string]interface{})["value"] != "100" {
		t.Fatalf("read value = %#v", flattened[0])
	}
}

func TestCreateEmptyIdentityDeletesLoadTest(t *testing.T) {
	var deleted bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodDelete {
			deleted = true
			_, _ = w.Write([]byte(`{"success":true}`))
			return
		}
		if r.Method == http.MethodPost {
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{}`))
			return
		}
		http.NotFound(w, r)
	}))
	defer server.Close()

	d := loadTestData(t, map[string]interface{}{
		"identifier": "checkout", "name": "Checkout", "org_id": "default", "project_id": "demo",
		"tool_type": "K6", "target_type": "kubernetes", "cleanup_policy": "delete", "managed_infra": false,
		"environment_id": "prod", "infrastructure_id": "payments",
		"tool_config": `{"mode":"script","script":{"content":"a2s="}}`,
	})
	diags := resourceChaosLoadTestCreate(t.Context(), d, loadTestSession(server))
	if !diags.HasError() || !strings.Contains(diagnosticSummary(diags), "was deleted") {
		t.Fatalf("diags = %v", diags)
	}
	if !deleted || d.Id() != "" {
		t.Fatalf("deleted %v id %q", deleted, d.Id())
	}
}

func TestNormalizeKeepsAuthoredTunables(t *testing.T) {
	raw := `{"mode":"script","variables":[{"name":"HOST"}],"defaultedTunables":["targetUsers"],"tunables":{"workerMode":"recommended","workerCount":2,"targetUsers":10}}`
	got, err := normalizeToolConfig(raw)
	if err != nil {
		t.Fatal(err)
	}
	for _, kept := range []string{"defaultedTunables", "workerCount", "targetUsers", "workerMode"} {
		if !strings.Contains(got, kept) {
			t.Fatalf("normalize dropped %s: %s", kept, got)
		}
	}
	if strings.Contains(got, "variables") {
		t.Fatalf("normalize kept variables: %s", got)
	}
}

func TestCreateReadFailureDeletesLoadTest(t *testing.T) {
	var deleted bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost:
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"identity":"checkout","name":"Checkout","toolType":"K6","targetType":"kubernetes"}`))
		case strings.HasSuffix(r.URL.Path, "/variables"):
			http.NotFound(w, r)
		case r.Method == http.MethodDelete:
			deleted = true
			_, _ = w.Write([]byte(`{"success":true}`))
		default:
			_, _ = w.Write([]byte(`{"identity":"checkout","name":"Checkout","toolType":"K6","targetType":"kubernetes","toolConfig":{"mode":"script","script":{"content":"a2s="}}}`))
		}
	}))
	defer server.Close()

	d := loadTestData(t, map[string]interface{}{
		"identifier": "checkout", "name": "Checkout", "org_id": "default", "project_id": "demo",
		"tool_type": "K6", "target_type": "kubernetes", "cleanup_policy": "delete", "managed_infra": false,
		"environment_id": "prod", "infrastructure_id": "payments",
		"tool_config": `{"mode":"script","script":{"content":"a2s="}}`,
	})
	diags := resourceChaosLoadTestCreate(t.Context(), d, loadTestSession(server))
	if !diags.HasError() || !strings.Contains(diagnosticSummary(diags), "was deleted") {
		t.Fatalf("diags = %v", diags)
	}
	if !deleted || d.Id() != "" {
		t.Fatalf("deleted %v id %q", deleted, d.Id())
	}
}

func TestCreateReadNotFoundDeletesLoadTest(t *testing.T) {
	var deleted bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && !strings.Contains(r.URL.Path, "/runs"):
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"identity":"checkout","name":"Checkout","toolType":"K6","targetType":"kubernetes"}`))
		case r.Method == http.MethodDelete:
			deleted = true
			_, _ = w.Write([]byte(`{"success":true}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	d := loadTestData(t, map[string]interface{}{
		"identifier": "checkout", "name": "Checkout", "org_id": "default", "project_id": "demo",
		"tool_type": "K6", "target_type": "kubernetes", "cleanup_policy": "delete", "managed_infra": false,
		"environment_id": "prod", "infrastructure_id": "payments",
		"tool_config": `{"mode":"script","script":{"content":"a2s="}}`,
	})
	diags := resourceChaosLoadTestCreate(t.Context(), d, loadTestSession(server))
	if !diags.HasError() || !strings.Contains(diagnosticSummary(diags), "was deleted") {
		t.Fatalf("diags = %v", diags)
	}
	if !deleted || d.Id() != "" {
		t.Fatalf("deleted %v id %q", deleted, d.Id())
	}
}

func TestImportRejectsOtherToolTypes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/variables") {
			_, _ = w.Write([]byte(`{"variables":[]}`))
			return
		}
		_, _ = w.Write([]byte(`{"identity":"checkout","name":"Checkout","toolType":"JMeter","targetType":"kubernetes","toolConfig":{"jmeter":{"mode":"script"}}}`))
	}))
	defer server.Close()

	d := loadTestData(t, map[string]interface{}{
		"identifier": "checkout", "name": "Checkout", "org_id": "default", "project_id": "demo",
		"tool_type": "K6", "target_type": "kubernetes", "cleanup_policy": "delete", "managed_infra": false,
		"environment_id": "prod", "infrastructure_id": "payments",
		"tool_config": `{"mode":"script","script":{"content":"a2s="}}`,
	})
	d.SetId("default/demo/checkout")
	if _, err := importChaosLoadTest(t.Context(), d, loadTestSession(server)); err == nil || !strings.Contains(err.Error(), "JMeter") {
		t.Fatal(err)
	}
	if d.Id() != "" {
		t.Fatalf("id = %q", d.Id())
	}
}

func TestReadNotFoundClearsID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.NotFound(w, r)
	}))
	defer server.Close()

	d := loadTestData(t, map[string]interface{}{
		"identifier": "missing", "name": "Missing", "org_id": "default", "project_id": "demo",
		"tool_type": "Locust", "cleanup_policy": "delete", "managed_infra": false,
		"tool_config": `{"locust":{"mode":"script"}}`,
	})
	d.SetId("missing")
	diags := resourceChaosLoadTestRead(t.Context(), d, loadTestSession(server))
	if diags.HasError() {
		t.Fatal(diags)
	}
	if d.Id() != "" {
		t.Fatalf("id = %q", d.Id())
	}
}

func TestVariablesReadFailureKeepsID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/variables") {
			http.NotFound(w, r)
			return
		}
		_, _ = w.Write([]byte(`{"identity":"checkout","name":"Checkout","toolType":"K6","targetType":"kubernetes","toolConfig":{"mode":"script","script":{"content":"a2s="}}}`))
	}))
	defer server.Close()

	d := loadTestData(t, map[string]interface{}{
		"identifier": "checkout", "name": "Checkout", "org_id": "default", "project_id": "demo",
		"tool_type": "K6", "target_type": "kubernetes", "cleanup_policy": "delete", "managed_infra": false,
		"environment_id": "prod", "infrastructure_id": "payments",
		"tool_config": `{"mode":"script","script":{"content":"a2s="}}`,
	})
	d.SetId("checkout")
	diags := resourceChaosLoadTestRead(t.Context(), d, loadTestSession(server))
	if !diags.HasError() {
		t.Fatal("expected the variables read to fail")
	}
	if d.Id() != "checkout" {
		t.Fatalf("id = %q", d.Id())
	}
}

func TestDataSourceReadsReferenceImport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/variables") {
			_, _ = w.Write([]byte(`{"variables":[]}`))
			return
		}
		_, _ = w.Write([]byte(`{"identity":"from-template","name":"From template","toolType":"Locust","importType":"REFERENCE","targetType":"linux","toolConfig":{"locust":{"mode":"script","script":{"content":"cHk="}}}}`))
	}))
	defer server.Close()

	d := schema.TestResourceDataRaw(t, dataSourceChaosLoadTestSchema(), map[string]interface{}{
		"identifier": "from-template", "org_id": "default", "project_id": "demo",
	})
	diags := dataSourceChaosLoadTestRead(t.Context(), d, loadTestSession(server))
	if diags.HasError() {
		t.Fatal(diags)
	}
	if d.Id() != "from-template" || d.Get("name") != "From template" || d.Get("tool_type") != "Locust" {
		t.Fatalf("state id %q name %q tool %q", d.Id(), d.Get("name"), d.Get("tool_type"))
	}
}

func TestReadRejectsReferenceImport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/variables") {
			_, _ = w.Write([]byte(`{"variables":[]}`))
			return
		}
		_, _ = w.Write([]byte(`{"identity":"from-template","name":"From template","toolType":"Locust","importType":"REFERENCE","targetType":"linux","toolConfig":{"locust":{"mode":"script"}}}`))
	}))
	defer server.Close()

	d := loadTestData(t, map[string]interface{}{
		"identifier": "from-template", "name": "From template", "org_id": "default", "project_id": "demo",
		"tool_type": "Locust", "cleanup_policy": "delete", "managed_infra": false,
		"environment_id": "prod", "infrastructure_id": "vm",
		"tool_config": `{"locust":{"mode":"script"}}`,
	})
	d.SetId("from-template")
	diags := resourceChaosLoadTestRead(t.Context(), d, loadTestSession(server))
	if !diags.HasError() || !strings.Contains(diags[0].Summary, "REFERENCE") {
		t.Fatalf("diags = %v", diags)
	}
}

func TestRefreshStripsServerOwnedToolConfig(t *testing.T) {
	server := newLoadTestServer()
	defer server.Close()

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: loadTestProviderFactories(server),
		Steps: []resource.TestStep{
			{
				Config: locustConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("harness_chaos_load_test.locust", "target_type", "linux"),
				),
			},
			{Config: locustConfig(), PlanOnly: true, ExpectNonEmptyPlan: false},
			{
				Config: k6Config(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("harness_chaos_load_test.k6", "variable.0.name", "HOST"),
					resource.TestCheckResourceAttr("harness_chaos_load_test.k6", "tool_config", `{"mode":"script","script":{"content":"a2s="}}`),
				),
			},
			{Config: k6Config(), PlanOnly: true, ExpectNonEmptyPlan: false},
		},
	})
}

func TestPlanRejectsIncompleteConfig(t *testing.T) {
	server := newLoadTestServer()
	defer server.Close()

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: loadTestProviderFactories(server),
		Steps: []resource.TestStep{
			{Config: k6NoTargetConfig(), PlanOnly: true, ExpectError: regexp.MustCompile("target_type must be kubernetes")},
			{Config: locustNoInfraConfig(), PlanOnly: true, ExpectError: regexp.MustCompile("environment_id and infrastructure_id are required")},
			{Config: jmeterConfig(), PlanOnly: true, ExpectError: regexp.MustCompile("JMeter")},
			{Config: k6VariablesInToolConfig(), PlanOnly: true, ExpectError: regexp.MustCompile("variables")},
		},
	})
}

func TestServiceReferenceReadBackMismatch(t *testing.T) {
	server := newLoadTestServer()
	defer server.Close()
	server.dropServiceReferences.Store(true)

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: loadTestProviderFactories(server),
		Steps: []resource.TestStep{{
			Config:      k6WithServiceConfig(),
			ExpectError: regexp.MustCompile("service_references were not stored.*was deleted"),
		}},
	})
}

func locustConfig() string {
	return `
provider "harness" {}
resource "harness_chaos_load_test" "locust" {
  org_id     = "default"
  project_id = "demo"
  identifier = "browse"
  name       = "Browse"
  tool_type  = "Locust"
  environment_id    = "prod"
  infrastructure_id = "vm-1"
  tool_config = jsonencode({ locust = { mode = "script", script = { content = "cHk=" } } })
}
`
}

func k6Config() string {
	return `
provider "harness" {}
resource "harness_chaos_load_test" "k6" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout"
  name       = "Checkout"
  tool_type   = "K6"
  target_type = "kubernetes"
  environment_id    = "prod"
  infrastructure_id = "payments"
  tool_config = jsonencode({ mode = "script", script = { content = "a2s=" } })
  variable {
    name  = "HOST"
    type  = "String"
    value = "https://example.com"
  }
}
`
}

func jmeterConfig() string {
	return `
provider "harness" {}
resource "harness_chaos_load_test" "jmeter" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout-jmeter"
  name       = "Checkout jmeter"
  tool_type   = "JMeter"
  target_type = "kubernetes"
  environment_id    = "prod"
  infrastructure_id = "payments"
  tool_config = jsonencode({ jmeter = { mode = "script", script = { content = "abc" } } })
}
`
}

func k6NoTargetConfig() string {
	return `
provider "harness" {}
resource "harness_chaos_load_test" "k6" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout"
  name       = "Checkout"
  tool_type  = "K6"
  environment_id    = "prod"
  infrastructure_id = "payments"
  tool_config = jsonencode({ mode = "script", script = { content = "a2s=" } })
}
`
}

func locustNoInfraConfig() string {
	return `
provider "harness" {}
resource "harness_chaos_load_test" "locust" {
  org_id     = "default"
  project_id = "demo"
  identifier = "browse"
  name       = "Browse"
  tool_type  = "Locust"
  tool_config = jsonencode({ locust = { mode = "script", script = { content = "cHk=" } } })
}
`
}

func k6VariablesInToolConfig() string {
	return `
provider "harness" {}
resource "harness_chaos_load_test" "k6" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout"
  name       = "Checkout"
  tool_type   = "K6"
  target_type = "kubernetes"
  environment_id    = "prod"
  infrastructure_id = "payments"
  tool_config = jsonencode({
    mode = "script"
    script = { content = "a2s=" }
    variables = [{ name = "HOST", type = "String", value = "https://example.com" }]
  })
}
`
}

func k6WithServiceConfig() string {
	return `
provider "harness" {}
resource "harness_chaos_load_test" "k6" {
  org_id     = "default"
  project_id = "demo"
  identifier = "checkout-svc"
  name       = "Checkout"
  tool_type   = "K6"
  target_type = "kubernetes"
  environment_id    = "prod"
  infrastructure_id = "payments"
  service_references = ["payments-api"]
  tool_config = jsonencode({ mode = "script", script = { content = "a2s=" } })
}
`
}

type loadTestAPI struct {
	mu                    sync.Mutex
	tests                 map[string]map[string]interface{}
	dropServiceReferences atomicBool
}

type atomicBool struct {
	mu    sync.Mutex
	value bool
}

func (b *atomicBool) Store(v bool) {
	b.mu.Lock()
	b.value = v
	b.mu.Unlock()
}

func (b *atomicBool) Load() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.value
}

func newLoadTestServer() *loadTestAPIServer {
	api := &loadTestAPI{tests: map[string]map[string]interface{}{}}
	server := httptest.NewServer(api)
	return &loadTestAPIServer{Server: server, loadTestAPI: api}
}

type loadTestAPIServer struct {
	*httptest.Server
	*loadTestAPI
}

func (a *loadTestAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if strings.HasSuffix(r.URL.Path, "/variables") {
		identity := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/loadTest/manager/api/v1/load-tests/"), "/variables")
		a.mu.Lock()
		body := a.tests[identity]
		a.mu.Unlock()
		vars := []interface{}{}
		if body != nil {
			if existing, ok := body["variables"].([]interface{}); ok {
				vars = existing
			}
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"loadTestIdentity": identity, "variables": vars})
		return
	}
	if r.URL.Path != "/loadTest/manager/api/v1/load-tests" && !strings.HasPrefix(r.URL.Path, "/loadTest/manager/api/v1/load-tests/") {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodPost:
		var body map[string]interface{}
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &body)
		identity, _ := body["identity"].(string)
		a.mu.Lock()
		a.tests[identity] = body
		a.mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(a.view(identity))
	case http.MethodDelete:
		_, _ = w.Write([]byte(`{"success":true}`))
	case http.MethodGet:
		identity := strings.TrimPrefix(r.URL.Path, "/loadTest/manager/api/v1/load-tests/")
		if a.view(identity) == nil {
			http.NotFound(w, r)
			return
		}
		_ = json.NewEncoder(w).Encode(a.view(identity))
	default:
		http.NotFound(w, r)
	}
}

func (a *loadTestAPI) view(identity string) map[string]interface{} {
	a.mu.Lock()
	defer a.mu.Unlock()
	body := a.tests[identity]
	if body == nil {
		return nil
	}
	raw, _ := json.Marshal(body)
	var view map[string]interface{}
	_ = json.Unmarshal(raw, &view)
	view["uniqueId"] = "uid-1"
	if view["cleanupPolicy"] == nil {
		view["cleanupPolicy"] = "delete"
	}
	if a.dropServiceReferences.Load() {
		delete(view, "serviceReferences")
	}
	tool, _ := view["toolType"].(string)
	cfg, _ := view["toolConfig"].(map[string]interface{})
	switch tool {
	case "K6":
		if vars, ok := view["variables"].([]interface{}); ok && len(vars) > 0 {
			cfg["variables"] = vars
		}
	case "Locust":
		if view["targetType"] == nil || view["targetType"] == "" {
			view["targetType"] = "linux"
		}
	}
	return view
}

func loadTestProviderFactories(server *loadTestAPIServer) map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"harness": func() (*schema.Provider, error) {
			client := sdkloadtest.NewAPIClient(&sdkloadtest.Configuration{
				BasePath:   server.URL + "/loadTest/manager/api",
				AccountId:  "acc",
				HTTPClient: server.Client(),
			})
			return &schema.Provider{
				ResourcesMap: map[string]*schema.Resource{
					"harness_chaos_load_test": ResourceChaosLoadTest(),
				},
				ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
					return &internal.Session{LoadTestClient: client}, nil
				},
			}, nil
		},
	}
}
