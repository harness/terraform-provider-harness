package scan

import (
	"context"
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
	diags := resourceChaosScanCreate(t.Context(), d, scanSession(server))
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
	}, &chaos.ScanReport{ScannedRisks: []chaos.ScanRisk{{Identity: "risk-a", Name: "Kept"}}}, true); err != nil {
		t.Fatal(err)
	}

	diags := readScan(t.Context(), d, scanSession(server), "scan_1")
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
	}, &chaos.ScanReport{ScannedRisks: []chaos.ScanRisk{{Identity: "risk-a"}}}, true); err != nil {
		t.Fatal(err)
	}

	if diags := readScan(t.Context(), d, scanSession(server), "scan_1"); diags.HasError() {
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

	d := schema.TestResourceDataRaw(t, resourceChaosScanSchema(), map[string]interface{}{})
	d.SetId("org/proj/scan_1")
	imported, err := resourceChaosScanImport(t.Context(), d, scanSession(server))
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
	if _, err := resourceChaosScanImport(t.Context(), d, scanSession(server)); err == nil {
		t.Fatal("short import id was accepted")
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

func scanResourceData(t *testing.T) *schema.ResourceData {
	t.Helper()
	return schema.TestResourceDataRaw(t, resourceChaosScanSchema(), map[string]interface{}{
		"org_id":     "org",
		"project_id": "proj",
		"identity":   "scan_1",
		"name":       "Scan",
		"pipeline_scan": []interface{}{
			map[string]interface{}{
				"pipeline_identity":     "pipe",
				"pipeline_execution_id": "exec",
			},
		},
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
