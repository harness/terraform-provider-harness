package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
	"github.com/micahlmartin/terraform-provider-harness/internal/envvar"
)

func resourceEncryptedText() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating an encrypted text secret",
		CreateContext: resourceEncryptedTextCreate,
		ReadContext:   resourceEncryptedTextRead,
		UpdateContext: resourceEncryptedTextUpdate,
		DeleteContext: resourceEncryptedTextDelete,

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
			"inherit_scopes_from_secret_manager": {
				Description: "Boolean that indicates whether or not to inherit the usage scopes from the secret manager",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"scoped_to_account": {
				Description: "Boolean that indicates whether or not the secret is scoped to the account",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"secret_manager_id": {
				Description: fmt.Sprintf("The id of the secret manager to associate the secret with. Defaults to the value of %s", envvar.HarnessAccountId),
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: func() (interface{}, error) {
					if v := os.Getenv(envvar.HarnessAccountId); v != "" {
						return v, nil
					}

					return "", nil
				},
			},
			"value": {
				Description: "The value of the secret",
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
			},
			"usage_scope": usageScopeSchema(),
		},
	}
}

func resourceEncryptedTextRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceEncryptedTextCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceEncryptedTextUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceEncryptedTextDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.ApiClient)

	err := c.Secrets().DeleteSecret(&client.DeleteSecretInput{
		SecretId:   d.Get("id").(string),
		SecretType: client.SecretTypes.EncryptedText,
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
