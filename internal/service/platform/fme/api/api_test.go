package api

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name        string
		apiKey      string
		expectError bool
	}{
		{
			name:        "valid api key",
			apiKey:      "test-api-key",
			expectError: false,
		},
		{
			name:        "empty api key",
			apiKey:      "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := New(
				APIKey(tt.apiKey),
				APIBaseURL("https://api.split.io/internal/api/v2"),
				UserAgent("test-client"),
			)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				if client != nil {
					t.Errorf("expected nil client but got %v", client)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if client == nil {
					t.Errorf("expected client but got nil")
				}
				if client.config.APIKey != tt.apiKey {
					t.Errorf("expected API key %s but got %s", tt.apiKey, client.config.APIKey)
				}
			}
		})
	}
}

func TestClientServices(t *testing.T) {
	client, err := New(
		APIKey("test-key"),
		APIBaseURL("https://api.split.io/internal/api/v2"),
	)
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}

	// Test that all services are initialized
	if client.ApiKeys == nil {
		t.Error("ApiKeys service not initialized")
	}
	if client.Environments == nil {
		t.Error("Environments service not initialized")
	}
	if client.FlagSets == nil {
		t.Error("FlagSets service not initialized")
	}
	if client.Splits == nil {
		t.Error("Splits service not initialized")
	}
	if client.Workspaces == nil {
		t.Error("Workspaces service not initialized")
	}
}

func TestConfigOptions(t *testing.T) {
	client, err := New(
		APIKey("test-key"),
		APIBaseURL("https://custom.api.url"),
		UserAgent("custom-agent"),
		ClientTimeout(600),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if client.config.APIBaseURL != "https://custom.api.url" {
		t.Errorf("expected base URL https://custom.api.url but got %s", client.config.APIBaseURL)
	}
	if client.config.UserAgent != "custom-agent" {
		t.Errorf("expected user agent custom-agent but got %s", client.config.UserAgent)
	}
	if client.config.ClientTimeout != 600 {
		t.Errorf("expected timeout 600 but got %d", client.config.ClientTimeout)
	}
}