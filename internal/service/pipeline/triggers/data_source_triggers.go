package triggers

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceTriggers() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness trigger.",

		ReadContext: dataTriggersRead,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "identifier of the trigger.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"target_id": {
				Description: "Identifier of the target pipeline",
				Type:        schema.TypeString,
				Required:    true,
			},
			"ignore_error": {
				Description: "ignore error default false",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"yaml": {
				Description: "trigger yaml",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"webhook_url": {
				Description: "webhook url",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"executor_info": {
				Description: "Executor identity for this trigger.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description: "Type of principal that executes the pipeline when this trigger fires.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"identifier": {
							Description: "Identifier of the executor principal.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"org_id": {
							Description: "Org identifier of the SERVICE_ACCOUNT executor.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"project_id": {
							Description: "Project identifier of the SERVICE_ACCOUNT executor.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
	helpers.SetProjectLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataTriggersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)

	resp, httpResp, err := c.TriggersApi.GetTrigger(ctx, c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string), d.Get("target_id").(string), id)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if err := readTriggers(d, resp.Data); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
