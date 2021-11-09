package ng

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCurrentUser() *schema.Resource {
	return &schema.Resource{
		Description: utils.GetNextgenDescription("Data source for retrieving the current user based on the API key."),

		ReadContext: dataSourceCurrentUserRead,

		Schema: map[string]*schema.Schema{
			"uuid": {
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
			"2fa_enabled": {
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
		},
	}
}

func dataSourceCurrentUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)
	resp, _, err := c.NGClient.UsersApi.GetUserInfo(ctx, c.AccountId)
	if err != nil {
		return diag.Errorf(err.(nextgen.GenericSwaggerError).Error())
	}

	user := resp.Data
	d.SetId(user.Uuid)
	d.Set("uuid", user.Uuid)
	d.Set("name", user.Name)
	d.Set("email", user.Email)
	d.Set("token", user.Token)
	d.Set("default_account_id", user.DefaultAccountId)
	d.Set("intent", user.Intent)
	d.Set("admin", user.Admin)
	d.Set("2fa_enabled", user.TwoFactorAuthenticationEnabled)
	d.Set("email_verified", user.EmailVerified)
	d.Set("locked", user.Locked)
	d.Set("signup_action", user.SignupAction)
	d.Set("edition", user.Edition)
	d.Set("billing_frequency", user.BillingFrequency)

	return nil
}
