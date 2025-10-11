package image_registry

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	resourceName = "harness_chaos_image_registry"
)

func ResourceChaosImageRegistry() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for managing a Harness Chaos Image Registry",
		CreateContext: resourceChaosImageRegistryUpdate,
		ReadContext:   resourceChaosImageRegistryRead,
		UpdateContext: resourceChaosImageRegistryUpdate,
		DeleteContext: resourceChaosImageRegistryDelete,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description:  "The organization ID of the image registry",
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"project_id": {
				Description:  "The project ID of the image registry",
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"infra_id": {
				Description:  "The infrastructure ID to set up the image registry",
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"registry_server": {
				Description:  "The registry server URL",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"registry_account": {
				Description:  "The registry account name",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"is_private": {
				Description: "Whether the registry is private",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"is_default": {
				Description: "Whether this is the default registry",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"is_override_allowed": {
				Description: "Whether override is allowed for this registry",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"use_custom_images": {
				Description: "Whether to use custom images",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"secret_name": {
				Description:  "The name of the secret for authentication",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"custom_images": {
				Description: "Custom images configuration",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_watcher": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"ddcr": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"ddcr_lib": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"ddcr_fault": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.StringIsNotEmpty,
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
