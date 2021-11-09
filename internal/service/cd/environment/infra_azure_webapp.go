package environment

import (
	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func infraDetailsAzureWebApp() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"resource_group": {
				Description: "The name of the resource group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"subscription_id": {
				Description: "The unique id of the azure subscription.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func expandAzureWebAppConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAzureWebApp{}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["resource_group"]; attr != nil {
		details.ResourceGroup = attr.(string)
	}

	if attr := config["subscription_id"]; attr != nil {
		details.SubscriptionId = attr.(string)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenAzureWebAppConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.AzureWebApp {
		return results
	}

	detailConfig := map[string]interface{}{}
	infraDetail := detail.ToAzureWebApp()

	detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
	detailConfig["resource_group"] = infraDetail.ResourceGroup
	detailConfig["subscription_id"] = infraDetail.SubscriptionId

	return append(results, detailConfig)
}
