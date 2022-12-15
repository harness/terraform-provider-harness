package connector

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceConnectorGithub() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Github connector.",
		ReadContext:   resourceConnectorGithubRead,
		CreateContext: resourceConnectorGithubCreateOrUpdate,
		UpdateContext: resourceConnectorGithubCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Githubhub repository or account.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"connection_type": {
				Description:  fmt.Sprintf("Whether the connection we're making is to a github repository or a github account. Valid values are %s.", strings.Join(nextgen.GitConnectorTypeValues, ", ")),
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
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"api_authentication": {
				Description: "Configuration for using the github api. API Access is required for using “Git Experience”, for creation of Git based triggers, Webhooks management and updating Git statuses.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"github_app": {
							Description:   "Configuration for using the github app for interacting with the github api.",
							Type:          schema.TypeList,
							Optional:      true,
							MaxItems:      1,
							AtLeastOneOf:  []string{"api_authentication.0.token_ref", "api_authentication.0.github_app"},
							ConflictsWith: []string{"api_authentication.0.token_ref"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"installation_id": {
										Description: "Enter the Installation ID located in the URL of the installed GitHub App.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"application_id": {
										Description: "Enter the GitHub App ID from the GitHub App General tab.",
										Type:        schema.TypeString,
										Required:    true,
									},
									"private_key_ref": {
										Description: "Reference to the secret containing the private key." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"token_ref": {
							Description:   "Personal access token for interacting with the github api." + secret_ref_text,
							Type:          schema.TypeString,
							Optional:      true,
							AtLeastOneOf:  []string{"api_authentication.0.token_ref", "api_authentication.0.github_app"},
							ConflictsWith: []string{"api_authentication.0.github_app"},
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
							ConflictsWith: []string{"credentials.0.ssh"},
							ExactlyOneOf:  []string{"credentials.0.ssh", "credentials.0.http"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Description:   "Username to use for authentication.",
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.http.0.username_ref"},
										ExactlyOneOf:  []string{"credentials.0.http.0.username", "credentials.0.http.0.username_ref"},
									},
									"username_ref": {
										Description:   "Reference to a secret containing the username to use for authentication." + secret_ref_text,
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.http.0.username"},
										ExactlyOneOf:  []string{"credentials.0.http.0.username", "credentials.0.http.0.username_ref"},
									},
									"token_ref": {
										Description: "Reference to a secret containing the personal access to use for authentication." + secret_ref_text,
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
							ConflictsWith: []string{"credentials.0.http"},
							ExactlyOneOf:  []string{"credentials.0.ssh", "credentials.0.http"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssh_key_ref": {
										Description: "Reference to the Harness secret containing the ssh key." + secret_ref_text,
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
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorGithubRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Github)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorGithub(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorGithubCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorGithub(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorGithub(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorGithub(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:  nextgen.ConnectorTypes.Github,
		Github: &nextgen.GithubConnector{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.Github.Url = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Github.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("validation_repo"); ok {
		connector.Github.ValidationRepo = attr.(string)
	}

	if attr, ok := d.GetOk("connection_type"); ok {
		connector.Github.Type_ = nextgen.GitConnectorType(attr.(string))
	}

	if attr, ok := d.GetOk("credentials"); ok {
		credConfig := attr.([]interface{})[0].(map[string]interface{})
		connector.Github.Authentication = &nextgen.GithubAuthentication{}

		if attr := credConfig["http"].([]interface{}); len(attr) > 0 {
			httpConfig := attr[0].(map[string]interface{})
			connector.Github.Authentication.Type_ = nextgen.GitAuthTypes.Http
			connector.Github.Authentication.Http = &nextgen.GithubHttpCredentials{
				Type_:         nextgen.GithubHttpCredentialTypes.UsernameToken,
				UsernameToken: &nextgen.GithubUsernameToken{},
			}

			if attr, ok := httpConfig["username"]; ok {
				connector.Github.Authentication.Http.UsernameToken.Username = attr.(string)
			}

			if attr, ok := httpConfig["username_ref"]; ok {
				connector.Github.Authentication.Http.UsernameToken.UsernameRef = attr.(string)
			}

			if attr, ok := httpConfig["token_ref"]; ok {
				connector.Github.Authentication.Http.UsernameToken.TokenRef = attr.(string)
			}
		}

		if attr := credConfig["ssh"].([]interface{}); len(attr) > 0 {
			sshConfig := attr[0].(map[string]interface{})
			connector.Github.Authentication.Type_ = nextgen.GitAuthTypes.Ssh
			connector.Github.Authentication.Ssh = &nextgen.GithubSshCredentials{}

			if attr, ok := sshConfig["ssh_key_ref"]; ok {
				connector.Github.Authentication.Ssh.SshKeyRef = attr.(string)
			}
		}
	}

	if attr, ok := d.GetOk("api_authentication"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.Github.ApiAccess = &nextgen.GithubApiAccess{}

		if attr, ok := config["token_ref"]; ok && attr != "" {
			connector.Github.ApiAccess.Type_ = nextgen.GithubApiAccessTypes.Token
			connector.Github.ApiAccess.Token = &nextgen.GithubTokenSpec{
				TokenRef: attr.(string),
			}
		}

		if attr, ok := config["github_app"].([]interface{}); ok && len(attr) > 0 {
			config := attr[0].(map[string]interface{})
			connector.Github.ApiAccess.Type_ = nextgen.GithubApiAccessTypes.GithubApp
			connector.Github.ApiAccess.GithubApp = &nextgen.GithubAppSpec{}

			if attr, ok := config["installation_id"]; ok {
				connector.Github.ApiAccess.GithubApp.InstallationId = attr.(string)
			}

			if attr, ok := config["application_id"]; ok {
				connector.Github.ApiAccess.GithubApp.ApplicationId = attr.(string)
			}

			if attr, ok := config["private_key_ref"]; ok {
				connector.Github.ApiAccess.GithubApp.PrivateKeyRef = attr.(string)
			}

		}
	}

	return connector
}

func readConnectorGithub(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {

	d.Set("url", connector.Github.Url)
	d.Set("connection_type", connector.Github.Type_.String())
	d.Set("delegate_selectors", connector.Github.DelegateSelectors)
	d.Set("validation_repo", connector.Github.ValidationRepo)

	if connector.Github.Authentication != nil {
		switch connector.Github.Authentication.Type_ {
		case nextgen.GitAuthTypes.Http:
			switch connector.Github.Authentication.Http.Type_ {
			case nextgen.GithubHttpCredentialTypes.UsernameToken:
				d.Set("credentials", []map[string]interface{}{
					{
						"http": []map[string]interface{}{
							{
								"username":     connector.Github.Authentication.Http.UsernameToken.Username,
								"username_ref": connector.Github.Authentication.Http.UsernameToken.UsernameRef,
								"token_ref":    connector.Github.Authentication.Http.UsernameToken.TokenRef,
							},
						},
					},
				})
			default:
				return fmt.Errorf("unsupported github http authentication type: %s", connector.Github.Authentication.Type_)
			}

		case nextgen.GitAuthTypes.Ssh:
			d.Set("credentials", []map[string]interface{}{
				{
					"ssh": []map[string]interface{}{
						{
							"ssh_key_ref": connector.Github.Authentication.Ssh.SshKeyRef,
						},
					},
				},
			})
		default:
			return fmt.Errorf("unsupported git auth type: %s", connector.Github.Type_)
		}
	}

	if connector.Github.ApiAccess != nil {
		switch connector.Github.ApiAccess.Type_ {
		case nextgen.GithubApiAccessTypes.GithubApp:
			d.Set("api_authentication", []map[string]interface{}{
				{
					"github_app": []map[string]interface{}{
						{
							"installation_id": connector.Github.ApiAccess.GithubApp.InstallationId,
							"application_id":  connector.Github.ApiAccess.GithubApp.ApplicationId,
							"private_key_ref": connector.Github.ApiAccess.GithubApp.PrivateKeyRef,
						},
					},
				},
			})
		case nextgen.GithubApiAccessTypes.Token:
			d.Set("api_authentication", []map[string]interface{}{
				{
					"token_ref": connector.Github.ApiAccess.Token.TokenRef,
				},
			})
		}
	}

	return nil
}
