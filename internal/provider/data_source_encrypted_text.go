package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
)

func dataSourceEncryptedText() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving a Harness application",

		ReadContext: dataSourceEncryptedTextRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the encrypted secret",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the encrypted secret",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"secret_manager_id": {
				Description: "The id of the associated secret manager",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"usage_scopes": usageScopeSchema(),
		},
	}
}

func dataSourceEncryptedTextRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	c := meta.(*client.ApiClient)

	var secret *client.EncryptedText
	var err error

	if id := d.Get("id").(string); id != "" {
		// Try lookup by Id first
		secret, err = c.Secrets().GetEncryptedTextById(id)
		if err != nil {
			return diag.FromErr(err)
		}
	} else if name := d.Get("name").(string); name != "" {
		// Fallback to lookup by name
		name := d.Get("name").(string)
		secret, err = c.Secrets().GetEncryptedTextByName(name)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		// Throw error if neither are set
		return diag.Errorf("id or name must be set")
	}

	d.SetId(secret.Id)
	d.Set("name", secret.Name)
	d.Set("secret_manager_id", secret.SecretManagerId)

	if secret.UsageScope != nil {
		usageScopes := flattenAppEnvScopes(secret.UsageScope.AppEnvScopes)
		if err := d.Set("usage_scopes", usageScopes); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}
