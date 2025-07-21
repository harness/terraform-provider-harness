package chaos_hub

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	resourceName = "harness_chaos_hub"
)

func ResourceChaosHub() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for managing a Harness Chaos Hub",
		CreateContext: resourceChaosHubCreate,
		ReadContext:   resourceChaosHubRead,
		UpdateContext: resourceChaosHubUpdate,
		DeleteContext: resourceChaosHubDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceChaosHubImport,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "ID of the chaos hub",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"org_id": {
				Description:  "The organization ID of the chaos hub",
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"project_id": {
				Description:  "The project ID of the chaos hub",
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"name": {
				Description:  "Name of the chaos hub",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"connector_id": {
				Description:  "ID of the Git connector",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"connector_scope": {
				Description:  "Scope of the Git connector (PROJECT, ORGANISATION, or ACCOUNT)",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"PROJECT", "ORGANISATION", "ACCOUNT"}, false),
				Default:      "PROJECT",
			},
			"repo_branch": {
				Description:  "Git repository branch",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"repo_name": {
				Description: "Name of the Git repository (required for account-level connectors)",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": {
				Description: "Description of the chaos hub",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Description: "Tags to associate with the chaos hub",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return !d.HasChange("tags")
				},
			},
			"is_default": {
				Description: "Whether this is the default chaos hub",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
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
			"last_synced_at": {
				Description: "Timestamp of the last sync",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"is_available": {
				Description: "Whether the chaos hub is available",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"total_experiments": {
				Description: "Total number of experiments in the hub",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"total_faults": {
				Description: "Total number of faults in the hub",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
}
