package dashboards

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

func DataSourceDashboard() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness Dashboard.",

		ReadContext: dataSourceDashboardRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Identifier of the dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "Created at timestamp of the Dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"dashboard_id": {
				Description: "Unique identifier of the Dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"data_source": {
				Description: "Data Sources within the Dashboard.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeList,
				Computed:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "Description of the Dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"folder_id": {
				Description: "Unique identifier of the Folder.",
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
			},
			"models": {
				Description: "Data Models within the Dashboard.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeList,
				Computed:    true,
			},
			"name": {
				Description: "Name of the Dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"resource_identifier": {
				Description: "Resource identifier of the dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"title": {
				Description: "Title of the Dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "Resource identifier of the dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"view_count": {
				Description: "View count of the dashboard.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}

	helpers.SetCommonDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceDashboardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)

	if id == "" {
		return diag.FromErr(errors.New("identifier must be specified"))
	}

	var err error
	var dashboard *nextgen.Dashboard
	var httpResp *http.Response
	var resp nextgen.GetDashboardResponse

	resp, httpResp, err = c.DashboardsApi.GetDashboard(ctx, id, &nextgen.DashboardsApiGetDashboardOpts{
		AccountId: optional.NewString(c.AccountId),
	})
	dashboard = resp.Resource

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if dashboard == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readDashboard(d, dashboard)

	return nil
}
