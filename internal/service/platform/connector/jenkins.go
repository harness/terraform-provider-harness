package connector

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceConnectorJenkins() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating an Jenkins connector.",
		ReadContext:   resourceConnectorJenkinsRead,
		CreateContext: resourceConnectorCreateOrUpdate,
		UpdateContext: resourceConnectorCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"jenkins_url": {
				Description: "Jenkins Url.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"auth": {
				Description: "This entity contains the details for Jenkins Authentication.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description:  "Can be one of UsernamePassword, Anonymous, Bearer Token(HTTP Header)",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"UsernamePassword", "Anonymous", "Bearer Token(HTTP Header)"}, false),
						},
						"jenkins_bearer_token": {
							Description:   "Authenticate to App Dynamics using bearer token.",
							Type:          schema.TypeList,
							Optional:      true,
							MaxItems:      1,
							ConflictsWith: []string{"auth.0.jenkins_user_name_password"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"token_ref": {
										Description: "Reference of the token." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"jenkins_user_name_password": {
							Description:   "Authenticate to App Dynamics using user name and password.",
							Type:          schema.TypeList,
							Optional:      true,
							MaxItems:      1,
							ConflictsWith: []string{"auth.0.jenkins_bearer_token"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Description: "Username to use for authentication.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"username_ref": {
										Description: "Username reference to use for authentication.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"password_ref": {
										Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
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

func resourceConnectorJenkinsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.Jenkins)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorJenkins(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorJenkins(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorJenkins(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorJenkins(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:   nextgen.ConnectorTypes.Jenkins,
		Jenkins: &nextgen.JenkinsConnector{ConnectorType: nextgen.ConnectorTypes.Jenkins.String()},
	}

	if attr, ok := d.GetOk("jenkins_url"); ok {
		connector.Jenkins.JenkinsUrl = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.Jenkins.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("auth"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		connector.Jenkins.Auth = &nextgen.JenkinsAuthentication{}

		if attr, ok := config["type"]; ok {
			connector.Jenkins.Auth.Type_ = attr.(string)
		}

		if attr, ok := config["jenkins_bearer_token"]; ok {
			if len(attr.([]interface{})) != 0 {
				config := attr.([]interface{})[0].(map[string]interface{})

				connector.Jenkins.Auth.JenkinsBearerToken = &nextgen.JenkinsBearerTokenDto{}
				if attr, ok := config["token_ref"]; ok {
					connector.Jenkins.Auth.JenkinsBearerToken.TokenRef = attr.(string)
				}
			}
		}

		if attr, ok := config["jenkins_user_name_password"]; ok {
			if len(attr.([]interface{})) != 0 {
				config := attr.([]interface{})[0].(map[string]interface{})

				connector.Jenkins.Auth.JenkinsUserNamePassword = &nextgen.JenkinsUserNamePasswordDto{}
				if attr, ok := config["username"]; ok {
					connector.Jenkins.Auth.JenkinsUserNamePassword.Username = attr.(string)
				}

				if attr, ok := config["username_ref"]; ok {
					connector.Jenkins.Auth.JenkinsUserNamePassword.UsernameRef = attr.(string)
				}

				if attr, ok := config["password_ref"]; ok {
					connector.Jenkins.Auth.JenkinsUserNamePassword.PasswordRef = attr.(string)
				}
			}
		}
	}

	return connector
}

func readConnectorJenkins(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("jenkins_url", connector.Jenkins.JenkinsUrl)
	d.Set("auth", []interface{}{
		map[string]interface{}{
			"type":                       connector.Jenkins.Auth.Type_,
			"jenkins_user_name_password": readJenkinsUsernamePassword(d, connector),
			"jenkins_bearer_token":       readJenkinsBearerToken(d, connector),
		},
	})

	d.Set("delegate_selectors", connector.Jenkins.DelegateSelectors)

	return nil
}

func readJenkinsUsernamePassword(d *schema.ResourceData, connector *nextgen.ConnectorInfo) []interface{} {
	var spec []interface{}
	switch connector.Jenkins.Auth.Type_ {
	case "UsernamePassword":
		spec = []interface{}{
			map[string]interface{}{
				"username":     connector.Jenkins.Auth.JenkinsUserNamePassword.Username,
				"username_ref": connector.Jenkins.Auth.JenkinsUserNamePassword.UsernameRef,
				"password_ref": connector.Jenkins.Auth.JenkinsUserNamePassword.PasswordRef,
			},
		}
	case "Bearer Token(HTTP Header)":
		//noop
	}
	return spec
}

func readJenkinsBearerToken(d *schema.ResourceData, connector *nextgen.ConnectorInfo) []interface{} {
	var spec []interface{}
	switch connector.Jenkins.Auth.Type_ {
	case "UsernamePassword":
		//noop
	case "Bearer Token(HTTP Header)":
		spec = []interface{}{
			map[string]interface{}{
				"token_ref": connector.Jenkins.Auth.JenkinsBearerToken.TokenRef,
			},
		}
	}
	return spec
}
