package connector

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceConnectorJira() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Jira connector.",
		ReadContext:   resourceConnectorJiraRead,
		CreateContext: resourceConnectorJiraCreateOrUpdate,
		UpdateContext: resourceConnectorJiraCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "URL of the Jira server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"username": {
				Description:   "Username to use for authentication.",
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
			},
			"username_ref": {
				Description:   "Reference to a secret containing the username to use for authentication." + secret_ref_text,
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,				
			},
			"password_ref": {
				Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
				Type:        schema.TypeString,
				Optional:      true,
				Computed:      true,				
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"auth": {
				Description: "The credentials to use for the jira authentication.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Description: "Authentication types for Jira connector",
							Type:        schema.TypeString,
							Required:      true,	
							ValidateFunc: validation.StringInSlice([]string{"UsernamePassword","PersonalAccessToken"}, false),						
						},
						"username_password": {
							Description:   "Authenticate using username password.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"auth.0.personal_access_token"},
							AtLeastOneOf: []string{
								"auth.0.username_password",
								"auth.0.personal_access_token",
							},							
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Description:   "Username to use for authentication.",
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"auth.0.username_password.0.username_ref"},
										AtLeastOneOf: []string{
											"auth.0.username_password.0.username",
											"auth.0.username_password.0.username_ref",
										},
									},
									"username_ref": {
										Description:   "Reference to a secret containing the username to use for authentication." + secret_ref_text,
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"auth.0.username_password.0.username"},
										AtLeastOneOf: []string{
											"auth.0.username_password.0.username",
											"auth.0.username_password.0.username_ref",
										},
									},
									"password_ref": {
										Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},	
						"personal_access_token": {
							Description:   "Authenticate using personal access token.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"auth.0.username_password"},
							AtLeastOneOf: []string{
								"auth.0.username_password",
								"auth.0.personal_access_token",
							},	
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pat_ref": {
										Description: "Reference to a secret containing the personal access token to use for authentication." + secret_ref_text,
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

func resourceConnectorJiraRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Jira)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorJira(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorJiraCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorJira(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorJira(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorJira(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.Jira,
		Jira:  &nextgen.JiraConnector{},
	}

	if attr, ok := d.GetOk("url"); ok {
		connector.Jira.JiraUrl = attr.(string)
	}

	if attr, ok := d.GetOk("username"); ok {
		connector.Jira.Username = attr.(string)
	}

	if attr, ok := d.GetOk("username_ref"); ok {
		connector.Jira.UsernameRef = attr.(string)
	}

	if attr, ok := d.GetOk("password_ref"); ok {
		connector.Jira.PasswordRef = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Jira.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("auth"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.Jira.Auth = &nextgen.JiraAuthentication{}
		if attrAuthType, ok := config["auth_type"]; ok {
			if attrAuthType.(string) == "UsernamePassword" {
				connector.Jira.Auth.Type_ = nextgen.JiraAuthTypes.UsernamePassword
			}
			if attrAuthType.(string) == "PersonalAccessToken" {
				connector.Jira.Auth.Type_ = nextgen.JiraAuthTypes.PersonalAccessToken
			}
		}
		if config["auth_type"] == "UsernamePassword" {
			if attrUsernamePassword, ok := config["username_password"]; ok {
				configUsernamePassword := attrUsernamePassword.([]interface{})[0].(map[string]interface{})
				connector.Jira.Auth.UsernamePassword = &nextgen.JiraUserNamePassword{}

				if attr, ok := configUsernamePassword["username"]; ok {
					connector.Jira.Auth.UsernamePassword.Username = attr.(string)
				}

				if attr, ok := configUsernamePassword["username_ref"]; ok {
					connector.Jira.Auth.UsernamePassword.UsernameRef = attr.(string)
				}

				if attr, ok := configUsernamePassword["password_ref"]; ok {
					connector.Jira.Auth.UsernamePassword.PasswordRef = attr.(string)
				}
			}
		}

		if config["auth_type"] == "PersonalAccessToken" {
			if attrPatPassword, ok := config["personal_access_token"]; ok {
				configPatPassword := attrPatPassword.([]interface{})[0].(map[string]interface{})
				connector.Jira.Auth.PersonalAccessToken = &nextgen.JiraPatPassword{}

				if attr, ok := configPatPassword["pat_ref"]; ok {
					connector.Jira.Auth.PersonalAccessToken.PatRef = attr.(string)
				}
			}
		}		

	}

	return connector
}

func readConnectorJira(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {

	d.Set("url", connector.Jira.JiraUrl)
	d.Set("username", connector.Jira.Username)
	d.Set("username_ref", connector.Jira.UsernameRef)
	d.Set("password_ref", connector.Jira.PasswordRef)
	d.Set("delegate_selectors", connector.Jira.DelegateSelectors)

	switch connector.Jira.Auth.Type_ {
	case nextgen.JiraAuthTypes.UsernamePassword:
		d.Set("auth", []map[string]interface{}{
			{
			"auth_type" : "UsernamePassword",
			"username_password" : []map[string]interface{}{
				{
					"username":     connector.Jira.Auth.UsernamePassword.Username,
					"username_ref": connector.Jira.Auth.UsernamePassword.UsernameRef,
					"password_ref": connector.Jira.Auth.UsernamePassword.PasswordRef,
				},
			},
		},
		})
	case nextgen.JiraAuthTypes.PersonalAccessToken:
		d.Set("auth", []map[string]interface{}{
			{
			"auth_type" : "PersonalAccessToken",
			"personal_access_token" : []map[string]interface{}{
				{
					"pat_ref": connector.Jira.Auth.PersonalAccessToken.PatRef,
				},
			},
		},
		})		
	default:
		return fmt.Errorf("unsupported jira auth type: %s", connector.Jira.Auth.Type_)
	}

	return nil
}
