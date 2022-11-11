package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorKubernetesCloudCost() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Kubernetes Cloud Cost connector.",
		ReadContext: resourceConnectorKubernetesCloudCostRead,

		Schema: map[string]*schema.Schema{
			"connector_ref": {
				Description: "Reference of the Connector.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"features_enabled": {
				Description: "Indicates which feature to enable among Billing, Optimization, and Visibility.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
