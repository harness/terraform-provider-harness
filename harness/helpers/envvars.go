package helpers

import "os"

type EnvVar string

var EnvVars = struct {
	AccountId      EnvVar
	ApiKey         EnvVar
	BearerToken    EnvVar
	DelegateSecret EnvVar
	Endpoint       EnvVar
	NGApiKey       EnvVar
	NGEndpoint     EnvVar
}{
	AccountId:      "HARNESS_ACCOUNT_ID",
	ApiKey:         "HARNESS_API_KEY",
	BearerToken:    "HARNESS_BEARER_TOKEN",
	DelegateSecret: "HARNESS_DELEGATE_SECRET",
	Endpoint:       "HARNESS_ENDPOINT",
	NGApiKey:       "HARNESS_NG_API_KEY",
	NGEndpoint:     "HARNESS_NG_ENDPOINT",
}

func (e EnvVar) String() string {
	return string(e)
}

func (e EnvVar) Get() string {
	return os.Getenv(string(e))
}

func (e EnvVar) GetWithDefault(fallback string) string {
	if value, ok := os.LookupEnv(string(e)); ok {
		return value
	}
	return fallback
}
