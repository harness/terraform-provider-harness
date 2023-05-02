package connector

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorTerraformCloud() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Terraform Cloud connector.",
		ReadContext:   resourceConnectorTerraformCloudRead,
		CreateContext: resourceConnectorTerraformCloudCreateOrUpdate,
		UpdateContext: resourceConnectorTerraformCloudCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,
		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Terraform Cloud platform.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"credentials": {
				Description: "Credentials to connect to the Terraform Cloud platform.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_token": {
							Description: "API token credentials to use for authentication.",
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"api_token_ref": {
										Description: "Reference to a secret containing the API token to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
					},
				},
			},
			"delegate_selectors": {
				Description: "Connect only using delegates with these tags.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorTerraformCloudRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.TerraformCloud)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorTerraformCloud(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorTerraformCloudCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorTerraformCloud(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorTerraformCloud(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorTerraformCloud(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.TerraformCloud,
		TerraformCloud: &nextgen.TerraformCloudConnector{
			Credential: &nextgen.TerraformCloudCredential{
				Type_: nextgen.TerraformCloudAuthTypes.ApiToken,
			},
		},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.TerraformCloud.TerraformCloudUrl = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.TerraformCloud.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("credentials"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.TerraformCloud.Credential = &nextgen.TerraformCloudCredential{}

		if attr := config["api_token"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
			connector.TerraformCloud.Credential.Type_ = nextgen.TerraformCloudAuthTypes.ApiToken
			connector.TerraformCloud.Credential.ApiToken = &nextgen.TerraformCloudTokenCredentials{}

			if attr, ok := config["api_token_ref"]; ok {
				connector.TerraformCloud.Credential.ApiToken.ApiToken = attr.(string)
			}
		}
	}

	return connector
}

func readConnectorTerraformCloud(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.TerraformCloud.TerraformCloudUrl)
	d.Set("delegate_selectors", connector.TerraformCloud.DelegateSelectors)

	switch connector.TerraformCloud.Credential.Type_ {
	case nextgen.TerraformCloudAuthTypes.ApiToken:
		d.Set("credentials", []interface{}{
			map[string]interface{}{
				"api_token": []interface{}{
					map[string]interface{}{
						"api_token_ref": connector.TerraformCloud.Credential.ApiToken.ApiToken,
					},
				},
			},
		})
	default:
		return fmt.Errorf("unsupported Terraform Cloud credentials type: %s", connector.TerraformCloud.Credential.Type_)
	}

	return nil
}
