package connector

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func getGitSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "Git connector",
		Type:          schema.TypeList,
		MaxItems:      1,
		Optional:      true,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "git"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"url": {
					Description: "Url of the git repository or account.",
					Type:        schema.TypeString,
					Required:    true,
				},
				"connection_type": {
					Description:  fmt.Sprintf("Whether the connection we're making is to a git repository or a git account. Valid values are %s.", strings.Join(nextgen.GitConnectorTypeValues, ", ")),
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
								ConflictsWith: []string{"git.0.credentials.0.ssh"},
								ExactlyOneOf:  []string{"git.0.credentials.0.ssh", "git.0.credentials.0.http"},
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"username": {
											Description:   "Username to use for authentication.",
											Type:          schema.TypeString,
											Optional:      true,
											ConflictsWith: []string{"git.0.credentials.0.http.0.username_ref"},
											ExactlyOneOf:  []string{"git.0.credentials.0.http.0.username", "git.0.credentials.0.http.0.username_ref"},
										},
										"username_ref": {
											Description:   "Reference to a secret containing the username to use for authentication.",
											Type:          schema.TypeString,
											Optional:      true,
											ConflictsWith: []string{"git.0.credentials.0.http.0.username"},
											ExactlyOneOf:  []string{"git.0.credentials.0.http.0.username", "git.0.credentials.0.http.0.username_ref"},
										},
										"password_ref": {
											Description: "Reference to a secret containing the password to use for authentication.",
											Type:        schema.TypeString,
											Required:    true,
										},
									},
								},
							},
							"ssh": {
								Description:   "Authenticate using SSH for the connection.",
								Type:          schema.TypeList,
								MaxItems:      1,
								Optional:      true,
								ConflictsWith: []string{"git.0.credentials.0.http"},
								ExactlyOneOf:  []string{"git.0.credentials.0.ssh", "git.0.credentials.0.http"},
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

func expandGitConfig(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.Git
	connector.Git = &nextgen.GitConfig{}

	if attr, ok := config["url"]; ok {
		connector.Git.Url = attr.(string)
	}

	if attr, ok := config["delegate_selectors"]; ok {
		connector.Git.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := config["validation_repo"]; ok {
		connector.Git.ValidationRepo = attr.(string)
	}

	if attr, ok := config["connection_type"]; ok {
		connector.Git.ConnectionType = attr.(string)
	}

	if attr, ok := config["credentials"]; ok {
		credConfig := attr.([]interface{})[0].(map[string]interface{})

		if attr := credConfig["http"].([]interface{}); len(attr) > 0 {
			httpConfig := attr[0].(map[string]interface{})
			connector.Git.Type_ = nextgen.GitAuthTypes.Http
			connector.Git.Http = &nextgen.GitHttpAuthenticationDto{}

			if attr, ok := httpConfig["username"]; ok {
				connector.Git.Http.Username = attr.(string)
			}

			if attr, ok := httpConfig["username_ref"]; ok {
				connector.Git.Http.UsernameRef = attr.(string)
			}

			if attr, ok := httpConfig["password_ref"]; ok {
				connector.Git.Http.PasswordRef = attr.(string)
			}
		}

		if attr := credConfig["ssh"].([]interface{}); len(attr) > 0 {
			sshConfig := attr[0].(map[string]interface{})
			connector.Git.Type_ = nextgen.GitAuthTypes.Ssh
			connector.Git.Ssh = &nextgen.GitSshAuthentication{}

			if attr, ok := sshConfig["ssh_key_ref"]; ok {
				connector.Git.Ssh.SshKeyRef = attr.(string)
			}
		}
	}
}

func flattenGitConfig(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.Git {
		return nil
	}

	results := map[string]interface{}{}

	results["url"] = connector.Git.Url
	results["connection_type"] = connector.Git.ConnectionType
	results["delegate_selectors"] = connector.Git.DelegateSelectors
	results["validation_repo"] = connector.Git.ValidationRepo

	switch connector.Git.Type_ {
	case nextgen.GitAuthTypes.Http:
		results["credentials"] = []map[string]interface{}{
			{
				"http": []map[string]interface{}{
					{
						"username":     connector.Git.Http.Username,
						"username_ref": connector.Git.Http.UsernameRef,
						"password_ref": connector.Git.Http.PasswordRef,
					},
				},
			},
		}
	case nextgen.GitAuthTypes.Ssh:
		results["credentials"] = []map[string]interface{}{
			{
				"ssh": []map[string]interface{}{
					{
						"ssh_key_ref": connector.Git.Ssh.SshKeyRef,
					},
				},
			},
		}
	default:
		return fmt.Errorf("unsupported git auth type: %s", connector.Git.Type_)
	}

	d.Set("git", []interface{}{results})

	return nil
}
