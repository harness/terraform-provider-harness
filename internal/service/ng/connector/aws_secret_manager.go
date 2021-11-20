package connector

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getAwsSecretManagerSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "The AWS Secret Manager configuration.",
		Type:          schema.TypeList,
		Optional:      true,
		MaxItems:      1,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "aws_secret_manager"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"secret_name_prefix": {
					Description: "A prefix to be added to all secrets.",
					Type:        schema.TypeString,
					Optional:    true,
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
								ConflictsWith: []string{"aws_secret_manager.0.credentials.0.assume_role", "aws_secret_manager.0.credentials.0.inherit_from_delegate"},
								AtLeastOneOf:  []string{"aws_secret_manager.0.credentials.0.manual", "aws_secret_manager.0.credentials.0.assume_role", "aws_secret_manager.0.credentials.0.inherit_from_delegate"},
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
								ConflictsWith: []string{"aws_secret_manager.0.credentials.0.manual", "aws_secret_manager.0.credentials.0.assume_role"},
								AtLeastOneOf:  []string{"aws_secret_manager.0.credentials.0.manual", "aws_secret_manager.0.credentials.0.assume_role", "aws_secret_manager.0.credentials.0.inherit_from_delegate"},
								RequiredWith:  []string{"aws_secret_manager.0.delegate_selectors"},
							},
							"assume_role": {
								Description:   "Connect using STS assume role.",
								Type:          schema.TypeList,
								Optional:      true,
								MaxItems:      1,
								ConflictsWith: []string{"aws_secret_manager.0.credentials.0.manual", "aws_secret_manager.0.credentials.0.inherit_from_delegate"},
								AtLeastOneOf:  []string{"aws_secret_manager.0.credentials.0.manual", "aws_secret_manager.0.credentials.0.assume_role", "aws_secret_manager.0.credentials.0.inherit_from_delegate"},
								RequiredWith:  []string{"aws_secret_manager.0.delegate_selectors"},
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

func expandAwsSecretManagerConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.AwsSecretManager
	connector.AwsSecretManager = &nextgen.AwsSecretManagerDto{}

	if attr := config["secret_name_prefix"].(string); attr != "" {
		connector.AwsSecretManager.SecretNamePrefix = attr
	}

	if attr := config["region"].(string); attr != "" {
		connector.AwsSecretManager.Region = attr
	}

	if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
		connector.AwsSecretManager.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
	}

	if attr := config["credentials"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.AwsSecretManager.Credential = &nextgen.AwsSecretManagerCredentialDto{}

		if attr := config["inherit_from_delegate"].(bool); attr {
			connector.AwsSecretManager.Credential.Type_ = nextgen.AwsSecretManagerAuthTypes.AssumeIAMRole
			connector.AwsSecretManager.Credential.AssumeIamRole = &nextgen.AwsSmCredentialSpecAssumeIamdto{
				// Spec: nil,
			}
		}

		if attr := config["manual"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
			connector.AwsSecretManager.Credential.Type_ = nextgen.AwsSecretManagerAuthTypes.ManualConfig
			connector.AwsSecretManager.Credential.ManualConfig = &nextgen.AwsSmCredentialSpecManualConfigDto{}

			if attr := config["access_key_ref"].(string); attr != "" {
				connector.AwsSecretManager.Credential.ManualConfig.AccessKey = attr
			}

			if attr := config["secret_key_ref"].(string); attr != "" {
				connector.AwsSecretManager.Credential.ManualConfig.SecretKey = attr
			}
		}

		if attr := config["assume_role"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
			connector.AwsSecretManager.Credential.Type_ = nextgen.AwsSecretManagerAuthTypes.AssumeSTSRole
			connector.AwsSecretManager.Credential.AssumeStsRole = &nextgen.AwsSmCredentialSpecAssumeStsdto{}

			if attr := config["role_arn"].(string); attr != "" {
				connector.AwsSecretManager.Credential.AssumeStsRole.RoleArn = attr
			}

			if attr := config["external_id"].(string); attr != "" {
				connector.AwsSecretManager.Credential.AssumeStsRole.ExternalId = attr
			}

			if attr, ok := config["duration"]; ok {
				connector.AwsSecretManager.Credential.AssumeStsRole.AssumeStsRoleDuration = int32(attr.(int))
			}
		}
	}
}

func flattenAwsSecretManagerConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.AwsSecretManager {
		return nil
	}

	results := map[string]interface{}{}

	results["secret_name_prefix"] = connector.AwsSecretManager.SecretNamePrefix
	results["region"] = connector.AwsSecretManager.Region
	results["delegate_selectors"] = connector.AwsSecretManager.DelegateSelectors

	switch connector.AwsSecretManager.Credential.Type_ {
	case nextgen.AwsSecretManagerAuthTypes.AssumeIAMRole:
		results["credentials"] = []interface{}{
			map[string]interface{}{
				"inherit_from_delegate": true,
			},
		}
	case nextgen.AwsSecretManagerAuthTypes.ManualConfig:
		results["credentials"] = []interface{}{
			map[string]interface{}{
				"manual": []interface{}{
					map[string]interface{}{
						"access_key_ref": connector.AwsSecretManager.Credential.ManualConfig.AccessKey,
						"secret_key_ref": connector.AwsSecretManager.Credential.ManualConfig.SecretKey,
					},
				},
			},
		}
	case nextgen.AwsSecretManagerAuthTypes.AssumeSTSRole:
		results["credentials"] = []interface{}{
			map[string]interface{}{
				"assume_role": []interface{}{
					map[string]interface{}{
						"role_arn":    connector.AwsSecretManager.Credential.AssumeStsRole.RoleArn,
						"external_id": connector.AwsSecretManager.Credential.AssumeStsRole.ExternalId,
						"duration":    connector.AwsSecretManager.Credential.AssumeStsRole.AssumeStsRoleDuration,
					},
				},
			},
		}
	default:
		return fmt.Errorf("unsupported aws kms auth type: %s", connector.AwsSecretManager.Credential.Type_)
	}

	d.Set("aws_secret_manager", []interface{}{results})

	return nil
}
