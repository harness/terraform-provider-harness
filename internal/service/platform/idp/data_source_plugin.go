package idp

import (
	"context"

	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourcePlugin() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving IDP plugin configuration.",
		ReadContext: dataSourcePluginRead,
		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of the plugin.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Display name of the plugin.",
			},
			"configs": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Backstage YAML configuration for the plugin.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the plugin is currently enabled.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the plugin.",
			},
			"category": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Category of the plugin.",
			},
			"source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Source of the plugin.",
			},
			"icon_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Icon URL of the plugin.",
			},
			"documentation": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Documentation URL for the plugin.",
			},
			"plugin_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the plugin.",
			},
			"core": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this is a core plugin.",
			},
			"env_variables": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Secret environment variables injected into the plugin runtime.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"env_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the environment variable.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the environment variable source.",
						},
						"harness_secret_identifier": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Harness secret identifier used as the value.",
						},
						"identifier": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Server-generated unique identifier for this env variable entry.",
						},
					},
				},
			},
			"proxy": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Delegate-based proxy configuration for outbound HTTP calls.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Proxy host.",
						},
						"proxy": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether proxy is enabled for this host.",
						},
						"selectors": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Delegate selectors.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"identifier": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Server-generated unique identifier for this proxy entry.",
						},
						"health_check_path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Health check path for the proxy endpoint.",
						},
					},
				},
			},
		},
	}
	return resource
}

func dataSourcePluginRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	pluginId := d.Get("identifier").(string)

	resp, err := c.PluginAppConfigApi.GetPluginInfo(ctx, pluginId)
	if err != nil {
		return diag.Errorf("failed to read plugin %s: %s", pluginId, getPluginErrorMessage(err))
	}

	if resp.Plugin == nil {
		return diag.Errorf("plugin %s not found", pluginId)
	}

	if err := readPluginState(d, resp); err != nil {
		return diag.Errorf("failed to read plugin %s state: %s", pluginId, err.Error())
	}

	if resp.Plugin.PluginDetails != nil {
		details := resp.Plugin.PluginDetails
		d.Set("description", details.Description)
		d.Set("category", details.Category)
		d.Set("source", details.Source)
		d.Set("icon_url", details.IconUrl)
		d.Set("documentation", details.Documentation)
		d.Set("plugin_type", details.PluginType)
		d.Set("core", details.Core)
	}

	return nil
}
