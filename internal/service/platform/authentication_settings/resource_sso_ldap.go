// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authentication_settings

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/harness/harness-go-sdk/harness/ldap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
)

// ResourceSSOLDAP manages the account-level LDAP SSO settings of a Harness account via the
// NextGen LDAP settings APIs (gated by the `PL_ENABLE_NG_LDAP_SETTINGS` feature flag).
//
// Exactly one LDAP configuration exists per account; every API endpoint resolves it from the
// account. The Terraform resource ID is therefore the account ID. The API requires a non-null
// identifier in the request body and never generates one server-side; when `identifier` is
// not configured it is derived from `name` (trimmed, spaces replaced with underscores),
// mirroring the Harness CG->NG migration convention, and it is immutable afterwards.
func ResourceSSOLDAP() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for managing account-level LDAP SSO settings via the NextGen LDAP APIs. " +
			"Requires the `PL_ENABLE_NG_LDAP_SETTINGS` feature flag to be enabled on the account. " +
			"Exactly one LDAP configuration exists per account.\n\n" +
			"The Terraform resource ID is the account ID. The API requires a non-null identifier in the request body " +
			"and never generates one server-side; when `identifier` is not configured it is derived from `name` " +
			"(trimmed, spaces replaced with underscores), and it is immutable afterwards. If `name` contains special " +
			"characters, set `identifier` explicitly. Destroying this resource deletes the LDAP configuration in " +
			"Harness; as in the UI, deletion is rejected while user groups are linked to the configuration.",
		ReadContext:   resourceSSOLDAPRead,
		CreateContext: resourceSSOLDAPCreate,
		UpdateContext: resourceSSOLDAPUpdate,
		DeleteContext: resourceSSOLDAPDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceSSOLDAPImport,
		},

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the LDAP SSO configuration. Immutable once created. When not set, it is derived from `name` (leading/trailing whitespace trimmed, spaces replaced with underscores) — the same convention the Harness backend uses when migrating legacy LDAP configurations to the NextGen store. If `name` contains special characters (anything other than letters, digits, spaces and underscores), set `identifier` explicitly to control the stored value.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"name": {
				Description: "Name of the LDAP SSO configuration. When `identifier` is not set, the identifier is derived from the name used at create time.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"disabled": {
				Description: "Whether the LDAP SSO configuration is disabled. LDAP sync and LDAP login do not run while disabled.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"cron_expression": {
				Description: "Cron expression for periodic synchronization of user groups from the LDAP directory.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "0 0 */1 ? * * *",
			},
			"connection_settings": ldapConnectionSettingsSchema(),
			"user_settings":       ldapUserSettingsSchema(),
			"group_settings":      ldapGroupSettingsSchema(),
			"url": {
				Description: "URL of the LDAP server (e.g. `ldaps://ldap.example.com:636`). Stored as-is; the server does not derive it.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
		},
	}

	return resource
}

func ldapConnectionSettingsSchema() *schema.Schema {
	return &schema.Schema{
		Description: "LDAP connection settings.",
		Type:        schema.TypeList,
		Required:    true,
		MinItems:    1,
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
					Description: "Whether SSL is used when connecting to the LDAP server.",
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
					Description: "Maximum number of referral hops to follow.",
					Type:        schema.TypeInt,
					Optional:    true,
					Default:     1,
				},
				"connection_timeout": {
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
				"bind_dn": {
					Description: "Distinguished name used to bind to the LDAP directory.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"secret_ref_path": {
					// This is a reference path (e.g. `account.my-secret`) to the secret holding the
					// bind password, not secret material itself, so it is intentionally not marked
					// Sensitive — mirroring the `*_ref` attributes of connector resources.
					Description: "Reference path to the account-level Harness secret containing the LDAP bind password (e.g. `account.my-secret`).",
					Type:        schema.TypeString,
					Required:    true,
				},
				"delegate_selectors": {
					Description: "Delegate selectors used to route LDAP connectivity.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}

func ldapUserSettingsSchema() *schema.Schema {
	return &schema.Schema{
		Description: "LDAP user query settings. At least one block is required.",
		Type:        schema.TypeList,
		Required:    true,
		MinItems:    1,
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
					Description: "Attribute containing the user's sAMAccountName (Active Directory).",
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
	}
}

func ldapGroupSettingsSchema() *schema.Schema {
	return &schema.Schema{
		Description: "LDAP group query settings.",
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
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
					Description: "Attribute on the group entry listing its members.",
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
	}
}

func resourceSSOLDAPCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLDAPClientWithContext(ctx)

	resp, httpResp, err := c.LdapSettingsApi.CreateLdapSettings(ctx, buildLdapSettingsDTO(d, meta))
	if err != nil {
		return handleLdapApiError(err, d, httpResp)
	}

	d.SetId(meta.(*internal.Session).AccountId)
	if resp.LdapSettings != nil {
		d.Set("identifier", resp.LdapSettings.Identifier)
	}

	return resourceSSOLDAPRead(ctx, d, meta)
}

func resourceSSOLDAPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLDAPClientWithContext(ctx)

	resp, httpResp, err := c.LdapSettingsApi.GetLdapSettings(ctx)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
		return handleLdapApiError(err, d, httpResp)
	}

	readLdapSettings(d, resp)
	return nil
}

func resourceSSOLDAPUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLDAPClientWithContext(ctx)

	// The {identifier} path parameter is ignored server-side (the configuration is resolved
	// from the account); it is passed only for forward compatibility.
	_, httpResp, err := c.LdapSettingsApi.UpdateLdapSettings(ctx, ldapSettingsIdentifier(d, meta), buildLdapSettingsDTO(d, meta))
	if err != nil {
		return handleLdapApiError(err, d, httpResp)
	}

	return resourceSSOLDAPRead(ctx, d, meta)
}

func resourceSSOLDAPDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLDAPClientWithContext(ctx)

	httpResp, err := c.LdapSettingsApi.DeleteLdapSettings(ctx, ldapSettingsIdentifier(d, meta))
	if err != nil {
		// A 404 means the account has no LDAP configuration (anymore); treat it as already deleted.
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		return handleLdapApiError(err, d, httpResp)
	}

	d.SetId("")
	return nil
}

func resourceSSOLDAPImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// Exactly one LDAP configuration exists per account, so the resource ID is the account ID.
	accountId := meta.(*internal.Session).AccountId
	if d.Id() != accountId {
		return nil, fmt.Errorf("LDAP SSO settings are account-level; import using the account ID (%s), got %q", accountId, d.Id())
	}
	return []*schema.ResourceData{d}, nil
}

func ldapSettingsIdentifier(d *schema.ResourceData, meta interface{}) string {
	if identifier := d.Get("identifier").(string); identifier != "" {
		return identifier
	}
	// The identifier path parameter is ignored server-side; fall back to the account ID so
	// the path segment is never empty.
	return meta.(*internal.Session).AccountId
}

func buildLdapSettingsDTO(d *schema.ResourceData, meta interface{}) ldap.LdapSettingsDTO {
	// The API validates identifier and accountIdentifier as non-null and never generates an
	// identifier server-side, so derive it from the name when not configured (mirroring the
	// Harness CG->NG migration convention: trimmed, spaces replaced with underscores).
	identifier := d.Get("identifier").(string)
	if identifier == "" {
		identifier = ldapIdentifierFromName(d.Get("name").(string))
	}
	return ldap.LdapSettingsDTO{
		AccountIdentifier:      meta.(*internal.Session).AccountId,
		Identifier:             identifier,
		Name:                   d.Get("name").(string),
		Url:                    d.Get("url").(string),
		Disabled:               d.Get("disabled").(bool),
		CronExpression:         d.Get("cron_expression").(string),
		LdapConnectionSettings: expandLdapConnectionSettings(d.Get("connection_settings").([]interface{})),
		LdapUserSettings:       expandLdapUserSettings(d.Get("user_settings").([]interface{})),
		LdapGroupSettings:      expandLdapGroupSettings(d.Get("group_settings").([]interface{})),
	}
}

func ldapIdentifierFromName(name string) string {
	return strings.ReplaceAll(strings.TrimSpace(name), " ", "_")
}

func expandLdapConnectionSettings(list []interface{}) *ldap.LdapConnectionSettingsDTO {
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	m := list[0].(map[string]interface{})
	referrals := m["referrals_enabled"].(bool)
	maxReferralHops := int32(m["max_referral_hops"].(int))
	recursive := m["use_recursive_group_membership_search"].(bool)
	return &ldap.LdapConnectionSettingsDTO{
		Host:                              m["host"].(string),
		Port:                              int32(m["port"].(int)),
		SslEnabled:                        m["ssl_enabled"].(bool),
		ReferralsEnabled:                  &referrals,
		MaxReferralHops:                   &maxReferralHops,
		ConnectionTimeout:                 int32(m["connection_timeout"].(int)),
		ResponseTimeout:                   int32(m["response_timeout"].(int)),
		UseRecursiveGroupMembershipSearch: &recursive,
		BindDN:                            m["bind_dn"].(string),
		SecretRefPath:                     m["secret_ref_path"].(string),
		DelegateSelectors:                 expandLdapStringSet(m["delegate_selectors"].(*schema.Set)),
	}
}

func expandLdapUserSettings(list []interface{}) []ldap.LdapUserSettingsDTO {
	settings := make([]ldap.LdapUserSettingsDTO, 0, len(list))
	for _, item := range list {
		if item == nil {
			continue
		}
		m := item.(map[string]interface{})
		settings = append(settings, ldap.LdapUserSettingsDTO{
			BaseDN:              m["base_dn"].(string),
			SearchFilter:        m["search_filter"].(string),
			UidAttr:             m["uid_attr"].(string),
			SamAccountNameAttr:  m["sam_account_name_attr"].(string),
			EmailAttr:           m["email_attr"].(string),
			DisplayNameAttr:     m["display_name_attr"].(string),
			GroupMembershipAttr: m["group_membership_attr"].(string),
		})
	}
	return settings
}

