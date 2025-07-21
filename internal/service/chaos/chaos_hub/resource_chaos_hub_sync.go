package chaos_hub

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
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
	if len(parts) == 4 {
		// If hub_id is in full path format, use its components
		accountID = parts[0]
		orgID = parts[1]
		projectID = parts[2]
		hubID = parts[3]
	}

	// Convert to model.IdentifiersRequest
	if len(parts) == 4 {
		// If hub_id is in full path format, use its components
		accountID = parts[0]
		orgID = parts[1]
		projectID = parts[2]
		hubID = parts[3]
	} else {
		// Otherwise use the provided org_id/project_id and account_id from provider
		accountID = c.AccountId
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
		return diag.Errorf("failed to sync chaos hub: %v", err)
	}

	// Set the ID as a combination of hub ID and timestamp to force recreation on each sync
	d.SetId(fmt.Sprintf("%s-%d", hubID, time.Now().Unix()))
	d.Set("last_synced_at", time.Now().Format(time.RFC3339))

	log.Printf("[DEBUG] Synced Chaos Hub %s with sync ID: %s", hubID, syncID)
	return resourceChaosHubSyncRead(ctx, d, meta)
}

func resourceChaosHubSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*internal.Session).ChaosClient
	hubID := d.Get("hub_id").(string)
	accountID := c.AccountId
	orgID := d.Get("org_id").(string)
	projectID := d.Get("project_id").(string)

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
		return diag.Errorf("failed to get chaos hub: %v", err)
	}

	// Update the status
	d.Set("last_synced_at", hub.LastSyncedAt)

	return nil
}

func resourceChaosHubSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Nothing to do here as the sync is a one-time operation
	d.SetId("")
	return nil
}
