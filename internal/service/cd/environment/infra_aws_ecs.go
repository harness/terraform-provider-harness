package environment

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func infraDetailsAwsEcs() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"assign_public_ip": {
				Description: "Flag to assign a public IP address.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cluster_name": {
				Description: "The name of the ECS cluster to use.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"execution_role": {
				Description: "The ARN of the role to use for execution.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"launch_type": {
				Description: fmt.Sprintf("The type of launch configuration to use. Valid options are %s", strings.Join(cac.AwsEcsLaunchTypesSlice, ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"region": {
				Description: "The region to deploy to.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"security_group_ids": {
				Description: "The security group ids to apply to the ecs service.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"subnet_ids": {
				Description: "The subnet ids to apply to the ecs service.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"vpc_id": {
				Description: "The VPC ids to use when selecting the instances.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func expandAwsEcsConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAwsEcs{}

	if attr := config["assign_public_ip"]; attr != nil {
		details.AssignPublicIp = attr.(bool)
	}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["cluster_name"]; attr != "" {
		details.ClusterName = attr.(string)
	}

	if attr := config["execution_role"]; attr != "" {
		details.ExecutionRole = attr.(string)
	}

	if attr := config["launch_type"]; attr != "" {
		details.LaunchType = cac.AwsEcsLaunchType(attr.(string))
	}

	if attr := config["region"]; attr != "" {
		details.Region = attr.(string)
	}

	if attr := config["security_group_ids"]; attr != nil {
		sgs := []string{}
		for _, sg := range attr.(*schema.Set).List() {
			sgs = append(sgs, sg.(string))
		}
		if len(sgs) > 0 {
			details.SecurityGroupIds = sgs
		}
	}

	if attr := config["subnet_ids"]; attr != nil {
		subnets := []string{}
		for _, subnet := range attr.(*schema.Set).List() {
			subnets = append(subnets, subnet.(string))
		}
		if len(subnets) > 0 {
			details.SubnetIds = subnets
		}
	}

	if attr := config["vpc_id"]; attr != "" {
		details.VpcId = attr.(string)
	}

	infraDef.InfrastructureDetail = details.ToInfrastructureDetail()
}

func flattenAwsEcsConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.AwsEcs {
		return results
	}

	detailConfig := map[string]interface{}{}
	infraDetail := detail.ToAwsEcs()

	detailConfig["assign_public_ip"] = infraDetail.AssignPublicIp
	detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
	detailConfig["cluster_name"] = infraDetail.ClusterName
	detailConfig["execution_role"] = infraDetail.ExecutionRole
	detailConfig["launch_type"] = infraDetail.LaunchType
	detailConfig["region"] = infraDetail.Region
	detailConfig["security_group_ids"] = infraDetail.SecurityGroupIds
	detailConfig["subnet_ids"] = infraDetail.SubnetIds
	detailConfig["vpc_id"] = infraDetail.VpcId

	return append(results, detailConfig)
}
