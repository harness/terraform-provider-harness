package loadtest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	sdkloadtest "github.com/harness/harness-go-sdk/harness/loadtest"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestChaosLoadTest_SchemaValid(t *testing.T) {
	resource := ResourceChaosLoadTest()
	if err := resource.InternalValidate(resource.Schema, true); err != nil {
		t.Fatal(err)
	}
	if resource.Importer == nil {
		t.Fatal("expected resource to support import")
	}
	dataSource := DataSourceChaosLoadTest()
	if err := dataSource.InternalValidate(dataSource.Schema, false); err != nil {
		t.Fatal(err)
	}
}

func TestLoadTestCreateRequestK6IsFlat(t *testing.T) {
	d := loadTestData(t, map[string]interface{}{
		"identifier":        "checkout",
		"name":              "Checkout",
		"org_id":            "default",
		"project_id":        "demo",
		"tool_type":         "K6",
		"target_type":       "kubernetes",
		"environment_id":    "prod",
		"infrastructure_id": "payments",
		"cleanup_policy":    "retain",
		"tool_config":       `{"mode":"script","script":{"content":"a2s="}}`,
		"tags":              []interface{}{"checkout"},
	})
	req, err := loadTestCreateRequest(d)
	if err != nil {
		t.Fatal(err)
	}
	if req.ToolType != sdkloadtest.LoadTestToolK6 || req.ToolConfig["mode"] != "script" {
		t.Fatalf("request = %+v", req)
	}
	if _, nested := req.ToolConfig["k6"]; nested {
		t.Fatal("k6 toolConfig must be flat")
	}
	if req.EnvironmentIdentifier != "prod" || req.InfraIdentifier != "payments" || req.CleanupPolicy != "retain" {
		t.Fatalf("infra = %+v", req)
	}
}

func TestLoadTestCreateRequestLocustAndManagedInfra(t *testing.T) {
	locust := loadTestData(t, map[string]interface{}{
		"identifier":        "browse",
		"name":              "Browse",
		"org_id":            "default",
		"project_id":        "demo",
		"tool_type":         "Locust",
		"target_type":       "linux",
		"environment_id":    "prod",
		"infrastructure_id": "vm-1",
		"managed_infra":     false,
		"cleanup_policy":    "delete",
		"tool_config":       `{"locust":{"mode":"script","script":{"content":"cHk="}}}`,
	})
	req, err := loadTestCreateRequest(locust)
	if err != nil {
		t.Fatal(err)
	}
	nested, _ := req.ToolConfig["locust"].(map[string]interface{})
	if nested["mode"] != "script" {
		t.Fatalf("locust toolConfig = %#v", req.ToolConfig)
	}

	managed := loadTestData(t, map[string]interface{}{
		"identifier":     "hosted",
		"name":           "Hosted",
		"org_id":         "default",
		"project_id":     "demo",
		"tool_type":      "K6",
		"target_type":    "kubernetes",
		"managed_infra":  true,
		"cleanup_policy": "delete",
		"tool_config":    `{"mode":"script","script":{"content":"a2s="}}`,
	})
	created, err := loadTestCreateRequest(managed)
	if err != nil {
		t.Fatal(err)
	}
	if !created.ManagedInfra || created.EnvironmentIdentifier != "" || created.InfraIdentifier != "" {
		t.Fatalf("managed = %+v", created)
	}

	missing := loadTestData(t, map[string]interface{}{
		"identifier":     "missing",
		"name":           "Missing",
		"org_id":         "default",
		"project_id":     "demo",
		"tool_type":      "K6",
		"managed_infra":  false,
		"cleanup_policy": "delete",
		"tool_config":    `{"mode":"script"}`,
	})
	if _, err := loadTestCreateRequest(missing); err == nil {
		t.Fatal("expected owned infra identifiers to be required")
	}
}

