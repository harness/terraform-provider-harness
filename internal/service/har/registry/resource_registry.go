package registry

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/har"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"net/http"
	"regexp"
)

func ResourceRegistry() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for creating and managing Harness Registries.",

		ReadContext:   resourceRegistryRead,
		CreateContext: resourceRegistryCreateOrUpdate,
		UpdateContext: resourceRegistryCreateOrUpdate,
		DeleteContext: resourceRegistryDelete,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the registry",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Description: "Description of the registry",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"org_id": {
				Description: "Unique identifier of the organization",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Unique identifier of the project",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"parent_ref": {
				Description: "Parent reference for the registry",
				Type:        schema.TypeString,
				Optional:    true,
				Deprecated:  "This field is deprecated and will be removed in a future version. Use org_id and/or project_id instead",
			},
			"space_ref": {
				Description: "Space reference for the registry",
				Type:        schema.TypeString,
				Optional:    true,
				Deprecated:  "This field is deprecated and will be removed in a future version. Use org_id and/or project_id instead",
			},
			"package_type": {
				Description: "Type of package (DOCKER, HELM, MAVEN, etc.)",
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validation.StringInSlice([]string{
					(string)(har.DOCKER_PackageType),
					(string)(har.MAVEN_PackageType),
					(string)(har.PYTHON_PackageType),
					(string)(har.GENERIC_PackageType),
					(string)(har.HELM_PackageType),
					(string)(har.NUGET_PackageType),
					(string)(har.NPM_PackageType),
					(string)(har.RPM_PackageType),
					(string)(har.CARGO_PackageType),
				}, false),
			},
			"config": {
				Description: "Configuration for the registry",
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Description: "Type of registry (VIRTUAL or UPSTREAM)",
							Type:        schema.TypeString,
							Required:    true,
							ValidateFunc: validation.StringInSlice([]string{
								(string)(har.VIRTUAL_RegistryType),
								(string)(har.UPSTREAM_RegistryType),
							}, false),
						},
						// Virtual Config
						"upstream_proxies": {
							Description: "List of upstream proxies for VIRTUAL registry type",
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							ConflictsWith: []string{
								"config.0.source",
								"config.0.url",
								"config.0.auth",
								"config.0.auth_type",
							},
						},

						// Upstream Config
						"source": {
							Description: "Source of the upstream (only for UPSTREAM type)",
							Type:        schema.TypeString,
							Optional:    true,
							ConflictsWith: []string{
								"config.0.upstream_proxies",
							},
						},
						"url": {
							Description: "URL of the upstream",
							Type:        schema.TypeString,
							Optional:    true,
							ValidateFunc: validation.All(
								validation.StringMatch(
									regexp.MustCompile(`^https?://`),
									"URL must start with http:// or https://",
								),
							),
							ConflictsWith: []string{
								"config.0.upstream_proxies",
							},
						},
						"auth": {
							Description: "Authentication configuration for UPSTREAM registry type",
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							ConflictsWith: []string{
								"config.0.upstream_proxies",
							},
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_type": {
										Description: "Type of authentication (UserPassword, Anonymous)",
										Type:        schema.TypeString,
										Required:    false,
										ValidateFunc: validation.StringInSlice([]string{
											(string)(har.USER_PASSWORD_AuthType),
											(string)(har.ANONYMOUS_AuthType),
											(string)(har.ACCESS_KEY_SECRET_KEY_AuthType),
										},
											false),
										Deprecated: "This field is deprecated and will be removed in a future version. Use auth_type in config instead.",
									},
									"secret_identifier": {
										Description: "Secret identifier for UserPassword auth type",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"secret_space_path": {
										Description: "Secret space path for UserPassword auth type",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"user_name": {
										Description: "User name for UserPassword auth type",
										Type:        schema.TypeString,
										Optional:    true,
									},
								},
							},
						},
						"auth_type": {
							Description: "Type of authentication for UPSTREAM registry type (UserPassword, Anonymous)",
							Type:        schema.TypeString,
							Required:    true,
							ValidateFunc: validation.StringInSlice([]string{
								(string)(har.USER_PASSWORD_AuthType),
								(string)(har.ANONYMOUS_AuthType),
								(string)(har.ACCESS_KEY_SECRET_KEY_AuthType),
							}, false),
							ConflictsWith: []string{
								"config.0.upstream_proxies",
							},
						},
					},
					CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, i interface{}) error {
						configType := d.Get("config.0.type").(string)
						packageType := d.Get("package_type").(string)

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
							}
						}

						return nil
					},
				},
			},
			"url": {
				Description: "URL of the registry",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"created_at": {
				Description: "Timestamp when the registry was created",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"allowed_pattern": {
				Description: "Allowed pattern for the registry",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"blocked_pattern": {
				Description: "Blocked pattern for the registry",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceRegistryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetHarClientWithContext(ctx)

	registryRef := getParentRef(c.AccountId, d.Get("org_id").(string), d.Get("project_id").(string),
		d.Get("parent_ref").(string)) + "/" + d.Get("identifier").(string)
	resp, httpResp, err := c.RegistriesApi.GetRegistry(ctx, registryRef)

	if err != nil {
		return helpers.HandleReadApiError(err, d, httpResp)
	}

	readRegistry(d, resp.Data)
	return nil
}

func resourceRegistryCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetHarClientWithContext(ctx)
	if c == nil {
		return diag.Errorf("Harness client is not initialized. Check provider configuration.")
	}

	var err error
	var resp har.InlineResponse201
	var httpResp *http.Response

	registry := buildRegistry(d, c.AccountId)
	spaceRef := getParentRef(c.AccountId, d.Get("org_id").(string), d.Get("project_id").(string),
		d.Get("space_ref").(string))

	if d.Id() == "" {
		resp, httpResp, err = c.RegistriesApi.CreateRegistry(ctx, &har.RegistriesApiCreateRegistryOpts{
			Body: optional.NewInterface(registry), SpaceRef: optional.NewString(spaceRef),
		})
	} else {
		resp, httpResp, err = c.RegistriesApi.ModifyRegistry(ctx, d.Id(), &har.RegistriesApiModifyRegistryOpts{
			Body: optional.NewInterface(registry),
		})
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readRegistry(d, resp.Data)
	return nil
}

func resourceRegistryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetHarClientWithContext(ctx)

	registryRef := getParentRef(c.AccountId, d.Get("org_id").(string), d.Get("project_id").(string),
		d.Get("parent_ref").(string)) + "/" + d.Get("identifier").(string)

	_, httpResp, err := c.RegistriesApi.DeleteRegistry(ctx, registryRef)

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildRegistry(d *schema.ResourceData, accountID string) *har.RegistryRequest {
	registry := &har.RegistryRequest{}

	if attr, ok := d.GetOk("identifier"); ok {
		registry.Identifier = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		registry.Description = attr.(string)
	}

	parentRef := getParentRef(accountID, d.Get("org_id").(string), d.Get("project_id").(string),
		d.Get("parent_ref").(string))

	registry.ParentRef = parentRef

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
								}
								*upstreamConfig.AuthType = har.USER_PASSWORD_AuthType

							} else if authType == (string)(har.ANONYMOUS_AuthType) {
								upstreamConfig.Auth = &har.OneOfUpstreamConfigAuth{
									Anonymous: har.Anonymous{},
								}
							} else if authType == (string)(har.ACCESS_KEY_SECRET_KEY_AuthType) {
								accessKeySecretKey := har.AccessKeySecretKey{}
								if val, ok := authConfig["access_key"].(string); ok {
									accessKeySecretKey.AccessKey = val
								}
								if accessKeySecretKey.AccessKey == "" {
									if val, ok := authConfig["access_key_identifier"].(string); ok {
										accessKeySecretKey.AccessKeySecretIdentifier = val
									}
									if val, ok := authConfig["access_key_secret_path"].(string); ok {
										accessKeySecretKey.AccessKeySecretSpacePath = val
									}
								}
								if val, ok := authConfig["secret_key_identifier"].(string); ok {
									accessKeySecretKey.SecretKeyIdentifier = val
								}
								if val, ok := authConfig["secret_key_secret_path"].(string); ok {
									accessKeySecretKey.SecretKeySpacePath = val
								}

								upstreamConfig.Auth = &har.OneOfUpstreamConfigAuth{
									AccessKeySecretKey: accessKeySecretKey,
								}
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
			if registry.Config.UpstreamConfig.AuthType != nil {
				authMap := map[string]interface{}{}

				authMap["auth_type"] = *registry.Config.UpstreamConfig.AuthType
				switch *registry.Config.UpstreamConfig.AuthType {
				case har.USER_PASSWORD_AuthType:
					authMap["user_name"] = registry.Config.UpstreamConfig.Auth.UserPassword.UserName
					authMap["secret_identifier"] = registry.Config.UpstreamConfig.Auth.UserPassword.SecretIdentifier
					authMap["secret_space_path"] = registry.Config.UpstreamConfig.Auth.UserPassword.SecretSpacePath
				case har.ANONYMOUS_AuthType:
					break
				case har.ACCESS_KEY_SECRET_KEY_AuthType:
					authMap["access_key"] = registry.Config.UpstreamConfig.Auth.AccessKey
					authMap["access_key_identifier"] = registry.Config.UpstreamConfig.Auth.AccessKeySecretKey.AccessKeySecretIdentifier
					authMap["access_key_secret_path"] = registry.Config.UpstreamConfig.Auth.AccessKeySecretKey.AccessKeySecretSpacePath
					authMap["secret_key_identifier"] = registry.Config.UpstreamConfig.Auth.AccessKeySecretKey.SecretKeyIdentifier
					authMap["secret_key_secret_path"] = registry.Config.UpstreamConfig.Auth.AccessKeySecretKey.SecretKeySpacePath
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
