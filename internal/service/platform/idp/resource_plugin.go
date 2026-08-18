package idp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourcePlugin() *schema.Resource {
	resource := &schema.Resource{
		Description:   "Resource for managing IDP plugin configurations.",
		ReadContext:   resourcePluginRead,
		CreateContext: resourcePluginCreateOrUpdate,
		UpdateContext: resourcePluginCreateOrUpdate,
		DeleteContext: resourcePluginDelete,
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				d.Set("identifier", d.Id())
				c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)
				resp, err := c.PluginAppConfigApi.GetPluginInfo(ctx, d.Id())
				if err != nil {
					return nil, fmt.Errorf("failed to read plugin for import: %w", err)
				}
				if resp.Plugin == nil {
					return nil, fmt.Errorf("plugin %s not found", d.Id())
				}
				readPluginState(d, resp)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unique identifier of the plugin.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Display name of the plugin configuration.",
			},
			"configs": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Backstage YAML configuration for the plugin.",
				DiffSuppressFunc: helpers.YamlDiffSuppressFunction,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the plugin should be enabled. Defaults to true.",
			},
			"env_variables": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Secret environment variables injected into the plugin runtime.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"env_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the environment variable.",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of the environment variable source. Valid values: Secret, Config.",
						},
						"harness_secret_identifier": {
							Type:        schema.TypeString,
							Required:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "Delegate-based proxy configuration for outbound HTTP calls.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Proxy host.",
						},
						"proxy": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Whether proxy is enabled for this host.",
						},
						"selectors": {
							Type:        schema.TypeList,
							Optional:    true,
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
							Optional:    true,
							Description: "Health check path for the proxy endpoint.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourcePluginCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	pluginId := d.Get("identifier").(string)
	request := buildPluginAppConfigRequest(d)

	_, err := c.PluginAppConfigApi.SaveOrUpdate(ctx, request)
	if err != nil {
		return diag.Errorf("failed to save plugin configuration for %s: %s", pluginId, getPluginErrorMessage(err))
	}

	enabled := d.Get("enabled").(bool)
	_, err = c.PluginAppConfigApi.Toggle(ctx, pluginId, enabled)
	if err != nil {
		return diag.Errorf("failed to toggle plugin %s: %s", pluginId, getPluginErrorMessage(err))
	}

	d.SetId(pluginId)
	return resourcePluginRead(ctx, d, meta)
}

func resourcePluginRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	pluginId := d.Id()
	if pluginId == "" {
		pluginId = d.Get("identifier").(string)
	}

	resp, err := c.PluginAppConfigApi.GetPluginInfo(ctx, pluginId)
	if err != nil {
		if isPluginNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("failed to read plugin %s: %s", pluginId, getPluginErrorMessage(err))
	}

	if resp.Plugin == nil {
		d.SetId("")
		return nil
	}

	readPluginState(d, resp)
	return nil
}

func resourcePluginDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetIDPClientWithContext(ctx)

	pluginId := d.Id()

	_, err := c.PluginAppConfigApi.Toggle(ctx, pluginId, false)
	if err != nil {
		if isPluginNotFoundError(err) {
			return nil
		}
		return diag.Errorf("failed to disable plugin %s: %s", pluginId, getPluginErrorMessage(err))
	}

	return nil
}