func TestLoadTestResourceLifecycle(t *testing.T) {
	var method, path string
	var gotBody map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		method = r.Method
		path = r.URL.Path
		if r.Body != nil && (r.Method == http.MethodPost || r.Method == http.MethodPut) {
			body, _ := io.ReadAll(r.Body)
			_ = json.Unmarshal(body, &gotBody)
		}
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/variables") {
			_, _ = w.Write([]byte(`{"loadTestIdentity":"checkout","variables":[]}`))
			return
		}
		switch {
		case r.Method == http.MethodPost:
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(loadTestJSON("Checkout")))
		case r.Method == http.MethodGet:
			_, _ = w.Write([]byte(loadTestJSON("Checkout")))
		case r.Method == http.MethodPut:
			_, _ = w.Write([]byte(loadTestJSON("Checkout v2")))
		case r.Method == http.MethodDelete:
			_, _ = w.Write([]byte(`{"success":true}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

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
		"tool_config":       `{"mode":"script","script":{"content":"a2s="}}`,
	})
	meta := loadTestSession(server)
	diags := resourceChaosLoadTestCreate(t.Context(), d, meta)
	if diags.HasError() {
		t.Fatal(diags)
	}
	if method != http.MethodGet || path != "/loadTest/manager/api/v1/load-tests/checkout/variables" {
		t.Fatalf("after create, last call %s %s", method, path)
	}
	if d.Id() != "checkout" || d.Get("unique_id") != "uid-1" {
		t.Fatalf("state id %s unique %s", d.Id(), d.Get("unique_id"))
	}
	if _, nested := gotBody["toolConfig"].(map[string]interface{})["k6"]; nested {
		t.Fatalf("create body = %#v", gotBody["toolConfig"])
	}
	if gotBody["environmentIdentifier"] != "prod" || gotBody["infraIdentifier"] != "payments" {
		t.Fatalf("create body = %#v", gotBody)
	}

	if err := d.Set("name", "Checkout v2"); err != nil {
		t.Fatal(err)
	}
	diags = resourceChaosLoadTestUpdate(t.Context(), d, meta)
	if diags.HasError() {
		t.Fatal(diags)
	}
	if gotBody["name"] != "Checkout v2" || gotBody["toolConfig"] == nil {
		t.Fatalf("update body = %#v", gotBody)
	}
	if _, ok := gotBody["variables"].([]interface{}); !ok {
		t.Fatalf("update variables = %#v", gotBody["variables"])
	}

	diags = resourceChaosLoadTestDelete(t.Context(), d, meta)
	if diags.HasError() || d.Id() != "" || method != http.MethodDelete {
		t.Fatalf("delete id %q method %s diags %v", d.Id(), method, diags)
	}
}

func TestParseLoadTestImportID(t *testing.T) {
	org, project, id, err := parseLoadTestImportID("default/demo/checkout")
	if err != nil || org != "default" || project != "demo" || id != "checkout" {
		t.Fatalf("parsed %s %s %s err %v", org, project, id, err)
	}
	if _, _, _, err := parseLoadTestImportID("default/checkout"); err == nil {
		t.Fatal("expected an invalid import id to fail")
	}
	if _, _, _, err := parseLoadTestImportID("default/demo/checkout?v2"); err == nil {
		t.Fatal("expected an identifier containing ? to fail")
	}
	if _, _, _, err := parseLoadTestImportID("default/demo/team/checkout"); err == nil {
		t.Fatal("expected an identifier containing / to fail")
	}
}

func TestCreateTimeoutSaysToDeleteBeforeRetry(t *testing.T) {
	client := sdkloadtest.NewAPIClient(&sdkloadtest.Configuration{
		BasePath:   "http://loadtest.invalid/loadTest/manager/api",
		AccountId:  "acc",
		HTTPClient: &http.Client{Transport: roundTripError{}},
	})
	d := loadTestData(t, map[string]interface{}{
		"identifier":        "checkout",
		"name":              "Checkout",
		"org_id":            "default",
		"project_id":        "demo",
		"tool_type":         "K6",
		"target_type":       "kubernetes",
		"cleanup_policy":    "delete",
		"managed_infra":     false,
		"environment_id":    "prod",
		"infrastructure_id": "payments",
		"tool_config":       `{"mode":"script","script":{"content":"a2s="}}`,
	})
	diags := resourceChaosLoadTestCreate(t.Context(), d, &internal.Session{LoadTestClient: client})
	if !diags.HasError() || d.Id() != "" {
		t.Fatalf("id %q diags %v", d.Id(), diags)
	}
	if !strings.Contains(diagnosticSummary(diags), "delete that load test before the next apply") {
		t.Fatalf("diags = %v", diags)
	}
}

type roundTripError struct{}

func (roundTripError) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("connection timed out")
}

func TestNestedK6ModeIsAccepted(t *testing.T) {
	err := validateToolMode(sdkloadtest.LoadTestToolK6, map[string]interface{}{
		"k6": map[string]interface{}{
			"mode":   "image",
			"script": map[string]interface{}{"image": "ghcr.io/example/k6:1"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestScriptValidationListsEveryFinding(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"Script validation failed","message":"first problem","errorCode":"E2022","validation":{"findings":[{"severity":"error","message":"first problem","line":3,"column":1},{"severity":"error","message":"second problem","line":9,"column":2}]}}`))
	}))
	defer server.Close()
	d := loadTestData(t, map[string]interface{}{
		"identifier": "checkout", "name": "Checkout", "org_id": "default", "project_id": "demo",
		"tool_type": "K6", "target_type": "kubernetes", "cleanup_policy": "delete", "managed_infra": false,
		"environment_id": "prod", "infrastructure_id": "payments",
		"tool_config": `{"mode":"script","script":{"content":"a2s="}}`,
	})
	diags := resourceChaosLoadTestCreate(t.Context(), d, loadTestSession(server))
	summary := diagnosticSummary(diags)
	if !strings.Contains(summary, "first problem") || !strings.Contains(summary, "second problem") {
		t.Fatalf("diags = %v", diags)
	}
	if d.Id() != "" {
		t.Fatalf("id = %q", d.Id())
	}
}

