package provider

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	id := d.Get("id").(string)
	appId := d.Get("app_id").(string)

	err := c.Services().DeleteService(appId, id)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func commonServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: "Id of the service",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"app_id": {
			Description: "The id of the application the service belongs to",
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Description: "Name of the service",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "Description of th service",
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}
