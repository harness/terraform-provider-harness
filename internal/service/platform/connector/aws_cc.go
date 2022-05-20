package connector

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	u "github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorAwsCC() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an AWS Cloud Cost connector.",
		ReadContext:   resourceConnectorAwsCCRead,
		CreateContext: resourceConnectorAwsCCCreateOrUpdate,
		UpdateContext: resourceConnectorAwsCCCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

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
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorAwsCCRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.CEAws)
	if err != nil {
		return err
	}

	if err := readConnectorAwsCC(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorAwsCCCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorAwsCC(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorAwsCC(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorAwsCC(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.CEAws,
		AwsCC: &nextgen.CeAwsConnector{
			CrossAccountAccess: &nextgen.CrossAccountAccess{},
			CurAttributes:      &nextgen.AwsCurAttributes{},
		},
	}

	if attr, ok := d.GetOk("account_id"); ok {
		connector.AwsCC.AwsAccountId = attr.(string)
	}

	if attr, ok := d.GetOk("report_name"); ok {
		connector.AwsCC.CurAttributes.ReportName = attr.(string)
	}

	if attr, ok := d.GetOk("s3_bucket"); ok {
		connector.AwsCC.CurAttributes.S3BucketName = attr.(string)
	}

	// if attr, ok := d.GetOk("s3_prefix"); ok {
	// 	connector.AwsCC.CurAttributes.S3Prefix = attr.(string)
	// }

	// if attr, ok := d.GetOk("region"); ok {
	// 	connector.AwsCC.CurAttributes.Region = attr.(string)
	// }

	if attr, ok := d.GetOk("cross_account_access"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.AwsCC.CrossAccountAccess = &nextgen.CrossAccountAccess{}

		if attr, ok := config["role_arn"]; ok {
			connector.AwsCC.CrossAccountAccess.CrossAccountRoleArn = attr.(string)
		}

		if attr, ok := config["external_id"]; ok {
			connector.AwsCC.CrossAccountAccess.ExternalId = attr.(string)
		}
	}

	if attr, ok := d.GetOk("features_enabled"); ok {
		connector.AwsCC.FeaturesEnabled = u.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorAwsCC(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("account_id", connector.AwsCC.AwsAccountId)
	d.Set("report_name", connector.AwsCC.CurAttributes.ReportName)
	d.Set("s3_bucket", connector.AwsCC.CurAttributes.S3BucketName)
	// d.Set("s3_prefix", connector.AwsCC.CurAttributes.S3Prefix)
	// d.Set("region", connector.AwsCC.CurAttributes.Region)
	d.Set("features_enabled", connector.AwsCC.FeaturesEnabled)
	d.Set("cross_account_access", []map[string]interface{}{
		{
			"role_arn":    connector.AwsCC.CrossAccountAccess.CrossAccountRoleArn,
			"external_id": connector.AwsCC.CrossAccountAccess.ExternalId,
		},
	})

	return nil
}
