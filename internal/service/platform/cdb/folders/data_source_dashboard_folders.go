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

func DataSourceDashboardFolders() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness Dashboard Folder.",

		ReadContext: dataSourceFolderRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Identifier of the folder.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the folder.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
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

	if id == "" {
		return diag.FromErr(errors.New("Id must be specified"))
	}

	var err error
	var folder *nextgen.Folder
	var httpResp *http.Response
	var resp nextgen.GetFolderResponse

	resp, httpResp, err = c.DashboardsFolderApi.GetFolder(ctx, id, &nextgen.DashboardsFoldersApiGetFolderOpts{
		AccountId: optional.NewString(c.AccountId),
	})
	folder = resp.Resource

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
