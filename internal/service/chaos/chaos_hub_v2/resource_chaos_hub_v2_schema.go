package chaos_hub_v2

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceChaosHubV2Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		// Required Fields
		"identity": {
			Description:  "Unique identifier for the chaos hub.",
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"name": {
			Description:  "Name of the chaos hub.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		// Optional Fields
		"org_id": {
			Description:  "The ID of the organization.",
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"project_id": {
			Description:  "The ID of the project.",
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"connector_ref": {
			Description:  "Reference to the Git connector (format: scope.connectorId, e.g., org.myconnector or account.myconnector).",
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"repo_branch": {
			Description:  "Git repository branch.",
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"repo_name": {
			Description:  "Name of the Git repository (required for account-level connectors).",
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
		"description": {
			Description: "Description of the chaos hub.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Description: "Tags to associate with the chaos hub.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		// Computed Fields
		"hub_id": {
			Description: "Internal hub ID returned by the API.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"account_id": {
			Description: "Account ID.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"is_default": {
			Description: "Whether this is the default chaos hub.",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"is_removed": {
			Description: "Whether the chaos hub has been removed.",
			Type:        schema.TypeBool,
			Computed:    true,
		},
		"created_at": {
			Description: "Creation timestamp (Unix epoch).",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"updated_at": {
			Description: "Last update timestamp (Unix epoch).",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"created_by": {
			Description: "User who created the chaos hub.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"updated_by": {
			Description: "User who last updated the chaos hub.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"last_synced_at": {
			Description: "Timestamp of the last sync (Unix epoch).",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"repo_url": {
			Description: "Git repository URL.",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"connector_id": {
			Description: "Connector ID (deprecated, use connector_ref).",
			Type:        schema.TypeString,
			Computed:    true,
		},
		"action_template_count": {
			Description: "Number of action templates in the hub.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"experiment_template_count": {
			Description: "Number of experiment templates in the hub.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"fault_template_count": {
			Description: "Number of fault templates in the hub.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
		"probe_template_count": {
			Description: "Number of probe templates in the hub.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
	}
}
