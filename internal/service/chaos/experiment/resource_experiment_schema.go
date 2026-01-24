package experiment

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceChaosExperimentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// ========== REQUIRED FIELDS ==========
		"org_id": {
			Description: "Organization identifier where the experiment will be created",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"project_id": {
			Description: "Project identifier where the experiment will be created",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"template_identity": {
			Description: "Identity of the experiment template to launch from",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"hub_identity": {
			Description: "Identity of the hub where the experiment template resides (no prefix needed)",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"hub_org_id": {
			Description: "Organization identifier where the hub/template resides (leave empty for account-level hubs). This is used to locate the template, not where the experiment will be created.",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"hub_project_id": {
			Description: "Project identifier where the hub/template resides (leave empty for org-level or account-level hubs). This is used to locate the template, not where the experiment will be created.",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"name": {
			Description: "Name of the chaos experiment",
			Type:        schema.TypeString,
			Required:    true,
		},
		"infra_ref": {
			Description: "Infrastructure reference (ID or identity) to bind the experiment to",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		// ========== OPTIONAL INPUT FIELDS ==========
		"identity": {
			Description: "Unique identifier for the experiment (auto-generated if not provided)",
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
		},
		"description": {
			Description: "Description of the chaos experiment",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Description: "Tags to categorize the experiment",
			Type:        schema.TypeSet,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"revision": {
			Description: "Template revision to use (default: v1)",
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "v1",
			ForceNew:    true,
		},
		"import_type": {
			Description: "Import type: REFERENCE (template reference) or LOCAL (full copy). Default: REFERENCE",
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "REFERENCE",
			ForceNew:    true,
			ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
				v := val.(string)
				if v != "REFERENCE" && v != "LOCAL" {
					errs = append(errs, fmt.Errorf("%q must be either REFERENCE or LOCAL, got: %s", key, v))
				}
				return
			},
		},

		// ========== COMPUTED FIELDS ==========
		"experiment_id": {
			Description: "Full experiment ID",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"infra_id": {
			Description: "Resolved infrastructure ID",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"infra_type": {
			Description: "Infrastructure type (e.g., KubernetesV2)",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"experiment_type": {
			Description: "Type of the experiment",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"is_custom_experiment": {
			Description: "Whether this is a custom experiment",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"fault_ids": {
			Description: "List of fault IDs used in the experiment",
			Type:        schema.TypeList,
			Computed:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"cron_syntax": {
			Description: "Cron expression for scheduled execution",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"is_cron_enabled": {
			Description: "Whether cron scheduling is enabled",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"is_single_run_cron_enabled": {
			Description: "Whether single-run cron is enabled",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"last_executed_at": {
			Description: "Timestamp of last execution",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"total_experiment_runs": {
			Description: "Total number of experiment runs",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"target_network_map_id": {
			Description: "Target network map ID",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"created_at": {
			Description: "Creation timestamp (Unix)",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"created_by": {
			Description: "Username of the creator",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"updated_at": {
			Description: "Last update timestamp (Unix)",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"updated_by": {
			Description: "Username of the last updater",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"manifest": {
			Description: "Full experiment manifest YAML (populated for LOCAL imports)",
			Type:        schema.TypeString,
			Computed:    true,
		},

		// ========== NESTED BLOCKS ==========
		"template_details": {
			Description: "Details about the experiment template used",
			Type:        schema.TypeList,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"identity": {
						Description: "Template identity",
						Type:        schema.TypeString,
						Computed:    true,
					},
					"hub_reference": {
						Description: "Hub reference where template resides",
						Type:        schema.TypeString,
						Computed:    true,
					},
					"reference": {
						Description: "Full template reference",
						Type:        schema.TypeString,
						Computed:    true,
					},
					"revision": {
						Description: "Template revision used",
						Type:        schema.TypeString,
						Computed:    true,
					},
				},
			},
		},
	}
}
