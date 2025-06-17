package folders

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceDashboardFolders() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for a Harness Custom Dashboard Folder.",

		CreateContext: resourceFolderCreate,
		ReadContext:   resourceFolderRead,
		UpdateContext: resourceFolderUpdate,
		DeleteContext: resourceFolderDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Identifier of the folder.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the folder.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"created_at": {
				Description: "Created DateTime of the folder.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}

func resourceFolderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	if id == "" || id == "0" {
		d.SetId("")
		return nil
	}

	resp, httpResp, err := c.DashboardsFolderApi.GetFolder(ctx, id, &nextgen.DashboardsFoldersApiGetFolderOpts{
		AccountId: optional.NewString(c.AccountId),
	})

	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			// Folder not found, remove from state
			d.SetId("")
			return nil
		}
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Resource == nil {
		d.SetId("")
		return nil
	}

	readFolder(d, resp.Resource)

	return nil
}

func resourceFolderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	updateRequestBody := buildFolderUpdateRequest(d)

	resp, httpResp, err := c.DashboardsFolderApi.UpdateFolder(ctx, *updateRequestBody, d.Id(), &nextgen.DashboardsFoldersApiUpdateFolderOpts{
		AccountId: optional.NewString(c.AccountId),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readFolder(d, resp.Resource)

	return nil
}

func resourceFolderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	createRequestBody := buildFolderCreateRequest(d)

	resp, httpResp, err := c.DashboardsFolderApi.CreateFolder(ctx, *createRequestBody, &nextgen.DashboardsFoldersApiCreateFolderOpts{
		AccountId: optional.NewString(c.AccountId),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readFolder(d, resp.Resource)

	return nil
}

func resourceFolderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	folderId := d.Id()

	_, httpResp, err := c.DashboardsFolderApi.DeleteFolder(ctx, folderId, &nextgen.DashboardsFoldersApiDeleteFolderOpts{
		AccountId: optional.NewString(c.AccountId),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildFolderCreateRequest(d *schema.ResourceData) *nextgen.CreateFolderRequestBody {
	req := &nextgen.CreateFolderRequestBody{}
	if v, ok := d.GetOk("name"); ok {
		req.Name = v.(string)
	}
	return req
}

func buildFolderUpdateRequest(d *schema.ResourceData) *nextgen.UpdateFolderRequestBody {
	req := &nextgen.UpdateFolderRequestBody{}
	if v, ok := d.GetOk("name"); ok {
		req.Name = v.(string)
	}
	return req
}

func readFolder(d *schema.ResourceData, folder *nextgen.Folder) {
	d.SetId(folder.Id)
	d.Set("Name", folder.Name)
}
