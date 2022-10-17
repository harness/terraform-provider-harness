package permissions

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/utils"
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
			"permissions": {
				Description: "Response of the api",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"identifier": {
							Description: "Identifier of the permission",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name of the permission",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status": {
							Description: "Status of the permission",
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
						},
						"resource_type": {
							Description: "Resource type for the given permission",
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
						},
						"action": {
							Description: "Action performed by the permission",
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
						},
						"allowed_scope_levels": {
							Description: "The scope levels at which this resource group can be used",
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"include_in_all_roles": {
							Description: "Is included in all roles",
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
						},
					}},
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
	var id = c.AccountId
	
	id = fmt.Sprintf("%d", utils.StringHashcode(c.AccountId+d.Get("org_id").(string)+d.Get("project_id").(string)))
	readPermissions(d, resp.Data, id)

	return nil
}

func FlattenPermissions(permissionResponse []nextgen.PermissionResponse) []map[string]interface{} {
	if permissionResponse == nil {
		return make([]map[string]interface{}, 0)
	}
	results := make([]map[string]interface{}, len(permissionResponse))

	for i, res := range permissionResponse {
		results[i] = map[string]interface{}{
			"identifier":           res.Permission.Identifier,
			"name":                 res.Permission.Name,
			"status":               res.Permission.Status,
			"resource_type":        res.Permission.ResourceType,
			"action":               res.Permission.Action,
			"allowed_scope_levels": res.Permission.AllowedScopeLevels,
			"include_in_all_roles": res.Permission.IncludeInAllRoles,
		}
	}

	return results
}

func readPermissions(d *schema.ResourceData, permissionResponse []nextgen.PermissionResponse, id string) {
	d.SetId(id)
	d.Set("permissions", FlattenPermissions(permissionResponse))
}