func TestLoadTestIdentifierRejectsReservedURLCharacters(t *testing.T) {
	resource := ResourceChaosLoadTest()
	for _, identifier := range []string{"checkout?v2", "checkout#v2", "team/checkout", "checkout v2"} {
		_, errs := resource.Schema["identifier"].ValidateFunc(identifier, "identifier")
		if len(errs) == 0 {
			t.Fatalf("identifier %q was accepted", identifier)
		}
	}
	if _, errs := resource.Schema["identifier"].ValidateFunc("sample-products-api", "identifier"); len(errs) != 0 {
		t.Fatal(errs)
	}
}

func loadTestData(t *testing.T, raw map[string]interface{}) *schema.ResourceData {
	t.Helper()
	return schema.TestResourceDataRaw(t, resourceChaosLoadTestSchema(), raw)
}

func loadTestSession(server *httptest.Server) *internal.Session {
	client := sdkloadtest.NewAPIClient(&sdkloadtest.Configuration{
		BasePath:   server.URL + "/loadTest/manager/api",
		AccountId:  "acc",
		HTTPClient: server.Client(),
	})
	return &internal.Session{LoadTestClient: client}
}

func loadTestJSON(name string) string {
	body := map[string]interface{}{
		"identity":                 "checkout",
		"uniqueId":                 "uid-1",
		"name":                     name,
		"toolType":                 "K6",
		"targetType":               "kubernetes",
		"environmentIdentifier":    "prod",
		"infraIdentifier":          "payments",
		"cleanupPolicy":            "delete",
		"latestRevisionIdentifier": "rev-1",
		"toolConfig":               map[string]interface{}{"mode": "script", "script": map[string]interface{}{"content": "a2s="}},
	}
	buf, _ := json.Marshal(body)
	return string(buf)
}
