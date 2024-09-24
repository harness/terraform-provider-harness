package governance_rule

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceRule() *schema.Resource {
	return &schema.Resource{
		Description: "Datasource for looking up a rule.",

		ReadContext: resourceRuleReadDataSource,

		Schema: map[string]*schema.Schema{
			"rule_id": {
				Description: "Id of rule.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"rules_yaml": {
				Description: "Policy YAML of the rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cloud_provider": {
				Description: "The cloud provider for the rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description for rule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceRuleReadDataSource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("rule_id").(string)
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

	d.SetId(id)

	return nil
}
