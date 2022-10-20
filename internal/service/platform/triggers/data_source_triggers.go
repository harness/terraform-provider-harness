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
				Description: "identifier of the cluster.",
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

	readTriggers(d, resp.Data)

	return nil
}
