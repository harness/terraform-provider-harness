package environment

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func infraDetailsAwsWinRM() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"autoscaling_group_name": {
				Description: "The name of the autoscaling group.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"desired_capacity": {
				Description: "The desired capacity of the autoscaling group.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"host_connection_attrs_name": {
				Description: "The name of the host connection attributes to use.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"host_connection_type": {
				Description: fmt.Sprintf("The type of host connection to use. Valid options are %s", strings.Join(cac.HostConnectionTypesSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"hostname_convention": {
				Description: fmt.Sprintf("The naming convention to use for the hostname. Defaults to %s", DefaultHostnameConvention),
				Type:        schema.TypeString,
				Optional:    true,
				Default:     DefaultHostnameConvention,
			},
			"loadbalancer_name": {
				Description: "The name of the load balancer to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"region": {
				Description: "The region to deploy to.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func expandAwsWinRMConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAwsWinRM{}

	if attr := config["autoscaling_group_name"]; attr != nil {
		details.AutoscalingGroupName = attr.(string)
		details.UseAutoScalingGroup = true
	}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["desired_capacity"]; attr != "" {
		details.DesiredCapacity = attr.(int)
		if details.DesiredCapacity > 0 {
			details.SetDesiredCapacity = true
		}
	}

	if attr := config["host_connection_attrs_name"]; attr != "" {
		details.HostConnectionAttrsName = attr.(string)
	}

	if attr := config["host_connection_type"]; attr != "" {
		details.HostConnectionType = cac.HostConnectionType(attr.(string))
	}

	if attr := config["hostname_convention"]; attr != "" {
		details.HostNameConvention = attr.(string)
	}

	if attr := config["loadbalancer_name"]; attr != "" {
		details.LoadBalancerName = attr.(string)
	}

	if attr := config["region"]; attr != "" {
		details.Region = attr.(string)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenAwsWinRMConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type == cac.InfrastructureTypes.AwsSSH && infraDef.DeploymentType == cac.DeploymentTypes.WinRM {
		detailConfig := map[string]interface{}{}
		infraDetail := detail.ToAwsWinRm()

		detailConfig["autoscaling_group_name"] = infraDetail.AutoscalingGroupName
		detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
		detailConfig["desired_capacity"] = infraDetail.DesiredCapacity
		detailConfig["host_connection_attrs_name"] = infraDetail.HostConnectionAttrsName
		detailConfig["host_connection_type"] = infraDetail.HostConnectionType
		detailConfig["hostname_convention"] = infraDetail.HostNameConvention
		detailConfig["loadbalancer_name"] = infraDetail.LoadBalancerName
		detailConfig["region"] = infraDetail.Region

		results = append(results, detailConfig)
	}

	return results
}
