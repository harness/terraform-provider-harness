package provider

import (
	"context"
	"log"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func commonCloudProviderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: "The id of the cloud provider.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"name": {
			Description: "The name of the cloud provider.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"usage_scope": usageScopeSchema(),
	}
}

func resourceCloudProviderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Deleting cloud provider %s", d.Get("name"))
	c := meta.(*api.Client)

	id := d.Get("id").(string)
	err := c.CloudProviders().DeleteCloudProvider(id)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
