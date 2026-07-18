package folders

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

func DataSourceDashboardFolder() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness Custom Dashboard Folder by id or name.",

		ReadContext: dataSourceFolderRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Identifier of the folder. Required if name is not provided.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Name of the folder. Required if id is not provided.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"created_at": {
				Description: "Created DateTime of the folder.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	helpers.SetCommonDataSourceSchema(resource.Schema)

	return resource
}

func dataSourceFolderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("id").(string)
	name := d.Get("name").(string)

	var folder *nextgen.Folder
	var httpResp *http.Response
	var err error

	if id != "" {
		var resp nextgen.GetFolderResponse
		resp, httpResp, err = c.DashboardsFolderApi.GetFolder(ctx, id, &nextgen.DashboardsFoldersApiGetFolderOpts{
			AccountId: optional.NewString(c.AccountId),
		})
		folder = resp.Resource
	} else if name != "" {
		// Fallback to list + filter by name (SDK has no GetFolderByName)
		var all []nextgen.Folder
		all, httpResp, err = listDashboardFolders(c)
		if err == nil {
			for i := range all {
				if all[i].Name == name {
					folder = &all[i]
					break
				}
			}
			if folder == nil {
				// also search inside subfolders (listDashboardFolders already flattens? but call flatten to be sure)
				flat := flattenFolders(all)
				for i := range flat {
					if flat[i].Name == name {
						folder = &flat[i]
						break
					}
				}
			}
		}
	} else {
		return diag.FromErr(errors.New("either id or name must be specified"))
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if folder == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readFolder(d, folder)

	return nil
}
