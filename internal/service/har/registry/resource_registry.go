package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/har"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// injectFirewallMode marshals a *har.RegistryRequest to a map and injects the
// firewallMode field into config, since the SDK struct does not include this
// field yet. Returns the map to pass to optional.NewInterface().
func injectFirewallMode(registry *har.RegistryRequest, firewallMode string) (interface{}, error) {
	data, err := json.Marshal(registry)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal registry request: %w", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to unmarshal registry request: %w", err)
	}
	cfg, ok := m["config"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to inject firewallMode: 'config' key missing from marshaled registry request")
	}
	cfg["firewallMode"] = firewallMode
	return m, nil
}

func ResourceRegistry() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating and managing Harness Registries.",
		ReadContext:   resourceRegistryRead,
		CreateContext: resourceRegistryCreateOrUpdate,
		UpdateContext: resourceRegistryCreateOrUpdate,
		DeleteContext: resourceRegistryDelete,
		Schema:        resourceRegistrySchema(false),
		Importer: &schema.ResourceImporter{
			StateContext: resourceRegistryImport,
		},
		CustomizeDiff: resourceRegistryCustomizeDiff,
	}
}

func isHarnessSecretManager(secretMgrId string) bool {
	return secretMgrId == "harnessSecretManager" ||
		secretMgrId == "account.harnessSecretManager" ||
		secretMgrId == "org.harnessSecretManager" ||
		secretMgrId == "project.harnessSecretManager"
}

func validateUpstreamRegistrySecrets(ctx context.Context, d *schema.ResourceData, meta interface{}) error {
	if configType, _ := d.Get("config.0.type").(string); configType != "UPSTREAM" {
		return nil
	}

	auth, ok := d.GetOk("config.0.auth")
	if !ok {
		return nil
	}

	authList, ok := auth.([]interface{})
	if !ok || len(authList) == 0 {
		return nil
	}
	authConfig, ok := authList[0].(map[string]interface{})
	if !ok {
		return nil
	}
	authType, _ := authConfig["auth_type"].(string)
	if authType == "Anonymous" {
		return nil
	}

	session := meta.(*internal.Session)
	platformClient, ctx := session.GetPlatformClientWithContext(ctx)

	// Parse space_ref for org/project scope
	var orgId, projectId string
	if parts := strings.Split(d.Get("space_ref").(string), "/"); len(parts) >= 3 {
		orgId = parts[1]
		projectId = parts[2]
	} else if len(parts) >= 2 {
		orgId = parts[1]
	}

	// Collect secrets to validate
	var secretsToValidate []string
	if authType == "UserPassword" {
		if secretId, ok := authConfig["secret_identifier"].(string); ok && secretId != "" {
			secretsToValidate = append(secretsToValidate, secretId)
		}
	} else if authType == "AccessKeySecretKey" {
		if accessKeyId, ok := authConfig["access_key_identifier"].(string); ok && accessKeyId != "" {
			secretsToValidate = append(secretsToValidate, accessKeyId)
		}
		if secretKeyId, ok := authConfig["secret_key_identifier"].(string); ok && secretKeyId != "" {
			secretsToValidate = append(secretsToValidate, secretKeyId)
		}
	}

	// Validate each secret
	for _, secretId := range secretsToValidate {
		opts := &nextgen.SecretsApiGetSecretV2Opts{}
		if orgId != "" {
			opts.OrgIdentifier = optional.NewString(orgId)
		}
		if projectId != "" {
			opts.ProjectIdentifier = optional.NewString(projectId)
		}

		resp, _, err := platformClient.SecretsApi.GetSecretV2(ctx, secretId, platformClient.AccountId, opts)
		if err != nil {
			return fmt.Errorf("failed to validate secret '%s': %v", secretId, err)
		}

		var secretMgrId string
		if resp.Data != nil && resp.Data.Secret != nil {
			if resp.Data.Secret.Text != nil {
				secretMgrId = resp.Data.Secret.Text.SecretManagerIdentifier
			} else if resp.Data.Secret.File != nil {
				secretMgrId = resp.Data.Secret.File.SecretManagerIdentifier
			}
		}

		if !isHarnessSecretManager(secretMgrId) {
			return fmt.Errorf("secret '%s' uses secret manager '%s', but upstream registry authentication requires secrets to be stored in Harness Secret Manager", secretId, secretMgrId)
		}
	}

	return nil
}

