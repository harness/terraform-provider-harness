package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/micahlmartin/terraform-provider-harness/internal/client"
)

func dataSourceGitConnector() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving a Harness application",

		ReadContext: dataSourceGitConnectorRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the encrypted secret",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The name of the encrypted secret",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"usage_scopes": usageScopeSchema(),
		},
	}
}

func dataSourceGitConnectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	c := meta.(*client.ApiClient)

	id := d.Get("id").(string)
	conn, err := c.Connectors().GetGitConnectorById(id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", conn.Name)
	d.Set("usage_scopes", flattenUsageScope(conn.UsageScope))

	return nil
}
