package permissions

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourcePermissions() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving permissions",

		ReadContext: dataSourcePermissionsRead,
		Schema: map[string]*schema.Schema{
						"org_id": {
							Description: "Organization Identifier",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"project_id": {
							Description: "Project Identifier",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"permissions":{
							Description: "Response of the api",
							Type:        schema.TypeList,
							Computed:    true,
						},
		},
	}

	return resource
}

func dataSourcePermissionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error

	permissionsApiGetRoleOpts := &nextgen.PermissionsApiGetPermissionListOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	}

	resp, httpResp, err := c.PermissionsApi.GetPermissionList(ctx, permissionsApiGetRoleOpts)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readPermissions(d, resp.Data)
    
	return nil
}

func readPermissions(d *schema.ResourceData, permissionResponse []nextgen.PermissionResponse) {
	d.SetId(permissionResponse[0].Permission.Identifier)
	d.Set("permissions", permissionResponse)
}