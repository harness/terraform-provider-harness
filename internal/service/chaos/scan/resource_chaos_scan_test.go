package scan

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestChaosScan_ResourceSchemaValid(t *testing.T) {
	for _, resource := range []*schema.Resource{ResourceChaosPipelineScan(), ResourceChaosInfrastructureScan()} {
		if err := resource.InternalValidate(resource.Schema, true); err != nil {
			t.Fatalf("resource schema failed InternalValidate: %v", err)
		}
		if resource.Importer == nil {
			t.Fatal("expected resource to support import")
		}
	}
}

func TestChaosScan_DataSourceSchemaValid(t *testing.T) {
	for _, dataSource := range []*schema.Resource{DataSourceChaosPipelineScan(), DataSourceChaosInfrastructureScan()} {
		if err := dataSource.InternalValidate(dataSource.Schema, false); err != nil {
			t.Fatalf("data source schema failed InternalValidate: %v", err)
		}
	}
}

func TestScanCreateRequestPipeline(t *testing.T) {
	d := schema.TestResourceDataRaw(t, resourceChaosPipelineScanSchema(), map[string]interface{}{
		"identity":              "cart_scan",
		"name":                  "Cart scan",
		"org_id":                "default",
		"project_id":            "demo",
		"description":           "pipeline execution",
		"detection_method":      "ruleBased",
		"tags":                  []interface{}{"checkout"},
		"pipeline_identity":     "cart_deploy",
		"pipeline_execution_id": "exec-1",
	})

	req, err := pipelineScanRequest(d)
	if err != nil {
		t.Fatal(err)
	}
	if req.DetectionMethod != chaos.DetectionMethodRiskAnalyzer {
		t.Fatalf("method = %s", req.DetectionMethod)
	}
	if req.PipelineIdentity != "cart_deploy" || req.PipelineExecutionID != "exec-1" {
		t.Fatalf("pipeline = %+v", req)
	}
}

func TestInfrastructureScanFromChaosInfrastructure(t *testing.T) {
	d := schema.TestResourceDataRaw(t, resourceChaosInfrastructureScanSchema(), map[string]interface{}{
		"identity":          "infra_scan",
		"name":              "Infra scan",
		"description":       "from chaos infrastructure",
		"detection_method":  "ruleBased",
		"tags":              []interface{}{"payments"},
		"infrastructure_id": "payments-infra",
		"environment_id":    "prod",
	})

	req, ok, err := infrastructureScanRequest(d)
	if err != nil {
		t.Fatal(err)
	}
	if !ok || req == nil {
		t.Fatal("expected the chaos infrastructure form")
	}
	if req.InfrastructureID != "payments-infra" || req.EnvironmentID != "prod" || req.DetectionMethod != "RiskAnalyzer" || req.Name != "Infra scan" {
		t.Fatalf("request = %+v", req)
	}
	if len(req.Tags) != 1 || req.Tags[0] != "payments" {
		t.Fatalf("tags = %#v", req.Tags)
	}
}

func TestScanCreateRequestInfrastructure(t *testing.T) {
	d := schema.TestResourceDataRaw(t, resourceChaosInfrastructureScanSchema(), map[string]interface{}{
		"identity":                     "infra_scan",
		"name":                         "Infra scan",
		"detection_method":             "ai",
		"agent_identity":               "agent-1",
		"environment_identity":         "env-1",
		"environment_parent_unique_id": "parent-1",
	})

	req, err := discoveryAgentScanRequest(d)
	if err != nil {
		t.Fatal(err)
	}
	if req.ScanType != chaos.ScanTypeDiscoveryAgent {
		t.Fatalf("scan type = %s", req.ScanType)
	}
	if req.DetectionMethod != "AI" {
		t.Fatalf("detection method = %s", req.DetectionMethod)
	}
	infra := req.Source.Infrastructure
	if infra == nil || infra.AgentIdentity != "agent-1" || infra.EnvironmentIdentity != "env-1" || infra.EnvironmentParentUniqueID != "parent-1" {
		t.Fatalf("infrastructure = %+v", infra)
	}
}

