package secret

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceSecretWinRM() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for creating a WinRM credential secret.",
		ReadContext:   resourceSecretWinRMRead,
		CreateContext: resourceSecretWinRMCreateOrUpdate,
		UpdateContext: resourceSecretWinRMCreateOrUpdate,
		DeleteContext: resourceSecretDelete,
		Importer:      helpers.MultiLevelResourceImporter,

		Schema: map[string]*schema.Schema{
			"port": {
				Description: "WinRM port. Default is 5986 for HTTPS, 5985 for HTTP.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"ntlm": {
				Description:  "NTLM authentication scheme",
				Type:         schema.TypeList,
				MaxItems:     1,
				Optional:     true,
				ExactlyOneOf: []string{"ntlm", "kerberos"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Description: "Domain name for NTLM authentication.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"username": {
							Description: "Username to use for authentication.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"password_ref": {
							Description: "Reference to a secret containing the password to use for authentication. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account' to the expression: account.{identifier}.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"use_ssl": {
							Description: "Use SSL/TLS for WinRM communication.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"skip_cert_check": {
							Description: "Skip certificate verification.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
						},
						"use_no_profile": {
							Description: "Use no profile.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
					},
				},
			},
			"kerberos": {
				Description:  "Kerberos authentication scheme",
				Type:         schema.TypeList,
				MaxItems:     1,
				Optional:     true,
				ExactlyOneOf: []string{"ntlm", "kerberos"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"principal": {
							Description: "Kerberos principal.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"realm": {
							Description: "Kerberos realm.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"tgt_generation_method": {
							Description: "Method to generate TGT (Ticket Granting Ticket).",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"use_ssl": {
							Description: "Use SSL/TLS for WinRM communication.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"skip_cert_check": {
							Description: "Skip certificate verification.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
						},
						"use_no_profile": {
							Description: "Use no profile.",
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
						},
						"tgt_key_tab_file_path_spec": {
							Description:  "TGT generation using key tab file.",
							Type:         schema.TypeList,
							MaxItems:     1,
							Optional:     true,
							ExactlyOneOf: []string{"kerberos.0.tgt_key_tab_file_path_spec", "kerberos.0.tgt_password_spec"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_path": {
										Description: "Path to the key tab file.",
										Type:        schema.TypeString,
										Required:    true,
									},
								},
							},
						},
						"tgt_password_spec": {
							Description:  "TGT generation using password.",
							Type:         schema.TypeList,
							MaxItems:     1,
							Optional:     true,
							ExactlyOneOf: []string{"kerberos.0.tgt_key_tab_file_path_spec", "kerberos.0.tgt_password_spec"},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"password_ref": {
										Description: "Reference to a secret containing the password. To reference a secret at the organization scope, prefix 'org' to the expression: org.{identifier}. To reference a secret at the account scope, prefix 'account' to the expression: account.{identifier}.",
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

func resourceSecretWinRMRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secret, diags := resourceSecretReadBase(ctx, d, meta, nextgen.SecretTypes.WinRmCredentials)
	if diags.HasError() || secret == nil {
		return diags
	}

	if err := readSecretWinRM(d, secret); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceSecretWinRMCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secret := buildSecretWinRM(d)
	secret, diags := resourceSecretCreateOrUpdateBase(ctx, d, meta, secret)
	if diags.HasError() {
		return diags
	}

	if err := readSecretWinRM(d, secret); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func buildSecretWinRM(d *schema.ResourceData) *nextgen.Secret {
	secret := &nextgen.Secret{
		Type_: nextgen.SecretTypes.WinRmCredentials,
		WinRmCredentials: &nextgen.WinRmCredentialsSpec{
			Auth: &nextgen.WinRmAuth{},
		},
	}

	// Set port if specified
	if attr, ok := d.GetOk("port"); ok {
		secret.WinRmCredentials.Port = int32(attr.(int))
	}

	// Handle NTLM authentication
	if attr, ok := d.GetOk("ntlm"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		secret.WinRmCredentials.Auth.Type_ = "NTLM"
		secret.WinRmCredentials.Auth.NtlmConfig = &nextgen.NtlmConfig{
			Type_:          "NTLM",
			Domain:         config["domain"].(string),
			Username:       config["username"].(string),
			Password:       config["password_ref"].(string),
			UseSSL:         d.Get("ntlm.0.use_ssl").(bool),
			SkipCertChecks: d.Get("ntlm.0.skip_cert_check").(bool),
			UseNoProfile:   d.Get("ntlm.0.use_no_profile").(bool),
		}
	}

	// Handle Kerberos authentication
	if attr, ok := d.GetOk("kerberos"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})

		secret.WinRmCredentials.Auth.Type_ = "Kerberos"
		secret.WinRmCredentials.Auth.KerberosConfig = &nextgen.KerberosWinRmConfigDto{
			Type_:               "Kerberos",
			Principal:           config["principal"].(string),
			Realm:               config["realm"].(string),
			TgtGenerationMethod: config["tgt_generation_method"].(string),
			UseSSL:              d.Get("kerberos.0.use_ssl").(bool),
			SkipCertChecks:      d.Get("kerberos.0.skip_cert_check").(bool),
			UseNoProfile:        d.Get("kerberos.0.use_no_profile").(bool),
		}

		// Handle TGT key tab file path spec
		if _, ok := d.GetOk("kerberos.0.tgt_key_tab_file_path_spec"); ok {
			if val, ok := config["tgt_key_tab_file_path_spec"]; ok {
				tgtConfig := val.([]interface{})[0].(map[string]interface{})
				secret.WinRmCredentials.Auth.KerberosConfig.TgtGenerationMethod = string(nextgen.TgtGenerationMethodTypes.TGTKeyTabFilePathSpecDTO)
				secret.WinRmCredentials.Auth.KerberosConfig.KeyTabFilePathSpec = &nextgen.TgtKeyTabFilePathSpecDto{}

				if keyPath, ok := tgtConfig["key_path"]; ok {
					secret.WinRmCredentials.Auth.KerberosConfig.KeyTabFilePathSpec.KeyPath = keyPath.(string)
				}
			}
		}

		// Handle TGT password spec
		if _, ok := d.GetOk("kerberos.0.tgt_password_spec"); ok {
			if val, ok := config["tgt_password_spec"]; ok {
				tgtConfig := val.([]interface{})[0].(map[string]interface{})
				secret.WinRmCredentials.Auth.KerberosConfig.TgtGenerationMethod = string(nextgen.TgtGenerationMethodTypes.TGTPasswordSpecDTO)
				secret.WinRmCredentials.Auth.KerberosConfig.PasswordSpec = &nextgen.TgtPasswordSpecDto{}

				if password, ok := tgtConfig["password_ref"]; ok {
					secret.WinRmCredentials.Auth.KerberosConfig.PasswordSpec.Password = password.(string)
				}
			}
		}
	}

	return secret
}

func readSecretWinRM(d *schema.ResourceData, secret *nextgen.Secret) error {
	// Set port
	if secret.WinRmCredentials.Port > 0 {
		d.Set("port", secret.WinRmCredentials.Port)
	}

	// Read authentication configuration based on type
	switch secret.WinRmCredentials.Auth.Type_ {
	case "NTLM":
		if secret.WinRmCredentials.Auth.NtlmConfig != nil {
			ntlmMap := map[string]interface{}{
				"domain":          secret.WinRmCredentials.Auth.NtlmConfig.Domain,
				"username":        secret.WinRmCredentials.Auth.NtlmConfig.Username,
				"password_ref":    secret.WinRmCredentials.Auth.NtlmConfig.Password,
				"use_ssl":         secret.WinRmCredentials.Auth.NtlmConfig.UseSSL,
				"skip_cert_check": secret.WinRmCredentials.Auth.NtlmConfig.SkipCertChecks,
				"use_no_profile":  secret.WinRmCredentials.Auth.NtlmConfig.UseNoProfile,
			}
			d.Set("ntlm", []map[string]interface{}{ntlmMap})
		}

	case "Kerberos":
		if secret.WinRmCredentials.Auth.KerberosConfig != nil {
			d.Set("kerberos", []map[string]interface{}{
				{
					"principal":                  secret.WinRmCredentials.Auth.KerberosConfig.Principal,
					"realm":                      secret.WinRmCredentials.Auth.KerberosConfig.Realm,
					"tgt_generation_method":      secret.WinRmCredentials.Auth.KerberosConfig.TgtGenerationMethod,
					"use_ssl":                    secret.WinRmCredentials.Auth.KerberosConfig.UseSSL,
					"skip_cert_check":            secret.WinRmCredentials.Auth.KerberosConfig.SkipCertChecks,
					"use_no_profile":             secret.WinRmCredentials.Auth.KerberosConfig.UseNoProfile,
					"tgt_key_tab_file_path_spec": readWinRmTgtKeyFilePathSpec(secret),
					"tgt_password_spec":          readWinRmTgtPasswordSpec(secret),
				},
			})
		}
	}

	return nil
}

func readWinRmTgtKeyFilePathSpec(secret *nextgen.Secret) []map[string]interface{} {
	var spec []map[string]interface{}
	if secret.WinRmCredentials.Auth.KerberosConfig.KeyTabFilePathSpec != nil {
		spec = []map[string]interface{}{
			{"key_path": secret.WinRmCredentials.Auth.KerberosConfig.KeyTabFilePathSpec.KeyPath},
		}
	}
	return spec
}

func readWinRmTgtPasswordSpec(secret *nextgen.Secret) []map[string]interface{} {
	var spec []map[string]interface{}
	if secret.WinRmCredentials.Auth.KerberosConfig.PasswordSpec != nil {
		spec = []map[string]interface{}{
			{"password_ref": secret.WinRmCredentials.Auth.KerberosConfig.PasswordSpec.Password},
		}
	}
	return spec
}
