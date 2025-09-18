// internal/service/chaos/image_registry/data_source_chaos_image_registry.go
package image_registry

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceChaosImageRegistry() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Image Registry and checking override status",
		ReadContext: dataSourceChaosImageRegistryRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "The organization ID of the image registry",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "The project ID of the image registry",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"infra_id": {
				Description: "The infrastructure ID to set up the image registry",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"check_override": {
				Description: "Whether to check if override is allowed",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"override_blocked_by_scope": {
				Description: "Indicates if override is blocked by scope (only populated if check_override is true)",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"registry_server": {
				Description: "The registry server URL",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"registry_account": {
				Description: "The registry account name",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"is_private": {
				Description: "Whether the registry is private",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_default": {
				Description: "Whether this is the default registry",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_override_allowed": {
				Description: "Whether override is allowed for this registry",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"use_custom_images": {
				Description: "Whether custom images are used",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"secret_name": {
				Description: "The name of the secret for authentication",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"custom_images": {
				Description: "Custom images configuration",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_watcher": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ddcr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ddcr_lib": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ddcr_fault": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"created_at": {
				Description: "Creation timestamp",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"updated_at": {
				Description: "Last update timestamp",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceChaosImageRegistryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient

	identifiers := model.ScopedIdentifiersRequest{
		AccountIdentifier: c.AccountId,
	}

	if v, ok := d.GetOk("org_id"); ok {
		orgID := v.(string)
		identifiers.OrgIdentifier = &orgID
	}

	if v, ok := d.GetOk("project_id"); ok {
		projectID := v.(string)
		identifiers.ProjectIdentifier = &projectID
	}

	var infraID *string
	if v, ok := d.GetOk("infra_id"); ok {
		infraIDVal := v.(string)
		infraID = &infraIDVal
	}

	// If check_override is true, call CheckOverride API
	if d.Get("check_override").(bool) {
		response, err := c.ImageRegistryApi.CheckOverride(ctx, identifiers, infraID)
		if err != nil {
			return diag.Errorf("failed to check image registry override: %v", err)
		}

		d.Set("override_blocked_by_scope", response.OverrideBlockedByScope)

		// If we have registry data from the override check, use it
		if response.ImageRegistry != nil {
			return setImageRegistryData(d, response.ImageRegistry, identifiers)
		}
	}

	// Otherwise, use the standard Get API
	registry, err := c.ImageRegistryApi.Get(ctx, identifiers, infraID)
	if err != nil {
		return diag.Errorf("failed to read image registry: %v", err)
	}

	return setImageRegistryData(d, registry, identifiers)
}

func setImageRegistryData(d *schema.ResourceData, registry *model.ImageRegistryResponse, identifiers model.ScopedIdentifiersRequest) diag.Diagnostics {
	d.SetId(generateID(identifiers, registry.RegistryAccount))
	d.Set("registry_server", registry.RegistryServer)
	d.Set("org_id", identifiers.OrgIdentifier)
	d.Set("project_id", identifiers.ProjectIdentifier)
	d.Set("registry_account", registry.RegistryAccount)
	d.Set("infra_id", registry.InfraID)
	d.Set("is_private", registry.IsPrivate)
	d.Set("is_default", registry.IsDefault)
	d.Set("is_override_allowed", registry.IsOverrideAllowed)
	d.Set("use_custom_images", registry.UseCustomImages)
	d.Set("created_at", registry.CreatedAt)
	d.Set("updated_at", registry.UpdatedAt)

	if registry.SecretName != nil {
		d.Set("secret_name", *registry.SecretName)
	}

	if registry.CustomImages != nil {
		customImages := map[string]interface{}{}
		if registry.CustomImages.LogWatcher != nil {
			customImages["log_watcher"] = *registry.CustomImages.LogWatcher
		}
		if registry.CustomImages.Ddcr != nil {
			customImages["ddcr"] = *registry.CustomImages.Ddcr
		}
		if registry.CustomImages.DdcrLib != nil {
			customImages["ddcr_lib"] = *registry.CustomImages.DdcrLib
		}
		if registry.CustomImages.DdcrFault != nil {
			customImages["ddcr_fault"] = *registry.CustomImages.DdcrFault
		}
		if len(customImages) > 0 {
			d.Set("custom_images", []map[string]interface{}{customImages})
		}
	}

	return nil
}
