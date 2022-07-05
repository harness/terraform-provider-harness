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

func ResourceRoles() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating roles.",
		ReadContext:   resourceRolesRead,
		UpdateContext: resourceRolesCreateOrUpdate,
		DeleteContext: resourceRolesDelete,
		CreateContext: resourceRolesCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

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

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceRolesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	rolesApiGetRoleOpts := &nextgen.RolesApiGetRoleOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	}

	resp, _, err := c.RolesApi.GetRole(ctx, id, rolesApiGetRoleOpts)

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	readRoles(d, resp.Data.Role)

	return nil
}

func resourceRolesCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoRoleResponse
	id := d.Id()

	role := buildRoles(d)

	if id == "" {
		resp, _, err = c.RolesApi.PostRole(ctx, *role, &nextgen.RolesApiPostRoleOpts{
			AccountIdentifier: optional.NewString(c.AccountId),
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		resp, _, err = c.RolesApi.PutRole(ctx, *role, id, &nextgen.RolesApiPutRoleOpts{
			AccountIdentifier: optional.NewString(c.AccountId),
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	}

	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	readRoles(d, resp.Data.Role)

	return nil
}

func resourceRolesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, _, err := c.RolesApi.DeleteRole(ctx, d.Id(), &nextgen.RolesApiDeleteRoleOpts{
		AccountIdentifier: optional.NewString(c.AccountId),
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}

func buildRoles(d *schema.ResourceData) *nextgen.Role {
	return &nextgen.Role{
		Identifier:         d.Get("identifier").(string),
		Name:               d.Get("name").(string),
		Description:        d.Get("description").(string),
		Tags:               helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
		Permissions:        helpers.ExpandField(d.Get("permissions").(*schema.Set).List()),
		AllowedScopeLevels: helpers.ExpandField(d.Get("allowed_scope_levels").(*schema.Set).List()),
	}
}

func readRoles(d *schema.ResourceData, role *nextgen.Role) {
	d.SetId(role.Identifier)
	d.Set("identifier", role.Identifier)
	d.Set("name", role.Name)
	d.Set("description", role.Description)
	d.Set("tags", helpers.FlattenTags(role.Tags))
	d.Set("permissions", role.Permissions)
	d.Set("allowed_scope_levels", role.AllowedScopeLevels)
}