func resourceRegistryCustomizeDiff(ctx context.Context, d *schema.ResourceDiff, i interface{}) error {
	configType, _ := d.Get("config.0.type").(string)
	packageType, _ := d.Get("package_type").(string)

	// firewall_mode is only supported on UPSTREAM registries with certain package types
	if fm, ok := d.GetOk("config.0.firewall_mode"); ok && fm.(string) != "" {
		if configType != "UPSTREAM" {
			return fmt.Errorf("'firewall_mode' is only valid for UPSTREAM registry type")
		}
		if packageType == "DOCKER" || packageType == "HELM" {
			return fmt.Errorf("'firewall_mode' is not supported for %s package type", packageType)
		}
	}

	if configType == "UPSTREAM" {
		// Source is required for UPSTREAM
		if source, ok := d.GetOk("config.0.source"); !ok || source.(string) == "" {
			return fmt.Errorf("'source' is required for UPSTREAM registry type")
		}

		// URL is required for HELM package type
		if packageType == "HELM" {
			if url, ok := d.GetOk("config.0.url"); !ok || url.(string) == "" {
				return fmt.Errorf("'url' is required for UPSTREAM registry type with HELM package type")
			}
		}

		// Authentication is required for UPSTREAM registry type
		hasAuth := false
		if _, ok := d.GetOk("config.0.auth"); ok {
			hasAuth = true
		}
		if _, ok := d.GetOk("config.0.auth_type"); ok {
			hasAuth = true
		}
		if !hasAuth {
			return fmt.Errorf("authentication is required for UPSTREAM registry type. Provide either 'config.auth_type' field or 'config.auth' block with authentication details")
		}

		// Validate auth configuration
		if auth, ok := d.GetOk("config.0.auth"); ok {
			authConfig := auth.([]interface{})[0].(map[string]interface{})
			authType := authConfig["auth_type"].(string)

			if authType == "UserPassword" {
				// Check required fields for UserPassword auth
				if userName, ok := authConfig["user_name"].(string); !ok || userName == "" {
					return fmt.Errorf("'user_name' is required for UserPassword authentication")
				}
				if secretId, ok := authConfig["secret_identifier"].(string); !ok || secretId == "" {
					return fmt.Errorf("'secret_identifier' is required for UserPassword authentication")
				}
			}

			if authType == "AccessKeySecretKey" {
				// Check required fields for AccessKeySecretKey auth
				if secretKeyId, ok := authConfig["secret_key_identifier"].(string); !ok || secretKeyId == "" {
					return fmt.Errorf("'secret_key_identifier' is required for AccessKeySecretKey authentication")
				}
			}
		}
	}

	return nil
}

