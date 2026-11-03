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

func resourceChaosPipelineScanSchema() map[string]*schema.Schema {
	return mergeScanSchema(resourceChaosScanSchema(), pipelineScanFields(true))
}

func resourceChaosInfrastructureScanSchema() map[string]*schema.Schema {
	return mergeScanSchema(resourceChaosScanSchema(), infrastructureScanFields(true))
}

func resourceChaosScanSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"org_id": {
			Description:  "Organization identifier of the scan.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"project_id": {
			Description:  "Project identifier of the scan.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"identity": {
			Description:  "Identifier of the scan.",
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
			Description: "Change this to run the same scan again.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"abort_token": {
			Description: "Change this to abort the scan.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"detection_method": {
			Description: "Scan method. `ai` or `ruleBased`.",
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
		"scan_type": {
			Description: "Scan type recorded by the API.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"status": {
			Description: "Scan lifecycle status.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"result": {
			Description: "Outcome of a completed scan.",
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
			Description: "Completion time as a Unix timestamp.",
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
			Description: "Risks produced by this scan.",
			Type:        schema.TypeList,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: scanReportRiskSchema(),
			},
		},
	}
}

func dataSourceChaosPipelineScanSchema() map[string]*schema.Schema {
	return mergeScanSchema(dataSourceChaosScanSchema(), pipelineScanFields(false))
}

func dataSourceChaosInfrastructureScanSchema() map[string]*schema.Schema {
	return mergeScanSchema(dataSourceChaosScanSchema(), infrastructureScanFields(false))
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
		"tags":             stringListSchema("Tags associated with the scan.", false, true),
		"detection_method": {Description: "Detection method used by the scan.", Type: schema.TypeString, Computed: true},
		"scan_type":        {Description: "Scan type recorded by the API.", Type: schema.TypeString, Computed: true},
		"status":           {Description: "Scan lifecycle status.", Type: schema.TypeString, Computed: true},
		"result":           {Description: "Outcome of a completed scan.", Type: schema.TypeString, Computed: true},
		"error_message":    {Description: "Error recorded when the scan fails.", Type: schema.TypeString, Computed: true},
		"retry_count":      {Description: "Number of times the scan has been retried.", Type: schema.TypeInt, Computed: true},
		"account_id":       {Description: "Account identifier.", Type: schema.TypeString, Computed: true},
		"created_at":       {Description: "Creation time as a Unix timestamp.", Type: schema.TypeInt, Computed: true},
		"updated_at":       {Description: "Last update time as a Unix timestamp.", Type: schema.TypeInt, Computed: true},
		"completed_at":     {Description: "Completion time as a Unix timestamp.", Type: schema.TypeInt, Computed: true},
		"created_by":       {Description: "User who created the scan.", Type: schema.TypeString, Computed: true},
		"updated_by":       {Description: "User who last updated the scan.", Type: schema.TypeString, Computed: true},
		"summary":          heatmapSummarySchema(),
		"risk_summary":     riskSummarySchema(),
		"collected_data":   collectedDataSchema(),
		"risks": {
			Description: "Risks produced by this scan.",
			Type:        schema.TypeList,
			Computed:    true,
			Elem:        &schema.Resource{Schema: scanReportRiskSchema()},
		},
	}
}

func mergeScanSchema(base, extra map[string]*schema.Schema) map[string]*schema.Schema {
	out := make(map[string]*schema.Schema, len(base)+len(extra))
	for key, value := range base {
		out[key] = value
	}
	for key, value := range extra {
		out[key] = value
	}
	return out
}

