package registry

import (
	"context"
	"fmt"
	"github.com/harness/harness-go-sdk/harness/har"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"regexp"
)

func resourceRegistrySchema(readOnly bool) map[string]*schema.Schema {
	mainSchema := map[string]*schema.Schema{
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
		"parent_ref": {
			Description: "Parent reference for the registry",
			Type:        schema.TypeString,
			Required:    true,
		},
		"space_ref": {
			Description: "Space reference for the registry",
			Type:        schema.TypeString,
			Required:    true,
		},
		"config": {
			Description: "Configuration for the registry",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Description: "Type of registry (VIRTUAL or UPSTREAM)",
						Type:        schema.TypeString,
						Required:    true,
						ValidateFunc: validation.StringInSlice([]string{
							"VIRTUAL",
							"UPSTREAM",
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
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Upstream source",
						ValidateFunc: validation.StringInSlice([]string{
							"Dockerhub", "Custom", "AwsEcr", "MavenCentral", "PyPi", "NpmJs", "NugetOrg", "Crates",
						}, false),
					},
					"url": {
						Description: "URL of the upstream (required if type=UPSTREAM & package_type=HELM)",
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
									Required:    true,
									ValidateFunc: validation.StringInSlice([]string{
										(string)(har.USER_PASSWORD_AuthType),
										(string)(har.ANONYMOUS_AuthType),
										(string)(har.ACCESS_KEY_SECRET_KEY_AuthType),
									},
										false),
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
								"access_key": {
									Type:      schema.TypeString,
									Optional:  true,
									Sensitive: true,
								},
								"access_key_identifier": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"access_key_secret_path": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"secret_key_identifier": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"secret_key_secret_path": {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
					"auth_type": {
						Description: "Type of authentication for UPSTREAM registry type (UserPassword, Anonymous)",
						Type:        schema.TypeString,
						Optional:    true,
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
		"url": {
			Description: "URL of the registry",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"created_at": {
			Description: "Creation timestamp",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"allowed_pattern": {
			Description: "Allowed artifact patterns",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"blocked_pattern": {
			Description: "Blocked artifact patterns",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}

	if readOnly {
		mainSchema["package_type"] = &schema.Schema{
			Description: "Type of package (DOCKER, HELM, MAVEN, etc.)",
			Type:        schema.TypeString,
			Optional:    true,
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
		}
	}
	return mainSchema
}

func getUpstreamRegistrySchema() *schema.Resource {
	urlRe := regexp.MustCompile(`^https?://`)
	urlValidator := validation.StringMatch(urlRe, "URL must start with http:// or https://")

	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"source": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Upstream source",
				ValidateFunc: validation.StringInSlice([]string{
					"Dockerhub", "Custom", "AwsEcr", "MavenCentral", "PyPi", "NpmJs", "NugetOrg", "Crates",
				}, false),
			},
			"url": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Upstream URL (required when package_type=HELM)",
				ValidateFunc: validation.All(urlValidator),
				DiffSuppressFunc: func(k, o, n string, _ *schema.ResourceData) bool {
					return trimSlash(o) == trimSlash(n)
				},
			},

			// ---- Auth (choose 0 or 1; 0 => default anonymous server-side) ----
			"anonymous": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				ExactlyOneOf: []string{
					"upstream.0.anonymous",
					"upstream.0.user_password",
					"upstream.0.access_key_secret_key",
				},
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{}},
			},
			"user_password": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				ExactlyOneOf: []string{
					"upstream.0.anonymous",
					"upstream.0.user_password",
					"upstream.0.access_key_secret_key",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:     schema.TypeString,
							Required: true,
						},
						"secret_identifier": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"secret_space_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"access_key_secret_key": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				ExactlyOneOf: []string{
					"upstream.0.anonymous",
					"upstream.0.user_password",
					"upstream.0.access_key_secret_key",
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_key": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"access_key_identifier": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"access_key_secret_path": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"secret_key_identifier": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
						"secret_key_secret_path": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func trimSlash(s string) string {
	for len(s) > 0 && s[len(s)-1] == '/' {
		s = s[:len(s)-1]
	}
	return s
}
