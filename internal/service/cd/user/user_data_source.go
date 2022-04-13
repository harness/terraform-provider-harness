package user

import (
	"context"
	"strings"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceUser() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving a Harness user",

		ReadContext: dataSourceUserRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description:   "Unique identifier of the user",
				Type:          schema.TypeString,
				Optional:      true,
				AtLeastOneOf:  []string{"id", "email"},
				ConflictsWith: []string{"email"},
			},
			"email": {
				Description:   "The email of the user.",
				Type:          schema.TypeString,
				Optional:      true,
				AtLeastOneOf:  []string{"id", "email"},
				ConflictsWith: []string{"id"},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new)
				},
			},
			"name": {
				Description: "The name of the user.",
				Type:        schema.TypeString,
				Computed:    true,
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
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	c := meta.(*internal.Session).CDClient

	var user *graphql.User
	var err error

	if id := d.Get("id").(string); id != "" {
		// Try lookup by Id first
		user, err = c.UserClient.GetUserById(id)
		if err != nil {
			return diag.FromErr(err)
		}
	} else if email := d.Get("email").(string); email != "" {
		// Fallback to lookup by email
		user, err = c.UserClient.GetUserByEmail(email)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		// Throw error if neither are set
		return diag.Errorf("id or name must be set")
	}

	d.SetId(user.Id)
	d.Set("email", user.Email)
	d.Set("name", user.Name)
	d.Set("is_email_verified", user.IsEmailVerified)
	d.Set("is_imported_from_identity_provider", user.IsImportedFromIdentityProvider)
	d.Set("is_password_expired", user.IsPasswordExpired)
	d.Set("is_two_factor_auth_enabled", user.IsTwoFactorAuthenticationEnabled)
	d.Set("is_user_locked", user.IsUserLocked)

	return nil
}
