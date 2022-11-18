package application

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/harness/terraform-provider-harness/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceApplicationGitSync() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Resource for configuring application git sync.",

		CreateContext: resourceApplicationGitSyncCreateOrUpdate,
		ReadContext:   resourceApplicationGitSyncRead,
		UpdateContext: resourceApplicationGitSyncCreateOrUpdate,
		DeleteContext: resourceApplicationGitSyncDelete,

		Schema: map[string]*schema.Schema{
			"app_id": {
				Description: "The id of the application.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"branch": {
				Description: "The branch of the git repository to sync to.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"connector_id": {
				Description: "The id of the git connector to use.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"repository_name": {
				Description: "The name of the git repository to sync to. This is only used if the git connector is for an account and not an individual repository.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"enabled": {
				Description: "Whether or not to enable git sync.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				d.Set("app_id", d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func resourceApplicationGitSyncCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}
	input := &graphql.UpdateApplicationGitSyncConfigInput{}

	if attr := d.Get("branch").(string); attr != "" {
		input.Branch = attr
	}

	if attr := d.Get("connector_id").(string); attr != "" {
		input.GitConnectorId = attr
	}

	if attr := d.Get("repository_name").(string); attr != "" {
		input.RepositoryName = attr
	}

	if attr := d.Get("enabled").(bool); attr {
		input.SyncEnabled = attr
	}

	if attr := d.Get("app_id").(string); attr != "" {
		input.ApplicationId = attr
	}

	config, err := c.ApplicationClient.UpdateGitSyncConfig(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(input.ApplicationId)
	readGitSyncConfig(d, config)

	return nil
}

func readGitSyncConfig(d *schema.ResourceData, config *graphql.GitSyncConfig) {
	if config == nil {
		return
	}

	d.Set("branch", config.Branch)

	if config.GitConnector != nil {
		d.Set("connector_id", config.GitConnector.Id)
	}

	d.Set("repository_name", config.RepositoryName)
	d.Set("enabled", config.SyncEnabled)
}

func resourceApplicationGitSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}
	appId := d.Get("app_id").(string)

	app, err := c.ApplicationClient.GetApplicationById(appId)
	if err != nil {
		return diag.FromErr(err)
	}

	if app == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	d.SetId(appId)
	readGitSyncConfig(d, app.GitSyncConfig)

	return nil
}

func resourceApplicationGitSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).CDClient
	if c == nil {
		return diag.Errorf(utils.CDClientAPIKeyError)
	}
	if err := c.ApplicationClient.RemoveGitSyncConfig(d.Get("app_id").(string)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
