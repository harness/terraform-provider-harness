package user

import (
	"context"
	"log"
	"strings"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceUser() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for creating a Harness user",

		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the user.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the user.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"email": {
				Description: "The email of the user.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
			},
			"group_ids": {
				Description: "The groups the user belongs to. This is only used during the creation of the user. The groups are not updated after the user is created. When using this option you should also set `lifecycle = { ignore_changes = [\"group_ids\"] }`.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"is_email_verified": {
				Description: "Flag indicating whether or not the users email has been verified.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_imported_from_identity_provider": {
				Description: "Flag indicating whether or not the user was imported from an identity provider.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_password_expired": {
				Description: "Flag indicating whether or not the users password has expired.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_two_factor_auth_enabled": {
				Description: "Flag indicating whether or not two-factor authentication is enabled for the user.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_user_locked": {
				Description: "Flag indicating whether or not the user is locked out.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				d.Set("email", d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}
	log.Printf("[DEBUG] Creating user %s", d.Get("email").(string))

	input := &graphql.CreateUserInput{
		Name:         d.Get("name").(string),
		Email:        d.Get("email").(string),
		UserGroupIds: utils.InterfaceSliceToStringSlice(d.Get("group_ids").(*schema.Set).List()),
	}

	user, err := c.UserClient.CreateUser(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readUser(d, user)
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}
	email := d.Get("email").(string)

	log.Printf("[DEBUG] Looking for user by email %s", email)

	user, err := c.UserClient.GetUserByEmail(email)
	if err != nil {
		return diag.FromErr(err)
	}

	if user == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readUser(d, user)
}

func readUser(d *schema.ResourceData, user *graphql.User) diag.Diagnostics {
	d.SetId(user.Id)
	d.Set("name", user.Name)
	d.Set("email", user.Email)
	d.Set("is_email_verified", user.IsEmailVerified)
	d.Set("is_imported_from_identity_provider", user.IsImportedFromIdentityProvider)
	d.Set("is_password_expired", user.IsPasswordExpired)
	d.Set("is_two_factor_auth_enabled", user.IsTwoFactorAuthenticationEnabled)
	d.Set("is_user_locked", user.IsUserLocked)

	return nil
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}
	log.Printf("[DEBUG] Updating user %s", d.Id())

	input := &graphql.UpdateUserInput{
		Name: d.Get("name").(string),
		Id:   d.Id(),
	}

	user, err := c.UserClient.UpdateUser(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readUser(d, user)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}
	if err := c.UserClient.DeleteUser(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