func TestScanUpdateRequestKeepsSourceOffTheBody(t *testing.T) {
	d := schema.TestResourceDataRaw(t, resourceChaosPipelineScanSchema(), map[string]interface{}{
		"identity":              "scan_1",
		"name":                  "Renamed",
		"pipeline_identity":     "pipe",
		"pipeline_execution_id": "exec",
	})

	req := scanUpdateRequest(d)
	if req.Name == nil || *req.Name != "Renamed" {
		t.Fatalf("name = %+v", req.Name)
	}
	if req.Tags == nil {
		t.Fatal("tags should be an empty list so the API can clear them")
	}
}

func TestSetScanStateFlattensRisks(t *testing.T) {
	d := schema.TestResourceDataRaw(t, resourceChaosPipelineScanSchema(), map[string]interface{}{
		"identity":              "scan_1",
		"name":                  "Scan",
		"pipeline_identity":     "pipe",
		"pipeline_execution_id": "exec",
	})

	err := setScanState(d, &chaos.Scan{
		Identity:        "scan_1",
		Name:            "Scan",
		ScanType:        chaos.ScanTypePipelineExecution,
		Status:          chaos.ScanStatusCompleted,
		DetectionMethod: chaos.DetectionMethodAI,
		Source: chaos.ScanSource{Pipeline: &chaos.PipelineScan{
			PipelineIdentity:    "pipe",
			PipelineExecutionID: "exec",
			Infrastructures: []chaos.PipelineInfrastructure{{
				EnvironmentRef:         "prod",
				InfrastructureIdentity: "k8s",
				ChaosOnboarded:         false,
				ChaosSupported:         true,
			}},
		}},
		CreatedAt: 1710000000,
		Summary:   &chaos.HeatmapSummary{TotalRisks: 1, High: 1},
		RiskSummary: &chaos.RiskSummary{
			RiskCount:      1,
			SeverityCounts: map[string]int{"High": 1},
			CategoryCounts: map[string]int{"Availability": 1},
		},
		CollectedData: &chaos.CollectionSummary{
			TotalResources: 3,
			Namespaces:     []string{"shop"},
			Errors:         []chaos.CollectionError{{Source: "pipeline", Message: "no kubernetes cd stage"}},
		},
	}, &chaos.ScanReport{
		ScannedRisks: []chaos.ScanRisk{{
			Identity:        "risk-b",
			Name:            "Later risk",
			Severity:        "Low",
			Status:          "Open",
			ServiceIdentity: "pay",
		}, {
			Identity:        "risk-a",
			Name:            "Single replica",
			Severity:        "High",
			Status:          "Open",
			ServiceIdentity: "cart",
			ServiceName:     "cart",
			ServiceType:     "DiscoveredService",
			Metadata:        map[string]interface{}{"rule": "replicas"},
			Occurrences: []chaos.ScanRiskOccurrence{{
				OccurrenceID: "occ-1",
				DetectedAt:   20,
				SourceID:     "scan_1",
				Context:      chaos.OccurrenceContext{Evidence: []string{"replicas=1"}},
			}},
		}},
	}, true, true)
	if err != nil {
		t.Fatal(err)
	}

	risks := d.Get("risks").([]interface{})
	if len(risks) != 2 {
		t.Fatalf("risks = %#v", risks)
	}
	risk := risks[0].(map[string]interface{})
	if risk["identity"] != "risk-a" || risk["severity"] != "High" || risk["metadata_json"] != `{"rule":"replicas"}` {
		t.Fatalf("risk = %#v", risk)
	}
	infras := d.Get("infrastructures").([]interface{})
	if len(infras) != 1 || infras[0].(map[string]interface{})["environment_ref"] != "prod" {
		t.Fatalf("infrastructures = %#v", infras)
	}
	if infras[0].(map[string]interface{})["chaos_onboarded"] != false {
		t.Fatalf("chaos_onboarded = %#v", infras[0])
	}
	if fmt.Sprint(d.Get("created_at")) != "1710000000" {
		t.Fatalf("created_at = %#v", d.Get("created_at"))
	}
	summary := d.Get("risk_summary").([]interface{})
	if len(summary) != 1 {
		t.Fatalf("risk_summary = %#v", summary)
	}
	counts := summary[0].(map[string]interface{})["severity_counts"].(map[string]interface{})
	if fmt.Sprint(counts["High"]) != "1" {
		t.Fatalf("severity_counts = %#v", counts)
	}
	if d.Get("status") != chaos.ScanStatusCompleted {
		t.Fatalf("status = %s", d.Get("status"))
	}

	infraData := schema.TestResourceDataRaw(t, resourceChaosInfrastructureScanSchema(), map[string]interface{}{
		"identity":          "infra_scan",
		"name":              "Infra",
		"infrastructure_id": "payments-infra",
		"environment_id":    "prod",
	})
	if err := setScanState(infraData, &chaos.Scan{
		Identity: "infra_scan",
		Name:     "Infra",
		ScanType: chaos.ScanTypeDiscoveryAgent,
		Status:   chaos.ScanStatusCompleted,
		Source: chaos.ScanSource{Infrastructure: &chaos.InfrastructureScan{
			AgentIdentity:             "agent-1",
			EnvironmentIdentity:       "env-agent",
			EnvironmentParentUniqueID: "parent-1",
		}},
	}, nil, true, true); err != nil {
		t.Fatal(err)
	}
	if infraData.Get("infrastructure_id") != "payments-infra" || infraData.Get("environment_id") != "prod" {
		t.Fatalf("chaos infrastructure was dropped: %#v", infraData.State())
	}
	if infraData.Get("agent_identity") != "agent-1" || infraData.Get("environment_identity") != "env-agent" || infraData.Get("environment_parent_unique_id") != "parent-1" {
		t.Fatalf("resolved agent = %s %s %s", infraData.Get("agent_identity"), infraData.Get("environment_identity"), infraData.Get("environment_parent_unique_id"))
	}
	collected := d.Get("collected_data").([]interface{})
	if len(collected) != 1 {
		t.Fatalf("collected_data = %#v", collected)
	}
	collectionErrors := collected[0].(map[string]interface{})["errors"].([]interface{})
	if len(collectionErrors) != 1 || collectionErrors[0].(map[string]interface{})["message"] != "no kubernetes cd stage" {
		t.Fatalf("collection errors = %#v", collectionErrors)
	}
	if risks[0].(map[string]interface{})["identity"] != "risk-a" || risks[1].(map[string]interface{})["identity"] != "risk-b" {
		t.Fatalf("risks were not sorted by identity: %#v", risks)
	}
}

