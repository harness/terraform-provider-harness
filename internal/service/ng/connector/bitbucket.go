package connector

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func getBitBucketSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "BitBucket connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "bitbucket"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "Url of the BitBucket repository or account.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"connection_type": {
					Description:  fmt.Sprintf("Whether the connection we're making is to a BitBucket repository or a BitBucket account. Valid values are %s.", strings.Join(nextgen.GitConnectorTypeValues, ", ")),
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
					Description: "Configuration for using the BitBucket api. API Access is required for using “Git Experience”, for creation of Git based triggers, Webhooks management and updating Git statuses.",
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"username": {
								Description:   "The username used for connecting to the api.",
								Type:          schema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"bitbucket.0.api_authentication.0.username_ref"},
								AtLeastOneOf:  []string{"bitbucket.0.api_authentication.0.username", "bitbucket.0.api_authentication.0.username_ref"},
							},
							"username_ref": {
								Description:   "The name of the Harness secret containing the username.",
								Type:          schema.TypeString,
								Optional:      true,
								ConflictsWith: []string{"bitbucket.0.api_authentication.0.username"},
								AtLeastOneOf:  []string{"bitbucket.0.api_authentication.0.username", "bitbucket.0.api_authentication.0.username_ref"},
							},
							"token_ref": {
								Description: "Personal access token for interacting with the BitBucket api.",
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
								ConflictsWith: []string{"bitbucket.0.credentials.0.ssh"},
								ExactlyOneOf:  []string{"bitbucket.0.credentials.0.ssh", "bitbucket.0.credentials.0.http"},
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"username": {
											Description:   "Username to use for authentication.",
											Type:          schema.TypeString,
											Optional:      true,
											ConflictsWith: []string{"bitbucket.0.credentials.0.http.0.username_ref"},
											ExactlyOneOf:  []string{"bitbucket.0.credentials.0.http.0.username", "bitbucket.0.credentials.0.http.0.username_ref"},
										},
										"username_ref": {
											Description:   "Reference to a secret containing the username to use for authentication.",
											Type:          schema.TypeString,
											Optional:      true,
											ConflictsWith: []string{"bitbucket.0.credentials.0.http.0.username"},
											ExactlyOneOf:  []string{"bitbucket.0.credentials.0.http.0.username", "bitbucket.0.credentials.0.http.0.username_ref"},
										},
										"password_ref": {
											Description: "Reference to a secret containing the password to use for authentication.",
											Type:        schema.TypeString,
											Optional:    true,
										},
									},
								},
							},
							"ssh": {
								Description:   "Authenticate using SSH for the connection.",
								Type:          schema.TypeList,
								MaxItems:      1,
								Optional:      true,
								ConflictsWith: []string{"bitbucket.0.credentials.0.http"},
								ExactlyOneOf:  []string{"bitbucket.0.credentials.0.ssh", "bitbucket.0.credentials.0.http"},
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

func expandBitBucketConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Bitbucket
	connector.BitBucket = &nextgen.BitbucketConnector{}

	if attr, ok := config["url"]; ok {
		connector.BitBucket.Url = attr.(string)
	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.BitBucket.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := config["validation_repo"]; ok {
		connector.BitBucket.ValidationRepo = attr.(string)
	}

	if attr, ok := config["connection_type"]; ok {
		connector.BitBucket.Type_ = nextgen.GitConnectorType(attr.(string))
	}

	if attr, ok := config["credentials"]; ok {
		credConfig := attr.([]interface{})[0].(map[string]interface{})
		connector.BitBucket.Authentication = &nextgen.BitbucketAuthentication{}

		if attr := credConfig["http"].([]interface{}); len(attr) > 0 {
			httpConfig := attr[0].(map[string]interface{})
			connector.BitBucket.Authentication.Type_ = nextgen.GitAuthTypes.Http
			connector.BitBucket.Authentication.Http = &nextgen.BitbucketHttpCredentials{
				Type_:            nextgen.BitBucketHttpCredentialTypes.UsernamePassword,
				UsernamePassword: &nextgen.BitbucketUsernamePassword{},
			}

			if attr := httpConfig["username"].(string); attr != "" {
				connector.BitBucket.Authentication.Http.UsernamePassword.Username = attr
			}

			if attr := httpConfig["username_ref"].(string); attr != "" {
				connector.BitBucket.Authentication.Http.UsernamePassword.UsernameRef = attr
			}

			if attr := httpConfig["password_ref"].(string); attr != "" {
				connector.BitBucket.Authentication.Http.UsernamePassword.PasswordRef = attr
			}
		}

		if attr := credConfig["ssh"].([]interface{}); len(attr) > 0 {
			sshConfig := attr[0].(map[string]interface{})
			connector.BitBucket.Authentication.Type_ = nextgen.GitAuthTypes.Ssh
			connector.BitBucket.Authentication.Ssh = &nextgen.BitbucketSshCredentials{}

			if attr := sshConfig["ssh_key_ref"].(string); attr != "" {
				connector.BitBucket.Authentication.Ssh.SshKeyRef = attr
			}
		}
	}

	if attr := config["api_authentication"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.BitBucket.ApiAccess = &nextgen.BitbucketApiAccess{
			Type_:         nextgen.BitBucketApiAccessTypes.UsernameToken,
			UsernameToken: &nextgen.BitbucketUsernameTokenApiAccess{},
		}

		if attr := config["username"].(string); attr != "" {
			connector.BitBucket.ApiAccess.UsernameToken.Username = attr
		}

		if attr := config["username_ref"].(string); attr != "" {
			connector.BitBucket.ApiAccess.UsernameToken.UsernameRef = attr
		}

		if attr := config["token_ref"].(string); attr != "" {
			connector.BitBucket.ApiAccess.UsernameToken.TokenRef = attr
		}
	}

}

func flattenBitBucketConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.Bitbucket {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.BitBucket.Url
	results["connection_type"] = connector.BitBucket.Type_.String()
	results["delegate_selectors"] = connector.BitBucket.DelegateSelectors
	results["validation_repo"] = connector.BitBucket.ValidationRepo

	if connector.BitBucket.Authentication != nil {
		switch connector.BitBucket.Authentication.Type_ {
		case nextgen.GitAuthTypes.Http:
			switch connector.BitBucket.Authentication.Http.Type_ {
			case nextgen.BitBucketHttpCredentialTypes.UsernamePassword:
				results["credentials"] = []map[string]interface{}{
					{
						"http": []map[string]interface{}{
							{
								"username":     connector.BitBucket.Authentication.Http.UsernamePassword.Username,
								"username_ref": connector.BitBucket.Authentication.Http.UsernamePassword.UsernameRef,
								"password_ref": connector.BitBucket.Authentication.Http.UsernamePassword.PasswordRef,
							},
						},
					},
				}
			default:
				return fmt.Errorf("unsupported BitBucket http authentication type: %s", connector.BitBucket.Authentication.Type_)
			}

		case nextgen.GitAuthTypes.Ssh:
			results["credentials"] = []map[string]interface{}{
				{
					"ssh": []map[string]interface{}{
						{
							"ssh_key_ref": connector.BitBucket.Authentication.Ssh.SshKeyRef,
						},
					},
				},
			}
		default:
			return fmt.Errorf("unsupported Bitbucket auth type: %s", connector.BitBucket.Type_)
		}
	}

	if connector.BitBucket.ApiAccess != nil {
		switch connector.BitBucket.ApiAccess.Type_ {
		case nextgen.BitBucketApiAccessTypes.UsernameToken:
			results["api_authentication"] = []map[string]interface{}{
				{
					"username":     connector.BitBucket.ApiAccess.UsernameToken.Username,
					"username_ref": connector.BitBucket.ApiAccess.UsernameToken.UsernameRef,
					"token_ref":    connector.BitBucket.ApiAccess.UsernameToken.TokenRef,
				},
			}
		default:
			return fmt.Errorf("unsupported BitBucket api access type: %s", connector.BitBucket.ApiAccess.Type_)
		}
	}

	d.Set("bitbucket", []interface{}{results})

	return nil
}
