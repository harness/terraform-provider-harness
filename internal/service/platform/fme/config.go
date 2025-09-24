package fme

import (
	"github.com/harness/terraform-provider-harness/internal/service/platform/fme/api"
)

// FMEConfig represents the configuration for the FME service
type FMEConfig struct {
	APIClient *api.Client
}

// NewFMEConfig creates a new FME configuration
func NewFMEConfig(apiKey string) (*FMEConfig, error) {
	client, err := api.New(
		api.APIKey(apiKey),
		api.APIBaseURL("https://api.split.io/internal/api/v2"),
		api.UserAgent("terraform-provider-harness-fme"),
		api.ClientTimeout(300),
	)
	if err != nil {
		return nil, err
	}

	return &FMEConfig{
		APIClient: client,
	}, nil
}