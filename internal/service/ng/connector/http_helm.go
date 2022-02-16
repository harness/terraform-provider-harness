package connector

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getHttpHelmSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Helm connector.",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "http_helm"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "URL of the helm server.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"delegate_selectors": {
					Description: "Connect using only the delegates which have these tags.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
				"credentials": {
					Description: "Credentials to use for authentication.",
					Type:        schema.TypeList,
					MaxItems:    1,
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"username": {
								Description:   "Username to use for authentication.",
								Type:          schema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"http_helm.0.credentials.0.username_ref"},
								ExactlyOneOf:  []string{"http_helm.0.credentials.0.username", "http_helm.0.credentials.0.username_ref"},
							},
							"username_ref": {
								Description:   "Reference to a secret containing the username to use for authentication.",
								Type:          schema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"http_helm.0.credentials.0.username"},
								ExactlyOneOf:  []string{"http_helm.0.credentials.0.username", "http_helm.0.credentials.0.username_ref"},
							},
							"password_ref": {
								Description: "Reference to a secret containing the password to use for authentication.",
								Type:        schema.TypeString,
								Required:    true,
							},
						},
					},
				},
			},
		},
	}
}

func expandHttpHelmConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.HttpHelmRepo
	connector.HttpHelm = &nextgen.HttpHelmConnector{}

	if attr := config["url"].(string); attr != "" {
		connector.HttpHelm.HelmRepoUrl = attr
	}

	if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
		connector.HttpHelm.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
	}

	connector.HttpHelm.Auth = &nextgen.HttpHelmAuthentication{
		Type_: nextgen.HttpHelmAuthTypes.Anonymous,
	}

	if attr := config["credentials"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.HttpHelm.Auth.Type_ = nextgen.HttpHelmAuthTypes.UsernamePassword
		connector.HttpHelm.Auth.UsernamePassword = &nextgen.HttpHelmUsernamePassword{}

		if attr := config["username"].(string); attr != "" {
			connector.HttpHelm.Auth.UsernamePassword.Username = attr
		}

		if attr := config["username_ref"].(string); attr != "" {
			connector.HttpHelm.Auth.UsernamePassword.UsernameRef = attr
		}

		if attr := config["password_ref"].(string); attr != "" {
			connector.HttpHelm.Auth.UsernamePassword.PasswordRef = attr
		}
	}
}

func flattenHttpHelmConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.HttpHelmRepo {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.HttpHelm.HelmRepoUrl
	results["delegate_selectors"] = connector.HttpHelm.DelegateSelectors

	switch connector.HttpHelm.Auth.Type_ {
	case nextgen.HttpHelmAuthTypes.UsernamePassword:
		results["credentials"] = []map[string]interface{}{
			{
				"username":     connector.HttpHelm.Auth.UsernamePassword.Username,
				"username_ref": connector.HttpHelm.Auth.UsernamePassword.UsernameRef,
				"password_ref": connector.HttpHelm.Auth.UsernamePassword.PasswordRef,
			},
		}
	case nextgen.HttpHelmAuthTypes.Anonymous:
		// noop
	default:
		return fmt.Errorf("unsupported http helm auth type: %s", connector.HttpHelm.Auth.Type_)
	}

	d.Set("http_helm", []interface{}{results})

	return nil
}
