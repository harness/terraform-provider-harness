package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceService() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness service.",

		ReadContext: dataSourceServiceRead,

		Schema: map[string]*schema.Schema{
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

func dataSourceServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var svc *nextgen.ServiceResponseDetails
	var httpResp *http.Response

	id := d.Get("identifier").(string)
	name := d.Get("name").(string)

	if id != "" {
		var resp nextgen.ResponseDtoServiceResponse
		resp, httpResp, err = c.ServicesApi.GetServiceV2(ctx, d.Get("identifier").(string), c.AccountId, &nextgen.ServicesApiGetServiceV2Opts{
			OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		})
		svc = resp.Data.Service
	} else if name != "" {
		svc, httpResp, err = c.ServicesApi.GetServiceByName(ctx, c.AccountId, name, nextgen.GetServiceByNameOpts{
			OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		})
	} else {
		return diag.FromErr(errors.New("either identifier or name must be specified"))
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if svc == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readService(d, svc)

	return nil
}
