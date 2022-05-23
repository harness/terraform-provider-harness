package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorJira() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Jira connector.",
		ReadContext: resourceConnectorJiraRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "Url of the Jira server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"username": {
				Description: "Username to use for authentication.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"username_ref": {
				Description: "Reference to a secret containing the username to use for authentication.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"password_ref": {
				Description: "Reference to a secret containing the password to use for authentication.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