func TestChaosScanForceNewIsOnInputsOnly(t *testing.T) {
	pipeline := resourceChaosPipelineScanSchema()
	infra := resourceChaosInfrastructureScanSchema()
	if !pipeline["org_id"].Required || !pipeline["project_id"].Required || !infra["org_id"].Required {
		t.Fatal("org_id and project_id must be required")
	}
	if _, ok := pipeline["infrastructure_id"]; ok {
		t.Fatal("a pipeline scan must not carry infrastructure scan fields")
	}
	if _, ok := infra["pipeline_identity"]; ok {
		t.Fatal("an infrastructure scan must not carry pipeline scan fields")
	}
	if !pipeline["pipeline_identity"].ForceNew || !pipeline["pipeline_execution_id"].ForceNew {
		t.Fatal("pipeline inputs must be ForceNew")
	}
	if pipeline["notify_id"].ForceNew || pipeline["infrastructures"].ForceNew {
		t.Fatal("API-filled pipeline fields must not be ForceNew")
	}
	if !infra["infrastructure_id"].ForceNew || !infra["environment_id"].ForceNew {
		t.Fatal("chaos infrastructure inputs must be ForceNew")
	}
	if len(infra["infrastructure_id"].RequiredWith) != 1 || infra["infrastructure_id"].RequiredWith[0] != "environment_id" {
		t.Fatal("infrastructure_id requires environment_id")
	}
	if len(infra["environment_id"].RequiredWith) != 1 || infra["environment_id"].RequiredWith[0] != "infrastructure_id" {
		t.Fatal("environment_id is required with infrastructure_id")
	}
	if strings.Contains(pipeline["identity"].Description, "infrastructure") || strings.Contains(pipeline["retry_token"].Description, "infrastructure") {
		t.Fatal("pipeline scan docs must not describe the infrastructure scan retry")
	}
	if infra["agent_identity"].ForceNew || infra["environment_identity"].ForceNew || infra["environment_parent_unique_id"].ForceNew {
		t.Fatal("agent fields are filled by the API; CustomizeDiff replaces the scan only when a configured value changes")
	}
	if !infra["agent_identity"].Optional || !infra["agent_identity"].Computed {
		t.Fatal("agent_identity must stay optional so existing configs and imports keep working")
	}
	if !infra["infrastructure_id"].Optional || len(infra["infrastructure_id"].ConflictsWith) == 0 || len(infra["infrastructure_id"].AtLeastOneOf) != 2 {
		t.Fatal("infrastructure_id must be an alternative to agent_identity")
	}
	if pipeline["timeouts"] != nil {
		t.Fatal("timeouts are declared on the resource, not the schema map")
	}
	if ResourceChaosPipelineScan().Timeouts == nil || ResourceChaosInfrastructureScan().Timeouts.Create == nil {
		t.Fatal("create timeout is required so apply waits for a terminal scan")
	}
}

