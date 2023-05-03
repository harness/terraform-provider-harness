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
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"credentials_ref": {
				Description: "Reference to the secret containing credentials of IAM service account for Google Secret Manager." + secret_ref_text,
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
