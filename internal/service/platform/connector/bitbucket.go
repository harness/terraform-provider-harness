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

func ResourceConnectorBitbucket() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Bitbucket connector.",
		ReadContext:   resourceConnectorBitbucketRead,
		CreateContext: resourceConnectorBitbucketCreateOrUpdate,
		UpdateContext: resourceConnectorBitbucketCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

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
						"username": {
							Description:   "The username used for connecting to the api.",
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"api_authentication.0.username_ref"},
							AtLeastOneOf:  []string{"api_authentication.0.username", "api_authentication.0.username_ref"},
						},
						"username_ref": {
							Description:   "The name of the Harness secret containing the username." + secret_ref_text,
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"api_authentication.0.username"},
							AtLeastOneOf:  []string{"api_authentication.0.username", "api_authentication.0.username_ref"},
						},
						"token_ref": {
							Description: "Personal access token for interacting with the BitBucket api." + secret_ref_text,
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
										Description:   "Reference to a secret containing the username to use for authentication." + secret_ref_text,
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.http.0.username"},
										ExactlyOneOf:  []string{"credentials.0.http.0.username", "credentials.0.http.0.username_ref"},
									},
									"password_ref": {
										Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
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
		connector.BitBucket.ApiAccess = &nextgen.BitbucketApiAccess{
			Type_:         nextgen.BitBucketApiAccessTypes.UsernameToken,
			UsernameToken: &nextgen.BitbucketUsernameTokenApiAccess{},
		}

		if attr, ok := config["username"]; ok {
			connector.BitBucket.ApiAccess.UsernameToken.Username = attr.(string)
		}

		if attr, ok := config["username_ref"]; ok {
			connector.BitBucket.ApiAccess.UsernameToken.UsernameRef = attr.(string)
		}

		if attr, ok := config["token_ref"]; ok {
			connector.BitBucket.ApiAccess.UsernameToken.TokenRef = attr.(string)
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
					"username":     connector.BitBucket.ApiAccess.UsernameToken.Username,
					"username_ref": connector.BitBucket.ApiAccess.UsernameToken.UsernameRef,
					"token_ref":    connector.BitBucket.ApiAccess.UsernameToken.TokenRef,
				},
			})
		default:
			return fmt.Errorf("unsupported BitBucket api access type: %s", connector.BitBucket.ApiAccess.Type_)
		}
	}

	return nil
}
