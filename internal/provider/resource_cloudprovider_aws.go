package provider

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudProviderAws() *schema.Resource {

	commonSchema := commonCloudProviderSchema()
	providerSchema := map[string]*schema.Schema{
		"credentials": {
			Description: "Credential configuration for connecting to AWS",
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"assume_cross_account_role": {
						Description: "Configuration for assuming a cross account role.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"role_arn": {
									Description: "This is an IAM role in the target deployment AWS account.",
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
					"iam_role_service_account": {
						Description:   "Configure the use of IAM role for service accounts.",
						Type:          schema.TypeList,
						Optional:      true,
						MaxItems:      1,
						ConflictsWith: []string{"credentials.0.access_keys", "credentials.0.delegate"},
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"delegate_selector": {
									Description: "Select the Delegate to use via one of its Selectors.",
									Type:        schema.TypeString,
									Required:    true,
								},
							},
						},
					},
					"delegate": {
						Description:   "Configuration for acquiring credentials through the IAM profile associated with the delegate.",
						Type:          schema.TypeList,
						Optional:      true,
						ConflictsWith: []string{"credentials.0.access_keys", "credentials.0.iam_role_service_account"},
						MaxItems:      1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"selector": {
									Description: "Select the Delegate to use via one of its Selectors.",
									Type:        schema.TypeString,
									Required:    true,
								},
							},
						},
					},
					"access_keys": {
						Description:   "Configuration for acquiring credentials with an api key and secret",
						Type:          schema.TypeList,
						Optional:      true,
						MaxItems:      1,
						ConflictsWith: []string{"credentials.0.delegate", "credentials.0.iam_role_service_account"},
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"access_key_id": {
									Description: "The plain text AWS access key id.",
									Type:        schema.TypeString,
									Optional:    true,
									ExactlyOneOf: []string{
										"credentials.0.access_keys.0.access_key_id",
										"credentials.0.access_keys.0.encrypted_access_key_name",
									},
								},
								"encrypted_access_key_name": {
									Description: "The name of the encrypted text secret in Harness containing the AWS access key id",
									Type:        schema.TypeString,
									Optional:    true,
									ExactlyOneOf: []string{
										"credentials.0.access_keys.0.access_key_id",
										"credentials.0.access_keys.0.encrypted_access_key_name",
									},
								},
								"encrypted_secret_access_key_secret_name": {
									Description: "The name of the encrypted text secret in Harness containing the AWS secret access key.",
									Type:        schema.TypeString,
									Required:    true,
								},
							},
						},
					},
				},
			},
		},
	}

	helpers.MergeSchemas(commonSchema, providerSchema)

	return &schema.Resource{
		Description:   "Resource for creating a physical data center cloud provider",
		CreateContext: resourceCloudProviderAwsCreate,
		ReadContext:   resourceCloudProviderAwsRead,
		UpdateContext: resourceCloudProviderAwsUpdate,
		DeleteContext: resourceCloudProviderDelete,

		Schema: providerSchema,
	}
}

func resourceCloudProviderAwsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	name := d.Get("name").(string)

	cp := &cac.AwsCloudProvider{}
	err := c.ConfigAsCode().GetCloudProviderByName(name, cp)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cp.Id)
	d.Set("name", cp.Name)

	scope, err := flattenUsageRestrictions(c, cp.UsageRestrictions)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("usage_scope", scope)

	return nil
}

func resourceCloudProviderAwsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	input := cac.NewEntity(cac.ObjectTypes.AwsCloudProvider).(*cac.AwsCloudProvider)
	input.Name = d.Get("name").(string)

	restrictions, err := expandUsageRestrictions(c, d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	if restrictions != nil {
		input.UsageRestrictions = restrictions
	}

	expandAwsCloudProviderCredentials(d.Get("credentials").([]interface{}), input)

	cp, err := c.ConfigAsCode().UpsertAwsCloudProvider(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cp.Id)

	return nil
}

func resourceCloudProviderAwsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	cp := cac.NewEntity(cac.ObjectTypes.AwsCloudProvider).(*cac.AwsCloudProvider)
	cp.Name = d.Get("name").(string)

	usageRestrictions, err := expandUsageRestrictions(c, d.Get("usage_scope").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(err)
	}
	cp.UsageRestrictions = usageRestrictions

	_, err = c.ConfigAsCode().UpsertAwsCloudProvider(cp)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func expandAwsCloudProviderCredentials(d []interface{}, input *cac.AwsCloudProvider) {
	config := d[0].(map[string]interface{})

	if attr, ok := config["access_keys"]; ok && attr != nil {
		expandAwsCloudProviderAccessKeys(attr.([]interface{}), input)
	}

	if attr, ok := config["delegate"]; ok && attr != nil {
		expandAwsCloudProviderDelegateConfig(attr.([]interface{}), input)
	}

	if attr, ok := config["iam_role_service_account"]; ok && attr != nil {
		expandAwsCloudProviderServiceAccount(attr.([]interface{}), input)
	}

	if attr, ok := config["assume_cross_account_role"]; ok && attr != nil {
		expndAwsCloudProviderCrossAccountAttributes(attr.([]interface{}), input)
	}
}

func expndAwsCloudProviderCrossAccountAttributes(d []interface{}, input *cac.AwsCloudProvider) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})

	input.CrossAccountAttributes = &cac.AwsCrossAccountAttributes{}

	if attr, ok := config["cross_account_role_arn"]; ok && attr != nil {
		input.CrossAccountAttributes.CrossAccountRoleArn = attr.(string)
	}

	if attr, ok := config["external_id"]; ok && attr != nil {
		input.CrossAccountAttributes.ExternalId = attr.(string)
	}
}

func expandAwsCloudProviderServiceAccount(d []interface{}, input *cac.AwsCloudProvider) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})

	input.UseIRSA = true

	if attr, ok := config["delegate_selector"]; ok && attr != nil {
		input.DelegateSelector = attr.(string)
	}
}

func expandAwsCloudProviderDelegateConfig(d []interface{}, input *cac.AwsCloudProvider) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})

	if attr, ok := config["selector"]; ok && attr != "" {
		input.DelegateSelector = attr.(string)
		input.UseEc2IamCredentials = true
	}
}

func expandAwsCloudProviderAccessKeys(d []interface{}, input *cac.AwsCloudProvider) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})

	if attr, ok := config["access_key_id"]; ok && attr != "" {
		input.AccessKey = attr.(string)
	}

	if attr, ok := config["encrypted_access_key_name"]; ok && attr != "" {
		input.AccessKeySecretId = &cac.SecretRef{
			Name: attr.(string),
		}
	}

	if attr, ok := config["encrypted_secret_access_key_secret_name"]; ok && attr != "" {
		input.SecretKey = &cac.SecretRef{
			Name: attr.(string),
		}
	}
}
