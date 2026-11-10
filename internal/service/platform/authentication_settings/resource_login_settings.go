package authentication_settings

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceLoginSettings manages the login settings of a Harness account: the
// password strength policy, the user lockout policy and the password expiration
// policy. Login settings are a singleton per account: they always exist
// server-side, so "create" applies the desired configuration and "delete" resets
// the policies to the Harness defaults. The resource ID is the account identifier.
//
// The update API (PUT /ng/login-settings/{uuid}) requires the full login settings
// document, so create/update/delete all read the current settings first and send
// the merged document back; policy blocks not managed by this resource keep their
// existing values.
func ResourceLoginSettings() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing the login settings of a Harness account " +
			"(password strength policy, user lockout policy and password expiration policy). " +
			"Destroying this resource resets the policies to the Harness defaults.",

		ReadContext:   resourceLoginSettingsRead,
		CreateContext: resourceLoginSettingsCreateOrUpdate,
		UpdateContext: resourceLoginSettingsCreateOrUpdate,
		DeleteContext: resourceLoginSettingsDelete,
		Importer:      helpers.AccountLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"password_strength_policy": {
				Description: "Password strength policy of the account. If omitted, the existing policy is left untouched and its current values are readable from state (same behavior as the UI).",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Description: "Whether the password strength policy is enabled.",
							Type:        schema.TypeBool,
							Required:    true,
						},
						"min_number_of_characters": {
							Description: "Minimum number of characters required in a password.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"min_number_of_uppercase_characters": {
							Description: "Minimum number of uppercase characters required in a password.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"min_number_of_lowercase_characters": {
							Description: "Minimum number of lowercase characters required in a password.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"min_number_of_special_characters": {
							Description: "Minimum number of special characters required in a password.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
						"min_number_of_digits": {
							Description: "Minimum number of digits required in a password.",
							Type:        schema.TypeInt,
							Optional:    true,
						},
					},
				},
			},
			"user_lockout_policy": {
				Description: "User lockout policy of the account. If omitted, the existing policy is left untouched and its current values are readable from state (same behavior as the UI).",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_lockout_policy": {
							Description: "Whether the lockout policy is enabled.",
							Type:        schema.TypeBool,
							Required:    true,
						},
						"number_of_failed_attempts_before_lockout": {
							Description: "Number of failed login attempts before a user is locked out.",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
						},
						"lockout_period": {
							Description: "Lockout period in hours.",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     24,
						},
						"notify_user": {
							Description: "Whether the user is notified about the lockout.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
					},
				},
			},
			"password_expiration_policy": {
				Description: "Password expiration policy of the account. If omitted, the existing policy is left untouched and its current values are readable from state (same behavior as the UI).",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Description: "Whether the password expiration policy is enabled.",
							Type:        schema.TypeBool,
							Required:    true,
						},
						"days_before_password_expires": {
							Description: "Number of days before a password expires.",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     90,
						},
						"days_before_user_notified_of_password_expiration": {
							Description: "Number of days before expiration the user is notified.",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5,
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceLoginSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.AuthenticationSettingsApi.GetLoginSettings(ctx, c.AccountId)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Resource == nil {
		d.SetId("")
		return nil
	}

	readLoginSettings(d, c.AccountId, resp.Resource)

	return nil
}

func resourceLoginSettingsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	current, httpResp, err := c.AuthenticationSettingsApi.GetLoginSettings(ctx, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if current.Resource == nil {
		return diag.Errorf("no login settings found for account %s", c.AccountId)
	}

	// The update API requires the full login settings document; overlay the policy
	// blocks present in the configuration onto the current settings, so blocks not
	// managed by this resource keep their existing values.
	settings := buildLoginSettings(d, current.Resource)

	_, httpResp, err = c.AuthenticationSettingsApi.UpdateLoginSettings(ctx, settings, current.Resource.Uuid, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	// Singleton per account: the account identifier is the resource ID.
	d.SetId(c.AccountId)

	return resourceLoginSettingsRead(ctx, d, meta)
}

// resourceLoginSettingsDelete resets the login settings to the Harness defaults:
// password strength and expiration policies disabled, lockout policy disabled with
// the platform default thresholds. The settings themselves cannot be deleted.
func resourceLoginSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	current, httpResp, err := c.AuthenticationSettingsApi.GetLoginSettings(ctx, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if current.Resource == nil {
		// Nothing to reset.
		return nil
	}

	// userGroupsToNotify is not managed by this resource; preserve it on reset,
	// the same way the UI carries it over on every save of the lockout policy.
	var userGroupsToNotify []map[string]interface{}
	if current.Resource.UserLockoutPolicy != nil {
		userGroupsToNotify = current.Resource.UserLockoutPolicy.UserGroupsToNotify
	}

	defaults := nextgen.LoginSettings{
		Uuid:      current.Resource.Uuid,
		AccountId: c.AccountId,
		PasswordStrengthPolicy: &nextgen.PasswordStrengthPolicy{
			Enabled: false,
		},
		UserLockoutPolicy: &nextgen.UserLockoutPolicy{
			EnableLockoutPolicy:                 false,
			NumberOfFailedAttemptsBeforeLockout: 5,
			LockOutPeriod:                       24,
			NotifyUser:                          true,
			UserGroupsToNotify:                  userGroupsToNotify,
		},
		PasswordExpirationPolicy: &nextgen.PasswordExpirationPolicy{
			Enabled:                   false,
			DaysBeforePasswordExpires: 90,
			DaysBeforeUserNotifiedOfPasswordExpiration: 5,
		},
	}

	_, httpResp, err = c.AuthenticationSettingsApi.UpdateLoginSettings(ctx, defaults, current.Resource.Uuid, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

// buildLoginSettings merges the policy blocks present in the configuration onto
// the current login settings, preserving any block the resource does not manage.
func buildLoginSettings(d *schema.ResourceData, current *nextgen.LoginSettings) nextgen.LoginSettings {
	settings := nextgen.LoginSettings{
		Uuid:                     current.Uuid,
		AccountId:                current.AccountId,
		PasswordStrengthPolicy:   current.PasswordStrengthPolicy,
		UserLockoutPolicy:        current.UserLockoutPolicy,
		PasswordExpirationPolicy: current.PasswordExpirationPolicy,
	}

	if v, ok := d.GetOk("password_strength_policy"); ok {
		policy := expandPasswordStrengthPolicy(v.([]interface{}))
		settings.PasswordStrengthPolicy = &policy
	}

	if v, ok := d.GetOk("user_lockout_policy"); ok {
		policy := expandUserLockoutPolicy(v.([]interface{}))
		// userGroupsToNotify is not exposed in the schema (the API stores full
		// embedded UserGroup objects, manageable only via the legacy CG surface).
		// The write API replaces the whole embedded policy, so carry the groups
		// over from the current settings — the NG UI does the same on save.
		if current.UserLockoutPolicy != nil {
			policy.UserGroupsToNotify = current.UserLockoutPolicy.UserGroupsToNotify
		}
		settings.UserLockoutPolicy = &policy
	}

	if v, ok := d.GetOk("password_expiration_policy"); ok {
		policy := expandPasswordExpirationPolicy(v.([]interface{}))
		settings.PasswordExpirationPolicy = &policy
	}

	return settings
}

func readLoginSettings(d *schema.ResourceData, accountId string, settings *nextgen.LoginSettings) {
	// Singleton per account: the account identifier is the resource ID. Setting it
	// here covers the data source read as well, which has no prior ID.
	d.SetId(accountId)
	d.Set("password_strength_policy", flattenPasswordStrengthPolicy(settings.PasswordStrengthPolicy))
	d.Set("user_lockout_policy", flattenUserLockoutPolicy(settings.UserLockoutPolicy))
	d.Set("password_expiration_policy", flattenPasswordExpirationPolicy(settings.PasswordExpirationPolicy))
}

func expandPasswordStrengthPolicy(input []interface{}) nextgen.PasswordStrengthPolicy {
	if len(input) == 0 || input[0] == nil {
		return nextgen.PasswordStrengthPolicy{}
	}

	m := input[0].(map[string]interface{})

	return nextgen.PasswordStrengthPolicy{
		Enabled:                        m["enabled"].(bool),
		MinNumberOfCharacters:          int32(m["min_number_of_characters"].(int)),
		MinNumberOfUppercaseCharacters: int32(m["min_number_of_uppercase_characters"].(int)),
		MinNumberOfLowercaseCharacters: int32(m["min_number_of_lowercase_characters"].(int)),
		MinNumberOfSpecialCharacters:   int32(m["min_number_of_special_characters"].(int)),
		MinNumberOfDigits:              int32(m["min_number_of_digits"].(int)),
	}
}

func expandUserLockoutPolicy(input []interface{}) nextgen.UserLockoutPolicy {
	if len(input) == 0 || input[0] == nil {
		return nextgen.UserLockoutPolicy{}
	}

	m := input[0].(map[string]interface{})

	return nextgen.UserLockoutPolicy{
		EnableLockoutPolicy:                 m["enable_lockout_policy"].(bool),
		NumberOfFailedAttemptsBeforeLockout: int32(m["number_of_failed_attempts_before_lockout"].(int)),
		LockOutPeriod:                       int32(m["lockout_period"].(int)),
		NotifyUser:                          m["notify_user"].(bool),
	}
}

func expandPasswordExpirationPolicy(input []interface{}) nextgen.PasswordExpirationPolicy {
	if len(input) == 0 || input[0] == nil {
		return nextgen.PasswordExpirationPolicy{}
	}

	m := input[0].(map[string]interface{})

	return nextgen.PasswordExpirationPolicy{
		Enabled:                   m["enabled"].(bool),
		DaysBeforePasswordExpires: int32(m["days_before_password_expires"].(int)),
		DaysBeforeUserNotifiedOfPasswordExpiration: int32(m["days_before_user_notified_of_password_expiration"].(int)),
	}
}

func flattenPasswordStrengthPolicy(policy *nextgen.PasswordStrengthPolicy) []interface{} {
	if policy == nil {
		return nil
	}

	return []interface{}{map[string]interface{}{
		"enabled":                            policy.Enabled,
		"min_number_of_characters":           policy.MinNumberOfCharacters,
		"min_number_of_uppercase_characters": policy.MinNumberOfUppercaseCharacters,
		"min_number_of_lowercase_characters": policy.MinNumberOfLowercaseCharacters,
		"min_number_of_special_characters":   policy.MinNumberOfSpecialCharacters,
		"min_number_of_digits":               policy.MinNumberOfDigits,
	}}
}

func flattenUserLockoutPolicy(policy *nextgen.UserLockoutPolicy) []interface{} {
	if policy == nil {
		return nil
	}

	return []interface{}{map[string]interface{}{
		"enable_lockout_policy":                    policy.EnableLockoutPolicy,
		"number_of_failed_attempts_before_lockout": policy.NumberOfFailedAttemptsBeforeLockout,
		"lockout_period":                           policy.LockOutPeriod,
		"notify_user":                              policy.NotifyUser,
	}}
}

func flattenPasswordExpirationPolicy(policy *nextgen.PasswordExpirationPolicy) []interface{} {
	if policy == nil {
		return nil
	}

	return []interface{}{map[string]interface{}{
		"enabled":                      policy.Enabled,
		"days_before_password_expires": policy.DaysBeforePasswordExpires,
		"days_before_user_notified_of_password_expiration": policy.DaysBeforeUserNotifiedOfPasswordExpiration,
	}}
}
