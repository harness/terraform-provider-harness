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
			"created_at": {
				Description: "Created DateTime of the folder.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	helpers.SetCommonDataSourceSchema(resource.Schema)

	// Override common schema: lookup by exactly one of id or name.
	resource.Schema["id"] = &schema.Schema{
		Description:  "Identifier of the folder. Required if name is not provided.",
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: []string{"id", "name"},
	}
	resource.Schema["name"] = &schema.Schema{
		Description:  "Name of the folder. Required if id is not provided.",
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: []string{"id", "name"},
	}

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
			flat := flattenFolders(all)
			for i := range flat {
				if flat[i].Name == name {
					folder = &flat[i]
					break
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
		if id != "" {
			return diag.Errorf("no dashboard folder found with id %q", id)
		}
		return diag.Errorf("no dashboard folder found with name %q", name)
	}

	readFolder(d, folder)

	return nil
}
