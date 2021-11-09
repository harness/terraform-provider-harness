package environment

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func infraDetailsAwsSSH() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tag": {
				Description: "The tags to use when selecting the instances.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Description: "The key of the tag.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"value": {
							Description: "The value of the tag.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"vpc_ids": {
				Description: "The VPC ids to use when selecting the instances.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"autoscaling_group_name": {
				Description: "The name of the autoscaling group.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"desired_capacity": {
				Description: "The desired capacity of the auto scaling group.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"host_connection_attrs_name": {
				Description: "The name of the host connection attributes to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"host_connection_type": {
				Description: fmt.Sprintf("The type of host connection to use. Valid options are %s", strings.Join(cac.HostConnectionTypesSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"loadbalancer_name": {
				Description: "The name of the load balancer to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hostname_convention": {
				Description: fmt.Sprintf("The naming convention to use for the hostname. Defaults to %s", DefaultHostnameConvention),
				Type:        schema.TypeString,
				Optional:    true,
				Default:     DefaultHostnameConvention,
			},
			"region": {
				Description: "The region to deploy to.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func expandAwsSSHConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAwsSSH{}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["autoscaling_group_name"]; attr != "" {
		details.AutoscalingGroupName = attr.(string)
		details.UseAutoScalingGroup = true
	}

	if attr := config["desired_capacity"]; attr != nil {
		details.DesiredCapacity = attr.(int)
		details.SetDesiredCapacity = details.DesiredCapacity > 0
	}

	if attr := config["host_connection_attrs_name"]; attr != nil {
		details.HostConnectionAttrsName = attr.(string)
	}

	if attr := config["host_connection_type"]; attr != nil {
		details.HostConnectionType = cac.HostConnectionType(attr.(string))
	}

	if attr := config["loadbalancer_name"]; attr != nil {
		details.LoadBalancerName = attr.(string)
	}

	if attr := config["hostname_convention"]; attr != nil {
		details.HostNameConvention = attr.(string)
	}

	if attr := config["region"]; attr != nil {
		details.Region = attr.(string)
	}

	expandAwsInstanceFilter(config, details)

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenAwsSSHConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type == cac.InfrastructureTypes.AwsSSH && infraDef.DeploymentType == cac.DeploymentTypes.SSH {

		detailConfig := map[string]interface{}{}
		infraDetail := detail.ToAwsSSH()

		detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
		detailConfig["autoscaling_group_name"] = infraDetail.AutoscalingGroupName
		detailConfig["desired_capacity"] = infraDetail.DesiredCapacity
		detailConfig["host_connection_attrs_name"] = infraDetail.HostConnectionAttrsName
		detailConfig["host_connection_type"] = infraDetail.HostConnectionType
		detailConfig["loadbalancer_name"] = infraDetail.LoadBalancerName
		detailConfig["hostname_convention"] = infraDetail.HostNameConvention
		detailConfig["region"] = infraDetail.Region

		if infraDetail.AwsInstanceFilter != nil {
			if tags := flattenAwsTags(infraDetail.AwsInstanceFilter.Tags); len(tags) > 0 {
				detailConfig["tag"] = tags
			}

			if vpcIds := flattenVpcIds(infraDetail.AwsInstanceFilter.VpcIds); len(vpcIds) > 0 {
				detailConfig["vpc_ids"] = vpcIds
			}
		}

		results = append(results, detailConfig)
	}

	return results
}

func expandAwsInstanceFilter(config map[string]interface{}, details *cac.InfrastructureAwsSSH) {
	details.AwsInstanceFilter = &cac.AwsInstanceFilter{}

	if attr := config["tag"]; attr != nil {
		tagConfig := attr.(*schema.Set).List()
		if len(tagConfig) > 0 {
			details.AwsInstanceFilter.Tags = make([]*cac.AwsTag, len(tagConfig))
			for i, tag := range tagConfig {
				tConfig := tag.(map[string]interface{})
				t := &cac.AwsTag{
					Key:   tConfig["key"].(string),
					Value: tConfig["value"].(string),
				}
				details.AwsInstanceFilter.Tags[i] = t
			}
		}
	}

	if attr := config["vpc_ids"]; attr != nil {
		vpcIds := attr.(*schema.Set).List()
		if len(vpcIds) > 0 {
			details.AwsInstanceFilter.VpcIds = make([]string, len(vpcIds))
			for i, vpcId := range vpcIds {
				details.AwsInstanceFilter.VpcIds[i] = vpcId.(string)
			}
		}
	}
}

func flattenAwsTags(tags []*cac.AwsTag) []interface{} {
	results := []interface{}{}

	for _, tag := range tags {
		t := map[string]interface{}{
			"key":   tag.Key,
			"value": tag.Value,
		}
		results = append(results, t)
	}

	return results
}

func flattenVpcIds(vpcIds []string) []interface{} {
	var results []interface{}

	for _, vpcId := range vpcIds {
		results = append(results, vpcId)
	}

	return results
}
