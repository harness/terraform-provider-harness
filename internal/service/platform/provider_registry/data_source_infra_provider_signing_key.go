package provider_registry

import (
	"context"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInfraProviderSigningKey() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a GPG Signing Key from the IaCM Provider Registry.",
		ReadContext: dataSourceInfraProviderSigningKeyRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the signing key.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"key_id": {
				Description: "GPG key ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"key_name": {
				Description: "GPG key name.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"ascii_armor": {
				Description: "ASCII-armored GPG public key.",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
			"user": {
				Description: "User who uploaded the key.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "Creation timestamp.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"updated_at": {
				Description: "Last updated timestamp.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
	return resource
}

func dataSourceInfraProviderSigningKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Get("id").(string)

	// List all signing keys and find the one with matching ID
	resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryListSigningKeys(
		ctx,
		c.AccountId,
		nil,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}

	// Find the signing key by ID
	for _, key := range resp {
		if key.Id == id {
			d.SetId(key.Id)
			d.Set("key_id", key.KeyId)
			d.Set("key_name", key.KeyName)
			d.Set("ascii_armor", key.AsciiArmor)
			d.Set("user", key.User)
			d.Set("created_at", key.CreatedAt)
			d.Set("updated_at", key.UpdatedAt)
			return nil
		}
	}

	return diag.Errorf("Signing key with ID %s not found", id)
}
