package connector

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getNexusSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Nexus connector.",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "nexus"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "URL of the Nexus server.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"delegate_selectors": {
					Description: "Connect using only the delegates which have these tags.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
				"version": {
					Description: fmt.Sprintf("Version of the Nexus server. Valid values are %s", strings.Join(nextgen.NexusVersionSlice, ", ")),
					Type:        schema.TypeString,
					Required:    true,
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
								ConflictsWith: []string{"nexus.0.credentials.0.username_ref"},
								ExactlyOneOf:  []string{"nexus.0.credentials.0.username", "nexus.0.credentials.0.username_ref"},
							},
							"username_ref": {
								Description:   "Reference to a secret containing the username to use for authentication.",
								Type:          schema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"nexus.0.credentials.0.username"},
								ExactlyOneOf:  []string{"nexus.0.credentials.0.username", "nexus.0.credentials.0.username_ref"},
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

func expandNexusConfig(d []interface{}, connector *nextgen.ConnectorInfoDto) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Nexus.String()
	connector.Nexus = &nextgen.NexusConnectorDto{}

	if attr := config["url"].(string); attr != "" {
		connector.Nexus.NexusServerUrl = attr
	}

	if attr := config["version"].(string); attr != "" {
		connector.Nexus.Version = attr
	}

	if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
		connector.Nexus.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
	}

	connector.Nexus.Auth = &nextgen.NexusAuthenticationDto{
		Type_: nextgen.NexusAuthTypes.Anonymous.String(),
	}

	if attr := config["credentials"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.Nexus.Auth.Type_ = nextgen.NexusAuthTypes.UsernamePassword.String()
		connector.Nexus.Auth.UsernamePassword = &nextgen.NexusUsernamePasswordAuthDto{}

		if attr := config["username"].(string); attr != "" {
			connector.Nexus.Auth.UsernamePassword.Username = attr
		}

		if attr := config["username_ref"].(string); attr != "" {
			connector.Nexus.Auth.UsernamePassword.UsernameRef = attr
		}

		if attr := config["password_ref"].(string); attr != "" {
			connector.Nexus.Auth.UsernamePassword.PasswordRef = attr
		}
	}
}

func flattenNexusConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfoDto) error {
	if connector.Type_ != nextgen.ConnectorTypes.Nexus.String() {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.Nexus.NexusServerUrl
	results["delegate_selectors"] = connector.Nexus.DelegateSelectors
	results["version"] = connector.Nexus.Version

	switch connector.Nexus.Auth.Type_ {
	case nextgen.NexusAuthTypes.UsernamePassword.String():
		results["credentials"] = []map[string]interface{}{
			{
				"username":     connector.Nexus.Auth.UsernamePassword.Username,
				"username_ref": connector.Nexus.Auth.UsernamePassword.UsernameRef,
				"password_ref": connector.Nexus.Auth.UsernamePassword.PasswordRef,
			},
		}
	case nextgen.NexusAuthTypes.Anonymous.String():
		// noop
	default:
		return fmt.Errorf("unsupported nexus auth type: %s", connector.Nexus.Auth.Type_)
	}

	d.Set("nexus", []interface{}{results})

	return nil
}
