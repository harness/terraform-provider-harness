package provider

import (
	"context"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	k8sAuthTypes = []string{
		"authentication.0.delegate_selectors",
		"authentication.0.username_password",
		"authentication.0.oidc",
		"authentication.0.service_account",
		// "authentication.0.custom",
	}
)

func resourceCloudProviderK8s() *schema.Resource {

	providerSchema := map[string]*schema.Schema{
		"skip_validation": {
			Description: "Skip validation of Kubernetes configuration.",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
		},
		"authentication": {
			Description: "Authentication configuration for the Kubernetes cluster",
			Type:        schema.TypeList,
			MaxItems:    1,
			Required:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"delegate_selectors": {
						Description:  "Delegate selectors to inherit the GCP credentials from.",
						Type:         schema.TypeList,
						Optional:     true,
						Elem:         &schema.Schema{Type: schema.TypeString},
						ExactlyOneOf: k8sAuthTypes,
					},
					"username_password": {
						Description:  "Username and password for authentication to the cluster",
						Type:         schema.TypeList,
						Optional:     true,
						MaxItems:     1,
						ExactlyOneOf: k8sAuthTypes,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"username": {
									Description:   "Username for authentication to the cluster",
									Type:          schema.TypeString,
									Optional:      true,
									ConflictsWith: []string{"authentication.0.username_password.0.username_secret_name"},
								},
								"username_secret_name": {
									Description:   "Name of the Harness secret containing the username for authentication to the cluster",
									Type:          schema.TypeString,
									Optional:      true,
									ConflictsWith: []string{"authentication.0.username_password.0.username"},
								},
								"password_secret_name": {
									Description: "Name of the Harness secret containing the password for the cluster.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"master_url": {
									Description: "URL of the Kubernetes master to connect to.",
									Type:        schema.TypeString,
									Required:    true,
								},
							},
						},
					},
					"oidc": {
						Description:  "Service account configuration for connecting to the Kubernetes cluster",
						Type:         schema.TypeList,
						Optional:     true,
						MaxItems:     1,
						ExactlyOneOf: k8sAuthTypes,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"identity_provider_url": {
									Description: "URL of the identity provider to use.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"username": {
									Description: "Username for authentication to the cluster. This can be the username itself or the ID of a harness secret.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"password_secret_name": {
									Description: "Name of the Harness secret containing the password for the cluster.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"client_id_secret_name": {
									Description: "Name of the Harness secret containing the client ID for the cluster.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"client_secret_secret_name": {
									Description: "Name of the Harness secret containing the client secret for the cluster.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"scopes": {
									Description: "Scopes to request from the identity provider.",
									Type:        schema.TypeList,
									Optional:    true,
									Elem:        &schema.Schema{Type: schema.TypeString},
								},
								"master_url": {
									Description: "URL of the Kubernetes master to connect to.",
									Type:        schema.TypeString,
									Required:    true,
								},
							},
						},
					},
					"service_account": {
						Description:  "Username and password for authentication to the cluster",
						Type:         schema.TypeList,
						Optional:     true,
						MaxItems:     1,
						ExactlyOneOf: k8sAuthTypes,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"service_account_token_secret_name": {
									Description: "Name of the Harness secret containing the service account token for the cluster.",
									Type:        schema.TypeString,
									Required:    true,
								},
								"ca_certificate_secret_name": {
									Description: "Name of the Harness secret containing the CA certificate for the cluster.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"master_url": {
									Description: "URL of the Kubernetes master to connect to.",
									Type:        schema.TypeString,
									Required:    true,
								},
							},
						},
					},
					// "custom": {
					// 	Description:  "Custom configuration for connecting to the Kubernetes cluster",
					// 	Type:         schema.TypeList,
					// 	Optional:     true,
					// 	MaxItems:     1,
					// 	ExactlyOneOf: k8sAuthTypes,
					// 	Elem: &schema.Resource{
					// 		Schema: map[string]*schema.Schema{
					// 			"username": {
					// 				Description:   "Username for authentication to the cluster",
					// 				Type:          schema.TypeString,
					// 				Optional:      true,
					// 				ConflictsWith: []string{"authentication.0.custom.0.username_secret_name"},
					// 				AtLeastOneOf:  []string{"authentication.0.custom.0.username", "authentication.0.custom.0.username_secret_name"},
					// 			},
					// 			"username_secret_name": {
					// 				Description:   "Name of the Harness secret containing the username for authentication to the cluster",
					// 				Type:          schema.TypeString,
					// 				Optional:      true,
					// 				ConflictsWith: []string{"authentication.0.custom.0.username"},
					// 				AtLeastOneOf:  []string{"authentication.0.custom.0.username", "authentication.0.custom.0.username_secret_name"},
					// 			},
					// 			"password_secret_name": {
					// 				Description: "Name of the Harness secret containing the password for the cluster.",
					// 				Type:        schema.TypeString,
					// 				Optional:    true,
					// 			},
					// 			"ca_certificate_secret_name": {
					// 				Description: "Name of the Harness secret containing the CA certificate for the cluster.",
					// 				Type:        schema.TypeString,
					// 				Optional:    true,
					// 			},
					// 			"client_certificate_secret_name": {
					// 				Description: "Name of the Harness secret containing the client certificate for the cluster.",
					// 				Type:        schema.TypeString,
					// 				Optional:    true,
					// 			},
					// 			"client_key_secret_name": {
					// 				Description: "Name of the Harness secret containing the client key for the cluster.",
					// 				Type:        schema.TypeString,
					// 				Optional:    true,
					// 			},
					// 			"client_key_passphrase_secret_name": {
					// 				Description: "Name of the Harness secret containing the client key passphrase for the cluster.",
					// 				Type:        schema.TypeString,
					// 				Optional:    true,
					// 			},
					// 			"client_key_algorithm": {
					// 				Description: "Algorithm for the client key.",
					// 				Type:        schema.TypeString,
					// 				Optional:    true,
					// 			},
					// 			"service_account_token_secret_name": {
					// 				Description: "Name of the Harness secret containing the service account token for the cluster.",
					// 				Type:        schema.TypeString,
					// 				Optional:    true,
					// 			},
					// 			"master_url": {
					// 				Description: "URL of the Kubernetes master to connect to.",
					// 				Type:        schema.TypeString,
					// 				Required:    true,
					// 			},
					// 		},
					// 	},
					// },
				},
			},
		},
	}

	helpers.MergeSchemas(commonCloudProviderSchema(), providerSchema)

	return &schema.Resource{
		Description:   "Resource for creating a Kubernetes cloud provider",
		CreateContext: resourceCloudProviderK8sCreateOrUpdate,
		ReadContext:   resourceCloudProviderK8sRead,
		UpdateContext: resourceCloudProviderK8sCreateOrUpdate,
		DeleteContext: resourceCloudProviderDelete,

		Schema: providerSchema,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceCloudProviderK8sRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	cp := &cac.KubernetesCloudProvider{}
	if err := c.ConfigAsCode().GetCloudProviderById(d.Id(), cp); err != nil {
		return diag.FromErr(err)
	} else if cp.IsEmpty() {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readCloudProviderK8s(c, d, cp)
}

func readCloudProviderK8s(c *api.Client, d *schema.ResourceData, cp *cac.KubernetesCloudProvider) diag.Diagnostics {
	d.SetId(cp.Id)
	d.Set("name", cp.Name)
	d.Set("skip_validation", cp.SkipValidation)
	d.Set("authentication", flattenK8sAuth(d, cp))

	scope, err := flattenUsageRestrictions(c, cp.UsageRestrictions)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("usage_scope", scope)

	return nil
}

func resourceCloudProviderK8sCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	var input *cac.KubernetesCloudProvider
	var err error

	if d.IsNewResource() {
		input = cac.NewEntity(cac.ObjectTypes.KubernetesCloudProvider).(*cac.KubernetesCloudProvider)
	} else {
		input = &cac.KubernetesCloudProvider{}
		if err = c.ConfigAsCode().GetCloudProviderById(d.Id(), input); err != nil {
			return diag.FromErr(err)
		} else if input.IsEmpty() {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	}

	input.Name = d.Get("name").(string)
	input.SkipValidation = d.Get("skip_validation").(bool)

	expandK8sAuth(d.Get("authentication").([]interface{}), input)

	if err := expandUsageRestrictions(c, d.Get("usage_scope").(*schema.Set).List(), input.UsageRestrictions); err != nil {
		return diag.FromErr(err)
	}

	cp, err := c.ConfigAsCode().UpsertKubernetesCloudProvider(input)

	if err != nil {
		return diag.FromErr(err)
	}

	return readCloudProviderK8s(c, d, cp)
}

func flattenK8sAuth(d *schema.ResourceData, cp *cac.KubernetesCloudProvider) []map[string]interface{} {

	authConfig := make([]map[string]interface{}, 0)

	r := make(map[string]interface{})

	switch cp.AuthType {
	case cac.KubernetesAuthTypes.Custom:
		r["custom"] = flattenK8sCustomAuth(d, cp)
	case cac.KubernetesAuthTypes.UsernameAndPassword:
		r["username_password"] = flattenK8sUsernameAndPasswordAuth(d, cp)
	case cac.KubernetesAuthTypes.OIDC:
		r["oidc"] = flattenK8sOIDCAuth(d, cp)
	case cac.KubernetesAuthTypes.ServiceAccount:
		r["service_account"] = flattenK8sServiceAccountAuth(d, cp)
	default:
		// Inherit from delegate
		r["delegate_selectors"] = cp.DelegateSelectors
	}

	return append(authConfig, r)
}

func flattenK8sServiceAccountAuth(d *schema.ResourceData, cp *cac.KubernetesCloudProvider) []map[string]interface{} {
	results := make([]map[string]interface{}, 0)

	r := make(map[string]interface{})

	if cp.MasterUrl != "" {
		r["master_url"] = cp.MasterUrl
	}

	if cp.CACert != nil {
		r["ca_certificate_secret_name"] = cp.CACert.Name
	}

	if cp.ServiceAccountToken != nil {
		r["service_account_token_secret_name"] = cp.ServiceAccountToken.Name
	}

	return append(results, r)
}

func flattenK8sOIDCAuth(d *schema.ResourceData, cp *cac.KubernetesCloudProvider) []map[string]interface{} {
	results := make([]map[string]interface{}, 0)

	r := make(map[string]interface{})

	if cp.OIDCIdentityProviderUrl != "" {
		r["identity_provider_url"] = cp.OIDCIdentityProviderUrl
	}

	if cp.OIDCUsername != "" {
		r["username"] = cp.OIDCUsername
	}

	if cp.OIDCPassword != nil {
		r["password_secret_name"] = cp.OIDCPassword.Name
	}

	if cp.OIDCClientId != nil {
		r["client_id_secret_name"] = cp.OIDCClientId.Name
	}

	if cp.OIDCSecret != nil {
		r["client_secret_secret_name"] = cp.OIDCSecret.Name
	}

	if cp.OIDCScopes != "" {
		r["scopes"] = strings.Split(cp.OIDCScopes, " ")
	}

	if cp.MasterUrl != "" {
		r["master_url"] = cp.MasterUrl
	}

	return append(results, r)
}

func flattenK8sUsernameAndPasswordAuth(d *schema.ResourceData, cp *cac.KubernetesCloudProvider) []map[string]interface{} {
	results := make([]map[string]interface{}, 0)

	r := map[string]interface{}{}

	if cp.Username != "" {
		r["username"] = cp.Username
	}

	if cp.UsernameSecretId != nil {
		r["username_secret_name"] = cp.UsernameSecretId.Name
	}

	if cp.Password != nil {
		r["password_secret_name"] = cp.Password.Name
	}

	r["master_url"] = cp.MasterUrl

	return append(results, r)
}

func flattenK8sCustomAuth(d *schema.ResourceData, cp *cac.KubernetesCloudProvider) []map[string]interface{} {
	results := make([]map[string]interface{}, 0)

	r := make(map[string]interface{})

	if cp.Username != "" {
		r["username"] = cp.Username
	}

	if cp.UsernameSecretId != nil {
		r["username_secret_name"] = cp.UsernameSecretId.Name
	}

	if cp.Password != nil {
		r["password_secret_name"] = cp.Password.Name
	}

	if cp.MasterUrl != "" {
		r["master_url"] = cp.MasterUrl
	}

	if cp.CACert != nil {
		r["ca_certificate_secret_name"] = cp.CACert.Name
	}

	if cp.CACert != nil {
		r["ca_certificate_secret_name"] = cp.CACert.Name
	}

	if cp.ClientKey != nil {
		r["client_key_secret_name"] = cp.ClientKey.Name
	}

	if cp.ClientKey != nil {
		r["client_key_secret_name"] = cp.ClientKey.Name
	}

	if cp.ClientKeyPassPhrase != nil {
		r["client_key_passphrase_secret_name"] = cp.ClientKeyPassPhrase.Name
	}

	if cp.ClientKeyAlgorithm != "" {
		r["client_key_algorithm"] = cp.ClientKeyAlgorithm
	}

	if cp.ServiceAccountToken != nil {
		r["service_account_token_secret_name"] = cp.ServiceAccountToken.Name
	}

	return append(results, r)
}

func expandK8sAuth(d []interface{}, cp *cac.KubernetesCloudProvider) {

	if len(d) == 0 {
		return
	}

	auth := d[0].(map[string]interface{})

	if attr, ok := auth["delegate_selectors"]; ok {
		for _, v := range attr.([]interface{}) {
			cp.DelegateSelectors = append(cp.DelegateSelectors, v.(string))
			cp.UseKubernetesDelegate = true
		}
	}

	if attr, ok := auth["username_password"]; ok {
		for _, v := range attr.([]interface{}) {
			data := v.(map[string]interface{})

			cp.AuthType = cac.KubernetesAuthTypes.UsernameAndPassword

			if attr, ok := data["master_url"]; ok && attr != "" {
				cp.MasterUrl = attr.(string)
			}

			if attr, ok := data["username"]; ok && attr != "" {
				cp.Username = attr.(string)
			}

			if attr, ok := data["username_secret_name"]; ok && attr != "" {
				cp.UsernameSecretId = &cac.SecretRef{
					Name: attr.(string),
				}
				cp.UseEncryptedUsername = true
			}

			if attr, ok := data["password_secret_name"]; ok && attr != "" {
				cp.Password = &cac.SecretRef{
					Name: attr.(string),
				}
			}
		}
	}

	if attr, ok := auth["oidc"]; ok {
		for _, v := range attr.([]interface{}) {
			data := v.(map[string]interface{})

			cp.AuthType = cac.KubernetesAuthTypes.OIDC

			if attr, ok := data["master_url"]; ok && attr != "" {
				cp.MasterUrl = attr.(string)
			}

			if attr, ok := data["username"]; ok && attr != "" {
				cp.OIDCPassword = &cac.SecretRef{
					Name: attr.(string),
				}
			}

			if attr, ok := data["identity_provider_url"]; ok && attr != "" {
				cp.OIDCIdentityProviderUrl = attr.(string)
			}

			if attr, ok := data["username"]; ok && attr != "" {
				cp.OIDCUsername = attr.(string)
			}

			if attr, ok := data["password_secret_name"]; ok && attr != "" {
				cp.OIDCPassword = &cac.SecretRef{
					Name: attr.(string),
				}
			}

			if attr, ok := data["client_id_secret_name"]; ok && attr != "" {
				cp.OIDCClientId = &cac.SecretRef{
					Name: attr.(string),
				}
			}

			if attr, ok := data["client_secret_secret_name"]; ok && attr != "" {
				cp.OIDCSecret = &cac.SecretRef{
					Name: attr.(string),
				}
			}

			if attr, ok := data["scopes"]; ok && attr != nil {
				scopes := []string{}
				for _, v := range attr.([]interface{}) {
					scopes = append(scopes, v.(string))
				}
				cp.OIDCScopes = strings.Join(scopes, " ")
			}
		}
	}

	if attr, ok := auth["service_account"]; ok {
		for _, v := range attr.([]interface{}) {
			data := v.(map[string]interface{})

			cp.AuthType = cac.KubernetesAuthTypes.ServiceAccount

			if attr, ok := data["master_url"]; ok && attr != "" {
				cp.MasterUrl = attr.(string)
			}

			if attr, ok := data["service_account_token_secret_name"]; ok && attr != "" {
				cp.ServiceAccountToken = &cac.SecretRef{
					Name: attr.(string),
				}
			}

			if attr, ok := data["ca_certificate_secret_name"]; ok && attr != "" {
				cp.CACert = &cac.SecretRef{
					Name: attr.(string),
				}
			}
		}
	}

	if attr, ok := auth["custom"]; ok {
		for _, v := range attr.([]interface{}) {
			data := v.(map[string]interface{})

			cp.AuthType = cac.KubernetesAuthTypes.Custom

			if attr, ok := data["master_url"]; ok && attr != "" {
				cp.MasterUrl = attr.(string)
			}

			if attr, ok := data["username"]; ok && attr != "" {
				cp.Username = attr.(string)
			}

			if attr, ok := data["username_secret_name"]; ok && attr != "" {
				cp.UsernameSecretId = &cac.SecretRef{
					Name: attr.(string),
				}
			}

			if attr, ok := data["password_secret_name"]; ok && attr != "" {
				cp.Password = &cac.SecretRef{
					Name: attr.(string),
				}
			}

			if attr, ok := data["ca_certificate_secret_name"]; ok && attr != "" {
				cp.CACert = &cac.SecretRef{
					Name: attr.(string),
				}
			}

			if attr, ok := data["client_certificate_secret_name"]; ok && attr != "" {
				cp.ClientCert = &cac.SecretRef{
					Name: attr.(string),
				}
			}

			if attr, ok := data["client_key_secret_name"]; ok && attr != "" {
				cp.ClientKey = &cac.SecretRef{
					Name: attr.(string),
				}
			}

			if attr, ok := data["client_key_passphrase_secret_name"]; ok && attr != "" {
				cp.ClientKeyPassPhrase = &cac.SecretRef{
					Name: attr.(string),
				}
			}

			if attr, ok := data["client_key_algorithm"]; ok && attr != nil {
				cp.ClientKeyAlgorithm = attr.(string)
			}

			if attr, ok := data["service_account_token_secret_name"]; ok && attr != nil {
				cp.ServiceAccountToken = &cac.SecretRef{
					Name: attr.(string),
				}
			}
		}
	}

}
