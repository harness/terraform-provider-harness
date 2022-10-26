package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorGcpSM() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up GCP Secret Manager connector.",
		ReadContext: resourceConnectorGcpSMRead,

		Schema: map[string]*schema.Schema{
			"is_default": {
				Description: "Indicative if this is default Secret manager for secrets.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"credentials_ref": {
				Description: "Reference to the secret containing credentials of IAM service account for Google Secret Manager.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
