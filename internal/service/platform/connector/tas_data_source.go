package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceConnectorTas() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up an Tas Connector.",
		ReadContext: resourceConnectorTasRead,

		Schema: map[string]*schema.Schema{
			"credentials": {
				Description: "Contains Tas connector credentials.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description: "Type can be ManualConfig.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tas_manual_details": {
							Description: "Authenticate to Tas using manual details.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"endpoint_url": {
										Description: "URL of the Tas server.",
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
										Description: "Reference of the secret for the password.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
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
			"execute_on_delegate": {
				Description: "Execute on delegate or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
