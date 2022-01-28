package user

import (
	"context"
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceUserGroupPermissionsPermissions() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for adding permissions to an existing Harness user group.",

		CreateContext: resourceUserGroupPermissionsCreateOrUpdate,
		ReadContext:   resourceUserGroupPermissionsRead,
		UpdateContext: resourceUserGroupPermissionsCreateOrUpdate,
		DeleteContext: resourceUserGroupPermissionsDelete,

		Schema: map[string]*schema.Schema{
			"user_group_id": {
				Description: "Unique identifier of the user group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"account_permissions": getUserGroupAccountPermissionsSchema(),
			"app_permissions":     getUserGroupAccountPermissionsSchema(),
		},

		Importer: &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}

func resourceUserGroupPermissionsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Get("user_group_id").(string)
	ug, err := c.CDClient.UserClient.GetUserGroupById(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if ug == nil {
		return diag.FromErr(fmt.Errorf("user group %s does not exist", id))
	}

	input := &graphql.UserGroup{
		Id:          ug.Id,
		Permissions: &graphql.UserGroupPermissions{},
	}

	expandAccountPermissions(d.Get("account_permissions").(*schema.Set).List(), input.Permissions)
	expandAppPermissions(d.Get("app_permissions").(*schema.Set).List(), input.Permissions)

	updatedUg, err := c.CDClient.UserClient.UpdateUserGroup(ug)
	if err != nil {
		return diag.FromErr(err)
	}

	// Computed fields
	readUserGroupPermissions(d, updatedUg)

	return nil
}

func resourceUserGroupPermissionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Get("user_group_id").(string)

	userGroup, err := c.CDClient.UserClient.GetUserGroupById(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if userGroup == nil {
		return diag.FromErr(fmt.Errorf("user group %s does not exist", id))
	}

	return readUserGroupPermissions(d, userGroup)
}

func readUserGroupPermissions(d *schema.ResourceData, userGroup *graphql.UserGroup) diag.Diagnostics {
	d.SetId(userGroup.Id)

	if accountPermission := flattenAccountPermissions(userGroup.Permissions); len(accountPermission) > 0 {
		d.Set("account_permissions", accountPermission)
	}

	if appPermissions := flattenAppPermissions(userGroup.Permissions); len(appPermissions) > 0 {
		d.Set("app_permissions", appPermissions)
	}

	return nil
}

func resourceUserGroupPermissionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	ug, err := c.CDClient.UserClient.GetUserGroupById(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	ug.Permissions.AccountPermissions = &graphql.AccountPermissions{}
	ug.Permissions.AppPermissions = []*graphql.AppPermission{}

	updatedUg, err := c.CDClient.UserClient.UpdateUserGroup(ug)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(updatedUg.Permissions.AccountPermissions.AccountPermissionTypes) > 0 {
		return diag.FromErr(fmt.Errorf("failed to delete user group permissions"))
	}

	if len(updatedUg.Permissions.AppPermissions) > 0 {
		return diag.FromErr(fmt.Errorf("failed to delete user group permissions"))
	}

	d.SetId("")

	return nil
}
