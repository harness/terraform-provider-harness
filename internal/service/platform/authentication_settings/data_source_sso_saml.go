package authentication_settings

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSSOSAML() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a SAML SSO configuration of a Harness account. " +
			"The SAML metadata XML and the Azure client secret are not returned by the API.",

		ReadContext: dataSourceSSOSAMLRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the SAML SSO configuration.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"display_name": {
				Description: "Display name of the SAML SSO configuration.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"friendly_saml_name": {
				Description: "Friendly name of the app on the SAML SSO provider end in Harness.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"entity_identifier": {
				Description: "Entity identifier of the SAML service provider.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"origin": {
				Description: "Origin URL derived from the uploaded SAML metadata.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"logout_url": {
				Description: "Logout URL of the SAML identity provider.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"group_membership_attr": {
				Description: "Attribute in the SAML assertion containing the user's group memberships.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"authorization_enabled": {
				Description: "Whether group-based authorization is enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"authentication_enabled": {
				Description: "Whether login via this SAML configuration is enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"saml_provider_type": {
				Description: "SAML provider type.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"client_id": {
				Description: "Optional clientId for Azure SSO.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"jit_enabled": {
				Description: "Whether Just in time user provisioning is enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"jit_validation_key": {
				Description: "Optional key to match in the SAML assertion for Just in time user provisioning.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"jit_validation_value": {
				Description: "Optional value to match in the SAML assertion for Just in time user provisioning.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}
