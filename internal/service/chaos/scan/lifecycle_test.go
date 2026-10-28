package scan

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceChaosScanCreateFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"pipeline identity is required"}`))
	}))
	defer server.Close()

	d := scanResourceData(t)
	diags := resourcePipelineScanCreate(t.Context(), d, scanSession(server))
	if !diags.HasError() || !strings.Contains(diags[0].Summary, "pipeline identity is required") {
		t.Fatalf("diags = %#v", diags)
	}
	if d.Id() != "" {
		t.Fatalf("failed create stored id %q", d.Id())
	}
}

func TestResourceChaosScanReportFailureKeepsRisks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/report") {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"message":"report store unavailable"}`))
			return
		}
		_, _ = w.Write([]byte(`{"identity":"scan_1","name":"Scan","orgID":"org","projectID":"proj","scanType":"PipelineExecution","status":"COMPLETED","source":{"pipeline":{"pipelineIdentity":"pipe","pipelineExecutionID":"exec"}}}`))
	}))
	defer server.Close()

	d := scanResourceData(t)
	d.SetId("scan_1")
	if err := setScanState(d, &chaos.Scan{
		Identity: "scan_1",
		Name:     "Scan",
		ScanType: chaos.ScanTypePipelineExecution,
		Status:   chaos.ScanStatusCompleted,
		Source:   chaos.ScanSource{Pipeline: &chaos.PipelineScan{PipelineIdentity: "pipe", PipelineExecutionID: "exec"}},
	}, &chaos.ScanReport{ScannedRisks: []chaos.ScanRisk{{Identity: "risk-a", Name: "Kept"}}}, true, true); err != nil {
		t.Fatal(err)
	}

	diags := readScan(t.Context(), d, scanSession(server), "scan_1", chaos.ScanTypePipelineExecution, true)
	if diags.HasError() {
		t.Fatalf("report failure removed the scan: %#v", diags)
	}
	if len(diags) != 1 || diags[0].Severity != diag.Warning {
		t.Fatalf("diags = %#v", diags)
	}
	risks := d.Get("risks").([]interface{})
	if len(risks) != 1 || risks[0].(map[string]interface{})["identity"] != "risk-a" {
		t.Fatalf("risks = %#v", risks)
	}
}

func TestResourceChaosScanEmptyReportClearsRisks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/report") {
			_, _ = w.Write([]byte(`{"scannedRisks":[],"summary":{"totalRisks":0}}`))
			return
		}
		_, _ = w.Write([]byte(`{"identity":"scan_1","name":"Scan","orgID":"org","projectID":"proj","scanType":"PipelineExecution","status":"COMPLETED","source":{"pipeline":{"pipelineIdentity":"pipe","pipelineExecutionID":"exec"}}}`))
	}))
	defer server.Close()

	d := scanResourceData(t)
	d.SetId("scan_1")
	if err := setScanState(d, &chaos.Scan{
		Identity: "scan_1",
		ScanType: chaos.ScanTypePipelineExecution,
		Status:   chaos.ScanStatusCompleted,
		Source:   chaos.ScanSource{Pipeline: &chaos.PipelineScan{PipelineIdentity: "pipe", PipelineExecutionID: "exec"}},
	}, &chaos.ScanReport{ScannedRisks: []chaos.ScanRisk{{Identity: "risk-a"}}}, true, true); err != nil {
		t.Fatal(err)
	}

	if diags := readScan(t.Context(), d, scanSession(server), "scan_1", chaos.ScanTypePipelineExecution, true); diags.HasError() {
		t.Fatal(diags)
	}
	if risks := d.Get("risks").([]interface{}); len(risks) != 0 {
		t.Fatalf("risks = %#v", risks)
	}
}

func TestResourceChaosScanDeleteNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message":"risk scan not found"}`))
	}))
	defer server.Close()

	d := scanResourceData(t)
	d.SetId("scan_1")
	diags := resourceChaosScanDelete(t.Context(), d, scanSession(server))
	if diags.HasError() {
		t.Fatalf("diags = %#v", diags)
	}
	if d.Id() != "" {
		t.Fatalf("id = %q", d.Id())
	}
}

