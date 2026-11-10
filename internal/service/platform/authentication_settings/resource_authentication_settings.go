package authentication_settings

import (
	"context"
	"fmt"
	"net"
	"strings"

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
		// The resource manages the account the provider is authenticated with, so the
		// only valid import ID is that account's ID. AccountLevelResourceImporter is
		// not used here because it accepts any non-empty ID; the read would then
		// silently re-ID the resource to the provider's account.
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				accountId := meta.(*internal.Session).AccountId
				if d.Id() == "" {
					return nil, fmt.Errorf("invalid identifier: %s", d.Id())
				}
				if d.Id() != accountId {
					return nil, fmt.Errorf("cannot import harness_platform_authentication_settings with id %q: the resource manages the account the provider is authenticated with (%s); import using that account ID", d.Id(), accountId)
				}
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"two_factor_enabled": {
				Description: "Whether two-factor authentication is enforced for all users in the account.",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"whitelisted_domains": {
				Description: "List of whitelisted domains for the account. Domain names only - " +
					"IP addresses, CIDR ranges and URLs with a scheme prefix are rejected. " +
					"A `www.` prefix is allowed and is stripped by the API before storing; this " +
					"schema hashes entries the same way, so `www.example.com` in config matches " +
					"`example.com` returned by the API without a spurious diff. " +
					"When not set in the configuration, the current account value is kept (read from the API, like the Harness UI). " +
					"Set to an empty list to clear all whitelisted domains.",
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      hashNormalizedDomain,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validateWhitelistedDomain,
				},
			},
			"session_timeout_minutes": {
				Description: "Users are logged out after this many minutes of inactivity. Valid range: 30 to 4320. " +
					"When not set in the configuration, the current account value is kept (read from the API and sent back on updates, like the Harness UI).",
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(30, 4320),
			},
			"absolute_session_timeout_minutes": {
				Description: "Users are logged out after this many minutes regardless of activity. Maximum: 4320. " +
					"When not set in the configuration, the current account value is kept (read from the API and sent back on updates, like the Harness UI).",
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntAtMost(4320),
				RequiredWith: []string{"session_timeout_minutes"},
			},
			"public_access_enabled": {
				Description: "Whether public access is enabled for the account. " +
					"When not set in the configuration, the current account value is kept (read from the API, like the Harness UI).",
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"authentication_mechanism": {
				Description: "Authentication mechanism configured for the account. One of `USER_PASSWORD`, `SAML`, `LDAP`, `OAUTH`, `OIDC`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}

func resourceAuthenticationSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	// Read via the v2 API: the v1 list endpoint is rejected with a 400 on accounts
	// where the PL_ENABLE_MULTIPLE_IDP_SUPPORT feature flag is enabled, while v2
	// works regardless of the flag and returns the same response shape.
	resp, httpResp, err := c.AuthenticationSettingsApi.GetAuthenticationSettingsV2(ctx, c.AccountId)
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

	// Whitelisted domains and public access are written only when they change, so
	// configurations that do not declare them leave the current account values
	// untouched (the read API values are carried in state via Computed).
	if d.HasChange("whitelisted_domains") {
		if diags := updateWhitelistedDomains(ctx, c, expandStringSet(d.Get("whitelisted_domains").(*schema.Set))); diags != nil {
			return diags
		}
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

	if d.HasChange("public_access_enabled") {
		if diags := setPublicAccess(ctx, c, d.Get("public_access_enabled").(bool)); diags != nil {
			return diags
		}
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

	resp, httpResp, err := c.AuthenticationSettingsApi.SetTwoFactorAuth(ctx, req, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, nil, httpResp)
	}

	// The backend applies the change before returning, but only returns resource=true
	// when 2FA is enabled; disabling always yields resource=false (see
	// TwoFactorAuthenticationManager.overrideTwoFactorAuthentication). Only an
	// unsuccessful enable is a failure. The subsequent read verifies the state.
	if enabled && !resp.Resource {
		return diag.Errorf("failed to set two-factor authentication to %t", enabled)
	}

	return nil
}

func updateWhitelistedDomains(ctx context.Context, c *nextgen.APIClient, domains []string) diag.Diagnostics {
	if domains == nil {
		// The API expects a JSON array; send an empty one to clear all domains.
		domains = []string{}
	}

	resp, httpResp, err := c.AuthenticationSettingsApi.SetWhitelistedDomains(ctx, domains, c.AccountId)
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

func readAuthenticationSettings(d *schema.ResourceData, accountId string, settings *nextgen.AuthenticationSettingsV2Response) {
	// Singleton per account: the account identifier is the resource ID. Setting it here
	// covers the data source read as well, which has no prior ID.
	d.SetId(accountId)
	d.Set("two_factor_enabled", settings.TwoFactorEnabled)
	d.Set("whitelisted_domains", settings.WhitelistedDomains)
	d.Set("authentication_mechanism", settings.AuthenticationMechanism)
	d.Set("public_access_enabled", settings.PublicAccessEnabled)
	// The API returns null for an unset timeout, which decodes to 0. Write the raw
	// value instead of guarding on > 0: the fields are Computed, so a 0 for an
	// undeclared config produces no diff, while an out-of-band change from a
	// configured value to unset (e.g. the UI sends 0 to clear the absolute
	// timeout) is detected as drift instead of leaving a stale value in state.
	d.Set("session_timeout_minutes", settings.SessionTimeoutInMinutes)
	d.Set("absolute_session_timeout_minutes", settings.AbsoluteSessionTimeoutInMinutes)
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

// normalizeWhitelistedDomain mirrors the server-side canonicalization in
// AccountServiceImpl.updateWhitelistedDomains (scheme and "www." prefixes are
// stripped before the value is stored).
func normalizeWhitelistedDomain(domain string) string {
	domain = strings.ToLower(strings.TrimSpace(domain))
	domain = strings.TrimPrefix(domain, "http://")
	domain = strings.TrimPrefix(domain, "https://")
	domain = strings.TrimPrefix(domain, "www.")
	return strings.TrimSuffix(domain, "/")
}

// hashNormalizedDomain is the Set hash for whitelisted_domains. Hashing the
// normalized form makes an entry like "https://www.example.com" in config hash
// equal to "example.com" read back from the API, suppressing spurious diffs
// (TypeSet elements do not honor DiffSuppressFunc reliably).
func hashNormalizedDomain(v interface{}) int {
	return schema.HashString(normalizeWhitelistedDomain(v.(string)))
}

// validateWhitelistedDomain rejects values the API would fail at apply time
// (the backend validates with DomainValidator before sanitizing, see
// AccountServiceImpl.updateWhitelistedDomains): IP addresses, CIDR ranges and
// URLs with a scheme prefix. A "www." prefix is allowed - the API strips it.
// The scheme check runs on the raw value, before normalization: normalizing
// first would strip http(s):// and let scheme-prefixed URLs through to fail
// at apply time.
func validateWhitelistedDomain(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)
	if strings.Contains(value, "://") {
		errors = append(errors, fmt.Errorf("%q contains a scheme prefix; %s accepts domain names only (e.g. example.com or www.example.com)", value, k))
		return warnings, errors
	}
	normalized := normalizeWhitelistedDomain(value)
	if net.ParseIP(normalized) != nil {
		errors = append(errors, fmt.Errorf("%q is an IP address; %s accepts domain names only (e.g. example.com)", value, k))
	}
	if _, _, err := net.ParseCIDR(normalized); err == nil {
		errors = append(errors, fmt.Errorf("%q is a CIDR range; %s accepts domain names only (e.g. example.com)", value, k))
	}
	return warnings, errors
}