func buildPluginAppConfigRequest(d *schema.ResourceData) idp.PluginAppConfigRequest {
	appConfig := idp.PluginAppConfig{
		ConfigId:     d.Get("identifier").(string),
		ConfigName:   d.Get("name").(string),
		Enabled:      true,
		Configs:      d.Get("configs").(string),
		EnvVariables: []idp.PluginAppConfigEnvVar{},
		Proxy:        []idp.PluginAppConfigProxy{},
	}

	if v, ok := d.GetOk("env_variables"); ok {
		envVars := v.([]interface{})
		appConfig.EnvVariables = make([]idp.PluginAppConfigEnvVar, len(envVars))
		for i, ev := range envVars {
			envMap := ev.(map[string]interface{})
			appConfig.EnvVariables[i] = idp.PluginAppConfigEnvVar{
				EnvName:                 envMap["env_name"].(string),
				Type:                    envMap["type"].(string),
				HarnessSecretIdentifier: envMap["harness_secret_identifier"].(string),
			}
		}
	}

	if v, ok := d.GetOk("proxy"); ok {
		proxies := v.([]interface{})
		appConfig.Proxy = make([]idp.PluginAppConfigProxy, len(proxies))
		for i, p := range proxies {
			proxyMap := p.(map[string]interface{})
			proxy := idp.PluginAppConfigProxy{
				Host:  proxyMap["host"].(string),
				Proxy: proxyMap["proxy"].(bool),
			}
			if hcp, ok := proxyMap["health_check_path"].(string); ok && hcp != "" {
				proxy.HealthCheckPath = &hcp
			}
			if selectors, ok := proxyMap["selectors"].([]interface{}); ok && len(selectors) > 0 {
				proxy.Selectors = make([]string, len(selectors))
				for j, s := range selectors {
					proxy.Selectors[j] = s.(string)
				}
			}
			appConfig.Proxy[i] = proxy
		}
	}

	return idp.PluginAppConfigRequest{AppConfig: appConfig}
}

func readPluginState(d *schema.ResourceData, resp idp.PluginInfoResponse) {
	plugin := resp.Plugin

	d.SetId(plugin.PluginDetails.Id)
	d.Set("identifier", plugin.PluginDetails.Id)
	d.Set("name", plugin.PluginDetails.Name)
	d.Set("enabled", plugin.PluginDetails.Enabled)

	if plugin.Config != nil {
		d.Set("configs", *plugin.Config)
	}

	if len(plugin.EnvVariables) > 0 {
		envVars := make([]map[string]interface{}, 0, len(plugin.EnvVariables))
		for _, ev := range plugin.EnvVariables {
			if ev.IsDeleted {
				continue
			}
			envVars = append(envVars, map[string]interface{}{
				"env_name":                  ev.EnvName,
				"type":                      ev.Type,
				"harness_secret_identifier": ev.HarnessSecretIdentifier,
				"identifier":               ev.Identifier,
			})
		}
		d.Set("env_variables", envVars)
	} else {
		d.Set("env_variables", []map[string]interface{}{})
	}

	if len(plugin.Proxy) > 0 {
		proxies := make([]map[string]interface{}, len(plugin.Proxy))
		for i, p := range plugin.Proxy {
			proxy := map[string]interface{}{
				"host":       p.Host,
				"proxy":      p.Proxy,
				"identifier": p.Identifier,
			}
			if p.HealthCheckPath != nil {
				proxy["health_check_path"] = *p.HealthCheckPath
			} else {
				proxy["health_check_path"] = ""
			}
			if len(p.Selectors) > 0 {
				proxy["selectors"] = p.Selectors
			} else {
				proxy["selectors"] = []string{}
			}
			proxies[i] = proxy
		}
		d.Set("proxy", proxies)
	} else {
		d.Set("proxy", []map[string]interface{}{})
	}
}

func getPluginErrorMessage(err error) string {
	if msg := idpAPIErrorMessage(err); msg != "" {
		return msg
	}
	return err.Error()
}

func isPluginNotFoundError(err error) bool {
	if isNotFoundError(err) {
		return true
	}
	swaggerErr, ok := err.(interface{ Body() []byte })
	if !ok {
		return false
	}
	_ = swaggerErr
	return false
}

func isPluginHTTPNotFound(httpResp *http.Response) bool {
	return httpResp != nil && httpResp.StatusCode == http.StatusNotFound
}

// pluginNotFoundFromError checks both error body and the fact that the SDK
// doesn't return httpResp directly from PluginAppConfigApi methods.
func pluginNotFoundFromError(err error) bool {
	if err == nil {
		return false
	}
	if isNotFoundError(err) {
		return true
	}
	swaggerErr, ok := err.(interface {
		Error() string
	})
	if ok {
		errStr := swaggerErr.Error()
		return errStr == fmt.Sprintf("%d %s", http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	return false
}
