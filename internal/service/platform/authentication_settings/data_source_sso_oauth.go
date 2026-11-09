package authentication_settings

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSSOOAuth() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving the OAuth SSO configuration of a Harness account.",

		ReadContext: resourceSSOOAuthRead,

		Schema: map[string]*schema.Schema{
			"allowed_providers": {
				Description: "OAuth providers users are allowed to sign in with.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"filter": {
				Description: "Filter applied to the available OAuth providers.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}
