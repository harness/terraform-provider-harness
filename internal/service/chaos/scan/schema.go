package scan

import (
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func stringListSchema(description string, optional, computed bool) *schema.Schema {
	return &schema.Schema{
		Description: description,
		Type:        schema.TypeList,
		Optional:    optional,
		Computed:    computed,
		Elem:        &schema.Schema{Type: schema.TypeString},
	}
}

func resourceChaosScanSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"org_id": {
			Description:  "Organization identifier. Scans are project scoped, matching the UI. The agent that runs the scan is called with the account, organization, and project.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"project_id": {
			Description:  "Project identifier. Scans are project scoped, matching the UI. A scan created outside a project does not appear on the project Risks pages.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"identity": {
			Description:  "Unique identifier of this scan. This becomes the resource id. The API rejects an identifier that already exists. Another apply keeps this identifier and does not create a second scan. Change `retry_token` to run this same scan again after the infrastructure changes. Change `identity` only to replace this scan with a different one; Terraform then deletes the current scan.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"name": {
			Description:  "Name of the scan.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"description": {
			Description: "Description of the scan.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": stringListSchema("Tags associated with the scan.", true, false),
		"retry_token": {
			Description: "Change this to a new value to run this scan again, using the same API as the UI retry action. The scan identifier stays the same, so this does not add another scan on the UI. Set it to `harness_chaos_infrastructure_v2.<name>.updated_at` to run again only when that Terraform infrastructure is created or updated. An apply that leaves this value unchanged does not scan again. Terraform waits until the scan is COMPLETED, ERRORED, or ABORTED. ERRORED and ABORTED fail the apply.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"abort_token": {
			Description: "Change this on an existing scan to abort it, using the same API as the UI. `ABORTED` is a successful apply. Create holds the state lock, so this cannot abort a create that is still running. Do not change this and `retry_token` in the same apply.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"detection_method": {
			Description: "Scan method, matching the UI dropdown. `ai` is AI Scan. `ruleBased` is Rule-based Scan. `AI` and `RiskAnalyzer` are the same choices and remain accepted. Leave this unset to let the API choose AI Scan. Changing between AI Scan and Rule-based Scan creates a new scan. Another apply that keeps the same method does not.",
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			DiffSuppressFunc: func(_, oldValue, newValue string, _ *schema.ResourceData) bool {
				oldMethod, oldErr := chaos.APIDetectionMethod(oldValue)
				newMethod, newErr := chaos.APIDetectionMethod(newValue)
				return oldErr == nil && newErr == nil && oldMethod == newMethod
			},
			ValidateFunc: func(v interface{}, k string) (ws []string, errs []error) {
				value, _ := v.(string)
				if _, err := chaos.APIDetectionMethod(value); err != nil {
					errs = append(errs, err)
				}
				return
			},
		},
		"pipeline_scan":       pipelineScanSchema(true),
		"infrastructure_scan": infrastructureScanSchema(true),

		"scan_type": {
			Description: "Scan type recorded by the API. `PipelineExecution` for a pipeline scan, `DiscoveryAgent` for an infrastructure scan.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"status": {
			Description: "Scan lifecycle status: QUEUED, COLLECTING, COLLECTED, ANALYSING, COMPLETED, ERRORED, or ABORTED. Risks are populated once the scan completes.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"result": {
			Description: "Outcome of a completed scan, such as SUCCESS or NO_RESOURCES_FOUND.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"error_message": {
			Description: "Error recorded when the scan fails.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"retry_count": {
			Description: "Number of times the scan has been retried.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"account_id": {
			Description: "Account identifier.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"created_at": {
			Description: "Creation time as a Unix timestamp.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"updated_at": {
			Description: "Last update time as a Unix timestamp.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"completed_at": {
			Description: "Completion time as a Unix timestamp. Empty until the scan finishes.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"created_by": {
			Description: "User who created the scan.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"updated_by": {
			Description: "User who last updated the scan.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"summary":        heatmapSummarySchema(),
		"risk_summary":   riskSummarySchema(),
		"collected_data": collectedDataSchema(),
		"risks": {
			Description: "Computed view of the risks embedded in this scan's report, sorted by identity. Empty until analysis finishes, including a COMPLETED scan that found nothing. This is not the durable scanned risk. A future risk resource should use the scanned-risk API, which is a different object keyed by the same identity.",
			Type:        schema.TypeList,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: scanReportRiskSchema(),
			},
		},
	}
}

func dataSourceChaosScanSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"org_id": {
			Description:  "Organization identifier of the project that owns the scan.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"project_id": {
			Description:  "Project identifier of the scan.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"identity": {
			Description:  "Identifier of the scan to read.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"name": {
			Description: "Name of the scan.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"description": {
			Description: "Description of the scan.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"tags":                stringListSchema("Tags associated with the scan.", false, true),
		"detection_method":    {Description: "Detection method used by the scan.", Type: schema.TypeString, Computed: true},
		"pipeline_scan":       pipelineScanSchema(false),
		"infrastructure_scan": infrastructureScanSchema(false),
		"scan_type":           {Description: "Scan type recorded by the API.", Type: schema.TypeString, Computed: true},
		"status":              {Description: "Scan lifecycle status.", Type: schema.TypeString, Computed: true},
		"result":              {Description: "Outcome of a completed scan.", Type: schema.TypeString, Computed: true},
		"error_message":       {Description: "Error recorded when the scan fails.", Type: schema.TypeString, Computed: true},
		"retry_count":         {Description: "Number of times the scan has been retried.", Type: schema.TypeInt, Computed: true},
		"account_id":          {Description: "Account identifier.", Type: schema.TypeString, Computed: true},
		"created_at":          {Description: "Creation time as a Unix timestamp.", Type: schema.TypeInt, Computed: true},
		"updated_at":          {Description: "Last update time as a Unix timestamp.", Type: schema.TypeInt, Computed: true},
		"completed_at":        {Description: "Completion time as a Unix timestamp.", Type: schema.TypeInt, Computed: true},
		"created_by":          {Description: "User who created the scan.", Type: schema.TypeString, Computed: true},
		"updated_by":          {Description: "User who last updated the scan.", Type: schema.TypeString, Computed: true},
		"summary":             heatmapSummarySchema(),
		"risk_summary":        riskSummarySchema(),
		"collected_data":      collectedDataSchema(),
		"risks": {
			Description: "Risks produced by this scan.",
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        &schema.Resource{Schema: scanReportRiskSchema()},
		},
	}
}

func pipelineScanSchema(forResource bool) *schema.Schema {
	pipelineIdentity := &schema.Schema{
		Description: "Identifier of the pipeline to scan.",
		Type:        schema.TypeString,
	}
	executionID := &schema.Schema{
		Description: "Execution id (`planExecutionId`) of a pipeline run whose status is Success. The UI only offers those executions. Collection fails when the execution has no Kubernetes CD stage.",
		Type:        schema.TypeString,
	}
	notifyID := &schema.Schema{
		Description: "Notification id used when a pipeline step triggers the scan. Omit this for scans created the same way as the UI.",
		Type:        schema.TypeString,
	}
	if forResource {
		pipelineIdentity.Required = true
		pipelineIdentity.ForceNew = true
		pipelineIdentity.ValidateFunc = validation.StringIsNotEmpty
		executionID.Required = true
		executionID.ForceNew = true
		executionID.ValidateFunc = validation.StringIsNotEmpty
		notifyID.Optional = true
		notifyID.Computed = true
	} else {
		pipelineIdentity.Computed = true
		executionID.Computed = true
		notifyID.Computed = true
	}

	description := "Pipeline execution recorded on the scan. Set when the scan type is PipelineExecution."
	if forResource {
		description = "Pipeline scan. Scans one pipeline execution and records scan type PipelineExecution. Exactly one of `pipeline_scan` or `infrastructure_scan` must be set."
	}
	s := &schema.Schema{
		Description: description,
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"pipeline_identity":     pipelineIdentity,
				"pipeline_execution_id": executionID,
				"notify_id":             notifyID,
				"infrastructures": {
					Description: "CD environment and infrastructure pairs observed while collecting manifests. Populated by the scan.",
					Type:        schema.TypeList,
					Computed:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"environment_ref":         {Description: "Environment reference.", Type: schema.TypeString, Computed: true},
							"infrastructure_identity": {Description: "Infrastructure identifier.", Type: schema.TypeString, Computed: true},
							"chaos_onboarded":         {Description: "Whether chaos is already onboarded for this infrastructure.", Type: schema.TypeBool, Computed: true},
							"chaos_supported":         {Description: "Whether this infrastructure can be onboarded to chaos.", Type: schema.TypeBool, Computed: true},
						},
					},
				},
			},
		},
	}
	if forResource {
		s.Optional = true
		s.MaxItems = 1
		s.ExactlyOneOf = []string{"pipeline_scan", "infrastructure_scan"}
	} else {
		s.Computed = true
	}
	return s
}