func TestResourceChaosScanImportScope(t *testing.T) {
	var gotQuery string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/report") {
			_, _ = w.Write([]byte(`{"scannedRisks":[],"summary":{}}`))
			return
		}
		_, _ = w.Write([]byte(`{"identity":"scan_1","name":"Scan","orgID":"org","projectID":"proj","scanType":"DiscoveryAgent","status":"COMPLETED","source":{"discoveryAgent":{"agentIdentity":"agent","environmentIdentity":"env"}}}`))
	}))
	defer server.Close()

	d := schema.TestResourceDataRaw(t, resourceChaosInfrastructureScanSchema(), map[string]interface{}{})
	d.SetId("org/proj/scan_1")
	imported, err := importChaosScan(t.Context(), d, scanSession(server), chaos.ScanTypeDiscoveryAgent)
	if err != nil {
		t.Fatal(err)
	}
	if len(imported) != 1 || imported[0].Id() != "scan_1" {
		t.Fatalf("imported = %#v", imported)
	}
	if imported[0].Get("org_id") != "org" || imported[0].Get("project_id") != "proj" {
		t.Fatalf("scope org=%v project=%v", imported[0].Get("org_id"), imported[0].Get("project_id"))
	}
	if !strings.Contains(gotQuery, "organizationIdentifier=org") || !strings.Contains(gotQuery, "projectIdentifier=proj") {
		t.Fatalf("query = %s", gotQuery)
	}

	d.SetId("scan_1")
	if _, err := importChaosScan(t.Context(), d, scanSession(server), chaos.ScanTypeDiscoveryAgent); err == nil {
		t.Fatal("short import id was accepted")
	}
}

func TestResourceChaosScanImportRejectsOtherScanType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"identity":"scan_1","name":"Scan","orgID":"org","projectID":"proj","scanType":"PipelineExecution","status":"COMPLETED","source":{"pipelineExecution":{"pipelineIdentity":"pipe","pipelineExecutionID":"exec"}}}`))
	}))
	defer server.Close()

	d := schema.TestResourceDataRaw(t, resourceChaosInfrastructureScanSchema(), map[string]interface{}{})
	d.SetId("org/proj/scan_1")
	_, err := importChaosScan(t.Context(), d, scanSession(server), chaos.ScanTypeDiscoveryAgent)
	if err == nil || !strings.Contains(err.Error(), "PipelineExecution") {
		t.Fatalf("error = %v", err)
	}
}

func TestAbortAndRetryHTTP(t *testing.T) {
	previous := scanPollInterval
	scanPollInterval = time.Millisecond
	t.Cleanup(func() { scanPollInterval = previous })

	var paths []string
	afterRetry := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.Method+" "+r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/abort"):
			_, _ = w.Write([]byte(`{"status":"ok"}`))
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/retry"):
			afterRetry = true
			_, _ = w.Write([]byte(`{"status":"ok"}`))
		case afterRetry:
			_, _ = w.Write([]byte(`{"identity":"scan_1","status":"ERRORED","errorMessage":"still broken"}`))
		default:
			_, _ = w.Write([]byte(`{"identity":"scan_1","status":"ABORTED"}`))
		}
	}))
	defer server.Close()

	client := chaos.NewScanClient(scanSession(server).ChaosClient)
	scope := chaos.Scope{AccountID: "acc", OrgID: "org", ProjectID: "proj"}
	if _, err := client.Abort(t.Context(), scope, "scan_1"); err != nil {
		t.Fatal(err)
	}
	aborted := waitForTerminalScan(t.Context(), true, func(ctx context.Context) (*chaos.Scan, *http.Response, error) {
		return client.Get(ctx, scope, "scan_1")
	})
	if aborted.HasError() {
		t.Fatalf("abort wait = %#v", aborted)
	}

	if _, err := client.Retry(t.Context(), scope, "scan_1"); err != nil {
		t.Fatal(err)
	}
	retried := waitForTerminalScan(t.Context(), false, func(ctx context.Context) (*chaos.Scan, *http.Response, error) {
		return client.Get(ctx, scope, "scan_1")
	})
	if !retried.HasError() || !strings.Contains(retried[0].Summary, "still broken") {
		t.Fatalf("retry wait = %#v", retried)
	}
	joined := strings.Join(paths, ",")
	if !strings.Contains(joined, "POST /v3/risk-scans/scan_1/abort") || !strings.Contains(joined, "POST /v3/risk-scans/scan_1/retry") {
		t.Fatalf("paths = %s", joined)
	}
}

func TestDataSourcePipelineScanRejectsInfrastructureScan(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"identity":"scan_1","name":"Scan","orgID":"org","projectID":"proj","scanType":"DiscoveryAgent","status":"COMPLETED","source":{"discoveryAgent":{"agentIdentity":"agent","environmentIdentity":"env"}}}`))
	}))
	defer server.Close()

	d := scanResourceData(t)
	d.SetId("scan_1")
	diags := readScan(t.Context(), d, scanSession(server), "scan_1", chaos.ScanTypePipelineExecution, true)
	if !diags.HasError() || !strings.Contains(diags[0].Summary, "DiscoveryAgent") {
		t.Fatalf("diags = %#v", diags)
	}
}

