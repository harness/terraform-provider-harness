package connector

import (
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getK8sClusterSchema() *schema.Schema {
	return &schema.Schema{
		Description:   "The k8s cluster to use for the connector.",
		Type:          schema.TypeList,
		Optional:      true,
		MaxItems:      1,
		ConflictsWith: utils.GetConflictsWithSlice(connectorConfigNames, "k8s_cluster"),
		ExactlyOneOf:  connectorConfigNames,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"inherit_from_delegate": {
					Description: "Credentials are inherited from the delegate.",
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					AtLeastOneOf: []string{
						"k8s_cluster.0.inherit_from_delegate",
						"k8s_cluster.0.username_password",
						"k8s_cluster.0.service_account",
						"k8s_cluster.0.openid_connect",
						"k8s_cluster.0.client_key_cert",
					},
					ConflictsWith: []string{
						"k8s_cluster.0.username_password",
						"k8s_cluster.0.service_account",
						"k8s_cluster.0.openid_connect",
						"k8s_cluster.0.client_key_cert",
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
				"username_password": {
					Description: "Username and password for the connector.",
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					AtLeastOneOf: []string{
						"k8s_cluster.0.inherit_from_delegate",
						"k8s_cluster.0.username_password",
						"k8s_cluster.0.service_account",
						"k8s_cluster.0.openid_connect",
						"k8s_cluster.0.client_key_cert",
					},
					ConflictsWith: []string{
						"k8s_cluster.0.inherit_from_delegate",
						"k8s_cluster.0.service_account",
						"k8s_cluster.0.openid_connect",
						"k8s_cluster.0.client_key_cert",
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
									"k8s_cluster.0.username_password.0.username",
									"k8s_cluster.0.username_password.0.username_ref",
								},
								ConflictsWith: []string{"k8s_cluster.0.username_password.0.username_ref"},
							},
							"username_ref": {
								Description: "Reference to the secret containing the username for the connector.",
								Type:        schema.TypeString,
								Optional:    true,
								ExactlyOneOf: []string{
									"k8s_cluster.0.username_password.0.username",
									"k8s_cluster.0.username_password.0.username_ref",
								},
								ConflictsWith: []string{"k8s_cluster.0.username_password.0.username"},
							},
							"password_ref": {
								Description: "Reference to the secret containing the password for the connector.",
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
						"k8s_cluster.0.inherit_from_delegate",
						"k8s_cluster.0.username_password",
						"k8s_cluster.0.service_account",
						"k8s_cluster.0.openid_connect",
						"k8s_cluster.0.client_key_cert",
					},
					ConflictsWith: []string{
						"k8s_cluster.0.inherit_from_delegate",
						"k8s_cluster.0.username_password",
						"k8s_cluster.0.openid_connect",
						"k8s_cluster.0.client_key_cert",
					},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"master_url": {
								Description: "The URL of the Kubernetes cluster.",
								Type:        schema.TypeString,
								Required:    true,
							},
							"service_account_token_ref": {
								Description: "Reference to the secret containing the service account token for the connector.",
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
						"k8s_cluster.0.inherit_from_delegate",
						"k8s_cluster.0.username_password",
						"k8s_cluster.0.service_account",
						"k8s_cluster.0.openid_connect",
						"k8s_cluster.0.client_key_cert",
					},
					ConflictsWith: []string{
						"k8s_cluster.0.inherit_from_delegate",
						"k8s_cluster.0.username_password",
						"k8s_cluster.0.service_account",
						"k8s_cluster.0.client_key_cert",
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
									"k8s_cluster.0.openid_connect.0.username",
									"k8s_cluster.0.openid_connect.0.username_ref",
								},
								ConflictsWith: []string{"k8s_cluster.0.openid_connect.0.username_ref"},
							},
							"username_ref": {
								Description: "Reference to the secret containing the username for the connector.",
								Type:        schema.TypeString,
								Optional:    true,
								ExactlyOneOf: []string{
									"k8s_cluster.0.openid_connect.0.username",
									"k8s_cluster.0.openid_connect.0.username_ref",
								},
								ConflictsWith: []string{"k8s_cluster.0.openid_connect.0.username"},
							},
							"client_id_ref": {
								Description: "Reference to the secret containing the client ID for the connector.",
								Type:        schema.TypeString,
								Required:    true,
							},
							"password_ref": {
								Description: "Reference to the secret containing the password for the connector.",
								Type:        schema.TypeString,
								Required:    true,
							},
							"secret_ref": {
								Description: "Reference to the secret containing the client secret for the connector.",
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
						"k8s_cluster.0.inherit_from_delegate",
						"k8s_cluster.0.username_password",
						"k8s_cluster.0.service_account",
						"k8s_cluster.0.openid_connect",
						"k8s_cluster.0.client_key_cert",
					},
					ConflictsWith: []string{
						"k8s_cluster.0.inherit_from_delegate",
						"k8s_cluster.0.username_password",
						"k8s_cluster.0.service_account",
						"k8s_cluster.0.openid_connect",
					},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"master_url": {
								Description: "The URL of the Kubernetes cluster.",
								Type:        schema.TypeString,
								Required:    true,
							},
							"ca_cert_ref": {
								Description: "Reference to the secret containing the CA certificate for the connector.",
								Type:        schema.TypeString,
								Optional:    true,
							},
							"client_cert_ref": {
								Description: "Reference to the secret containing the client certificate for the connector.",
								Type:        schema.TypeString,
								Required:    true,
							},
							"client_key_ref": {
								Description: "Reference to the secret containing the client key for the connector.",
								Type:        schema.TypeString,
								Required:    true,
							},
							"client_key_passphrase_ref": {
								Description: "Reference to the secret containing the client key passphrase for the connector.",
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
		},
	}
}

func expandK8sCluster(d []interface{}, connector *nextgen.ConnectorInfo) {
	if len(d) == 0 {
		return
	}

	config := d[0].(map[string]interface{})
	connector.Type_ = nextgen.ConnectorTypes.K8sCluster
	connector.K8sCluster = &nextgen.KubernetesClusterConfig{}

	if attr := config["inherit_from_delegate"].([]interface{}); len(attr) > 0 {
		config := attr[0].(map[string]interface{})
		connector.K8sCluster.Credential = &nextgen.KubernetesCredential{
			Type_: nextgen.KubernetesCredentialTypes.InheritFromDelegate,
		}

		if attr := config["delegate_selectors"].(*schema.Set).List(); len(attr) > 0 {
			connector.K8sCluster.DelegateSelectors = utils.InterfaceSliceToStringSlice(attr)
		}

	} else {

		connector.K8sCluster.Credential = &nextgen.KubernetesCredential{
			Type_: nextgen.KubernetesCredentialTypes.ManualConfig,
			ManualConfig: &nextgen.KubernetesClusterDetails{
				Auth: &nextgen.KubernetesAuth{},
			},
		}

		if attr := config["client_key_cert"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
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

		if attr := config["username_password"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
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

		if attr := config["service_account"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
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

		if attr := config["service_account"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
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

		if attr := config["openid_connect"].([]interface{}); len(attr) > 0 {
			config := attr[0].(map[string]interface{})
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
}

func flattenK8sCluster(d *schema.ResourceData, connector *nextgen.ConnectorInfo) error {
	if connector.Type_ != nextgen.ConnectorTypes.K8sCluster {
		return nil
	}

	results := map[string]interface{}{}

	switch connector.K8sCluster.Credential.Type_ {
	case nextgen.KubernetesCredentialTypes.InheritFromDelegate:
		results["inherit_from_delegate"] = []map[string]interface{}{
			{
				"delegate_selectors": connector.K8sCluster.DelegateSelectors,
			},
		}
	case nextgen.KubernetesCredentialTypes.ManualConfig:
		auth := connector.K8sCluster.Credential.ManualConfig.Auth
		switch auth.Type_ {
		case nextgen.KubernetesAuthTypes.ClientKeyCert:
			results["client_key_cert"] = []map[string]interface{}{
				{
					"master_url":                connector.K8sCluster.Credential.ManualConfig.MasterUrl,
					"ca_cert_ref":               auth.ClientKeyCert.CaCertRef,
					"client_cert_ref":           auth.ClientKeyCert.ClientCertRef,
					"client_key_ref":            auth.ClientKeyCert.ClientKeyRef,
					"client_key_passphrase_ref": auth.ClientKeyCert.ClientKeyPassphraseRef,
					"client_key_algorithm":      auth.ClientKeyCert.ClientKeyAlgo,
				},
			}
		case nextgen.KubernetesAuthTypes.UsernamePassword:
			results["username_password"] = []map[string]interface{}{
				{
					"master_url":   connector.K8sCluster.Credential.ManualConfig.MasterUrl,
					"username":     auth.UsernamePassword.Username,
					"username_ref": auth.UsernamePassword.UsernameRef,
					"password_ref": auth.UsernamePassword.PasswordRef,
				},
			}
		case nextgen.KubernetesAuthTypes.ServiceAccount:
			results["service_account"] = []map[string]interface{}{
				{
					"master_url":                connector.K8sCluster.Credential.ManualConfig.MasterUrl,
					"service_account_token_ref": auth.ServiceAccount.ServiceAccountTokenRef,
				},
			}
		case nextgen.KubernetesAuthTypes.OpenIdConnect:
			results["openid_connect"] = []map[string]interface{}{
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
			}
		default:
			return fmt.Errorf("unsupported auth method: %s", auth.Type_)
		}
	default:
		return fmt.Errorf("unsupported k8s_cluster.credential.type_: %s", connector.K8sCluster.Credential.Type_)
	}

	d.Set("k8s_cluster", []interface{}{results})

	return nil
}
