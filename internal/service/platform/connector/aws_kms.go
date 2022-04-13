package connector

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/gitsync"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorAwsKms() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an AWS KMS connector.",
		ReadContext:   resourceConnectorAwsKmsRead,
		CreateContext: resourceConnectorAwsKmsCreateOrUpdate,
		UpdateContext: resourceConnectorAwsKmsCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

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
							ConflictsWith: []string{"credentials.0.assume_role", "credentials.0.inherit_from_delegate"},
							AtLeastOneOf:  []string{"credentials.0.manual", "credentials.0.assume_role", "credentials.0.inherit_from_delegate"},
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
							ConflictsWith: []string{"credentials.0.manual", "credentials.0.assume_role"},
							AtLeastOneOf:  []string{"credentials.0.manual", "credentials.0.assume_role", "credentials.0.inherit_from_delegate"},
							RequiredWith:  []string{"delegate_selectors"},
						},
						"assume_role": {
							Description:   "Connect using STS assume role.",
							Type:          schema.TypeList,
							Optional:      true,
							MaxItems:      1,
							ConflictsWith: []string{"credentials.0.manual", "credentials.0.inherit_from_delegate"},
							AtLeastOneOf:  []string{"credentials.0.manual", "credentials.0.assume_role", "credentials.0.inherit_from_delegate"},
							RequiredWith:  []string{"delegate_selectors"},
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
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)
	gitsync.SetGitSyncSchema(resource.Schema, false)

	return resource
}

func resourceConnectorAwsKmsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.AwsKms)
	if err != nil {
		return err
	}

	if err := readConnectorAwsKms(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorAwsKmsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorAwsKms(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorAwsKms(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorAwsKms(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:  nextgen.ConnectorTypes.AwsKms,
		AwsKms: &nextgen.AwsKmsConnector{},
	}

	if attr, ok := d.GetOk("arn_ref"); ok {
		connector.AwsKms.KmsArn = attr.(string)
	}

	if attr, ok := d.GetOk("region"); ok {
		connector.AwsKms.Region = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.AwsKms.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("credentials"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.AwsKms.Credential = &nextgen.AwsKmsConnectorCredential{}

		if attr, ok := config["inherit_from_delegate"]; ok && attr.(bool) {
			connector.AwsKms.Credential.Type_ = nextgen.AwsKmsAuthTypes.AssumeIAMRole
			connector.AwsKms.Credential.AssumeIamRole = &nextgen.AwsKmsCredentialSpecAssumeIam{
				DelegateSelectors: connector.AwsKms.DelegateSelectors,
			}
		}

		if attr := config["manual"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
			connector.AwsKms.Credential.Type_ = nextgen.AwsKmsAuthTypes.ManualConfig
			connector.AwsKms.Credential.ManualConfig = &nextgen.AwsKmsCredentialSpecManualConfig{}

			if attr, ok := config["access_key_ref"]; ok {
				connector.AwsKms.Credential.ManualConfig.AccessKey = attr.(string)
			}

			if attr, ok := config["secret_key_ref"]; ok {
				connector.AwsKms.Credential.ManualConfig.SecretKey = attr.(string)
			}
		}

		if attr := config["assume_role"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
			connector.AwsKms.Credential.Type_ = nextgen.AwsKmsAuthTypes.AssumeSTSRole
			connector.AwsKms.Credential.AssumeStsRole = &nextgen.AwsKmsCredentialSpecAssumeSts{
				DelegateSelectors: connector.AwsKms.DelegateSelectors,
			}

			if attr, ok := config["role_arn"]; ok {
				connector.AwsKms.Credential.AssumeStsRole.RoleArn = attr.(string)
			}

			if attr, ok := config["external_id"]; ok {
				connector.AwsKms.Credential.AssumeStsRole.ExternalName = attr.(string)
			}

			if attr, ok := config["duration"]; ok {
				connector.AwsKms.Credential.AssumeStsRole.AssumeStsRoleDuration = int32(attr.(int))
			}
		}
	}

	return connector
}

func readConnectorAwsKms(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("arn_ref", connector.AwsKms.KmsArn)
	d.Set("region", connector.AwsKms.Region)
	d.Set("delegate_selectors", connector.AwsKms.DelegateSelectors)

	switch connector.AwsKms.Credential.Type_ {
	case nextgen.AwsKmsAuthTypes.AssumeIAMRole:
		d.Set("credentials", []interface{}{
			map[string]interface{}{
				"inherit_from_delegate": true,
			},
		})
	case nextgen.AwsKmsAuthTypes.ManualConfig:
		d.Set("credentials", []interface{}{
			map[string]interface{}{
				"manual": []interface{}{
					map[string]interface{}{
						"access_key_ref": connector.AwsKms.Credential.ManualConfig.AccessKey,
						"secret_key_ref": connector.AwsKms.Credential.ManualConfig.SecretKey,
					},
				},
			},
		})
	case nextgen.AwsKmsAuthTypes.AssumeSTSRole:
		d.Set("credentials", []interface{}{
			map[string]interface{}{
				"assume_role": []interface{}{
					map[string]interface{}{
						"role_arn":    connector.AwsKms.Credential.AssumeStsRole.RoleArn,
						"external_id": connector.AwsKms.Credential.AssumeStsRole.ExternalName,
						"duration":    connector.AwsKms.Credential.AssumeStsRole.AssumeStsRoleDuration,
					},
				},
			},
		})
	default:
		return fmt.Errorf("unsupported aws kms auth type: %s", connector.AwsKms.Credential.Type_)
	}

	return nil
}
