package security_governance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	ruleResourceName = "harness_chaos_security_governance_rule"
)

func ResourceChaosSecurityGovernanceRule() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for managing a Harness Chaos Security Governance Rule",
		CreateContext: resourceChaosSecurityGovernanceRuleCreate,
		ReadContext:   resourceChaosSecurityGovernanceRuleRead,
		UpdateContext: resourceChaosSecurityGovernanceRuleUpdate,
		DeleteContext: resourceChaosSecurityGovernanceRuleDelete,
		Importer:      &schema.ResourceImporter{StateContext: resourceChaosSecurityGovernanceRuleImport},

		CustomizeDiff: validateTimeWindows,
		Schema: map[string]*schema.Schema{
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
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"condition_ids": {
				Description: "List of condition IDs associated with this rule",
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_windows": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time_zone": {
							Type:     schema.TypeString,
							Required: true,
						},
						"start_time": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"end_time": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							// ValidateFunc: validation.IntAtLeast(1),
						},
						"duration": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringMatch(regexp.MustCompile(`^\d+[smh]$`), "must be a valid duration (e.g., 30m, 1h)"),
						},
						"recurrence": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice([]string{
											"None",
											"Daily",
											"Weekly",
											"Monthly",
											"Yearly",
										}, false),
									},
									"until": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"value": {
										Type:     schema.TypeInt,
										Optional: true,
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

// Custom validation function for time windows
func validateTimeWindows(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	// if d.HasChange("time_windows") {
	// 	timeWindows := d.Get("time_windows").([]interface{})

	// 	for i, tw := range timeWindows {
	// 		window := tw.(map[string]interface{})
	// 		hasEndTime := window["end_time"] != nil && window["end_time"].(int) > 0
	// 		hasDuration := window["duration"] != nil && window["duration"].(string) != ""

	// 		if hasEndTime && hasDuration {
	// 			return fmt.Errorf("time_windows[%d]: only one of 'end_time' or 'duration' can be specified", i)
	// 		}

	// 		if !hasEndTime && !hasDuration {
	// 			return fmt.Errorf("time_windows[%d]: one of 'end_time' or 'duration' must be specified", i)
	// 		}
	// 	}
	// }

	return nil
}

// Add this helper function at the top of the file
func validateResourceData(d *schema.ResourceData) error {
	if d == nil {
		return fmt.Errorf("resource data is nil")
	}

	if _, ok := d.GetOk("name"); !ok {
		return fmt.Errorf("name is required")
	}

	if _, ok := d.GetOk("condition_ids"); !ok {
		return fmt.Errorf("at least one condition_id is required")
	}

	return nil
}

// Update the buildTimeWindows function
func buildTimeWindows(d *schema.ResourceData) ([]*model.TimeWindowInput, error) {
	timeWindows := d.Get("time_windows").([]interface{})
	if len(timeWindows) == 0 {
		return nil, fmt.Errorf("at least one time window is required")
	}

	result := make([]*model.TimeWindowInput, 0, len(timeWindows))

	for i, tw := range timeWindows {
		twMap, ok := tw.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid time window format at index %d", i)
		}

		timeZone, ok := twMap["time_zone"].(string)
		if !ok || timeZone == "" {
			return nil, fmt.Errorf("time_zone is required for time window %d", i)
		}

		startTime, ok := twMap["start_time"].(int)
		if !ok {
			return nil, fmt.Errorf("start_time is required for time window %d", i)
		}

		timeWindow := &model.TimeWindowInput{
			TimeZone:  timeZone,
			StartTime: startTime,
		}

		// Handle end_time or duration
		if endTime, ok := twMap["end_time"].(int); ok && endTime > 0 {
			timeWindow.EndTime = &endTime
		} else if duration, ok := twMap["duration"].(string); ok && duration != "" {
			timeWindow.Duration = &duration
		} else {
			return nil, fmt.Errorf("either end_time or duration is required for time window %d", i)
		}

		// Handle recurrence
		if recurrences, ok := twMap["recurrence"].([]interface{}); ok && len(recurrences) > 0 {
			rec, ok := recurrences[0].(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid recurrence format in time window %d", i)
			}

			until, ok := rec["until"].(int)
			if !ok {
				return nil, fmt.Errorf("until is required for recurrence in time window %d", i)
			}

			recurrence := &model.RecurrenceInput{
				Type: model.RecurrenceType(rec["type"].(string)),
				Spec: &model.RecurrenceSpecInput{
					Until: &until,
				},
			}

			// Only set Value for MONTHLY recurrence
			if rec["type"].(string) == "Monthly" {
				if value, ok := rec["value"].(int); ok {
					valueInt := int(value)
					recurrence.Spec.Value = &valueInt
				}
			}

			timeWindow.Recurrence = recurrence
		}

		result = append(result, timeWindow)
	}

	return result, nil
}

