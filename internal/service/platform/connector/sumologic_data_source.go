package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorSumologic() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Sumologic connector.",
		ReadContext: resourceConnectorSumologicRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the SumoLogic server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"access_id_ref": {
				Description: "Reference to the Harness secret containing the access id." + secret_ref_text,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"access_key_ref": {
				Description: "Reference to the Harness secret containing the access key." + secret_ref_text,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
