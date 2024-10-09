package governance_rule

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceRule() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating, updating, and managing rule.",
		ReadContext:   resourceRuleRead,
		CreateContext: resourceRuleCreateOrUpdate,
		UpdateContext: resourceRuleCreateOrUpdate,
		DeleteContext: resourceRuleDelete,
		Importer:      helpers.AccountLevelResourceImporter,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the rule.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description for rule.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cloud_provider": {
				Description:  "The cloud provider for the rule. It should be either AWS, AZURE or GCP.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"AWS", "GCP", "AZURE"}, false),
			},
			"rules_yaml": {
				Description: "The policy YAML of the rule",
				Type:        schema.TypeString,
				Required:    true,
			},
			"rule_id": {
				Description: "Id of the rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}

func resourceRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	resp, httpResp, err := c.RuleApi.GetPolicies(ctx, readRuleRequest(id), c.AccountId, nil)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Data != nil {
		err := readRuleResponse(d, resp.Data)
		if err != nil {
			return helpers.HandleReadApiError(err, d, httpResp)
		}
	}

	return nil
}

func readRuleRequest(id string) nextgen.ListDto {
	return nextgen.ListDto{
		Query: &nextgen.RuleRequest{
			PolicyIds: []string{id},
		},
	}
}

func readRuleResponse(d *schema.ResourceData, ruleList *nextgen.RuleList) error {
	rule := ruleList.Rules[0]

	d.Set("name", rule.Name)
	d.Set("cloud_provider", rule.CloudProvider)
	d.Set("description", rule.Description)
	d.Set("rules_yaml", rule.RulesYaml)

	return nil
}

func resourceRuleCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoRule
	var httpResp *http.Response

	id := d.Id()

	if id == "" {
		resp, httpResp, err = c.RuleApi.CreateNewRule(ctx, buildRule(d, false), c.AccountId)
	} else {
		resp, httpResp, err = c.RuleApi.UpdateRule(ctx, buildRule(d, true), c.AccountId)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data != nil {
		createOrUpdateRuleResponse(d, resp.Data)
	}

	return nil
}

func buildRule(d *schema.ResourceData, update bool) nextgen.CreateRuleDto {
	rule := &nextgen.CcmRule{
		Name:          d.Get("name").(string),
		CloudProvider: d.Get("cloud_provider").(string),
		Description:   d.Get("description").(string),
		RulesYaml:     d.Get("rules_yaml").(string),
		IsOOTB:        false,
	}

	if update {
		rule.Uuid = d.Id()
	}

	return nextgen.CreateRuleDto{
		Rule: rule,
	}
}

func createOrUpdateRuleResponse(d *schema.ResourceData, rule *nextgen.CcmRule) error {
	d.SetId(rule.Uuid)
	d.Set("rule_id", rule.Uuid)
	d.Set("name", rule.Name)
	d.Set("cloud_provider", rule.CloudProvider)
	d.Set("description", rule.Description)
	d.Set("rules_yaml", rule.RulesYaml)

	return nil
}

func resourceRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	_, httpResp, err := c.RuleApi.DeleteRule(ctx, c.AccountId, id)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}
