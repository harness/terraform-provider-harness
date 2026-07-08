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

func ResourceConnectorJDBC() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a JDBC connector.",
		ReadContext:   resourceConnectorJDBCRead,
		CreateContext: resourceConnectorJDBCCreateOrUpdate,
		UpdateContext: resourceConnectorJDBCCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"url": {
				Description: "The URL of the database server.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"delegate_selectors": {
				Description: "Tags to filter delegates for connection.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"credentials": {
				Description: "The credentials to use for the database server.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_type": {
							Description: "Authentication types for JDBC connector",
							Type:        schema.TypeString,
							Optional:    true,
							Default:     nextgen.JDBCAuthTypes.UsernamePassword.String(),
							ValidateFunc: validation.StringInSlice([]string{
								nextgen.JDBCAuthTypes.UsernamePassword.String(),
								nextgen.JDBCAuthTypes.ServiceAccount.String(),
								nextgen.JDBCAuthTypes.KeyPair.String(),
								nextgen.JDBCAuthTypes.InheritFromDelegate.String(),
								nextgen.JDBCAuthTypes.Oidc.String(),
							}, false),
						},
						"username": {
							Description:   "The username to use for the database server.",
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username_ref", "credentials.0.username_password", "credentials.0.service_account", "credentials.0.key_pair", "credentials.0.inherit_from_delegate", "credentials.0.oidc"},
							AtLeastOneOf: []string{
								"credentials.0.username",
								"credentials.0.username_ref",
								"credentials.0.username_password",
								"credentials.0.service_account",
								"credentials.0.key_pair",
								"credentials.0.inherit_from_delegate",
								"credentials.0.oidc",
							},
						},
						"username_ref": {
							Description:   "The reference to the Harness secret containing the username to use for the database server." + secret_ref_text,
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username", "credentials.0.username_password", "credentials.0.service_account", "credentials.0.key_pair", "credentials.0.inherit_from_delegate", "credentials.0.oidc"},
							AtLeastOneOf: []string{
								"credentials.0.username",
								"credentials.0.username_ref",
								"credentials.0.username_password",
								"credentials.0.service_account",
								"credentials.0.key_pair",
								"credentials.0.inherit_from_delegate",
								"credentials.0.oidc",
							},
						},
						"password_ref": {
							Description:   "The reference to the Harness secret containing the password to use for the database server." + secret_ref_text,
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username_password", "credentials.0.service_account", "credentials.0.key_pair", "credentials.0.inherit_from_delegate", "credentials.0.oidc"},
							AtLeastOneOf: []string{
								"credentials.0.password_ref",
								"credentials.0.username_password",
								"credentials.0.service_account",
								"credentials.0.key_pair",
								"credentials.0.inherit_from_delegate",
								"credentials.0.oidc",
							},
						},
						"username_password": {
							Description:   "Authenticate using username password.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username", "credentials.0.username_ref", "credentials.0.password_ref", "credentials.0.service_account", "credentials.0.key_pair", "credentials.0.inherit_from_delegate", "credentials.0.oidc"},
							AtLeastOneOf: []string{
								"credentials.0.username_password",
								"credentials.0.service_account",
								"credentials.0.password_ref",
								"credentials.0.key_pair",
								"credentials.0.inherit_from_delegate",
								"credentials.0.oidc",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Description:   "Username to use for authentication.",
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.username_password.0.username_ref"},
										AtLeastOneOf: []string{
											"credentials.0.username_password.0.username",
											"credentials.0.username_password.0.username_ref",
										},
									},
									"username_ref": {
										Description:   "Reference to a secret containing the username to use for authentication." + secret_ref_text,
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.username_password.0.username"},
										AtLeastOneOf: []string{
											"credentials.0.username_password.0.username",
											"credentials.0.username_password.0.username_ref",
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
						"service_account": {
							Description:   "Authenticate using service account.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username", "credentials.0.username_ref", "credentials.0.password_ref", "credentials.0.username_password", "credentials.0.key_pair", "credentials.0.inherit_from_delegate", "credentials.0.oidc"},
							AtLeastOneOf: []string{
								"credentials.0.username_password",
								"credentials.0.service_account",
								"credentials.0.password_ref",
								"credentials.0.key_pair",
								"credentials.0.inherit_from_delegate",
								"credentials.0.oidc",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"token_ref": {
										Description: "Reference to a secret containing the token to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"key_pair": {
							Description:   "Authenticate using key pair.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username", "credentials.0.username_ref", "credentials.0.password_ref", "credentials.0.username_password", "credentials.0.service_account", "credentials.0.inherit_from_delegate", "credentials.0.oidc"},
							AtLeastOneOf: []string{
								"credentials.0.username_password",
								"credentials.0.service_account",
								"credentials.0.password_ref",
								"credentials.0.key_pair",
								"credentials.0.inherit_from_delegate",
								"credentials.0.oidc",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Description:   "Username to use for authentication.",
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.key_pair.0.username_ref"},
										AtLeastOneOf: []string{
											"credentials.0.key_pair.0.username",
											"credentials.0.key_pair.0.username_ref",
										},
									},
									"username_ref": {
										Description:   "Reference to a secret containing the username to use for authentication." + secret_ref_text,
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.key_pair.0.username"},
										AtLeastOneOf: []string{
											"credentials.0.key_pair.0.username",
											"credentials.0.key_pair.0.username_ref",
										},
									},
									"private_key_file_ref": {
										Description: "Reference to a secret containing the private key file to use for authentication." + secret_ref_text,
										Type:        schema.TypeString,
										Required:    true,
									},
									"private_key_passphrase_ref": {
										Description: "Reference to a secret containing the passphrase for the private key file." + secret_ref_text,
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"inherit_from_delegate": {
							Description:   "Authenticate using credentials inherited from the Harness delegate runtime identity (e.g. GCP ADC, AWS IAM).",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username", "credentials.0.username_ref", "credentials.0.password_ref", "credentials.0.username_password", "credentials.0.service_account", "credentials.0.key_pair", "credentials.0.oidc"},
							AtLeastOneOf: []string{
								"credentials.0.username_password",
								"credentials.0.service_account",
								"credentials.0.password_ref",
								"credentials.0.key_pair",
								"credentials.0.inherit_from_delegate",
								"credentials.0.oidc",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Description:   "Username to use for authentication.",
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.inherit_from_delegate.0.username_ref"},
									},
									"username_ref": {
										Description:   "Reference to a secret containing the username to use for authentication." + secret_ref_text,
										Type:          schema.TypeString,
										Optional:      true,
										ConflictsWith: []string{"credentials.0.inherit_from_delegate.0.username"},
									},
								},
							},
						},
						"oidc": {
							Description:   "Authenticate using OIDC.",
							Type:          schema.TypeList,
							MaxItems:      1,
							Optional:      true,
							ConflictsWith: []string{"credentials.0.username", "credentials.0.username_ref", "credentials.0.password_ref", "credentials.0.username_password", "credentials.0.service_account", "credentials.0.key_pair", "credentials.0.inherit_from_delegate"},
							AtLeastOneOf: []string{
								"credentials.0.username_password",
								"credentials.0.service_account",
								"credentials.0.password_ref",
								"credentials.0.key_pair",
								"credentials.0.inherit_from_delegate",
								"credentials.0.oidc",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"provider_type": {
										Description:  "The OIDC provider type. Currently supported: Gcp.",
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"Gcp"}, false),
									},
									"gcp_oidc": {
										Description: "GCP OIDC configuration. Required when provider_type is Gcp.",
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"project_number": {
													Description: "The GCP project number (numeric).",
													Type:        schema.TypeString,
													Required:    true,
												},
												"workload_pool_id": {
													Description: "The Workload Identity Pool ID.",
													Type:        schema.TypeString,
													Required:    true,
												},
												"provider_id": {
													Description: "The OIDC Provider ID within the pool.",
													Type:        schema.TypeString,
													Required:    true,
												},
												"service_account_email": {
													Description: "The GCP Service Account email for impersonation.",
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
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorJDBCRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.JDBC)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorJDBC(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorJDBCCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorJDBC(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorJDBC(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorJDBC(d *schema.ResourceData) *nextgen.ConnectorInfo {
	config := d.Get("credentials").([]interface{})[0].(map[string]interface{})
	authType := config["auth_type"].(string)

	connector1 := &nextgen.ConnectorInfo{
		Type_: nextgen.ConnectorTypes.JDBC,
		JDBC: &nextgen.JdbcConnector{
			Url: d.Get("url").(string),
			// As currently we support through delegate only
			ExecuteOnDelegate:    true,
			IgnoreTestConnection: false,
		},
	}

	switch authType {
	case nextgen.JDBCAuthTypes.UsernamePassword.String():
		{
			connector1.JDBC.Auth = &nextgen.JdbcAuthenticationDto{
				Type_:            nextgen.JDBCAuthTypes.UsernamePassword,
				UsernamePassword: &nextgen.JdbcUserNamePasswordDto{},
			}
			if usernamePasswordConfig, ok := config["username_password"]; ok && len(usernamePasswordConfig.([]interface{})) > 0 {
				config = usernamePasswordConfig.([]interface{})[0].(map[string]interface{})
				if attr, ok := config["username"]; ok {
					connector1.JDBC.Auth.UsernamePassword.Username = attr.(string)
				}
				if attr, ok := config["username_ref"]; ok {
					connector1.JDBC.Auth.UsernamePassword.UsernameRef = attr.(string)
				}
				if attr, ok := config["password_ref"]; ok {
					connector1.JDBC.Auth.UsernamePassword.PasswordRef = attr.(string)
				}
			} else {
				if attr, ok := config["username"]; ok {
					connector1.JDBC.Auth.UsernamePassword.Username = attr.(string)
				}

				if attr, ok := config["username_ref"]; ok {
					connector1.JDBC.Auth.UsernamePassword.UsernameRef = attr.(string)
				}

				if attr, ok := config["password_ref"]; ok {
					connector1.JDBC.Auth.UsernamePassword.PasswordRef = attr.(string)
				}
			}
		}
	case nextgen.JDBCAuthTypes.ServiceAccount.String():
		{
			connector1.JDBC.Auth = &nextgen.JdbcAuthenticationDto{
				Type_:          nextgen.JDBCAuthTypes.ServiceAccount,
				ServiceAccount: &nextgen.JdbcServiceAccountDto{},
			}
			if serviceAccountConfig, ok := config["service_account"]; ok && len(serviceAccountConfig.([]interface{})) > 0 {
				config = serviceAccountConfig.([]interface{})[0].(map[string]interface{})
				if attr, ok := config["token_ref"]; ok {
					connector1.JDBC.Auth.ServiceAccount.ServiceAccountTokenRef = attr.(string)
				}
			}
		}
	case nextgen.JDBCAuthTypes.KeyPair.String():
		{
			connector1.JDBC.Auth = &nextgen.JdbcAuthenticationDto{
				Type_:   nextgen.JDBCAuthTypes.KeyPair,
				KeyPair: &nextgen.JdbcKeyPairDto{},
			}
			if keyPairConfig, ok := config["key_pair"]; ok && len(keyPairConfig.([]interface{})) > 0 {
				config = keyPairConfig.([]interface{})[0].(map[string]interface{})
				if attr, ok := config["username"]; ok {
					connector1.JDBC.Auth.KeyPair.Username = attr.(string)
				}
				if attr, ok := config["username_ref"]; ok {
					connector1.JDBC.Auth.KeyPair.UsernameRef = attr.(string)
				}
				if attr, ok := config["private_key_file_ref"]; ok {
					connector1.JDBC.Auth.KeyPair.PrivateKeyFileRef = attr.(string)
				}
				if attr, ok := config["private_key_passphrase_ref"]; ok {
					connector1.JDBC.Auth.KeyPair.PrivateKeyPassphraseRef = attr.(string)
				}
			}
		}
	case nextgen.JDBCAuthTypes.Oidc.String():
		{
			connector1.JDBC.Auth = &nextgen.JdbcAuthenticationDto{
				Type_: nextgen.JDBCAuthTypes.Oidc,
				Oidc:  &nextgen.JdbcOidcDto{},
			}
			if oidcConfig, ok := config["oidc"]; ok && len(oidcConfig.([]interface{})) > 0 {
				oidcMap := oidcConfig.([]interface{})[0].(map[string]interface{})
				if attr, ok := oidcMap["provider_type"]; ok {
					connector1.JDBC.Auth.Oidc.ProviderType = nextgen.JDBCOidcProviderType(attr.(string))
				}
				if gcpOidcConfig, ok := oidcMap["gcp_oidc"]; ok && len(gcpOidcConfig.([]interface{})) > 0 {
					gcpMap := gcpOidcConfig.([]interface{})[0].(map[string]interface{})
					connector1.JDBC.Auth.Oidc.GcpOidcSpec = &nextgen.JdbcGcpOidcSpecDto{}
					if attr, ok := gcpMap["project_number"]; ok {
						connector1.JDBC.Auth.Oidc.GcpOidcSpec.ProjectNumber = attr.(string)
					}
					if attr, ok := gcpMap["workload_pool_id"]; ok {
						connector1.JDBC.Auth.Oidc.GcpOidcSpec.WorkloadPoolId = attr.(string)
					}
					if attr, ok := gcpMap["provider_id"]; ok {
						connector1.JDBC.Auth.Oidc.GcpOidcSpec.ProviderId = attr.(string)
					}
					if attr, ok := gcpMap["service_account_email"]; ok {
						connector1.JDBC.Auth.Oidc.GcpOidcSpec.ServiceAccountEmail = attr.(string)
					}
				}
			}
		}
	case nextgen.JDBCAuthTypes.InheritFromDelegate.String():
		{
			connector1.JDBC.Auth = &nextgen.JdbcAuthenticationDto{
				Type_:               nextgen.JDBCAuthTypes.InheritFromDelegate,
				InheritFromDelegate: &nextgen.JdbcDelegateAccessDto{},
			}
			if inheritConfig, ok := config["inherit_from_delegate"]; ok && len(inheritConfig.([]interface{})) > 0 {
				if inheritMap, ok := inheritConfig.([]interface{})[0].(map[string]interface{}); ok {
					if attr, ok := inheritMap["username"]; ok {
						connector1.JDBC.Auth.InheritFromDelegate.Username = attr.(string)
					}
					if attr, ok := inheritMap["username_ref"]; ok {
						connector1.JDBC.Auth.InheritFromDelegate.UsernameRef = attr.(string)
					}
				}
			}
		}
	default:
		panic(fmt.Sprintf("unknown jdbc auth method type %s", authType))
	}

	if attr, ok := d.GetOk("delegate_selectors"); ok {
		connector1.JDBC.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr.(*schema.Set).List())
	}

	return connector1
}

func readConnectorJDBC(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	d.Set("url", connector.JDBC.Url)
	d.Set("delegate_selectors", connector.JDBC.DelegateSelectors)
	switch connector.JDBC.Auth.Type_ {
	case nextgen.JDBCAuthTypes.UsernamePassword:
		useFirstClassUserAndPasswordVars := false
		_, ok := d.GetOk("credentials")
		if ok {
			cred := d.Get("credentials").([]interface{})[0].(map[string]interface{})
			passwordRefValue, passwordRefKeyPresent := cred["password_ref"]
			if passwordRefKeyPresent && passwordRefValue != "" {
				useFirstClassUserAndPasswordVars = true
			}
		}
		if useFirstClassUserAndPasswordVars {
			d.Set("credentials", []map[string]interface{}{
				{
					"auth_type":    nextgen.JDBCAuthTypes.UsernamePassword.String(),
					"username":     connector.JDBC.Auth.UsernamePassword.Username,
					"username_ref": connector.JDBC.Auth.UsernamePassword.UsernameRef,
					"password_ref": connector.JDBC.Auth.UsernamePassword.PasswordRef,
				},
			})
		} else {
			d.Set("credentials", []map[string]interface{}{
				{
					"auth_type": nextgen.JDBCAuthTypes.UsernamePassword.String(),
					"username_password": []map[string]interface{}{
						{
							"username":     connector.JDBC.Auth.UsernamePassword.Username,
							"username_ref": connector.JDBC.Auth.UsernamePassword.UsernameRef,
							"password_ref": connector.JDBC.Auth.UsernamePassword.PasswordRef,
						},
					},
				},
			})
		}
	case nextgen.JDBCAuthTypes.ServiceAccount:
		d.Set("credentials", []map[string]interface{}{
			{
				"auth_type": nextgen.JDBCAuthTypes.ServiceAccount.String(),
				"service_account": []map[string]interface{}{
					{
						"token_ref": connector.JDBC.Auth.ServiceAccount.ServiceAccountTokenRef,
					},
				},
			},
		})
	case nextgen.JDBCAuthTypes.KeyPair:
		d.Set("credentials", []map[string]interface{}{
			{
				"auth_type": nextgen.JDBCAuthTypes.KeyPair.String(),
				"key_pair": []map[string]interface{}{
					{
						"username":                   connector.JDBC.Auth.KeyPair.Username,
						"username_ref":               connector.JDBC.Auth.KeyPair.UsernameRef,
						"private_key_file_ref":       connector.JDBC.Auth.KeyPair.PrivateKeyFileRef,
						"private_key_passphrase_ref": connector.JDBC.Auth.KeyPair.PrivateKeyPassphraseRef,
					},
				},
			},
		})
	case nextgen.JDBCAuthTypes.Oidc:
		credMap := map[string]interface{}{
			"auth_type": nextgen.JDBCAuthTypes.Oidc.String(),
			"oidc": []map[string]interface{}{
				{
					"provider_type": string(connector.JDBC.Auth.Oidc.ProviderType),
				},
			},
		}
		if connector.JDBC.Auth.Oidc.GcpOidcSpec != nil {
			credMap["oidc"] = []map[string]interface{}{
				{
					"provider_type": string(connector.JDBC.Auth.Oidc.ProviderType),
					"gcp_oidc": []map[string]interface{}{
						{
							"project_number":        connector.JDBC.Auth.Oidc.GcpOidcSpec.ProjectNumber,
							"workload_pool_id":      connector.JDBC.Auth.Oidc.GcpOidcSpec.WorkloadPoolId,
							"provider_id":           connector.JDBC.Auth.Oidc.GcpOidcSpec.ProviderId,
							"service_account_email": connector.JDBC.Auth.Oidc.GcpOidcSpec.ServiceAccountEmail,
						},
					},
				},
			}
		}
		d.Set("credentials", []map[string]interface{}{credMap})
	case nextgen.JDBCAuthTypes.InheritFromDelegate:
		inheritMap := map[string]interface{}{}
		if connector.JDBC.Auth.InheritFromDelegate != nil {
			inheritMap = map[string]interface{}{
				"username":     connector.JDBC.Auth.InheritFromDelegate.Username,
				"username_ref": connector.JDBC.Auth.InheritFromDelegate.UsernameRef,
			}
		}
		d.Set("credentials", []map[string]interface{}{
			{
				"auth_type": nextgen.JDBCAuthTypes.InheritFromDelegate.String(),
				"inherit_from_delegate": []map[string]interface{}{
					inheritMap,
				},
			},
		})
	default:
		return fmt.Errorf("unknown jdbc auth method type %s", connector.JDBC.Auth.Type_)
	}

	return nil
}
