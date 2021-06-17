package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/micahlmartin/terraform-provider-harness/harness/api"
	"github.com/micahlmartin/terraform-provider-harness/harness/api/graphql"
)

func dataSourceApplication() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Data source for retrieving a Harness application",

		ReadContext: dataSourceApplicationRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the application",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "The name of the application",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "The application description",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_manual_trigger_authorized": {
				Description: "When this is set to true, all manual triggers will require API Key authorization",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"git_sync_enabled": {
				Description: "True if git sync is enabled on this application",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"git_sync_connector_id": {
				Description: "The id of the git sync connector",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	c := meta.(*api.Client)

	var app *graphql.Application
	var err error

	if id := d.Get("id").(string); id != "" {
		// Try lookup by Id first
		app, err = c.Applications().GetApplicationById(id)
		if err != nil {
			return diag.FromErr(err)
		}
	} else if name := d.Get("name").(string); name != "" {
		// Fallback to lookup by name
		name := d.Get("name").(string)
		app, err = c.Applications().GetApplicationByName(name)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		// Throw error if neither are set
		return diag.Errorf("id or name must be set")
	}

	d.SetId(app.Id)
	d.Set("name", app.Name)
	d.Set("description", app.Description)
	d.Set("is_manual_trigger_authorized", app.IsManualTriggerAuthorized)

	if app.GitSyncConfig != nil {
		d.Set("git_sync_enabled", app.GitSyncConfig.SyncEnabled)
		d.Set("git_sync_connector_id", app.GitSyncConfig.GitConnector.Id)
	}

	return nil
}