// fetchFirewallMode makes a raw HTTP GET to the registry endpoint and extracts
// the firewallMode field from the JSON response, since the SDK struct does not
// include this field yet.
func fetchFirewallMode(c *har.APIClient, ctx context.Context, registryRef string) string {
	// The /+ suffix tells the HAR API to return the full registry detail including config fields.
	rawURL := c.Endpoint + "/registry/" + registryRef + "/+"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		tflog.Warn(ctx, "fetchFirewallMode: failed to create request", map[string]interface{}{"error": err.Error()})
		return ""
	}
	req.Header.Set("x-api-key", c.ApiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		tflog.Warn(ctx, "fetchFirewallMode: HTTP request failed", map[string]interface{}{"error": err.Error(), "url": rawURL})
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		tflog.Warn(ctx, "fetchFirewallMode: unexpected status code", map[string]interface{}{"status": resp.StatusCode, "url": rawURL})
		return ""
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Warn(ctx, "fetchFirewallMode: failed to read response body", map[string]interface{}{"error": err.Error()})
		return ""
	}

	var raw struct {
		Data struct {
			Config struct {
				FirewallMode string `json:"firewallMode"`
			} `json:"config"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &raw); err != nil {
		tflog.Warn(ctx, "fetchFirewallMode: failed to unmarshal response", map[string]interface{}{"error": err.Error()})
		return ""
	}
	return raw.Data.Config.FirewallMode
}

func resourceRegistryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetHarClientWithContext(ctx)

	registryRef := d.Get("parent_ref").(string) + "/" + d.Get("identifier").(string)
	resp, httpResp, err := c.RegistriesApi.GetRegistry(ctx, registryRef)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readRegistry(d, resp.Data)

	// Fetch firewall_mode from raw API response (SDK doesn't support this field yet)
	if resp.Data != nil && resp.Data.Config != nil &&
		resp.Data.Config.Type_ != nil && *resp.Data.Config.Type_ == har.UPSTREAM_RegistryType {
		firewallMode := fetchFirewallMode(c, ctx, registryRef)
		if firewallMode != "" {
			// Update the config in state to include firewall_mode
			if configRaw, ok := d.GetOk("config"); ok {
				configList := configRaw.([]interface{})
				if len(configList) > 0 {
					configMap := configList[0].(map[string]interface{})
					configMap["firewall_mode"] = firewallMode
					d.Set("config", []interface{}{configMap})
				}
			}
		}
	}

	return nil
}

func resourceRegistryCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetHarClientWithContext(ctx)
	if c == nil {
		return diag.Errorf("Harness client is not initialized. Check provider configuration.")
	}

	// Validate secrets are from Harness Secret Manager for upstream registries
	if err := validateUpstreamRegistrySecrets(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}

	var err error
	var resp har.InlineResponse201
	var httpResp *http.Response

	registry := buildRegistry(d)
	spaceRef := d.Get("space_ref").(string)

	var body interface{} = registry
	if fm, ok := d.GetOk("config.0.firewall_mode"); ok && fm.(string) != "" {
		injected, fwErr := injectFirewallMode(registry, fm.(string))
		if fwErr != nil {
			return diag.FromErr(fwErr)
		}
		body = injected
	}

	if d.Id() == "" {
		resp, httpResp, err = c.RegistriesApi.CreateRegistry(ctx, &har.RegistriesApiCreateRegistryOpts{
			Body: optional.NewInterface(body), SpaceRef: optional.NewString(spaceRef),
		})
	} else {
		registryRef := d.Get("parent_ref").(string) + "/" + d.Get("identifier").(string)
		resp, httpResp, err = c.RegistriesApi.ModifyRegistry(ctx, registryRef, &har.RegistriesApiModifyRegistryOpts{
			Body: optional.NewInterface(body),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readRegistry(d, resp.Data)

	// Fetch firewall_mode from raw API response (SDK doesn't support this field yet)
	if resp.Data != nil && resp.Data.Config != nil &&
		resp.Data.Config.Type_ != nil && *resp.Data.Config.Type_ == har.UPSTREAM_RegistryType {
		registryRef := d.Get("parent_ref").(string) + "/" + d.Get("identifier").(string)
		firewallMode := fetchFirewallMode(c, ctx, registryRef)
		if firewallMode != "" {
			if configRaw, ok := d.GetOk("config"); ok {
				configList := configRaw.([]interface{})
				if len(configList) > 0 {
					configMap := configList[0].(map[string]interface{})
					configMap["firewall_mode"] = firewallMode
					d.Set("config", []interface{}{configMap})
				}
			}
		}
	}

	return nil
}

func resourceRegistryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetHarClientWithContext(ctx)

	registryRef := d.Get("parent_ref").(string) + "/" + d.Get("identifier").(string)

	_, httpResp, err := c.RegistriesApi.DeleteRegistry(ctx, registryRef)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func resourceRegistryImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	id := d.Id()

	if id == "" {
		return nil, fmt.Errorf("import ID cannot be empty")
	}

	parts := strings.Split(id, "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid import ID format. Expected: accountId/registry-name or accountId/orgId/projectId/registry-name")
	}

	for i, part := range parts {
		if part == "" {
			return nil, fmt.Errorf("component at position %d cannot be empty", i)
		}
	}

	identifier := parts[len(parts)-1]
	spaceRef := strings.Join(parts[:len(parts)-1], "/")

	d.Set("identifier", identifier)
	d.Set("space_ref", spaceRef)
	d.Set("parent_ref", spaceRef)
	d.SetId(identifier)

	diags := resourceRegistryRead(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to read registry '%s' at '%s': %v", identifier, spaceRef, diags)
	}

	return []*schema.ResourceData{d}, nil
}

func buildRegistry(d *schema.ResourceData) *har.RegistryRequest {
	registry := &har.RegistryRequest{}

	if attr, ok := d.GetOk("identifier"); ok {
		registry.Identifier = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		registry.Description = attr.(string)
	}

	if attr, ok := d.GetOk("parent_ref"); ok {
		registry.ParentRef = attr.(string)
	}

	if attr, ok := d.GetOk("package_type"); ok {
		pt := har.PackageType(attr.(string))
		registry.PackageType = &pt
	}
	if attr, ok := d.GetOk("config"); ok {
		configList := attr.([]interface{})
		if len(configList) > 0 { // Ensure config is not empty before accessing index
			config := configList[0].(map[string]interface{})

			if t, ok := config["type"].(string); ok {
				registryType := har.RegistryType(t)
				registry.Config = &har.RegistryConfig{Type_: &registryType}

				// Handle VIRTUAL type
				if registryType == har.VIRTUAL_RegistryType {
					// Initialize VirtualConfig to prevent nil pointer access
					registry.Config.VirtualConfig = har.VirtualConfig{}
					if proxies, ok := config["upstream_proxies"].([]interface{}); ok {
						for _, proxy := range proxies {
							registry.Config.VirtualConfig.UpstreamProxies = append(
								registry.Config.VirtualConfig.UpstreamProxies, proxy.(string))
						}
					}
				}

				// Handle UPSTREAM type
				if registryType == har.UPSTREAM_RegistryType {
					upstreamConfig := &har.UpstreamConfig{}

					if source, ok := config["source"].(string); ok {
						upstreamConfig.Source = source
					}
					if url, ok := config["url"].(string); ok {
						upstreamConfig.Url = url
					} else {
						upstreamConfig.Url = ""
					}

					// Handle authType at the top level
					if authType, ok := config["auth_type"].(string); ok {
						upstreamConfig.AuthType = (*har.AuthType)(&authType)
					}

					// Handle Authentication block
					if authAttr, ok := config["auth"].([]interface{}); ok && len(authAttr) > 0 {
						authConfig := authAttr[0].(map[string]interface{}) // Extract first element as map

						if authType, ok := authConfig["auth_type"].(string); ok {
							upstreamConfig.AuthType = (*har.AuthType)(&authType)

							if authType == (string)(har.USER_PASSWORD_AuthType) {
								userPassword := har.UserPassword{}

								if val, ok := authConfig["user_name"].(string); ok {
									userPassword.UserName = val
								}
								if val, ok := authConfig["secret_identifier"].(string); ok {
									userPassword.SecretIdentifier = val
								}
								if val, ok := authConfig["secret_space_path"].(string); ok {
									userPassword.SecretSpacePath = val
								}

								upstreamConfig.Auth = &har.OneOfUpstreamConfigAuth{
									UserPassword: userPassword,
									AuthType:     (*har.AuthType)(&authType),
								}
								*upstreamConfig.AuthType = har.USER_PASSWORD_AuthType

							} else if authType == (string)(har.ANONYMOUS_AuthType) {
								upstreamConfig.Auth = &har.OneOfUpstreamConfigAuth{
									Anonymous: har.Anonymous{},
								}
							} else if authType == (string)(har.ACCESS_KEY_SECRET_KEY_AuthType) {
								aksk := har.AccessKeySecretKey{}

								if val, ok := authConfig["access_key"].(string); ok {
									aksk.AccessKey = val
								}
								if val, ok := authConfig["access_key_identifier"].(string); ok {
									aksk.AccessKeySecretIdentifier = val
								}
								if val, ok := authConfig["access_key_secret_path"].(string); ok {
									aksk.AccessKeySecretSpacePath = val
								}
								if val, ok := authConfig["secret_key_identifier"].(string); ok {
									aksk.SecretKeyIdentifier = val
								}
								if val, ok := authConfig["secret_key_secret_path"].(string); ok {
									aksk.SecretKeySpacePath = val
								}

								upstreamConfig.Auth = &har.OneOfUpstreamConfigAuth{
									AccessKeySecretKey: aksk,
									AuthType:           (*har.AuthType)(&authType),
								}
								*upstreamConfig.AuthType = har.ACCESS_KEY_SECRET_KEY_AuthType

							}
						}
					}

					registry.Config.UpstreamConfig = *upstreamConfig
				}
			}
		}
	}

	if attr, ok := d.GetOk("allowed_pattern"); ok {
		registry.AllowedPattern = convertListToString(attr.([]interface{}))
	}

	if attr, ok := d.GetOk("blocked_pattern"); ok {
		registry.BlockedPattern = convertListToString(attr.([]interface{}))
	}

	return registry
}

func readRegistry(d *schema.ResourceData, registry *har.Registry) {
	d.SetId(registry.Identifier)
	d.Set("identifier", registry.Identifier)
	d.Set("description", registry.Description)
	d.Set("url", registry.Url)
	d.Set("package_type", registry.PackageType)
	d.Set("created_at", registry.CreatedAt)
	d.Set("allowed_pattern", registry.AllowedPattern)
	d.Set("blocked_pattern", registry.BlockedPattern)

	if registry.Config != nil {
		configMap := map[string]interface{}{
			"type": registry.Config.Type_,
		}

		// Handle VIRTUAL type configuration
		if registry.Config.Type_ != nil && *registry.Config.Type_ == har.VIRTUAL_RegistryType {
			if len(registry.Config.VirtualConfig.UpstreamProxies) > 0 {
				configMap["upstream_proxies"] = registry.Config.VirtualConfig.UpstreamProxies
			}
		}

		// Handle UPSTREAM type configuration
		if registry.Config.Type_ != nil && *registry.Config.Type_ == har.UPSTREAM_RegistryType {
			if registry.Config.UpstreamConfig.Source != "" {
				configMap["source"] = registry.Config.UpstreamConfig.Source
			}
			if registry.Config.UpstreamConfig.Url != "" {
				configMap["url"] = registry.Config.UpstreamConfig.Url
			}
			if registry.Config.UpstreamConfig.AuthType != nil {
				configMap["auth_type"] = *registry.Config.UpstreamConfig.AuthType
			}

			// Handle Authentication
			if registry.Config.UpstreamConfig.Auth != nil {
				authMap := map[string]interface{}{}
				if registry.Config.UpstreamConfig.Auth.UserPassword.UserName != "" {
					authMap["auth_type"] = "UserPassword"
					authMap["user_name"] = registry.Config.UpstreamConfig.Auth.UserPassword.UserName
					authMap["secret_identifier"] = registry.Config.UpstreamConfig.Auth.UserPassword.SecretIdentifier
					authMap["secret_space_path"] = registry.Config.UpstreamConfig.Auth.UserPassword.SecretSpacePath
				} else if registry.Config.UpstreamConfig.Auth.AccessKeySecretKey.SecretKeyIdentifier != "" {
					authMap["auth_type"] = "AccessKeySecretKey"
					authMap["access_key"] = registry.Config.UpstreamConfig.Auth.AccessKeySecretKey.AccessKey
					authMap["access_key_identifier"] = registry.Config.UpstreamConfig.Auth.AccessKeySecretKey.AccessKeySecretIdentifier
					authMap["access_key_secret_path"] = registry.Config.UpstreamConfig.Auth.AccessKeySecretKey.AccessKeySecretSpacePath
					authMap["secret_key_identifier"] = registry.Config.UpstreamConfig.Auth.AccessKeySecretKey.SecretKeyIdentifier
					authMap["secret_key_secret_path"] = registry.Config.UpstreamConfig.Auth.AccessKeySecretKey.SecretKeySpacePath
				} else {
					authMap["auth_type"] = "Anonymous"
				}
				configMap["auth"] = authMap
			}
		}

		// Set the updated config map to Terraform state
		d.Set("config", []interface{}{configMap})
	}
}

func convertListToString(list []interface{}) []string {
	result := make([]string, len(list))
	for i, v := range list {
		result[i] = v.(string)
	}
	return result
}
