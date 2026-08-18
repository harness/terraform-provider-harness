package idp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/harness/harness-go-sdk/harness/idp"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildPluginAppConfigRequest(t *testing.T) {
	raw := map[string]interface{}{
		"identifier": "harness-proxy",
		"name":       "Configure Backend Proxies",
		"configs":    "proxy:\n  /harness-api:\n    target: https://app.harness.io\n",
		"env_variables": []interface{}{
			map[string]interface{}{
				"env_name":                  "HARNESS_API_KEY",
				"type":                      "FROM_HARNESS_SECRET_MANAGER",
				"harness_secret_identifier": "my_secret",
			},
		},
		"proxy": []interface{}{
			map[string]interface{}{
				"host":      "app.harness.io",
				"proxy":     true,
				"selectors": []interface{}{"delegate1"},
			},
		},
	}

	resourceSchema := ResourcePlugin().Schema
	d := schema.TestResourceDataRaw(t, resourceSchema, raw)

	req := buildPluginAppConfigRequest(d)

	assert.Equal(t, "harness-proxy", req.AppConfig.ConfigId)
	assert.Equal(t, "Configure Backend Proxies", req.AppConfig.ConfigName)
	assert.True(t, req.AppConfig.Enabled)
	assert.Contains(t, req.AppConfig.Configs, "proxy:")
	require.Len(t, req.AppConfig.EnvVariables, 1)
	assert.Equal(t, "HARNESS_API_KEY", req.AppConfig.EnvVariables[0].EnvName)
	assert.Equal(t, "FROM_HARNESS_SECRET_MANAGER", req.AppConfig.EnvVariables[0].Type)
	assert.Equal(t, "my_secret", req.AppConfig.EnvVariables[0].HarnessSecretIdentifier)
	require.Len(t, req.AppConfig.Proxy, 1)
	assert.Equal(t, "app.harness.io", req.AppConfig.Proxy[0].Host)
	assert.True(t, req.AppConfig.Proxy[0].Proxy)
	assert.Equal(t, []string{"delegate1"}, req.AppConfig.Proxy[0].Selectors)
}

func TestBuildPluginAppConfigRequest_Minimal(t *testing.T) {
	raw := map[string]interface{}{
		"identifier": "my-plugin",
		"name":       "My Plugin",
		"configs":    "key: value",
	}

	resourceSchema := ResourcePlugin().Schema
	d := schema.TestResourceDataRaw(t, resourceSchema, raw)

	req := buildPluginAppConfigRequest(d)

	assert.Equal(t, "my-plugin", req.AppConfig.ConfigId)
	assert.Equal(t, "My Plugin", req.AppConfig.ConfigName)
	assert.True(t, req.AppConfig.Enabled)
	assert.Equal(t, "key: value", req.AppConfig.Configs)
	assert.Nil(t, req.AppConfig.EnvVariables)
	assert.Nil(t, req.AppConfig.Proxy)
}

func TestReadPluginState(t *testing.T) {
	config := "proxy:\n  /harness-api:\n    target: https://app.harness.io\n"
	resp := idp.PluginInfoResponse{
		Plugin: &idp.PluginInfoData{
			PluginDetails: &idp.PluginDetails{
				Id:      "harness-proxy",
				Name:    "Configure Backend Proxies",
				Enabled: true,
			},
			Config: &config,
			EnvVariables: []idp.PluginAppConfigEnvVar{
				{
					EnvName:                 "HARNESS_API_KEY",
					Type:                    "FROM_HARNESS_SECRET_MANAGER",
					HarnessSecretIdentifier: "my_secret",
				},
				{
					EnvName:                 "DELETED_VAR",
					Type:                    "FROM_HARNESS_SECRET_MANAGER",
					HarnessSecretIdentifier: "old_secret",
					IsDeleted:               true,
				},
			},
			Proxy: []idp.PluginAppConfigProxy{
				{
					Host:      "app.harness.io",
					Proxy:     true,
					Selectors: []string{"delegate1"},
				},
			},
		},
	}

	resourceSchema := ResourcePlugin().Schema
	d := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{
		"identifier": "harness-proxy",
		"name":       "Configure Backend Proxies",
		"configs":    "",
	})

	readPluginState(d, resp)

	assert.Equal(t, "harness-proxy", d.Id())
	assert.Equal(t, "harness-proxy", d.Get("identifier"))
	assert.Equal(t, "Configure Backend Proxies", d.Get("name"))
	assert.Equal(t, true, d.Get("enabled"))
	assert.Contains(t, d.Get("configs"), "proxy:")

	envVars := d.Get("env_variables").([]interface{})
	require.Len(t, envVars, 1)
	ev := envVars[0].(map[string]interface{})
	assert.Equal(t, "HARNESS_API_KEY", ev["env_name"])
	assert.Equal(t, "my_secret", ev["harness_secret_identifier"])

	proxies := d.Get("proxy").([]interface{})
	require.Len(t, proxies, 1)
	p := proxies[0].(map[string]interface{})
	assert.Equal(t, "app.harness.io", p["host"])
}

