package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
)

func resourceSSHCredential() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating an encrypted text secret",
		CreateContext: resourceSSHCredentialCreate,
		ReadContext:   resourceSSHCredentialRead,
		UpdateContext: resourceSSHCredentialUpdate,
		DeleteContext: resourceSSHCredentialDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Id of the encrypted text secret",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Name of the encrypted text secret",
				Type:        schema.TypeString,
				Required:    true,
			},
			"kerberos_authentication": {
				Description: "Kerberos authentication for SSH. Cannot be used if ssh_authentication is specified",
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Description: "Port to use for Kerberos authentication",
							Type:        schema.TypeInt,
							Required:    true,
						},
						"principal": {
							Description: "Name of the principal for authentication",
							Type:        schema.TypeString,
							Required:    true,
						},
						"realm": {
							Description: "Realm associated with the Kerberos authentication",
							Type:        schema.TypeString,
							Required:    true,
						},
						"tgt_generation_method": {
							Description: "TGT generation method",
							Type:        schema.TypeSet,
							Optional:    true,
							MaxItems:    1,
							Elem:        &schema.Resource{
								// Schema: ,
							},
						},
					},
				},
			},
			"usage_scope": usageScopeSchema(),
		},
	}
}

func resourceSSHCredentialRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.ApiClient)

	secretId := d.Get("id").(string)

	secret, err := c.Secrets().GetEncryptedTextById(secretId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", secret.Name)
	d.Set("inherit_scopes_from_secret_manager", secret.InheritScopesFromSM)
	d.Set("scoped_to_account", secret.ScopedToAccount)
	d.Set("secret_manager_id", secret.SecretManagerId)

	if secret.UsageScope != nil {
		scopes := flattenAppEnvScopes(secret.UsageScope.AppEnvScopes)
		d.Set("usage_scope", scopes)
	}

	return nil
}

func resourceSSHCredentialCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.ApiClient)

	input := &client.CreateSecretInput{
		EncryptedText: &client.EncryptedTextInput{
			InheritScopesFromSM: d.Get("inherit_scopes_from_secret_manager").(bool),
			Name:                d.Get("name").(string),
			ScopedToAccount:     d.Get("scoped_to_account").(bool),
			SecretManagerId:     d.Get("secret_manager_id").(string),
			Value:               d.Get("value").(string),
		},
	}

	scopes := d.Get("usage_scope").(*schema.Set)
	var usageScopes []*client.AppEnvScope
	for _, sc := range scopes.List() {
		usageScopes = append(usageScopes, expandUsageScopeObject(sc))
	}

	if len(usageScopes) > 0 {
		input.EncryptedText.UsageScope = &client.UsageScope{
			AppEnvScopes: usageScopes,
		}
	}

	secret, err := c.Secrets().CreateEncryptedText(input)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(secret.Id)

	return nil
}

func resourceSSHCredentialUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.ApiClient)

	// Validation
	if d.HasChange("secret_manager_id") {
		return diag.Errorf("secret_manager_id is immutable and cannot be changed once set")
	}

	input := &client.UpdateSecretInput{
		SecretId: d.Get("id").(string),
		EncryptedText: &client.UpdateEncryptedText{
			InheritScopesFromSM: d.Get("inherit_scopes_from_secret_manager").(bool),
			Name:                d.Get("name").(string),
			ScopedToAccount:     d.Get("scoped_to_account").(bool),
			Value:               d.Get("value").(string),
		},
	}

	scopes := d.Get("usage_scope").(*schema.Set)
	var usageScopes []*client.AppEnvScope
	for _, sc := range scopes.List() {
		usageScopes = append(usageScopes, expandUsageScopeObject(sc))
	}

	if len(usageScopes) > 0 {
		input.EncryptedText.UsageScope = &client.UsageScope{
			AppEnvScopes: usageScopes,
		}
	}

	_, err := c.Secrets().UpdateEncryptedText(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceSSHCredentialDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.ApiClient)

	err := c.Secrets().DeleteSecret(d.Get("id").(string), client.SecretTypes.EncryptedText)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
