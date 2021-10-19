package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateInfraDefinition_KubernetesDirect_KubernetesManifests(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	app, err := createApplication(name)
	require.NoError(t, err)
	require.NotNil(t, app)

	env, err := createEnvironment(app.Id, name)
	require.NoError(t, err)
	require.NotNil(t, env)

	input := &graphql.KubernetesCloudProvider{}
	input.Name = name
	input.ClusterDetailsType = graphql.ClusterDetailsTypes.InheritClusterDetails
	input.InheritClusterDetails = &graphql.InheritClusterDetails{
		DelegateSelectors: []string{"k8s"},
	}

	cp, err := c.CloudProviders().CreateKubernetesCloudProvider(input)
	require.NoError(t, err)
	require.NotNil(t, cp)

	infraDef := cac.NewEntity(cac.ObjectTypes.InfrastructureDefinition).(*cac.InfrastructureDefinition)
	infraDef.Name = name
	infraDef.ApplicationId = app.Id
	infraDef.EnvironmentId = env.Id
	infraDef.CloudProviderType = cac.CloudProviderTypes.KubernetesCluster
	infraDef.DeploymentType = cac.DeploymentTypes.Kubernetes
	infraDef.InfrastructureDetail = (&cac.InfrastructureKubernetesDirect{
		CloudProviderName: cp.Name,
		Namespace:         "default",
		ReleaseName:       "test",
	}).ToInfrastructureDetail()

	ifraDef, err := c.ConfigAsCode().UpsertInfraDefinition(infraDef)
	require.NoError(t, err)
	require.NotNil(t, ifraDef)

	err = c.Applications().DeleteApplication(app.Id)
	require.NoError(t, err)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateInfraDefinition_KubernetesDirect_Helm(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	app, err := createApplication(name)
	require.NoError(t, err)
	require.NotNil(t, app)

	env, err := createEnvironment(app.Id, name)
	require.NoError(t, err)
	require.NotNil(t, env)

	cpInput := &graphql.KubernetesCloudProvider{}
	cpInput.Name = name
	cpInput.ClusterDetailsType = graphql.ClusterDetailsTypes.InheritClusterDetails
	cpInput.InheritClusterDetails = &graphql.InheritClusterDetails{
		DelegateSelectors: []string{"k8s"},
	}

	cp, err := c.CloudProviders().CreateKubernetesCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	infraDefInput := cac.NewEntity(cac.ObjectTypes.InfrastructureDefinition).(*cac.InfrastructureDefinition)
	infraDefInput.Name = name
	infraDefInput.ApplicationId = app.Id
	infraDefInput.EnvironmentId = env.Id
	infraDefInput.CloudProviderType = cac.CloudProviderTypes.KubernetesCluster
	infraDefInput.DeploymentType = cac.DeploymentTypes.Helm
	infraDefInput.InfrastructureDetail = (&cac.InfrastructureKubernetesDirect{
		CloudProviderName: cp.Name,
		Namespace:         "default",
		ReleaseName:       "test",
	}).ToInfrastructureDetail()

	infraDef, err := c.ConfigAsCode().UpsertInfraDefinition(infraDefInput)
	require.NoError(t, err)
	require.NotNil(t, infraDef)

	err = c.ConfigAsCode().DeleteInfraDefinition(infraDef.ApplicationId, infraDef.EnvironmentId, infraDef.Id)
	require.NoError(t, err)

	id, err := c.ConfigAsCode().GetInfraDefinitionById(infraDef.ApplicationId, infraDef.EnvironmentId, infraDef.Id)
	require.NoError(t, err)
	require.Nil(t, id)

	err = c.Applications().DeleteApplication(app.Id)
	require.NoError(t, err)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateInfraDefinition_Datacenter_SSH(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))

	app, err := createApplication(name)
	require.NoError(t, err)
	require.NotNil(t, app)

	env, err := createEnvironment(app.Id, name)
	require.NoError(t, err)
	require.NotNil(t, env)

	cpInput := &graphql.PhysicalDataCenterCloudProvider{}
	cpInput.Name = name

	cp, err := c.CloudProviders().CreatePhysicalDataCenterCloudProvider(cpInput)
	require.NoError(t, err)
	require.NotNil(t, cp)

	infraDefInput := cac.NewEntity(cac.ObjectTypes.InfrastructureDefinition).(*cac.InfrastructureDefinition)
	infraDefInput.Name = name
	infraDefInput.ApplicationId = app.Id
	infraDefInput.EnvironmentId = env.Id
	infraDefInput.CloudProviderType = cac.CloudProviderTypes.DataCenter
	infraDefInput.DeploymentType = cac.DeploymentTypes.SSH
	infraDefInput.InfrastructureDetail = (&cac.InfrastructureDataCenterSSH{
		CloudProviderName:       cp.Name,
		HostConnectionAttrsName: "test-ssh-cred",
		HostNames:               []string{"localhost", "127.0.0.1", "loopback"},
	}).ToInfrastructureDetail()

	ifraDef, err := c.ConfigAsCode().UpsertInfraDefinition(infraDefInput)
	require.NoError(t, err)
	require.NotNil(t, ifraDef)

	err = c.Applications().DeleteApplication(app.Id)
	require.NoError(t, err)

	err = c.CloudProviders().DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

// func TestGetInfraDefinitionById_KubernetesDirect(t *testing.T) {
// 	c := getClient()

// 	app, err := c.Applications().GetApplicationByName("TestAccDataSourceApplication_idIu")
// 	require.NoError(t, err)
// 	require.NotNil(t, app)

// 	infra, err := c.ConfigAsCode().GetInfraDefinitionById(app.Id,  "iYCnfWzkS72p8OGCDkidyw")
// 	require.NoError(t, err)
// 	require.NotNil(t, infra)
// 	require.Len(t, infra.InfrastructureDetail, 1)

// 	k8sDetail := infra.InfrastructureDetail[0].ToKubernetesDirect()
// 	require.NotNil(t, k8sDetail)

// 	infraDetail := infra.InfrastructureDetail[0]
// 	require.Equal(t, infraDetail.Type, cac.InfrastructureTypes.KubernetesDirect)
// 	require.Equal(t, k8sDetail.CloudProviderName, infraDetail.CloudProviderName)
// 	require.Equal(t, k8sDetail.Namespace, infraDetail.Namespace)
// 	require.Equal(t, k8sDetail.ReleaseName, infraDetail.ReleaseName)
// }

// func TestGetInfraDefinitionByName(t *testing.T) {
// 	c := getClient()

// 	infra, err := c.ConfigAsCode().GetInfraDefinitionByName("TestAccDataSourceApplication_idIu", "test", "k8s")
// 	require.NoError(t, err)
// 	require.NotNil(t, infra)
// }
