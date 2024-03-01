package environment

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceEnvironmentList() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness environment List.",

		ReadContext: dataSourceEnvironmentListRead,

		Schema: map[string]*schema.Schema{
			"environments": {
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
		},
	}

	helpers.SetOptionalOrgAndProjectLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceEnvironmentListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var httpResp *http.Response

	var resp nextgen.ResponseDtoPageResponseEnvironmentResponse
	resp, httpResp, err = c.EnvironmentsApi.GetEnvironmentList(ctx, c.AccountId, &nextgen.EnvironmentsApiGetEnvironmentListOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	var output = resp.Data.Content
	var environments []map[string]interface{}
	for _, v := range output {
		newEnvironment := map[string]interface{}{
			"identifier": v.Environment.Identifier,
			"name":       v.Environment.Name,
		}

		environments = append(environments, newEnvironment)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if environments == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	d.SetId(resp.CorrelationId)
	d.Set("environments", environments)

	return nil
}
