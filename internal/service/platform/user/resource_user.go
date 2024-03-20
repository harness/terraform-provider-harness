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
		Description: "Resource for creating a Harness User. This requires your authentication mechanism to be set to SAML, LDAP, or OAuth, and the feature flag AUTO_ACCEPT_SAML_ACCOUNT_INVITES to be enabled.",

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
				Optional:    true,
			},
			"project_id": {
				Description:  "Project identifier of the user.",
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"org_id"},
			},
			"name": {
				Description: "Name of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"email": {
				Description: "The email of the user.",
				Type:        schema.TypeString,
				Required:    true,
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
			"user_groups": {
				Description: "The user group of the user. Cannot be updated.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"role_bindings": {
				Description: "Role Bindings of the user. Cannot be updated.",
				Type:        schema.TypeList,
				Optional:    true,
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

	var email = ""
	if attr, ok := d.GetOk("email"); ok {
		email = attr.(string)
	}

	resp, httpResp, err := c.UserApi.GetAggregatedUsers(ctx, c.AccountId, &nextgen.UserApiGetAggregatedUsersOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		SearchTerm:        optional.NewString(email),
	})

	if &resp == nil || resp.Data == nil || resp.Data.Empty {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readUserList(d, resp.Data)

	return nil
}

var creationSemaphore = make(chan struct{}, 1)

func resourceUserCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	creationSemaphore <- struct{}{}
	defer func() { <-creationSemaphore }()
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
		// Extract user_groups from request and call user group update API
		updateUserBody := createUpdateUserBody(d)
		userOpts := nextgen.UserApiAddUserToUserGroupsOpts{}
		body := nextgen.UserAddToUserGroupDto{}
		userOpts.OrgIdentifier = helpers.BuildField(d, "org_id")
		userOpts.ProjectIdentifier = helpers.BuildField(d, "project_id")
		body.UserGroupIdsToAdd = updateUserBody.UserGroups
		userOpts.Body = body
		_, httpResp, err = c.UserApi.AddUserToUserGroups(ctx, c.AccountId, d.Get("identifier").(string), &userOpts)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	var email = ""
	if attr, ok := d.GetOk("email"); ok {
		email = attr.(string)
	}

	resp, httpResp, err := c.UserApi.GetAggregatedUsers(ctx, c.AccountId, &nextgen.UserApiGetAggregatedUsersOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
		SearchTerm:        optional.NewString(email),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readUserList(d, resp.Data)

	return nil
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	creationSemaphore <- struct{}{}
	defer func() { <-creationSemaphore }()
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	uuid := d.Get("identifier").(string)
	orgIdentifier := optional.NewString(d.Get("org_id").(string))
	projectIdentifier := optional.NewString(d.Get("project_id").(string))
	var removeUserOpts = &nextgen.UserApiRemoveUserOpts{}

	if orgIdentifier.IsSet() && len(orgIdentifier.Value()) > 0 && projectIdentifier.IsSet() && len(projectIdentifier.Value()) > 0 {
		removeUserOpts = &nextgen.UserApiRemoveUserOpts{
			OrgIdentifier:     orgIdentifier,
			ProjectIdentifier: projectIdentifier,
		}
	} else if orgIdentifier.IsSet() && len(orgIdentifier.Value()) > 0 {
		removeUserOpts = &nextgen.UserApiRemoveUserOpts{
			OrgIdentifier: orgIdentifier,
		}
	}

	_, httpResp, err := c.UserApi.RemoveUser(ctx, uuid, c.AccountId, removeUserOpts)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func createAddUserBody(d *schema.ResourceData) *nextgen.AddUsersDto {

	var addUsersDto nextgen.AddUsersDto
	if attr, ok := d.GetOk("email"); ok {
		addUsersDto.Emails = []string{attr.(string)}
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
	} else {
		addUsersDto.RoleBindings = []nextgen.RoleBinding{}
	}

	return &addUsersDto
}

func createUpdateUserBody(d *schema.ResourceData) *nextgen.AddUsersDto {

	var addUsersDto nextgen.AddUsersDto

	if attr, ok := d.GetOk("user_groups"); ok {
		addUsersDto.UserGroups = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return &addUsersDto
}

func readUserList(d *schema.ResourceData, userInfo *nextgen.PageResponseUserAggregate) {
	userInfoList := userInfo.Content
	for _, value := range userInfoList {
		readUser(d, &value)
	}
}

func readUser(d *schema.ResourceData, UserAggregate *nextgen.UserAggregate) {
	d.SetId(UserAggregate.User.Email)
	d.Set("identifier", UserAggregate.User.Uuid)
	d.Set("name", UserAggregate.User.Name)
	d.Set("email", UserAggregate.User.Email)
	d.Set("locked", UserAggregate.User.Locked)
	d.Set("disabled", UserAggregate.User.Disabled)
	d.Set("externally_managed", UserAggregate.User.ExternallyManaged)
}
