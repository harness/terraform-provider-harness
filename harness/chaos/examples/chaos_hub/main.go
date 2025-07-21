// examples/chaos_hub/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
)

// List all Chaos Hubs
func listChaosHubs(hubClient *chaos.ChaosHubClient, identifiers model.IdentifiersRequest) ([]*model.ChaosHubStatus, error) {
	fmt.Println("\n=== Listing Chaos Hubs ===")

	hubs, err := hubClient.List(
		context.Background(),
		identifiers,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to list chaos hubs: %w", err)
	}

	fmt.Printf("Found %d Chaos Hubs:\n", len(hubs))
	for i, hub := range hubs {
		fmt.Printf("%d. %s (ID: %s)\n", i+1, hub.Name, hub.ID)
		if hub.RepoName != nil {
			fmt.Printf("   Repo Name: %s\n", *hub.RepoName)
		}
		fmt.Printf("   Repo URL: %s\n", hub.RepoURL)
		fmt.Printf("   Branch: %s\n", hub.RepoBranch)
		fmt.Printf("   Is Available: %v\n", hub.IsAvailable)
		fmt.Printf("   Total Faults: %d\n", hub.TotalFaults)
		fmt.Printf("   Total Experiments: %d\n", hub.TotalExperiments)
		fmt.Printf("   Is Default: %v\n", hub.IsDefault)
		if len(hub.Tags) > 0 {
			fmt.Printf("   Tags: %v\n", hub.Tags)
		}
		if hub.Description != nil {
			fmt.Printf("   Description: %s\n", *hub.Description)
		}
		fmt.Println()
	}

	return hubs, nil
}

// Get details of a specific Chaos Hub
func getChaosHubDetails(hubClient *chaos.ChaosHubClient, hubID string, identifiers model.IdentifiersRequest) (*model.ChaosHubStatus, error) {
	fmt.Printf("\n=== Getting Details for Chaos Hub ID: %s ===\n", hubID)

	hub, err := hubClient.Get(
		context.Background(),
		identifiers,
		hubID,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get chaos hub: %w", err)
	}

	fmt.Println("Chaos Hub details:")
	fmt.Printf("  ID: %s\n", hub.ID)
	fmt.Printf("  Name: %s\n", hub.Name)
	if hub.RepoName != nil {
		fmt.Printf("  Repo Name: %s\n", *hub.RepoName)
	}
	fmt.Printf("  Repo URL: %s\n", hub.RepoURL)
	fmt.Printf("  Branch: %s\n", hub.RepoBranch)
	fmt.Printf("  Connector ID: %s\n", hub.ConnectorID)
	fmt.Printf("  Connector Scope: %s\n", hub.ConnectorScope)
	fmt.Printf("  Auth Type: %s\n", hub.AuthType)
	fmt.Printf("  Is Available: %v\n", hub.IsAvailable)
	fmt.Printf("  Total Faults: %d\n", hub.TotalFaults)
	fmt.Printf("  Total Experiments: %d\n", hub.TotalExperiments)
	fmt.Printf("  Is Default: %v\n", hub.IsDefault)
	if len(hub.Tags) > 0 {
		fmt.Printf("  Tags: %v\n", hub.Tags)
	}
	if hub.Description != nil {
		fmt.Printf("  Description: %s\n", *hub.Description)
	}
	if hub.CreatedBy != nil {
		fmt.Printf("  Created By: %s (%s)\n", hub.CreatedBy.Username, hub.CreatedBy.Email)
	}
	if hub.UpdatedBy != nil {
		fmt.Printf("  Updated By: %s (%s)\n", hub.UpdatedBy.Username, hub.UpdatedBy.Email)
	}
	fmt.Printf("  Created At: %s\n", hub.CreatedAt)
	fmt.Printf("  Updated At: %s\n", hub.UpdatedAt)
	fmt.Printf("  Last Synced At: %s\n", hub.LastSyncedAt)

	return hub, nil
}

// Create a new Chaos Hub
func createChaosHub(hubClient *chaos.ChaosHubClient, hubName, repoBranch, connectorID string, identifiers model.IdentifiersRequest) (*model.ChaosHub, error) {
	fmt.Printf("\n=== Creating New Chaos Hub: %s ===\n", hubName)

	hub, err := hubClient.Create(
		context.Background(),
		hubName,
		repoBranch,
		connectorID,
		identifiers,
		chaos.WithChaosHubDescription("Created via Harness Go SDK"),
		chaos.WithChaosHubTags([]string{"sdk-created", "example"}),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create chaos hub: %w", err)
	}

	fmt.Printf("Successfully created Chaos Hub with ID: %s\n", hub.ID)
	return hub, nil
}

