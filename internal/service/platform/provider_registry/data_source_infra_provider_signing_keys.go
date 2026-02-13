package provider_registry

import (
	"context"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInfraProviderSigningKeys() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for listing GPG Signing Keys from the IaCM Provider Registry.",
		ReadContext: dataSourceInfraProviderSigningKeysRead,

		Schema: map[string]*schema.Schema{
			"signing_keys": {
				Description: "List of signing keys.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "Unique identifier of the signing key.",
							Type:        schema.TypeString,
							Computed:    true,
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
				},
			},
		},
	}
	return resource
}

func dataSourceInfraProviderSigningKeysRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryListSigningKeys(
		ctx,
		c.AccountId,
		nil,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}
	readSigningKeysList(d, resp)
	return nil
}

func readSigningKeysList(d *schema.ResourceData, keys []nextgen.SigningKey) {
	d.SetId("signing_keys")

	if len(keys) > 0 {
		keysList := make([]interface{}, len(keys))
		for i, key := range keys {
			keyMap := map[string]interface{}{
				"id":          key.Id,
				"key_id":      key.KeyId,
				"key_name":    key.KeyName,
				"ascii_armor": key.AsciiArmor,
				"user":        key.User,
				"created_at":  key.CreatedAt,
				"updated_at":  key.UpdatedAt,
			}
			keysList[i] = keyMap
		}
		d.Set("signing_keys", keysList)
	}
}
