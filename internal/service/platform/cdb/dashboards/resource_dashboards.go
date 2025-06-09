package dashboards

import (
	"context"
	"strconv"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceDashboards() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for a Harness Custom Dashboard.",

		CreateContext: resourceDashboardClone,
		ReadContext:   resourceDashboardRead,
		UpdateContext: resourceDashboardUpdate,
		DeleteContext: resourceDashboardDelete,
		Importer:      helpers.OrgResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "Created at timestamp of the Dashboard.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"dashboard_id": {
				Description: "Unique identifier of the Dashboard.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"data_source": {
				Description: "Data Sources within the Dashboard.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeList,
				Optional:    true,
			},
			"description": {
				Description: "Description of the Dashboard.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"folder_id": {
				Description: "Unique identifier of the Folder.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"models": {
				Description: "Data Models within the Dashboard.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeList,
				Optional:    true,
			},
			"name": {
				Description: "Name of the Dashboard.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resource_identifier": {
				Description: "Resource identifier of the dashboard.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"title": {
				Description: "Title of the Dashboard.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "Resource identifier of the dashboard.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"view_count": {
				Description: "View count of the dashboard.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
		},
	}

	return resource
}

func resourceDashboardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" {
		d.MarkNewResource()
		return nil
	}

	resp, httpResp, err := c.DashboardsApi.GetDashboard(ctx, id, &nextgen.DashboardsApiGetDashboardOpts{
		AccountId: optional.NewString(c.AccountId),
	})

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Resource == nil {
		d.SetId("")
		return nil
	}

	readDashboard(d, resp.Resource)

	return nil
}

func resourceDashboardUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	updateRequestBody := buildDashboardCreateRequest(d)

	resp, httpResp, err := c.DashboardsApi.UpdateDashboard(ctx, *updateRequestBody, &nextgen.DashboardsApiUpdateDashboardOpts{
		AccountId: optional.NewString(c.AccountId),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readUpdateDashboard(d, resp.Resource)

	return nil
}

func resourceDashboardClone(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	cloneRequestBody := buildDashboardCloneRequest(d)

	resp, httpResp, err := c.DashboardsApi.CloneDashboard(ctx, *cloneRequestBody, &nextgen.DashboardsApiCloneDashboardOpts{
		AccountId: optional.NewString(c.AccountId),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readClonedDashboard(d, resp.Resource)

	return nil
}

func resourceDashboardDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	deleteRequest := buildDashboardDeleteRequest(d)
	folderId := d.Get("folderId").(string)

	_, httpResp, err := c.DashboardsApi.DeleteDashboard(ctx, *deleteRequest, folderId, &nextgen.DashboardsApiDeleteDashboardOpts{
		AccountId: optional.NewString(c.AccountId),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildDashboardCloneRequest(d *schema.ResourceData) *nextgen.CloneDashboardRequestBody {
	return &nextgen.CloneDashboardRequestBody{
		DashboardId: d.Get("dashboard_id").(string),
		FolderId:    d.Get("folder_id").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}
}

func buildDashboardCreateRequest(d *schema.ResourceData) *nextgen.CreateDashboardRequest {
	return &nextgen.CreateDashboardRequest{
		DashboardId: d.Get("dashboard_id").(int32),
		FolderId:    d.Get("folder_id").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}
}

func buildDashboardDeleteRequest(d *schema.ResourceData) *nextgen.DeleteDashboardRequest {
	return &nextgen.DeleteDashboardRequest{
		DashboardId: d.Get("dashboard_id").(string),
	}
}

func readDashboard(d *schema.ResourceData, dashboard *nextgen.Dashboard) {
	d.SetId(dashboard.Id)
	d.Set("type", dashboard.Type_)
	d.Set("description", dashboard.Description)
	d.Set("title", dashboard.Title)
	d.Set("view_count", dashboard.ViewCount)
	d.Set("favorite_count", dashboard.FavoriteCount)
	d.Set("created_at", dashboard.CreatedAt)
	d.Set("data_source", dashboard.DataSource)
	d.Set("models", dashboard.Models)
	d.Set("last_accessed_at", dashboard.LastAccessedAt)
	d.Set("resource_identifier", dashboard.ResourceIdentifier)
	d.Set("folder_id", dashboard.Folder.Id)
}

func readClonedDashboard(d *schema.ResourceData, dashboard *nextgen.ClonedDashboard) {
	d.SetId(dashboard.Id)
	d.Set("description", dashboard.Description)
	d.Set("title", dashboard.Title)
	d.Set("resource_identifier", dashboard.ResourceIdentifier)
}

func readUpdateDashboard(d *schema.ResourceData, dashboard *nextgen.UpdateDashboardResponseResource) {
	d.SetId(strconv.FormatInt(int64(dashboard.Id), 10))
	d.Set("description", dashboard.Description)
	d.Set("title", dashboard.Title)
	d.Set("resource_identifier", dashboard.ResourceIdentifier)
}
