package monitored_service

import (
	"context"
	hh "github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceMonitoredService() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a monitored service.",

		ReadContext: dataSourceMonitoredServiceRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "Account Identifier of the monitored service.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"org_id": {
				Description: "Organization Identifier of the monitored service.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project Identifier of the monitored service.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "Identifier of the monitored service.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}

	return resource
}

func dataSourceMonitoredServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	ctx = context.WithValue(ctx, nextgen.ContextAccessToken, hh.EnvVars.BearerToken.Get())
	var accountIdentifier, orgIdentifier, projectIdentifier string
	identifier := d.Get("identifier").(string)
	if attr, ok := d.GetOk("account_id"); ok {
		accountIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("org_id"); ok {
		orgIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		projectIdentifier = attr.(string)
	}
	resp, httpResp, err := c.MonitoredServiceApi.GetMonitoredService(ctx, identifier, accountIdentifier, orgIdentifier, projectIdentifier)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if &resp == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}
	readMonitoredService(d, &resp.Data, accountIdentifier)
	return nil
}
