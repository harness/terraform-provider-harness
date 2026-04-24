package codeRepositories

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

func ResourceConnectorBitbucket() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Bitbucket connector.",
		ReadContext:   resourceConnectorBitbucketRead,
		CreateContext: resourceConnectorBitbucketCreateOrUpdate,
		UpdateContext: resourceConnectorBitbucketCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,
		CustomizeDiff: validateBitbucketApiAuthentication,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the BitBucket repository or account.",
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
				Description: "Tags to filter delegates for connection.",
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
						"auth_type": {
							Description:  fmt.Sprintf("Type of API authentication. Valid values are %s. Defaults to `UsernameToken` for backward compatibility.", strings.Join(nextgen.BitBucketApiAccessTypeValues, ", ")),
							Type:         schema.TypeString,
							Optional:     true,
							Default:      nextgen.BitBucketApiAccessTypes.UsernameToken.String(),
							ValidateFunc: validation.StringInSlice(nextgen.BitBucketApiAccessTypeValues, false),
						},
						"username": {
							Description:   "The username used for connecting to the api. Applicable when `auth_type` is `UsernameToken`.",
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"api_authentication.0.username_ref"},
						},
						"username_ref": {
							Description:   "The name of the Harness secret containing the username. Applicable when `auth_type` is `UsernameToken`." + secretRefText,
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"api_authentication.0.username"},
						},
						"email": {
							Description:   "The email used for connecting to the api. Applicable when `auth_type` is `EmailAndApiToken`.",
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"api_authentication.0.email_ref"},
						},
						"email_ref": {
							Description:   "The name of the Harness secret containing the email. Applicable when `auth_type` is `EmailAndApiToken`." + secretRefText,
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"api_authentication.0.email"},
						},
						"token_ref": {
							Description: "Reference to a Harness secret containing the personal access token (or API token for `EmailAndApiToken`/`AccessToken`) for interacting with the BitBucket api." + secretRefText,
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
										Description:   "Reference to a secret containing the username to use for authentication." + secretRefText,
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.http.0.username"},
										ExactlyOneOf:  []string{"credentials.0.http.0.username", "credentials.0.http.0.username_ref"},
									},
									"password_ref": {
										Description: "Reference to a secret containing the password to use for authentication." + secretRefText,
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
							ConflictsWith: []string{"credentials.0.http"},
							ExactlyOneOf:  []string{"credentials.0.ssh", "credentials.0.http"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssh_key_ref": {
										Description: "Reference to the Harness secret containing the ssh key." + secretRefText,
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

func resourceConnectorBitbucketRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Bitbucket)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorBitbucket(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorBitbucketCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorBitbucket(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorBitbucket(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorBitbucket(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:     nextgen.ConnectorTypes.Bitbucket,
		BitBucket: &nextgen.BitbucketConnector{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.BitBucket.Url = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.BitBucket.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("validation_repo"); ok {
		connector.BitBucket.ValidationRepo = attr.(string)
	}

	if attr, ok := d.GetOk("connection_type"); ok {
		connector.BitBucket.Type_ = nextgen.GitConnectorType(attr.(string))
	}

	if attr, ok := d.GetOk("credentials"); ok {
		credConfig := attr.([]interface{})[0].(map[string]interface{})
		connector.BitBucket.Authentication = &nextgen.BitbucketAuthentication{}

		if attr := credConfig["http"].([]interface{}); len(attr) > 0 {
			httpConfig := attr[0].(map[string]interface{})
			connector.BitBucket.Authentication.Type_ = nextgen.GitAuthTypes.Http
			connector.BitBucket.Authentication.Http = &nextgen.BitbucketHttpCredentials{
				Type_:            nextgen.BitBucketHttpCredentialTypes.UsernamePassword,
				UsernamePassword: &nextgen.BitbucketUsernamePassword{},
			}

			if attr, ok := httpConfig["username"]; ok {
				connector.BitBucket.Authentication.Http.UsernamePassword.Username = attr.(string)
			}

			if attr, ok := httpConfig["username_ref"]; ok {
				connector.BitBucket.Authentication.Http.UsernamePassword.UsernameRef = attr.(string)
			}

			if attr, ok := httpConfig["password_ref"]; ok {
				connector.BitBucket.Authentication.Http.UsernamePassword.PasswordRef = attr.(string)
			}
		}

		if attr := credConfig["ssh"].([]interface{}); len(attr) > 0 {
			sshConfig := attr[0].(map[string]interface{})
			connector.BitBucket.Authentication.Type_ = nextgen.GitAuthTypes.Ssh
			connector.BitBucket.Authentication.Ssh = &nextgen.BitbucketSshCredentials{}

			if attr, ok := sshConfig["ssh_key_ref"]; ok {
				connector.BitBucket.Authentication.Ssh.SshKeyRef = attr.(string)
			}
		}
	}

	if attr, ok := d.GetOk("api_authentication"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		authType := nextgen.BitBucketApiAccessTypes.UsernameToken
		if v, ok := config["auth_type"].(string); ok && v != "" {
			authType = nextgen.BitBucketApiAccessType(v)
		}

		connector.BitBucket.ApiAccess = &nextgen.BitbucketApiAccess{
			Type_: authType,
		}

		tokenRef, _ := config["token_ref"].(string)

		switch authType {
		case nextgen.BitBucketApiAccessTypes.UsernameToken:
			connector.BitBucket.ApiAccess.UsernameToken = &nextgen.BitbucketUsernameTokenApiAccess{
				Username:    config["username"].(string),
				UsernameRef: config["username_ref"].(string),
				TokenRef:    tokenRef,
			}
		case nextgen.BitBucketApiAccessTypes.AccessToken:
			connector.BitBucket.ApiAccess.AccessToken = &nextgen.BitbucketAccessTokenApiAccess{
				TokenRef: tokenRef,
			}
		case nextgen.BitBucketApiAccessTypes.EmailAndApiToken:
			connector.BitBucket.ApiAccess.EmailApiToken = &nextgen.BitbucketEmailApiTokenApiAccess{
				Email:    config["email"].(string),
				EmailRef: config["email_ref"].(string),
				TokenRef: tokenRef,
			}
		}
	}

	return connector

}

func readConnectorBitbucket(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.BitBucket.Url)
	d.Set("connection_type", connector.BitBucket.Type_.String())
	d.Set("delegate_selectors", connector.BitBucket.DelegateSelectors)
	d.Set("validation_repo", connector.BitBucket.ValidationRepo)

	if connector.BitBucket.Authentication != nil {
		switch connector.BitBucket.Authentication.Type_ {
		case nextgen.GitAuthTypes.Http:
			switch connector.BitBucket.Authentication.Http.Type_ {
			case nextgen.BitBucketHttpCredentialTypes.UsernamePassword:
				d.Set("credentials", []map[string]interface{}{
					{
						"http": []map[string]interface{}{
							{
								"username":     connector.BitBucket.Authentication.Http.UsernamePassword.Username,
								"username_ref": connector.BitBucket.Authentication.Http.UsernamePassword.UsernameRef,
								"password_ref": connector.BitBucket.Authentication.Http.UsernamePassword.PasswordRef,
							},
						},
					},
				})
			default:
				return fmt.Errorf("unsupported BitBucket http authentication type: %s", connector.BitBucket.Authentication.Type_)
			}

		case nextgen.GitAuthTypes.Ssh:
			d.Set("credentials", []map[string]interface{}{
				{
					"ssh": []map[string]interface{}{
						{
							"ssh_key_ref": connector.BitBucket.Authentication.Ssh.SshKeyRef,
						},
					},
				},
			})
		default:
			return fmt.Errorf("unsupported Bitbucket auth type: %s", connector.BitBucket.Type_)
		}
	}

	if connector.BitBucket.ApiAccess != nil {
		switch connector.BitBucket.ApiAccess.Type_ {
		case nextgen.BitBucketApiAccessTypes.UsernameToken:
			d.Set("api_authentication", []map[string]interface{}{
				{
					"auth_type":    connector.BitBucket.ApiAccess.Type_.String(),
					"username":     connector.BitBucket.ApiAccess.UsernameToken.Username,
					"username_ref": connector.BitBucket.ApiAccess.UsernameToken.UsernameRef,
					"token_ref":    connector.BitBucket.ApiAccess.UsernameToken.TokenRef,
				},
			})
		case nextgen.BitBucketApiAccessTypes.AccessToken:
			d.Set("api_authentication", []map[string]interface{}{
				{
					"auth_type": connector.BitBucket.ApiAccess.Type_.String(),
					"token_ref": connector.BitBucket.ApiAccess.AccessToken.TokenRef,
				},
			})
		case nextgen.BitBucketApiAccessTypes.EmailAndApiToken:
			d.Set("api_authentication", []map[string]interface{}{
				{
					"auth_type": connector.BitBucket.ApiAccess.Type_.String(),
					"email":     connector.BitBucket.ApiAccess.EmailApiToken.Email,
					"email_ref": connector.BitBucket.ApiAccess.EmailApiToken.EmailRef,
					"token_ref": connector.BitBucket.ApiAccess.EmailApiToken.TokenRef,
				},
			})
		default:
			return fmt.Errorf("unsupported BitBucket api access type: %s", connector.BitBucket.ApiAccess.Type_)
		}
	}

	return nil
}

func validateBitbucketApiAuthentication(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	raw, ok := d.GetOk("api_authentication")
	if !ok {
		return nil
	}
	list := raw.([]interface{})
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	cfg := list[0].(map[string]interface{})

	authType := cfg["auth_type"].(string)
	if authType == "" {
		authType = nextgen.BitBucketApiAccessTypes.UsernameToken.String()
	}

	username, _ := cfg["username"].(string)
	usernameRef, _ := cfg["username_ref"].(string)
	email, _ := cfg["email"].(string)
	emailRef, _ := cfg["email_ref"].(string)
	tokenRef, _ := cfg["token_ref"].(string)

	expected := map[string]string{
		nextgen.BitBucketApiAccessTypes.UsernameToken.String(): `Expected api_authentication block for auth_type "UsernameToken":
  api_authentication {
    auth_type    = "UsernameToken"
    username     = "<username>"      # OR username_ref (exactly one)
    # username_ref = "account.<secret>"
    token_ref    = "account.<secret>"
  }`,
		nextgen.BitBucketApiAccessTypes.AccessToken.String(): `Expected api_authentication block for auth_type "AccessToken":
  api_authentication {
    auth_type = "AccessToken"
    token_ref = "account.<secret>"
  }`,
		nextgen.BitBucketApiAccessTypes.EmailAndApiToken.String(): `Expected api_authentication block for auth_type "EmailAndApiToken":
  api_authentication {
    auth_type  = "EmailAndApiToken"
    email      = "user@example.com"   # OR email_ref (exactly one)
    # email_ref = "account.<secret>"
    token_ref  = "account.<secret>"
  }`,
	}

	fail := func(reason string) error {
		return fmt.Errorf("%s\n\n%s", reason, expected[authType])
	}

	switch authType {
	case nextgen.BitBucketApiAccessTypes.UsernameToken.String():
		if email != "" || emailRef != "" {
			return fail(`"email"/"email_ref" are not allowed for auth_type "UsernameToken".`)
		}
		if username == "" && usernameRef == "" {
			return fail(`Missing credential: one of "username" or "username_ref" is required.`)
		}
		if username != "" && usernameRef != "" {
			return fail(`Set exactly one of "username" or "username_ref", not both.`)
		}
		if tokenRef == "" {
			return fail(`Missing required field "token_ref".`)
		}

	case nextgen.BitBucketApiAccessTypes.AccessToken.String():
		var extras []string
		if username != "" {
			extras = append(extras, `"username"`)
		}
		if usernameRef != "" {
			extras = append(extras, `"username_ref"`)
		}
		if email != "" {
			extras = append(extras, `"email"`)
		}
		if emailRef != "" {
			extras = append(extras, `"email_ref"`)
		}
		if len(extras) > 0 {
			return fail(fmt.Sprintf(`Unsupported field(s) for auth_type "AccessToken": %s.`, strings.Join(extras, ", ")))
		}
		if tokenRef == "" {
			return fail(`Missing required field "token_ref".`)
		}

	case nextgen.BitBucketApiAccessTypes.EmailAndApiToken.String():
		if username != "" || usernameRef != "" {
			return fail(`"username"/"username_ref" are not allowed for auth_type "EmailAndApiToken".`)
		}
		if email == "" && emailRef == "" {
			return fail(`Missing credential: one of "email" or "email_ref" is required.`)
		}
		if email != "" && emailRef != "" {
			return fail(`Set exactly one of "email" or "email_ref", not both.`)
		}
		if tokenRef == "" {
			return fail(`Missing required field "token_ref".`)
		}
	}
	return nil
}
