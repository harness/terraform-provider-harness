package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorSplunk() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Splunk connector.",
		ReadContext: resourceConnectorSplunkRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Splunk server.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account_id": {
				Description: "Splunk account id.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			// Deprecated fields - kept for backward compatibility
			"username": {
				Description: "The username used for connecting to Splunk. Deprecated: Use 'username_password' block instead.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"password_ref": {
				Description: "The reference to the Harness secret containing the Splunk password. Deprecated: Use 'username_password' block instead." + secret_ref_text,
				Type:        schema.TypeString,
				Computed:    true,
			},
			// New authentication blocks
			"username_password": {
				Description: "Authenticate to Splunk using username and password.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Description: "Username to use for authentication.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password_ref": {
							Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"bearer_token": {
				Description: "Authenticate to Splunk using bearer token.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bearer_token_ref": {
							Description: "Reference to the Harness secret containing the Splunk bearer token." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"hec_token": {
				Description: "Authenticate to Splunk using HEC (HTTP Event Collector) token.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hec_token_ref": {
							Description: "Reference to the Harness secret containing the Splunk HEC token." + secret_ref_text,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"no_authentication": {
				Description: "No authentication required for Splunk.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Resource{},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
