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
			"name": {
				Description: "Name of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"email": {
				Description: "Email address of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"token": {
				Description: "Token used to authenticate the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"default_account_id": {
				Description: "Default account ID of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"intent": {
				Description: "Intent of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"admin": {
				Description: "Whether the user is an administrator.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_two_factor_auth_enabled": {
				Description: "Whether 2FA is enabled for the user.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"email_verified": {
				Description: "Whether the user's email address has been verified.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"locked": {
				Description: "Whether or not the user account is locked.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"signup_action": {
				Description: "Signup action of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"edition": {
				Description: "Edition of the platform being used.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"billing_frequency": {
				Description: "Billing frequency of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"org_id": {
				Description: "Organization identifier of the user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project identifier of the user.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"emails": {
				Description: "The email of the user.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"role_bindings": {
				Description: "Role Bindings of the user.",
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

	resp, httpResp, err := c.UserApi.GetCurrentUserInfo(ctx, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readUser(d, resp.Data)

	return nil
}

func resourceUserCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()
	addUserBody := createAddUserBody(d)

	var err error
	var httpResp *http.Response

	if id == "" {
		_, httpResp, err = c.UserApi.AddUsers(ctx, *addUserBody, c.AccountId, &nextgen.UserApiAddUsersOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
	} else {
		_, httpResp, err = c.UserApi.UpdateUserInfo(ctx, c.AccountId, &nextgen.UserApiUpdateUserInfoOpts{
			Body: optional.NewInterface(addUserBody),
		})
	}

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	resp, httpResp, err := c.UserApi.GetCurrentUserInfo(ctx, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readUser(d, resp.Data)

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

func readUser(d *schema.ResourceData, user *nextgen.UserInfo) {
	d.SetId(user.Uuid)
	d.Set("identifier", user.Uuid)
	d.Set("name", user.Name)
	d.Set("email", user.Email)
	d.Set("token", user.Token)
	d.Set("default_account_id", user.DefaultAccountId)
	d.Set("intent", user.Intent)
	d.Set("admin", user.Admin)
	d.Set("is_two_factor_auth_enabled", user.TwoFactorAuthenticationEnabled)
	d.Set("email_verified", user.EmailVerified)
	d.Set("locked", user.Locked)
	d.Set("signup_action", user.SignupAction)
	d.Set("edition", user.Edition)
	d.Set("billing_frequency", user.BillingFrequency)
}
