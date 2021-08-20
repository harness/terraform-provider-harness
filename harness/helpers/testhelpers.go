package helpers

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
