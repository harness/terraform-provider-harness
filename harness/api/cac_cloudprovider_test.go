package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestCacSpotInstCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.SpotInstCloudProvider).(*cac.SpotInstCloudProvider)

	secret, err := createEncryptedTextSecret(expectedName, TestEnvVars.SpotInstToken.Get())
	require.NoError(t, err)

	cpInput.Name = expectedName
	cpInput.AccountId = TestEnvVars.SpotInstAccountId.Get()
	cpInput.Token = &cac.SecretRef{
		SecretManagerType: cac.SecretManagerTypes.GcpKMS,
		SecretId:          secret.Id,
	}

	c := getClient()
	cp, err := c.ConfigAsCode().UpsertSpotInstCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	cp.Token.SecretManagerType = cac.SecretManagerTypes.GcpKMS
	require.Equal(t, cpInput, cp)

	err = c.ConfigAsCode().DeleteCloudProvider(cpInput.Name)
	require.NoError(t, err)
}

func TestCacPcfCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.PcfCloudProvider).(*cac.PcfCloudProvider)

	secret, err := createEncryptedTextSecret(expectedName, "pcftest")
	require.NoError(t, err)

	cpInput.Name = expectedName
	cpInput.SkipValidation = true
	cpInput.EndpointUrl = "https://example.com"
	cpInput.Username = "username"
	cpInput.Password = &cac.SecretRef{
		SecretManagerType: cac.SecretManagerTypes.GcpKMS,
		SecretId:          secret.Id,
	}

	c := getClient()
	cp, err := c.ConfigAsCode().UpsertPcfCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	cp.Password.SecretManagerType = cac.SecretManagerTypes.GcpKMS
	require.Equal(t, cpInput, cp)

	err = c.ConfigAsCode().DeleteCloudProvider(cpInput.Name)
	require.NoError(t, err)
}

func TestCacKubernetesCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.KubernetesCloudProvider).(*cac.KubernetesCloudProvider)

	secret, err := createEncryptedTextSecret(expectedName, TestEnvVars.AzureClientSecret.Get())
	require.NoError(t, err)

	cpInput.Name = expectedName
	cpInput.SkipValidation = true
	cpInput.MasterUrl = "https://example.com"
	cpInput.ServiceAccountToken = &cac.SecretRef{
		SecretManagerType: cac.SecretManagerTypes.GcpKMS,
		SecretId:          secret.Id,
	}

	c := getClient()
	cp, err := c.ConfigAsCode().UpsertKubernetesCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	cp.ServiceAccountToken.SecretManagerType = cac.SecretManagerTypes.GcpKMS
	require.Equal(t, cpInput, cp)

	err = c.ConfigAsCode().DeleteCloudProvider(cpInput.Name)
	require.NoError(t, err)
}

func TestCacUpsertAzureCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.AzureCloudProvider).(*cac.AzureCloudProvider)

	secret, err := createEncryptedTextSecret(expectedName, TestEnvVars.AzureClientSecret.Get())
	require.NoError(t, err)

	cpInput.Name = expectedName
	cpInput.AzureEnvironmentType = cac.AzureEnvironmentTypes.AzureGlobal
	cpInput.ClientId = TestEnvVars.AzureClientId.Get()
	cpInput.TenantId = TestEnvVars.AzureTenantId.Get()
	cpInput.Key = &cac.SecretRef{
		SecretManagerType: cac.SecretManagerTypes.GcpKMS,
		SecretId:          secret.Id,
	}

	c := getClient()
	cp, err := c.ConfigAsCode().UpsertAzureCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	cp.Key.SecretManagerType = cac.SecretManagerTypes.GcpKMS
	require.Equal(t, cpInput, cp)

	err = c.ConfigAsCode().DeleteCloudProvider(cpInput.Name)
	require.NoError(t, err)
}

func TestCacUpsertGcpCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.GcpCloudProvider).(*cac.GcpCloudProvider)
	cpInput.Name = expectedName
	cpInput.DelegateSelectors = []string{"Primary"}
	cpInput.SkipValidation = true
	cpInput.UseDelegateSelectors = true

	c := getClient()
	cp, err := c.ConfigAsCode().UpsertGcpCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	require.Equal(t, cpInput, cp)

	err = c.ConfigAsCode().DeleteCloudProvider(cpInput.Name)
	require.NoError(t, err)
}

func TestUpsertPhysicalDataCenterCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.PhysicalDataCenterCloudProvider).(*cac.PhysicalDatacenterCloudProvider)
	cpInput.Name = expectedName

	c := getClient()

	cp, err := c.ConfigAsCode().UpsertPhysicalDataCenterCloudProvider(cpInput)
	require.NoError(t, err)

	cpInput.Id = cp.Id
	require.Equal(t, cpInput, cp)

	err = c.ConfigAsCode().DeleteCloudProvider(cpInput.Name)
	require.NoError(t, err)
}

func TestUpsertAwsCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	secret, err := createEncryptedTextSecret(expectedName, TestEnvVars.AwsSecretAccessKey.Get())
	require.NoError(t, err)

	cpInput := cac.NewEntity(cac.ObjectTypes.AwsCloudProvider).(*cac.AwsCloudProvider)
	cpInput.Name = expectedName
	cpInput.AccessKey = TestEnvVars.AwsAccessKeyId.Get()
	cpInput.SecretKey = &cac.SecretRef{
		SecretManagerType: cac.SecretManagerTypes.GcpKMS,
		SecretId:          secret.Id,
	}

	c := getClient()

	cp, err := c.ConfigAsCode().UpsertAwsCloudProvider(cpInput)
	require.NoError(t, err)

	cpInput.Id = cp.Id
	cp.SecretKey.SecretManagerType = cac.SecretManagerTypes.GcpKMS
	require.Equal(t, cpInput, cp)

	err = c.ConfigAsCode().DeleteCloudProvider(cpInput.Name)
	require.NoError(t, err)
}