func infrastructureScanSchema(forResource bool) *schema.Schema {
	infrastructureID := &schema.Schema{
		Description: "Identifier of the chaos infrastructure to scan. Same value as `harness_chaos_infrastructure_v2.infra_id`. Terraform reads the discovery agent linked to this infrastructure, which is what the UI selects for Scan infrastructure.",
		Type:        schema.TypeString,
	}
	chaosEnvironmentID := &schema.Schema{
		Description: "Environment of the chaos infrastructure. Same value as `harness_chaos_infrastructure_v2.environment_id`. Required with `infrastructure_id`.",
		Type:        schema.TypeString,
	}
	agent := &schema.Schema{
		Description: "Identity of the service discovery agent that observed the Kubernetes infrastructure. Set this when the agent is already known, including imports. Leave it unset when `infrastructure_id` is set; Terraform fills it from the linked agent.",
		Type:        schema.TypeString,
	}
	environment := &schema.Schema{
		Description: "Environment identifier of the discovery agent. Required with `agent_identity`. When `infrastructure_id` is set, Terraform fills this from the agent record, matching the UI.",
		Type:        schema.TypeString,
	}
	parent := &schema.Schema{
		Description: "Parent unique id of the discovery agent's environment. The UI sends this from the agent record. Terraform fills it when `infrastructure_id` is set.",
		Type:        schema.TypeString,
	}
	if forResource {
		infrastructureID.Optional = true
		infrastructureID.ForceNew = true
		infrastructureID.ValidateFunc = validation.StringIsNotEmpty
		infrastructureID.RequiredWith = []string{"infrastructure_scan.0.environment_id"}
		infrastructureID.ConflictsWith = []string{
			"infrastructure_scan.0.agent_identity",
			"infrastructure_scan.0.environment_identity",
		}
		infrastructureID.AtLeastOneOf = []string{
			"infrastructure_scan.0.infrastructure_id",
			"infrastructure_scan.0.agent_identity",
		}
		chaosEnvironmentID.Optional = true
		chaosEnvironmentID.ForceNew = true
		chaosEnvironmentID.ValidateFunc = validation.StringIsNotEmpty
		chaosEnvironmentID.RequiredWith = []string{"infrastructure_scan.0.infrastructure_id"}
		chaosEnvironmentID.ConflictsWith = []string{
			"infrastructure_scan.0.agent_identity",
			"infrastructure_scan.0.environment_identity",
		}
		agent.Optional = true
		agent.Computed = true
		agent.ValidateFunc = validation.StringIsNotEmpty
		agent.RequiredWith = []string{"infrastructure_scan.0.environment_identity"}
		agent.ConflictsWith = []string{
			"infrastructure_scan.0.infrastructure_id",
			"infrastructure_scan.0.environment_id",
		}
		agent.AtLeastOneOf = []string{
			"infrastructure_scan.0.infrastructure_id",
			"infrastructure_scan.0.agent_identity",
		}
		environment.Optional = true
		environment.Computed = true
		environment.ValidateFunc = validation.StringIsNotEmpty
		environment.RequiredWith = []string{"infrastructure_scan.0.agent_identity"}
		environment.ConflictsWith = []string{
			"infrastructure_scan.0.infrastructure_id",
			"infrastructure_scan.0.environment_id",
		}
		parent.Optional = true
		parent.Computed = true
	} else {
		infrastructureID.Computed = true
		chaosEnvironmentID.Computed = true
		agent.Computed = true
		environment.Computed = true
		parent.Computed = true
	}

	description := "Discovery agent recorded on the scan. Set when the scan type is DiscoveryAgent."
	if forResource {
		description = "Infrastructure scan. Scans the Kubernetes infrastructure known to a discovery agent and records scan type DiscoveryAgent. Set `infrastructure_id` and `environment_id` from the chaos infrastructure, or set `agent_identity` and `environment_identity` when the discovery agent is already known. Exactly one of `pipeline_scan` or `infrastructure_scan` must be set."
	}
	s := &schema.Schema{
		Description: description,
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"infrastructure_id":            infrastructureID,
				"environment_id":               chaosEnvironmentID,
				"agent_identity":               agent,
				"environment_identity":         environment,
				"environment_parent_unique_id": parent,
			},
		},
	}
	if forResource {
		s.Optional = true
		s.MaxItems = 1
		s.ExactlyOneOf = []string{"pipeline_scan", "infrastructure_scan"}
	} else {
		s.Computed = true
	}
	return s
}

func heatmapSummarySchema() *schema.Schema {
	num := func(desc string) *schema.Schema {
		return &schema.Schema{Description: desc, Type: schema.TypeInt, Computed: true}
	}
	return &schema.Schema{
		Description: "Severity rollup for the scan.",
		Type:        schema.TypeList,
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"score":            num("Resilience score."),
				"max_score":        num("Maximum resilience score."),
				"services_scanned": num("Number of services scanned."),
				"total_risks":      num("Number of risks found."),
				"rule_count":       num("Number of risk rules evaluated."),
				"critical":         num("Critical risks."),
				"high":             num("High risks."),
				"medium":           num("Medium risks."),
				"low":              num("Low risks."),
			},
		},
	}
}

