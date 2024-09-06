package governance_enforcement

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceRuleEnforcement() *schema.Resource {
	return &schema.Resource{
		Description: "Datasource for looking up a rule enforcement.",

		ReadContext: resourceRuleEnforcementReadDataSource,

		Schema: map[string]*schema.Schema{
			"enforcement_id": {
				Description: "Id of rule enforcement.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the rule enforcement.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cloud_provider": {
				Description: "The cloud provider for the rule enforcement.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"rule_ids": {
				Description: "List of rule IDs.",
				Type:        schema.TypeMap,
				Computed:    true,
			},
			"rule_set_ids": {
				Description: "List of rule set IDs.",
				Type:        schema.TypeMap,
				Computed:    true,
			},
			"execution_schedule": {
				Description: "Execution schedule in cron format.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"execution_timezone": {
				Description: "Timezone for the execution schedule.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"is_enabled": {
				Description: "Indicates if the rule enforcement is enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"target_accounts": {
				Description: "List of target accounts.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_regions": {
				Description: "List of target regions.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_dry_run": {
				Description: "Indicates if the rule enforcement is a dry run.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"description": {
				Description: "Description for rule enforcement.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceRuleEnforcementReadDataSource(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("enforcement_id").(string)
	resp, httpResp, err := c.RuleEnforcementApi.EnforcementDetails(ctx, c.AccountId, readRuleEnforcementRequest(id))

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Data != nil {
		err := readRuleEnforcementResponse(d, resp.Data)
		if err != nil {
			return helpers.HandleReadApiError(err, d, httpResp)
		}
	}

	d.SetId(id)

	return nil
}