// Update a Chaos Hub
func updateChaosHub(hubClient *chaos.ChaosHubClient, hubID, hubName, repoBranch, connectorID string, identifiers model.IdentifiersRequest, opts ...func(*model.ChaosHubRequest) *model.ChaosHubRequest) (*model.ChaosHub, error) {
	fmt.Printf("\n=== Updating Chaos Hub: %s ===\n", hubID)

	hub, err := hubClient.Update(
		context.Background(),
		hubID,
		hubName,
		repoBranch,
		connectorID,
		identifiers,
		opts...,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update chaos hub: %w", err)
	}

	fmt.Printf("Successfully updated Chaos Hub with ID: %s\n", hub.ID)
	return hub, nil
}

// Sync a Chaos Hub
func syncChaosHub(hubClient *chaos.ChaosHubClient, hubID string, identifiers model.IdentifiersRequest) error {
	fmt.Printf("\n=== Syncing Chaos Hub: %s ===\n", hubID)

	// First, get the hub details to check its status
	hub, err := getChaosHubDetails(hubClient, hubID, identifiers)
	if err != nil {
		return fmt.Errorf("failed to get hub details before sync: %w", err)
	}

	// Check if the hub is available
	if !hub.IsAvailable {
		return fmt.Errorf("cannot sync: hub is not available. Please check the connector configuration")
	}

	// Check if the connector is properly configured
	if hub.ConnectorID == "" {
		return fmt.Errorf("cannot sync: no connector is associated with this hub")
	}

	syncID, err := hubClient.Sync(
		context.Background(),
		hubID,
		identifiers,
	)

	if err != nil {
		// Provide more helpful error messages for common issues
		if err.Error() == "graphql error: error while decrypting connector: git connector authentication type is not supported" {
			return fmt.Errorf("failed to sync: the Git connector's authentication type is not supported. " +
				"Please check the connector configuration in the Harness UI")
		}
		return fmt.Errorf("failed to sync chaos hub: %w", err)
	}

	fmt.Printf("Started sync with ID: %s\n", syncID)
	return nil
}

// Delete a Chaos Hub
func deleteChaosHub(hubClient *chaos.ChaosHubClient, hubID string, identifiers model.IdentifiersRequest) error {
	fmt.Printf("\n=== Deleting Chaos Hub: %s ===\n", hubID)

	deleted, err := hubClient.Delete(
		context.Background(),
		hubID,
		identifiers,
	)

	if err != nil {
		return fmt.Errorf("failed to delete chaos hub: %w", err)
	}

	if deleted {
		fmt.Println("Successfully deleted Chaos Hub")
	} else {
		fmt.Println("Failed to delete Chaos Hub")
	}

	return nil
}

