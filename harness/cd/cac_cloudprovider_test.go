package cd

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetCloudProviderById(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.PhysicalDataCenterCloudProvider).(*cac.PhysicalDatacenterCloudProvider)
	cpInput.Name = expectedName

	c := getClient()

	cp, err := c.ConfigAsCodeClient.UpsertPhysicalDataCenterCloudProvider(cpInput)
	require.NoError(t, err)

	testCP := &cac.PhysicalDatacenterCloudProvider{}
	err = c.ConfigAsCodeClient.GetCloudProviderById(cp.Id, testCP)
	require.NoError(t, err)

	require.Equal(t, cp, testCP)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestDeleteCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.PhysicalDataCenterCloudProvider).(*cac.PhysicalDatacenterCloudProvider)
	cpInput.Name = expectedName

	c := getClient()

	cp, err := c.ConfigAsCodeClient.UpsertPhysicalDataCenterCloudProvider(cpInput)
	require.NoError(t, err)

	testCP := &cac.PhysicalDatacenterCloudProvider{}
	err = c.ConfigAsCodeClient.GetCloudProviderById(cp.Id, testCP)
	require.NoError(t, err)

	require.Equal(t, cp, testCP)

	err = c.ConfigAsCodeClient.DeleteCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP := &cac.PhysicalDatacenterCloudProvider{}
	err = c.ConfigAsCodeClient.GetCloudProviderByName(cpInput.Name, foundCP)
	require.NoError(t, err)
	require.Equal(t, &cac.PhysicalDatacenterCloudProvider{}, foundCP)
}

func TestCacSpotInstCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.SpotInstCloudProvider).(*cac.SpotInstCloudProvider)

	secret, err := createEncryptedTextSecret(expectedName, helpers.TestEnvVars.SpotInstToken.Get())
	require.NoError(t, err)

	cpInput.Name = expectedName
	cpInput.AccountId = helpers.TestEnvVars.SpotInstAccountId.Get()
	cpInput.Token = &cac.SecretRef{
		Name: secret.Name,
	}

	c := getClient()
	cp, err := c.ConfigAsCodeClient.UpsertSpotInstCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	require.Equal(t, cpInput, cp)

	err = c.CloudProviderClient.DeleteCloudProvider(cpInput.Id)
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
		Name: secret.Name,
	}

	c := getClient()
	cp, err := c.ConfigAsCodeClient.UpsertPcfCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	require.Equal(t, cpInput, cp)

	err = c.CloudProviderClient.DeleteCloudProvider(cpInput.Id)
	require.NoError(t, err)
}

func TestCacKubernetesCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.KubernetesCloudProvider).(*cac.KubernetesCloudProvider)

	secret, err := createEncryptedTextSecret(expectedName, helpers.TestEnvVars.AzureClientSecret.Get())
	require.NoError(t, err)

	cpInput.Name = expectedName
	cpInput.SkipValidation = true
	cpInput.MasterUrl = "https://example.com"
	cpInput.ServiceAccountToken = &cac.SecretRef{
		Name: secret.Name,
	}

	c := getClient()
	cp, err := c.ConfigAsCodeClient.UpsertKubernetesCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	require.Equal(t, cpInput, cp)

	err = c.CloudProviderClient.DeleteCloudProvider(cpInput.Id)
	require.NoError(t, err)
}

func TestCacUpsertAzureCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.AzureCloudProvider).(*cac.AzureCloudProvider)

	secret, err := createEncryptedTextSecret(expectedName, helpers.TestEnvVars.AzureClientSecret.Get())
	require.NoError(t, err)

	cpInput.Name = expectedName
	cpInput.AzureEnvironmentType = cac.AzureEnvironmentTypes.AzureGlobal
	cpInput.ClientId = helpers.TestEnvVars.AzureClientId.Get()
	cpInput.TenantId = helpers.TestEnvVars.AzureTenantId.Get()
	cpInput.Key = &cac.SecretRef{
		Name: secret.Name,
	}

	c := getClient()
	cp, err := c.ConfigAsCodeClient.UpsertAzureCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	require.Equal(t, cpInput, cp)

	err = c.CloudProviderClient.DeleteCloudProvider(cpInput.Id)
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
	cp, err := c.ConfigAsCodeClient.UpsertGcpCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	cpInput.Id = cp.Id
	require.Equal(t, cpInput, cp)

	err = c.CloudProviderClient.DeleteCloudProvider(cpInput.Id)
	require.NoError(t, err)
}

func TestUpsertPhysicalDataCenterCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	cpInput := cac.NewEntity(cac.ObjectTypes.PhysicalDataCenterCloudProvider).(*cac.PhysicalDatacenterCloudProvider)
	cpInput.Name = expectedName

	c := getClient()

	cp, err := c.ConfigAsCodeClient.UpsertPhysicalDataCenterCloudProvider(cpInput)
	require.NoError(t, err)

	cpInput.Id = cp.Id
	require.Equal(t, cpInput, cp)

	err = c.CloudProviderClient.DeleteCloudProvider(cpInput.Id)
	require.NoError(t, err)
}

func TestUpsertAwsCloudProvider(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	secret, err := createEncryptedTextSecret(expectedName, helpers.TestEnvVars.AwsSecretAccessKey.Get())
	require.NoError(t, err)

	cpInput := cac.NewEntity(cac.ObjectTypes.AwsCloudProvider).(*cac.AwsCloudProvider)
	cpInput.Name = expectedName
	cpInput.AccessKey = helpers.TestEnvVars.AwsAccessKeyId.Get()
	cpInput.SecretKey = &cac.SecretRef{
		Name: secret.Name,
	}

	c := getClient()

	cp, err := c.ConfigAsCodeClient.UpsertAwsCloudProvider(cpInput)
	require.NoError(t, err)

	cpInput.Id = cp.Id
	require.Equal(t, cpInput, cp)

	err = c.CloudProviderClient.DeleteCloudProvider(cpInput.Id)
	require.NoError(t, err)
}

func TestListCloudProviders(t *testing.T) {
	client := getClient()
	limit := 100
	offset := 0
	hasMore := true

	for hasMore {
		cps, pagination, err := client.CloudProviderClient.ListCloudProviders(limit, offset)
		require.NoError(t, err, "Failed to list cloud providers: %s", err)
		require.NotEmpty(t, cps, "No cloud providers found")
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(cps) == limit
		offset += limit
	}
}
