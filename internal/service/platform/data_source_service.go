package platform

import (
	"context"

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

		Schema: map[string]*schema.Schema{},
	}

	helpers.SetProjectLevelDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, _, err := c.ServicesApi.GetServiceV2(ctx, d.Get("identifier").(string), c.AccountId, &nextgen.ServicesApiGetServiceV2Opts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	if resp.Data == nil || resp.Data.Service == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readService(d, resp.Data.Service)

	return nil
}
