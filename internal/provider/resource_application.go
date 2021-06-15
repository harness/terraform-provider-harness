package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/micahlmartin/terraform-provider-harness/harness/graphql"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for creating a Harness application",

		CreateContext: resourceApplicationCreate,
		ReadContext:   resourceApplicationRead,
		UpdateContext: resourceApplicationUpdate,
		DeleteContext: resourceApplicationDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Unique identifier of the application",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the application",
				Type:        schema.TypeString,
				Required:    true,
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
				Computed:    true,
			},
			"git_sync_connector_id": {
				Description: "The id of the git sync connector",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*graphql.ApiClient)

	input := &graphql.Application{
		Name:                      d.Get("name").(string),
		Description:               d.Get("description").(string),
		IsManualTriggerAuthorized: d.Get("is_manual_trigger_authorized").(bool),
	}

	app, err := c.Applications().CreateApplication(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(app.Id)

	return nil
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*graphql.ApiClient)

	appId := d.Get("id").(string)

	app, err := c.Applications().GetApplicationById(appId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", app.Name)
	d.Set("description", app.Description)
	d.Set("is_manual_trigger_authorized", app.IsManualTriggerAuthorized)

	if app.GitSyncConfig != nil {
		d.Set("git_sync_enabled", app.GitSyncConfig.SyncEnabled)

		if app.GitSyncConfig.GitConnector != nil {
			d.Set("git_sync_connector_id", app.GitSyncConfig.GitConnector.Id)
		}
	}

	return nil
}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*graphql.ApiClient)

	input := &graphql.UpdateApplicationInput{
		ApplicationId:             d.Id(),
		Description:               d.Get("description").(string),
		IsManualTriggerAuthorized: d.Get("is_manual_trigger_authorized").(bool),
		Name:                      d.Get("name").(string),
	}

	_, err := c.Applications().UpdateApplication(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*graphql.ApiClient)

	if err := c.Applications().DeleteApplication(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
