package authentication_settings

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceAuthenticationSettings() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving the account level authentication settings.",

		ReadContext: resourceAuthenticationSettingsRead,
		Schema: map[string]*schema.Schema{
			"two_factor_enabled": {
				Description: "Whether two-factor authentication is enforced for all users in the account.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"whitelisted_domains": {
				Description: "List of whitelisted domains for the account.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"authentication_mechanism": {
				Description: "Authentication mechanism configured for the account. One of `USER_PASSWORD`, `SAML`, `LDAP`, `OAUTH`, `OIDC`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"session_timeout_minutes": {
				Description: "Users are logged out after this many minutes of inactivity.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"absolute_session_timeout_minutes": {
				Description: "Users are logged out after this many minutes regardless of activity.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"public_access_enabled": {
				Description: "Whether public access is enabled for the account.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}

	return resource
}
