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
)

func ResourceConnectorK8s() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a K8s connector.",
		ReadContext:   resourceConnectorK8sRead,
		CreateContext: resourceConnectorK8sCreateOrUpdate,
		UpdateContext: resourceConnectorK8sCreateOrUpdate,
		DeleteContext: resourceConnectorDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"inherit_from_delegate": {
				Description: "Credentials are inherited from the delegate.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				AtLeastOneOf: []string{
					"inherit_from_delegate",
					"username_password",
					"service_account",
					"openid_connect",
					"client_key_cert",
				},
				ConflictsWith: []string{
					"username_password",
					"service_account",
					"openid_connect",
					"client_key_cert",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delegate_selectors": {
							Description: "Selectors to use for the delegate.",
							Type:        schema.TypeSet,
							Required:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"delegate_selectors": {
				Description: "Selectors to use for the delegate.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				ConflictsWith: []string{
					"inherit_from_delegate",
				},
			},
			"username_password": {
				Description: "Username and password for the connector.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				AtLeastOneOf: []string{
					"inherit_from_delegate",
					"username_password",
					"service_account",
					"openid_connect",
					"client_key_cert",
				},
				ConflictsWith: []string{
					"inherit_from_delegate",
					"service_account",
					"openid_connect",
					"client_key_cert",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_url": {
							Description: "The URL of the Kubernetes cluster.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"username": {
							Description: "Username for the connector.",
							Type:        schema.TypeString,
							Optional:    true,
							ExactlyOneOf: []string{
								"username_password.0.username",
								"username_password.0.username_ref",
							},
							ConflictsWith: []string{"username_password.0.username_ref"},
						},
						"username_ref": {
							Description: "Reference to the secret containing the username for the connector." + secret_ref_text,
							Type:        schema.TypeString,
							Optional:    true,
							ExactlyOneOf: []string{
								"username_password.0.username",
								"username_password.0.username_ref",
							},
							ConflictsWith: []string{"username_password.0.username"},
						},
						"password_ref": {
							Description: "Reference to the secret containing the password for the connector." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"service_account": {
				Description: "Service account for the connector.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				AtLeastOneOf: []string{
					"inherit_from_delegate",
					"username_password",
					"service_account",
					"openid_connect",
					"client_key_cert",
				},
				ConflictsWith: []string{
					"inherit_from_delegate",
					"username_password",
					"openid_connect",
					"client_key_cert",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_url": {
							Description: "The URL of the Kubernetes cluster.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"service_account_token_ref": {
							Description: "Reference to the secret containing the service account token for the connector." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
			"openid_connect": {
				Description: "OpenID configuration for the connector.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				AtLeastOneOf: []string{
					"inherit_from_delegate",
					"username_password",
					"service_account",
					"openid_connect",
					"client_key_cert",
				},
				ConflictsWith: []string{
					"inherit_from_delegate",
					"username_password",
					"service_account",
					"client_key_cert",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_url": {
							Description: "The URL of the Kubernetes cluster.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"issuer_url": {
							Description: "The URL of the OpenID Connect issuer.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"username": {
							Description: "Username for the connector.",
							Type:        schema.TypeString,
							Optional:    true,
							ExactlyOneOf: []string{
								"openid_connect.0.username",
								"openid_connect.0.username_ref",
							},
							ConflictsWith: []string{"openid_connect.0.username_ref"},
						},
						"username_ref": {
							Description: "Reference to the secret containing the username for the connector." + secret_ref_text,
							Type:        schema.TypeString,
							Optional:    true,
							ExactlyOneOf: []string{
								"openid_connect.0.username",
								"openid_connect.0.username_ref",
							},
							ConflictsWith: []string{"openid_connect.0.username"},
						},
						"client_id_ref": {
							Description: "Reference to the secret containing the client ID for the connector." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
						"password_ref": {
							Description: "Reference to the secret containing the password for the connector." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
						"secret_ref": {
							Description: "Reference to the secret containing the client secret for the connector." + secret_ref_text,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"scopes": {
							Description: "Scopes to request for the connector.",
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"client_key_cert": {
				Description: "Client key and certificate config for the connector.",
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				AtLeastOneOf: []string{
					"inherit_from_delegate",
					"username_password",
					"service_account",
					"openid_connect",
					"client_key_cert",
				},
				ConflictsWith: []string{
					"inherit_from_delegate",
					"username_password",
					"service_account",
					"openid_connect",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master_url": {
							Description: "The URL of the Kubernetes cluster.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"ca_cert_ref": {
							Description: "Reference to the secret containing the CA certificate for the connector." + secret_ref_text,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"client_cert_ref": {
							Description: "Reference to the secret containing the client certificate for the connector." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
						"client_key_ref": {
							Description: "Reference to the secret containing the client key for the connector." + secret_ref_text,
							Type:        schema.TypeString,
							Required:    true,
						},
						"client_key_passphrase_ref": {
							Description: "Reference to the secret containing the client key passphrase for the connector." + secret_ref_text,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"client_key_algorithm": {
							Description: fmt.Sprintf("The algorithm used to generate the client key for the connector. Valid values are %s", strings.Join(nextgen.ClientKeyAlgorithmsSlice, ", ")),
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}

	helpers.SetMultiLevelResourceSchema(resource.Schema)

	return resource
}

func resourceConnectorK8sRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn, err := resourceConnectorReadBase(ctx, d, meta, nextgen.ConnectorTypes.K8sCluster)
	if err != nil {
		return err
	}

	if conn == nil {
		return nil
	}

	if err := readConnectorK8s(d, conn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceConnectorK8sCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := buildConnectorK8s(d)

	newConn, err := resourceConnectorCreateOrUpdateBase(ctx, d, meta, conn)
	if err != nil {
		return err
	}

	if err := readConnectorK8s(d, newConn); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildConnectorK8s(d *schema.ResourceData) *nextgen.ConnectorInfo {
	connector := &nextgen.ConnectorInfo{
		Type_:      nextgen.ConnectorTypes.K8sCluster,
		K8sCluster: &nextgen.KubernetesClusterConfig{},
	}

	if attr, ok := d.GetOk("inherit_from_delegate"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		connector.K8sCluster.Credential = &nextgen.KubernetesCredential{
			Type_: nextgen.KubernetesCredentialTypes.InheritFromDelegate,
		}

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.K8sCluster.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}

	} else {

		if attr, ok := d.GetOk("delegate_selectors"); ok {
			delegate_selectors := attr.(*schema.Set).List()
			if len(delegate_selectors) > 0 {
				connector.K8sCluster.DelegateSelectors = utils.InterfaceSliceToStringSlice(delegate_selectors)
			}
		}

		connector.K8sCluster.Credential = &nextgen.KubernetesCredential{
			Type_: nextgen.KubernetesCredentialTypes.ManualConfig,
			ManualConfig: &nextgen.KubernetesClusterDetails{
				Auth: &nextgen.KubernetesAuth{},
			},
		}

		if attr, ok := d.GetOk("client_key_cert"); ok {
			config := attr.([]interface{})[0].(map[string]interface{})
			clientKeyCert := &nextgen.KubernetesClientKeyCert{}
			connector.K8sCluster.Credential.ManualConfig.Auth.Type_ = nextgen.KubernetesAuthTypes.ClientKeyCert
			connector.K8sCluster.Credential.ManualConfig.Auth.ClientKeyCert = clientKeyCert

			if attr := config["master_url"].(string); attr != "" {
				connector.K8sCluster.Credential.ManualConfig.MasterUrl = attr
			}

			if attr := config["ca_cert_ref"].(string); attr != "" {
				clientKeyCert.CaCertRef = attr
			}

			if attr := config["client_cert_ref"].(string); attr != "" {
				clientKeyCert.ClientCertRef = attr
			}

			if attr := config["client_key_ref"].(string); attr != "" {
				clientKeyCert.ClientKeyRef = attr
			}

			if attr := config["client_key_passphrase_ref"].(string); attr != "" {
				clientKeyCert.ClientKeyPassphraseRef = attr
			}

			if attr := config["client_key_algorithm"].(string); attr != "" {
				clientKeyCert.ClientKeyAlgo = attr
			}
		}

		if attr, ok := d.GetOk("username_password"); ok {
			config := attr.([]interface{})[0].(map[string]interface{})
			usernamePasswordConfig := &nextgen.KubernetesUserNamePassword{}
			connector.K8sCluster.Credential.ManualConfig.Auth.Type_ = nextgen.KubernetesAuthTypes.UsernamePassword
			connector.K8sCluster.Credential.ManualConfig.Auth.UsernamePassword = usernamePasswordConfig

			if attr := config["master_url"].(string); attr != "" {
				connector.K8sCluster.Credential.ManualConfig.MasterUrl = attr
			}

			if attr := config["username"].(string); attr != "" {
				usernamePasswordConfig.Username = attr
			}

			if attr := config["username_ref"].(string); attr != "" {
				usernamePasswordConfig.UsernameRef = attr
			}

			if attr := config["password_ref"].(string); attr != "" {
				usernamePasswordConfig.PasswordRef = attr
			}
		}

		if attr, ok := d.GetOk("service_account"); ok {
			config := attr.([]interface{})[0].(map[string]interface{})
			saConfig := &nextgen.KubernetesServiceAccount{}
			connector.K8sCluster.Credential.ManualConfig.Auth.Type_ = nextgen.KubernetesAuthTypes.ServiceAccount
			connector.K8sCluster.Credential.ManualConfig.Auth.ServiceAccount = saConfig

			if attr := config["master_url"].(string); attr != "" {
				connector.K8sCluster.Credential.ManualConfig.MasterUrl = attr
			}

			if attr := config["service_account_token_ref"].(string); attr != "" {
				saConfig.ServiceAccountTokenRef = attr
			}
		}

		if attr, ok := d.GetOk("service_account"); ok {
			config := attr.([]interface{})[0].(map[string]interface{})
			saConfig := &nextgen.KubernetesServiceAccount{}
			connector.K8sCluster.Credential.ManualConfig.Auth.Type_ = nextgen.KubernetesAuthTypes.ServiceAccount
			connector.K8sCluster.Credential.ManualConfig.Auth.ServiceAccount = saConfig

			if attr := config["master_url"].(string); attr != "" {
				connector.K8sCluster.Credential.ManualConfig.MasterUrl = attr
			}

			if attr := config["service_account_token_ref"].(string); attr != "" {
				saConfig.ServiceAccountTokenRef = attr
			}
		}

		if attr, ok := d.GetOk("openid_connect"); ok {
			config := attr.([]interface{})[0].(map[string]interface{})
			oidcConfig := &nextgen.KubernetesOpenIdConnect{}
			connector.K8sCluster.Credential.ManualConfig.Auth.Type_ = nextgen.KubernetesAuthTypes.OpenIdConnect
			connector.K8sCluster.Credential.ManualConfig.Auth.OpenIdConnect = oidcConfig

			if attr := config["master_url"].(string); attr != "" {
				connector.K8sCluster.Credential.ManualConfig.MasterUrl = attr
			}

			if attr := config["issuer_url"].(string); attr != "" {
				oidcConfig.OidcIssuerUrl = attr
			}

			if attr := config["username"].(string); attr != "" {
				oidcConfig.OidcUsername = attr
			}

			if attr := config["username_ref"].(string); attr != "" {
				oidcConfig.OidcUsernameRef = attr
			}

			if attr := config["client_id_ref"].(string); attr != "" {
				oidcConfig.OidcClientIdRef = attr
			}

			if attr := config["password_ref"].(string); attr != "" {
				oidcConfig.OidcPasswordRef = attr
			}

			if attr := config["secret_ref"].(string); attr != "" {
				oidcConfig.OidcSecretRef = attr
			}

			if attr := config["scopes"].([]interface{}); len(attr) > 0 {
				oidcConfig.OidcScopes = strings.Join(utils.InterfaceSliceToStringSlice(attr), ",")
			}
		}
	}

	return connector
}

func readConnectorK8s(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {

	switch connector.K8sCluster.Credential.Type_ {
	case nextgen.KubernetesCredentialTypes.InheritFromDelegate:
		d.Set("inherit_from_delegate", []map[string]interface{}{
			{
				"delegate_selectors": connector.K8sCluster.DelegateSelectors,
			},
		})
	case nextgen.KubernetesCredentialTypes.ManualConfig:
		auth := connector.K8sCluster.Credential.ManualConfig.Auth
		switch auth.Type_ {
		case nextgen.KubernetesAuthTypes.ClientKeyCert:
			d.Set("client_key_cert", []map[string]interface{}{
				{
					"master_url":                connector.K8sCluster.Credential.ManualConfig.MasterUrl,
					"ca_cert_ref":               auth.ClientKeyCert.CaCertRef,
					"client_cert_ref":           auth.ClientKeyCert.ClientCertRef,
					"client_key_ref":            auth.ClientKeyCert.ClientKeyRef,
					"client_key_passphrase_ref": auth.ClientKeyCert.ClientKeyPassphraseRef,
					"client_key_algorithm":      auth.ClientKeyCert.ClientKeyAlgo,
				},
			})
			d.Set("delegate_selectors", connector.K8sCluster.DelegateSelectors)
		case nextgen.KubernetesAuthTypes.UsernamePassword:
			d.Set("username_password", []map[string]interface{}{
				{
					"master_url":   connector.K8sCluster.Credential.ManualConfig.MasterUrl,
					"username":     auth.UsernamePassword.Username,
					"username_ref": auth.UsernamePassword.UsernameRef,
					"password_ref": auth.UsernamePassword.PasswordRef,
				},
			})
			d.Set("delegate_selectors", connector.K8sCluster.DelegateSelectors)
		case nextgen.KubernetesAuthTypes.ServiceAccount:
			d.Set("service_account", []map[string]interface{}{
				{
					"master_url":                connector.K8sCluster.Credential.ManualConfig.MasterUrl,
					"service_account_token_ref": auth.ServiceAccount.ServiceAccountTokenRef,
				},
			})
			d.Set("delegate_selectors", connector.K8sCluster.DelegateSelectors)
		case nextgen.KubernetesAuthTypes.OpenIdConnect:
			d.Set("openid_connect", []map[string]interface{}{
				{
					"master_url":    connector.K8sCluster.Credential.ManualConfig.MasterUrl,
					"issuer_url":    auth.OpenIdConnect.OidcIssuerUrl,
					"username":      auth.OpenIdConnect.OidcUsername,
					"username_ref":  auth.OpenIdConnect.OidcUsernameRef,
					"client_id_ref": auth.OpenIdConnect.OidcClientIdRef,
					"password_ref":  auth.OpenIdConnect.OidcPasswordRef,
					"secret_ref":    auth.OpenIdConnect.OidcSecretRef,
					"scopes":        strings.Split(auth.OpenIdConnect.OidcScopes, ","),
				},
			})
			d.Set("delegate_selectors", connector.K8sCluster.DelegateSelectors)
		default:
			return fmt.Errorf("unsupported auth method: %s", auth.Type_)
		}
	default:
		return fmt.Errorf("unsupported k8s_cluster.credential.type_: %s", connector.K8sCluster.Credential.Type_)
	}

	return nil
}
