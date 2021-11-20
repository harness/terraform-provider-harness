package connector

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	u "github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getAwsCCSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Aws cloud cost account configuration.",
		Type:          schema.TypeList,
		Optional:      true,
		MaxItems:      1,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "aws_cloudcost"),
		AtLeastOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"account_id": {
					Description: "The AWS account id.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"report_name": {
					Description: "The cost and usage report name. Provided in the delivery options when the template is opened in the AWS console.",
					Type:        schema.TypeString,
					Required:    true,
				},
				// "region": {
				// 	Description: "The AWS region.",
				// 	Type:        schema.TypeString,
				// 	Required:    true,
				// },
				"s3_bucket": {
					Description: "The name of s3 bucket.",
					Type:        schema.TypeString,
					Required:    true,
				},
				// "s3_prefix": {

				// }
				"cross_account_access": {
					Description: "Harness uses the secure cross-account role to access your AWS account. The role includes a restricted policy to access the cost and usage reports and resources for the sole purpose of cost analysis and cost optimization.",
					Type:        schema.TypeList,
					Required:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"role_arn": {
								Description: "The ARN of the role to use for cross-account access.",
								Type:        schema.TypeString,
								Required:    true,
							},
							"external_id": {
								Description: "The external id of the role to use for cross-account access. This is a random unique value to provide additional secure authentication.",
								Type:        schema.TypeString,
								Required:    true,
							},
						},
					},
				},
				"features_enabled": {
					Description: fmt.Sprintf("The features enabled for the connector. Valid values are %s.", strings.Join(nextgen.CCMFeaturesSlice, ", ")),
					Type:        schema.TypeSet,
					Required:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}

func expandAwsCCConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.CEAws
	connector.AwsCC = &nextgen.CeAwsConnectorDto{
		CrossAccountAccess: &nextgen.CrossAccountAccess{},
		CurAttributes:      &nextgen.AwsCurAttributesDto{},
	}

	if attr, ok := config["account_id"]; ok {
		connector.AwsCC.AwsAccountId = attr.(string)
	}

	if attr, ok := config["report_name"]; ok {
		connector.AwsCC.CurAttributes.ReportName = attr.(string)
	}

	if attr, ok := config["s3_bucket"]; ok {
		connector.AwsCC.CurAttributes.S3BucketName = attr.(string)
	}

	// if attr, ok := config["s3_prefix"]; ok {
	// 	connector.AwsCC.CurAttributes.S3Prefix = attr.(string)
	// }

	// if attr, ok := config["region"]; ok {
	// 	connector.AwsCC.CurAttributes.Region = attr.(string)
	// }

	if attr, ok := config["cross_account_access"]; ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.AwsCC.CrossAccountAccess = &nextgen.CrossAccountAccess{}

		if attr, ok := config["role_arn"]; ok {
			connector.AwsCC.CrossAccountAccess.CrossAccountRoleArn = attr.(string)
		}

		if attr, ok := config["external_id"]; ok {
			connector.AwsCC.CrossAccountAccess.ExternalId = attr.(string)
		}
	}

	if attr, ok := config["features_enabled"]; ok {
		connector.AwsCC.FeaturesEnabled = u.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

}

func flattenAwsCCConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.CEAws {
		return nil
	}

	results := map[string]interface{}{}

	results["account_id"] = connector.AwsCC.AwsAccountId
	results["report_name"] = connector.AwsCC.CurAttributes.ReportName
	results["s3_bucket"] = connector.AwsCC.CurAttributes.S3BucketName
	// results["s3_prefix"] = connector.AwsCC.CurAttributes.S3Prefix
	// results["region"] = connector.AwsCC.CurAttributes.Region
	results["features_enabled"] = connector.AwsCC.FeaturesEnabled
	results["cross_account_access"] = []map[string]interface{}{
		{
			"role_arn":    connector.AwsCC.CrossAccountAccess.CrossAccountRoleArn,
			"external_id": connector.AwsCC.CrossAccountAccess.ExternalId,
		},
	}

	d.Set("aws_cloudcost", []interface{}{results})

	return nil
}