// Update the buildRuleInput function
func buildRuleInput(d *schema.ResourceData) (*model.RuleInput, error) {
	if d == nil {
		return nil, fmt.Errorf("resource data is nil")
	}

	// Parse condition IDs
	var conditionIDs []string
	if v, ok := d.GetOk("condition_ids"); ok {
		for _, id := range v.([]interface{}) {
			if id == nil {
				continue
			}
			idStr := id.(string)
			if idStr == "" {
				continue
			}
			parts := strings.Split(idStr, "/")
			conditionIDs = append(conditionIDs, parts[len(parts)-1])
		}
	}

	if len(conditionIDs) == 0 {
		return nil, fmt.Errorf("at least one condition_id is required")
	}

	// Parse time windows
	timeWindows, err := buildTimeWindows(d)
	if err != nil {
		return nil, fmt.Errorf("failed to build time windows: %w", err)
	}

	// Parse tags
	var tags []*string
	if v, ok := d.GetOk("tags"); ok {
		for _, tag := range v.([]interface{}) {
			if tag == nil {
				continue
			}
			tagStr := tag.(string)
			if tagStr != "" {
				tags = append(tags, &tagStr)
			}
		}
	}

	// Parse user group IDs
	var userGroupIDs []string
	if v, ok := d.GetOk("user_group_ids"); ok {
		for _, id := range v.([]interface{}) {
			if id == nil {
				continue
			}
			if idStr, ok := id.(string); ok && idStr != "" {
				userGroupIDs = append(userGroupIDs, idStr)
			}
		}
	}

	// Get description
	var description *string
	if v, ok := d.GetOk("description"); ok && v != "" {
		desc := v.(string)
		description = &desc
	}

	return &model.RuleInput{
		Name:         d.Get("name").(string),
		Description:  description,
		IsEnabled:    d.Get("is_enabled").(bool),
		RuleID:       d.Id(),
		UserGroupIds: userGroupIDs,
		TimeWindows:  timeWindows,
		Tags:         tags,
		ConditionIds: conditionIDs,
	}, nil
}

func getRuleIdentifiers(d *schema.ResourceData, accountID string) model.IdentifiersRequest {
	return model.IdentifiersRequest{
		AccountIdentifier: accountID,
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
	}
}

func resourceChaosSecurityGovernanceRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient
	client := chaos.NewSecurityGovernanceRuleClient(c)

	accountID := c.AccountId
	if accountID == "" {
		return diag.Errorf("account ID must be configured in the provider")
	}

	identifiers := getRuleIdentifiers(d, accountID)
	ruleInput, err := buildRuleInput(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to build rule input: %w", err))
	}

	// Generate a unique rule ID
	ruleID := fmt.Sprintf("tf-rule-%d", time.Now().UnixNano())
	ruleInput.RuleID = ruleID

	log.Printf("[DEBUG] Creating rule with input: %+v", ruleInput)

	resp, err := client.Create(ctx, identifiers, *ruleInput)
	if err != nil {
		return diag.Errorf("failed to create security governance rule (account: %s, org: %s, project: %s): %v",
			identifiers.AccountIdentifier,
			identifiers.OrgIdentifier,
			identifiers.ProjectIdentifier,
			err,
		)
	}

	log.Printf("[DEBUG] Created rule with response: %+v", resp)

	// The response contains the rule in the Response field as a JSON string
	var ruleIDStr string
	if resp.Response != "" {
		// Parse the response JSON to get the rule ID
		var responseObj struct {
			Rule struct {
				RuleID string `json:"ruleId"`
			} `json:"rule"`
		}
		if err := json.Unmarshal([]byte(resp.Response), &responseObj); err == nil && responseObj.Rule.RuleID != "" {
			ruleIDStr = responseObj.Rule.RuleID
		}
	}

	if ruleIDStr == "" {
		// Fallback to the generated rule ID if we couldn't extract it from the response
		ruleIDStr = ruleID
		log.Printf("[WARN] Could not extract rule ID from response, using generated ID: %s", ruleIDStr)
	}

	d.SetId(ruleIDStr)

	// Read the rule to populate all fields
	return resourceChaosSecurityGovernanceRuleRead(ctx, d, meta)
}

func resourceChaosSecurityGovernanceRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient
	client := chaos.NewSecurityGovernanceRuleClient(c)

	ruleID := d.Id()
	accountID := c.AccountId
	if accountID == "" {
		err := "account ID must be configured in the provider"
		log.Printf("[ERROR] %s", err)
		return diag.Errorf(err)
	}
	identifiers := getRuleIdentifiers(d, accountID)

	log.Printf("[DEBUG] Reading rule with ID: %s, Account: %s, Org: %v, Project: %v",
		ruleID, accountID, identifiers.OrgIdentifier, identifiers.ProjectIdentifier)

	resp, err := client.Get(ctx, identifiers, ruleID)
	if err != nil {
		// If the rule is not found, it might have been deleted outside of Terraform
		if strings.Contains(err.Error(), "no documents in result") ||
			strings.Contains(err.Error(), "not found") {
			log.Printf("[WARN] Rule not found, removing from state: %s", d.Id())
			d.SetId("")
			return nil
		}
		return diag.Errorf("failed to read security governance rule: %v", err)
	}

	if resp == nil || resp.Rule == nil {
		log.Printf("[WARN] Rule not found, removing from state: %s", d.Id())
		d.SetId("")
		return nil
	}

	// Set the attributes from the response
	if err := setRuleAttributes(d, resp, c.AccountId); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set rule attributes: %w", err))
	}

	return nil
}

// Update the setRuleAttributes function
func setRuleAttributes(d *schema.ResourceData, resp *model.RuleResponse, accountID string) error {
	if resp == nil || resp.Rule == nil {
		return fmt.Errorf("nil rule in response")
	}

	rule := resp.Rule

	// Set basic fields
	if err := d.Set("name", rule.Name); err != nil {
		return fmt.Errorf("failed to set name: %w", err)
	}

	if rule.Description != nil {
		if err := d.Set("description", *rule.Description); err != nil {
			return fmt.Errorf("failed to set description: %w", err)
		}
	}

	if err := d.Set("is_enabled", rule.IsEnabled); err != nil {
		return fmt.Errorf("failed to set is_enabled: %w", err)
	}

	// Set user group IDs
	if len(rule.UserGroupIds) > 0 {
		if err := d.Set("user_group_ids", rule.UserGroupIds); err != nil {
			return fmt.Errorf("failed to set user_group_ids: %w", err)
		}
	}

	// Set tags
	if len(rule.Tags) > 0 {
		tags := make([]string, 0, len(rule.Tags))
		for _, tag := range rule.Tags {
			if len(tag) > 0 && tag != "" {
				tags = append(tags, tag)
			}
		}
		if err := d.Set("tags", tags); err != nil {
			return fmt.Errorf("failed to set tags: %w", err)
		}
	}

	// Set condition IDs
	if len(rule.Conditions) > 0 {
		conditionIDs := make([]string, 0, len(rule.Conditions))

		for _, condition := range rule.Conditions {
			if condition != nil && condition.ConditionID != "" {
				conditionID := condition.ConditionID
				conditionIDs = append(conditionIDs, conditionID)
			}
		}
		if err := d.Set("condition_ids", conditionIDs); err != nil {
			return fmt.Errorf("failed to set condition_ids: %w", err)
		}
	}

	// Set time windows
	if len(rule.TimeWindows) > 0 {
		timeWindows := make([]map[string]interface{}, 0, len(rule.TimeWindows))
		for _, tw := range rule.TimeWindows {
			if tw == nil {
				continue
			}

			timeWindow := map[string]interface{}{
				"time_zone":  tw.TimeZone,
				"start_time": tw.StartTime,
				"end_time":   nil,
				"duration":   nil,
				"recurrence": nil,
			}

			if tw.EndTime != nil {
				timeWindow["end_time"] = *tw.EndTime
			}
			if tw.Duration != nil {
				timeWindow["duration"] = *tw.Duration
			}

			if tw.Recurrence != nil {
				recurrence := map[string]interface{}{
					"type": tw.Recurrence.Type,
				}

				if tw.Recurrence.Spec != nil {
					if tw.Recurrence.Spec.Until != nil {
						recurrence["until"] = *tw.Recurrence.Spec.Until
					}
					if tw.Recurrence.Spec.Value != nil {
						recurrence["value"] = *tw.Recurrence.Spec.Value
					}
				}

				timeWindow["recurrence"] = []map[string]interface{}{recurrence}
			}

			timeWindows = append(timeWindows, timeWindow)
		}
		if err := d.Set("time_windows", timeWindows); err != nil {
			return fmt.Errorf("failed to set time_windows: %w", err)
		}
	}

	return nil
}

func resourceChaosSecurityGovernanceRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient
	client := chaos.NewSecurityGovernanceRuleClient(c)

	ruleID := d.Id()
	accountID := c.AccountId
	if accountID == "" {
		err := "account ID must be configured in the provider"
		log.Printf("[ERROR] %s", err)
		return diag.Errorf(err)
	}
	identifiers := getRuleIdentifiers(d, accountID)

	ruleInput, err := buildRuleInput(d)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to build rule input: %w", err))
	}

	ruleInput.RuleID = ruleID

	log.Printf("[DEBUG] Updating rule with input: %+v", ruleInput)

	_, err = client.Update(ctx, identifiers, *ruleInput)
	if err != nil {
		return diag.Errorf("failed to update security governance rule: %v", err)
	}

	return resourceChaosSecurityGovernanceRuleRead(ctx, d, meta)
}

func resourceChaosSecurityGovernanceRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient
	client := chaos.NewSecurityGovernanceRuleClient(c)

	ruleID := d.Id()
	accountID := c.AccountId
	if accountID == "" {
		err := "account ID must be configured in the provider"
		log.Printf("[ERROR] %s", err)
		return diag.Errorf(err)
	}
	identifiers := getRuleIdentifiers(d, accountID)

	log.Printf("[DEBUG] Deleting rule with ID: %s, Account: %s, Org: %v, Project: %v",
		ruleID, identifiers.AccountIdentifier, identifiers.OrgIdentifier, identifiers.ProjectIdentifier)

	_, err := client.Delete(ctx, identifiers, ruleID)
	if err != nil {
		return diag.Errorf("failed to delete security governance rule: %v", err)
	}

	d.SetId("")
	return nil
}

func resourceChaosSecurityGovernanceRuleImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// Parse the import ID which can be in one of these formats:
	// 1. Account level: "rule_id"
	// 2. Org level: "org_id/rule_id"
	// 3. Project level: "org_id/project_id/rule_id"
	importID := d.Id()
	parts := strings.Split(importID, "/")

	var ruleID, orgID, projectID string

	switch len(parts) {
	case 1:
		// Account level: "rule_id"
		ruleID = parts[0]
	case 2:
		// Org level: "org-id/rule_id"
		orgID = parts[0]
		ruleID = parts[1]
	case 3:
		// Project level: "org-id/project-id/rule_id"
		orgID = parts[0]
		projectID = parts[1]
		ruleID = parts[2]
	default:
		return nil, fmt.Errorf("invalid import ID format. Expected \"<rule-id>\", \"<org-id>/<rule-id>\", or \"<org-id>/<project-id>/<rule-id>\"")
	}

	if ruleID == "" {
		return nil, fmt.Errorf("rule id cannot be empty")
	}
	d.SetId(ruleID)

	// Set the required fields for the resource
	if err := d.Set("org_id", orgID); err != nil {
		return nil, fmt.Errorf("failed to set org_id: %w", err)
	}
	if err := d.Set("project_id", projectID); err != nil {
		return nil, fmt.Errorf("failed to set project_id: %w", err)
	}

	// Read the rule to populate all other fields
	client := meta.(*internal.Session).ChaosClient
	sgClient := chaos.NewSecurityGovernanceRuleClient(client)

	identifiers := model.IdentifiersRequest{
		AccountIdentifier: client.AccountId,
		OrgIdentifier:     orgID,
		ProjectIdentifier: projectID,
	}

	resp, err := sgClient.Get(ctx, identifiers, ruleID)
	if err != nil {
		return nil, fmt.Errorf("failed to read security governance rule during import: %w", err)
	}

	if resp == nil || resp.Rule == nil {
		return nil, fmt.Errorf("imported security governance rule not found: %s", ruleID)
	}

	// Set all other attributes from the API response
	if err := setRuleAttributes(d, resp, client.AccountId); err != nil {
		return nil, fmt.Errorf("failed to set rule attributes during import: %w", err)
	}

	return []*schema.ResourceData{d}, nil
}

func parseRuleID(id string) (model.IdentifiersRequest, string, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 3 {
		return model.IdentifiersRequest{}, "", fmt.Errorf("invalid rule ID format: %s, expected format: org_id/project_id/rule_id", id)
	}

	return model.IdentifiersRequest{
		OrgIdentifier:     parts[0],
		ProjectIdentifier: parts[1],
	}, parts[2], nil
}
