package input_set

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInputSet() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness input set.",

		ReadContext: dataSourceInputSetRead,

		Schema: map[string]*schema.Schema{
			"pipeline_id": {
				Description: "Identifier of the pipeline",
				Type:        schema.TypeString,
				Required:    true,
			},
			"yaml": {
				Description: "Input Set YAML",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	helpers.SetProjectLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceInputSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.InputSetsApi.GetInputSet(ctx,
		d.Get("identifier").(string),
		c.AccountId,
		d.Get("org_id").(string),
		d.Get("project_id").(string),
		d.Get("pipeline_id").(string),
		&nextgen.InputSetsApiGetInputSetOpts{},
	)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readInputSet(d, resp.Data)

	return nil
}