func TestDataSourcesReadMatchingScanType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/report") {
			_, _ = w.Write([]byte(`{"scannedRisks":[],"summary":{"totalRisks":0}}`))
			return
		}
		if strings.Contains(r.URL.Path, "pipe_scan") {
			_, _ = w.Write([]byte(`{"identity":"pipe_scan","name":"Pipe","orgID":"org","projectID":"proj","scanType":"PipelineExecution","status":"COMPLETED","source":{"pipeline":{"pipelineIdentity":"pipe","pipelineExecutionID":"exec"}}}`))
			return
		}
		_, _ = w.Write([]byte(`{"identity":"infra_scan","name":"Infra","orgID":"org","projectID":"proj","scanType":"DiscoveryAgent","status":"COMPLETED","source":{"discoveryAgent":{"agentIdentity":"agent-1","environmentIdentity":"env-1","environmentParentUniqueId":"parent-1"}}}`))
	}))
	defer server.Close()

	pipeline := schema.TestResourceDataRaw(t, dataSourceChaosPipelineScanSchema(), map[string]interface{}{
		"org_id":     "org",
		"project_id": "proj",
		"identity":   "pipe_scan",
	})
	if diags := dataSourcePipelineScanRead(t.Context(), pipeline, scanSession(server)); diags.HasError() {
		t.Fatal(diags)
	}
	if pipeline.Get("pipeline_identity") != "pipe" || pipeline.Get("status") != chaos.ScanStatusCompleted {
		t.Fatalf("pipeline scan = %#v", pipeline.State())
	}

	infra := schema.TestResourceDataRaw(t, dataSourceChaosInfrastructureScanSchema(), map[string]interface{}{
		"org_id":     "org",
		"project_id": "proj",
		"identity":   "infra_scan",
	})
	if diags := dataSourceInfrastructureScanRead(t.Context(), infra, scanSession(server)); diags.HasError() {
		t.Fatal(diags)
	}
	if infra.Get("agent_identity") != "agent-1" || infra.Get("environment_identity") != "env-1" || infra.Get("environment_parent_unique_id") != "parent-1" {
		t.Fatalf("infrastructure scan = %#v", infra.State())
	}
}

