package scan

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func scopeFrom(d *schema.ResourceData, accountID string) chaos.Scope {
	return chaos.Scope{
		AccountID: accountID,
		OrgID:     d.Get("org_id").(string),
		ProjectID: d.Get("project_id").(string),
	}
}

func pipelineScanRequest(d *schema.ResourceData) (chaos.PipelineScanRequest, error) {
	method, err := chaos.APIDetectionMethod(d.Get("detection_method").(string))
	if err != nil {
		return chaos.PipelineScanRequest{}, err
	}
	return chaos.PipelineScanRequest{
		Identity:            d.Get("identity").(string),
		Name:                d.Get("name").(string),
		Description:         d.Get("description").(string),
		Tags:                expandStringList(d.Get("tags")),
		DetectionMethod:     method,
		PipelineIdentity:    d.Get("pipeline_identity").(string),
		PipelineExecutionID: d.Get("pipeline_execution_id").(string),
		NotifyID:            d.Get("notify_id").(string),
	}, nil
}

// infrastructureScanRequest is set when the user points at an existing chaos
// infrastructure. The scan client resolves the discovery agent from it.
func infrastructureScanRequest(d *schema.ResourceData) (*chaos.InfrastructureScanRequest, bool, error) {
	infrastructureID := d.Get("infrastructure_id").(string)
	if infrastructureID == "" {
		return nil, false, nil
	}
	if d.Get("environment_id").(string) == "" {
		return nil, false, fmt.Errorf("environment_id is required with infrastructure_id")
	}
	if d.Get("agent_identity").(string) != "" || d.Get("environment_identity").(string) != "" {
		return nil, false, fmt.Errorf("set infrastructure_id, or agent_identity and environment_identity, not both")
	}
	method, err := chaos.APIDetectionMethod(d.Get("detection_method").(string))
	if err != nil {
		return nil, false, err
	}
	return &chaos.InfrastructureScanRequest{
		Identity:         d.Get("identity").(string),
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		Tags:             expandStringList(d.Get("tags")),
		DetectionMethod:  method,
		InfrastructureID: infrastructureID,
		EnvironmentID:    d.Get("environment_id").(string),
	}, true, nil
}

func discoveryAgentScanRequest(d *schema.ResourceData) (chaos.ScanCreateRequest, error) {
	method, err := chaos.APIDetectionMethod(d.Get("detection_method").(string))
	if err != nil {
		return chaos.ScanCreateRequest{}, err
	}
	agentID := d.Get("agent_identity").(string)
	environmentID := d.Get("environment_identity").(string)
	if agentID == "" || environmentID == "" {
		return chaos.ScanCreateRequest{}, fmt.Errorf("set infrastructure_id, or agent_identity and environment_identity")
	}
	return chaos.ScanCreateRequest{
		Identity:        d.Get("identity").(string),
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		Tags:            expandStringList(d.Get("tags")),
		DetectionMethod: method,
		ScanType:        chaos.ScanTypeDiscoveryAgent,
		Source: chaos.ScanSource{Infrastructure: &chaos.InfrastructureScan{
			AgentIdentity:             agentID,
			EnvironmentIdentity:       environmentID,
			EnvironmentParentUniqueID: d.Get("environment_parent_unique_id").(string),
		}},
	}, nil
}

func scanUpdateRequest(d *schema.ResourceData) chaos.ScanUpdateRequest {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	tags := expandStringList(d.Get("tags"))
	if tags == nil {
		tags = []string{}
	}
	return chaos.ScanUpdateRequest{
		Name:        &name,
		Description: &description,
		Tags:        tags,
	}
}

