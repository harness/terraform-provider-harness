package governance_rule_set

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceRuleSet() *schema.Resource {
	return &schema.Resource{
		Description: "Datasource for looking up a rule.",

		ReadContext: resourceRuleSetReadDataSource,

		Schema: map[string]*schema.Schema{
			"rule_set_id": {
				Description: "Id of rule set.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the rule set.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cloud_provider": {
				Description: "The cloud provider for the rule set.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description for rule set.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"rule_ids": {
				Description: "List of rule IDs.",
				Type:        schema.TypeMap,
				Computed:    true,
			},
		},
	}
}

func resourceRuleSetReadDataSource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("rule_set_id").(string)
	resp, httpResp, err := c.RuleSetsApi.ListRuleSets(ctx, readRuleSetRequest(id), c.AccountId)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Data != nil {
		err := readRuleSetResponse(d, resp.Data)
		if err != nil {
			return helpers.HandleReadApiError(err, d, httpResp)
		}
	}

	d.SetId(id)

	return nil
}
