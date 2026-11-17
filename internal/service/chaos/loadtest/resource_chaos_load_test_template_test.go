package loadtest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	sdkloadtest "github.com/harness/harness-go-sdk/harness/loadtest"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestTemplateDeleteConflictKeepsTheResource(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"message":"load tests reference this template"}`))
	}))
	defer server.Close()
	d := schema.TestResourceDataRaw(t, resourceChaosLoadTestTemplateSchema(), map[string]interface{}{
		"org_id": "default", "project_id": "demo", "hub_identity": "hub", "identifier": "checkout",
		"name": "Checkout", "revision": "1", "tool_type": "K6",
	})
	d.SetId("checkout")
	diags := resourceChaosLoadTestTemplateDelete(t.Context(), d, loadTestSession(server))
	if !diags.HasError() {
		t.Fatal("expected delete to fail")
	}
	if d.Id() != "checkout" {
		t.Fatalf("id = %q", d.Id())
	}
	if !strings.Contains(diags[0].Summary, "reference") && !strings.Contains(strings.ToLower(diags[0].Summary), "conflict") && !strings.Contains(diags[0].Summary, "409") {
		t.Fatalf("summary = %q", diags[0].Summary)
	}
}

func TestTemplateUpdateFailureDoesNotRecordAuthoredConfig(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPut {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"message":"rejected"}`))
			return
		}
		_, _ = w.Write([]byte(`{"identity":"checkout","name":"Checkout","revision":"1","toolType":"JMeter"}`))
	}))
	defer server.Close()
	d := schema.TestResourceDataRaw(t, resourceChaosLoadTestTemplateSchema(), map[string]interface{}{
		"org_id": "default", "project_id": "demo", "hub_identity": "hub", "identifier": "checkout",
		"name": "Checkout", "revision": "1", "tool_type": "JMeter",
		"tool_config": `{"jmeter":{"mode":"script","script":{"content":"bmV3"}}}`,
	})
	d.SetId("checkout")
	diags := resourceChaosLoadTestTemplateUpdate(t.Context(), d, loadTestSession(server))
	if !diags.HasError() {
		t.Fatal("expected the update to fail")
	}
	if d.Get("authored_tool_config").(string) != "" {
		t.Fatalf("authored_tool_config = %q", d.Get("authored_tool_config"))
	}
}

func TestTemplateReadKeepsTheRevisionFromCreate(t *testing.T) {
	var variableRevision string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.Path, "/variables") {
			variableRevision = r.URL.Query().Get("revision")
			_, _ = w.Write([]byte(`{"variables":[]}`))
			return
		}
		if strings.Contains(r.URL.Path, "/revisions/1") {
			_, _ = w.Write([]byte(`{"identity":"checkout","hubIdentity":"hub","revision":"1","name":"Checkout","toolType":"K6"}`))
			return
		}
		_, _ = w.Write([]byte(`{"identity":"checkout","hubIdentity":"hub","revision":"2","name":"Checkout v2","toolType":"K6"}`))
	}))
	defer server.Close()
	d := schema.TestResourceDataRaw(t, resourceChaosLoadTestTemplateSchema(), map[string]interface{}{
		"org_id": "default", "project_id": "demo", "hub_identity": "hub", "identifier": "checkout",
		"name": "Checkout", "revision": "1", "tool_type": "K6",
	})
	d.SetId("checkout")
	diags := resourceChaosLoadTestTemplateRead(t.Context(), d, loadTestSession(server))
	if diags.HasError() {
		t.Fatalf("diags = %v", diags)
	}
	if d.Get("revision") != "1" || d.Get("name") != "Checkout" {
		t.Fatalf("revision = %q name = %q", d.Get("revision"), d.Get("name"))
	}
	if variableRevision != "1" {
		t.Fatalf("variable revision = %q", variableRevision)
	}
}

func TestTemplateDataSourceReadsTheLatestRevision(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.Path, "/revisions/") {
			t.Errorf("data source read %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}
		if strings.Contains(r.URL.Path, "/variables") {
			if r.URL.Query().Get("revision") != "" {
				t.Errorf("data source variables revision = %q", r.URL.Query().Get("revision"))
			}
			_, _ = w.Write([]byte(`{"variables":[]}`))
			return
		}
		_, _ = w.Write([]byte(`{"identity":"checkout","hubIdentity":"hub","revision":"2","name":"Checkout v2","toolType":"K6"}`))
	}))
	defer server.Close()
	d := schema.TestResourceDataRaw(t, dataSourceChaosLoadTestTemplateSchema(), map[string]interface{}{
		"org_id": "default", "project_id": "demo", "hub_identity": "hub", "identifier": "checkout",
	})
	diags := dataSourceChaosLoadTestTemplateRead(t.Context(), d, loadTestSession(server))
	if diags.HasError() {
		t.Fatalf("diags = %v", diags)
	}
	if d.Get("revision") != "2" {
		t.Fatalf("revision = %q", d.Get("revision"))
	}
}