func setScanState(d *schema.ResourceData, scan *chaos.Scan, report *chaos.ScanReport, includeRisks, keepInfrastructureIDs bool) error {
	if scan == nil {
		return fmt.Errorf("scan response is empty")
	}
	tags := scan.Tags
	if tags == nil {
		tags = []string{}
	}
	values := map[string]interface{}{
		"identity":         scan.Identity,
		"name":             scan.Name,
		"description":      scan.Description,
		"tags":             tags,
		"detection_method": scan.DetectionMethod,
		"scan_type":        scan.ScanType,
		"status":           scan.Status,
		"result":           scan.Result,
		"error_message":    scan.ErrorMessage,
		"retry_count":      scan.RetryCount,
		"account_id":       scan.AccountID,
		"created_at":       scan.CreatedAt,
		"updated_at":       scan.UpdatedAt,
		"completed_at":     scan.CompletedAt,
		"created_by":       scan.CreatedBy.DisplayName(),
		"updated_by":       scan.UpdatedBy.DisplayName(),
	}
	if scan.OrgID != "" {
		values["org_id"] = scan.OrgID
	}
	if scan.ProjectID != "" {
		values["project_id"] = scan.ProjectID
	}

	switch scan.ScanType {
	case chaos.ScanTypePipelineExecution:
		for key, value := range pipelineScanState(scan.Source.Pipeline) {
			values[key] = value
		}
	case chaos.ScanTypeDiscoveryAgent:
		for key, value := range infrastructureScanState(d, scan.Source.Infrastructure, keepInfrastructureIDs) {
			values[key] = value
		}
	}

	values["summary"] = heatmapState(scan.Summary)
	values["risk_summary"] = riskSummaryState(scan.RiskSummary)
	values["collected_data"] = collectedDataState(scan.CollectedData)
	if includeRisks {
		risks := []chaos.ScanRisk{}
		if report != nil {
			risks = report.ScannedRisks
		}
		values["risks"] = scanRisksState(risks)
	}

	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("set %s: %w", key, err)
		}
	}
	return nil
}

func pipelineScanState(src *chaos.PipelineScan) map[string]interface{} {
	if src == nil {
		return map[string]interface{}{
			"pipeline_identity":     "",
			"pipeline_execution_id": "",
			"notify_id":             "",
			"infrastructures":       []interface{}{},
		}
	}
	infras := make([]interface{}, 0, len(src.Infrastructures))
	for _, infra := range src.Infrastructures {
		infras = append(infras, map[string]interface{}{
			"environment_ref":         infra.EnvironmentRef,
			"infrastructure_identity": infra.InfrastructureIdentity,
			"chaos_onboarded":         infra.ChaosOnboarded,
			"chaos_supported":         infra.ChaosSupported,
		})
	}
	return map[string]interface{}{
		"pipeline_identity":     src.PipelineIdentity,
		"pipeline_execution_id": src.PipelineExecutionID,
		"notify_id":             src.NotifyID,
		"infrastructures":       infras,
	}
}

func infrastructureScanState(d *schema.ResourceData, src *chaos.InfrastructureScan, keepInfrastructureIDs bool) map[string]interface{} {
	if src == nil {
		src = &chaos.InfrastructureScan{}
	}
	// The API stores the discovery agent, not the chaos infrastructure id.
	// The resource schema keeps the infrastructure the user configured so the
	// next plan does not drop it and replace the scan. The data source schema
	// does not have those keys.
	values := map[string]interface{}{
		"agent_identity":               src.AgentIdentity,
		"environment_identity":         src.EnvironmentIdentity,
		"environment_parent_unique_id": src.EnvironmentParentUniqueID,
	}
	if keepInfrastructureIDs {
		values["infrastructure_id"] = d.Get("infrastructure_id").(string)
		values["environment_id"] = d.Get("environment_id").(string)
	}
	return values
}

func heatmapState(summary *chaos.HeatmapSummary) []interface{} {
	if summary == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"score":            int(summary.Score),
		"max_score":        int(summary.MaxScore),
		"services_scanned": summary.ServicesScanned,
		"total_risks":      summary.TotalRisks,
		"rule_count":       summary.RuleCount,
		"critical":         summary.Critical,
		"high":             summary.High,
		"medium":           summary.Medium,
		"low":              summary.Low,
	}}
}