func pipelineScanFields(forResource bool) map[string]*schema.Schema {
	pipelineIdentity := &schema.Schema{
		Description: "Identifier of the existing pipeline to scan.",
		Type:        schema.TypeString,
	}
	executionID := &schema.Schema{
		Description: "Execution id of a successful pipeline run.",
		Type:        schema.TypeString,
	}
	notifyID := &schema.Schema{
		Description: "Notification id when a pipeline step triggers the scan.",
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
	return map[string]*schema.Schema{
		"pipeline_identity":     pipelineIdentity,
		"pipeline_execution_id": executionID,
		"notify_id":             notifyID,
		"infrastructures": {
			Description: "CD environment and infrastructure pairs observed by the scan.",
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
	}
}

func infrastructureScanFields(forResource bool) map[string]*schema.Schema {
	infrastructureID := &schema.Schema{
		Description: "Identifier of the chaos infrastructure to scan. Required with `environment_id`.",
		Type:        schema.TypeString,
	}
	chaosEnvironmentID := &schema.Schema{
		Description: "Environment of the chaos infrastructure. Required with `infrastructure_id`.",
		Type:        schema.TypeString,
	}
	agent := &schema.Schema{
		Description: "Identity of the service discovery agent. Required with `environment_identity`.",
		Type:        schema.TypeString,
	}
	environment := &schema.Schema{
		Description: "Environment identifier of the discovery agent. Required with `agent_identity`.",
		Type:        schema.TypeString,
	}
	parent := &schema.Schema{
		Description: "Parent unique id of the discovery agent's environment.",
		Type:        schema.TypeString,
	}
	if forResource {
		infrastructureID.Optional = true
		infrastructureID.ForceNew = true
		infrastructureID.ValidateFunc = validation.StringIsNotEmpty
		infrastructureID.ConflictsWith = []string{"agent_identity", "environment_identity"}
		infrastructureID.AtLeastOneOf = []string{"infrastructure_id", "agent_identity"}
		infrastructureID.RequiredWith = []string{"environment_id"}
		chaosEnvironmentID.Optional = true
		chaosEnvironmentID.ForceNew = true
		chaosEnvironmentID.ValidateFunc = validation.StringIsNotEmpty
		chaosEnvironmentID.ConflictsWith = []string{"agent_identity", "environment_identity"}
		chaosEnvironmentID.RequiredWith = []string{"infrastructure_id"}
		agent.Optional = true
		agent.Computed = true
		agent.ValidateFunc = validation.StringIsNotEmpty
		agent.RequiredWith = []string{"environment_identity"}
		agent.ConflictsWith = []string{"infrastructure_id", "environment_id"}
		agent.AtLeastOneOf = []string{"infrastructure_id", "agent_identity"}
		environment.Optional = true
		environment.Computed = true
		environment.ValidateFunc = validation.StringIsNotEmpty
		environment.RequiredWith = []string{"agent_identity"}
		environment.ConflictsWith = []string{"infrastructure_id", "environment_id"}
		parent.Optional = true
		parent.Computed = true
	} else {
		agent.Computed = true
		environment.Computed = true
		parent.Computed = true
		return map[string]*schema.Schema{
			"agent_identity":               agent,
			"environment_identity":         environment,
			"environment_parent_unique_id": parent,
		}
	}
	return map[string]*schema.Schema{
		"infrastructure_id":            infrastructureID,
		"environment_id":               chaosEnvironmentID,
		"agent_identity":               agent,
		"environment_identity":         environment,
		"environment_parent_unique_id": parent,
	}
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
					Description: "Collection failures returned by the API.",
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

func scanReportRiskSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"identity": {
			Description: "Identity of the risk.",
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
		"fingerprint":      {Description: "Fingerprint of the risk.", Type: schema.TypeString, Computed: true},
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
					"source_id":              {Description: "Source id of the occurrence.", Type: schema.TypeString, Computed: true},
					"service_name":           {Description: "Service name recorded on the occurrence.", Type: schema.TypeString, Computed: true},
					"evidence":               stringListSchema("Evidence for the occurrence.", false, true),
					"remediation":            stringListSchema("Remediation guidance.", false, true),
					"resilience_testing_rec": stringListSchema("Resilience testing recommendations.", false, true),
				},
			},
		},
	}
}