func main() {
	// Get configuration from environment variables
	apiKey := os.Getenv("HARNESS_API_KEY")
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	orgID := os.Getenv("HARNESS_ORG_ID")
	projectID := os.Getenv("HARNESS_PROJECT_ID")
	connectorID := os.Getenv("HARNESS_CONNECTOR_ID")

	if apiKey == "" || accountID == "" {
		log.Fatal("HARNESS_API_KEY and HARNESS_ACCOUNT_ID environment variables are required")
	}

	// Create a new Chaos client
	cfg := &chaos.Configuration{
		ApiKey:    apiKey,
		AccountId: accountID,
		BasePath:  "https://app.harness.io/gateway/chaos/manager/api",
		UserAgent: "Harness-Go-SDK-Chaos-Hub-Example/1.0.0",
	}

	client := chaos.NewAPIClient(cfg)

	// Create Chaos Hub client
	hubClient := chaos.NewChaosHubClient(client)

	// Create identifiers
	identifiers := model.IdentifiersRequest{
		AccountIdentifier: accountID,
	}

	// Set organization and project IDs if provided
	if orgID != "" {
		identifiers.OrgIdentifier = orgID
	}
	if projectID != "" {
		identifiers.ProjectIdentifier = projectID
	}

	// List all chaos hubs first
	hubs, err := listChaosHubs(hubClient, identifiers)
	if err != nil {
		log.Printf("Warning: Failed to list chaos hubs: %v", err)
	} else if len(hubs) > 0 {
		fmt.Printf("\n=== Found %d existing Chaos Hubs ===\n", len(hubs))
	}

	// Test create, update, and delete operations
	testHubName := fmt.Sprintf("test-hub-%d", time.Now().Unix())
	fmt.Printf("\n=== Testing Create/Update/Delete Operations ===\n")

	// 1. Create a new hub
	fmt.Printf("\n--- Creating new Chaos Hub: %s ---\n", testHubName)
	newHub, err := createChaosHub(
		hubClient,
		testHubName,
		"main",
		connectorID,
		identifiers,
	)
	if err != nil {
		log.Printf("Failed to create hub: %v", err)
		return
	}

	// Small delay to ensure the hub is fully created
	time.Sleep(2 * time.Second)

	// 2. First update - basic fields
	updatedName := testHubName + "-updated"
	fmt.Printf("\n--- Updating Chaos Hub (Basic Fields): %s ---\n", newHub.ID)
	updatedHub, err := updateChaosHub(
		hubClient,
		newHub.ID,
		updatedName,
		"main",
		connectorID, // Using same connector for first update
		identifiers,
		chaos.WithChaosHubDescription("Updated via SDK test - basic fields"),
		chaos.WithChaosHubTags([]string{"sdk-test", "temporary", "first-update"}),
	)
	if err != nil {
		log.Printf("Failed to update hub (basic fields): %v", err)
		return
	}
	fmt.Printf("Successfully updated hub to name: %s\n", updatedHub.Name)

	// 3. Second update - change connector (if a second connector is provided)
	if updatedConnectorID := os.Getenv("HARNESS_CONNECTOR_ID_2"); updatedConnectorID != "" {
		fmt.Printf("\n--- Updating Chaos Hub (Connector): %s ---\n", updatedHub.ID)
		time.Sleep(2 * time.Second) // Small delay between updates

		doubleUpdatedHub, err := updateChaosHub(
			hubClient,
			updatedHub.ID,
			updatedName, // Keep the same name
			"main",
			updatedConnectorID, // Use the second connector
			identifiers,
			chaos.WithChaosHubDescription("Updated via SDK test - connector changed"),
			chaos.WithChaosHubTags([]string{"sdk-test", "temporary", "connector-updated"}),
		)
		if err != nil {
			log.Printf("Warning: Failed to update hub connector: %v", err)
		} else {
			fmt.Printf("Successfully updated hub connector to: %s\n", updatedConnectorID)
			updatedHub = doubleUpdatedHub
		}
	} else {
		fmt.Println("\nSkipping connector update: HARNESS_CONNECTOR_ID_2 not set")
	}

	// 4. Get the final hub details
	fmt.Printf("\n--- Getting final hub details ---\n")
	hubDetails, err := getChaosHubDetails(hubClient, newHub.ID, identifiers)
	if err != nil {
		log.Printf("Failed to get updated hub details: %v", err)
	} else {
		fmt.Printf("Updated Hub Details:\n")
		fmt.Printf("  Name: %s\n", hubDetails.Name)
		if hubDetails.Description != nil {
			fmt.Printf("  Description: %s\n", *hubDetails.Description)
		}
		fmt.Printf("  Tags: %v\n", hubDetails.Tags)
		fmt.Printf("  Connector ID: %s\n", hubDetails.ConnectorID)
	}

	// 5. Test sync operation
	fmt.Printf("\n--- Testing Sync Operation ---\n")
	err = syncChaosHub(hubClient, newHub.ID, identifiers)
	if err != nil {
		log.Printf("Warning: Failed to sync hub: %v", err)
	} else {
		fmt.Println("Successfully triggered sync operation")
		
		// Wait a moment for sync to potentially complete
		time.Sleep(3 * time.Second)
		
		// Get updated details to check sync status
		syncedHub, err := getChaosHubDetails(hubClient, newHub.ID, identifiers)
		if err != nil {
			log.Printf("Warning: Failed to get hub details after sync: %v", err)
		} else {
			fmt.Printf("Last Synced At: %s\n", syncedHub.LastSyncedAt)
			if syncedHub.IsAvailable {
				fmt.Printf("Hub is available with %d experiments and %d faults\n", 
					syncedHub.TotalExperiments, syncedHub.TotalFaults)
			} else {
				fmt.Println("Note: Hub is not currently available")
			}
		}
	}

	// 6. Clean up - delete the test hub
	fmt.Printf("\n--- Cleaning up: Deleting test hub ---\n")
	err = deleteChaosHub(hubClient, newHub.ID, identifiers)
	if err != nil {
		log.Printf("Failed to delete test hub: %v", err)
		return // Exit if we can't delete to avoid leaving test artifacts
	}

	fmt.Printf("Successfully deleted test hub: %s\n", newHub.ID)

	// 7. Verify deletion
	fmt.Printf("\n--- Verifying deletion ---\n")
	_, err = getChaosHubDetails(hubClient, newHub.ID, identifiers)
	if err != nil {
		fmt.Printf("Verification: Hub %s no longer exists (expected)\n", newHub.ID)
	} else {
		log.Printf("Warning: Hub %s still exists after deletion", newHub.ID)
	}

	// 8. Final list of hubs
	hubs, err = listChaosHubs(hubClient, identifiers)
	if err != nil {
		log.Printf("Warning: Failed to list chaos hubs after cleanup: %v", err)
	} else {
		fmt.Printf("\n=== Found %d Chaos Hubs after cleanup ===\n", len(hubs))
	}
}
