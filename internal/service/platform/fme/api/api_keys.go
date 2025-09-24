package api

import (
	"fmt"
)

// KeysService handles communication with the API keys related
// methods of the Split.io APIv2.
type KeysService service

// Key represents an API key in Split.io
type Key struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
	Type *string `json:"type"`
	Key  *string `json:"key"`
}

// KeyWorkspaceRequest represents the workspace reference in an API key request
type KeyWorkspaceRequest struct {
	Type *string `json:"type,omitempty"`
	ID   *string `json:"id,omitempty"`
}

// KeyEnvironmentRequest represents an environment reference in an API key request
type KeyEnvironmentRequest struct {
	Type *string `json:"type,omitempty"`
	ID   *string `json:"id,omitempty"`
}

// KeyRequest represents a request to create/update an API key
type KeyRequest struct {
	Name         *string                   `json:"name,omitempty"`
	APIKeyType   *string                   `json:"apiKeyType,omitempty"`
	Environments []*KeyEnvironmentRequest  `json:"environments,omitempty"`
	Workspace    *KeyWorkspaceRequest      `json:"workspace,omitempty"`
}

// Create creates a new API key
func (k *KeysService) Create(opts *KeyRequest) (*Key, error) {
	var result Key
	url := "https://api.split.io/internal/api/v2/apiKeys"
	err := k.client.post(url, opts, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes an API key
func (k *KeysService) Delete(keyID string) error {
	url := fmt.Sprintf("https://api.split.io/internal/api/v2/apiKeys/%s", keyID)
	err := k.client.delete(url)
	if err != nil {
		return err
	}
	return nil
}