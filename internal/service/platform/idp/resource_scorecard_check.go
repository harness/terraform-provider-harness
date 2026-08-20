package idp

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func scorecardCheckSchema(dataSource bool) map[string]*schema.Schema {
	computed := dataSource
	return map[string]*schema.Schema{
		"identifier": helpers.GetIdentifierSchema(helpers.SchemaFlagTypes.Required),
		"name": func() *schema.Schema {
			if dataSource {
				return helpers.GetNameSchema(helpers.SchemaFlagTypes.Computed)
			}
			return helpers.GetNameSchema(helpers.SchemaFlagTypes.Required)
		}(),
		"description": helpers.GetDescriptionSchema(func() helpers.SchemaFlagType {
			if dataSource {
				return helpers.SchemaFlagTypes.Computed
			}
			return helpers.SchemaFlagTypes.Optional
		}()),
		"tags": helpers.GetTagsSchema(func() helpers.SchemaFlagType {
			if dataSource {
				return helpers.SchemaFlagTypes.Computed
			}
			return helpers.SchemaFlagTypes.Optional
		}()),
		"expression": {
			Type:        schema.TypeString,
			Optional:    !computed,
			Computed:    true,
			Description: "JEXL expression used for advanced check rules. Required when rule_strategy is ADVANCED.",
		},
		"rule_strategy": func() *schema.Schema {
			s := &schema.Schema{
				Type:        schema.TypeString,
				Required:    !computed,
				Computed:    computed,
				Description: "How multiple rules are combined. Valid values are ALL_OF, ANY_OF, and ADVANCED.",
			}
			if !computed {
				s.ValidateFunc = validation.StringInSlice([]string{"ALL_OF", "ANY_OF", "ADVANCED"}, false)
			}
			return s
		}(),
		"default_behaviour": func() *schema.Schema {
			s := &schema.Schema{
				Type:        schema.TypeString,
				Required:    !computed,
				Computed:    computed,
				Description: "Default behaviour when the check cannot be evaluated. Valid values are PASS and FAIL.",
			}
			if !computed {
				s.ValidateFunc = validation.StringInSlice([]string{"PASS", "FAIL"}, false)
			}
			return s
		}(),
		"rule_description": {
			Type:        schema.TypeString,
			Optional:    !computed,
			Computed:    true,
			Description: "Description of the check rule set.",
		},
		"fail_message": {
			Type:        schema.TypeString,
			Optional:    !computed,
			Computed:    true,
			Description: "Message shown when the check fails.",
		},
		"custom": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether this is a custom check. Always true for checks created through Terraform.",
		},
		"harness_managed": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the check is managed by Harness.",
		},
		"percentage": {
			Type:        schema.TypeFloat,
			Computed:    true,
			Description: "Pass percentage for the check.",
		},
		"rules": {
			Type:        schema.TypeList,
			Optional:    !computed,
			Computed:    true,
			Description: "Basic rules evaluated by the check. Required when rule_strategy is ALL_OF or ANY_OF.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"identifier": {
						Type:        schema.TypeString,
						Optional:    !computed,
						Computed:    true,
						Description: "Identifier of the rule.",
					},
					"data_source_identifier": {
						Type:        schema.TypeString,
						Required:    !computed,
						Computed:    computed,
						Description: "Identifier of the data source used by the rule.",
					},
					"data_point_identifier": {
						Type:        schema.TypeString,
						Required:    !computed,
						Computed:    computed,
						Description: "Identifier of the data point evaluated by the rule.",
					},
					"operator": {
						Type:        schema.TypeString,
						Required:    !computed,
						Computed:    computed,
						Description: "Comparison operator.",
					},
					"value": {
						Type:        schema.TypeString,
						Optional:    !computed,
						Computed:    true,
						Description: "Value to compare against.",
					},
					"rule_description": {
						Type:        schema.TypeString,
						Optional:    !computed,
						Computed:    true,
						Description: "Description of the rule.",
					},
					"input_values": {
						Type:        schema.TypeList,
						Optional:    !computed,
						Computed:    true,
						Description: "Input values passed to the data point.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"key": {
									Type:        schema.TypeString,
									Required:    !computed,
									Computed:    computed,
									Description: "Input key.",
								},
								"value": {
									Type:        schema.TypeString,
									Required:    !computed,
									Computed:    computed,
									Description: "Input value.",
								},
							},
						},
					},
				},
			},
		},
	}
}

func ResourceScorecardCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating IDP scorecard checks.",
		ReadContext:   resourceScorecardCheckRead,
		CreateContext: resourceScorecardCheckCreateOrUpdate,
		UpdateContext: resourceScorecardCheckCreateOrUpdate,
		DeleteContext: resourceScorecardCheckDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema:        scorecardCheckSchema(false),
		CustomizeDiff: resourceScorecardCheckCustomizeDiff,
	}
}

func resourceScorecardCheckCustomizeDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	return validateScorecardCheckConfig(d.Get("rule_strategy").(string), d.Get("expression").(string), d.Get("rules"))
}

func validateScorecardCheckConfig(ruleStrategy, expression string, rules interface{}) error {
	switch ruleStrategy {
	case "ALL_OF", "ANY_OF":
		if countCheckRules(rules) == 0 {
			return fmt.Errorf("rules is required when rule_strategy is %s", ruleStrategy)
		}
	case "ADVANCED":
		if strings.TrimSpace(expression) == "" {
			return fmt.Errorf("expression is required when rule_strategy is ADVANCED")
		}
	}
	return nil
}

func countCheckRules(v interface{}) int {
	items, ok := v.([]interface{})
	if !ok {
		return 0
	}
	count := 0
	for _, item := range items {
		if item != nil {
			count++
		}
	}
	return count
}

func resourceScorecardCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)
	id := d.Id()
	if id == "" {
		id = d.Get("identifier").(string)
	}

	resp, httpResp, err := c.ChecksApi.GetCheck(ctx, id, &idp.ChecksApiGetCheckOpts{
		HarnessAccount: optional.NewString(c.AccountId),
		Custom:         optional.NewBool(true),
	})
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		if isNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if err := readScorecardCheck(d, resp.CheckDetails); err != nil {
		return diag.Errorf("failed to read IDP scorecard check: %v", err)
	}
	return nil
}

func resourceScorecardCheckCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if err := validateScorecardCheckConfig(d.Get("rule_strategy").(string), d.Get("expression").(string), d.Get("rules")); err != nil {
		return diag.FromErr(err)
	}

	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)
	req := idp.CheckRequest{CheckDetails: expandCheckDetails(d)}

	var httpResp *http.Response
	var err error
	if d.Id() == "" {
		_, httpResp, err = c.ChecksApi.CreateCheck(ctx, req, &idp.ChecksApiCreateCheckOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
	} else {
		_, httpResp, err = c.ChecksApi.UpdateCheck(ctx, req, d.Id(), &idp.ChecksApiUpdateCheckOpts{
			HarnessAccount: optional.NewString(c.AccountId),
		})
	}
	if err != nil {
		return handleIDPWriteApiError("harness_platform_idp_scorecard_check", err, d, httpResp)
	}

	d.SetId(req.CheckDetails.Identifier)
	return resourceScorecardCheckRead(ctx, d, meta)
}

func resourceScorecardCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)
	opts := &idp.ChecksApiDeleteCheckOpts{
		HarnessAccount: optional.NewString(c.AccountId),
		ForceDelete:    optional.NewBool(false),
	}
	httpResp, err := c.ChecksApi.DeleteCheck(ctx, d.Id(), opts)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return nil
		}
		if isNotFoundError(err) {
			return nil
		}
		return handleIDPApiError(err, d, httpResp)
	}
	return nil
}

func expandCheckDetails(d *schema.ResourceData) *idp.CheckDetails {
	details := &idp.CheckDetails{
		Identifier:       d.Get("identifier").(string),
		Name:             d.Get("name").(string),
		Custom:           true,
		RuleStrategy:     d.Get("rule_strategy").(string),
		DefaultBehaviour: d.Get("default_behaviour").(string),
	}
	if v, ok := d.GetOk("description"); ok {
		details.Description = v.(string)
	}
	if v, ok := d.GetOk("expression"); ok {
		details.Expression = v.(string)
	}
	if v, ok := d.GetOk("rule_description"); ok {
		details.RuleDescription = v.(string)
	}
	if v, ok := d.GetOk("fail_message"); ok {
		details.FailMessage = v.(string)
	}
	if v, ok := d.GetOk("tags"); ok {
		details.Tags = expandStringSet(v)
	}
	if v, ok := d.GetOk("rules"); ok {
		details.Rules = expandCheckRules(v.([]interface{}))
	}
	return details
}

func expandCheckRules(items []interface{}) []idp.CheckRule {
	rules := make([]idp.CheckRule, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}
		m := item.(map[string]interface{})
		rule := idp.CheckRule{
			Identifier:           m["identifier"].(string),
			DataSourceIdentifier: m["data_source_identifier"].(string),
			DataPointIdentifier:  m["data_point_identifier"].(string),
			Operator:             m["operator"].(string),
			Value:                m["value"].(string),
			RuleDescription:      m["rule_description"].(string),
		}
		if v, ok := m["input_values"].([]interface{}); ok {
			rule.InputValues = expandInputValues(v)
		}
		rules = append(rules, rule)
	}
	return rules
}

func expandInputValues(items []interface{}) []idp.InputValue {
	values := make([]idp.InputValue, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}
		m := item.(map[string]interface{})
		values = append(values, idp.InputValue{
			Key:   m["key"].(string),
			Value: m["value"].(string),
		})
	}
	return values
}

func expandStringSet(v interface{}) []string {
	set := v.(*schema.Set).List()
	out := make([]string, 0, len(set))
	for _, item := range set {
		out = append(out, item.(string))
	}
	return out
}

func readScorecardCheck(d *schema.ResourceData, details *idp.CheckDetails) error {
	if details == nil {
		return nil
	}
	d.SetId(details.Identifier)
	d.Set("identifier", details.Identifier)
	d.Set("name", details.Name)
	d.Set("description", details.Description)
	d.Set("expression", details.Expression)
	d.Set("rule_description", details.RuleDescription)
	d.Set("rule_strategy", details.RuleStrategy)
	d.Set("default_behaviour", details.DefaultBehaviour)
	d.Set("fail_message", details.FailMessage)
	d.Set("custom", details.Custom)
	d.Set("harness_managed", details.HarnessManaged)
	d.Set("percentage", details.Percentage)
	d.Set("tags", details.Tags)

	rules := make([]map[string]interface{}, 0, len(details.Rules))
	for _, rule := range details.Rules {
		inputValues := make([]map[string]interface{}, 0, len(rule.InputValues))
		for _, input := range rule.InputValues {
			inputValues = append(inputValues, map[string]interface{}{
				"key":   input.Key,
				"value": input.Value,
			})
		}
		rules = append(rules, map[string]interface{}{
			"identifier":             rule.Identifier,
			"data_source_identifier": rule.DataSourceIdentifier,
			"data_point_identifier":  rule.DataPointIdentifier,
			"operator":               rule.Operator,
			"value":                  rule.Value,
			"rule_description":       rule.RuleDescription,
			"input_values":           inputValues,
		})
	}
	return d.Set("rules", rules)
}
