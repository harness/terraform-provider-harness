package provider_registry

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInfraProviders() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for listing Terraform/OpenTofu Providers from the IaCM Provider Registry.",
		ReadContext: dataSourceInfraProvidersRead,

		Schema: map[string]*schema.Schema{
			"providers": {
				Description: "List of providers.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "Unique identifier of the provider.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": {
							Description: "Provider type.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": {
							Description: "Description of the provider.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"account": {
							Description: "Account that owns the provider.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"created": {
							Description: "Timestamp when the provider was created.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"updated": {
							Description: "Timestamp when the provider was last updated.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
	return resource
}

func dataSourceInfraProvidersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryListProviders(
		ctx,
		c.AccountId,
		nil,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}
	readProviders(d, resp)
	return nil
}

func readProviders(d *schema.ResourceData, providers []nextgen.Provider) {
	d.SetId("providers")

	if len(providers) > 0 {
		providersList := make([]interface{}, len(providers))
		for i, p := range providers {
			providerMap := map[string]interface{}{
				"id":          p.Identifier,
				"type":        string(p.Type_),
				"description": p.Description,
				"account":     p.AccountIdentifier,
				"created":     int64(0),
				"updated":     p.LastModifiedAt,
			}
			providersList[i] = providerMap
		}
		d.Set("providers", providersList)
	}
}