func TestPluginResourceCRUD_MockServer(t *testing.T) {
	pluginConfig := "proxy:\n  /test:\n    target: https://test.harness.io\n"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/v1/app-config":
			resp := idp.PluginAppConfigResponse{
				AppConfig: &idp.PluginAppConfigResponseData{
					ConfigId:   "test-plugin",
					ConfigName: "Test Plugin",
					Configs:    pluginConfig,
					Enabled:    true,
				},
			}
			json.NewEncoder(w).Encode(resp)

		case r.Method == http.MethodGet && r.URL.Path == "/v1/plugins-info/test-plugin":
			resp := idp.PluginInfoResponse{
				Plugin: &idp.PluginInfoData{
					PluginDetails: &idp.PluginDetails{
						Id:      "test-plugin",
						Name:    "Test Plugin",
						Enabled: true,
					},
					Config:       &pluginConfig,
					EnvVariables: []idp.PluginAppConfigEnvVar{},
					Proxy:        []idp.PluginAppConfigProxy{},
				},
			}
			json.NewEncoder(w).Encode(resp)

		case r.Method == http.MethodPost && r.URL.Path == "/v1/plugin-toggle/test-plugin":
			enabled := r.URL.Query().Get("enabled")
			resp := idp.PluginAppConfigResponse{
				AppConfig: &idp.PluginAppConfigResponseData{
					ConfigId:   "test-plugin",
					ConfigName: "Test Plugin",
					Configs:    pluginConfig,
					Enabled:    enabled == "true",
				},
			}
			json.NewEncoder(w).Encode(resp)

		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := idp.NewAPIClient(&idp.Configuration{
		BasePath:   server.URL + "/v1",
		AccountId:  "test-account",
		ApiKey:     "test-key",
		HTTPClient: retryablehttp.NewClient(),
	})

	_, ctx := client.WithAuthContext(context.Background())

	// Test SaveOrUpdate
	saveResp, err := client.PluginAppConfigApi.SaveOrUpdate(ctx, idp.PluginAppConfigRequest{
		AppConfig: idp.PluginAppConfig{
			ConfigId:   "test-plugin",
			ConfigName: "Test Plugin",
			Enabled:    true,
			Configs:    pluginConfig,
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "test-plugin", saveResp.AppConfig.ConfigId)
	assert.True(t, saveResp.AppConfig.Enabled)

	// Test GetPluginInfo
	infoResp, err := client.PluginAppConfigApi.GetPluginInfo(ctx, "test-plugin")
	require.NoError(t, err)
	assert.Equal(t, "test-plugin", infoResp.Plugin.PluginDetails.Id)
	assert.True(t, infoResp.Plugin.PluginDetails.Enabled)

	// Test Toggle disable
	toggleResp, err := client.PluginAppConfigApi.Toggle(ctx, "test-plugin", false)
	require.NoError(t, err)
	assert.False(t, toggleResp.AppConfig.Enabled)
}

func TestGetPluginErrorMessage(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "plain error",
			err:      fmt.Errorf("connection refused"),
			expected: "connection refused",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := getPluginErrorMessage(tt.err)
			assert.Equal(t, tt.expected, msg)
		})
	}
}
