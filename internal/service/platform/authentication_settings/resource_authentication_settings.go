package authentication_settings

import (
	"context"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceAuthenticationSettings manages account-level authentication settings.
// Authentication settings are a singleton per account: they always exist server-side,
// so "create" applies the desired configuration and "delete" resets the settings to
// the Harness defaults. The resource ID is the account identifier.
//
// The GET /authentication-settings API returns all readable fields, including
// two_factor_enabled, whitelisted_domains, authentication_mechanism,
// session_timeout_minutes, absolute_session_timeout_minutes and public_access_enabled,
// so out-of-band changes are detected by terraform plan.
func ResourceAuthenticationSettings() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing account level authentication settings. " +
			"Destroying this resource resets two-factor authentication, whitelisted domains and " +
			"public access to the Harness defaults. Session timeout settings are left unchanged, " +
			"since the API does not expose a way to restore the platform default.",
		ReadContext:   resourceAuthenticationSettingsRead,
		CreateContext: resourceAuthenticationSettingsCreateOrUpdate,
		UpdateContext: resourceAuthenticationSettingsCreateOrUpdate,
		DeleteContext: resourceAuthenticationSettingsDelete,
		Importer:      helpers.AccountLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"two_factor_enabled": {
				Description: "Whether two-factor authentication is enforced for all users in the account.",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"whitelisted_domains": {
				Description: "List of whitelisted domains for the account.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"session_timeout_minutes": {
				Description:  "Users are logged out after this many minutes of inactivity. Valid range: 30 to 4320.",
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(30, 4320),
			},
			"absolute_session_timeout_minutes": {
				Description:  "Users are logged out after this many minutes regardless of activity. Maximum: 4320.",
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntAtMost(4320),
				RequiredWith: []string{"session_timeout_minutes"},
			},
			"public_access_enabled": {
				Description: "Whether public access is enabled for the account.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"authentication_mechanism": {
				Description: "Authentication mechanism configured for the account, e.g. `SSO` or `NON_SSO`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}

func resourceAuthenticationSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.AuthenticationSettingsApi.GetAuthenticationSettings(ctx, c.AccountId)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Resource == nil {
		d.SetId("")
		return nil
	}

	readAuthenticationSettings(d, c.AccountId, resp.Resource)

	return nil
}

func resourceAuthenticationSettingsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	if diags := setTwoFactorAuth(ctx, c, d.Get("two_factor_enabled").(bool)); diags != nil {
		return diags
	}

	if diags := updateWhitelistedDomains(ctx, c, expandStringSet(d.Get("whitelisted_domains").(*schema.Set))); diags != nil {
		return diags
	}

	if v, ok := d.GetOk("session_timeout_minutes"); ok {
		settings := nextgen.SessionTimeoutSettings{
			SessionTimeOutInMinutes: int32(v.(int)),
		}
		if abs, ok := d.GetOk("absolute_session_timeout_minutes"); ok {
			settings.AbsoluteSessionTimeOutInMinutes = int32(abs.(int))
		}
		if diags := setSessionTimeout(ctx, c, settings); diags != nil {
			return diags
		}
	}

	if diags := setPublicAccess(ctx, c, d.Get("public_access_enabled").(bool)); diags != nil {
		return diags
	}

	// Singleton per account: the account identifier is the resource ID.
	d.SetId(c.AccountId)

	return resourceAuthenticationSettingsRead(ctx, d, meta)
}

// resourceAuthenticationSettingsDelete resets the account authentication settings to the
// Harness defaults: two-factor authentication disabled, whitelisted domains cleared and
// public access disabled. The settings themselves cannot be deleted. Session timeout
// settings are left unchanged because the API does not expose the platform default.
func resourceAuthenticationSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	if diags := setTwoFactorAuth(ctx, c, false); diags != nil {
		return diags
	}

	if diags := updateWhitelistedDomains(ctx, c, nil); diags != nil {
		return diags
	}

	return setPublicAccess(ctx, c, false)
}

func setTwoFactorAuth(ctx context.Context, c *nextgen.APIClient, enabled bool) diag.Diagnostics {
	req := nextgen.TwoFactorAdminOverrideSettings{
		AdminOverrideTwoFactorEnabled: enabled,
	}

	resp, httpResp, err := c.AuthenticationSettingsApi.SetTwoFactorAuthAtAccountLevel(ctx, req, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, nil, httpResp)
	}

	if !resp.Resource {
		return diag.Errorf("failed to set two-factor authentication to %t", enabled)
	}

	return nil
}

func updateWhitelistedDomains(ctx context.Context, c *nextgen.APIClient, domains []string) diag.Diagnostics {
	opts := &nextgen.AuthenticationSettingsApiUpdateWhitelistedDomainsOpts{
		Body: optional.NewInterface(domains),
	}

	resp, httpResp, err := c.AuthenticationSettingsApi.UpdateWhitelistedDomains(ctx, c.AccountId, opts)
	if err != nil {
		return helpers.HandleApiError(err, nil, httpResp)
	}

	if !resp.Resource {
		return diag.Errorf("failed to update whitelisted domains")
	}

	return nil
}

func setSessionTimeout(ctx context.Context, c *nextgen.APIClient, settings nextgen.SessionTimeoutSettings) diag.Diagnostics {
	resp, httpResp, err := c.AuthenticationSettingsApi.SetSessionTimeoutAtAccountLevel(ctx, settings, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, nil, httpResp)
	}

	if !resp.Resource {
		return diag.Errorf("failed to set session timeout")
	}

	return nil
}

func setPublicAccess(ctx context.Context, c *nextgen.APIClient, enabled bool) diag.Diagnostics {
	resp, httpResp, err := c.AuthenticationSettingsApi.SetPublicAccess(ctx, enabled, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, nil, httpResp)
	}

	if !resp.Resource {
		return diag.Errorf("failed to set public access to %t", enabled)
	}

	return nil
}

func readAuthenticationSettings(d *schema.ResourceData, accountId string, settings *nextgen.AuthenticationSettingsResponse) {
	// Singleton per account: the account identifier is the resource ID. Setting it here
	// covers the data source read as well, which has no prior ID.
	d.SetId(accountId)
	d.Set("two_factor_enabled", settings.TwoFactorEnabled)
	d.Set("whitelisted_domains", settings.WhitelistedDomains)
	d.Set("authentication_mechanism", settings.AuthenticationMechanism)
	d.Set("public_access_enabled", settings.PublicAccessEnabled)
	if settings.SessionTimeoutInMinutes > 0 {
		d.Set("session_timeout_minutes", settings.SessionTimeoutInMinutes)
	}
	if settings.AbsoluteSessionTimeoutInMinutes > 0 {
		d.Set("absolute_session_timeout_minutes", settings.AbsoluteSessionTimeoutInMinutes)
	}
}

func expandStringSet(set *schema.Set) []string {
	if set == nil || set.Len() == 0 {
		return nil
	}

	result := make([]string, 0, set.Len())
	for _, v := range set.List() {
		result = append(result, v.(string))
	}

	return result
}
