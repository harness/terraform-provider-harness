package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorAws() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up an AWS connector.",
		ReadContext: resourceConnectorAwsRead,

		Schema: map[string]*schema.Schema{
			"manual": {
				Description: "Use IAM role for service accounts.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_key": {
							Description: "AWS access key.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"access_key_ref": {
							Description: "Reference to the Harness secret containing the aws access key." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"secret_key_ref": {
							Description: "Reference to the Harness secret containing the aws secret key." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"delegate_selectors": {
							Description: "Connect only use delegates with these tags.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"irsa": {
				Description: "Use IAM role for service accounts.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delegate_selectors": {
							Description: "The delegates to inherit the credentials from.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"inherit_from_delegate": {
				Description: "Inherit credentials from the delegate.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delegate_selectors": {
							Description: "The delegates to inherit the credentials from.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"cross_account_access": {
				Description: "Select this option if you want to use one AWS account for the connection, but you want to deploy or build in a different AWS account. In this scenario, the AWS account used for AWS access in Credentials will assume the IAM role you specify in Cross-account role ARN setting. This option uses the AWS Security Token Service (STS) feature.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_arn": {
							Description: "The Amazon Resource Name (ARN) of the role that you want to assume. This is an IAM role in the target AWS account.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"external_id": {
							Description: "If the administrator of the account to which the role belongs provided you with an external ID, then enter that value.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
