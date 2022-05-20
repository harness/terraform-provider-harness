package usergroup

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceUserGroup() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness User Group.",

		ReadContext:   resourceUserGroupRead,
		UpdateContext: resourceUserGroupCreateOrUpdate,
		DeleteContext: resourceUserGroupDelete,
		CreateContext: resourceUserGroupCreateOrUpdate,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			// "is_sso_linked": {
			// 	Description: "Whether the user group is linked to an SSO account.",
			// 	Type:        schema.TypeBool,
			// 	Optional:    true,
			// },
			// "linked_sso_id": {
			// 	Description: "The SSO account ID that the user group is linked to.",
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// },
			// "linked_sso_name": {
			// 	Description: "The SSO account name that the user group is linked to.",
			// 	Type:        schema.TypeString,
			// 	Optional:    true,
			// },
			"externally_managed": {
				Description: "Whether the user group is externally managed.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}

	helpers.SetProjectLevelResourceSchema(resource.Schema)

	return resource
}

func resourceUserGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	resp, _, err := c.UserGroupApi.GetUserGroup(ctx, c.AccountId, id, &nextgen.UserGroupApiGetUserGroupOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	// Soft delete lookup error handling
	// https://harness.atlassian.net/browse/PL-23765
	if resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readUserGroup(d, resp.Data)

	return nil
}

func resourceUserGroupCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoUserGroup

	id := d.Id()
	ug := buildUserGroup(d)
	ug.AccountIdentifier = c.AccountId

	if id == "" {
		resp, _, err = c.UserGroupApi.PostUserGroup(ctx, ug, c.AccountId, &nextgen.UserGroupApiPostUserGroupOpts{
			OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		})
	} else {
		resp, _, err = c.UserGroupApi.PutUserGroup(ctx, ug, c.AccountId, &nextgen.UserGroupApiPutUserGroupOpts{
			OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
			ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d)
	}

	readUserGroup(d, resp.Data)

	return nil
}

func resourceUserGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, _, err := c.UserGroupApi.DeleteUserGroup(ctx, c.AccountId, d.Id(), &nextgen.UserGroupApiDeleteUserGroupOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	return nil
}

func buildUserGroup(d *schema.ResourceData) nextgen.UserGroup {
	return nextgen.UserGroup{
		Identifier:        d.Get("identifier").(string),
		OrgIdentifier:     d.Get("org_id").(string),
		ProjectIdentifier: d.Get("project_id").(string),
		Name:              d.Get("name").(string),
		Description:       d.Get("description").(string),
		Tags:              helpers.ExpandTags(d.Get("tags").(*schema.Set).List()),
	}
}

func readUserGroup(d *schema.ResourceData, env *nextgen.UserGroup) {
	d.SetId(env.Identifier)
	d.Set("identifier", env.Identifier)
	d.Set("org_id", env.OrgIdentifier)
	d.Set("name", env.Name)
	d.Set("description", env.Description)
	d.Set("tags", helpers.FlattenTags(env.Tags))
}
