// examples/image_registry/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/harness/harness-go-sdk/harness/chaos"
	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
)

// Get and display registry details
func getRegistryDetails(registryClient *chaos.ImageRegistryClient, identifiers model.ScopedIdentifiersRequest, infraID *string) (*model.ImageRegistryResponse, error) {
	fmt.Println("\n=== Getting Registry Details ===")

	registry, err := registryClient.Get(
		context.Background(),
		identifiers,
		infraID, // infraID
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get registry: %w", err)
	}

	fmt.Println("Registry details:")
	fmt.Printf("  Registry Server: %s\n", registry.RegistryServer)
	fmt.Printf("  Registry Account: %s\n", registry.RegistryAccount)
	fmt.Printf("  Is Private: %v\n", registry.IsPrivate)
	fmt.Printf("  Is Default: %v\n", registry.IsDefault)
	fmt.Printf("  Is Override Allowed: %v\n", registry.IsOverrideAllowed)
	if registry.SecretName != nil {
		fmt.Printf("  Secret Name: %s\n", *registry.SecretName)
	}
	if registry.CustomImages != nil {
		fmt.Println("  Custom Images:")
		if registry.CustomImages.LogWatcher != nil {
			fmt.Printf("    Log Watcher: %s\n", *registry.CustomImages.LogWatcher)
		}
		if registry.CustomImages.Ddcr != nil {
			fmt.Printf("    DDCR: %s\n", *registry.CustomImages.Ddcr)
		}
		if registry.CustomImages.DdcrLib != nil {
			fmt.Printf("    DDCR Lib: %s\n", *registry.CustomImages.DdcrLib)
		}
		if registry.CustomImages.DdcrFault != nil {
			fmt.Printf("    DDCR Fault: %s\n", *registry.CustomImages.DdcrFault)
		}
	}

	return registry, nil
}

// Example of checking registry override
func exampleCheckOverride(registryClient *chaos.ImageRegistryClient, identifiers model.ScopedIdentifiersRequest, infraID *string) {
	fmt.Println("\n=== Example: Check Registry Override ===")

	// First, get the current registry details to show what we're checking
	fmt.Println("Current registry details being checked:")
	fmt.Printf("  Account: %s\n", identifiers.AccountIdentifier)
	if identifiers.OrgIdentifier != nil {
		fmt.Printf("  Org: %s\n", *identifiers.OrgIdentifier)
	}
	if identifiers.ProjectIdentifier != nil {
		fmt.Printf("  Project: %s\n", *identifiers.ProjectIdentifier)
	}

	// Check override for the registry
	fmt.Println("\nChecking override status...")
	override, err := registryClient.CheckOverride(
		context.Background(),
		identifiers,
		nil, // infraID
	)

	if err != nil {
		log.Printf("Failed to check registry override: %v\n", err)
		return
	}

	fmt.Println("\nOverride check results:")
	if override.OverrideBlockedByScope != "" {
		fmt.Printf("  Override Blocked By Scope: %s\n", override.OverrideBlockedByScope)
	} else {
		fmt.Println("  No override restrictions found")
	}

	if override.ImageRegistry != nil {
		fmt.Println("\n  Registry details from override check:")
		fmt.Printf("    Registry Server: %s\n", override.ImageRegistry.RegistryServer)
		fmt.Printf("    Registry Account: %s\n", override.ImageRegistry.RegistryAccount)
		fmt.Printf("    Is Override Allowed: %v\n", override.ImageRegistry.IsOverrideAllowed)

		if override.ImageRegistry.IsOverrideAllowed {
			fmt.Println("\n  Override is allowed for this registry")
		} else {
			fmt.Println("\n  Override is NOT allowed for this registry")
			if override.OverrideBlockedByScope != "" {
				fmt.Printf("  Reason: %s\n", override.OverrideBlockedByScope)
			}
		}
	} else {
		fmt.Println("\n  No registry configuration found at this scope")
		if override.OverrideBlockedByScope == "NO_OVERRIDE_ALLOWED" {
			fmt.Println("  Override is not allowed at this scope")
		} else if override.OverrideBlockedByScope != "" {
			fmt.Printf("  Override blocked: %s\n", override.OverrideBlockedByScope)
		} else {
			fmt.Println("  No registry configuration found to override")
		}
	}
}

func main() {
	// Get configuration from environment variables
	apiKey := os.Getenv("HARNESS_API_KEY")
	accountID := os.Getenv("HARNESS_ACCOUNT_ID")
	orgID := os.Getenv("HARNESS_ORG_ID")
	projectID := os.Getenv("HARNESS_PROJECT_ID")
	infraID := os.Getenv("HARNESS_INFRA_ID")

	if apiKey == "" || accountID == "" {
		log.Fatal("HARNESS_API_KEY and HARNESS_ACCOUNT_ID environment variables are required")
	}

	// Create a new Chaos client
	cfg := &chaos.Configuration{
		ApiKey:        apiKey,
		AccountId:     accountID,
		BasePath:      "https://app.harness.io/gateway/chaos/manager/api", // or your custom base URL
		UserAgent:     "Harness-Go-SDK-Example/1.0.0",
		DefaultHeader: map[string]string{"X-Api-Key": apiKey},
	}

	client := chaos.NewAPIClient(cfg)

	// Create image registry client
	registryClient := chaos.NewImageRegistryClient(client)

	// Create scoped identifiers
	identifiers := model.ScopedIdentifiersRequest{
		AccountIdentifier: accountID,
	}
	// Only set if they have values
	if orgID != "" {
		orgIDCopy := orgID // Create a new variable to take the address of
		identifiers.OrgIdentifier = &orgIDCopy
	}
	if projectID != "" {
		projectIDCopy := projectID // Create a new variable to take the address of
		identifiers.ProjectIdentifier = &projectIDCopy
	}
	infraIDCopy := infraID // Create a new variable to take the address of

	// Run examples
	exampleCreateWithCustomImages(registryClient, identifiers, &infraIDCopy)
	getRegistryDetails(registryClient, identifiers, &infraIDCopy)
	exampleCheckOverride(registryClient, identifiers, &infraIDCopy) // Check override after creation

	exampleUpdateWithCustomImages(registryClient, identifiers, &infraIDCopy)
	getRegistryDetails(registryClient, identifiers, &infraIDCopy)
	exampleCheckOverride(registryClient, identifiers, &infraIDCopy) // Check override after update
}

