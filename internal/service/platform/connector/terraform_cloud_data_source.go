package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorTerraformCloud() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for looking up a Terraform Cloud connector.",
		ReadContext: resourceConnectorTerraformCloudRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Terraform Cloud platform.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"credentials": {
				Description: "Credentials to connect to the Terraform Cloud platform.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_token": {
							Description: "API token credentials to use for authentication.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"api_token_ref": {
										Description: "Reference to a secret containing the API token to use for authentication." + secret_ref_text,
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

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
