package chaos

import "github.com/harness/harness-go-sdk/harness/chaos/graphql/model"

// CreateImageRegistryRequest creates a new ImageRegistryRequest with the required fields
func CreateImageRegistryRequest(registryServer, registryAccount string, isPrivate bool) model.ImageRegistryRequest {
	return model.ImageRegistryRequest{
		RegistryServer:  registryServer,
		RegistryAccount: registryAccount,
		IsPrivate:       isPrivate,
	}
}

// WithImageRegistryInfraID sets the infrastructure ID in the request
func WithImageRegistryInfraID(req model.ImageRegistryRequest, infraID string) model.ImageRegistryRequest {
	req.InfraID = &infraID
	return req
}

// WithImageRegistrySecretName sets the secret name in the request
func WithImageRegistrySecretName(req model.ImageRegistryRequest, secretName string) model.ImageRegistryRequest {
	req.SecretName = &secretName
	return req
}

// WithImageRegistryIsDefault sets whether this is the default registry
func WithImageRegistryIsDefault(req model.ImageRegistryRequest, isDefault bool) model.ImageRegistryRequest {
	req.IsDefault = isDefault
	return req
}

// WithImageRegistryIsOverrideAllowed sets whether override is allowed
func WithImageRegistryIsOverrideAllowed(req model.ImageRegistryRequest, isOverrideAllowed bool) model.ImageRegistryRequest {
	req.IsOverrideAllowed = isOverrideAllowed
	return req
}

// WithImageRegistryCustomImages adds custom images to the request using individual parameters
// This is the recommended way to set custom images
func WithImageRegistryCustomImages(logWatcher, ddcr, ddcrLib, ddcrFault string) func(model.ImageRegistryRequest) model.ImageRegistryRequest {
	return func(req model.ImageRegistryRequest) model.ImageRegistryRequest {
		req.CustomImages = &model.CustomImages{
			LogWatcher: stringPtr(logWatcher),
			Ddcr:       stringPtr(ddcr),
			DdcrLib:    stringPtr(ddcrLib),
			DdcrFault:  stringPtr(ddcrFault),
		}
		// Enable useCustomImages when custom images are provided
		req.UseCustomImages = true
		return req
	}
}

// WithImageRegistryCustomImagesFromMap sets the custom images from a map
// The map should have keys: logWatcher, ddcr, ddcrLib, ddcrFault
func WithImageRegistryCustomImagesFromMap(images map[string]string) func(model.ImageRegistryRequest) model.ImageRegistryRequest {
	return func(req model.ImageRegistryRequest) model.ImageRegistryRequest {
		customImages := &model.CustomImages{
			LogWatcher: stringPtr(images["logWatcher"]),
			Ddcr:       stringPtr(images["ddcr"]),
			DdcrLib:    stringPtr(images["ddcrLib"]),
			DdcrFault:  stringPtr(images["ddcrFault"]),
		}
		req.CustomImages = customImages
		req.UseCustomImages = true
		return req
	}
}

// WithImageRegistryUseCustomImages enables or disables the use of custom images
// Set to true to use custom images, false to use default images
func WithImageRegistryUseCustomImages(useCustomImages bool) func(model.ImageRegistryRequest) model.ImageRegistryRequest {
	return func(req model.ImageRegistryRequest) model.ImageRegistryRequest {
		req.UseCustomImages = useCustomImages
		return req
	}
}

// stringPtr returns a pointer to the string value
// Returns nil for empty strings to match the API's behavior with optional fields
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
