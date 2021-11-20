package connector

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getAppDynamicsSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "App Dynamics connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "app_dynamics"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
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
					ConflictsWith: []string{"app_dynamics.0.api_token"},
					AtLeastOneOf:  []string{"app_dynamics.0.username_password", "app_dynamics.0.api_token"},
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
					ConflictsWith: []string{"app_dynamics.0.username_password"},
					AtLeastOneOf:  []string{"app_dynamics.0.username_password", "app_dynamics.0.api_token"},
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
		},
	}
}

func expandAppDynamicsConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.AppDynamics
	connector.AppDynamics = &nextgen.AppDynamicsConnectorDto{}

	if attr, ok := config["url"]; ok {
		connector.AppDynamics.ControllerUrl = attr.(string)
	}

	if attr, ok := config["account_name"]; ok {
		connector.AppDynamics.Accountname = attr.(string)
	}

	if attr := config["username_password"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})

		connector.AppDynamics.AuthType = nextgen.AppDynamicsAuthTypes.UsernamePassword

		if attr, ok := config["username"]; ok {
			connector.AppDynamics.Username = attr.(string)
		}

		if attr, ok := config["password_ref"]; ok {
			connector.AppDynamics.PasswordRef = attr.(string)
		}

	}

	if attr := config["api_token"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})

		connector.AppDynamics.AuthType = nextgen.AppDynamicsAuthTypes.ApiClientToken

		if attr, ok := config["client_secret_ref"]; ok {
			connector.AppDynamics.ClientSecretRef = attr.(string)
		}

		if attr, ok := config["client_id"]; ok {
			connector.AppDynamics.ClientId = attr.(string)
		}

	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.AppDynamics.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

}

func flattenAppDynamicsConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.AppDynamics {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.AppDynamics.ControllerUrl
	results["account_name"] = connector.AppDynamics.Accountname

	switch connector.AppDynamics.AuthType {
	case nextgen.AppDynamicsAuthTypes.UsernamePassword:
		results["username_password"] = []interface{}{
			map[string]interface{}{
				"username":     connector.AppDynamics.Username,
				"password_ref": connector.AppDynamics.PasswordRef,
			},
		}
	case nextgen.AppDynamicsAuthTypes.ApiClientToken:
		results["api_token"] = []interface{}{
			map[string]interface{}{
				"client_secret_ref": connector.AppDynamics.ClientSecretRef,
				"client_id":         connector.AppDynamics.ClientId,
			},
		}
	default:
		return fmt.Errorf("Unknown App Dynamics auth type: %s", connector.AppDynamics.AuthType)
	}

	results["delegate_selectors"] = connector.AppDynamics.DelegateSelectors

	d.Set("app_dynamics", []interface{}{results})

	return nil
}
