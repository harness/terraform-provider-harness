package sso

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceSSOProvider() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving an SSO providers",

		ReadContext: dataSourceSSOProviderRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description:   "Unique identifier of the SSO provider.",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name"},
				AtLeastOneOf:  []string{"id", "name"},
			},
			"name": {
				Description:   "The name of the SSO provider.",
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"id"},
				AtLeastOneOf:  []string{"id", "name"},
			},
			"type": {
				Description: "The type of SSO provider.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceSSOProviderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	c := meta.(*internal.Session)

	var provider *graphql.SSOProvider
	var err error

	if id := d.Get("id").(string); id != "" {
		// Try lookup by Id first
		provider, err = c.CDClient.SSOClient.GetSSOProviderById(id)
		if err != nil {
			return diag.FromErr(err)
		}
	} else if name := d.Get("name").(string); name != "" {
		// Fallback to lookup by name
		name := d.Get("name").(string)
		provider, err = c.CDClient.SSOClient.GetSSOProviderByName(name)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		// Throw error if neither are set
		return diag.Errorf("id or name must be set")
	}

	d.SetId(provider.Id)
	d.Set("name", provider.Name)
	d.Set("type", provider.SSOType)

	return nil
}
