package helpers

import (
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
