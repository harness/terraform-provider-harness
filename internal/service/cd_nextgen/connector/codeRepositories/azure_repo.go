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

func ResourceConnectorAzureRepo() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Azure Repo connector.",
		ReadContext:   resourceConnectorAzureRepoRead,
		CreateContext: resourceConnectorAzureRepoCreateOrUpdate,
		UpdateContext: resourceConnectorAzureRepoCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the azure repository or account.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"connection_type": {
				Description:  fmt.Sprintf("Whether the connection we're making is to a azure repository or a azure account. Valid values are %s.", strings.Join(nextgen.GitConnectorTypeValues, ", ")),
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
			"execute_on_delegate": {
				Description: "Execute on delegate or not.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"api_authentication": {
				Description: "Configuration for using the azure api. API Access is required for using “Git Experience”, for creation of Git based triggers, Webhooks management and updating Git statuses.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"token_ref": {
							Description: "Personal access token for interacting with the azure api." + secretRefText,
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
									"token_ref": {
										Description:   "Reference to a secret containing the personal access to use for authentication." + secretRefText,
										Type:          schema.TypeString,
										Required:      true,
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

func resourceConnectorAzureRepoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Gitlab)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorAzureRepo(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorAzureRepoCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorAzureRepo(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorAzureRepo(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorAzureRepo(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:  nextgen.ConnectorTypes.AzureRepo,
		AzureRepo: &nextgen.AzureRepoConnector{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.AzureRepo.Url = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.AzureRepo.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("execute_on_delegate"); ok {
		connector.AzureRepo.ExecuteOnDelegate = attr.(bool)
	}

	if attr, ok := d.GetOk("validation_repo"); ok {
		connector.AzureRepo.ValidationRepo = attr.(string)
	}

	if attr, ok := d.GetOk("connection_type"); ok {
		connector.AzureRepo.Type_ = nextgen.GitConnectorType(attr.(string))
	}

	if attr, ok := d.GetOk("credentials"); ok {
		credConfig := attr.([]interface{})[0].(map[string]interface{})
		connector.AzureRepo.Authentication = &nextgen.AzureRepoAuthentication{}

		if attr := credConfig["http"].([]interface{}); len(attr) > 0 {
			httpConfig := attr[0].(map[string]interface{})
			connector.AzureRepo.Authentication.Type_ = nextgen.GitAuthTypes.Http
			connector.AzureRepo.Authentication.Http = &nextgen.AzureRepoHttpCredentials{}

			if attr := httpConfig["token_ref"].(string); attr != "" {
				connector.AzureRepo.Authentication.Http.Type_ = nextgen.AzureRepoHttpCredentialTypes.UsernameToken
				connector.AzureRepo.Authentication.Http.UsernameToken = &nextgen.AzureRepoUsernameToken{
					TokenRef: attr,
				}

				if attr := httpConfig["username"].(string); attr != "" {
					connector.AzureRepo.Authentication.Http.UsernameToken.Username = attr
				}

				if attr := httpConfig["username_ref"].(string); attr != "" {
					connector.AzureRepo.Authentication.Http.UsernameToken.UsernameRef = attr
				}
			}
		}

		if attr := credConfig["ssh"].([]interface{}); len(attr) > 0 {
			sshConfig := attr[0].(map[string]interface{})
			connector.AzureRepo.Authentication.Type_ = nextgen.GitAuthTypes.Ssh
			connector.AzureRepo.Authentication.Ssh = &nextgen.AzureRepoSshCredentials{}

			if attr := sshConfig["ssh_key_ref"].(string); attr != "" {
				connector.AzureRepo.Authentication.Ssh.SshKeyRef = attr
			}
		}
	}

	if attr, ok := d.GetOk("api_authentication"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.AzureRepo.ApiAccess = &nextgen.AzureRepoApiAccess{
			Type_: nextgen.AzureRepoApiAuthTypes.Token,
			Token: &nextgen.AzureRepoTokenSpec{},
		}

		if attr, ok := config["token_ref"]; ok {
			connector.AzureRepo.ApiAccess.Token.TokenRef = attr.(string)
		}
	}

	return connector
}

func readConnectorAzureRepo(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.AzureRepo.Url)
	d.Set("connection_type", string(connector.AzureRepo.Type_))
	d.Set("delegate_selectors", connector.AzureRepo.DelegateSelectors)
	d.Set("execute_on_delegate", connector.AzureRepo.ExecuteOnDelegate)
	d.Set("validation_repo", connector.AzureRepo.ValidationRepo)

	if connector.AzureRepo.Authentication != nil {
		switch connector.AzureRepo.Authentication.Type_ {
		case nextgen.GitAuthTypes.Http:
			switch connector.AzureRepo.Authentication.Http.Type_ {
			case nextgen.AzureRepoHttpCredentialTypes.UsernameToken:
				d.Set("credentials", []map[string]interface{}{
					{
						"http": []map[string]interface{}{
							{
								"username":     connector.AzureRepo.Authentication.Http.UsernameToken.Username,
								"username_ref": connector.AzureRepo.Authentication.Http.UsernameToken.UsernameRef,
								"token_ref":    connector.AzureRepo.Authentication.Http.UsernameToken.TokenRef,
							},
						},
					},
				})
			default:
				return fmt.Errorf("unsupported azure repo http authentication type: %s", connector.AzureRepo.Authentication.Type_)
			}

		case nextgen.GitAuthTypes.Ssh:
			d.Set("credentials", []map[string]interface{}{
				{
					"ssh": []map[string]interface{}{
						{
							"ssh_key_ref": connector.AzureRepo.Authentication.Ssh.SshKeyRef,
						},
					},
				},
			})
		default:
			return fmt.Errorf("unsupported git auth type: %s", connector.AzureRepo.Type_)
		}
	}

	if connector.AzureRepo.ApiAccess != nil {
		switch connector.AzureRepo.ApiAccess.Type_ {
		case nextgen.AzureRepoApiAuthTypes.Token:
			d.Set("api_authentication", []map[string]interface{}{
				{
					"token_ref": connector.AzureRepo.ApiAccess.Token.TokenRef,
				},
			})
		default:
			return fmt.Errorf("unsupported azure repo api access type: %s", connector.AzureRepo.ApiAccess.Type_)
		}
	}

	return nil
}