func TestWaitForTerminalScan(t *testing.T) {
	previous := scanPollInterval
	scanPollInterval = time.Millisecond
	t.Cleanup(func() { scanPollInterval = previous })

	calls := 0
	err := waitForTerminalScan(context.Background(), false, func(context.Context) (*chaos.Scan, *http.Response, error) {
		calls++
		if calls == 1 {
			return &chaos.Scan{Identity: "scan_1", Status: chaos.ScanStatusQueued}, nil, nil
		}
		return &chaos.Scan{Identity: "scan_1", Status: chaos.ScanStatusCompleted}, nil, nil
	})
	if err.HasError() {
		t.Fatal(err)
	}
	if calls != 2 {
		t.Fatalf("calls = %d", calls)
	}

	failed := waitForTerminalScan(context.Background(), false, func(context.Context) (*chaos.Scan, *http.Response, error) {
		return &chaos.Scan{Identity: "scan_1", Status: chaos.ScanStatusErrored, ErrorMessage: "no kubernetes cd stage"}, nil, nil
	})
	if !failed.HasError() || !strings.Contains(failed[0].Summary, "no kubernetes cd stage") {
		t.Fatalf("failed wait = %#v", failed)
	}

	aborted := waitForTerminalScan(context.Background(), true, func(context.Context) (*chaos.Scan, *http.Response, error) {
		return &chaos.Scan{Identity: "scan_1", Status: chaos.ScanStatusAborted}, nil, nil
	})
	if aborted.HasError() {
		t.Fatalf("requested abort = %#v", aborted)
	}
	unexpected := waitForTerminalScan(context.Background(), false, func(context.Context) (*chaos.Scan, *http.Response, error) {
		return &chaos.Scan{Identity: "scan_1", Status: chaos.ScanStatusAborted, ErrorMessage: "stopped"}, nil, nil
	})
	if !unexpected.HasError() {
		t.Fatal("an abort that was not requested must fail the apply")
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	timedOut := waitForTerminalScan(ctx, false, func(context.Context) (*chaos.Scan, *http.Response, error) {
		return &chaos.Scan{Identity: "scan_1", Status: chaos.ScanStatusQueued}, nil, nil
	})
	if !timedOut.HasError() || !strings.Contains(timedOut[0].Summary, "timed out") {
		t.Fatalf("timeout = %#v", timedOut)
	}
}

func TestParseScanImportID(t *testing.T) {
	orgID, projectID, identity, err := parseScanImportID("org/proj/scan_1")
	if err != nil {
		t.Fatal(err)
	}
	if orgID != "org" || projectID != "proj" || identity != "scan_1" {
		t.Fatalf("got %s %s %s", orgID, projectID, identity)
	}

	for _, id := range []string{"scan_1", "org/scan_1", "a/b/c/d", "org//scan_1", "/proj/scan_1"} {
		if _, _, _, err = parseScanImportID(id); err == nil {
			t.Fatalf("accepted invalid import id %q", id)
		}
	}
}
