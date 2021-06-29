package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetKubernetesCloudProviderById(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createKubernetesCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviders().GetKubernetesCloudProviderById(cp.Id)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, cp.Id, foundCP.Id)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestGetKubernetesCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createKubernetesCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviders().GetKubernetesCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateKubernetesCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createKubernetesCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func createKubernetesCloudProvider(name string) (*graphql.KubernetesCloudProvider, error) {
	c := getClient()

	input := &graphql.KubernetesCloudProvider{}
	input.Name = name
	input.ClusterDetailsType = graphql.ClusterDetailsTypes.InheritClusterDetails
	input.InheritClusterDetails = &graphql.InheritClusterDetails{
		DelegateSelectors: []string{"Primary"},
	}
	input.SkipValidation = true

	cp, err := c.CloudProviders().CreateKubernetesCloudProvider(input)
	if err != nil {
		return nil, err
	}

	return cp, nil
}