func TestTemplateUpdateLeavesANewerRevisionAlone(t *testing.T) {
	var updated bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPut {
			updated = true
			w.WriteHeader(http.StatusOK)
		}
		_, _ = w.Write([]byte(`{"identity":"checkout","hubIdentity":"hub","revision":"2","name":"Checkout v2","toolType":"K6"}`))
	}))
	defer server.Close()
	d := schema.TestResourceDataRaw(t, resourceChaosLoadTestTemplateSchema(), map[string]interface{}{
		"org_id": "default", "project_id": "demo", "hub_identity": "hub", "identifier": "checkout",
		"name": "Checkout", "revision": "1", "tool_type": "K6",
		"tool_config": `{"mode":"script","script":{"content":"a2s="}}`,
	})
	d.SetId("checkout")
	diags := resourceChaosLoadTestTemplateUpdate(t.Context(), d, loadTestSession(server))
	if !diags.HasError() || !strings.Contains(diags[0].Summary, "latest") {
		t.Fatalf("diags = %v", diags)
	}
	if updated {
		t.Fatal("update was sent to the latest revision")
	}
	if d.Id() != "checkout" {
		t.Fatalf("id = %q", d.Id())
	}
}

func TestTemplateRevisionRejectsLatest(t *testing.T) {
	d := schema.TestResourceDataRaw(t, resourceChaosLoadTestTemplateRevisionSchema(), map[string]interface{}{
		"org_id": "default", "project_id": "demo", "hub_identity": "hub", "template_identity": "checkout",
		"revision": "latest",
	})
	diags := resourceChaosLoadTestTemplateRevisionCreate(t.Context(), d, nil)
	if !diags.HasError() || !strings.Contains(diags[0].Summary, "latest") {
		t.Fatalf("diags = %v", diags)
	}
}

func TestParseTemplateImportIDAllowsAccountAndOrgHubs(t *testing.T) {
	org, project, hub, id, err := parseTemplateImportID("//hub/checkout")
	if err != nil || org != "" || project != "" || hub != "hub" || id != "checkout" {
		t.Fatalf("account hub = %q %q %q %q %v", org, project, hub, id, err)
	}
	org, project, hub, id, err = parseTemplateImportID("default//hub/checkout")
	if err != nil || org != "default" || project != "" || hub != "hub" || id != "checkout" {
		t.Fatalf("org hub = %q %q %q %q %v", org, project, hub, id, err)
	}
	if _, _, _, _, err = parseTemplateImportID("/demo/hub/checkout"); err == nil {
		t.Fatal("expected a project hub without an org to be rejected")
	}
	org, project, hub, template, revision, err := parseTemplateRevisionImportID("//hub/checkout/2")
	if err != nil || org != "" || project != "" || hub != "hub" || template != "checkout" || revision != "2" {
		t.Fatalf("revision account hub = %q %q %q %q %q %v", org, project, hub, template, revision, err)
	}
}

func TestRevisionDeleteRefusesPinnedRevision(t *testing.T) {
	var deleted bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.Contains(r.URL.Path, "/referenced-by"):
			_, _ = w.Write([]byte(`{"data":[{"identity":"pinned","kind":"LoadTest","orgIdentifier":"default","projectIdentifier":"demo"}],"pagination":{"totalPages":1}}`))
		case strings.Contains(r.URL.Path, "/v1/load-tests/"):
			_, _ = w.Write([]byte(`{"identity":"pinned","importType":"REFERENCE","templateReference":{"hubIdentity":"hub","identity":"checkout","revision":"2"}}`))
		case r.Method == http.MethodDelete:
			deleted = true
			w.WriteHeader(http.StatusNoContent)
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()
	client := sdkloadtest.NewAPIClient(&sdkloadtest.Configuration{BasePath: server.URL + "/loadTest/manager/api", AccountId: "acc"})
	d := schema.TestResourceDataRaw(t, resourceChaosLoadTestTemplateRevisionSchema(), map[string]interface{}{
		"org_id": "default", "project_id": "demo", "hub_identity": "hub", "template_identity": "checkout", "revision": "2",
	})
	d.SetId("2")
	diags := resourceChaosLoadTestTemplateRevisionDelete(t.Context(), d, &internal.Session{LoadTestClient: client})
	if !diags.HasError() || !strings.Contains(diags[0].Summary, "pinned") {
		t.Fatalf("diags = %v", diags)
	}
	if deleted || d.Id() == "" {
		t.Fatalf("deleted = %v id = %q", deleted, d.Id())
	}
}

