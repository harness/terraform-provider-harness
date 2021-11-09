package environment

import (
	"github.com/harness-io/harness-go-sdk/harness/cd/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func infraDetailsAwsLambda() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_provider_name": {
				Description: "The name of the cloud provider to connect with.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"iam_role": {
				Description: "The IAM role to use.",
				Type:        schema.TypeString,
				Optional:    true,
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

func expandAwsLambdaConfiguration(d []interface{}, infraDef *cac.InfrastructureDefinition) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	details := &cac.InfrastructureAwsLambda{}

	if attr := config["cloud_provider_name"]; attr != "" {
		details.CloudProviderName = attr.(string)
	}

	if attr := config["iam_role"]; attr != "" {
		details.IamRole = attr.(string)
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

func flattenAwsLambdaConfiguration(d *schema.ResourceData, infraDef *cac.InfrastructureDefinition) []interface{} {
	results := []interface{}{}

	if len(infraDef.InfrastructureDetail) == 0 {
		return results
	}

	detail := infraDef.InfrastructureDetail[0]

	if detail.Type != cac.InfrastructureTypes.AwsLambda {
		return results
	}

	detailConfig := map[string]interface{}{}
	infraDetail := detail.ToAwsLambda()

	detailConfig["cloud_provider_name"] = infraDetail.CloudProviderName
	detailConfig["iam_role"] = infraDetail.IamRole
	detailConfig["region"] = infraDetail.Region
	detailConfig["security_group_ids"] = infraDetail.SecurityGroupIds
	detailConfig["subnet_ids"] = infraDetail.SubnetIds
	detailConfig["vpc_id"] = infraDetail.VpcId

	return append(results, detailConfig)
}
