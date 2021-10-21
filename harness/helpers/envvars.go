package helpers

import "os"

type EnvVar string

var EnvVars = struct {
	HarnessAccountId   EnvVar
	HarnessApiKey      EnvVar
	HarnessEndpoint    EnvVar
	HarnessBearerToken EnvVar
	HarnessNGApiKey    EnvVar
	HarnessNGEndpoint  EnvVar
}{
	HarnessAccountId:   "HARNESS_ACCOUNT_ID",
	HarnessApiKey:      "HARNESS_API_KEY",
	HarnessEndpoint:    "HARNESS_ENDPOINT",
	HarnessBearerToken: "HARNESS_BEARER_TOKEN",
	HarnessNGApiKey:    "HARNESS_NG_API_KEY",
	HarnessNGEndpoint:  "HARNESS_NG_ENDPOINT",
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
