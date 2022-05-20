package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorPrometheus() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Prometheus connector.",
		ReadContext: resourceConnectorPrometheusRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "Url of the Prometheus server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
