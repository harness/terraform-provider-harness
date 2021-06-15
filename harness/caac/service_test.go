package caac

import (
	"fmt"
	"testing"

	"github.com/micahlmartin/terraform-provider-harness/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateService(t *testing.T) {
	// Setup
	c := getClient()
	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
	serviceName := fmt.Sprintf("svc-%s-%s", t.Name(), utils.RandStringBytes(5))
	app, err := createApplication(appName)
	require.NoError(t, err)
	require.NotNil(t, app)

	// Cleanup
	defer func() {
		c.ApiClient.Applications().DeleteApplication(app.Id)
	}()

	// Verify
	svc, _ := ServiceFactory(app.Id, serviceName, DeploymentTypes.Kubernetes, ArtifactTypes.Docker)
	svc.ApplicationId = app.Id
	newService, err := c.UpsertService(svc)
	require.NoError(t, err)
	require.NotEmpty(t, newService.Id)
	require.Equal(t, app.Id, newService.ApplicationId)
}

func TestGetServiceById(t *testing.T) {

	// Create application
	c := getClient()
	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
	app, err := createApplication(appName)
	require.NotNil(t, app)
	require.NoError(t, err)

	// Create service
	serviceName := fmt.Sprintf("%s-%s", getSafeTestName(t.Name()), utils.RandStringBytes(4))
	svcInput, err := ServiceFactory(app.Id, serviceName, DeploymentTypes.Kubernetes, ArtifactTypes.Docker)
	require.NoError(t, err)
	require.NotNil(t, svcInput)

	svc, err := c.UpsertService(svcInput)
	require.NoError(t, err)
	require.NotNil(t, svc)

	defer func() {
		c.ApiClient.Applications().DeleteApplication(app.Id)
	}()

	// Find service by id
	svcLookup, err := c.GetServiceById(app.Id, svc.Id)
	require.NoError(t, err)
	require.NotNil(t, svcLookup)
	require.Equal(t, ArtifactTypes.Docker, svcLookup.ArtifactType)
	require.Equal(t, DeploymentTypes.Kubernetes, svcLookup.DeploymentType)
	require.Equal(t, serviceName, svcLookup.Name)
	require.Equal(t, HelmVersions.V2, svcLookup.HelmVersion)
}

func TestServiceSerialization(t *testing.T) {
	// Setup
	c := getClient()
	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
	app, err := createApplication(appName)
	require.NotNil(t, app)
	require.NoError(t, err)

	defer func() {
		c.ApiClient.Applications().DeleteApplication(app.Id)
	}()

	t.Run("ssh_service", testServiceSerialization(app.Id, app.Name, DeploymentTypes.SSH, ArtifactTypes.Tar))
	t.Run("ami_service", testServiceSerialization(app.Id, app.Name, DeploymentTypes.AMI, ArtifactTypes.AMI))
	t.Run("aws_codedeploy", testServiceSerialization(app.Id, app.Name, DeploymentTypes.AWSCodeDeploy, ArtifactTypes.AWSCodeDeploy))
	t.Run("aws_lambda", testServiceSerialization(app.Id, app.Name, DeploymentTypes.AWSLambda, ArtifactTypes.AWSLambda))
	t.Run("aws_ecs", testServiceSerialization(app.Id, app.Name, DeploymentTypes.ECS, ArtifactTypes.Docker))
	t.Run("pcf", testServiceSerialization(app.Id, app.Name, DeploymentTypes.PCF, ArtifactTypes.PCF))
	t.Run("winrm_iis_website", testServiceSerialization(app.Id, app.Name, DeploymentTypes.WinRM, ArtifactTypes.IISWebsite))
	t.Run("kubernetes_service", testServiceSerializationWithAdditionalTests(app.Id, app.Name, DeploymentTypes.Kubernetes, ArtifactTypes.Docker, func(t *testing.T, svc *Service) {
		require.Equal(t, HelmVersions.V2, svc.HelmVersion)
	}))
	t.Run("helm", testServiceSerializationWithAdditionalTests(app.Id, app.Name, DeploymentTypes.Helm, ArtifactTypes.Docker, func(t *testing.T, svc *Service) {
		require.Equal(t, HelmVersions.V2, svc.HelmVersion)
	}))

	// t.Run("custom", testServiceSerialization(app.Id, app.Name, DeploymentTypes.Custom, ArtifactTypes.Tar))

	// Cleanup

}

func testServiceSerialization(applicationId string, applicationName string, deploymentType string, artifactType string) func(t *testing.T) {
	return testServiceSerializationWithAdditionalTests(applicationId, applicationName, deploymentType, artifactType, nil)
}

func testServiceSerializationWithAdditionalTests(applicationId string, applicationName string, deploymentType string, artifactType string, additionalTests func(t *testing.T, serviceUnderTest *Service)) func(t *testing.T) {
	return func(t *testing.T) {
		// Create service
		serviceName := fmt.Sprintf("%s-%s", getSafeTestName(t.Name()), utils.RandStringBytes(4))
		svc, err := createService(applicationId, serviceName, deploymentType, artifactType)
		require.NoError(t, err)
		require.NotNil(t, svc)

		// Verify
		require.Equal(t, deploymentType, svc.DeploymentType)
		require.Equal(t, serviceName, svc.Name)
		require.Equal(t, artifactType, svc.ArtifactType)
		require.Equal(t, ObjectTypes.Service, svc.Type)

		if additionalTests != nil {
			additionalTests(t, svc)
		}

	}

}
