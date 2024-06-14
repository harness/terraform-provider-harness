package account

import (
	"context"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCurrentAccount() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving information about the current Harness account",

		ReadContext: dataSourceCurrentAccount,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Id of the account.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"account_id": {
				Description: "Id of the account.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"endpoint": {
				Description: "The url of the Harness control plane.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceCurrentAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	d.SetId(c.AccountId)
	d.Set("account_id", c.AccountId)
	d.Set("endpoint", c.Endpoint)

	return nil
}
