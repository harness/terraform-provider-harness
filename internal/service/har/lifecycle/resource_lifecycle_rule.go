package lifecycle

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/har"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceLifecycleRule() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating and managing Harness Artifact Registry Lifecycle Rules.",
		ReadContext:   resourceLifecycleRuleRead,
		CreateContext: resourceLifecycleRuleCreate,
		UpdateContext: resourceLifecycleRuleUpdate,
		DeleteContext: resourceLifecycleRuleDelete,
		Schema:        resourceLifecycleRuleSchema(false),
		Importer: &schema.ResourceImporter{
			StateContext: resourceLifecycleRuleImport,
		},
	}
}

func resourceLifecycleRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetHarClientWithContext(ctx)

	accountId := d.Get("account_id").(string)
	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)
	ruleId := d.Id()

	resp, httpResp, err := c.LifecycleApi.GetLifecycleRule(ctx, accountId, orgId, projectId, ruleId)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readLifecycleRule(d, &resp)
	return nil
}

func resourceLifecycleRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetHarClientWithContext(ctx)

	accountId := d.Get("account_id").(string)
	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)

	body := buildLifecycleRuleRequest(d)

	resp, httpResp, err := c.LifecycleApi.CreateLifecycleRule(ctx, accountId, orgId, projectId, body)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readLifecycleRule(d, &resp)
	return nil
}

func resourceLifecycleRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetHarClientWithContext(ctx)

	accountId := d.Get("account_id").(string)
	orgId := d.Get("org_id").(string)
	projectId := d.Get("project_id").(string)
	ruleId := d.Id()

	body := buildLifecycleRuleRequest(d)

	resp, httpResp, err := c.LifecycleApi.UpdateLifecycleRule(ctx, accountId, orgId, projectId, ruleId, body)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readLifecycleRule(d, &resp)
	return nil
}

func resourceLifecycleRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetHarClientWithContext(ctx)

	accountId := d.Get("account_id").(string)
	ruleId := d.Id()

	httpResp, err := c.LifecycleApi.DeleteLifecycleRule(ctx, accountId, ruleId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func resourceLifecycleRuleImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()
	if id == "" {
		return nil, fmt.Errorf("import ID cannot be empty")
	}

	// Format: accountId/ruleId or accountId/orgId/ruleId or accountId/orgId/projectId/ruleId
	parts := strings.Split(id, "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid import ID format. Expected: accountId/ruleId or accountId/orgId/ruleId or accountId/orgId/projectId/ruleId")
	}

	ruleId := parts[len(parts)-1]
	accountId := parts[0]

	d.Set("account_id", accountId)
	d.Set("rule_id", ruleId)
	d.SetId(ruleId)

	switch len(parts) {
	case 3:
		d.Set("org_id", parts[1])
	case 4:
		d.Set("org_id", parts[1])
		d.Set("project_id", parts[2])
	}

	diags := resourceLifecycleRuleRead(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read lifecycle rule '%s': %v", ruleId, diags)
	}

	return []*schema.ResourceData{d}, nil
}

