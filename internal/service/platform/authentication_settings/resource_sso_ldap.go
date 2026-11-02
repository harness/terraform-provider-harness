package authentication_settings

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceSSOLDAP manages the LDAP SSO configuration of a Harness account.
// The resource ID is the LDAP settings identifier.
//
// Note: the GET /authentication-settings/ldap/settings API does not return the
// plaintext bind password; the value from the Terraform configuration is kept in
// state, and out-of-band changes to it are not detected by terraform plan.
func ResourceSSOLDAP() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing the LDAP SSO configuration of a Harness account. " +
			"The bind password can be provided inline or as a reference to a Harness secret " +
			"(e.g. `harness_platform_secret`), controlled by `password_type`.",

		ReadContext:   resourceSSOLDAPRead,
		CreateContext: resourceSSOLDAPCreate,
		UpdateContext: resourceSSOLDAPUpdate,
		DeleteContext: resourceSSOLDAPDelete,
		Importer:      helpers.AccountLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the LDAP settings.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"display_name": {
				Description: "Display name of the LDAP settings.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cron_expression": {
				Description: "Cron expression for the LDAP user group sync job.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"disabled": {
				Description: "Whether the LDAP settings are disabled. LDAP sync does not run in the disabled state.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"next_iterations": {
				Description: "List of upcoming iterations (epoch millis) for the next LDAP sync job.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"connection_settings": {
				Description: "LDAP connection settings.",
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Description: "Hostname of the LDAP server.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"port": {
							Description: "Port of the LDAP server.",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     636,
						},
						"ssl_enabled": {
							Description: "Whether SSL is enabled for the LDAP connection.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"referrals_enabled": {
							Description: "Whether LDAP referrals are followed.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
						},
						"max_referral_hops": {
							Description:  "Maximum number of referral hops to follow.",
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      1,
							ValidateFunc: validation.IntAtLeast(1),
						},
						"bind_dn": {
							Description: "Bind distinguished name used to authenticate with the LDAP server.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"bind_password": {
							Description: "Bind password, provided inline or as a Harness secret reference " +
								"depending on `password_type`. Not returned by the read API; the configured " +
								"value is kept in state.",
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"password_type": {
							Description:  "Type of the bind password: `INLINE` for a plaintext password or `ENCRYPTED_SECRET` for a Harness secret reference.",
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "INLINE",
							ValidateFunc: validation.StringInSlice([]string{"INLINE", "ENCRYPTED_SECRET"}, false),
						},
						"connect_timeout": {
							Description: "Connection timeout in milliseconds.",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5000,
						},
						"response_timeout": {
							Description: "Response timeout in milliseconds.",
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     5000,
						},
						"use_recursive_group_membership_search": {
							Description: "Whether recursive group membership search is used.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
						},
					},
				},
			},
			"user_settings": {
				Description: "LDAP user query settings.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base_dn": {
							Description: "Base distinguished name under which users are searched.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"search_filter": {
							Description: "LDAP search filter used to match users.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "(objectClass=person)",
						},
						"uid_attr": {
							Description: "Attribute containing the user's unique identifier.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "uid",
						},
						"sam_account_name_attr": {
							Description: "Attribute containing the user's sAMAccountName.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "sAMAccountName",
						},
						"email_attr": {
							Description: "Attribute containing the user's email address.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "mail",
						},
						"display_name_attr": {
							Description: "Attribute containing the user's display name.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "cn",
						},
						"group_membership_attr": {
							Description: "Attribute containing the user's group memberships.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "memberOf",
						},
					},
				},
			},
			"group_settings": {
				Description: "LDAP group query settings.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"base_dn": {
							Description: "Base distinguished name under which groups are searched.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"search_filter": {
							Description: "LDAP search filter used to match groups.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "(objectClass=group)",
						},
						"name_attr": {
							Description: "Attribute containing the group's name.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "cn",
						},
						"description_attr": {
							Description: "Attribute containing the group's description.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "description",
						},
						"user_membership_attr": {
							Description: "Attribute containing the group's members.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "member",
						},
						"referenced_user_attr": {
							Description: "Attribute on a member entry referencing the user.",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "dn",
						},
					},
				},
			},
		},
	}

	return resource
}

func resourceSSOLDAPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpResp, err := c.AuthenticationSettingsApi.GetLdapSettings(ctx, c.AccountId)
	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	if resp.Resource == nil || resp.Resource.Identifier == "" {
		d.SetId("")
		return nil
	}

	// For the data source the ID is not set yet; verify the identifier argument
	// matches the LDAP configuration returned for the account.
	if d.Id() == "" {
		if identifier := d.Get("identifier").(string); identifier != "" && identifier != resp.Resource.Identifier {
			return diag.Errorf("no LDAP SSO configuration found with identifier %q", identifier)
		}
	}

	// The data source schema has no bind_password field; only resources (and imports,
	// which have an ID by then) carry it in state.
	readSSOLDAP(d, resp.Resource, d.Id() != "")

	return nil
}

func resourceSSOLDAPCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	settings := buildLDAPSettings(d)

	_, httpResp, err := c.AuthenticationSettingsApi.CreateLdapSettings(ctx, settings, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	d.SetId(settings.Identifier)

	return resourceSSOLDAPRead(ctx, d, meta)
}

func resourceSSOLDAPUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	settings := buildLDAPSettings(d)

	_, httpResp, err := c.AuthenticationSettingsApi.UpdateLdapSettings(ctx, settings, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return resourceSSOLDAPRead(ctx, d, meta)
}

func resourceSSOLDAPDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	_, httpResp, err := c.AuthenticationSettingsApi.DeleteLdapSettings(ctx, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	d.SetId("")

	return nil
}

func readSSOLDAP(d *schema.ResourceData, settings *nextgen.LdapSettings, includeBindPassword bool) {
	d.SetId(settings.Identifier)
	d.Set("identifier", settings.Identifier)
	d.Set("display_name", settings.DisplayName)
	d.Set("cron_expression", settings.CronExpression)
	d.Set("disabled", settings.Disabled)
	d.Set("next_iterations", settings.NextIterations)
	d.Set("connection_settings", flattenLDAPConnectionSettings(settings.ConnectionSettings, d, includeBindPassword))
	d.Set("user_settings", flattenLDAPUserSettings(settings.UserSettingsList))
	d.Set("group_settings", flattenLDAPGroupSettings(settings.GroupSettingsList))
}

func buildLDAPSettings(d *schema.ResourceData) nextgen.LdapSettings {
	settings := nextgen.LdapSettings{
		SettingsType:       "LDAP",
		Identifier:         d.Get("identifier").(string),
		DisplayName:        d.Get("display_name").(string),
		CronExpression:     d.Get("cron_expression").(string),
		Disabled:           d.Get("disabled").(bool),
		ConnectionSettings: expandLDAPConnectionSettings(d.Get("connection_settings").([]interface{})),
		UserSettingsList:   expandLDAPUserSettings(d.Get("user_settings").([]interface{})),
		GroupSettingsList:  expandLDAPGroupSettings(d.Get("group_settings").([]interface{})),
	}
	return settings
}

func expandLDAPConnectionSettings(input []interface{}) *nextgen.LdapConnectionSettings {
	if len(input) == 0 || input[0] == nil {
		return nil
	}

	m := input[0].(map[string]interface{})

	return &nextgen.LdapConnectionSettings{
		Host:                              m["host"].(string),
		Port:                              int32(m["port"].(int)),
		SslEnabled:                        m["ssl_enabled"].(bool),
		ReferralsEnabled:                  m["referrals_enabled"].(bool),
		MaxReferralHops:                   int32(m["max_referral_hops"].(int)),
		BindDN:                            m["bind_dn"].(string),
		BindPassword:                      m["bind_password"].(string),
		PasswordType:                      m["password_type"].(string),
		ConnectTimeout:                    int32(m["connect_timeout"].(int)),
		ResponseTimeout:                   int32(m["response_timeout"].(int)),
		UseRecursiveGroupMembershipSearch: m["use_recursive_group_membership_search"].(bool),
	}
}

