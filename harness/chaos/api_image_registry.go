package chaos

import (
	"context"
	"fmt"

	"github.com/harness/harness-go-sdk/harness/chaos/graphql/model"
)

// ImageRegistryClient provides methods to interact with the Image Registry API
type ImageRegistryClient struct {
	client *APIClient
}

// NewImageRegistryClient creates a new Image Registry client
func NewImageRegistryClient(client *APIClient) *ImageRegistryClient {
	return &ImageRegistryClient{client: client}
}

// Get retrieves an image registry configuration
func (c *ImageRegistryClient) Get(
	ctx context.Context,
	identifiers model.ScopedIdentifiersRequest,
	infraID *string,
) (*model.ImageRegistryResponse, error) {
	query := `
        query GetImageRegistry($identifiers: ScopedIdentifiersRequest!, $infraID: String) {
            getImageRegistry(identifiers: $identifiers, infraID: $infraID) {
                identifier {
                    accountIdentifier
                    orgIdentifier
                    projectIdentifier
                }
                infraID
                registryServer
                registryAccount
                isOverrideAllowed
                isPrivate
                secretName
                isDefault
                useCustomImages
                customImages {
                    logWatcher
                    ddcr
                    ddcrLib
                    ddcrFault
                }
                createdBy {
                    userID
                    username
                    email
                }
                updatedBy {
                    userID
                    username
                    email
                }
                createdAt
                updatedAt
            }
        }`

	variables := map[string]interface{}{
		"identifiers": identifiers,
	}

	if infraID != nil {
		variables["infraID"] = *infraID
	}

	var response struct {
		GetImageRegistry *model.ImageRegistryResponse `json:"getImageRegistry"`
	}

	if err := c.client.ExecuteGraphQL(ctx, query, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to get image registry: %w", err)
	}

	return response.GetImageRegistry, nil
}

// Create creates a new image registry configuration
func (c *ImageRegistryClient) Create(
	ctx context.Context,
	identifiers model.ScopedIdentifiersRequest,
	registryServer string,
	registryAccount string,
	isPrivate bool,
	opts ...func(model.ImageRegistryRequest) model.ImageRegistryRequest,
) (*model.ImageRegistryResponse, error) {
	// Start with the basic request
	req := CreateImageRegistryRequest(registryServer, registryAccount, isPrivate)

	// Apply any additional options
	for _, opt := range opts {
		req = opt(req)
	}

	query := `
        mutation CreateImageRegistry($identifiers: ScopedIdentifiersRequest!, $input: ImageRegistryRequest!) {
            createImageRegistry(identifiers: $identifiers, request: $input) {
                identifier {
                    accountIdentifier
                    orgIdentifier
                    projectIdentifier
                }
                infraID
                registryServer
                registryAccount
                isOverrideAllowed
                isPrivate
                secretName
                isDefault
                useCustomImages
                customImages {
                    logWatcher
                    ddcr
                    ddcrLib
                    ddcrFault
                }
            }
        }`

	var response struct {
		CreateImageRegistry *model.ImageRegistryResponse `json:"createImageRegistry"`
	}

	// Ensure we're passing the variables correctly
	variables := map[string]interface{}{
		"identifiers": identifiers,
		"input":       req,
	}

	if err := c.client.ExecuteGraphQL(ctx, query, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to create image registry: %w", err)
	}

	return response.CreateImageRegistry, nil
}

