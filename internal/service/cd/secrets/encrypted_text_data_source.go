package secrets

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/terraform-provider-harness/internal/service/cd/usagescope"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceEncryptedText() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving a Harness application",

		ReadContext: dataSourceEncryptedTextRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the encrypted secret",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the encrypted secret",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"secret_manager_id": {
				Description: "The id of the associated secret manager",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"usage_scope": usagescope.Schema(),
		},
	}
}

func dataSourceEncryptedTextRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	c := meta.(*cd.ApiClient)

	var secret *graphql.EncryptedText
	var err error

	if id := d.Get("id").(string); id != "" {
		// Try lookup by Id first
		secret, err = c.SecretClient.GetEncryptedTextById(id)
		if err != nil {
			return diag.FromErr(err)
		}
	} else if name := d.Get("name").(string); name != "" {
		// Fallback to lookup by name
		name := d.Get("name").(string)
		secret, err = c.SecretClient.GetEncryptedTextByName(name)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		// Throw error if neither are set
		return diag.Errorf("id or name must be set")
	}

	d.SetId(secret.Id)
	d.Set("name", secret.Name)
	d.Set("secret_manager_id", secret.SecretManagerId)
	d.Set("usage_scope", usagescope.FlattenUsageScope(secret.UsageScope))

	return nil
}
