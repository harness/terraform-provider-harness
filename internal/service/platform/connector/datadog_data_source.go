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
				Description: "Url of the Datadog server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"application_key_ref": {
				Description: "Reference to the Harness secret containing the application key.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"api_key_ref": {
				Description: "Reference to the Harness secret containing the api key.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
