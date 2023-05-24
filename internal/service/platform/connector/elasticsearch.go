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

func ResourceConnectorElasticSearch() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an ElasticSearch connector.",
		ReadContext:   resourceConnectorElasticSearchRead,
		CreateContext: resourceConnectorElasticSearchCreateOrUpdate,
		UpdateContext: resourceConnectorElasticSearchCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the elasticsearch",
				Type:        schema.TypeString,
				Required:    true,
			},
			"username_password": {
				Description:  "Authenticate to ElasticSearch using username and password.",
				Type:         schema.TypeList,
				MaxItems:     1,
				Optional:     true,
				InputDefault: "username_password",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Description: "Username to use for authentication.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"password_ref": {
							Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"api_token": {
				Description:  "Authenticate to ElasticSearch using api token.",
				Type:         schema.TypeList,
				MaxItems:     1,
				Optional:     true,
				InputDefault: "api_token",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_secret_ref": {
							Description: "Reference to the Harness secret containing the ElasticSearch client secret." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
						"client_id": {
							Description: "The API Key id used for connecting to ElasticSearch.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"no_authentication": {
				Description:  "No Authentication to ElasticSearch",
				Type:         schema.TypeList,
				MaxItems:     1,
				Optional:     true,
				InputDefault: "no_authentication",
				Elem:         &schema.Resource{},
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorElasticSearchRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.ElasticSearch)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorElasticSearch(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorElasticSearchCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorElasticSearch(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorElasticSearch(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorElasticSearch(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:         nextgen.ConnectorTypes.ElasticSearch,
		ElasticSearch: &nextgen.ElkConnectorDto{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.ElasticSearch.Url = attr.(string)
	}

	if attr, ok := d.GetOk("username_password"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		connector.ElasticSearch.AuthType = nextgen.ElkAuthTypes.UsernamePassword

		if attr, ok := config["username"]; ok {
			connector.ElasticSearch.Username = attr.(string)
		}

		if attr, ok := config["password_ref"]; ok {
			connector.ElasticSearch.PasswordRef = attr.(string)
		}

	} else if attr, ok := d.GetOk("api_token"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		connector.ElasticSearch.AuthType = nextgen.ElkAuthTypes.ApiClientToken

		if attr, ok := config["client_secret_ref"]; ok {
			connector.ElasticSearch.ApiKeyRef = attr.(string)
		}

		if attr, ok := config["client_id"]; ok {
			connector.ElasticSearch.ApiKeyId = attr.(string)
		}

	} else {

		connector.ElasticSearch.AuthType = nextgen.ElkAuthTypes.None

	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.ElasticSearch.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorElasticSearch(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.ElasticSearch.Url)

	switch connector.ElasticSearch.AuthType {

	case nextgen.ElkAuthTypes.UsernamePassword:
		d.Set("username_password", []interface{}{
			map[string]interface{}{
				"username":     connector.ElasticSearch.Username,
				"password_ref": connector.ElasticSearch.PasswordRef,
			},
		})
	case nextgen.ElkAuthTypes.ApiClientToken:
		d.Set("api_token", []interface{}{
			map[string]interface{}{
				"client_id":         connector.ElasticSearch.ApiKeyId,
				"client_secret_ref": connector.ElasticSearch.ApiKeyRef,
			},
		})
	case nextgen.ElkAuthTypes.None:
		// no-op
	default:
		return fmt.Errorf("Unknown ElasticSearch auth type: %s", connector.ElasticSearch.AuthType)
	}

	d.Set("delegate_selectors", connector.ElasticSearch.DelegateSelectors)

	return nil
}
