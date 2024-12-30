package artifactRepositories

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorAzureArtifacts() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an Azure Artifacts connector.",
		ReadContext:   resourceConnectorAzureArtifactsRead,
		CreateContext: resourceConnectorAzureArtifactsCreateOrUpdate,
		UpdateContext: resourceConnectorAzureArtifactsCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,
		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Azure Artifacts server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"credentials": {
				Description: "Credentials to use for authentication.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"token_ref": {
							Description: "Reference to a secret containing the token to use for authentication." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorAzureArtifactsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.AzureArtifacts)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorAzureArtifacts(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func readConnectorAzureArtifacts(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {

	d.Set("url", connector.AzureArtifacts.Url)
	d.Set("delegate_selectors", connector.AzureArtifacts.DelegateSelectors)

	switch connector.AzureArtifacts.Auth.Creds.Type_ {
	case nextgen.AzureArtifactsAuthTypes.PersonalAccessToken:
		d.Set("credentials", []map[string]interface{}{
			{
				"token_ref": connector.AzureArtifacts.Auth.Creds.UserToken.TokenRef,
			},
		})
	default:
		return fmt.Errorf("unsupported auth type: %s", connector.AzureArtifacts.Auth.Creds.Type_)
	}

	return nil
}

func resourceConnectorAzureArtifactsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorAzureArtifacts(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorAzureArtifacts(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorAzureArtifacts(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.AzureArtifacts,
		AzureArtifacts: &nextgen.AzureArtifactsConnector{
			Auth: &nextgen.AzureArtifactsAuthentication{
				Creds: &nextgen.AzureArtifactsHttpCredentials{},
			},
		},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.AzureArtifacts.Url = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.AzureArtifacts.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("credentials"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		connector.AzureArtifacts.Auth.Creds.Type_ = nextgen.AzureArtifactsAuthTypes.PersonalAccessToken
		connector.AzureArtifacts.Auth.Creds.UserToken = &nextgen.AzureArtifactsUsernameToken{}

		if attr, ok := config["token_ref"]; ok {
			connector.AzureArtifacts.Auth.Creds.UserToken.TokenRef = attr.(string)
		}
	}

	return connector
}
