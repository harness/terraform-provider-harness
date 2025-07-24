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
		"org_id": {
			Description: "Unique identifier of the organization",
			Type:        schema.TypeString,
			Optional:    true,
			ConflictsWith: []string{
				"parent_ref", "space_ref",
			},
		},
		"project_id": {
			Description: "Unique identifier of the project",
			Type:        schema.TypeString,
			Optional:    true,
			ConflictsWith: []string{
				"parent_ref", "space_ref",
			},
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
		"config": {
			Description: "Configuration for the registry",
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Deprecated:  "This field is deprecated and will be removed in a future version. Use type and virtual or upstream instead at the root for new package types.",
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
						Description: "Source of the upstream (only for UPSTREAM type)",
						Type:        schema.TypeString,
						Optional:    true,
						ConflictsWith: []string{
							"config.0.upstream_proxies",
						},
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
									ValidateFunc: validation.StringInSlice([]string{"UserPassword", "Anonymous"},
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
							},
						},
					},
					"auth_type": {
						Description:  "Type of authentication for UPSTREAM registry type (UserPassword, Anonymous)",
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"UserPassword", "Anonymous"}, false),
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
		"type": {
			Description: "Type of registry (VIRTUAL or UPSTREAM)",
			Type:        schema.TypeString,
			// This should be required but we have to set it to optional for now for backwards compatibility
			Optional: true,
			ValidateFunc: validation.StringInSlice([]string{
				(string)(har.VIRTUAL_RegistryType),
				(string)(har.UPSTREAM_RegistryType),
			}, false),
			ConflictsWith: []string{
				"config",
			},
		},
		"virtual": {
			Type:     schema.TypeList, // or TypeSet
			Optional: true,
			MaxItems: 1,
			//ExactlyOneOf: []string{"virtual", "upstream"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"upstream_proxies": {
						Description: "List of upstream proxies",
						Type:        schema.TypeSet,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
				},
			},
			ConflictsWith: []string{
				"config",
			},
		},
		"upstream": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			//ExactlyOneOf: []string{"virtual", "upstream"},
			Elem: getUpstreamRegistrySchema(),
			ConflictsWith: []string{
				"config",
			},
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

func customizeRegistryDiff(ctx context.Context, d *schema.ResourceDiff, _ interface{}) error {
	// ensure type matches which block is set
	rType := d.Get("type").(string)

	if rType == string(har.VIRTUAL_RegistryType) {
		if d.Get("upstream") != nil {
			return fmt.Errorf("cannot set 'upstream' when type is VIRTUAL")
		}
	}

	if rType == string(har.UPSTREAM_RegistryType) {
		if d.Get("virtual") != nil {
			return fmt.Errorf("cannot set 'virtual' when type is UPSTREAM")
		}

		// HELM requires url
		if pt := d.Get("package_type").(string); pt == string(har.HELM_PackageType) {
			if u, ok := d.GetOk("upstream.0.url"); !ok || u.(string) == "" {
				return fmt.Errorf("'upstream.url' is required when package_type is HELM")
			}
		}
	}

	return nil
}
