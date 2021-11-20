package connector

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getAwsKmsSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "The AWS KMS configuration.",
		Type:          schema.TypeList,
		Optional:      true,
		MaxItems:      1,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "aws_kms"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"arn_ref": {
					Description: "A reference to the Harness secret containing the ARN of the AWS KMS.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"region": {
					Description: "The AWS region where the AWS Secret Manager is.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"delegate_selectors": {
					Description: "Connect using only the delegates which have these tags.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
				"credentials": {
					Description: "The credentials to use for connecting to aws.",
					Type:        schema.TypeList,
					MaxItems:    1,
					Required:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"manual": {
								Description:   "Specify the AWS key and secret used for authenticating.",
								Type:          schema.TypeList,
								MaxItems:      1,
								Optional:      true,
								ConflictsWith: []string{"aws_kms.0.credentials.0.assume_role", "aws_kms.0.credentials.0.inherit_from_delegate"},
								AtLeastOneOf:  []string{"aws_kms.0.credentials.0.manual", "aws_kms.0.credentials.0.assume_role", "aws_kms.0.credentials.0.inherit_from_delegate"},
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"access_key_ref": {
											Description: "The reference to the Harness secret containing the AWS access key.",
											Type:        schema.TypeString,
											Required:    true,
										},
										"secret_key_ref": {
											Description: "The reference to the Harness secret containing the AWS secret key.",
											Type:        schema.TypeString,
											Required:    true,
										},
									},
								},
							},
							"inherit_from_delegate": {
								Description:   "Inherit the credentials from from the delegate.",
								Type:          schema.TypeBool,
								Optional:      true,
								ConflictsWith: []string{"aws_kms.0.credentials.0.manual", "aws_kms.0.credentials.0.assume_role"},
								AtLeastOneOf:  []string{"aws_kms.0.credentials.0.manual", "aws_kms.0.credentials.0.assume_role", "aws_kms.0.credentials.0.inherit_from_delegate"},
								RequiredWith:  []string{"aws_kms.0.delegate_selectors"},
							},
							"assume_role": {
								Description:   "Connect using STS assume role.",
								Type:          schema.TypeList,
								Optional:      true,
								MaxItems:      1,
								ConflictsWith: []string{"aws_kms.0.credentials.0.manual", "aws_kms.0.credentials.0.inherit_from_delegate"},
								AtLeastOneOf:  []string{"aws_kms.0.credentials.0.manual", "aws_kms.0.credentials.0.assume_role", "aws_kms.0.credentials.0.inherit_from_delegate"},
								RequiredWith:  []string{"aws_kms.0.delegate_selectors"},
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"role_arn": {
											Description: "The ARN of the role to assume.",
											Type:        schema.TypeString,
											Required:    true,
										},
										"external_id": {
											Description: "If the administrator of the account to which the role belongs provided you with an external ID, then enter that value.",
											Type:        schema.TypeString,
											Optional:    true,
										},
										"duration": {
											Description: "The duration, in seconds, of the role session. The value can range from 900 seconds (15 minutes) to 3600 seconds (1 hour). By default, the value is set to 3600 seconds. An expiration can also be specified in the client request body. The minimum value is 1 hour.",
											Type:        schema.TypeInt,
											Required:    true,
											ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
												value := v.(int)
												if value < 900 || value > 3600 {
													errors = append(errors, fmt.Errorf("%q must be between 900 and 3600", k))
												}
												return
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func expandAwsKmsConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.AwsKms
	connector.AwsKms = &nextgen.AwsKmsConnectorDto{}

	if attr := config["arn_ref"].(string); attr != "" {
		connector.AwsKms.KmsArn = attr
	}

	if attr := config["region"].(string); attr != "" {
		connector.AwsKms.Region = attr
	}

	if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
		connector.AwsKms.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
	}

	if attr := config["credentials"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.AwsKms.Credential = &nextgen.AwsKmsConnectorCredentialDto{}

		if attr := config["inherit_from_delegate"].(bool); attr {
			connector.AwsKms.Credential.Type_ = nextgen.AwsKmsAuthTypes.AssumeIAMRole
			connector.AwsKms.Credential.AssumeIamRole = &nextgen.AwsKmsCredentialSpecAssumeIamdto{
				DelegateSelectors: connector.AwsKms.DelegateSelectors,
			}
		}

		if attr := config["manual"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
			connector.AwsKms.Credential.Type_ = nextgen.AwsKmsAuthTypes.ManualConfig
			connector.AwsKms.Credential.ManualConfig = &nextgen.AwsKmsCredentialSpecManualConfigDto{}

			if attr := config["access_key_ref"].(string); attr != "" {
				connector.AwsKms.Credential.ManualConfig.AccessKey = attr
			}

			if attr := config["secret_key_ref"].(string); attr != "" {
				connector.AwsKms.Credential.ManualConfig.SecretKey = attr
			}
		}

		if attr := config["assume_role"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
			connector.AwsKms.Credential.Type_ = nextgen.AwsKmsAuthTypes.AssumeSTSRole
			connector.AwsKms.Credential.AssumeStsRole = &nextgen.AwsKmsCredentialSpecAssumeStsdto{
				DelegateSelectors: connector.AwsKms.DelegateSelectors,
			}

			if attr := config["role_arn"].(string); attr != "" {
				connector.AwsKms.Credential.AssumeStsRole.RoleArn = attr
			}

			if attr := config["external_id"].(string); attr != "" {
				connector.AwsKms.Credential.AssumeStsRole.ExternalName = attr
			}

			if attr, ok := config["duration"]; ok {
				connector.AwsKms.Credential.AssumeStsRole.AssumeStsRoleDuration = int32(attr.(int))
			}
		}
	}
}

func flattenAwsKmsConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.AwsKms {
		return nil
	}

	results := map[string]interface{}{}

	results["arn_ref"] = connector.AwsKms.KmsArn
	results["region"] = connector.AwsKms.Region
	results["delegate_selectors"] = connector.AwsKms.DelegateSelectors

	switch connector.AwsKms.Credential.Type_ {
	case nextgen.AwsKmsAuthTypes.AssumeIAMRole:
		results["credentials"] = []interface{}{
			map[string]interface{}{
				"inherit_from_delegate": true,
			},
		}
	case nextgen.AwsKmsAuthTypes.ManualConfig:
		results["credentials"] = []interface{}{
			map[string]interface{}{
				"manual": []interface{}{
					map[string]interface{}{
						"access_key_ref": connector.AwsKms.Credential.ManualConfig.AccessKey,
						"secret_key_ref": connector.AwsKms.Credential.ManualConfig.SecretKey,
					},
				},
			},
		}
	case nextgen.AwsKmsAuthTypes.AssumeSTSRole:
		results["credentials"] = []interface{}{
			map[string]interface{}{
				"assume_role": []interface{}{
					map[string]interface{}{
						"role_arn":    connector.AwsKms.Credential.AssumeStsRole.RoleArn,
						"external_id": connector.AwsKms.Credential.AssumeStsRole.ExternalName,
						"duration":    connector.AwsKms.Credential.AssumeStsRole.AssumeStsRoleDuration,
					},
				},
			},
		}
	default:
		return fmt.Errorf("unsupported aws kms auth type: %s", connector.AwsKms.Credential.Type_)
	}

	d.Set("aws_kms", []interface{}{results})

	return nil
}
