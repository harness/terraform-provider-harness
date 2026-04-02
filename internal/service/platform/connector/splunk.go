package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceConnectorSplunk() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Splunk connector.",
		ReadContext:   resourceConnectorSplunkRead,
		CreateContext: resourceConnectorSplunkCreateOrUpdate,
		UpdateContext: resourceConnectorSplunkCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Splunk server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"account_id": {
				Description: "Splunk account id.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			// Deprecated fields - kept for backward compatibility
			"username": {
				Description: "The username used for connecting to Splunk. Deprecated: Use 'username_password' block instead.",
				Type:        schema.TypeString,
				Optional:    true,
				Deprecated:  "Use 'username_password' authentication block instead",
			},
			"password_ref": {
				Description: "The reference to the Harness secret containing the Splunk password. Deprecated: Use 'username_password' block instead." + secret_ref_text,
				Type:        schema.TypeString,
				Optional:    true,
				Deprecated:  "Use 'username_password' authentication block instead",
			},
			// New authentication blocks
			"username_password": {
				Description:  "Authenticate to Splunk using username and password.",
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
			"bearer_token": {
				Description:  "Authenticate to Splunk using bearer token.",
				Type:         schema.TypeList,
				MaxItems:     1,
				Optional:     true,
				InputDefault: "bearer_token",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bearer_token_ref": {
							Description: "Reference to the Harness secret containing the Splunk bearer token." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"hec_token": {
				Description:  "Authenticate to Splunk using HEC (HTTP Event Collector) token.",
				Type:         schema.TypeList,
				MaxItems:     1,
				Optional:     true,
				InputDefault: "hec_token",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hec_token_ref": {
							Description: "Reference to the Harness secret containing the Splunk HEC token." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"no_authentication": {
				Description:  "No authentication required for Splunk.",
				Type:         schema.TypeList,
				MaxItems:     1,
				Optional:     true,
				InputDefault: "no_authentication",
				Elem:         &schema.Resource{},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorSplunkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Splunk)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorSplunk(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorSplunkCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorSplunk(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorSplunk(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorSplunk(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:  nextgen.ConnectorTypes.Splunk,
		Splunk: &nextgen.SplunkConnector{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.Splunk.SplunkUrl = attr.(string)
	}

	if attr, ok := d.GetOk("account_id"); ok {
		connector.Splunk.AccountId = attr.(string)
	}

	// Handle authentication - check old fields first for backward compatibility
	// Priority: old fields > username_password > bearer_token > hec_token > no_authentication
	if username, usernameOk := d.GetOk("username"); usernameOk {
		// Old schema (deprecated) - username/password at root level
		connector.Splunk.AuthType = nextgen.SplunkAuthTypes.UsernamePassword
		connector.Splunk.Username = username.(string)
		if passwordRef, ok := d.GetOk("password_ref"); ok {
			connector.Splunk.PasswordRef = passwordRef.(string)
		}
	} else if attr, ok := d.GetOk("username_password"); ok {
		// New schema - username_password block
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.Splunk.AuthType = nextgen.SplunkAuthTypes.UsernamePassword

		if username, ok := config["username"]; ok {
			connector.Splunk.Username = username.(string)
		}
		if passwordRef, ok := config["password_ref"]; ok {
			connector.Splunk.PasswordRef = passwordRef.(string)
		}
	} else if attr, ok := d.GetOk("bearer_token"); ok {
		// Bearer token authentication
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.Splunk.AuthType = nextgen.SplunkAuthTypes.BearerToken

		if bearerTokenRef, ok := config["bearer_token_ref"]; ok {
			connector.Splunk.TokenRef = bearerTokenRef.(string)
		}
	} else if attr, ok := d.GetOk("hec_token"); ok {
		// HEC token authentication
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.Splunk.AuthType = nextgen.SplunkAuthTypes.HECToken

		if hecTokenRef, ok := config["hec_token_ref"]; ok {
			connector.Splunk.TokenRef = hecTokenRef.(string)
		}
	} else if _, ok := d.GetOk("no_authentication"); ok {
		// No authentication
		connector.Splunk.AuthType = nextgen.SplunkAuthTypes.None
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Splunk.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorSplunk(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.Splunk.SplunkUrl)
	d.Set("account_id", connector.Splunk.AccountId)
	d.Set("delegate_selectors", connector.Splunk.DelegateSelectors)

	// Handle authentication based on AuthType
	// Detect which schema format is being used in the configuration
	_, hasOldUsername := d.GetOk("username")
	_, hasOldPasswordRef := d.GetOk("password_ref")
	_, hasNewUsernamePassword := d.GetOk("username_password")

	usingOldSchema := (hasOldUsername || hasOldPasswordRef) && !hasNewUsernamePassword
	usingNewSchema := hasNewUsernamePassword && !hasOldUsername && !hasOldPasswordRef
	// If neither format is configured, this is likely a data source - populate both formats
	populateBothFormats := !usingOldSchema && !usingNewSchema

	switch connector.Splunk.AuthType {
	case nextgen.SplunkAuthTypes.UsernamePassword:
		if usingOldSchema || populateBothFormats {
			// Populate old flat schema fields
			d.Set("username", connector.Splunk.Username)
			d.Set("password_ref", connector.Splunk.PasswordRef)
		}
		if usingNewSchema || populateBothFormats {
			// Populate new block schema
			d.Set("username_password", []interface{}{
				map[string]interface{}{
					"username":     connector.Splunk.Username,
					"password_ref": connector.Splunk.PasswordRef,
				},
			})
		}
	case nextgen.SplunkAuthTypes.BearerToken:
		d.Set("bearer_token", []interface{}{
			map[string]interface{}{
				"bearer_token_ref": connector.Splunk.TokenRef,
			},
		})
	case nextgen.SplunkAuthTypes.HECToken:
		d.Set("hec_token", []interface{}{
			map[string]interface{}{
				"hec_token_ref": connector.Splunk.TokenRef,
			},
		})
	case nextgen.SplunkAuthTypes.None:
		d.Set("no_authentication", []interface{}{map[string]interface{}{}})
	default:
		// If AuthType is not set (old connectors), treat as username/password
		if connector.Splunk.Username != "" || connector.Splunk.PasswordRef != "" {
			if usingOldSchema || populateBothFormats {
				d.Set("username", connector.Splunk.Username)
				d.Set("password_ref", connector.Splunk.PasswordRef)
			}
			if usingNewSchema || populateBothFormats {
				d.Set("username_password", []interface{}{
					map[string]interface{}{
						"username":     connector.Splunk.Username,
						"password_ref": connector.Splunk.PasswordRef,
					},
				})
			}
		}
	}

	return nil
}