func TestPipelineScanCreateSendsExecutionAndWaits(t *testing.T) {
	var got chaos.ScanCreateRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/v3/risk-scans":
			if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
				t.Fatal(err)
			}
			_, _ = w.Write([]byte(`{"identity":"scan_1","scanType":"PipelineExecution","status":"QUEUED"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/v3/risk-scans/scan_1":
			_, _ = w.Write([]byte(`{"identity":"scan_1","name":"Scan","orgID":"org","projectID":"proj","scanType":"PipelineExecution","status":"COMPLETED","detectionMethod":"RiskAnalyzer","source":{"pipeline":{"pipelineIdentity":"pipe","pipelineExecutionID":"exec"}}}`))
		case r.Method == http.MethodGet && r.URL.Path == "/v3/risk-scans/scan_1/report":
			_, _ = w.Write([]byte(`{"scannedRisks":[{"identity":"risk-b","name":"B"},{"identity":"risk-a","name":"A"}],"summary":{"totalRisks":2}}`))
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	d := scanResourceData(t)
	if err := d.Set("detection_method", "ruleBased"); err != nil {
		t.Fatal(err)
	}
	diags := resourcePipelineScanCreate(t.Context(), d, scanSession(server))
	if diags.HasError() {
		t.Fatal(diags)
	}
	if got.ScanType != chaos.ScanTypePipelineExecution || got.DetectionMethod != chaos.DetectionMethodRiskAnalyzer {
		t.Fatalf("create = %+v", got)
	}
	if got.Source.Pipeline == nil || got.Source.Pipeline.PipelineIdentity != "pipe" || got.Source.Pipeline.PipelineExecutionID != "exec" {
		t.Fatalf("pipeline source = %+v", got.Source.Pipeline)
	}
	if d.Id() != "scan_1" || d.Get("status") != chaos.ScanStatusCompleted {
		t.Fatalf("state id=%s status=%s", d.Id(), d.Get("status"))
	}
	risks := d.Get("risks").([]interface{})
	if len(risks) != 2 || risks[0].(map[string]interface{})["identity"] != "risk-a" {
		t.Fatalf("risks = %#v", risks)
	}
}

func TestInfrastructureScanCreateResolvesDiscoveryAgent(t *testing.T) {
	var got chaos.ScanCreateRequest
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/chaos/manager/api/rest/v2/infrastructure/payments-infra":
			if r.URL.Query().Get("environmentIdentifier") != "prod" {
				t.Fatalf("environment = %s", r.URL.RawQuery)
			}
			_, _ = w.Write([]byte(`{"discoveryAgentID":"agent-1"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/servicediscovery/api/v1/agents/agent-1":
			if r.URL.Query().Get("environmentIdentifier") != "prod" {
				t.Fatalf("agent query = %s", r.URL.RawQuery)
			}
			_, _ = w.Write([]byte(`{"identity":"agent-1","environmentIdentifier":"prod","parentUniqueId":"parent-1"}`))
		case r.Method == http.MethodPost && r.URL.Path == "/chaos/manager/api/v3/risk-scans":
			if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
				t.Fatal(err)
			}
			_, _ = w.Write([]byte(`{"identity":"infra_scan","scanType":"DiscoveryAgent","status":"COMPLETED"}`))
		case r.Method == http.MethodGet && r.URL.Path == "/chaos/manager/api/v3/risk-scans/infra_scan":
			_, _ = w.Write([]byte(`{"identity":"infra_scan","name":"Infra","orgID":"org","projectID":"proj","scanType":"DiscoveryAgent","status":"COMPLETED","source":{"discoveryAgent":{"agentIdentity":"agent-1","environmentIdentity":"prod","environmentParentUniqueId":"parent-1"}}}`))
		case r.Method == http.MethodGet && r.URL.Path == "/chaos/manager/api/v3/risk-scans/infra_scan/report":
			_, _ = w.Write([]byte(`{"scannedRisks":[],"summary":{"totalRisks":0}}`))
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	d := schema.TestResourceDataRaw(t, resourceChaosInfrastructureScanSchema(), map[string]interface{}{
		"org_id":            "org",
		"project_id":        "proj",
		"identity":          "infra_scan",
		"name":              "Infra",
		"detection_method":  "ai",
		"infrastructure_id": "payments-infra",
		"environment_id":    "prod",
	})
	client := chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   server.URL + "/chaos/manager/api",
		AccountId:  "acc",
		HTTPClient: server.Client(),
	})
	diags := resourceInfrastructureScanCreate(t.Context(), d, &internal.Session{ChaosClient: client})
	if diags.HasError() {
		t.Fatal(diags)
	}
	agent := got.Source.Infrastructure
	if got.ScanType != chaos.ScanTypeDiscoveryAgent || got.DetectionMethod != chaos.DetectionMethodAI || agent == nil {
		t.Fatalf("create = %+v", got)
	}
	if agent.AgentIdentity != "agent-1" || agent.EnvironmentIdentity != "prod" || agent.EnvironmentParentUniqueID != "parent-1" {
		t.Fatalf("agent = %+v", agent)
	}
	if d.Get("infrastructure_id") != "payments-infra" || d.Get("environment_id") != "prod" || d.Get("agent_identity") != "agent-1" {
		t.Fatalf("state = %#v", d.State())
	}
}

func TestInfrastructureScanCreateFailsWhenInfrastructureHasNoAgent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/chaos/manager/api/rest/v2/infrastructure/payments-infra":
			_, _ = w.Write([]byte(`{"identity":"payments-infra"}`))
		case "/servicediscovery/api/v1/agents/payments-infra":
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message":"agent not found"}`))
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	d := schema.TestResourceDataRaw(t, resourceChaosInfrastructureScanSchema(), map[string]interface{}{
		"org_id":            "org",
		"project_id":        "proj",
		"identity":          "infra_scan",
		"name":              "Infra",
		"infrastructure_id": "payments-infra",
		"environment_id":    "prod",
	})
	client := chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   server.URL + "/chaos/manager/api",
		AccountId:  "acc",
		HTTPClient: server.Client(),
	})
	diags := resourceInfrastructureScanCreate(t.Context(), d, &internal.Session{ChaosClient: client})
	if !diags.HasError() || !strings.Contains(diags[0].Summary, "no discovery agent") {
		t.Fatalf("diags = %#v", diags)
	}
	if d.Id() != "" {
		t.Fatalf("failed create stored id %q", d.Id())
	}
}

func scanResourceData(t *testing.T) *schema.ResourceData {
	t.Helper()
	return schema.TestResourceDataRaw(t, resourceChaosPipelineScanSchema(), map[string]interface{}{
		"org_id":                "org",
		"project_id":            "proj",
		"identity":              "scan_1",
		"name":                  "Scan",
		"pipeline_identity":     "pipe",
		"pipeline_execution_id": "exec",
	})
}

func scanSession(server *httptest.Server) *internal.Session {
	client := chaos.NewAPIClient(&chaos.Configuration{
		BasePath:   server.URL,
		AccountId:  "acc",
		HTTPClient: server.Client(),
	})
	return &internal.Session{ChaosClient: client}
}
