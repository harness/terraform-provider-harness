package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorDynatrace() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Dynatrace connector.",
		ReadContext: resourceConnectorDynatraceRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Dynatrace server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"api_token_ref": {
				Description: "The reference to the Harness secret containing the api token." + secret_ref_text,
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
