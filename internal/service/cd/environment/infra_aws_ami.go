package environment

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func infraDetailsAwsAmi() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ami_deployment_type": {
				Description:  fmt.Sprintf("The ami deployment type to use. Valid options are %s", strings.Join(cac.AmiDeploymentTypesSlice, ", ")),
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(cac.AmiDeploymentTypesSlice, false),
			},
			"asg_identifies_workload": {
				Description: "Flag to indicate whether the autoscaling group identifies the workload.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"autoscaling_group_name": {
				Description: "The name of the autoscaling group.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"classic_loadbalancers": {
				Description: "The classic load balancers to use.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
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
			"spotinst_cloud_provider_name": {
				Description: "The name of the SpotInst cloud provider to connect with.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"spotinst_config_json": {
				Description: "The SpotInst configuration to use.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"stage_classic_loadbalancers": {
				Description: "The staging classic load balancers to use.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"stage_target_group_arns": {
				Description: "The staging classic load balancers to use.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"target_group_arns": {
				Description: "The ARN's of the target groups.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"use_traffic_shift": {
				Description: "Flag to enable traffic shifting.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func expandAwsAmiConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAwsAmi{}

	if attr := config["ami_deployment_type"]; attr != "" {
		details.AmiDeploymentType = cac.AmiDeploymentType(attr.(string))
	}

	if attr := config["asg_identifies_workload"]; attr != nil {
		details.ASGIdentifiesWorkload = attr.(bool)
	}

	if attr := config["autoscaling_group_name"]; attr != "" {
		details.AutoscalingGroupName = attr.(string)
	}

	if attr := config["classic_loadbalancers"]; attr != nil {
		lbs := []string{}
		for _, lb := range attr.(*schema.Set).List() {
			lbs = append(lbs, lb.(string))
		}
		if len(lbs) > 0 {
			details.ClassicLoadBalancers = lbs
		}
	}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["hostname_convention"]; attr != "" {
		details.HostNameConvention = attr.(string)
	}

	if attr := config["region"]; attr != "" {
		details.Region = attr.(string)
	}

	if attr := config["spotinst_cloud_provider_name"]; attr != "" {
		details.SpotinstCloudProviderName = attr.(string)
	}

	if attr := config["spotinst_config_json"]; attr != "" {
		details.SpotinstElastiGroupJson = attr.(string)
	}

	if attr := config["stage_classic_loadbalancers"]; attr != nil {
		lbs := []string{}
		for _, lb := range attr.(*schema.Set).List() {
			lbs = append(lbs, lb.(string))
		}
		if len(lbs) > 0 {
			details.StageClassicLoadBalancers = lbs
		}
	}

	if attr := config["stage_target_group_arns"]; attr != nil {
		groupsArns := []string{}
		for _, arn := range attr.(*schema.Set).List() {
			groupsArns = append(groupsArns, arn.(string))
		}
		if len(groupsArns) > 0 {
			details.StageTargetGroupArns = groupsArns
		}
	}

	if attr := config["target_group_arns"]; attr != nil {
		groupsArns := []string{}
		for _, arn := range attr.(*schema.Set).List() {
			groupsArns = append(groupsArns, arn.(string))
		}
		if len(groupsArns) > 0 {
			details.TargetGroupArns = groupsArns
		}
	}

	if attr := config["use_traffic_shift"]; attr != nil {
		details.UseTrafficShift = attr.(bool)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenAwsAmiConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.AwsAmi {
		return results
	}

	detailConfig := map[string]interface{}{}
	infraDetail := detail.ToAwsAmi()

	detailConfig["ami_deployment_type"] = infraDetail.AmiDeploymentType.String()
	detailConfig["asg_identifies_workload"] = infraDetail.ASGIdentifiesWorkload
	detailConfig["autoscaling_group_name"] = infraDetail.AutoscalingGroupName
	detailConfig["classic_loadbalancers"] = infraDetail.ClassicLoadBalancers
	detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
	detailConfig["hostname_convention"] = infraDetail.HostNameConvention
	detailConfig["region"] = infraDetail.Region
	detailConfig["spotinst_cloud_provider_name"] = infraDetail.SpotinstCloudProviderName
	detailConfig["spotinst_config_json"] = infraDetail.SpotinstElastiGroupJson
	detailConfig["stage_classic_loadbalancers"] = infraDetail.StageClassicLoadBalancers
	detailConfig["stage_target_group_arns"] = infraDetail.StageTargetGroupArns
	detailConfig["target_group_arns"] = infraDetail.TargetGroupArns
	detailConfig["use_traffic_shift"] = infraDetail.UseTrafficShift

	return append(results, detailConfig)
}
