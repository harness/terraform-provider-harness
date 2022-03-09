package cd

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

var configureClient sync.Once
var apiClient *ApiClient

func getClient() *ApiClient {
	configureClient.Do(func() {
		var err error

		apiClient, err = NewClient(DefaultConfig())

		if err != nil {
			panic(err)
		}
	})

	return apiClient
}

func GetUnauthorizedClient() *ApiClient {
	cfg := DefaultConfig()
	cfg.APIKey = "INVALID_API_KEY"

	c, _ := NewClient(cfg)
	return c
}

func TestClientRequireApiKey_Config(t *testing.T) {
	cfg := DefaultConfig()
	cfg.APIKey = ""

	_, err := NewClient(cfg)

	require.Error(t, err, InvalidConfigError{})
	require.Equal(t, "ApiKey", err.(InvalidConfigError).Field)

	cfg.APIKey = "APIKEY"
	c, err := NewClient(cfg)

	require.Equal(t, c.Configuration.APIKey, cfg.APIKey)
	require.NoError(t, err)
}

func TestClientRequireAccountId_Config(t *testing.T) {
	cfg := DefaultConfig()
	cfg.AccountId = ""

	_, err := NewClient(cfg)

	require.Error(t, err, InvalidConfigError{})
	require.Equal(t, "AccountId", err.(InvalidConfigError).Field)

	cfg.AccountId = "ACCOUNT_ID"
	c, err := NewClient(cfg)

	require.Equal(t, c.Configuration.AccountId, cfg.AccountId)
	require.NoError(t, err)
}
