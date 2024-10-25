package service

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceServiceList() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness service List.",

		ReadContext: dataSourceServiceListRead,

		Schema: map[string]*schema.Schema{
			"services": {
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

func dataSourceServiceListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var httpResp *http.Response

	var resp nextgen.ResponseDtoPageResponseServiceResponse
	resp, httpResp, err = c.ServicesApi.GetServiceList(ctx, c.AccountId, &nextgen.ServicesApiGetServiceListOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})
	var output []nextgen.ServiceResponse = resp.Data.Content
	var services []map[string]interface{}
	for _, v := range output {
		newService := map[string]interface{}{
			"identifier": v.Service.Identifier,
			"name":       v.Service.Name,
		}

		services = append(services, newService)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if services == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	d.SetId(resp.CorrelationId)
	d.Set("services", services)

	return nil
}
