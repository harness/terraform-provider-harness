package connector

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func getGitlabSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Gitlab connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "gitlab"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "Url of the gitlab repository or account.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"connection_type": {
					Description:  fmt.Sprintf("Whether the connection we're making is to a gitlab repository or a gitlab account. Valid values are %s.", strings.Join(nextgen.GitConnectorTypeValues, ", ")),
					Type:         schema.TypeString,
					Required:     true,
					ValidateFunc: validation.StringInSlice(nextgen.GitConnectorTypeValues, false),
				},
				"validation_repo": {
					Description: "Repository to test the connection with. This is only used when `connection_type` is `Account`.",
					Type:        schema.TypeString,
					Optional:    true,
				},
				"delegate_selectors": {
					Description: "Connect using only the delegates which have these tags.",
					Type:        schema.TypeSet,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
				},
				"api_authentication": {
					Description: "Configuration for using the gitlab api. API Access is required for using “Git Experience”, for creation of Git based triggers, Webhooks management and updating Git statuses.",
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"token_ref": {
								Description: "Personal access token for interacting with the gitlab api.",
								Type:        schema.TypeString,
								Required:    true,
							},
						},
					},
				},
				"credentials": {
					Description: "Credentials to use for the connection.",
					Type:        schema.TypeList,
					MaxItems:    1,
					Required:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"http": {
								Description:   "Authenticate using Username and password over http(s) for the connection.",
								Type:          schema.TypeList,
								MaxItems:      1,
								Optional:      true,
								ConflictsWith: []string{"gitlab.0.credentials.0.ssh"},
								ExactlyOneOf:  []string{"gitlab.0.credentials.0.ssh", "gitlab.0.credentials.0.http"},
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"username": {
											Description:   "Username to use for authentication.",
											Type:          schema.TypeString,
											Optional:      true,
											ConflictsWith: []string{"gitlab.0.credentials.0.http.0.username_ref"},
											ExactlyOneOf:  []string{"gitlab.0.credentials.0.http.0.username", "gitlab.0.credentials.0.http.0.username_ref"},
										},
										"username_ref": {
											Description:   "Reference to a secret containing the username to use for authentication.",
											Type:          schema.TypeString,
											Optional:      true,
											ConflictsWith: []string{"gitlab.0.credentials.0.http.0.username"},
											ExactlyOneOf:  []string{"gitlab.0.credentials.0.http.0.username", "gitlab.0.credentials.0.http.0.username_ref"},
										},
										"token_ref": {
											Description:   "Reference to a secret containing the personal access to use for authentication.",
											Type:          schema.TypeString,
											Optional:      true,
											ConflictsWith: []string{"gitlab.0.credentials.0.http.0.password_ref"},
											AtLeastOneOf:  []string{"gitlab.0.credentials.0.http.0.token_ref", "gitlab.0.credentials.0.http.0.password_ref"},
										},
										"password_ref": {
											Description:   "Reference to a secret containing the password to use for authentication.",
											Type:          schema.TypeString,
											Optional:      true,
											ConflictsWith: []string{"gitlab.0.credentials.0.http.0.token_ref"},
											AtLeastOneOf:  []string{"gitlab.0.credentials.0.http.0.token_ref", "gitlab.0.credentials.0.http.0.password_ref"},
										},
									},
								},
							},
							"ssh": {
								Description:   "Authenticate using SSH for the connection.",
								Type:          schema.TypeList,
								MaxItems:      1,
								Optional:      true,
								ConflictsWith: []string{"gitlab.0.credentials.0.http"},
								ExactlyOneOf:  []string{"gitlab.0.credentials.0.ssh", "gitlab.0.credentials.0.http"},
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"ssh_key_ref": {
											Description: "Reference to the Harness secret containing the ssh key.",
											Type:        schema.TypeString,
											Required:    true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func expandGitlabConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Gitlab
	connector.Gitlab = &nextgen.GitlabConnector{}

	if attr, ok := config["url"]; ok {
		connector.Gitlab.Url = attr.(string)
	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.Gitlab.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := config["validation_repo"]; ok {
		connector.Gitlab.ValidationRepo = attr.(string)
	}

	if attr, ok := config["connection_type"]; ok {
		connector.Gitlab.Type_ = nextgen.GitConnectorType(attr.(string))
	}

	if attr, ok := config["credentials"]; ok {
		credConfig := attr.([]interface{})[0].(map[string]interface{})
		connector.Gitlab.Authentication = &nextgen.GitlabAuthentication{}

		if attr := credConfig["http"].([]interface{}); len(attr) > 0 {
			httpConfig := attr[0].(map[string]interface{})
			connector.Gitlab.Authentication.Type_ = nextgen.GitAuthTypes.Http
			connector.Gitlab.Authentication.Http = &nextgen.GitlabHttpCredentials{}

			if attr := httpConfig["token_ref"].(string); attr != "" {
				connector.Gitlab.Authentication.Http.Type_ = nextgen.GitlabHttpCredentialTypes.UsernameToken
				connector.Gitlab.Authentication.Http.UsernameToken = &nextgen.GitlabUsernameToken{
					TokenRef: attr,
				}

				if attr := httpConfig["username"].(string); attr != "" {
					connector.Gitlab.Authentication.Http.UsernameToken.Username = attr
				}

				if attr := httpConfig["username_ref"].(string); attr != "" {
					connector.Gitlab.Authentication.Http.UsernameToken.UsernameRef = attr
				}
			}

			if attr := httpConfig["password_ref"].(string); attr != "" {
				connector.Gitlab.Authentication.Http.Type_ = nextgen.GitlabHttpCredentialTypes.UsernamePassword
				connector.Gitlab.Authentication.Http.UsernamePassword = &nextgen.GitlabUsernamePassword{
					PasswordRef: attr,
				}

				if attr := httpConfig["username"].(string); attr != "" {
					connector.Gitlab.Authentication.Http.UsernamePassword.Username = attr
				}

				if attr := httpConfig["username_ref"].(string); attr != "" {
					connector.Gitlab.Authentication.Http.UsernamePassword.UsernameRef = attr
				}
			}
		}

		if attr := credConfig["ssh"].([]interface{}); len(attr) > 0 {
			sshConfig := attr[0].(map[string]interface{})
			connector.Gitlab.Authentication.Type_ = nextgen.GitAuthTypes.Ssh
			connector.Gitlab.Authentication.Ssh = &nextgen.GitlabSshCredentials{}

			if attr := sshConfig["ssh_key_ref"].(string); attr != "" {
				connector.Gitlab.Authentication.Ssh.SshKeyRef = attr
			}
		}
	}

	if attr := config["api_authentication"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.Gitlab.ApiAccess = &nextgen.GitlabApiAccess{
			Type_: nextgen.GitlabApiAuthTypes.Token,
			Token: &nextgen.GitlabTokenSpec{},
		}

		if attr, ok := config["token_ref"]; ok {
			connector.Gitlab.ApiAccess.Token.TokenRef = attr.(string)
		}
	}
}

func flattenGitlabConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.Gitlab {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.Gitlab.Url
	results["connection_type"] = connector.Gitlab.Type_.String()
	results["delegate_selectors"] = connector.Gitlab.DelegateSelectors
	results["validation_repo"] = connector.Gitlab.ValidationRepo

	if connector.Gitlab.Authentication != nil {
		switch connector.Gitlab.Authentication.Type_ {
		case nextgen.GitAuthTypes.Http:
			switch connector.Gitlab.Authentication.Http.Type_ {
			case nextgen.GitlabHttpCredentialTypes.UsernameToken:
				results["credentials"] = []map[string]interface{}{
					{
						"http": []map[string]interface{}{
							{
								"username":     connector.Gitlab.Authentication.Http.UsernameToken.Username,
								"username_ref": connector.Gitlab.Authentication.Http.UsernameToken.UsernameRef,
								"token_ref":    connector.Gitlab.Authentication.Http.UsernameToken.TokenRef,
							},
						},
					},
				}
			case nextgen.GitlabHttpCredentialTypes.UsernamePassword:
				results["credentials"] = []map[string]interface{}{
					{
						"http": []map[string]interface{}{
							{
								"username":     connector.Gitlab.Authentication.Http.UsernamePassword.Username,
								"username_ref": connector.Gitlab.Authentication.Http.UsernamePassword.UsernameRef,
								"password_ref": connector.Gitlab.Authentication.Http.UsernamePassword.PasswordRef,
							},
						},
					},
				}
			default:
				return fmt.Errorf("unsupported gitlab http authentication type: %s", connector.Gitlab.Authentication.Type_)
			}

		case nextgen.GitAuthTypes.Ssh:
			results["credentials"] = []map[string]interface{}{
				{
					"ssh": []map[string]interface{}{
						{
							"ssh_key_ref": connector.Gitlab.Authentication.Ssh.SshKeyRef,
						},
					},
				},
			}
		default:
			return fmt.Errorf("unsupported git auth type: %s", connector.Gitlab.Type_)
		}
	}

	if connector.Gitlab.ApiAccess != nil {
		switch connector.Gitlab.ApiAccess.Type_ {
		case nextgen.GitlabApiAuthTypes.Token:
			results["api_authentication"] = []map[string]interface{}{
				{
					"token_ref": connector.Gitlab.ApiAccess.Token.TokenRef,
				},
			}
		default:
			return fmt.Errorf("unsupported gitlab api access type: %s", connector.Gitlab.ApiAccess.Type_)
		}
	}

	d.Set("gitlab", []interface{}{results})

	return nil
}