func expandLDAPUserSettings(input []interface{}) []nextgen.LdapUserSettings {
	if len(input) == 0 {
		return nil
	}

	result := make([]nextgen.LdapUserSettings, 0, len(input))
	for _, item := range input {
		m := item.(map[string]interface{})
		result = append(result, nextgen.LdapUserSettings{
			BaseDN:              m["base_dn"].(string),
			SearchFilter:        m["search_filter"].(string),
			UidAttr:             m["uid_attr"].(string),
			SamAccountNameAttr:  m["sam_account_name_attr"].(string),
			EmailAttr:           m["email_attr"].(string),
			DisplayNameAttr:     m["display_name_attr"].(string),
			GroupMembershipAttr: m["group_membership_attr"].(string),
		})
	}

	return result
}

func expandLDAPGroupSettings(input []interface{}) []nextgen.LdapGroupSettings {
	if len(input) == 0 {
		return nil
	}

	result := make([]nextgen.LdapGroupSettings, 0, len(input))
	for _, item := range input {
		m := item.(map[string]interface{})
		result = append(result, nextgen.LdapGroupSettings{
			BaseDN:             m["base_dn"].(string),
			SearchFilter:       m["search_filter"].(string),
			NameAttr:           m["name_attr"].(string),
			DescriptionAttr:    m["description_attr"].(string),
			UserMembershipAttr: m["user_membership_attr"].(string),
			ReferencedUserAttr: m["referenced_user_attr"].(string),
		})
	}

	return result
}

func flattenLDAPConnectionSettings(settings *nextgen.LdapConnectionSettings, d *schema.ResourceData, includeBindPassword bool) []interface{} {
	if settings == nil {
		return nil
	}

	m := map[string]interface{}{
		"host":                                  settings.Host,
		"port":                                  settings.Port,
		"ssl_enabled":                           settings.SslEnabled,
		"referrals_enabled":                     settings.ReferralsEnabled,
		"max_referral_hops":                     settings.MaxReferralHops,
		"bind_dn":                               settings.BindDN,
		"password_type":                         settings.PasswordType,
		"connect_timeout":                       settings.ConnectTimeout,
		"response_timeout":                      settings.ResponseTimeout,
		"use_recursive_group_membership_search": settings.UseRecursiveGroupMembershipSearch,
	}

	// The API does not return the plaintext bind password; keep the configured value.
	if includeBindPassword {
		if settings.BindPassword != "" {
			m["bind_password"] = settings.BindPassword
		} else if v, ok := d.GetOk("connection_settings.0.bind_password"); ok {
			m["bind_password"] = v.(string)
		}
	}

	return []interface{}{m}
}

func flattenLDAPUserSettings(input []nextgen.LdapUserSettings) []interface{} {
	if len(input) == 0 {
		return nil
	}

	result := make([]interface{}, 0, len(input))
	for _, settings := range input {
		result = append(result, map[string]interface{}{
			"base_dn":               settings.BaseDN,
			"search_filter":         settings.SearchFilter,
			"uid_attr":              settings.UidAttr,
			"sam_account_name_attr": settings.SamAccountNameAttr,
			"email_attr":            settings.EmailAttr,
			"display_name_attr":     settings.DisplayNameAttr,
			"group_membership_attr": settings.GroupMembershipAttr,
		})
	}

	return result
}

func flattenLDAPGroupSettings(input []nextgen.LdapGroupSettings) []interface{} {
	if len(input) == 0 {
		return nil
	}

	result := make([]interface{}, 0, len(input))
	for _, settings := range input {
		result = append(result, map[string]interface{}{
			"base_dn":              settings.BaseDN,
			"search_filter":        settings.SearchFilter,
			"name_attr":            settings.NameAttr,
			"description_attr":     settings.DescriptionAttr,
			"user_membership_attr": settings.UserMembershipAttr,
			"referenced_user_attr": settings.ReferencedUserAttr,
		})
	}

	return result
}
