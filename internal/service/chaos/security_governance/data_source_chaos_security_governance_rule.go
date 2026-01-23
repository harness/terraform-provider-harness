package security_governance

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceChaosSecurityGovernanceRule() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Security Governance Rule.",

		ReadContext: dataSourceChaosSecurityGovernanceRuleRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description:  "The ID of the rule.",
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"id", "name"},
			},
			"name": {
				Description:  "The name of the rule.",
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"id", "name"},
			},
			"org_id": {
				Description: "The organization identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "The project identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"is_enabled": {
				Description: "Whether the rule is enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"condition_ids": {
				Description: "List of condition IDs associated with the rule.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_group_ids": {
				Description: "List of user group IDs associated with the rule.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": {
				Description: "Tags associated with the rule.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_windows": {
				Description: "Time windows when the rule is active.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time_zone": {
							Description: "Time zone for the time window.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"start_time": {
							Description: "Start time of the time window in milliseconds since epoch.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"end_time": {
							Description: "End time of the time window in milliseconds since epoch.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"duration": {
							Description: "Duration of the time window (e.g., '30m', '2h').",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"recurrence": {
							Description: "Recurrence configuration for the time window.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "Type of recurrence (e.g., 'Daily', 'Weekly', 'Monthly').",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"until": {
										Description: "Unix timestamp in milliseconds until when the recurrence should continue.",
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"value": {
										Description: "Recurrence value (e.g., interval for daily recurrence).",
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceChaosSecurityGovernanceRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session)
	if c == nil {
		return diag.Errorf("provider configuration is nil")
	}
	securityGovernanceRuleClient := c.ChaosClient.SecurityGovernanceRuleApi

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID is required")
	}

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	var rule *model.Rule

	if ruleID, ok := d.GetOk("id"); ok {
		// Extract just the rule ID part if full path is provided
		idStr := ruleID.(string)
		parts := strings.Split(idStr, "/")
		actualRuleID := idStr
		if len(parts) == 3 {
			actualRuleID = parts[2] // Extract just the rule ID part
		}

		// Look up by ID
		identifiers := model.IdentifiersRequest{
			AccountIdentifier: accountID,
			OrgIdentifier:     orgID,
			ProjectIdentifier: projectID,
		}

		resp, err := securityGovernanceRuleClient.Get(ctx, identifiers, actualRuleID)
		if err != nil {
			if isNotFoundError(err) {
				d.SetId("")
				return diag.Diagnostics{
					diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  "Rule not found",
						Detail:   fmt.Sprintf("Security governance rule with ID '%s' not found. Removing from state.", ruleID),
					},
				}
			}
			return helpers.HandleChaosGraphQLReadError(err, d, "read_chaos_security_governance_rule")
		}
		if resp == nil || resp.Rule == nil {
			d.SetId("")
			return diag.Diagnostics{
				diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Rule not found",
					Detail:   fmt.Sprintf("Security governance rule with ID '%s' not found. Removing from state.", ruleID),
				},
			}
		}
		rule = resp.Rule
	} else if ruleName, ok := d.GetOk("name"); ok {
		nameStr := ruleName.(string)
		// Look up by name
		identifiers := model.IdentifiersRequest{
			AccountIdentifier: accountID,
			OrgIdentifier:     orgID,
			ProjectIdentifier: projectID,
		}

		rules, err := securityGovernanceRuleClient.List(ctx, identifiers, model.ListRuleRequest{})
		if err != nil {
			return helpers.HandleChaosGraphQLReadError(err, d, "read_chaos_security_governance_rule")
		}

		var found bool
		for _, r := range rules {
			if r.Rule != nil && r.Rule.Name == nameStr {
				rule = r.Rule
				found = true
				break
			}
		}

		if !found {
			d.SetId("")
			return diag.Diagnostics{
				diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Rule not found",
					Detail:   fmt.Sprintf("Security governance rule with name '%s' not found. Removing from state.", nameStr),
				},
			}
		}
	} else {
		return diag.Errorf("either 'id' or 'name' must be specified")
	}

	// Set the ID
	d.SetId(rule.RuleID)

	// Set the attributes
	if err := setRuleAttributes(d, &model.RuleResponse{Rule: rule}, accountID); err != nil {
		return helpers.HandleChaosGraphQLError(fmt.Errorf("failed to set rule attributes: %w", err), d, "read_chaos_security_governance_rule")
	}

	return nil
}

// Helper function to check if error is a "not found" error
func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "no documents in result") ||
		strings.Contains(err.Error(), "not found") ||
		strings.Contains(err.Error(), "does not exist")
}
