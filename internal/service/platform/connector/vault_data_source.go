package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceConnectorVault() *schema.Resource {
	resource := &schema.Resource{
		Description: "DataSource for looking up a Vault connector in Harness.",
		ReadContext: resourceConnectorVaultRead,

		Schema: map[string]*schema.Schema{
			"auth_token": {
				Description: "This is the authentication token for Vault.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"base_path": {
				Description: "This is the location of the Vault directory where Secret will be stored.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"vault_url": {
				Description: "URL of the HashiCorp Vault.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"is_read_only": {
				Description: "Read only or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"renewal_interval_minutes": {
				Description: "This is the time interval for token renewal.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"secret_engine_manually_configured": {
				Description: "Manually entered Secret Engine.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"secret_engine_name": {
				Description: "Name of the Secret Engine.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"app_role_id": {
				Description: "ID of App Role.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"secret_id": {
				Description: "ID of the Secret.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"is_default": {
				Description: "Is default or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"secret_engine_version": {
				Description: "Version of Secret Engine.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"namespace": {
				Description: "This is the Vault namespace where Secret will be created.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"use_k8s_auth": {
				Description: "Boolean value to indicate if K8s Auth is used for authentication.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"sink_path": {
				Description: "This is the location at which auth token is to be read from.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"use_vault_agent": {
				Description: "Boolean value to indicate if Vault Agent is used for authentication.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"use_aws_iam": {
				Description: "Boolean value to indicate if Aws Iam is used for authentication.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"aws_region": {
				Description: "This is the Aws region where aws iam auth will happen.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"vault_aws_iam_role": {
				Description: "This is the Vault role defined to bind to aws iam account/role being accessed.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"xvault_aws_iam_server_id": {
				Description: "This is the Aws Iam Header Server ID that has been configured for this Aws Iam instance.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"vault_k8s_auth_role": {
				Description: "This is the role where K8s auth will happen.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"service_account_token_path": {
				Description: "This is the SA token path where the token is mounted in the K8s Pod.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"k8s_auth_endpoint": {
				Description: "This is the path where kubernetes auth is enabled in Vault.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"renew_app_role_token": {
				Description: "Boolean value to indicate if appRole token renewal is enabled or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"access_type": {
				Description: "Access type.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"default": {
				Description: "Is default or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"read_only": {
				Description: "Read only.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
