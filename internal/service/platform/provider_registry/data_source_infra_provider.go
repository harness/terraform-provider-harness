package provider_registry

import (
	"context"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInfraProvider() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Terraform/OpenTofu Provider from the IaCM Provider Registry.",
		ReadContext: dataSourceInfraProviderRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the provider.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": {
				Description: "Provider type (e.g., aws, azurerm, google).",
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
			"versions": {
				Description: "List of provider versions.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": {
							Description: "Version number.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"synced": {
							Description: "Whether the version is synced.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"files": {
							Description: "List of uploaded files for this version.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func dataSourceInfraProviderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)
	id := d.Get("id").(string)

	resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryGetProvider(
		ctx,
		id,
		c.AccountId,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}
	readProvider(d, &resp)
	return nil
}
