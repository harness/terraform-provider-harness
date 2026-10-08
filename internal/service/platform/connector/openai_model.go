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

func ResourceConnectorOpenAIModel() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an OpenAI Model connector.",
		ReadContext:   resourceConnectorOpenAIModelRead,
		CreateContext: resourceConnectorOpenAIModelCreateOrUpdate,
		UpdateContext: resourceConnectorOpenAIModelCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,
		CustomizeDiff: validateOpenAIModelAuthConfig,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the OpenAI API endpoint.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"model": {
				Description: "The OpenAI model to use (e.g. gpt-4o).",
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
				Description: "The authentication configuration for the OpenAI connector.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Description:  "Authentication type for the OpenAI connector.",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"Token", "Vertex"}, false),
						},
						"token": {
							Description:   "Authenticate using an OpenAI API token.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"auth.0.vertex"},
							AtLeastOneOf: []string{
								"auth.0.token",
								"auth.0.vertex",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"token_ref": {
										Description: "Reference to a secret containing the OpenAI API token." + secret_ref_text,
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
							ConflictsWith: []string{"auth.0.token"},
							AtLeastOneOf: []string{
								"auth.0.token",
								"auth.0.vertex",
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
					},
				},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorOpenAIModelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.OpenAIModel)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorOpenAIModel(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorOpenAIModelCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorOpenAIModel(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorOpenAIModel(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorOpenAIModel(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:       nextgen.ConnectorTypes.OpenAIModel,
		OpenAIModel: &nextgen.OpenAIModelConnector{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.OpenAIModel.Url = attr.(string)
	}

	if attr, ok := d.GetOk("model"); ok {
		connector.OpenAIModel.Model = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.OpenAIModel.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("execute_on_delegate"); ok {
		connector.OpenAIModel.ExecuteOnDelegate = attr.(bool)
	}

	if attr, ok := d.GetOk("ignore_test_connection"); ok {
		connector.OpenAIModel.IgnoreTestConnection = attr.(bool)
	}

	if attr, ok := d.GetOk("auth"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.OpenAIModel.Authentication = &nextgen.OpenAIModelAuthentication{}

		authType := config["auth_type"].(string)

		switch authType {
		case "Token":
			connector.OpenAIModel.Authentication.Type_ = nextgen.OpenAIModelAuthTypes.Token
			if attrToken, ok := config["token"]; ok {
				tokenList := attrToken.([]interface{})
				if len(tokenList) > 0 {
					configToken := tokenList[0].(map[string]interface{})
					connector.OpenAIModel.Authentication.Token = &nextgen.OpenAIModelTokenSpec{}
					if attr, ok := configToken["token_ref"]; ok {
						connector.OpenAIModel.Authentication.Token.TokenRef = attr.(string)
					}
				}
			}

		case "Vertex":
			connector.OpenAIModel.Authentication.Type_ = nextgen.OpenAIModelAuthTypes.Vertex
			if attrVertex, ok := config["vertex"]; ok {
				vertexList := attrVertex.([]interface{})
				if len(vertexList) > 0 {
					configVertex := vertexList[0].(map[string]interface{})
					connector.OpenAIModel.Authentication.Vertex = &nextgen.OpenAIModelVertexSpec{}
					if attr, ok := configVertex["service_account_key_ref"]; ok {
						connector.OpenAIModel.Authentication.Vertex.ServiceAccountKeyRef = attr.(string)
					}
					if attr, ok := configVertex["project_id"]; ok {
						connector.OpenAIModel.Authentication.Vertex.ProjectId = attr.(string)
					}
					if attr, ok := configVertex["region"]; ok {
						connector.OpenAIModel.Authentication.Vertex.Region = attr.(string)
					}
				}
			}
		}
	}

	return connector
}

func readConnectorOpenAIModel(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.OpenAIModel == nil {
		return nil
	}
	d.Set("url", connector.OpenAIModel.Url)
	d.Set("model", connector.OpenAIModel.Model)
	d.Set("delegate_selectors", connector.OpenAIModel.DelegateSelectors)
	d.Set("execute_on_delegate", connector.OpenAIModel.ExecuteOnDelegate)
	d.Set("ignore_test_connection", connector.OpenAIModel.IgnoreTestConnection)

	if connector.OpenAIModel.Authentication == nil {
		return nil
	}

	switch connector.OpenAIModel.Authentication.Type_ {
	case nextgen.OpenAIModelAuthTypes.Token:
		if connector.OpenAIModel.Authentication.Token != nil {
			d.Set("auth", []map[string]interface{}{
				{
					"auth_type": "Token",
					"token": []map[string]interface{}{
						{
							"token_ref": connector.OpenAIModel.Authentication.Token.TokenRef,
						},
					},
				},
			})
		}
	case nextgen.OpenAIModelAuthTypes.Vertex:
		if connector.OpenAIModel.Authentication.Vertex != nil {
			d.Set("auth", []map[string]interface{}{
				{
					"auth_type": "Vertex",
					"vertex": []map[string]interface{}{
						{
							"service_account_key_ref": connector.OpenAIModel.Authentication.Vertex.ServiceAccountKeyRef,
							"project_id":              connector.OpenAIModel.Authentication.Vertex.ProjectId,
							"region":                  connector.OpenAIModel.Authentication.Vertex.Region,
						},
					},
				},
			})
		}
	default:
		return fmt.Errorf("unsupported openai model auth type: %s", connector.OpenAIModel.Authentication.Type_)
	}

	return nil
}

func validateOpenAIModelAuthConfig(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	auth, ok := d.GetOk("auth")
	if !ok {
		return nil
	}

	authList := auth.([]interface{})
	if len(authList) == 0 {
		return nil
	}

	config := authList[0].(map[string]interface{})
	authType := config["auth_type"].(string)

	switch authType {
	case "Token":
		if _, ok := d.GetOk("auth.0.token"); !ok {
			return fmt.Errorf("auth_type is \"Token\" but no token block provided")
		}
	case "Vertex":
		if _, ok := d.GetOk("auth.0.vertex"); !ok {
			return fmt.Errorf("auth_type is \"Vertex\" but no vertex block provided")
		}
	}

	return nil
}
