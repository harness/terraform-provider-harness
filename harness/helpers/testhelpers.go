package helpers

var TestEnvVars = struct {
	AwsAccessKeyId      EnvVar
	AwsSecretAccessKey  EnvVar
	AzureClientId       EnvVar
	AzureClientSecret   EnvVar
	AzureTenantId       EnvVar
	DelegateSecret      EnvVar
	DelegateProfileId   EnvVar
	DelegateWaitTimeout EnvVar
	SpotInstAccountId   EnvVar
	SpotInstToken       EnvVar
}{
	AwsAccessKeyId:      "HARNESS_TEST_AWS_ACCESS_KEY_ID",
	AwsSecretAccessKey:  "HARNESS_TEST_AWS_SECRET_ACCESS_KEY",
	AzureClientId:       "HARNESS_TEST_AZURE_CLIENT_ID",
	AzureClientSecret:   "HARNESS_TEST_AZURE_CLIENT_SECRET",
	AzureTenantId:       "HARNESS_TEST_AZURE_TENANT_ID",
	DelegateSecret:      "HARNESS_TEST_DELEGATE_SECRET",
	DelegateProfileId:   "HARNESS_TEST_DELEGATE_PROFILE_ID",
	DelegateWaitTimeout: "HARNESS_TEST_DELEGATE_WAIT_TIMEOUT",
	SpotInstAccountId:   "HARNESS_TEST_SPOT_ACCT_ID",
	SpotInstToken:       "HARNESS_TEST_SPOT_TOKEN",
}
