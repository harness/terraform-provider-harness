package authentication_settings

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSSOLDAP() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving the LDAP SSO configuration of a Harness account.",

		ReadContext: resourceSSOLDAPRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the LDAP settings.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"display_name": {
				Description: "Display name of the LDAP settings.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cron_expression": {
				Description: "Cron expression for the LDAP user group sync job.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"disabled": {
				Description: "Whether the LDAP settings are disabled. LDAP sync does not run in the disabled state.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"next_iterations": {
				Description: "List of upcoming iterations (epoch millis) for the next LDAP sync job.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
			"connection_settings": {
				Description: "LDAP connection settings. The bind password is not returned by the API.",
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
							Description: "Whether SSL is enabled for the LDAP connection.",
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
						"bind_dn": {
							Description: "Bind distinguished name used to authenticate with the LDAP server.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password_type": {
							Description: "Type of the bind password: `INLINE` or `ENCRYPTED_SECRET`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"connect_timeout": {
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
					},
				},
			},
			"user_settings": {
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
			},
			"group_settings": {
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
							Description: "Attribute containing the group's members.",
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
			},
		},
	}

	return resource
}
