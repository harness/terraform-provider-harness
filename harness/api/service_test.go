package api

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
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
		err = c.Applications().DeleteApplication(app.Id)
		require.Nil(t, err, "Failed to delete application: %s", err)
	}()

	// Verify
	svc, _ := ServiceFactory(app.Id, serviceName, cac.DeploymentTypes.Kubernetes, cac.ArtifactTypes.Docker)
	svc.ApplicationId = app.Id
	newService, err := c.Services().UpsertService(svc)
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
	svcInput, err := ServiceFactory(app.Id, serviceName, cac.DeploymentTypes.Kubernetes, cac.ArtifactTypes.Docker)
	require.NoError(t, err)
	require.NotNil(t, svcInput)

	svc, err := c.Services().UpsertService(svcInput)
	require.NoError(t, err)
	require.NotNil(t, svc)

	defer func() {
		err = c.Applications().DeleteApplication(app.Id)
		require.Nil(t, err, "Failed to delete application: %s", err)
	}()

	// Find service by id
	svcLookup, err := c.Services().GetServiceById(app.Id, svc.Id)
	require.NoError(t, err)
	require.NotNil(t, svcLookup)
	require.Equal(t, cac.ArtifactTypes.Docker, svcLookup.ArtifactType)
	require.Equal(t, cac.DeploymentTypes.Kubernetes, svcLookup.DeploymentType)
	require.Equal(t, serviceName, svcLookup.Name)
	require.Equal(t, cac.HelmVersions.V2, svcLookup.HelmVersion)
}

func TestServiceSerialization(t *testing.T) {
	// Setup
	c := getClient()
	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
	app, err := createApplication(appName)
	require.NotNil(t, app)
	require.NoError(t, err)

	defer func() {
		err = c.Applications().DeleteApplication(app.Id)
		require.Nil(t, err, "Failed to delete application: %s", err)
	}()

	t.Run("ssh_service", testServiceSerialization(app.Id, app.Name, cac.DeploymentTypes.SSH, cac.ArtifactTypes.Tar))
	t.Run("ami_service", testServiceSerialization(app.Id, app.Name, cac.DeploymentTypes.AMI, cac.ArtifactTypes.AMI))
	t.Run("aws_codedeploy", testServiceSerialization(app.Id, app.Name, cac.DeploymentTypes.AWSCodeDeploy, cac.ArtifactTypes.AWSCodeDeploy))
	t.Run("aws_lambda", testServiceSerialization(app.Id, app.Name, cac.DeploymentTypes.AWSLambda, cac.ArtifactTypes.AWSLambda))
	t.Run("aws_ecs", testServiceSerialization(app.Id, app.Name, cac.DeploymentTypes.ECS, cac.ArtifactTypes.Docker))
	t.Run("pcf", testServiceSerialization(app.Id, app.Name, cac.DeploymentTypes.PCF, cac.ArtifactTypes.PCF))
	t.Run("winrm_iis_website", testServiceSerialization(app.Id, app.Name, cac.DeploymentTypes.WinRM, cac.ArtifactTypes.IISWebsite))
	t.Run("kubernetes_service", testServiceSerializationWithAdditionalTests(app.Id, app.Name, cac.DeploymentTypes.Kubernetes, cac.ArtifactTypes.Docker, func(t *testing.T, svc *cac.Service) {
		require.Equal(t, cac.HelmVersions.V2, svc.HelmVersion)
	}))
	t.Run("helm", testServiceSerializationWithAdditionalTests(app.Id, app.Name, cac.DeploymentTypes.Helm, cac.ArtifactTypes.Docker, func(t *testing.T, svc *cac.Service) {
		require.Equal(t, cac.HelmVersions.V2, svc.HelmVersion)
	}))

}

func TestDeleteService(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(4))
	app, err := createApplication(expectedName)
	require.NoError(t, err)

	defer func() {
		err = c.Applications().DeleteApplication(app.Id)
		require.Nil(t, err, "Failed to delete application: %s", err)
	}()

	svc, err := createService(app.Id, expectedName, cac.DeploymentTypes.Kubernetes, cac.ArtifactTypes.Docker)
	require.NoError(t, err)
	require.NotNil(t, svc)

	svcLookup, err := c.Services().GetServiceById(app.Id, svc.Id)
	require.NoError(t, err)
	require.NotNil(t, svcLookup)

	err = c.Services().DeleteService(app.Id, svc.Id)
	require.NoError(t, err)

	svcLookup, err = c.Services().GetServiceById(app.Id, svc.Id)
	require.Error(t, err, "received http status code '403'")
	require.Nil(t, svcLookup)
}

func testServiceSerialization(applicationId string, applicationName string, deploymentType string, artifactType string) func(t *testing.T) {
	return testServiceSerializationWithAdditionalTests(applicationId, applicationName, deploymentType, artifactType, nil)
}

func testServiceSerializationWithAdditionalTests(applicationId string, applicationName string, deploymentType string, artifactType string, additionalTests func(t *testing.T, serviceUnderTest *cac.Service)) func(t *testing.T) {
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
		require.Equal(t, cac.ObjectTypes.Service, svc.Type)

		if additionalTests != nil {
			additionalTests(t, svc)
		}

	}

}

func createService(applicationId string, serviceName string, deploymentType string, artifactType string) (*cac.Service, error) {
	serviceInput, err := ServiceFactory(applicationId, serviceName, deploymentType, artifactType)
	if err != nil {
		return nil, err
	}

	serviceInput.ApplicationId = applicationId

	serviceInput.Description = "some description"
	return getClient().Services().UpsertService(serviceInput)
}

var safeTestNameRx = regexp.MustCompile("/")

func getSafeTestName(name string) string {
	return safeTestNameRx.ReplaceAllString(name, "_")
}
