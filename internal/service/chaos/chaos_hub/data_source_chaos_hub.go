package chaos_hub

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceChaosHub() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving a Harness Chaos Hub",
		ReadContext: dataSourceChaosHubRead,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Description: "The organization ID of the chaos hub",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "The project ID of the chaos hub",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "Name of the chaos hub",
				Type:        schema.TypeString,
				Required:    true,
			},
			"connector_id": {
				Description: "ID of the Git connector",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repo_branch": {
				Description: "Git repository branch",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repo_name": {
				Description: "Name of the Git repository",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "Description of the chaos hub",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"tags": {
				Description: "Tags associated with the chaos hub",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_default": {
				Description: "Whether this is the default chaos hub",
				Type:        schema.TypeBool,
				Computed:    true,
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

func dataSourceChaosHubRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient

	// Get all hubs and find the one with the matching name
	identifiers := model.IdentifiersRequest{
		AccountIdentifier: c.AccountId,
	}

	if v, ok := d.GetOk("org_id"); ok {
		orgID := v.(string)
		identifiers.OrgIdentifier = orgID
	}

	if v, ok := d.GetOk("project_id"); ok {
		projectID := v.(string)
		identifiers.ProjectIdentifier = projectID
	}

	hubClient := chaos.NewChaosHubClient(c)
	hubs, err := hubClient.List(ctx, identifiers)
	if err != nil {
		return diag.Errorf("failed to list chaos hubs: %v", err)
	}

	targetName := d.Get("name").(string)
	var foundHub *model.ChaosHubStatus

	for _, hub := range hubs {
		if hub.Name == targetName {
			foundHub = hub
			break
		}
	}

	if foundHub == nil {
		return diag.Errorf("chaos hub with name '%s' not found", targetName)
	}

	// Convert model.IdentifiersRequest to ScopedIdentifiersRequest for generateID
	scopedIdentifiers := ScopedIdentifiersRequest{
		AccountIdentifier: identifiers.AccountIdentifier,
	}
	if identifiers.OrgIdentifier != "" {
		scopedIdentifiers.OrgIdentifier = &identifiers.OrgIdentifier
	}
	if identifiers.ProjectIdentifier != "" {
		scopedIdentifiers.ProjectIdentifier = &identifiers.ProjectIdentifier
	}

	return setChaosHubData(d, foundHub, scopedIdentifiers)
}

func setChaosHubData(d *schema.ResourceData, hub *model.ChaosHubStatus, identifiers ScopedIdentifiersRequest) diag.Diagnostics {
	d.SetId(generateID(identifiers, hub.Identity))
	d.Set("name", hub.Name)
	d.Set("connector_id", hub.ConnectorID)
	d.Set("repo_branch", hub.RepoBranch)
	d.Set("is_default", hub.IsDefault)
	d.Set("created_at", hub.CreatedAt)
	d.Set("updated_at", hub.UpdatedAt)
	d.Set("last_synced_at", hub.LastSyncedAt)
	d.Set("is_available", hub.IsAvailable)
	d.Set("total_experiments", hub.TotalExperiments)
	d.Set("total_faults", hub.TotalFaults)

	if hub.RepoName != nil {
		d.Set("repo_name", *hub.RepoName)
	}
	if hub.Description != nil {
		d.Set("description", *hub.Description)
	}
	if len(hub.Tags) > 0 {
		d.Set("tags", hub.Tags)
	}

	return nil
}
