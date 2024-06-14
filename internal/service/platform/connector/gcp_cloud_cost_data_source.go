package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceConnectorGCPCloudCost() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a GCP Cloud Cost Connector.",
		ReadContext: resourceConnectorGCPCloudCostRead,

		Schema: map[string]*schema.Schema{
			"features_enabled": {
				Description: "Indicates which features to enable among Billing, Optimization, and Visibility.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"gcp_project_id": {
				Description: "GCP Project Id.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"service_account_email": {
				Description: "Email corresponding to the Service Account.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"billing_export_spec": {
				Description: "Returns billing details.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_set_id": {
							Description: "Data Set Id.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"table_id": {
							Description: "Table Id.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
