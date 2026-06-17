package security_governance_v3

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	hcty "github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceChaosSecurityGovernanceRuleV3() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for managing a Harness Chaos Security Governance Rule (V3 / REST API).",
		CreateContext: resourceChaosSecurityGovernanceRuleV3Create,
		ReadContext:   resourceChaosSecurityGovernanceRuleV3Read,
		UpdateContext: resourceChaosSecurityGovernanceRuleV3Update,
		DeleteContext: resourceChaosSecurityGovernanceRuleV3Delete,
		Importer:      &schema.ResourceImporter{StateContext: resourceChaosSecurityGovernanceRuleV3Import},

		Schema: resourceChaosSecurityGovernanceRuleV3Schema(),
	}
}

func resourceChaosSecurityGovernanceRuleV3Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"org_id": {
			Description:  "The organization ID of the security governance rule",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"project_id": {
			Description:  "The project ID of the security governance rule",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"name": {
			Description:  "Name of the security governance rule",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"description": {
			Description: "Description of the security governance rule",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Description: "Tags for the security governance rule",
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"is_enabled": {
			Description: "Whether the rule is enabled",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
		},
		"user_group_ids": {
			Description: "List of user group IDs associated with this rule",
			Type:        schema.TypeList,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"condition_ids": {
			Description: "List of condition IDs associated with this rule",
			Type:        schema.TypeList,
			Required:    true,
			MinItems:    1,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"time_windows": {
			Description: "Time windows during which the rule is active",
			Type:        schema.TypeList,
			Required:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"time_zone": {
						Description: "IANA time zone for the window (e.g., UTC, America/New_York).",
						Type:        schema.TypeString,
						Required:    true,
					},
					"start_time": {
						Description: "Start of the window as a Unix epoch timestamp in milliseconds.",
						Type:        schema.TypeInt,
						Required:    true,
					},
					"end_time": {
						Description: "End of the window as a Unix epoch timestamp in milliseconds. Computed from duration when not set.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"duration": {
						Description:  "Duration of the window (e.g., 30m, 1h). Computed from end_time when not set.",
						Type:         schema.TypeString,
						Optional:     true,
						Computed:     true,
						ValidateFunc: validation.StringMatch(regexp.MustCompile(`^\d+[smh]$`), "must be a valid duration (e.g., 30m, 1h)"),
					},
					"recurrence": {
						Description: "Recurrence specification for the time window.",
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"type": {
									Description: "Recurrence type (None, Daily, Weekly, Monthly, Yearly).",
									Type:        schema.TypeString,
									Required:    true,
									ValidateFunc: validation.StringInSlice([]string{
										"None",
										"Daily",
										"Weekly",
										"Monthly",
										"Yearly",
									}, false),
								},
								"until": {
									Description: "End of the recurrence as a Unix epoch timestamp in milliseconds. Use -1 for no end.",
									Type:        schema.TypeInt,
									Required:    true,
								},
								"value": {
									Description: "Day of month for Monthly recurrence. Only used when type is Monthly.",
									Type:        schema.TypeInt,
									Optional:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

// expandConditionIDsV3 extracts bare condition IDs, tolerating fully scoped
// references (org/project/condition-id) by using the trailing path segment.
func expandConditionIDsV3(d *schema.ResourceData) []string {
	var ids []string
	for _, v := range d.Get("condition_ids").([]interface{}) {
		s, ok := v.(string)
		if !ok || s == "" {
			continue
		}
		parts := strings.Split(s, "/")
		ids = append(ids, parts[len(parts)-1])
	}
	return ids
}

// rawTimeWindowAttrConfigured reports whether the given time_windows[idx].attr
// was explicitly set in the practitioner's configuration (as opposed to being a
// carried-over Computed value from prior state). This is used to keep end_time
// and duration mutually exclusive on the wire.
func rawTimeWindowAttrConfigured(d *schema.ResourceData, idx int, attr string) bool {
	v, diags := d.GetRawConfigAt(hcty.GetAttrPath("time_windows").IndexInt(idx).GetAttr(attr))
	return !diags.HasError() && v.IsKnown() && !v.IsNull()
}

func buildRuleTimeWindowsV3(d *schema.ResourceData) []chaos.SecurityGovernanceTimeWindow {
	tws := d.Get("time_windows").([]interface{})
	result := make([]chaos.SecurityGovernanceTimeWindow, 0, len(tws))
	for i, tw := range tws {
		m := tw.(map[string]interface{})
		w := chaos.SecurityGovernanceTimeWindow{
			TimeZone:  m["time_zone"].(string),
			StartTime: int64(m["start_time"].(int)),
		}

		et, _ := m["end_time"].(int)
		dur, _ := m["duration"].(string)

		// end_time and duration are mutually exclusive inputs; the backend
		// derives the unset one. Send only the value the practitioner explicitly
		// set in config so a carried-over Computed value (e.g. a prior duration)
		// cannot override a newly configured end_time and cause perpetual drift.
		durConfigured := rawTimeWindowAttrConfigured(d, i, "duration")
		endConfigured := rawTimeWindowAttrConfigured(d, i, "end_time")
		switch {
		case durConfigured && dur != "":
			w.Duration = dur
		case endConfigured && et > 0:
			w.EndTime = int64(et)
		case dur != "":
			w.Duration = dur
		case et > 0:
			w.EndTime = int64(et)
		}
		if recs, ok := m["recurrence"].([]interface{}); ok && len(recs) > 0 && recs[0] != nil {
			r := recs[0].(map[string]interface{})
			rec := &chaos.SecurityGovernanceRecurrence{
				Type_: r["type"].(string),
				Spec:  &chaos.SecurityGovernanceRecurrenceSpec{},
			}
			if until, ok := r["until"].(int); ok {
				rec.Spec.Until = int64(until)
			}
			// Value is only meaningful for Monthly recurrence.
			if r["type"].(string) == "Monthly" {
				if val, ok := r["value"].(int); ok {
					rec.Spec.Value = int32(val)
				}
			}
			w.Recurrence = rec
		}
		result = append(result, w)
	}
	return result
}

func resourceChaosSecurityGovernanceRuleV3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	ruleID := fmt.Sprintf("tf-rule-%d", time.Now().UnixNano())

	req := chaos.ChaosguardrulesCreateRuleRequest{
		RuleId:       ruleID,
		Name:         d.Get("name").(string),
		Description:  d.Get("description").(string),
		IsEnabled:    d.Get("is_enabled").(bool),
		ConditionIds: expandConditionIDsV3(d),
		UserGroupIds: interfaceSliceToStringSlice(d.Get("user_group_ids").([]interface{})),
		Tags:         interfaceSliceToStringSlice(d.Get("tags").([]interface{})),
		TimeWindows:  buildRuleTimeWindowsV3(d),
	}

	log.Printf("[DEBUG] Creating chaos security governance rule: %s (org: %s, project: %s)", ruleID, orgID, projectID)

	_, httpResp, err := c.CreateruleApi.CreateRule(ctx, req, c.AccountId, &chaos.CreateruleApiCreateRuleOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
	})
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	d.SetId(ruleID)

	return resourceChaosSecurityGovernanceRuleV3Read(ctx, d, meta)
}

func resourceChaosSecurityGovernanceRuleV3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	ruleID := d.Id()

	log.Printf("[DEBUG] Reading chaos security governance rule: %s", ruleID)

	resp, httpResp, err := c.GetruleApi.GetRule(ctx, c.AccountId, ruleID, &chaos.GetruleApiGetRuleOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
	})
	if err != nil {
		return helpers.HandleChaosReadApiErrorWithGracefulDestroy(err, d, httpResp, []string{
			"not found",
			"no documents in result",
		})
	}

	return setRuleV3Data(d, &resp, orgID, projectID)
}

func resourceChaosSecurityGovernanceRuleV3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	ruleID := d.Id()

	body := chaos.GithubComHarnessHceSaasGraphqlServerPkgDatabaseMongodbSecurityGovernanceRule{
		RuleId:       ruleID,
		Name:         d.Get("name").(string),
		Description:  d.Get("description").(string),
		IsEnabled:    d.Get("is_enabled").(bool),
		ConditionIds: expandConditionIDsV3(d),
		UserGroupIds: interfaceSliceToStringSlice(d.Get("user_group_ids").([]interface{})),
		Tags:         interfaceSliceToStringSlice(d.Get("tags").([]interface{})),
		TimeWindows:  buildRuleTimeWindowsV3(d),
	}

	log.Printf("[DEBUG] Updating chaos security governance rule: %s", ruleID)

	_, httpResp, err := c.UpdateruleApi.UpdateRule(ctx, body, c.AccountId, ruleID, &chaos.UpdateruleApiUpdateRuleOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
	})
	if err != nil {
		return helpers.HandleChaosApiError(err, d, httpResp)
	}

	return resourceChaosSecurityGovernanceRuleV3Read(ctx, d, meta)
}

func resourceChaosSecurityGovernanceRuleV3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	ruleID := d.Id()

	log.Printf("[DEBUG] Deleting chaos security governance rule: %s", ruleID)

	_, httpResp, err := c.DeleteruleApi.DeleteRule(ctx, c.AccountId, ruleID, &chaos.DeleteruleApiDeleteRuleOpts{
		OrganizationIdentifier: optional.NewString(orgID),
		ProjectIdentifier:      optional.NewString(projectID),
	})
	if err != nil {
		diags := helpers.HandleChaosReadApiErrorWithGracefulDestroy(err, d, httpResp, []string{
			"not found",
			"no documents in result",
		})
		if d.Id() == "" {
			return diags
		}
		return diags
	}

	d.SetId("")
	return nil
}

func resourceChaosSecurityGovernanceRuleV3Import(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// Expected format: <org_id>/<project_id>/<rule_id>
	importID := d.Id()
	parts := strings.Split(importID, "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid import ID format. Expected \"<org-id>/<project-id>/<rule-id>\", got: %s", importID)
	}

	orgID := parts[0]
	projectID := parts[1]
	ruleID := parts[2]

	if orgID == "" || projectID == "" || ruleID == "" {
		return nil, fmt.Errorf("org_id, project_id, and rule_id cannot be empty")
	}

	d.SetId(ruleID)
	if err := d.Set("org_id", orgID); err != nil {
		return nil, fmt.Errorf("failed to set org_id: %w", err)
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, fmt.Errorf("failed to set project_id: %w", err)
	}

	diags := resourceChaosSecurityGovernanceRuleV3Read(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read security governance rule during import: %v", diags)
	}

	return []*schema.ResourceData{d}, nil
}

// setRuleV3Data maps a GetRule REST response onto the Terraform state.
func setRuleV3Data(d *schema.ResourceData, resp *chaos.ChaosguardrulesGetRuleResponse, orgID, projectID string) diag.Diagnostics {
	d.Set("org_id", orgID)
	d.Set("project_id", projectID)
	d.Set("name", resp.Name)
	d.Set("description", resp.Description)
	d.Set("is_enabled", resp.IsEnabled)

	if len(resp.UserGroupIds) > 0 {
		d.Set("user_group_ids", resp.UserGroupIds)
	}
	if len(resp.Tags) > 0 {
		d.Set("tags", resp.Tags)
	}
	if len(resp.ConditionIds) > 0 {
		d.Set("condition_ids", resp.ConditionIds)
	}

	if len(resp.TimeWindows) > 0 {
		tws := make([]map[string]interface{}, 0, len(resp.TimeWindows))
		for _, tw := range resp.TimeWindows {
			m := map[string]interface{}{
				"time_zone":  tw.TimeZone,
				"start_time": int(tw.StartTime),
				"end_time":   int(tw.EndTime),
				"duration":   tw.Duration,
			}
			if tw.Recurrence != nil {
				rec := map[string]interface{}{
					"type": tw.Recurrence.Type_,
				}
				if tw.Recurrence.Spec != nil {
					rec["until"] = int(tw.Recurrence.Spec.Until)
					rec["value"] = int(tw.Recurrence.Spec.Value)
				}
				m["recurrence"] = []interface{}{rec}
			}
			tws = append(tws, m)
		}
		if err := d.Set("time_windows", tws); err != nil {
			return diag.FromErr(fmt.Errorf("failed to set time_windows: %w", err))
		}
	}

	return nil
}
