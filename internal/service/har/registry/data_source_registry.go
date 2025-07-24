package registry

import (
	"context"
	"github.com/harness/harness-go-sdk/harness/har"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
)

func DataSourceRegistry() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for creating and managing Harness Registries.",
		ReadContext: dataSourceRegistryRead,
		Schema:      resourceRegistrySchema(true),
	}
}

func dataSourceRegistryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetHarClientWithContext(ctx)
	if c == nil {
		return diag.Errorf("Harness client is not initialized. Check provider configuration.")
	}

	var registry *har.Registry
	var err error
	var resp har.InlineResponse201
	var httpResp *http.Response

	id := d.Get("identifier").(string)
	registryRef := d.Get("space_ref").(string) + "/" + id

	if id != "" && registryRef != "" {
		resp, httpResp, err = c.RegistriesApi.GetRegistry(ctx, registryRef)
		if err != nil {
			return helpers.HandleReadApiError(err, d, httpResp)
		}

		registry = resp.Data
	} else {
		return diag.Errorf("Registry identifier and Space reference are required to read the registry.")
	}

	if registry.Identifier == "" {
		return diag.Errorf("Registry not found.")
	}

	readRegistry(d, registry)
	return nil
}
