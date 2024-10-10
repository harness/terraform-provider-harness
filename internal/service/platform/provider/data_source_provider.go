package provider

import (
	"context"

	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceProvider() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for Harness Provider.",

		ReadContext: dataSourceProviderRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "The identifier of the provider entity.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
		},
	}
	return resource
}

func dataSourceProviderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)
	identifier := d.Get("identifier").(string)

	resp, httpResp, err := c.ProviderApi.GetProvider(ctx, identifier, c.AccountId)
	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readProvider(d, resp.Data)

	return nil
}