// Update updates an existing image registry configuration
func (c *ImageRegistryClient) Update(
	ctx context.Context,
	identifiers model.ScopedIdentifiersRequest,
	infraID *string,
	registryServer string,
	registryAccount string,
	isPrivate bool,
	opts ...func(model.ImageRegistryRequest) model.ImageRegistryRequest,
) (string, error) {
	// Start with the basic request with all required fields
	req := model.ImageRegistryRequest{
		RegistryServer:    registryServer,
		RegistryAccount:   registryAccount,
		IsPrivate:         isPrivate,
		IsDefault:         infraID == nil, // Default to true for project-level, false for infra-level
		IsOverrideAllowed: true,           // Default to true
		UseCustomImages:   false,          // Default to false
		SecretName:        new(string),    // Initialize as empty string
	}

	// Set empty string for secret name if not provided
	*req.SecretName = ""

	// Set infraID if provided
	if infraID != nil && *infraID != "" {
		req.InfraID = infraID
	}

	// Apply any additional options (these can override the defaults above)
	for _, opt := range opts {
		req = opt(req)
	}

	query := `
        mutation UpdateImageRegistry($identifiers: ScopedIdentifiersRequest!, $request: ImageRegistryRequest!) {
            updateImageRegistry(identifiers: $identifiers, request: $request)
        }`

	var response struct {
		UpdateImageRegistry string `json:"updateImageRegistry"`
	}

	// Prepare variables according to the working request
	variables := map[string]interface{}{
		"identifiers": map[string]interface{}{
			"accountIdentifier": identifiers.AccountIdentifier,
			"orgIdentifier":     "", // Initialize as empty string
			"projectIdentifier": "", // Initialize as empty string
		},
		"request": map[string]interface{}{
			"registryServer":    req.RegistryServer,
			"registryAccount":   req.RegistryAccount,
			"isPrivate":         req.IsPrivate,
			"isDefault":         req.IsDefault,
			"isOverrideAllowed": req.IsOverrideAllowed,
			"useCustomImages":   req.UseCustomImages,
			"secretName":        "", // Initialize as empty string
		},
	}

	// Set optional fields if they have values
	if identifiers.OrgIdentifier != nil && *identifiers.OrgIdentifier != "" {
		variables["identifiers"].(map[string]interface{})["orgIdentifier"] = *identifiers.OrgIdentifier
	}

	if identifiers.ProjectIdentifier != nil && *identifiers.ProjectIdentifier != "" {
		variables["identifiers"].(map[string]interface{})["projectIdentifier"] = *identifiers.ProjectIdentifier
	}

	if req.InfraID != nil && *req.InfraID != "" {
		variables["request"].(map[string]interface{})["infraID"] = *req.InfraID
	}

	if req.SecretName != nil {
		variables["request"].(map[string]interface{})["secretName"] = *req.SecretName
	}

	// Always include customImages, even if empty
	if req.CustomImages != nil {
		variables["request"].(map[string]interface{})["customImages"] = req.CustomImages
	} else {
		variables["request"].(map[string]interface{})["customImages"] = nil
	}

	if err := c.client.ExecuteGraphQL(ctx, query, variables, &response); err != nil {
		return "", fmt.Errorf("failed to update image registry: %w", err)
	}

	return response.UpdateImageRegistry, nil
}

// CheckOverride checks if an image registry override is allowed
func (c *ImageRegistryClient) CheckOverride(
	ctx context.Context,
	identifiers model.ScopedIdentifiersRequest,
	infraID *string,
) (*model.CheckImageRegistryOverrideResponse, error) {
	query := `
        query CheckImageRegistryOverride($identifiers: ScopedIdentifiersRequest!, $infraID: String) {
            checkImageRegistryOverride(identifiers: $identifiers, infraID: $infraID) {
                OverrideBlockedByScope
                ImageRegistry {
                    identifier {
                        accountIdentifier
                        orgIdentifier
                        projectIdentifier
                    }
                    infraID
                    registryServer
                    registryAccount
                    isOverrideAllowed
                    isPrivate
                    secretName
                    isDefault
                    useCustomImages
                    customImages {
                        logWatcher
                        ddcr
                        ddcrLib
                        ddcrFault
                    }
                }
            }
        }`

	variables := map[string]interface{}{
		"identifiers": identifiers,
	}

	if infraID != nil {
		variables["infraID"] = *infraID
	}

	var response struct {
		CheckImageRegistryOverride *model.CheckImageRegistryOverrideResponse `json:"checkImageRegistryOverride"`
	}

	if err := c.client.ExecuteGraphQL(ctx, query, variables, &response); err != nil {
		return nil, fmt.Errorf("failed to check image registry override: %w", err)
	}

	if response.CheckImageRegistryOverride == nil {
		// Return a default response instead of an error
		return &model.CheckImageRegistryOverrideResponse{
			OverrideBlockedByScope: "NO_OVERRIDE_ALLOWED",
			ImageRegistry:          nil,
		}, nil
	}

	return response.CheckImageRegistryOverride, nil
}
