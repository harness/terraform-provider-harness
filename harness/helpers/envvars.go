package helpers

import "os"

type EnvVar string

var EnvVars = struct {
	HarnessAccountId   EnvVar
	HarnessApiKey      EnvVar
	HarnessEndpoint    EnvVar
	HarnessBearerToken EnvVar
}{
	HarnessAccountId:   "HARNESS_ACCOUNT_ID",
	HarnessApiKey:      "HARNESS_API_KEY",
	HarnessEndpoint:    "HARNESS_ENDPOINT",
	HarnessBearerToken: "HARNESS_BEARER_TOKEN",
}

func (e EnvVar) String() string {
	return string(e)
}

func (e EnvVar) Get() string {
	return os.Getenv(string(e))
}

func (e EnvVar) GetDefault(fallback string) string {
	if value, ok := os.LookupEnv(string(e)); ok {
		return value
	}
	return fallback
}
