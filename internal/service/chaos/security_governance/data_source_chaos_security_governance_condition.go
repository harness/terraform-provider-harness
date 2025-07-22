package security_governance

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	dataSourceName = "harness_chaos_security_governance_condition"
)

func DataSourceChaosSecurityGovernanceCondition() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Security Governance Condition",
		ReadContext: dataSourceChaosSecurityGovernanceConditionRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "The organization ID of the security governance condition",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "The project ID of the security governance condition",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"id": {
				Description:  "The ID of the security governance condition. Either `id` or `name` must be specified.",
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"id", "name"},
				Computed:     true,
			},
			"name": {
				Description:  "The name of the security governance condition. Either `id` or `name` must be specified.",
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"id", "name"},
				Computed:     true,
			},
			"description": {
				Description: "The description of the security governance condition",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"infra_type": {
				Description: "The infrastructure type (KubernetesV2, Linux, Windows)",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"fault_spec": {
				Description: "Fault specification for the condition",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"operator": {
							Description: "Operator for the fault specification",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"faults": {
							Description: "List of fault specifications",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fault_type": {
										Description: "Type of the fault",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"name": {
										Description: "Name of the fault",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"k8s_spec": {
				Description: "Kubernetes specific configuration",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"infra_spec": {
							Description: "Infrastructure specification",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"operator": {
										Description: "Operator for comparing infrastructure IDs",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"infra_ids": {
										Description: "List of infrastructure IDs",
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"machine_spec": {
				Description: "Machine specific configuration",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"infra_spec": {
							Description: "Infrastructure specification",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"operator": {
										Description: "Operator for comparing infrastructure IDs",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"infra_ids": {
										Description: "List of infrastructure IDs",
										Type:        schema.TypeList,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"tags": {
				Description: "Tags associated with the condition",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

// safeString returns a string representation of a string pointer, or "<nil>" if the pointer is nil
func safeString(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}

// validationError creates a consistent error format for validation errors
func validationError(field, message string, args ...interface{}) diag.Diagnostics {
	msg := fmt.Sprintf("invalid %s: %s", field, fmt.Sprintf(message, args...))
	return diag.Diagnostics{
		{
			Severity: diag.Error,
			Summary:  "Validation Error",
			Detail:   msg,
		},
	}
}

// apiError creates a consistent error format for API errors
func apiError(operation string, err error) diag.Diagnostics {
	msg := fmt.Sprintf("failed to %s: %v", operation, err)
	return diag.Diagnostics{
		{
			Severity: diag.Error,
			Summary:  "API Error",
			Detail:   msg,
		},
	}
}

// logTrace logs a trace message with context
func logTrace(ctx context.Context, msg string, args ...interface{}) {
	tflog.Trace(ctx, fmt.Sprintf(msg, args...))
}

// logDebug logs a debug message with context
func logDebug(ctx context.Context, msg string, args ...interface{}) {
	tflog.Debug(ctx, fmt.Sprintf(msg, args...))
}

// logError logs an error message with context
func logError(ctx context.Context, msg string, args ...interface{}) {
	tflog.Error(ctx, fmt.Sprintf(msg, args...))
}

func dataSourceChaosSecurityGovernanceConditionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Add trace logs for function entry and exit
	logTrace(ctx, "Starting dataSourceChaosSecurityGovernanceConditionRead")
	defer logTrace(ctx, "Completed dataSourceChaosSecurityGovernanceConditionRead")

	c := meta.(*internal.Session).ChaosClient
	if c == nil {
		err := errors.New("Chaos client is not properly initialized")
		logError(ctx, err.Error())
		return diag.FromErr(err)
	}

	client := chaos.NewSecurityGovernanceConditionClient(c)

	// Get the account ID from the provider config
	accountID := c.AccountId
	if accountID == "" {
		err := errors.New("account ID must be configured in the provider")
		logError(ctx, err.Error())
		return validationError("account_id", "must be configured in the provider")
	}

	// Add account ID to context for logging
	ctx = tflog.SetField(ctx, "account_id", accountID)

	// Check if we're looking up by ID or name
	var conditionID string
	var orgID, projectID *string

	if id, ok := d.GetOk("id"); ok {
		// The ID might be in the format "account/org/project/condition-id" or just the condition ID
		idStr, ok := id.(string)
		if !ok {
			err := fmt.Errorf("expected 'id' to be a string, got %T", id)
			logError(ctx, err.Error())
			return validationError("id", "must be a string")
		}

		logDebug(ctx, "Looking up condition by ID", "id", idStr)

		if strings.Contains(idStr, "/") {
			// Parse the full ID format: account/org/project/condition-id
			parts := strings.Split(idStr, "/")
			if len(parts) >= 4 {
				// If we have a full path, use the org and project from it
				orgID = &parts[1]
				projectID = &parts[2]
				conditionID = parts[3]
				logDebug(ctx, "Parsed ID components", "org_id", *orgID, "project_id", *projectID, "condition_id", conditionID)
			} else {
				// Handle unexpected format
				conditionID = parts[len(parts)-1]
				logDebug(ctx, "Using last part of ID as condition ID", "condition_id", conditionID)
			}
		} else {
			conditionID = idStr
		}
	} else if name, ok := d.GetOk("name"); ok {
		// For name-based lookup, use the provided org/project or default to the one in the config
		nameStr, ok := name.(string)
		if !ok {
			err := fmt.Errorf("expected 'name' to be a string, got %T", name)
			logError(ctx, err.Error())
			return validationError("name", "must be a string")
		}

		logDebug(ctx, "Looking up condition by name", "name", nameStr)

		if orgIDVal, ok := d.GetOk("org_id"); ok {
			orgIDStr, ok := orgIDVal.(string)
			if !ok {
				err := fmt.Errorf("expected 'org_id' to be a string, got %T", orgIDVal)
				logError(ctx, err.Error())
				return validationError("org_id", "must be a string")
			}
			orgID = &orgIDStr
		}

		if projectIDVal, ok := d.GetOk("project_id"); ok {
			projectIDStr, ok := projectIDVal.(string)
			if !ok {
				err := fmt.Errorf("expected 'project_id' to be a string, got %T", projectIDVal)
				logError(ctx, err.Error())
				return validationError("project_id", "must be a string")
			}
			projectID = &projectIDStr
		}

		logDebug(ctx, "Listing all conditions for name-based lookup")

		// List all conditions and find the one with the matching name
		listReq := model.IdentifiersRequest{
			AccountIdentifier: accountID,
		}

		if orgID != nil {
			listReq.OrgIdentifier = *orgID
		}

		if projectID != nil {
			listReq.ProjectIdentifier = *projectID
		}

		// Log the request parameters for debugging
		logDebug(ctx, "Listing conditions with request",
			"account_id", listReq.AccountIdentifier,
			"org_id", safeString(orgID),
			"project_id", safeString(projectID))

		// Create an empty ListConditionRequest since it's required but all fields are optional
		listConditionReq := model.ListConditionRequest{}
		conditions, err := client.List(ctx, listReq, listConditionReq)
		if err != nil {
			logError(ctx, "Failed to list conditions", "error", err.Error())
			return apiError("list security governance conditions", err)
		}

		logDebug(ctx, "Received conditions list response", "count", len(conditions.Conditions))

		found := false
		if conditions != nil && conditions.Conditions != nil {
			for i, c := range conditions.Conditions {
				condition := c.Condition
				if condition == nil {
					logDebug(ctx, "Skipping nil condition in response", "index", i)
					continue
				}

				logTrace(ctx, "Checking condition",
					"index", i,
					"condition_id", condition.ConditionID,
					"name", condition.Name)

				if condition.Name == nameStr {
					conditionID = condition.ConditionID
					found = true
					logDebug(ctx, "Found matching condition",
						"condition_id", conditionID,
						"name", nameStr)
					break
				}
			}
		} else {
			logDebug(ctx, "No conditions found in the response")
		}

		if !found {
			err := fmt.Errorf("no security governance condition found with name: %s", nameStr)
			logError(ctx, err.Error())
			return diag.FromErr(err)
		}
	} else {
		err := errors.New("either 'id' or 'name' must be provided")
		logError(ctx, err.Error())
		return validationError("id/name", "either 'id' or 'name' must be provided")
	}

	// Set up the identifiers for the final lookup
	identifiers := ScopedIdentifiersRequest{
		AccountIdentifier: accountID,
	}

	// Use the org/project from the ID if available, otherwise from the config
	if orgID != nil {
		identifiers.OrgIdentifier = orgID
	} else if orgIDVal, ok := d.GetOk("org_id"); ok {
		orgIDStr := orgIDVal.(string)
		identifiers.OrgIdentifier = &orgIDStr
	}

	if projectID != nil {
		identifiers.ProjectIdentifier = projectID
	} else if projectIDVal, ok := d.GetOk("project_id"); ok {
		projectIDStr := projectIDVal.(string)
		identifiers.ProjectIdentifier = &projectIDStr
	}

	// Log the get request details
	logDebug(ctx, "Fetching condition details",
		"condition_id", conditionID,
		"account_id", identifiers.AccountIdentifier,
		"org_id", safeString(identifiers.OrgIdentifier),
		"project_id", safeString(identifiers.ProjectIdentifier))

	// Prepare the request
	getReq := model.IdentifiersRequest{
		AccountIdentifier: identifiers.AccountIdentifier,
	}

	// Only include org and project if they are not nil
	if identifiers.OrgIdentifier != nil {
		getReq.OrgIdentifier = *identifiers.OrgIdentifier
	}

	if identifiers.ProjectIdentifier != nil {
		getReq.ProjectIdentifier = *identifiers.ProjectIdentifier
	}

	// Make the API call
	resp, err := client.Get(ctx, getReq, conditionID)

	// Handle errors
	if err != nil {
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no documents in result") {
			logDebug(ctx, "Condition not found", "condition_id", conditionID)
			d.SetId("")
			return nil
		}
		logError(ctx, "Failed to get condition",
			"condition_id", conditionID,
			"error", err.Error())
		return apiError(fmt.Sprintf("get security governance condition %s", conditionID), err)
	}

	// Validate the response
	if resp == nil {
		err := errors.New("received nil response from API")
		logError(ctx, err.Error())
		return diag.FromErr(err)
	}

	if resp.Condition == nil {
		logDebug(ctx, "Condition not found in response", "condition_id", conditionID)
		d.SetId("")
		return nil
	}

	// Set the ID and other fields
	condition := resp.Condition
	logDebug(ctx, "Successfully retrieved condition",
		"condition_id", conditionID,
		"name", condition.Name,
		"infra_type", condition.InfraType)

	d.SetId(conditionID)
	d.Set("name", condition.Name)
	d.Set("description", condition.Description)
	d.Set("infra_type", condition.InfraType)

	// Set tags if they exist
	if len(condition.Tags) > 0 {
		tags := make([]string, 0, len(condition.Tags))
		for _, tag := range condition.Tags {
			if tag != "" { // Skip empty tags
				tags = append(tags, tag)
			}
		}
		if len(tags) > 0 {
			if err := d.Set("tags", tags); err != nil {
				logError(ctx, "Failed to set tags", "error", err.Error())
				return diag.FromErr(err)
			}
		}
	}

	// Set fault spec if it exists
	if condition.FaultSpec != nil {
		logDebug(ctx, "Setting fault spec", "operator", string(condition.FaultSpec.Operator))

		faultSpec := map[string]interface{}{
			"operator": string(condition.FaultSpec.Operator),
		}

		// Convert faults to the expected format
		if condition.FaultSpec.Faults != nil {
			faults := make([]map[string]string, 0, len(condition.FaultSpec.Faults))
			for i, f := range condition.FaultSpec.Faults {
				if f == nil {
					logDebug(ctx, "Skipping nil fault at index", "index", i)
					continue
				}

				faults = append(faults, map[string]string{
					"fault_type": string(f.FaultType),
					"name":       f.Name,
				})
			}

			if len(faults) > 0 {
				faultSpec["faults"] = faults
			}
		}

		if err := d.Set("fault_spec", []interface{}{faultSpec}); err != nil {
			logError(ctx, "Failed to set fault_spec", "error", err.Error())
			return diag.FromErr(err)
		}
	}

	// Set K8s spec if it exists
	if condition.K8sSpec != nil && condition.K8sSpec.InfraSpec != nil {
		logDebug(ctx, "Setting K8s spec",
			"operator", string(condition.K8sSpec.InfraSpec.Operator),
			"infra_ids_count", len(condition.K8sSpec.InfraSpec.InfraIds))

		k8sSpec := map[string]interface{}{
			"infra_spec": []map[string]interface{}{
				{
					"operator":  string(condition.K8sSpec.InfraSpec.Operator),
					"infra_ids": condition.K8sSpec.InfraSpec.InfraIds,
				},
			},
		}

		if err := d.Set("k8s_spec", []interface{}{k8sSpec}); err != nil {
			logError(ctx, "Failed to set k8s_spec", "error", err.Error())
			return diag.FromErr(err)
		}
	}

	// Set Machine spec if it exists
	if condition.MachineSpec != nil && condition.MachineSpec.InfraSpec != nil {
		logDebug(ctx, "Setting Machine spec",
			"operator", string(condition.MachineSpec.InfraSpec.Operator),
			"infra_ids_count", len(condition.MachineSpec.InfraSpec.InfraIds))

		machineSpec := map[string]interface{}{
			"infra_spec": []map[string]interface{}{
				{
					"operator":  string(condition.MachineSpec.InfraSpec.Operator),
					"infra_ids": condition.MachineSpec.InfraSpec.InfraIds,
				},
			},
		}

		if err := d.Set("machine_spec", []interface{}{machineSpec}); err != nil {
			logError(ctx, "Failed to set machine_spec", "error", err.Error())
			return diag.FromErr(err)
		}
	}

	logDebug(ctx, "Successfully processed condition", "condition_id", conditionID)
	return nil
}
