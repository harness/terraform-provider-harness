package cd

import (
	"os"
	"sync"
	"testing"

	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/stretchr/testify/require"
)

var AccTestConfigureClient sync.Once
var Client *ApiClient

func getClient() *ApiClient {
	AccTestConfigureClient.Do(func() {
		var err error
		Client, err = NewClient(&Config{
			AccountId:    helpers.EnvVars.AccountId.Get(),
			APIKey:       helpers.EnvVars.ApiKey.Get(),
			DebugLogging: true,
		})

		if err != nil {
			panic(err)
		}
	})

	return Client
}

func GetUnauthorizedClient() *ApiClient {
	c, _ := NewClient(&Config{
		AccountId: helpers.EnvVars.AccountId.Get(),
		APIKey:    "BAD KEY",
	})

	return c
}

func TestClientRequireApiKey_Config(t *testing.T) {
	os.Clearenv()

	cfg := &Config{
		AccountId: "ACCOUNT_ID",
	}
	_, err := NewClient(cfg)

	require.Error(t, err, InvalidConfigError{})
	require.Equal(t, "ApiKey", err.(InvalidConfigError).Field)

	cfg = &Config{
		AccountId: "ACCOUNT_ID",
		APIKey:    "APIKEY",
	}
	c, err := NewClient(cfg)

	require.Equal(t, c.Configuration.APIKey, cfg.APIKey)
	require.NoError(t, err)
}

func TestClientRequireApiKey_Envvar(t *testing.T) {
	os.Clearenv()
	os.Setenv(helpers.EnvVars.ApiKey.String(), "APIKEY")

	cfg := &Config{
		AccountId: "ACCOUNT_ID",
	}
	c, err := NewClient(cfg)

	require.NoError(t, err, InvalidConfigError{})
	require.Equal(t, "APIKEY", c.Configuration.APIKey)
}

func TestClientRequireAccountId_Config(t *testing.T) {
	os.Clearenv()

	cfg := &Config{
		APIKey: "APIKEY",
	}
	_, err := NewClient(cfg)

	require.Error(t, err, InvalidConfigError{})
	require.Equal(t, "AccountId", err.(InvalidConfigError).Field)

	cfg = &Config{
		AccountId: "ACCOUNT_ID",
		APIKey:    "APIKEY",
	}
	c, err := NewClient(cfg)

	require.Equal(t, c.Configuration.AccountId, cfg.AccountId)
	require.NoError(t, err)
}

func TestClientRequireAccountId_Envvar(t *testing.T) {
	os.Clearenv()
	os.Setenv(helpers.EnvVars.AccountId.String(), "ACCOUNT_ID")

	cfg := &Config{
		APIKey: "APIKEY",
	}
	c, err := NewClient(cfg)

	require.NoError(t, err, InvalidConfigError{})
	require.Equal(t, "ACCOUNT_ID", c.Configuration.AccountId)
}