func TestTemplateAndRevisionPlansStayQuiet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.Contains(r.URL.Path, "/referenced-by"):
			_, _ = w.Write([]byte(`{"data":[],"pagination":{"totalPages":0}}`))
		case strings.Contains(r.URL.Path, "/variables"):
			_, _ = w.Write([]byte(`{"variables":[]}`))
		case strings.Contains(r.URL.Path, "/revisions/2") || strings.HasSuffix(r.URL.Path, "/revisions"):
			_, _ = w.Write([]byte(`{"identity":"checkout","revision":"2","name":"Checkout","toolType":"K6","scriptContent":"a2s=","scriptSource":"inline","toolConfig":{"mode":"script","script":{"content":"a2s="}},"tags":["team"]}`))
		default:
			_, _ = w.Write([]byte(`{"identity":"checkout","hubIdentity":"hub","revision":"1","name":"Checkout","toolType":"K6","scriptContent":"a2s=","scriptSource":"inline","description":"filled","envId":"env","infraId":"infra","infraType":"Kubernetes","toolConfig":{"mode":"script","script":{"content":"a2s="}}}`))
		}
	}))
	defer server.Close()
	template := `
provider "harness" {}
resource "harness_chaos_load_test_template" "checkout" {
  org_id       = "default"
  project_id   = "demo"
  hub_identity = "hub"
  identifier   = "checkout"
  name         = "Checkout"
  revision     = "1"
  tool_type    = "K6"
  tool_config  = jsonencode({ mode = "script", script = { content = "a2s=" } })
}
`
	revision := template + `
resource "harness_chaos_load_test_template_revision" "v2" {
  org_id            = "default"
  project_id        = "demo"
  hub_identity      = "hub"
  template_identity = "checkout"
  revision          = "2"
}
`
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: templateProviderFactories(server),
		Steps: []resource.TestStep{
			{Config: template},
			{Config: template, PlanOnly: true, ExpectNonEmptyPlan: false},
			{Config: revision},
			{Config: revision, PlanOnly: true, ExpectNonEmptyPlan: false},
		},
	})
}

func TestNewerTemplateRevisionDoesNotReplaceTheTemplate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.Contains(r.URL.Path, "/variables"):
			_, _ = w.Write([]byte(`{"variables":[]}`))
		case strings.Contains(r.URL.Path, "/revisions/1"):
			_, _ = w.Write([]byte(`{"identity":"checkout","hubIdentity":"hub","revision":"1","name":"Checkout","toolType":"K6","scriptContent":"a2s=","scriptSource":"inline","toolConfig":{"mode":"script","script":{"content":"a2s="}}}`))
		default:
			_, _ = w.Write([]byte(`{"identity":"checkout","hubIdentity":"hub","revision":"2","name":"Checkout v2","toolType":"K6","scriptContent":"bmV3","scriptSource":"inline","toolConfig":{"mode":"script","script":{"content":"bmV3"}}}`))
		}
	}))
	defer server.Close()
	config := `
provider "harness" {}
resource "harness_chaos_load_test_template" "checkout" {
  org_id       = "default"
  project_id   = "demo"
  hub_identity = "hub"
  identifier   = "checkout"
  name         = "Checkout"
  revision     = "1"
  tool_type    = "K6"
  tool_config  = jsonencode({ mode = "script", script = { content = "a2s=" } })
}
`
	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: templateProviderFactories(server),
		Steps: []resource.TestStep{
			{Config: config},
			{Config: config, PlanOnly: true, ExpectNonEmptyPlan: false},
		},
	})
}

func templateProviderFactories(server *httptest.Server) map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"harness": func() (*schema.Provider, error) {
			client := sdkloadtest.NewAPIClient(&sdkloadtest.Configuration{
				BasePath:  server.URL + "/loadTest/manager/api",
				AccountId: "acc",
			})
			return &schema.Provider{
				ResourcesMap: map[string]*schema.Resource{
					"harness_chaos_load_test_template":          ResourceChaosLoadTestTemplate(),
					"harness_chaos_load_test_template_revision": ResourceChaosLoadTestTemplateRevision(),
				},
				ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
					return &internal.Session{LoadTestClient: client}, nil
				},
			}, nil
		},
	}
}
