package chaos_hub

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceChaosHubSync() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for syncing a Harness Chaos Hub",

		CreateContext: resourceChaosHubSyncCreate,
		ReadContext:   resourceChaosHubSyncRead,
		DeleteContext: resourceChaosHubSyncDelete,

		Schema: map[string]*schema.Schema{
			"org_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"hub_id": {
				Description: "The ID of the Chaos Hub to sync",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"last_synced_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last sync",
			},
		},
	}
}

func resourceChaosHubSyncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient
	hubID := d.Get("hub_id").(string)
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	accountID := c.AccountId

	// Parse the hub ID if it's in full path format
	parts := strings.Split(hubID, "/")
	// Convert to model.IdentifiersRequest
	switch len(parts) {
	case 3:
		// Project level hub ID
		orgID = parts[0]
		projectID = parts[1]
		hubID = parts[2]
	case 2:
		// Org level hub ID
		orgID = parts[0]
		hubID = parts[1]
	case 1:
		// Account level hub ID
		hubID = parts[0]
	default:
		return helpers.HandleChaosGraphQLError(fmt.Errorf("invalid hub ID format: expected org_id/project_id/hub_id or project_id/hub_id or hub_id got: %s",
			hubID), d, "sync_chaos_hub")
	}

	// Convert to model.IdentifiersRequest
	modelIdentifiers := model.IdentifiersRequest{
		AccountIdentifier: accountID,
		OrgIdentifier:     orgID,
		ProjectIdentifier: projectID,
	}

	// Create a new Chaos Hub client
	hubClient := chaos.NewChaosHubClient(c)

	// Trigger the sync
	syncID, err := hubClient.Sync(ctx, hubID, modelIdentifiers)
	if err != nil {
		return helpers.HandleChaosGraphQLError(err, d, "sync_chaos_hub")
	}

	d.SetId(hubID)

	log.Printf("[DEBUG] Synced Chaos Hub %s with sync ID: %s", hubID, syncID)
	return resourceChaosHubSyncRead(ctx, d, meta)
}

func resourceChaosHubSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient

	accountID := c.AccountId
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)
	hubID := d.Id()

	// Convert to model.IdentifiersRequest
	modelIdentifiers := model.IdentifiersRequest{
		AccountIdentifier: accountID,
		OrgIdentifier:     orgID,
		ProjectIdentifier: projectID,
	}

	// Get the hub to check sync status
	hubClient := chaos.NewChaosHubClient(c)
	hub, err := hubClient.Get(ctx, modelIdentifiers, hubID)
	if err != nil {
		return helpers.HandleChaosGraphQLReadError(err, d, "get_chaos_hub")
	}

	// Update the status
	d.SetId(hubID)
	d.Set("last_synced_at", hub.LastSyncedAt)
	d.Set("org_id", orgID)
	d.Set("project_id", projectID)

	return nil
}

func resourceChaosHubSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Nothing to do here as the sync is a one-time operation
	d.SetId("")
	return nil
}