func riskSummaryState(summary *chaos.RiskSummary) []interface{} {
	if summary == nil {
		return []interface{}{}
	}
	return []interface{}{map[string]interface{}{
		"risk_count":        summary.RiskCount,
		"analysed_at":       summary.AnalysedAt,
		"services_analysed": summary.ServicesAnalysed,
		"services_at_risk":  summary.ServicesAtRisk,
		"score":             summary.Score,
		"score_max":         summary.ScoreMax,
		"severity_counts":   intMap(summary.SeverityCounts),
		"category_counts":   intMap(summary.CategoryCounts),
	}}
}

func collectedDataState(data *chaos.CollectionSummary) []interface{} {
	if data == nil {
		return []interface{}{}
	}
	namespaces := data.Namespaces
	if namespaces == nil {
		namespaces = []string{}
	}
	return []interface{}{map[string]interface{}{
		"collected_at":        data.CollectedAt,
		"total_resources":     data.TotalResources,
		"primary_resources":   data.PrimaryResources,
		"dependent_resources": data.DependentResources,
		"service_groups":      data.ServiceGroups,
		"namespaces":          namespaces,
		"errors":              collectionErrorsState(data.Errors),
	}}
}

func collectionErrorsState(errors []chaos.CollectionError) []interface{} {
	out := make([]interface{}, 0, len(errors))
	for _, item := range errors {
		out = append(out, map[string]interface{}{
			"source":  item.Source,
			"message": item.Message,
		})
	}
	return out
}

func scanRisksState(risks []chaos.ScanRisk) []interface{} {
	ordered := append([]chaos.ScanRisk(nil), risks...)
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Identity < ordered[j].Identity
	})
	out := make([]interface{}, 0, len(ordered))
	for _, risk := range ordered {
		tags := risk.Tags
		if tags == nil {
			tags = []string{}
		}
		out = append(out, map[string]interface{}{
			"identity":         risk.Identity,
			"name":             risk.Name,
			"description":      risk.Description,
			"tags":             tags,
			"fingerprint":      risk.Fingerprint,
			"severity":         risk.Severity,
			"status":           risk.Status,
			"service_identity": risk.ServiceIdentity,
			"service_name":     risk.ServiceName,
			"service_type":     risk.ServiceType,
			"metadata_json":    metadataJSON(risk.Metadata),
			"occurrences":      occurrencesState(risk.Occurrences),
		})
	}
	return out
}

func occurrencesState(occurrences []chaos.ScanRiskOccurrence) []interface{} {
	out := make([]interface{}, 0, len(occurrences))
	for _, occurrence := range occurrences {
		out = append(out, map[string]interface{}{
			"occurrence_id":          occurrence.OccurrenceID,
			"detected_at":            occurrence.DetectedAt,
			"source_type":            occurrence.SourceType,
			"source_id":              occurrence.SourceID,
			"service_name":           occurrence.ServiceName,
			"evidence":               stringOrEmpty(occurrence.Context.Evidence),
			"remediation":            stringOrEmpty(occurrence.Context.Remediation),
			"resilience_testing_rec": stringOrEmpty(occurrence.Context.ResilienceTestingReco),
		})
	}
	return out
}

func metadataJSON(metadata map[string]interface{}) string {
	if len(metadata) == 0 {
		return ""
	}
	encoded, err := json.Marshal(metadata)
	if err != nil {
		return ""
	}
	return string(encoded)
}

func intMap(in map[string]int) map[string]interface{} {
	out := make(map[string]interface{}, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}

func expandStringList(value interface{}) []string {
	if value == nil {
		return nil
	}
	raw, ok := value.([]interface{})
	if !ok {
		return nil
	}
	out := make([]string, 0, len(raw))
	for _, item := range raw {
		if text, ok := item.(string); ok {
			out = append(out, text)
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func stringOrEmpty(values []string) []string {
	if values == nil {
		return []string{}
	}
	return values
}

func parseScanImportID(id string) (orgID, projectID, identity string, err error) {
	parts := strings.Split(id, "/")
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		return "", "", "", fmt.Errorf("invalid import ID %q, expected org_id/project_id/scan_identity", id)
	}
	return parts[0], parts[1], parts[2], nil
}
