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
			"id": {
				Description: "Identifier of the dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "Created at timestamp of the Dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"dashboard_id": {
				Description: "Unique identifier of the Template Dashboard to create from.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
			"data_source": {
				Description: "Data Sources within the Dashboard.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
			},
			"description": {
				Description: "Description of the Dashboard.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"folder_id": {
				Description: "The Folder ID that the Dashboard belongs to.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"models": {
				Description: "Data Models within the Dashboard.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
			},
			"name": {
				Description: "Name of the Dashboard.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resource_identifier": {
				Description: "The Folder ID that the Dashboard belongs to.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
			"title": {
				Description: "Title of the Dashboard.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"type": {
				Description: "Type of the dashboard.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"view_count": {
				Description: "View count of the dashboard.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
		},
	}

	return resource
}

func resourceDashboardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" || id == "0" {
		d.SetId("")
		return nil
	}

	resp, httpResp, err := c.DashboardsApi.GetDashboard(ctx, id, &nextgen.DashboardsApiGetDashboardOpts{
		AccountId: optional.NewString(c.AccountId),
	})

	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			// Dashboard not found, remove from state
			d.SetId("")
			return nil
		}
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

	updateRequestBody := buildDashboardUpdateRequest(d)

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
	folderId := d.Get("resource_identifier").(string)

	_, httpResp, err := c.DashboardsApi.DeleteDashboard(ctx, *deleteRequest, folderId, &nextgen.DashboardsApiDeleteDashboardOpts{
		AccountId: optional.NewString(c.AccountId),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildDashboardCloneRequest(d *schema.ResourceData) *nextgen.CloneDashboardRequestBody {
	req := &nextgen.CloneDashboardRequestBody{}

	if v, ok := d.GetOk("dashboard_id"); ok {
		req.DashboardId = v.(string)
	}

	// Use resource_identifier as folderId if provided, otherwise use folder_id
	if v, ok := d.GetOk("resource_identifier"); ok {
		req.FolderId = v.(string)
	} else if v, ok := d.GetOk("folder_id"); ok {
		req.FolderId = v.(string)
	}

	// Use title as name if name is not provided
	if v, ok := d.GetOk("name"); ok {
		req.Name = v.(string)
	} else if v, ok := d.GetOk("title"); ok {
		req.Name = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		req.Description = v.(string)
	}

	return req
}

func buildDashboardUpdateRequest(d *schema.ResourceData) *nextgen.CreateDashboardRequest {
	req := &nextgen.CreateDashboardRequest{}

	if v, ok := d.GetOk("id"); ok {
		if id, err := strconv.ParseInt(v.(string), 10, 32); err == nil {
			req.DashboardId = int32(id)
		}
	}

	// Use resource_identifier as folderId if provided, otherwise use folder_id
	if v, ok := d.GetOk("resource_identifier"); ok {
		req.FolderId = v.(string)
	} else if v, ok := d.GetOk("folder_id"); ok {
		req.FolderId = v.(string)
	}

	// Use title as name if name is not provided
	if v, ok := d.GetOk("name"); ok {
		req.Name = v.(string)
	} else if v, ok := d.GetOk("title"); ok {
		req.Name = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		req.Description = v.(string)
	}

	return req
}

func buildDashboardDeleteRequest(d *schema.ResourceData) *nextgen.DeleteDashboardRequest {
	return &nextgen.DeleteDashboardRequest{
		DashboardId: d.Id(),
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

	if dashboard.ResourceIdentifier != "" {
		d.Set("resource_identifier", dashboard.ResourceIdentifier)
	}

	if v, ok := d.GetOk("resource_identifier"); ok && v.(string) == dashboard.Folder.Id {
	} else if _, ok := d.GetOk("folder_id"); ok {
		d.Set("folder_id", dashboard.Folder.Id)
	}
}

func readClonedDashboard(d *schema.ResourceData, dashboard *nextgen.ClonedDashboard) {
	if dashboard.Id != "" && dashboard.Id != "0" {
		d.SetId(dashboard.Id)
	}
	d.Set("description", dashboard.Description)
	d.Set("title", dashboard.Title)
	d.Set("resource_identifier", dashboard.ResourceIdentifier)
}

func readUpdateDashboard(d *schema.ResourceData, dashboard *nextgen.UpdateDashboardResponseResource) {
	if dashboard.Id != "" && dashboard.Id != "0" {
		d.SetId(dashboard.Id)
	}
	d.Set("description", dashboard.Description)
	d.Set("title", dashboard.Title)
	d.Set("resource_identifier", dashboard.ResourceIdentifier)
}
