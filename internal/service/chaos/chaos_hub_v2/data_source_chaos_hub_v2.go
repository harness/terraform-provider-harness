package chaos_hub_v2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceChaosHubV2() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Hub.",

		ReadContext: dataSourceChaosHubV2Read,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "The ID of the organization.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "The ID of the project.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"identity": {
				Description:  "Unique identifier of the chaos hub.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"identity", "name"},
			},
			"name": {
				Description:  "Name of the chaos hub.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"identity", "name"},
			},

			// All other fields are computed
			"connector_ref": {
				Description: "Reference to the Git connector.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repo_branch": {
				Description: "Git repository branch.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repo_name": {
				Description: "Name of the Git repository.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repo_url": {
				Description: "Git repository URL.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the chaos hub.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"tags": {
				Description: "Tags associated with the chaos hub.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hub_id": {
				Description: "Internal hub ID.",
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
			"connector_id": {
				Description: "Connector ID (deprecated).",
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
		},
	}
}

func dataSourceChaosHubV2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	var hub *chaos.Chaoshubv2GetHubResponse
	var err error

	// Check if identity is provided
	if identity, ok := d.GetOk("identity"); ok {
		hubIdentity := identity.(string)
		var resp chaos.Chaoshubv2GetHubResponse
		var httpResp *http.Response
		resp, httpResp, err = c.DefaultApi.GetChaosHub(
			ctx,
			c.AccountId,
			orgID,
			projectID,
			hubIdentity,
		)
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}
		hub = &resp
	} else if name, ok := d.GetOk("name"); ok {
		// Search by name using list API
		hubName := name.(string)
		var listResp chaos.Chaoshubv2ListHubResponse
		var httpResp *http.Response
		listResp, httpResp, err = c.DefaultApi.ListChaosHub(
			ctx,
			c.AccountId,
			orgID,
			projectID,
			hubName,
			false,
			0,
			100,
		)
		if err != nil {
			return helpers.HandleChaosApiError(err, d, httpResp)
		}

		// Find the hub with exact name match
		var foundHub *chaos.Chaoshubv2ChaosHubResponse
		for _, h := range listResp.Items {
			if h.Name == hubName {
				foundHub = &h
				break
			}
		}

		if foundHub == nil {
			return diag.Errorf("chaos hub with name '%s' not found", hubName)
		}

		// Convert ChaosHubResponse to GetHubResponse
		hub = &chaos.Chaoshubv2GetHubResponse{
			AccountID:               foundHub.AccountID,
			ActionTemplateCount:     foundHub.ActionTemplateCount,
			ConnectorId:             foundHub.ConnectorId,
			CreatedAt:               foundHub.CreatedAt,
			CreatedBy:               foundHub.CreatedBy,
			Description:             foundHub.Description,
			ExperimentTemplateCount: foundHub.ExperimentTemplateCount,
			FaultTemplateCount:      foundHub.FaultTemplateCount,
			HubId:                   foundHub.HubId,
			Identity:                foundHub.Identity,
			IsDefault:               foundHub.IsDefault,
			IsRemoved:               foundHub.IsRemoved,
			LastSyncedAt:            foundHub.LastSyncedAt,
			Name:                    foundHub.Name,
			OrgID:                   foundHub.OrgID,
			ProbeTemplateCount:      foundHub.ProbeTemplateCount,
			ProjectID:               foundHub.ProjectID,
			RepoBranch:              foundHub.RepoBranch,
			RepoName:                foundHub.RepoName,
			RepoUrl:                 foundHub.RepoUrl,
			Tags:                    foundHub.Tags,
			UpdatedAt:               foundHub.UpdatedAt,
			UpdatedBy:               foundHub.UpdatedBy,
		}
	} else {
		return diag.Errorf("either 'identity' or 'name' must be provided")
	}

	// Set the data source ID to the hub identity
	d.SetId(hub.Identity)

	// Set all the computed fields
	if err := setChaosHubV2Data(d, hub, c.AccountId, orgID, projectID); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set chaos hub data: %v", err))
	}

	return nil
}
