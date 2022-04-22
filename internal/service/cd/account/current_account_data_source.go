package account

import (
	"context"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceCurrentAccountConnector() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving information about the current Harness account",

		ReadContext: dataSourceGitConnectorCurrentAccount,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Id of the git connector.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"account_id": {
				Description: "Id of the account.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"url": {
				Description: "The url of the Harness control plane.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceGitConnectorCurrentAccount(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient

	d.SetId(c.Configuration.AccountId)
	d.Set("account_id", c.Configuration.AccountId)
	d.Set("url", c.Configuration.Endpoint)

	return nil
}
