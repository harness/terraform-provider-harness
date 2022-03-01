package connector

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getArtifactorySchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Artifactory connector.",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "artifactory"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "URL of the Artifactory server.",
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
								ConflictsWith: []string{"artifactory.0.credentials.0.username_ref"},
								ExactlyOneOf:  []string{"artifactory.0.credentials.0.username", "artifactory.0.credentials.0.username_ref"},
							},
							"username_ref": {
								Description:   "Reference to a secret containing the username to use for authentication.",
								Type:          schema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"artifactory.0.credentials.0.username"},
								ExactlyOneOf:  []string{"artifactory.0.credentials.0.username", "artifactory.0.credentials.0.username_ref"},
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

func expandArtifactoryConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Artifactory
	connector.Artifactory = &nextgen.ArtifactoryConnector{}

	if attr := config["url"].(string); attr != "" {
		connector.Artifactory.ArtifactoryServerUrl = attr
	}

	if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
		connector.Artifactory.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
	}

	connector.Artifactory.Auth = &nextgen.ArtifactoryAuthentication{
		Type_: nextgen.ArtifactoryAuthTypes.Anonymous,
	}

	if attr := config["credentials"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.Artifactory.Auth.Type_ = nextgen.ArtifactoryAuthTypes.UsernamePassword
		connector.Artifactory.Auth.UsernamePassword = &nextgen.ArtifactoryUsernamePasswordAuth{}

		if attr := config["username"].(string); attr != "" {
			connector.Artifactory.Auth.UsernamePassword.Username = attr
		}

		if attr := config["username_ref"].(string); attr != "" {
			connector.Artifactory.Auth.UsernamePassword.UsernameRef = attr
		}

		if attr := config["password_ref"].(string); attr != "" {
			connector.Artifactory.Auth.UsernamePassword.PasswordRef = attr
		}
	}
}

func flattenArtifactoryConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.Artifactory {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.Artifactory.ArtifactoryServerUrl
	results["delegate_selectors"] = connector.Artifactory.DelegateSelectors

	switch connector.Artifactory.Auth.Type_ {
	case nextgen.ArtifactoryAuthTypes.UsernamePassword:
		results["credentials"] = []map[string]interface{}{
			{
				"username":     connector.Artifactory.Auth.UsernamePassword.Username,
				"username_ref": connector.Artifactory.Auth.UsernamePassword.UsernameRef,
				"password_ref": connector.Artifactory.Auth.UsernamePassword.PasswordRef,
			},
		}
	case nextgen.ArtifactoryAuthTypes.Anonymous:
		// noop
	default:
		return fmt.Errorf("unsupported artifactory auth type: %s", connector.Artifactory.Auth.Type_)
	}

	d.Set("artifactory", []interface{}{results})

	return nil
}
