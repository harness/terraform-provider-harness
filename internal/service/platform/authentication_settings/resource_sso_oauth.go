package authentication_settings

import (
	"context"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// validOAuthProviders mirrors io.harness.ng.core.account.OauthProviderType,
// excluding GOOGLE_CHAT which is not a login provider.
var validOAuthProviders = []string{"AZURE", "BITBUCKET", "GITHUB", "GITLAB", "GOOGLE", "LINKEDIN"}

// ResourceSSOOAuth manages the OAuth SSO configuration of a Harness account.
// OAuth settings are a singleton per account: "create" applies the desired
// configuration and "delete" removes the OAuth mechanism from the account.
// The resource ID is the account identifier.
func ResourceSSOOAuth() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing the OAuth SSO configuration of a Harness account, " +
			"i.e. the set of OAuth providers users are allowed to sign in with. " +
			"Destroying this resource removes the OAuth mechanism from the account. " +
			"Note: this resource only configures the allowed OAuth providers. Enabling OAuth " +
			"as the account's login mechanism (equivalent to the OAuth toggle in the UI) is " +
			"managed separately via the account authentication settings.",

		ReadContext:   resourceSSOOAuthRead,
		CreateContext: resourceSSOOAuthCreateOrUpdate,
		UpdateContext: resourceSSOOAuthCreateOrUpdate,
		DeleteContext: resourceSSOOAuthDelete,
		Importer:      helpers.AccountLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"allowed_providers": {
				Description: "OAuth providers users are allowed to sign in with. " +
					"Valid values: `AZURE`, `BITBUCKET`, `GITHUB`, `GITLAB`, `GOOGLE`, `LINKEDIN`.",
				Type:     schema.TypeSet,
				Required: true,
				// The API rejects an empty provider list ("At least one OAuth provider
				// must be selected.") and the UI blocks disabling the last provider;
				// fail at plan time instead.
				MinItems: 1,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice(validOAuthProviders, false),
				},
			},
			"filter": {
				Description: "Filter applied to the available OAuth providers.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	return resource
}

func resourceSSOOAuthRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.AuthenticationSettingsApi.GetAuthenticationSettingsV2(ctx, c.AccountId)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	settings := findOAuthSettings(resp.Resource)
	if settings == nil {
		if d.Id() == "" {
			// Data source read: report the account with no providers configured
			// instead of a "not found" (empty ID).
			d.SetId(c.AccountId)
			d.Set("allowed_providers", []string{})
			return nil
		}
		// Resource read: OAuth configuration was removed out of band.
		d.SetId("")
		return nil
	}

	// Singleton per account: the account identifier is the resource ID. Setting it
	// here covers the data source read as well, which has no prior ID.
	d.SetId(c.AccountId)
	d.Set("allowed_providers", settings.AllowedProviders)
	d.Set("filter", settings.Filter)

	return nil
}

func resourceSSOOAuthCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	settings := nextgen.OAuthSettings{
		SettingsType:     "OAUTH",
		Filter:           d.Get("filter").(string),
		AllowedProviders: expandOAuthProviders(d.Get("allowed_providers").(*schema.Set)),
	}

	resp, httpResp, err := c.AuthenticationSettingsApi.SetOauthProviders(ctx, settings, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if !resp.Resource {
		return diag.Errorf("failed to update OAuth providers")
	}

	// Singleton per account: the account identifier is the resource ID.
	d.SetId(c.AccountId)

	return resourceSSOOAuthRead(ctx, d, meta)
}

func resourceSSOOAuthDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	// UI parity: the backend rejects removing the OAuth mechanism while the
	// account's authentication mechanism is still OAUTH, so the UI first
	// switches the mechanism to USER_PASSWORD and only then deletes. Mirror
	// that two-step flow here.
	settings, httpResp, err := c.AuthenticationSettingsApi.GetAuthenticationSettingsV2(ctx, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if settings.Resource != nil && settings.Resource.AuthenticationMechanism == "OAUTH" {
		mechResp, httpResp, err := c.AuthenticationSettingsApi.UpdateAuthMechanism(ctx, c.AccountId,
			&nextgen.AuthenticationSettingsApiUpdateAuthMechanismOpts{
				AuthenticationMechanism: optional.NewString("USER_PASSWORD"),
			})
		if err != nil {
			return helpers.HandleApiError(err, d, httpResp)
		}
		if !mechResp.Resource {
			return diag.Errorf("failed to switch the account authentication mechanism to USER_PASSWORD before removing OAuth")
		}
	}

	resp, httpResp, err := c.AuthenticationSettingsApi.RemoveOauthMechanism(ctx, c.AccountId)
	if err != nil {
		// The backend returns 400 ("No Oauth settings found for this account.")
		// when no OAuth configuration exists; treat that as already deleted so
		// destroy stays idempotent.
		if httpResp != nil && httpResp.StatusCode == http.StatusBadRequest &&
			strings.Contains(err.Error(), "No Oauth settings found") {
			d.SetId("")
			return nil
		}
		return helpers.HandleApiError(err, d, httpResp)
	}

	if !resp.Resource {
		return diag.Errorf("failed to remove the OAuth mechanism")
	}

	d.SetId("")

	return nil
}

// expandOAuthProviders converts the allowed_providers set into a string slice.
// Kept local to this resource (rather than sharing a helper across the
// package) so each per-resource change set stays independently mergeable.
func expandOAuthProviders(set *schema.Set) []string {
	if set == nil || set.Len() == 0 {
		return nil
	}

	result := make([]string, 0, set.Len())
	for _, v := range set.List() {
		result = append(result, v.(string))
	}

	return result
}

// findOAuthSettings returns the OAUTH entry of the v2 authentication settings
// response, or nil if the account has no OAuth configuration.
func findOAuthSettings(resp *nextgen.AuthenticationSettingsV2Response) *nextgen.SamlSettings {
	if resp == nil {
		return nil
	}

	for i := range resp.NgAuthSettings {
		settings := &resp.NgAuthSettings[i]
		if settings.SettingsType == "OAUTH" {
			return settings
		}
	}

	return nil
}
