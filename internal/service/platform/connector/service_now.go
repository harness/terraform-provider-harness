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

func ResourceConnectorServiceNow() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a Service Now connector.",
		ReadContext:   resourceConnectorServiceNowRead,
		CreateContext: resourceConnectorServiceNowCreateOrUpdate,
		UpdateContext: resourceConnectorServiceNowCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"service_now_url": {
				Description: "URL of service now.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"username": {
				Description: "Username to use for authentication.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"username_ref": {
				Description: "Reference to a secret containing the username to use for authentication." + secret_ref_text,
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"password_ref": {
				Description: "Reference to a secret containing the password to use for authentication." + secret_ref_text,
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"auth": {
				Description: "The credentials to use for the service now authentication.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Description:  "Authentication types for Jira connector",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"UsernamePassword", "AdfsClientCredentialsWithCertificate"}, false),
						},
						"username_password": {
							Description:   "Authenticate using username password.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"auth.0.adfs"},
							AtLeastOneOf: []string{
								"auth.0.username_password",
								"auth.0.adfs",
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
						"adfs": {
							Description:   "Authenticate using adfs client credentials with certificate.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"auth.0.username_password"},
							AtLeastOneOf: []string{
								"auth.0.username_password",
								"auth.0.adfs",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate_ref": {
										Description: "Reference to a secret containing the certificate to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
									"private_key_ref": {
										Description: "Reference to a secret containing the privateKeyRef to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
									"client_id_ref": {
										Description: "Reference to a secret containing the clientIdRef to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
									"resource_id_ref": {
										Description: "Reference to a secret containing the resourceIdRef to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
									"adfs_url": {
										Description: "asdf URL.",
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

func resourceConnectorServiceNowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.ServiceNow)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorServiceNow(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorServiceNowCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorServiceNow(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorServiceNow(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorServiceNow(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:      nextgen.ConnectorTypes.ServiceNow,
		ServiceNow: &nextgen.ServiceNowConnector{},
	}

	if attr, ok := d.GetOk("service_now_url"); ok {
		connector.ServiceNow.ServiceNowUrl = attr.(string)
	}

	if attr, ok := d.GetOk("username"); ok {
		connector.ServiceNow.Username = attr.(string)
	}

	if attr, ok := d.GetOk("username_ref"); ok {
		connector.ServiceNow.UsernameRef = attr.(string)
	}

	if attr, ok := d.GetOk("password_ref"); ok {
		connector.ServiceNow.PasswordRef = attr.(string)
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector.ServiceNow.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	if attr, ok := d.GetOk("auth"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.ServiceNow.Auth = &nextgen.ServiceNowAuthentication{}
		if attrAuthType, ok := config["auth_type"]; ok {
			if attrAuthType.(string) == "UsernamePassword" {
				connector.ServiceNow.Auth.Type_ = nextgen.ServiceNowAuthTypes.ServiceNowUserNamePassword
			}
			if attrAuthType.(string) == "AdfsClientCredentialsWithCertificate" {
				connector.ServiceNow.Auth.Type_ = nextgen.ServiceNowAuthTypes.ServiceNowAdfs
			}
		}
		if config["auth_type"] == "UsernamePassword" {
			if attrUsernamePassword, ok := config["username_password"]; ok {
				configUsernamePassword := attrUsernamePassword.([]interface{})[0].(map[string]interface{})
				connector.ServiceNow.Auth.ServiceNowUserNamePassword = &nextgen.ServiceNowUserNamePassword{}

				if attr, ok := configUsernamePassword["username"]; ok {
					connector.ServiceNow.Auth.ServiceNowUserNamePassword.Username = attr.(string)
				}

				if attr, ok := configUsernamePassword["username_ref"]; ok {
					connector.ServiceNow.Auth.ServiceNowUserNamePassword.UsernameRef = attr.(string)
				}

				if attr, ok := configUsernamePassword["password_ref"]; ok {
					connector.ServiceNow.Auth.ServiceNowUserNamePassword.PasswordRef = attr.(string)
				}
			}
		}
		if config["auth_type"] == "AdfsClientCredentialsWithCertificate" {
			if attrAdsf, ok := config["adfs"]; ok {
				configAsdf := attrAdsf.([]interface{})[0].(map[string]interface{})
				connector.ServiceNow.Auth.ServiceNowAdfs = &nextgen.ServiceNowAdfs{}

				if attr, ok := configAsdf["certificate_ref"]; ok {
					connector.ServiceNow.Auth.ServiceNowAdfs.CertificateRef = attr.(string)
				}

				if attr, ok := configAsdf["client_id_ref"]; ok {
					connector.ServiceNow.Auth.ServiceNowAdfs.ClientIdRef = attr.(string)
				}

				if attr, ok := configAsdf["private_key_ref"]; ok {
					connector.ServiceNow.Auth.ServiceNowAdfs.PrivateKeyRef = attr.(string)
				}

				if attr, ok := configAsdf["resource_id_ref"]; ok {
					connector.ServiceNow.Auth.ServiceNowAdfs.ResourceIdRef = attr.(string)
				}

				if attr, ok := configAsdf["adfs_url"]; ok {
					connector.ServiceNow.Auth.ServiceNowAdfs.AdfsUrl = attr.(string)
				}
			}
		}
	}
	return connector
}

func readConnectorServiceNow(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("service_now_url", connector.ServiceNow.ServiceNowUrl)
	d.Set("username", connector.ServiceNow.Username)
	d.Set("username_ref", connector.ServiceNow.UsernameRef)
	d.Set("password_ref", connector.ServiceNow.PasswordRef)
	d.Set("delegate_selectors", connector.ServiceNow.DelegateSelectors)

	switch connector.ServiceNow.Auth.Type_ {
	case "UsernamePassword":
		d.Set("auth", []map[string]interface{}{
			{
				"auth_type": "UsernamePassword",
				"username_password": []map[string]interface{}{
					{
						"username":     connector.ServiceNow.Auth.ServiceNowUserNamePassword.Username,
						"username_ref": connector.ServiceNow.Auth.ServiceNowUserNamePassword.UsernameRef,
						"password_ref": connector.ServiceNow.Auth.ServiceNowUserNamePassword.PasswordRef,
					},
				},
			},
		})
	case "AdfsClientCredentialsWithCertificate":
		d.Set("auth", []map[string]interface{}{
			{
				"auth_type": "AdfsClientCredentialsWithCertificate",
				"adfs": []map[string]interface{}{
					{
						"certificate_ref": connector.ServiceNow.Auth.ServiceNowAdfs.CertificateRef,
						"client_id_ref":   connector.ServiceNow.Auth.ServiceNowAdfs.ClientIdRef,
						"private_key_ref": connector.ServiceNow.Auth.ServiceNowAdfs.PrivateKeyRef,
						"resource_id_ref": connector.ServiceNow.Auth.ServiceNowAdfs.ResourceIdRef,
						"adfs_url":        connector.ServiceNow.Auth.ServiceNowAdfs.AdfsUrl,
					},
				},
			},
		})
	}
	return nil
}
