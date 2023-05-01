package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorAppDynamics() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up an App Dynamics connector.",
		ReadContext: resourceConnectorAppDynamicsRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the App Dynamics controller.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account_name": {
				Description: "The App Dynamics account name.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"username_password": {
				Description: "Authenticate to App Dynamics using username and password.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Description: "Username to use for authentication.",
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
			"api_token": {
				Description: "Authenticate to App Dynamics using api token.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_secret_ref": {
							Description: "Reference to the Harness secret containing the App Dynamics client secret." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"client_id": {
							Description: "The client id used for connecting to App Dynamics.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
