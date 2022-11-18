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

func ResourceConnectorGit() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Git connector.",
		ReadContext:   resourceConnectorGitRead,
		CreateContext: resourceConnectorGitCreateOrUpdate,
		UpdateContext: resourceConnectorGitCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the git repository or account.",
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
				Description: "Tags to filter delegates for connection.",
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

func resourceConnectorGitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Git)
	if err != nil {
		return err
	}

	if err := readConnectorGit(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorGitCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorGit(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorGit(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorGit(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Git,
		Git:   &nextgen.GitConfig{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.Git.Url = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Git.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("validation_repo"); ok {
		connector.Git.ValidationRepo = attr.(string)
	}

	if attr, ok := d.GetOk("connection_type"); ok {
		connector.Git.ConnectionType = attr.(string)
	}

	if attr, ok := d.GetOk("credentials"); ok {
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

	return connector
}

func readConnectorGit(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {

	d.Set("url", connector.Git.Url)
	d.Set("connection_type", connector.Git.ConnectionType)
	d.Set("delegate_selectors", connector.Git.DelegateSelectors)
	d.Set("validation_repo", connector.Git.ValidationRepo)

	switch connector.Git.Type_ {
	case nextgen.GitAuthTypes.Http:
		d.Set("credentials", []map[string]interface{}{
			{
				"http": []map[string]interface{}{
					{
						"username":     connector.Git.Http.Username,
						"username_ref": connector.Git.Http.UsernameRef,
						"password_ref": connector.Git.Http.PasswordRef,
					},
				},
			},
		})
	case nextgen.GitAuthTypes.Ssh:
		d.Set("credentials", []map[string]interface{}{
			{
				"ssh": []map[string]interface{}{
					{
						"ssh_key_ref": connector.Git.Ssh.SshKeyRef,
					},
				},
			},
		})
	default:
		return fmt.Errorf("unsupported git auth type: %s", connector.Git.Type_)
	}

	return nil
}
