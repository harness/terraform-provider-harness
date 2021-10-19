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
		"access_key_id": {
			Description: "The plain text AWS access key id.",
			Type:        schema.TypeString,
			Optional:    true,
			ConflictsWith: []string{
				"access_key_id_secret_name",
			},
		},
		"access_key_id_secret_name": {
			Description: "The name of the Harness secret containing the AWS access key id",
			Type:        schema.TypeString,
			Optional:    true,
			ConflictsWith: []string{
				"access_key_id",
				"usage_scope",
			},
		},
		"secret_access_key_secret_name": {
			Description:   "The name of the Harness secret containing the AWS secret access key.",
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"usage_scope"},
		},
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
		"delegate_selector": {
			Description: "Select the Delegate to use via one of its Selectors.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"use_irsa": {
			Description:  "Use the AWS IAM Role for Service Accounts.",
			Type:         schema.TypeBool,
			Optional:     true,
			RequiredWith: []string{"delegate_selector"},
		},
		"use_ec2_iam_credentials": {
			Description:  "Use the EC2 Instance Profile for Service Accounts.",
			Type:         schema.TypeBool,
			Optional:     true,
			RequiredWith: []string{"delegate_selector"},
		},
	}

	helpers.MergeSchemas(commonSchema, providerSchema)

	return &schema.Resource{
		Description:   configAsCodeDescription("Resource for creating an AWS cloud provider."),
		CreateContext: resourceCloudProviderAwsCreateOrUpdate,
		ReadContext:   resourceCloudProviderAwsRead,
		UpdateContext: resourceCloudProviderAwsCreateOrUpdate,
		DeleteContext: resourceCloudProviderDelete,

		Schema: providerSchema,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCloudProviderAwsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	cp := &cac.AwsCloudProvider{}
	if err := c.ConfigAsCode().GetCloudProviderById(d.Id(), cp); err != nil {
		return diag.FromErr(err)
	} else if cp.IsEmpty() {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readCloudProviderAws(c, d, cp)
}

func resourceCloudProviderAwsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	var input *cac.AwsCloudProvider
	var err error

	if d.IsNewResource() {
		input = cac.NewEntity(cac.ObjectTypes.AwsCloudProvider).(*cac.AwsCloudProvider)
	} else {
		input = &cac.AwsCloudProvider{}
		if err = c.ConfigAsCode().GetCloudProviderById(d.Id(), input); err != nil {
			return diag.FromErr(err)
		} else if input.IsEmpty() {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	}

	if attr := d.Get("name"); attr != "" {
		input.Name = attr.(string)
	}

	if attr := d.Get("access_key_id"); attr != "" {
		input.AccessKey = attr.(string)
	}

	if attr := d.Get("access_key_id_secret_name"); attr != "" {
		input.AccessKeySecretId = &cac.SecretRef{
			Name: attr.(string),
		}
	}

	if attr := d.Get("secret_access_key_secret_name"); attr != "" {
		input.SecretKey = &cac.SecretRef{
			Name: attr.(string),
		}
	}

	if attr := d.Get("assume_cross_account_role"); attr != nil {
		crossAccountConfig := attr.([]interface{})
		if len(crossAccountConfig) > 0 {
			config := crossAccountConfig[0].(map[string]interface{})
			input.AssumeCrossAccountRole = true
			input.CrossAccountAttributes = &cac.AwsCrossAccountAttributes{}

			if attr := config["role_arn"]; attr != "" {
				input.CrossAccountAttributes.CrossAccountRoleArn = attr.(string)
			}

			if attr := config["external_id"]; attr != "" {
				input.CrossAccountAttributes.ExternalId = attr.(string)
			}
		}
	}

	if attr := d.Get("delegate_selector"); attr != "" {
		input.DelegateSelector = attr.(string)
	}

	if attr := d.Get("use_irsa"); attr != nil {
		input.UseIRSA = attr.(bool)
	}

	if attr := d.Get("use_ec2_iam_credentials"); attr != nil {
		input.UseEc2IamCredentials = attr.(bool)
	}

	if input.UsageRestrictions == nil {
		input.UsageRestrictions = &cac.UsageRestrictions{}
	}

	if err := expandUsageRestrictions(c, d.Get("usage_scope").(*schema.Set).List(), input.UsageRestrictions); err != nil {
		return diag.FromErr(err)
	}

	cp, err := c.ConfigAsCode().UpsertAwsCloudProvider(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readCloudProviderAws(c, d, cp)
}

func readCloudProviderAws(c *api.Client, d *schema.ResourceData, cp *cac.AwsCloudProvider) diag.Diagnostics {
	d.SetId(cp.Id)
	d.Set("name", cp.Name)
	d.Set("access_key_id", cp.AccessKey)
	d.Set("delegate_selector", cp.DelegateSelector)
	d.Set("use_irsa", cp.UseIRSA)

	if cp.AccessKeySecretId != nil {
		d.Set("access_key_id_secret_name", cp.AccessKeySecretId.Name)
	}

	if cp.SecretKey != nil {
		d.Set("secret_access_key_secret_name", cp.SecretKey.Name)
	}

	if cp.AssumeCrossAccountRole && cp.CrossAccountAttributes != nil {
		attrs := map[string]interface{}{
			"role_arn":    cp.CrossAccountAttributes.CrossAccountRoleArn,
			"external_id": cp.CrossAccountAttributes.ExternalId,
		}
		d.Set("assume_cross_account_role", attrs)
	}

	scope, err := flattenUsageRestrictions(c, cp.UsageRestrictions)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("usage_scope", scope)

	return nil
}