func riskSummarySchema() *schema.Schema {
	num := func(desc string) *schema.Schema {
		return &schema.Schema{Description: desc, Type: schema.TypeInt, Computed: true}
	}
	return &schema.Schema{
		Description: "Analysis summary written when the scan finishes.",
		Type:        schema.TypeList,
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"risk_count":        num("Number of risks."),
				"analysed_at":       num("Analysis time as a Unix timestamp."),
				"services_analysed": num("Services included in analysis."),
				"services_at_risk":  num("Services with at least one risk."),
				"score":             num("Resilience score."),
				"score_max":         num("Maximum resilience score."),
				"severity_counts":   {Description: "Risk counts keyed by severity.", Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeInt}},
				"category_counts":   {Description: "Risk counts keyed by category.", Type: schema.TypeMap, Computed: true, Elem: &schema.Schema{Type: schema.TypeInt}},
			},
		},
	}
}

func collectedDataSchema() *schema.Schema {
	num := func(desc string) *schema.Schema {
		return &schema.Schema{Description: desc, Type: schema.TypeInt, Computed: true}
	}
	return &schema.Schema{
		Description: "Resources collected before analysis.",
		Type:        schema.TypeList,
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"collected_at":        num("Collection time as a Unix timestamp."),
				"total_resources":     num("Total resources collected."),
				"primary_resources":   num("Primary resources collected."),
				"dependent_resources": num("Dependent resources collected."),
				"service_groups":      num("Service groups collected."),
				"namespaces":          stringListSchema("Namespaces collected.", false, true),
				"errors": {
					Description: "Collection failures returned by the API. A pipeline execution with no Kubernetes CD stage is reported here or as a scan error.",
					Type:        schema.TypeList,
					Computed:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"source":  {Description: "Collector that reported the error.", Type: schema.TypeString, Computed: true},
							"message": {Description: "Collection error message.", Type: schema.TypeString, Computed: true},
						},
					},
				},
			},
		},
	}
}

// scanReportRiskSchema is the flattened risk embedded in a scan report.
// It is not the schema for a future harness_chaos_scanned_risk resource.
// That resource should follow ScannedRisk on the scanned-risk API, whose
// spec holds the rule, fingerprint, severity, status, and target service.
func scanReportRiskSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"identity": {
			Description: "Identity of the scanned risk. Stable across scans that find the same risk.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"name": {
			Description: "Name of the risk.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"description": {
			Description: "Description of the risk.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"tags":             stringListSchema("Tags on the risk.", false, true),
		"fingerprint":      {Description: "Fingerprint that groups occurrences of the same risk.", Type: schema.TypeString, Computed: true},
		"severity":         {Description: "Severity: Critical, High, Medium, or Low.", Type: schema.TypeString, Computed: true},
		"status":           {Description: "Status: Open, InProgress, Resolved, or Ignored.", Type: schema.TypeString, Computed: true},
		"service_identity": {Description: "Identity of the service the risk applies to.", Type: schema.TypeString, Computed: true},
		"service_name":     {Description: "Name of the service the risk applies to.", Type: schema.TypeString, Computed: true},
		"service_type":     {Description: "Service kind, such as DiscoveredService, ChaosService, or CDService.", Type: schema.TypeString, Computed: true},
		"metadata_json": {
			Description: "Additional risk metadata as a JSON object string.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"occurrences": {
			Description: "Detections of this risk in the scan.",
			Type:        schema.TypeList,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"occurrence_id":          {Description: "Identity of this occurrence.", Type: schema.TypeString, Computed: true},
					"detected_at":            {Description: "Detection time as a Unix timestamp.", Type: schema.TypeInt, Computed: true},
					"source_type":            {Description: "Source type of the occurrence.", Type: schema.TypeString, Computed: true},
					"source_id":              {Description: "Source id of the occurrence. For a scan this is the scan identity.", Type: schema.TypeString, Computed: true},
					"service_name":           {Description: "Service name recorded on the occurrence.", Type: schema.TypeString, Computed: true},
					"evidence":               stringListSchema("Evidence for the occurrence.", false, true),
					"remediation":            stringListSchema("Remediation guidance.", false, true),
					"resilience_testing_rec": stringListSchema("Resilience testing recommendations.", false, true),
				},
			},
		},
	}
}
