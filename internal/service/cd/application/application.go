package application

import (
	"context"
	"fmt"
	"strings"

	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func ResourceApplication() *schema.Resource {
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
			"tags": helpers.GetTagsSchema(helpers.SchemaFlagTypes.Optional),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}
	input := &graphql.CreateApplicationInput{
		Name:                      d.Get("name").(string),
		Description:               d.Get("description").(string),
		IsManualTriggerAuthorized: d.Get("is_manual_trigger_authorized").(bool),
	}

	app, err := c.ApplicationClient.CreateApplication(input)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := attachTags(d, app, c); err != nil {
		return diag.FromErr(err)
	}

	app, err = c.ApplicationClient.GetApplicationById(app.Id)
	if err != nil {
		return diag.FromErr(err)
	}

	applicationRead(d, app)

	return nil
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}
	appId := d.Get("id").(string)

	app, err := c.ApplicationClient.GetApplicationById(appId)
	if err != nil {
		return diag.FromErr(err)
	}

	// In case the application was deleted since the state was lst updated.
	if app == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	applicationRead(d, app)

	return nil
}

func applicationRead(d *schema.ResourceData, app *graphql.Application) {
	if app == nil {
		return
	}

	d.SetId(app.Id)
	d.Set("name", app.Name)
	d.Set("description", app.Description)
	d.Set("is_manual_trigger_authorized", app.IsManualTriggerAuthorized)

	if app.GitSyncConfig != nil {
		d.Set("git_sync_enabled", app.GitSyncConfig.SyncEnabled)

		if app.GitSyncConfig.GitConnector != nil {
			d.Set("git_sync_connector_id", app.GitSyncConfig.GitConnector.Id)
		}
	}

	tags := []string{}
	for _, tag := range app.Tags {
		tags = append(tags, fmt.Sprintf("%s:%s", tag.Name, tag.Value))
	}
	if len(tags) > 0 {
		d.Set("tags", tags)
	}
}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}
	input := &graphql.UpdateApplicationInput{
		ApplicationId:             d.Id(),
		Description:               d.Get("description").(string),
		IsManualTriggerAuthorized: d.Get("is_manual_trigger_authorized").(bool),
		Name:                      d.Get("name").(string),
	}

	app, err := c.ApplicationClient.UpdateApplication(input)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChanges("tags") {
		for _, t := range app.Tags {
			if err := c.TagClient.DetachTag(&graphql.DetachTagInput{
				EntityId:   app.Id,
				EntityType: graphql.TagEntityTypes.Application,
				Name:       t.Name,
			}); err != nil {
				return diag.Errorf("failed to detach tag %s. %s", t.Name, err)
			}
		}

		if err := attachTags(d, app, c); err != nil {
			return diag.FromErr(err)
		}
	}

	app, err = c.ApplicationClient.GetApplicationById(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	applicationRead(d, app)

	return nil
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}
	if err := c.ApplicationClient.DeleteApplication(d.Id()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func attachTags(d *schema.ResourceData, app *graphql.Application, c *cd.ApiClient) error {
	for _, t := range d.Get("tags").(*schema.Set).List() {
		tagInput := &graphql.AttachTagInput{
			EntityId:   app.Id,
			EntityType: graphql.TagEntityTypes.Application,
		}
		tagStr := t.(string)
		parts := strings.Split(tagStr, ":")

		if len(parts) > 0 {
			tagInput.Name = parts[0]
		}

		if len(parts) > 1 {
			tagInput.Value = parts[1]
		}

		if _, err := c.TagClient.AttachTag(tagInput); err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to attach tag %s", tagInput.Name))
		}
	}
	return nil
}
