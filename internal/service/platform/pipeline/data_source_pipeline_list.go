package pipeline

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-openapi-go-client/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourcePipelineList() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retieving the Harness pipleine List",

		ReadContext: dataSourcePipelineListRead,

		Schema: map[string]*schema.Schema{
			"pipelines": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"page": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}

	helpers.SetProjectLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourcePipelineListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetClientWithContext(ctx)

	var err error
	var httpResp *http.Response

	org_id := d.Get("org_id").(string)
	project_id := d.Get("project_id").(string)
	page := d.Get("page").(int)
	limit := d.Get("limit").(int)

	resp, httpResp, err := c.PipelinesApi.ListPipelines(ctx,
		org_id,
		project_id,
		&nextgen.PipelinesApiListPipelinesOpts{HarnessAccount: optional.NewString(c.AccountId), Page: optional.NewInt32(int32(page)), Limit: optional.NewInt32(int32(limit))},
	)

	var output = resp
	var pipelines []map[string]interface{}
	for _, v := range output {
		newPipeline := map[string]interface{}{
			"identifier": v.Identifier,
			"name":       v.Name,
		}

		pipelines = append(pipelines, newPipeline)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if pipelines == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	d.SetId(project_id)
	d.Set("pipelines", pipelines)

	return nil
}