// Example of creating a registry with custom images
func exampleCreateWithCustomImages(registryClient *chaos.ImageRegistryClient, identifiers model.ScopedIdentifiersRequest, infraID *string) {
	fmt.Println("\n=== Example: Create Registry with Custom Images ===")

	// Define custom images with updated tags
	customImages := map[string]string{
		"logWatcher": "us-west1-docker.pkg.dev/gar-setup/docker/harness-tf-done/chaos-log-watcher:1.62.0",
		"ddcr":       "us-west1-docker.pkg.dev/gar-setup/docker/harness-tf-done/chaos-ddcr:1.62.0",
		"ddcrLib":    "us-west1-docker.pkg.dev/gar-setup/docker/harness-tf-done/chaos-ddcr-faults:1.62.0",
		"ddcrFault":  "us-west1-docker.pkg.dev/gar-setup/docker/harness-tf-done/chaos-ddcr-faults:1.62.0",
	}

	// Create registry with custom images using map
	fmt.Println("Creating registry with custom images from map...")
	_, err := registryClient.Create(
		context.Background(),
		identifiers,
		"registry-harness-done.io",
		"harness-tf-done",
		true, // isPrivate
		func(req model.ImageRegistryRequest) model.ImageRegistryRequest {
			req = chaos.WithImageRegistryIsDefault(req, true)
			req = chaos.WithImageRegistrySecretName(req, "my-registry-secret")
			req = chaos.WithImageRegistryIsOverrideAllowed(req, true)
			req = chaos.WithImageRegistryCustomImagesFromMap(customImages)(req)
			return req
		},
	)

	if err != nil {
		log.Printf("Warning: Failed to create registry with custom images: %v\n", err)
	} else {
		fmt.Println("Successfully created registry with custom images")
	}
}

// Example of updating a registry with custom images
func exampleUpdateWithCustomImages(registryClient *chaos.ImageRegistryClient, identifiers model.ScopedIdentifiersRequest, infraID *string) {
	fmt.Println("\n=== Example: Update Registry with Custom Images ===")

	// First, check if there's an existing default registry
	fmt.Println("Checking for existing default registry...")
	_, err := registryClient.Get(
		context.Background(),
		identifiers,
		nil, // infraID
	)

	// If there's no error, it means a registry exists and we should set isDefault to false
	// to avoid conflicts with the existing default registry
	isDefault := err != nil // Only set as default if no registry exists

	if isDefault {
		fmt.Println("No existing registry found - this will be set as the default")
	} else {
		fmt.Println("Existing registry found - this will be a non-default registry")
	}

	// Update registry with custom images using individual parameters
	fmt.Println("\nUpdating registry with custom images...")
	result, err := registryClient.Update(
		context.Background(),
		identifiers,
		nil, // infraID for project-level registry
		"registry-harness-done.io",
		"harness-tf-done",
		true, // isPrivate
		func(req model.ImageRegistryRequest) model.ImageRegistryRequest {
			req = chaos.WithImageRegistryIsDefault(req, isDefault)
			req = chaos.WithImageRegistrySecretName(req, "my-updated-secret")
			req = chaos.WithImageRegistryCustomImages(
				"us-west1-docker.pkg.dev/gar-setup/docker/harness-tf-done/chaos-log-watcher:1.62.0",
				"us-west1-docker.pkg.dev/gar-setup/docker/harness-tf-done/chaos-ddcr:1.62.0",
				"us-west1-docker.pkg.dev/gar-setup/docker/harness-tf-done/chaos-ddcr-faults:1.62.0",
				"us-west1-docker.pkg.dev/gar-setup/docker/harness-tf-done/chaos-ddcr-faults:1.62.0",
			)(req)
			return req
		},
	)

	if err != nil {
		log.Printf("Warning: Failed to update registry with custom images: %v\n", err)
	} else {
		fmt.Printf("Successfully updated registry. Result: %s\n", result)
	}

	// Fetch and display the updated registry
	fmt.Println("\nFetching updated registry details...")
	updatedRegistry, err := registryClient.Get(
		context.Background(),
		identifiers,
		nil, // infraID
	)
	if err != nil {
		log.Printf("Failed to get updated registry: %v\n", err)
	} else {
		fmt.Printf("Updated registry details: %+v\n", updatedRegistry)
	}
}
