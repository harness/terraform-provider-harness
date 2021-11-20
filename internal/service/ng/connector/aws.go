package connector

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getAwsSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Aws account configuration.",
		Type:          schema.TypeList,
		Optional:      true,
		MaxItems:      1,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "aws"),
		AtLeastOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"manual": {
					Description: "Use IAM role for service accounts.",
					Type:        schema.TypeList,
					MaxItems:    1,
					Optional:    true,
					ConflictsWith: []string{
						"aws.0.irsa",
						"aws.0.inherit_from_delegate",
					},
					ExactlyOneOf: []string{
						"aws.0.manual",
						"aws.0.irsa",
						"aws.0.inherit_from_delegate",
					},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"access_key": {
								Description:   "AWS access key.",
								Type:          schema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"aws.0.manual.0.access_key_ref"},
								AtLeastOneOf:  []string{"aws.0.manual.0.access_key", "aws.0.manual.0.access_key_ref"},
							},
							"access_key_ref": {
								Description:   "Reference to the Harness secret containing the aws access key.",
								Type:          schema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"aws.0.manual.0.access_key"},
								AtLeastOneOf:  []string{"aws.0.manual.0.access_key", "aws.0.manual.0.access_key_ref"},
							},
							"secret_key_ref": {
								Description: "Reference to the Harness secret containing the aws secret key.",
								Type:        schema.TypeString,
								Required:    true,
							},
							"delegate_selectors": {
								Description: "Connect only use delegates with these tags.",
								Type:        schema.TypeSet,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"irsa": {
					Description: "Use IAM role for service accounts.",
					Type:        schema.TypeList,
					MaxItems:    1,
					Optional:    true,
					ConflictsWith: []string{
						"aws.0.manual",
						"aws.0.inherit_from_delegate",
					},
					ExactlyOneOf: []string{
						"aws.0.manual",
						"aws.0.irsa",
						"aws.0.inherit_from_delegate",
					},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"delegate_selectors": {
								Description: "The delegates to inherit the credentials from.",
								Type:        schema.TypeSet,
								Required:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"inherit_from_delegate": {
					Description: "Inherit credentials from the delegate.",
					Type:        schema.TypeList,
					MaxItems:    1,
					Optional:    true,
					ConflictsWith: []string{
						"aws.0.irsa",
						"aws.0.manual",
					},
					ExactlyOneOf: []string{
						"aws.0.manual",
						"aws.0.irsa",
						"aws.0.inherit_from_delegate",
					},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"delegate_selectors": {
								Description: "The delegates to inherit the credentials from.",
								Type:        schema.TypeSet,
								Required:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"cross_account_access": {
					Description: "Select this option if you want to use one AWS account for the connection, but you want to deploy or build in a different AWS account. In this scenario, the AWS account used for AWS access in Credentials will assume the IAM role you specify in Cross-account role ARN setting. This option uses the AWS Security Token Service (STS) feature.",
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"role_arn": {
								Description: "The Amazon Resource Name (ARN) of the role that you want to assume. This is an IAM role in the target AWS account.",
								Type:        schema.TypeString,
								Required:    true,
							},
							"external_id": {
								Description: "If the administrator of the account to which the role belongs provided you with an external ID, then enter that value.",
								Type:        schema.TypeString,
								Optional:    true,
							},
						},
					},
				},
			},
		},
	}
}

func expandAwsConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Aws
	connector.Aws = &nextgen.AwsConnector{
		Credential: &nextgen.AwsCredential{},
	}

	if attr := config["manual"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.Aws.Credential.Type_ = nextgen.AwsAuthTypes.ManualConfig
		connector.Aws.Credential.ManualConfig = &nextgen.AwsManualConfigSpec{}

		if attr := config["access_key"].(string); attr != "" {
			connector.Aws.Credential.ManualConfig.AccessKey = attr
		}

		if attr := config["access_key_ref"].(string); attr != "" {
			connector.Aws.Credential.ManualConfig.AccessKeyRef = attr
		}

		if attr := config["secret_key_ref"].(string); attr != "" {
			connector.Aws.Credential.ManualConfig.SecretKeyRef = attr
		}

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.Aws.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}
	}

	if attr := config["irsa"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.Aws.Credential.Type_ = nextgen.AwsAuthTypes.Irsa

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.Aws.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}
	}

	if attr := config["inherit_from_delegate"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.Aws.Credential.Type_ = nextgen.AwsAuthTypes.InheritFromDelegate

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.Aws.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}
	}

	if attr := config["cross_account_access"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.Aws.Credential.CrossAccountAccess = &nextgen.CrossAccountAccess{}

		if attr := config["role_arn"].(string); attr != "" {
			connector.Aws.Credential.CrossAccountAccess.CrossAccountRoleArn = attr
		}

		if attr := config["external_id"].(string); attr != "" {
			connector.Aws.Credential.CrossAccountAccess.ExternalId = attr
		}
	}
}

func flattenAwsConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.Aws {
		return nil
	}

	results := map[string]interface{}{}

	switch connector.Aws.Credential.Type_ {
	case nextgen.AwsAuthTypes.ManualConfig:
		results["manual"] = []map[string]interface{}{
			{
				"access_key":         connector.Aws.Credential.ManualConfig.AccessKey,
				"access_key_ref":     connector.Aws.Credential.ManualConfig.AccessKeyRef,
				"secret_key_ref":     connector.Aws.Credential.ManualConfig.SecretKeyRef,
				"delegate_selectors": connector.Aws.DelegateSelectors,
			},
		}
	case nextgen.AwsAuthTypes.Irsa:
		results["irsa"] = []map[string]interface{}{
			{
				"delegate_selectors": connector.Aws.DelegateSelectors,
			},
		}
	case nextgen.AwsAuthTypes.InheritFromDelegate:
		results["inherit_from_delegate"] = []map[string]interface{}{
			{
				"delegate_selectors": connector.Aws.DelegateSelectors,
			},
		}
	default:
		return fmt.Errorf("unsupported aws credential type: %s", connector.Aws.Credential.Type_)
	}

	d.Set("aws", []interface{}{results})

	return nil
}
