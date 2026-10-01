package connector

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceConnectorAnthropicModel() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an Anthropic Model connector.",
		ReadContext:   resourceConnectorAnthropicModelRead,
		CreateContext: resourceConnectorAnthropicModelCreateOrUpdate,
		UpdateContext: resourceConnectorAnthropicModelCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Anthropic API endpoint.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"model": {
				Description: "The Anthropic model to use (e.g. claude-opus-4-6).",
				Type:        schema.TypeString,
				Required:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"execute_on_delegate": {
				Description: "Execute on delegate or not.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"ignore_test_connection": {
				Description: "Ignore test connection when creating or updating the connector.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"auth": {
				Description: "The authentication configuration for the Anthropic connector.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Description:  "Authentication type for the Anthropic connector.",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"Token", "BedrockApiKey", "Vertex", "CloudProvider"}, false),
						},
						"token": {
							Description:   "Authenticate using an Anthropic API token.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"auth.0.bedrock_api_key", "auth.0.vertex", "auth.0.cloud_provider"},
							AtLeastOneOf: []string{
								"auth.0.token",
								"auth.0.bedrock_api_key",
								"auth.0.vertex",
								"auth.0.cloud_provider",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"token_ref": {
										Description: "Reference to a secret containing the Anthropic API token." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"bedrock_api_key": {
							Description:   "Authenticate using an AWS Bedrock API key.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"auth.0.token", "auth.0.vertex", "auth.0.cloud_provider"},
							AtLeastOneOf: []string{
								"auth.0.token",
								"auth.0.bedrock_api_key",
								"auth.0.vertex",
								"auth.0.cloud_provider",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"api_key_ref": {
										Description: "Reference to a secret containing the Bedrock API key." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
									"region": {
										Description: "AWS region for the Bedrock endpoint.",
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"vertex": {
							Description:   "Authenticate using Google Vertex AI credentials.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"auth.0.token", "auth.0.bedrock_api_key", "auth.0.cloud_provider"},
							AtLeastOneOf: []string{
								"auth.0.token",
								"auth.0.bedrock_api_key",
								"auth.0.vertex",
								"auth.0.cloud_provider",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_account_key_ref": {
										Description: "Reference to a secret containing the GCP service account key." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
									"project_id": {
										Description: "GCP project ID.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"region": {
										Description: "GCP region for the Vertex AI endpoint.",
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"cloud_provider": {
							Description:   "Authenticate using an existing cloud provider connector.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"auth.0.token", "auth.0.bedrock_api_key", "auth.0.vertex"},
							AtLeastOneOf: []string{
								"auth.0.token",
								"auth.0.bedrock_api_key",
								"auth.0.vertex",
								"auth.0.cloud_provider",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Description:  "Cloud provider type.",
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"AWS", "GCP", "Azure"}, false),
									},
									"connector_ref": {
										Description: "Reference to an existing cloud provider connector.",
										Type:        schema.TypeString,
										Required:    true,
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

	return resource
}

func resourceConnectorAnthropicModelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.AnthropicModel)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorAnthropicModel(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorAnthropicModelCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorAnthropicModel(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorAnthropicModel(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorAnthropicModel(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:          nextgen.ConnectorTypes.AnthropicModel,
		AnthropicModel: &nextgen.AnthropicModelConnector{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.AnthropicModel.Url = attr.(string)
	}

	if attr, ok := d.GetOk("model"); ok {
		connector.AnthropicModel.Model = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.AnthropicModel.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("execute_on_delegate"); ok {
		connector.AnthropicModel.ExecuteOnDelegate = attr.(bool)
	}

	if attr, ok := d.GetOk("ignore_test_connection"); ok {
		connector.AnthropicModel.IgnoreTestConnection = attr.(bool)
	}

	if attr, ok := d.GetOk("auth"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.AnthropicModel.Authentication = &nextgen.AnthropicModelAuthentication{}

		authType := config["auth_type"].(string)

		switch authType {
		case "Token":
			connector.AnthropicModel.Authentication.Type_ = nextgen.AnthropicModelAuthTypes.Token
			if attrToken, ok := config["token"]; ok {
				tokenList := attrToken.([]interface{})
				if len(tokenList) > 0 {
					configToken := tokenList[0].(map[string]interface{})
					connector.AnthropicModel.Authentication.Token = &nextgen.AnthropicModelTokenSpec{}
					if attr, ok := configToken["token_ref"]; ok {
						connector.AnthropicModel.Authentication.Token.TokenRef = attr.(string)
					}
				}
			}

		case "BedrockApiKey":
			connector.AnthropicModel.Authentication.Type_ = nextgen.AnthropicModelAuthTypes.BedrockApiKey
			if attrBedrock, ok := config["bedrock_api_key"]; ok {
				bedrockList := attrBedrock.([]interface{})
				if len(bedrockList) > 0 {
					configBedrock := bedrockList[0].(map[string]interface{})
					connector.AnthropicModel.Authentication.BedrockApiKey = &nextgen.AnthropicModelBedrockApiKeySpec{}
					if attr, ok := configBedrock["api_key_ref"]; ok {
						connector.AnthropicModel.Authentication.BedrockApiKey.ApiKeyRef = attr.(string)
					}
					if attr, ok := configBedrock["region"]; ok {
						connector.AnthropicModel.Authentication.BedrockApiKey.Region = attr.(string)
					}
				}
			}

		case "Vertex":
			connector.AnthropicModel.Authentication.Type_ = nextgen.AnthropicModelAuthTypes.Vertex
			if attrVertex, ok := config["vertex"]; ok {
				vertexList := attrVertex.([]interface{})
				if len(vertexList) > 0 {
					configVertex := vertexList[0].(map[string]interface{})
					connector.AnthropicModel.Authentication.Vertex = &nextgen.AnthropicModelVertexSpec{}
					if attr, ok := configVertex["service_account_key_ref"]; ok {
						connector.AnthropicModel.Authentication.Vertex.ServiceAccountKeyRef = attr.(string)
					}
					if attr, ok := configVertex["project_id"]; ok {
						connector.AnthropicModel.Authentication.Vertex.ProjectId = attr.(string)
					}
					if attr, ok := configVertex["region"]; ok {
						connector.AnthropicModel.Authentication.Vertex.Region = attr.(string)
					}
				}
			}

		case "CloudProvider":
			connector.AnthropicModel.Authentication.Type_ = nextgen.AnthropicModelAuthTypes.CloudProvider
			if attrCP, ok := config["cloud_provider"]; ok {
				cpList := attrCP.([]interface{})
				if len(cpList) > 0 {
					configCP := cpList[0].(map[string]interface{})
					connector.AnthropicModel.Authentication.CloudProvider = &nextgen.AnthropicModelCloudProviderSpec{}
					if attr, ok := configCP["type"]; ok {
						connector.AnthropicModel.Authentication.CloudProvider.Type_ = nextgen.AnthropicModelCloudProviderType(attr.(string))
					}
					if attr, ok := configCP["connector_ref"]; ok {
						connector.AnthropicModel.Authentication.CloudProvider.ConnectorRef = attr.(string)
					}
				}
			}
		}
	}

	return connector
}

func readConnectorAnthropicModel(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.AnthropicModel.Url)
	d.Set("model", connector.AnthropicModel.Model)
	d.Set("delegate_selectors", connector.AnthropicModel.DelegateSelectors)
	d.Set("execute_on_delegate", connector.AnthropicModel.ExecuteOnDelegate)
	d.Set("ignore_test_connection", connector.AnthropicModel.IgnoreTestConnection)

	switch connector.AnthropicModel.Authentication.Type_ {
	case nextgen.AnthropicModelAuthTypes.Token:
		d.Set("auth", []map[string]interface{}{
			{
				"auth_type": "Token",
				"token": []map[string]interface{}{
					{
						"token_ref": connector.AnthropicModel.Authentication.Token.TokenRef,
					},
				},
			},
		})
	case nextgen.AnthropicModelAuthTypes.BedrockApiKey:
		d.Set("auth", []map[string]interface{}{
			{
				"auth_type": "BedrockApiKey",
				"bedrock_api_key": []map[string]interface{}{
					{
						"api_key_ref": connector.AnthropicModel.Authentication.BedrockApiKey.ApiKeyRef,
						"region":      connector.AnthropicModel.Authentication.BedrockApiKey.Region,
					},
				},
			},
		})
	case nextgen.AnthropicModelAuthTypes.Vertex:
		d.Set("auth", []map[string]interface{}{
			{
				"auth_type": "Vertex",
				"vertex": []map[string]interface{}{
					{
						"service_account_key_ref": connector.AnthropicModel.Authentication.Vertex.ServiceAccountKeyRef,
						"project_id":              connector.AnthropicModel.Authentication.Vertex.ProjectId,
						"region":                  connector.AnthropicModel.Authentication.Vertex.Region,
					},
				},
			},
		})
	case nextgen.AnthropicModelAuthTypes.CloudProvider:
		d.Set("auth", []map[string]interface{}{
			{
				"auth_type": "CloudProvider",
				"cloud_provider": []map[string]interface{}{
					{
						"type":          string(connector.AnthropicModel.Authentication.CloudProvider.Type_),
						"connector_ref": connector.AnthropicModel.Authentication.CloudProvider.ConnectorRef,
					},
				},
			},
		})
	default:
		return fmt.Errorf("unsupported anthropic model auth type: %s", connector.AnthropicModel.Authentication.Type_)
	}

	return nil
}
