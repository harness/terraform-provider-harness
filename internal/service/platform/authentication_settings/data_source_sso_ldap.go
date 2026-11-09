// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package authentication_settings

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/harness/terraform-provider-harness/internal"
)

// DataSourceSSOLDAP reads the account-level LDAP SSO settings of a Harness account via the
// NextGen LDAP settings APIs (gated by the `PL_ENABLE_NG_LDAP_SETTINGS` feature flag).
// Exactly one LDAP configuration exists per account, so the data source takes no arguments.
func DataSourceSSOLDAP() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving the account-level LDAP SSO settings via the NextGen LDAP APIs. " +
			"Requires the `PL_ENABLE_NG_LDAP_SETTINGS` feature flag to be enabled on the account.",
		ReadContext: dataSourceSSOLDAPRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Identifier of the LDAP SSO configuration.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the LDAP SSO configuration.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"disabled": {
				Description: "Whether the LDAP SSO configuration is disabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"cron_expression": {
				Description: "Cron expression for periodic synchronization of user groups from the LDAP directory.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"connection_settings": ldapConnectionSettingsDataSourceSchema(),
			"user_settings":       ldapUserSettingsDataSourceSchema(),
			"group_settings":      ldapGroupSettingsDataSourceSchema(),
			"url": {
				Description: "URL of the LDAP server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}

func ldapConnectionSettingsDataSourceSchema() *schema.Schema {
	return &schema.Schema{
		Description: "LDAP connection settings.",
		Type:        schema.TypeList,
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"host": {
					Description: "Hostname of the LDAP server.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"port": {
					Description: "Port of the LDAP server.",
					Type:        schema.TypeInt,
					Computed:    true,
				},
				"ssl_enabled": {
					Description: "Whether SSL is used when connecting to the LDAP server.",
					Type:        schema.TypeBool,
					Computed:    true,
				},
				"referrals_enabled": {
					Description: "Whether LDAP referrals are followed.",
					Type:        schema.TypeBool,
					Computed:    true,
				},
				"max_referral_hops": {
					Description: "Maximum number of referral hops to follow.",
					Type:        schema.TypeInt,
					Computed:    true,
				},
				"connection_timeout": {
					Description: "Connection timeout in milliseconds.",
					Type:        schema.TypeInt,
					Computed:    true,
				},
				"response_timeout": {
					Description: "Response timeout in milliseconds.",
					Type:        schema.TypeInt,
					Computed:    true,
				},
				"use_recursive_group_membership_search": {
					Description: "Whether recursive group membership search is used.",
					Type:        schema.TypeBool,
					Computed:    true,
				},
				"bind_dn": {
					Description: "Distinguished name used to bind to the LDAP directory.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"secret_ref_path": {
					Description: "Reference path to the account-level Harness secret containing the LDAP bind password.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"delegate_selectors": {
					Description: "Delegate selectors used to route LDAP connectivity.",
					Type:        schema.TypeSet,
					Computed:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}

func ldapUserSettingsDataSourceSchema() *schema.Schema {
	return &schema.Schema{
		Description: "LDAP user query settings.",
		Type:        schema.TypeList,
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"base_dn": {
					Description: "Base distinguished name under which users are searched.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"search_filter": {
					Description: "LDAP search filter used to match users.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"uid_attr": {
					Description: "Attribute containing the user's unique identifier.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"sam_account_name_attr": {
					Description: "Attribute containing the user's sAMAccountName.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"email_attr": {
					Description: "Attribute containing the user's email address.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"display_name_attr": {
					Description: "Attribute containing the user's display name.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"group_membership_attr": {
					Description: "Attribute containing the user's group memberships.",
					Type:        schema.TypeString,
					Computed:    true,
				},
			},
		},
	}
}

func ldapGroupSettingsDataSourceSchema() *schema.Schema {
	return &schema.Schema{
		Description: "LDAP group query settings.",
		Type:        schema.TypeList,
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"base_dn": {
					Description: "Base distinguished name under which groups are searched.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"search_filter": {
					Description: "LDAP search filter used to match groups.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"name_attr": {
					Description: "Attribute containing the group's name.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"description_attr": {
					Description: "Attribute containing the group's description.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"user_membership_attr": {
					Description: "Attribute on the group entry listing its members.",
					Type:        schema.TypeString,
					Computed:    true,
				},
				"referenced_user_attr": {
					Description: "Attribute on a member entry referencing the user.",
					Type:        schema.TypeString,
					Computed:    true,
				},
			},
		},
	}
}

func dataSourceSSOLDAPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetLDAPClientWithContext(ctx)

	resp, httpResp, err := c.LdapSettingsApi.GetLdapSettings(ctx)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusNotFound {
			return diag.Errorf("no LDAP SSO configuration found for the account")
		}
		return handleLdapApiError(err, d, httpResp)
	}

	d.SetId(meta.(*internal.Session).AccountId)
	readLdapSettings(d, resp)
	return nil
}
