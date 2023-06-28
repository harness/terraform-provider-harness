package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorRancher() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Rancher connector.",
		ReadContext: resourceConnectorRancherRead,

		Schema: map[string]*schema.Schema{
			"delegate_selectors": {
				Description: "Selectors to use for the delegate.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"bearer_token": {
				Description: "URL and bearer token for the rancher cluster.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rancher_url": {
							Description: "The URL of the Rancher cluster.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"password_ref": {
							Description: "Reference to the secret containing the bearer token for the rancher cluster." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)
	return resource
}
