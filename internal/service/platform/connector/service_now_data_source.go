package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceConnectorSerivceNow() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up a Service Now connector.",
		ReadContext: resourceConnectorServiceNowRead,

		Schema: map[string]*schema.Schema{
			"service_now_url": {
				Description: "URL of service now.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"username": {
				Description: "Username to use for authentication.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"username_ref": {
				Description: "Reference to a secret containing the username to use for authentication." + secret_ref_text,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"password_ref": {
				Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"auth": {
				Description: "The credentials to use for the service now authentication.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Description: "Authentication types for Jira connector",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"username_password": {
							Description: "Authenticate using username password.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Description: "Username to use for authentication.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"username_ref": {
										Description: "Reference to a secret containing the username to use for authentication." + secret_ref_text,
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
						"adfs": {
							Description: "Authenticate using adfs client credentials with certificate.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate_ref": {
										Description: "Reference to a secret containing the certificate to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"private_key_ref": {
										Description: "Reference to a secret containing the privateKeyRef to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"client_id_ref": {
										Description: "Reference to a secret containing the clientIdRef to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"resource_id_ref": {
										Description: "Reference to a secret containing the resourceIdRef to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"adfs_url": {
										Description: "asdf URL.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"refresh_token": {
							Description: "Authenticate using refresh token grant type. Currently, this feature is behind the feature flag CDS_SERVICENOW_REFRESH_TOKEN_AUTH. Contact Harness Support to enable the feature.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"token_url": {
										Description: "Token url to use for authentication.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"refresh_token_ref": {
										Description: "Reference to a secret containing the refresh token to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"client_id_ref": {
										Description: "Reference to a secret containing the client id to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"client_secret_ref": {
										Description: "Reference to a secret containing the client secret to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"scope": {
										Description: "Scope string to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},						
					},
				},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchemaIdentifierRequired(resource.Schema)

	return resource
}
