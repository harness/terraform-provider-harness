package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvVarsAreSet(t *testing.T) {
	assert.NotEmpty(t, TestEnvVars.AwsAccessKeyId.Get())
	assert.NotEmpty(t, TestEnvVars.AwsSecretAccessKey.Get())
	assert.NotEmpty(t, TestEnvVars.AzureClientId.Get())
	assert.NotEmpty(t, TestEnvVars.AzureClientSecret.Get())
	assert.NotEmpty(t, TestEnvVars.AzureTenantId.Get())
	assert.NotEmpty(t, TestEnvVars.SpotInstAccountId.Get())
	assert.NotEmpty(t, TestEnvVars.SpotInstToken.Get())
}

type EnvVar string

var TestEnvVars = struct {
	AwsAccessKeyId     EnvVar
	AwsSecretAccessKey EnvVar
	AzureClientId      EnvVar
	AzureClientSecret  EnvVar
	AzureTenantId      EnvVar
	SpotInstAccountId  EnvVar
	SpotInstToken      EnvVar
}{
	AwsAccessKeyId:     "HARNESS_TEST_AWS_ACCESS_KEY_ID",
	AwsSecretAccessKey: "HARNESS_TEST_AWS_SECRET_ACCESS_KEY",
	AzureClientId:      "HARNESS_TEST_AZURE_CLIENT_ID",
	AzureClientSecret:  "HARNESS_TEST_AZURE_CLIENT_SECRET",
	AzureTenantId:      "HARNESS_TEST_AZURE_TENANT_ID",
	SpotInstAccountId:  "HARNESS_TEST_SPOT_ACCT_ID",
	SpotInstToken:      "HARNESS_TEST_SPOT_TOKEN",
}

func (e EnvVar) Get() string {
	return os.Getenv(string(e))
}
