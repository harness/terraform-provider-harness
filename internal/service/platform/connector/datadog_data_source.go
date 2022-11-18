package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorDatadog() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Datadog connector.",
		ReadContext: resourceConnectorDatadogRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Datadog server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"application_key_ref": {
				Description: "Reference to the Harness secret containing the application key." + secret_ref_text,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"api_key_ref": {
				Description: "Reference to the Harness secret containing the api key." + secret_ref_text,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