func buildLifecycleRuleRequest(d *schema.ResourceData) har.LifecycleRuleRequest {
	req := har.LifecycleRuleRequest{
		Name:   d.Get("name").(string),
		Action: har.LifecycleRuleAction(d.Get("action").(string)),
	}

	if v, ok := d.GetOk("description"); ok {
		s := v.(string)
		req.Description = &s
	}

	if v, ok := d.GetOk("package_type"); ok {
		s := v.(string)
		req.PackageType = &s
	}

	if v, ok := d.GetOk("apply_to"); ok {
		applyToList := v.([]interface{})
		if len(applyToList) > 0 {
			applyToMap := applyToList[0].(map[string]interface{})
			req.ApplyTo = har.LifecycleRuleApplyTo{
				Mode: har.LifecycleRuleApplyToMode(applyToMap["mode"].(string)),
			}
			if regs, ok := applyToMap["registries"].([]interface{}); ok {
				for _, r := range regs {
					req.ApplyTo.Registries = append(req.ApplyTo.Registries, r.(string))
				}
			}
		}
	}

	if v, ok := d.GetOk("criteria"); ok {
		criteriaList := v.([]interface{})
		if len(criteriaList) > 0 {
			criteriaMap := criteriaList[0].(map[string]interface{})
			criteria := &har.LifecycleRuleCriteria{
				Match: har.LifecycleRuleCriteriaMatch(criteriaMap["match"].(string)),
			}
			if rules, ok := criteriaMap["rules"].([]interface{}); ok {
				for _, r := range rules {
					ruleMap := r.(map[string]interface{})
					item := har.LifecycleRuleCriteriaItem{
						Type: har.LifecycleRuleCriteriaType(ruleMap["type"].(string)),
					}
					cfg := har.LifecycleRuleCriteriaConfig{}
					if val, ok := ruleMap["value"].(int); ok && val != 0 {
						v64 := int64(val)
						cfg.Value = &v64
					}
					if unit, ok := ruleMap["unit"].(string); ok && unit != "" {
						u := har.LifecycleRuleCriteriaUnit(unit)
						cfg.Unit = &u
					}
					item.Config = cfg
					criteria.Rules = append(criteria.Rules, item)
				}
			}
			req.Criteria = criteria
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		schedList := v.([]interface{})
		if len(schedList) > 0 {
			schedMap := schedList[0].(map[string]interface{})
			req.Schedule = &har.LifecycleRuleSchedule{
				Expression: schedMap["expression"].(string),
				Timezone:   schedMap["timezone"].(string),
			}
		}
	}

	if v, ok := d.GetOk("filter_config"); ok {
		filterList := v.([]interface{})
		if len(filterList) > 0 {
			filterMap := filterList[0].(map[string]interface{})
			pkgType := filterMap["package_type"].(string)
			req.FilterConfig = buildFilterConfig(pkgType, filterMap)
		}
	}

	return req
}

func buildFilterConfig(pkgType string, m map[string]interface{}) interface{} {
	toStringSlice := func(key string) []string {
		raw, ok := m[key].([]interface{})
		if !ok {
			return nil
		}
		result := make([]string, 0, len(raw))
		for _, v := range raw {
			result = append(result, v.(string))
		}
		return result
	}

	switch pkgType {
	case "DOCKER":
		return har.LifecycleRuleFilterConfigDocker{
			PackageType:               "DOCKER",
			PackageNameAllowedPattern: toStringSlice("package_name_allowed_pattern"),
			TagNameAllowedPattern:     toStringSlice("tag_name_allowed_pattern"),
		}
	case "MAVEN":
		return har.LifecycleRuleFilterConfigMaven{
			PackageType:               "MAVEN",
			GroupIdAllowedPattern:     toStringSlice("group_id_allowed_pattern"),
			PackageNameAllowedPattern: toStringSlice("package_name_allowed_pattern"),
			VersionNameAllowedPattern: toStringSlice("version_name_allowed_pattern"),
		}
	case "HUGGINGFACE":
		return har.LifecycleRuleFilterConfigHuggingFace{
			PackageType:               "HUGGINGFACE",
			ModelAllowedPattern:       toStringSlice("model_allowed_pattern"),
			DatasetAllowedPattern:     toStringSlice("dataset_allowed_pattern"),
			VersionNameAllowedPattern: toStringSlice("version_name_allowed_pattern"),
		}
	default:
		return har.LifecycleRuleFilterConfigGeneric{
			PackageType:               pkgType,
			PackageNameAllowedPattern: toStringSlice("package_name_allowed_pattern"),
			VersionNameAllowedPattern: toStringSlice("version_name_allowed_pattern"),
		}
	}
}

func readLifecycleRule(d *schema.ResourceData, rule *har.LifecycleRuleResponse) {
	if rule == nil {
		return
	}

	d.SetId(rule.Id)
	d.Set("rule_id", rule.Id)
	d.Set("name", rule.Name)
	d.Set("action", string(rule.Action))
	d.Set("enabled", rule.Enabled)

	if rule.Description != nil {
		d.Set("description", *rule.Description)
	}
	if rule.PackageType != nil {
		d.Set("package_type", *rule.PackageType)
	}
	if rule.OrgIdentifier != nil {
		d.Set("org_id", *rule.OrgIdentifier)
	}
	if rule.ProjectIdentifier != nil {
		d.Set("project_id", *rule.ProjectIdentifier)
	}
	if rule.CreatedAt != nil {
		d.Set("created_at", int(*rule.CreatedAt))
	}
	if rule.UpdatedAt != nil {
		d.Set("updated_at", int(*rule.UpdatedAt))
	}
	if rule.LastRunAt != nil {
		d.Set("last_run_at", int(*rule.LastRunAt))
	}
	if rule.NextRunAt != nil {
		d.Set("next_run_at", int(*rule.NextRunAt))
	}

	applyToMap := map[string]interface{}{
		"mode":       string(rule.ApplyTo.Mode),
		"registries": rule.ApplyTo.Registries,
	}
	d.Set("apply_to", []interface{}{applyToMap})

	if rule.Criteria != nil {
		criteriaRules := make([]interface{}, 0, len(rule.Criteria.Rules))
		for _, r := range rule.Criteria.Rules {
			ruleMap := map[string]interface{}{
				"type": string(r.Type),
			}
			if r.Config.Value != nil {
				ruleMap["value"] = int(*r.Config.Value)
			}
			if r.Config.Unit != nil {
				ruleMap["unit"] = string(*r.Config.Unit)
			}
			criteriaRules = append(criteriaRules, ruleMap)
		}
		criteriaMap := map[string]interface{}{
			"match": string(rule.Criteria.Match),
			"rules": criteriaRules,
		}
		d.Set("criteria", []interface{}{criteriaMap})
	}

	if rule.Schedule != nil {
		schedMap := map[string]interface{}{
			"expression": rule.Schedule.Expression,
			"timezone":   rule.Schedule.Timezone,
		}
		d.Set("schedule", []interface{}{schedMap})
	}
}
