package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceConnectorAzureCloudCost() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up an Azure Cloud Cost Connector.",
		ReadContext: resourceConnectorAzureCloudCostRead,

		Schema: map[string]*schema.Schema{
			"features_enabled": {
				Description: "Which feature to enable among BILLING, OPTIMIZATION, VISIBILITY",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tenant_id": {
				Description: "Tenant id.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"subscription_id": {
				Description: "Subsription id.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"billing_export_spec": {
				Description: "Returns Billing details like StorageAccount's Name, container's Name, directory's Name, report Name and subscription Id.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"storage_account_name": {
							Description: "Storage Account Name.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"container_name": {
							Description: "Container Name.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"directory_name": {
							Description: "Directory Name.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"report_name": {
							Description: "Report Name.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"subscription_id": {
							Description: "Subsription Id.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
