package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceConnectorAnthropicModel() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up an Anthropic Model connector.",
		ReadContext: resourceConnectorAnthropicModelRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Anthropic API endpoint.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"model": {
				Description: "The Anthropic model to use.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"execute_on_delegate": {
				Description: "Execute on delegate or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"ignore_test_connection": {
				Description: "Ignore test connection when creating or updating the connector.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"auth": {
				Description: "The authentication configuration for the Anthropic connector.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Description: "Authentication type for the Anthropic connector.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"token": {
							Description: "Authenticate using an Anthropic API token.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"token_ref": {
										Description: "Reference to a secret containing the Anthropic API token." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"bedrock_api_key": {
							Description: "Authenticate using an AWS Bedrock API key.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"api_key_ref": {
										Description: "Reference to a secret containing the Bedrock API key." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"region": {
										Description: "AWS region for the Bedrock endpoint.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"vertex": {
							Description: "Authenticate using Google Vertex AI credentials.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_account_key_ref": {
										Description: "Reference to a secret containing the GCP service account key." + secret_ref_text,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"project_id": {
										Description: "GCP project ID.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"region": {
										Description: "GCP region for the Vertex AI endpoint.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"cloud_provider": {
							Description: "Authenticate using an existing cloud provider connector.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description: "Cloud provider type.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"connector_ref": {
										Description: "Reference to an existing cloud provider connector.",
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
