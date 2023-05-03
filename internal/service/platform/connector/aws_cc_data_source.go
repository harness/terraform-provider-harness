package connector

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorAwsCC() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up an AWS Cloud Cost connector.",
		ReadContext: resourceConnectorAwsCCRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Description: "The AWS account id.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"report_name": {
				Description: "The cost and usage report name. Provided in the delivery options when the template is opened in the AWS console.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			// "region": {
			// 	Description: "The AWS region.",
			// 	Type:        schema.TypeString,
			// 	Computed:    true,
			// },
			"s3_bucket": {
				Description: "The name of s3 bucket.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			// "s3_prefix": {

			// }
			"cross_account_access": {
				Description: "Harness uses the secure cross-account role to access your AWS account. The role includes a restricted policy to access the cost and usage reports and resources for the sole purpose of cost analysis and cost optimization.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_arn": {
							Description: "The ARN of the role to use for cross-account access.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"external_id": {
							Description: "The external id of the role to use for cross-account access. This is a random unique value to provide additional secure authentication.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"features_enabled": {
				Description: fmt.Sprintf("The features enabled for the connector. Valid values are %s.", strings.Join(nextgen.CCMFeaturesSlice, ", ")),
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