func expandLdapGroupSettings(list []interface{}) []ldap.LdapGroupSettingsDTO {
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	m := list[0].(map[string]interface{})
	return []ldap.LdapGroupSettingsDTO{
		{
			BaseDN:             m["base_dn"].(string),
			SearchFilter:       m["search_filter"].(string),
			NameAttr:           m["name_attr"].(string),
			DescriptionAttr:    m["description_attr"].(string),
			UserMembershipAttr: m["user_membership_attr"].(string),
			ReferencedUserAttr: m["referenced_user_attr"].(string),
		},
	}
}

func readLdapSettings(d *schema.ResourceData, resp ldap.LdapSettingsResponse) {
	settings := resp.LdapSettings
	if settings == nil {
		return
	}
	d.Set("identifier", settings.Identifier)
	d.Set("name", settings.Name)
	d.Set("disabled", settings.Disabled)
	d.Set("cron_expression", settings.CronExpression)
	d.Set("url", settings.Url)
	d.Set("connection_settings", flattenLdapConnectionSettings(settings.LdapConnectionSettings))
	d.Set("user_settings", flattenLdapUserSettings(settings.LdapUserSettings))
	d.Set("group_settings", flattenLdapGroupSettings(settings.LdapGroupSettings))
}

func flattenLdapConnectionSettings(cs *ldap.LdapConnectionSettingsDTO) []interface{} {
	if cs == nil {
		return nil
	}
	return []interface{}{
		map[string]interface{}{
			"host":                                  cs.Host,
			"port":                                  cs.Port,
			"ssl_enabled":                           cs.SslEnabled,
			"referrals_enabled":                     boolValue(cs.ReferralsEnabled),
			"max_referral_hops":                     intValue(cs.MaxReferralHops),
			"connection_timeout":                    cs.ConnectionTimeout,
			"response_timeout":                      cs.ResponseTimeout,
			"use_recursive_group_membership_search": boolValue(cs.UseRecursiveGroupMembershipSearch),
			"bind_dn":                               cs.BindDN,
			"secret_ref_path":                       cs.SecretRefPath,
			"delegate_selectors":                    flattenStringSlice(cs.DelegateSelectors),
		},
	}
}

func flattenLdapUserSettings(settings []ldap.LdapUserSettingsDTO) []interface{} {
	result := make([]interface{}, 0, len(settings))
	for _, s := range settings {
		result = append(result, map[string]interface{}{
			"base_dn":               s.BaseDN,
			"search_filter":         s.SearchFilter,
			"uid_attr":              s.UidAttr,
			"sam_account_name_attr": s.SamAccountNameAttr,
			"email_attr":            s.EmailAttr,
			"display_name_attr":     s.DisplayNameAttr,
			"group_membership_attr": s.GroupMembershipAttr,
		})
	}
	return result
}

func flattenLdapGroupSettings(settings []ldap.LdapGroupSettingsDTO) []interface{} {
	result := make([]interface{}, 0, len(settings))
	for _, s := range settings {
		result = append(result, map[string]interface{}{
			"base_dn":              s.BaseDN,
			"search_filter":        s.SearchFilter,
			"name_attr":            s.NameAttr,
			"description_attr":     s.DescriptionAttr,
			"user_membership_attr": s.UserMembershipAttr,
			"referenced_user_attr": s.ReferencedUserAttr,
		})
	}
	return result
}

func boolValue(b *bool) bool {
	return b != nil && *b
}

func intValue(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func expandLdapStringSet(set *schema.Set) []string {
	if set == nil {
		return nil
	}
	out := make([]string, 0, set.Len())
	for _, v := range set.List() {
		out = append(out, v.(string))
	}
	return out
}

func flattenStringSlice(in []string) []interface{} {
	out := make([]interface{}, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}

// handleLdapApiError converts the ldap client's GenericError into actionable diagnostics.
// The account-singleton API returns io.harness.eraro.ResponseMessage bodies on 4xx, which
// GenericError.Message() unwraps.
func handleLdapApiError(err error, d *schema.ResourceData, httpResp *http.Response) diag.Diagnostics {
	if ge, ok := err.(ldap.GenericError); ok && httpResp != nil {
		switch httpResp.StatusCode {
		case http.StatusBadRequest:
			return diag.Errorf("Bad Request: %s", ge.Message())
		case http.StatusUnauthorized:
			return diag.Errorf("%s\nHint:\n"+
				"1) Please check if the token has expired or is wrong.\n"+
				"2) Harness Provider is misconfigured. Please give the correct platform_api_key.", httpResp.Status)
		case http.StatusForbidden:
			return diag.Errorf("%s\nHint:\n"+
				"1) Please check if the token has the required permission for this operation.\n"+
				"2) Please check if the token has expired or is wrong.", httpResp.Status)
		case http.StatusNotFound:
			return diag.Errorf("LDAP SSO settings not found for the account: %s", ge.Message())
		}
	}
	return helpers.HandleApiError(err, d, httpResp)
}
