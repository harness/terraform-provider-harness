package cd

import (
	"github.com/harness-io/harness-go-sdk/harness/helpers"
)

func getClient() *ApiClient {
	return NewClient(&Configuration{
		AccountId:    helpers.EnvVars.AccountId.Get(),
		APIKey:       helpers.EnvVars.ApiKey.Get(),
		DebugLogging: false,
	})
}

func GetUnauthorizedClient() *ApiClient {
	return NewClient(&Configuration{
		AccountId: helpers.EnvVars.AccountId.Get(),
		APIKey:    "BAD KEY",
	})
}
