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
				Description: "URL of the Jira server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"username": {
				Description: "Username to use for authentication.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"username_ref": {
				Description: "Reference to a secret containing the username to use for authentication." + secret_ref_text,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"password_ref": {
				Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"auth": {
				Description: "The credentials to use for the jira authentication.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Description: "Authentication types for Jira connector",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"username_password": {
							Description: "Authenticate using username password.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Description: "Username to use for authentication.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"username_ref": {
										Description: "Reference to a secret containing the username to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"password_ref": {
										Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
