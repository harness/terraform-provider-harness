package roles

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRoles() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving roles",

		ReadContext: dataSourceRolesRead,
		Schema: map[string]*schema.Schema{
			"permissions": {
				Description: "List of the permission identifiers ",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allowed_scope_levels": {
				Description: "The scope levels at which this role can be used",
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}

func dataSourceRolesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Get("identifier").(string)

	var err error

	rolesApiGetRoleOpts := &nextgen.RolesApiGetRoleOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	}

	resp, httpResp, err := c.RolesApi.GetRole(ctx, id, rolesApiGetRoleOpts)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readRoles(d, resp.Data.Role)

	return nil
}
