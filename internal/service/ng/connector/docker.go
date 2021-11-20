package connector

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getDockerRegistrySchema() *schema.Schema {
	return &schema.Schema{
		Description:   "The docker registry to use for the connector.",
		Type:          schema.TypeList,
		Optional:      true,
		MaxItems:      1,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "docker_registry"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Description: fmt.Sprintf("The type of the docker registry. Valid options are %s", strings.Join(nextgen.DockerRegistryTypesSlice, ", ")),
					Type:        schema.TypeString,
					Required:    true,
				},
				"url": {
					Description: "The url of the docker registry.",
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
					Description: "The credentials to use for the docker registry. If not specified then the connection is made to the registry anonymously.",
					Type:        schema.TypeList,
					MaxItems:    1,
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"username": {
								Description:   "The username to use for the docker registry.",
								Type:          schema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"docker_registry.0.credentials.0.username_ref"},
								AtLeastOneOf: []string{
									"docker_registry.0.0.credentials.0.username",
									"docker_registry.0.0.credentials.0.username_ref",
								},
							},
							"username_ref": {
								Description:   "The reference to the username to use for the docker registry.",
								Type:          schema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"docker_registry.0.credentials.0.username"},
								AtLeastOneOf: []string{
									"docker_registry.0.credentials.0.username",
									"docker_registry.0.credentials.0.username_ref",
								},
							},
							"password_ref": {
								Description: "The reference to the password to use for the docker registry.",
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

func expandDockerRegistryConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.DockerRegistry
	connector.DockerRegistry = &nextgen.DockerConnector{}

	if attr := config["url"].(string); attr != "" {
		connector.DockerRegistry.DockerRegistryUrl = attr
	}

	if attr := config["type"].(string); attr != "" {
		connector.DockerRegistry.ProviderType = attr
	}

	if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
		connector.DockerRegistry.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
	}

	connector.DockerRegistry.Auth = &nextgen.DockerAuthentication{
		Type_: nextgen.DockerAuthTypes.Anonymous,
	}

	if attr := config["credentials"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.DockerRegistry.Auth.Type_ = nextgen.DockerAuthTypes.UsernamePassword
		connector.DockerRegistry.Auth.UsernamePassword = &nextgen.DockerUserNamePassword{}

		if attr := config["username"].(string); attr != "" {
			connector.DockerRegistry.Auth.UsernamePassword.Username = attr
		}

		if attr := config["username_ref"].(string); attr != "" {
			connector.DockerRegistry.Auth.UsernamePassword.UsernameRef = attr
		}

		if attr := config["password_ref"].(string); attr != "" {
			connector.DockerRegistry.Auth.UsernamePassword.PasswordRef = attr
		}
	}
}

func flattenDockerRegistryConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.DockerRegistry {
		return nil
	}

	results := map[string]interface{}{}

	results["type"] = connector.DockerRegistry.ProviderType
	results["url"] = connector.DockerRegistry.DockerRegistryUrl
	results["delegate_selectors"] = connector.DockerRegistry.DelegateSelectors

	switch connector.DockerRegistry.Auth.Type_ {
	case nextgen.DockerAuthTypes.UsernamePassword:
		results["credentials"] = []map[string]interface{}{
			{
				"username":     connector.DockerRegistry.Auth.UsernamePassword.Username,
				"username_ref": connector.DockerRegistry.Auth.UsernamePassword.UsernameRef,
				"password_ref": connector.DockerRegistry.Auth.UsernamePassword.PasswordRef,
			},
		}
	case nextgen.DockerAuthTypes.Anonymous:
		// noop
	default:
		return fmt.Errorf("unsupported docker registry auth type: %s", connector.DockerRegistry.Auth.Type_)
	}

	d.Set("docker_registry", []interface{}{results})

	return nil
}
