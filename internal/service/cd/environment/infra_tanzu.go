package environment

import (
	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func infraDetailsTanzu() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"organization": {
				Description: "The PCF organization to use.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"space": {
				Description: "The PCF space to deploy to.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func expandTanzuConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureTanzu{}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["organization"]; attr != nil {
		details.Organization = attr.(string)
	}

	if attr := config["space"]; attr != nil {
		details.Space = attr.(string)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenTanzuConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.Pcf {
		return results
	}

	detailConfig := map[string]interface{}{}
	infraDetail := detail.ToPcf()

	detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
	detailConfig["organization"] = infraDetail.Organization
	detailConfig["space"] = infraDetail.Space

	return append(results, detailConfig)
}
