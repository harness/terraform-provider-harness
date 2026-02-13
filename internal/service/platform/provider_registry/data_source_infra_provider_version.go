package provider_registry

import (
	"context"
	"fmt"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInfraProviderVersion() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a specific Terraform/OpenTofu Provider Version from the IaCM Provider Registry.",
		ReadContext: dataSourceInfraProviderVersionRead,

		Schema: map[string]*schema.Schema{
			"provider_id": {
				Description: "The ID of the provider this version belongs to.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"version": {
				Description: "Version number (e.g., 1.0.0).",
				Type:        schema.TypeString,
				Required:    true,
			},
			"gpg_key_id": {
				Description: "GPG key ID for signing.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"protocols": {
				Description: "Supported Terraform protocol versions.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
	return resource
}

func dataSourceInfraProviderVersionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	
	providerId := d.Get("provider_id").(string)
	version := d.Get("version").(string)

	resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryGetProviderVersion(
		ctx,
		providerId,
		version,
		c.AccountId,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}
	
	d.SetId(fmt.Sprintf("%s/%s", providerId, version))
	d.Set("provider_id", providerId)
	d.Set("version", version)
	d.Set("gpg_key_id", resp.GpgKeyId)
	d.Set("protocols", resp.Protocols)
	
	return nil
}
