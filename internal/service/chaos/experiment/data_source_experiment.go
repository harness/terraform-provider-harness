package experiment

import (
	"context"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceChaosExperiment() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for looking up chaos experiments",

		ReadContext: dataSourceExperimentRead,

		Schema: map[string]*schema.Schema{
			// ========== REQUIRED SCOPE ==========
			"org_id": {
				Description: "Organization identifier",
				Type:        schema.TypeString,
				Required:    true,
			},
			"project_id": {
				Description: "Project identifier",
				Type:        schema.TypeString,
				Required:    true,
			},

			// ========== LOOKUP FIELDS (one required) ==========
			"identity": {
				Description:  "Experiment identity to lookup",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"identity", "name"},
			},
			"name": {
				Description:  "Experiment name to lookup",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"identity", "name"},
			},

			// ========== ALL OUTPUT FIELDS (same as resource) ==========
			"template_identity": {
				Description: "Identity of the experiment template",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"hub_identity": {
				Description: "Identity of the hub where template resides",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"infra_ref": {
				Description: "Infrastructure reference",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the chaos experiment",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"tags": {
				Description: "Tags to categorize the experiment",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"revision": {
				Description: "Template revision used",
				Type:        schema.TypeString,
				Computed:    true,
			},
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
				Description: "Infrastructure type",
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
		},
	}
}

func dataSourceExperimentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetChaosClientWithContext(ctx)

	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

	// Option 1: Lookup by identity (direct GET)
	if identity, ok := d.GetOk("identity"); ok {
		log.Printf("[DEBUG] Looking up experiment by identity: %s", identity.(string))

		experiment, httpResp, err := c.DefaultApi.GetChaosV2Experiment(
			ctx,
			c.AccountId,
			orgID,
			projectID,
			identity.(string),
		)

		if err != nil {
			return helpers.HandleChaosReadApiError(err, d, httpResp)
		}

		// Set ID
		d.SetId(fmt.Sprintf("%s/%s/%s", orgID, projectID, experiment.Identity))

		// Set all fields
		return setExperimentFields(d, &experiment)
	}

	// Option 2: Lookup by name (LIST with filter)
	if name, ok := d.GetOk("name"); ok {
		log.Printf("[DEBUG] Looking up experiment by name: %s", name.(string))

		opts := &chaos.DefaultApiListChaosV2ExperimentOpts{
			ExperimentName: optional.NewString(name.(string)),
		}

		resp, httpResp, err := c.DefaultApi.ListChaosV2Experiment(
			ctx,
			c.AccountId,
			orgID,
			projectID,
			0,   // page
			100, // limit
			opts,
		)

		if err != nil {
			return helpers.HandleChaosReadApiError(err, d, httpResp)
		}

		if len(resp.Data) == 0 {
			return diag.Errorf("experiment with name '%s' not found in org '%s', project '%s'", name, orgID, projectID)
		}

		if len(resp.Data) > 1 {
			return diag.Errorf("multiple experiments found with name '%s', please use 'identity' instead", name)
		}

		// Get full details via GET API
		experiment := resp.Data[0]
		// TypesExperimentV2 uses ExperimentID field
		experimentIdentity := experiment.ExperimentID

		log.Printf("[DEBUG] Found experiment with identity: %s, fetching full details", experimentIdentity)

		fullExperiment, httpResp, err := c.DefaultApi.GetChaosV2Experiment(
			ctx,
			c.AccountId,
			orgID,
			projectID,
			experimentIdentity,
		)

		if err != nil {
			return helpers.HandleChaosReadApiError(err, d, httpResp)
		}

		d.SetId(fmt.Sprintf("%s/%s/%s", orgID, projectID, fullExperiment.Identity))
		return setExperimentFields(d, &fullExperiment)
	}

	return diag.Errorf("either 'identity' or 'name' must be specified")
}

// setExperimentFields sets all experiment fields in the data source
func setExperimentFields(d *schema.ResourceData, experiment *chaos.ChaosExperimentChaosExperimentRequest) diag.Diagnostics {
	// Scope
	d.Set("org_id", experiment.OrgID)
	d.Set("project_id", experiment.ProjectID)

	// Basic info
	d.Set("name", experiment.Name)
	d.Set("identity", experiment.Identity)
	d.Set("description", experiment.Description)
	if len(experiment.Tags) > 0 {
		d.Set("tags", flattenTags(experiment.Tags))
	}

	// Infrastructure
	d.Set("infra_id", experiment.InfraID)
	d.Set("infra_ref", experiment.InfraID) // Use infra_id as infra_ref
	if experiment.InfraType != nil {
		d.Set("infra_type", string(*experiment.InfraType))
	}

	// Computed fields
	d.Set("experiment_id", experiment.ExperimentID)
	d.Set("is_custom_experiment", experiment.IsCustomExperiment)

	if experiment.ExperimentType != nil {
		d.Set("experiment_type", string(*experiment.ExperimentType))
	}

	if len(experiment.FaultIDs) > 0 {
		d.Set("fault_ids", experiment.FaultIDs)
	}

	d.Set("cron_syntax", experiment.CronSyntax)
	d.Set("is_cron_enabled", experiment.IsCronEnabled)
	d.Set("is_single_run_cron_enabled", experiment.IsSingleRunCronEnabled)
	d.Set("last_executed_at", experiment.LastExecutedAt)
	d.Set("total_experiment_runs", experiment.TotalExperimentRuns)
	d.Set("target_network_map_id", experiment.TargetNetworkMapID)
	d.Set("created_at", experiment.CreatedAt)
	d.Set("created_by", experiment.CreatedBy)
	d.Set("updated_at", experiment.UpdatedAt)
	d.Set("updated_by", experiment.UpdatedBy)

	// Template details
	if experiment.TemplateDetails != nil {
		templateDetails := []map[string]interface{}{
			{
				"identity":      experiment.TemplateDetails.Identity,
				"hub_reference": experiment.TemplateDetails.HubReference,
				"reference":     experiment.TemplateDetails.Reference,
				"revision":      experiment.TemplateDetails.Revision,
			},
		}
		d.Set("template_details", templateDetails)

		// Also set input fields for reference
		d.Set("hub_identity", experiment.TemplateDetails.HubReference)
		d.Set("template_identity", experiment.TemplateDetails.Identity)
		if experiment.TemplateDetails.Revision != "" {
			d.Set("revision", experiment.TemplateDetails.Revision)
		}
	}

	log.Printf("[DEBUG] Successfully set all experiment fields")
	return nil
}
