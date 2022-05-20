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

func ResourceConnectorAppDynamics() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an App Dynamics connector.",
		ReadContext:   resourceConnectorAppDynamicsRead,
		CreateContext: resourceConnectorAppDynamicsCreateOrUpdate,
		UpdateContext: resourceConnectorAppDynamicsCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "Url of the App Dynamics controller.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"account_name": {
				Description: "The App Dynamics account name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"username_password": {
				Description:   "Authenticate to App Dynamics using username and password.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ConflictsWith: []string{"api_token"},
				AtLeastOneOf:  []string{"username_password", "api_token"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Description: "Username to use for authentication.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"password_ref": {
							Description: "Reference to a secret containing the password to use for authentication.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"api_token": {
				Description:   "Authenticate to App Dynamics using api token.",
				Type:          schema.TypeList,
				MaxItems:      1,
				Optional:      true,
				ConflictsWith: []string{"username_password"},
				AtLeastOneOf:  []string{"username_password", "api_token"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_secret_ref": {
							Description: "Reference to the Harness secret containing the App Dynamics client secret.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"client_id": {
							Description: "The client id used for connecting to App Dynamics.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorAppDynamicsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.AppDynamics)
	if err != nil {
		return err
	}

	if err := readConnectorAppDynamics(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorAppDynamicsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorAppDynamics(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorAppDynamics(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorAppDynamics(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:       nextgen.ConnectorTypes.AppDynamics,
		AppDynamics: &nextgen.AppDynamicsConnectorDto{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.AppDynamics.ControllerUrl = attr.(string)
	}

	if attr, ok := d.GetOk("account_name"); ok {
		connector.AppDynamics.Accountname = attr.(string)
	}

	if attr, ok := d.GetOk("username_password"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		connector.AppDynamics.AuthType = nextgen.AppDynamicsAuthTypes.UsernamePassword

		if attr, ok := config["username"]; ok {
			connector.AppDynamics.Username = attr.(string)
		}

		if attr, ok := config["password_ref"]; ok {
			connector.AppDynamics.PasswordRef = attr.(string)
		}

	}

	if attr, ok := d.GetOk("api_token"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		connector.AppDynamics.AuthType = nextgen.AppDynamicsAuthTypes.ApiClientToken

		if attr, ok := config["client_secret_ref"]; ok {
			connector.AppDynamics.ClientSecretRef = attr.(string)
		}

		if attr, ok := config["client_id"]; ok {
			connector.AppDynamics.ClientId = attr.(string)
		}

	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.AppDynamics.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector
}

func readConnectorAppDynamics(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.AppDynamics.ControllerUrl)
	d.Set("account_name", connector.AppDynamics.Accountname)

	switch connector.AppDynamics.AuthType {
	case nextgen.AppDynamicsAuthTypes.UsernamePassword:
		d.Set("username_password", []interface{}{
			map[string]interface{}{
				"username":     connector.AppDynamics.Username,
				"password_ref": connector.AppDynamics.PasswordRef,
			},
		})
	case nextgen.AppDynamicsAuthTypes.ApiClientToken:
		d.Set("api_token", []interface{}{
			map[string]interface{}{
				"client_secret_ref": connector.AppDynamics.ClientSecretRef,
				"client_id":         connector.AppDynamics.ClientId,
			},
		})
	default:
		return fmt.Errorf("Unknown App Dynamics auth type: %s", connector.AppDynamics.AuthType)
	}

	d.Set("delegate_selectors", connector.AppDynamics.DelegateSelectors)

	return nil
}
