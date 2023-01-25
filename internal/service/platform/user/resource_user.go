package user

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceUser() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating a Harness User.",

		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserCreateOrUpdate,
		DeleteContext: resourceUserDelete,
		CreateContext: resourceUserCreateOrUpdate,
		Importer:      helpers.UserResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"org_id": {
				Description: "Organization identifier of the user.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project identifier of the user.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"email": {
				Description: "The email of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"disabled": {
				Description: "Whether or not the user account is disabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"locked": {
				Description: "Whether or not the user account is locked.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"externally_managed": {
				Description: "Whether or not the user account is externally managed.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"emails": {
				Description: "The email of the user.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_groups": {
				Description: "The user group of the user.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"role_bindings": {
				Description: "Role Bindings of the user.",
				Type:        schema.TypeList,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_group_identifier": {
							Description: "Resource Group Identifier of the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"role_identifier": {
							Description: "Role Identifier of the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"role_name": {
							Description: "Role Name Identifier of the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"resource_group_name": {
							Description: "Resource Group Name of the user.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"managed_role": {
							Description: "Managed Role of the user.",
							Type:        schema.TypeBool,
							Optional:    true,
						},
					}},
			},
		},
	}

	return resource
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	emails := []string{}
	var email = ""
	if attr, ok := d.GetOk("emails"); ok {
		emails = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
		email = emails[0]
	}

	resp, httpResp, err := c.UserApi.GetAggregatedUsers(ctx, c.AccountId, &nextgen.UserApiGetAggregatedUsersOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		SearchTerm:        optional.NewString(email),
	})

	if resp.Data.Empty {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if &resp == nil || resp.Data == nil || resp.Data.Empty {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readUser(d, &resp.Data.Content[0])

	return nil
}

func resourceUserCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	var err error
	var httpResp *http.Response

	if id == "" {
		addUserBody := createAddUserBody(d)
		_, httpResp, err = c.UserApi.AddUsers(ctx, *addUserBody, c.AccountId, &nextgen.UserApiAddUsersOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		updateUSerBody := updateAddUserBody(d)
		_, httpResp, err = c.UserApi.UpdateUserInfo(ctx, c.AccountId, &nextgen.UserApiUpdateUserInfoOpts{
			Body: optional.NewInterface(updateUSerBody),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	emails := []string{}
	if attr, ok := d.GetOk("emails"); ok {
		emails = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	var email = emails[0]
	resp, httpResp, err := c.UserApi.GetAggregatedUsers(ctx, c.AccountId, &nextgen.UserApiGetAggregatedUsersOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		SearchTerm:        optional.NewString(email),
	})
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if &resp == nil || resp.Data == nil || resp.Data.Empty {
		return nil
	}

	readUser(d, &resp.Data.Content[0])

	return nil
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.UserApi.RemoveUser(ctx, d.Id(), c.AccountId, &nextgen.UserApiRemoveUserOpts{
		OrgIdentifier:     optional.NewString(d.Get("org_id").(string)),
		ProjectIdentifier: optional.NewString(d.Get("project_id").(string)),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func createAddUserBody(d *schema.ResourceData) *nextgen.AddUsersDto {

	var addUsersDto nextgen.AddUsersDto
	if attr, ok := d.GetOk("emails"); ok {
		addUsersDto.Emails = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("user_groups"); ok {
		addUsersDto.UserGroups = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	var ipRoleBindings map[string]interface{}

	if attr, ok := d.GetOk("role_bindings"); ok {
		var roleBindings []nextgen.RoleBinding
		ipRoleBindings = attr.([]interface{})[0].(map[string]interface{})
		var roleBindingObj nextgen.RoleBinding
		if ipRoleBindings["resource_group_identifier"] != nil && len(ipRoleBindings["resource_group_identifier"].(string)) > 0 {
			roleBindingObj.ResourceGroupIdentifier = ipRoleBindings["resource_group_identifier"].(string)
		}

		if ipRoleBindings["role_identifier"] != nil && len(ipRoleBindings["role_identifier"].(string)) > 0 {
			roleBindingObj.RoleIdentifier = ipRoleBindings["role_identifier"].(string)
		}

		if ipRoleBindings["role_name"] != nil && len(ipRoleBindings["role_name"].(string)) > 0 {
			roleBindingObj.RoleName = ipRoleBindings["role_name"].(string)
		}

		if ipRoleBindings["resource_group_name"] != nil && len(ipRoleBindings["resource_group_name"].(string)) > 0 {
			roleBindingObj.ResourceGroupName = ipRoleBindings["resource_group_name"].(string)
		}

		if ipRoleBindings["managed_role"] != nil {
			roleBindingObj.ManagedRole = ipRoleBindings["managed_role"].(bool)
		}

		roleBindings = append(roleBindings, roleBindingObj)
		addUsersDto.RoleBindings = roleBindings
	}

	return &addUsersDto
}

func updateAddUserBody(d *schema.ResourceData) *nextgen.UserInfo {
	return &nextgen.UserInfo{
		Uuid: d.Get("identifier").(string),
		Name: d.Get("name").(string),
	}
}

func readUser(d *schema.ResourceData, UserAggregate *nextgen.UserAggregate) {
	d.SetId(UserAggregate.User.Uuid)
	d.Set("identifier", UserAggregate.User.Uuid)
	d.Set("name", UserAggregate.User.Name)
	d.Set("email", UserAggregate.User.Email)
	d.Set("locked", UserAggregate.User.Locked)
	d.Set("disabled", UserAggregate.User.Disabled)
	d.Set("externally_managed", UserAggregate.User.ExternallyManaged)
}
