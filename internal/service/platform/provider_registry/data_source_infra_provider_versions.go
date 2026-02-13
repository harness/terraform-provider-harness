package provider_registry

import (
	"context"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInfraProviderVersions() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for listing Terraform/OpenTofu Provider Versions from the IaCM Provider Registry.",
		ReadContext: dataSourceInfraProviderVersionsRead,

		Schema: map[string]*schema.Schema{
			"account": {
				Description: "Account identifier.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": {
				Description: "Provider type (e.g., aws, azurerm, google).",
				Type:        schema.TypeString,
				Required:    true,
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
						"protocols": {
							Description: "Supported Terraform protocol versions.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"platforms": {
							Description: "Supported platforms.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"os": {
										Description: "Operating system.",
										Type:        schema.TypeString,
										Computed:    true,
									},
									"arch": {
										Description: "Architecture.",
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func dataSourceInfraProviderVersionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c, ctx := m.(*internal.Session).GetPlatformClientWithContext(ctx)

	account := d.Get("account").(string)
	providerType := d.Get("type").(string)

	resp, httpRes, err := c.ProviderRegistryApi.ProviderRegistryListProviderVersions(
		ctx,
		account,
		providerType,
	)
	if err != nil {
		return helpers.HandleApiError(err, d, httpRes)
	}
	readProviderVersionsList(d, resp, account, providerType)
	return nil
}

func readProviderVersionsList(d *schema.ResourceData, response nextgen.ListProviderVersionsResponse, account string, providerType string) {
	d.SetId(account + "/" + providerType)
	
	if len(response.Versions) > 0 {
		versionsList := make([]interface{}, len(response.Versions))
		for i, v := range response.Versions {
			versionMap := map[string]interface{}{
				"version":   v.Version,
				"protocols": v.Protocols,
			}
			
			if len(v.Platforms) > 0 {
				platformsList := make([]interface{}, len(v.Platforms))
				for j, p := range v.Platforms {
					platformMap := map[string]interface{}{
						"os":   p.Os,
						"arch": p.Arch,
					}
					platformsList[j] = platformMap
				}
				versionMap["platforms"] = platformsList
			}
			
			versionsList[i] = versionMap
		}
		d.Set("versions", versionsList)
	}
}
