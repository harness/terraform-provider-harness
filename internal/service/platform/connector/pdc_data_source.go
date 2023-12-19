package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorPdc() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Pdc connector.",
		ReadContext: resourceConnectorPdcRead,

		Schema: map[string]*schema.Schema{
			"host": {
				Description: "Host of the Physical data centers.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": {
							Description: "hostname",
							Type:        schema.TypeString,
							Required:    true,
						},
						"attributes": {
							Description: "attributes for current host",
							Type:        schema.TypeMap,
							Optional:    true,
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
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
