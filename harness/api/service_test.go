package api

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/micahlmartin/terraform-provider-harness/harness/api/caac"
	"github.com/micahlmartin/terraform-provider-harness/harness/utils"
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
		c.Applications().DeleteApplication(app.Id)
	}()

	// Verify
	svc, _ := ServiceFactory(app.Id, serviceName, caac.DeploymentTypes.Kubernetes, caac.ArtifactTypes.Docker)
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
	svcInput, err := ServiceFactory(app.Id, serviceName, caac.DeploymentTypes.Kubernetes, caac.ArtifactTypes.Docker)
	require.NoError(t, err)
	require.NotNil(t, svcInput)

	svc, err := c.Services().UpsertService(svcInput)
	require.NoError(t, err)
	require.NotNil(t, svc)

	defer func() {
		c.Applications().DeleteApplication(app.Id)
	}()

	// Find service by id
	svcLookup, err := c.Services().GetServiceById(app.Id, svc.Id)
	require.NoError(t, err)
	require.NotNil(t, svcLookup)
	require.Equal(t, caac.ArtifactTypes.Docker, svcLookup.ArtifactType)
	require.Equal(t, caac.DeploymentTypes.Kubernetes, svcLookup.DeploymentType)
	require.Equal(t, serviceName, svcLookup.Name)
	require.Equal(t, caac.HelmVersions.V2, svcLookup.HelmVersion)
}

func TestServiceSerialization(t *testing.T) {
	// Setup
	c := getClient()
	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
	app, err := createApplication(appName)
	require.NotNil(t, app)
	require.NoError(t, err)

	defer func() {
		c.Applications().DeleteApplication(app.Id)
	}()

	t.Run("ssh_service", testServiceSerialization(app.Id, app.Name, caac.DeploymentTypes.SSH, caac.ArtifactTypes.Tar))
	t.Run("ami_service", testServiceSerialization(app.Id, app.Name, caac.DeploymentTypes.AMI, caac.ArtifactTypes.AMI))
	t.Run("aws_codedeploy", testServiceSerialization(app.Id, app.Name, caac.DeploymentTypes.AWSCodeDeploy, caac.ArtifactTypes.AWSCodeDeploy))
	t.Run("aws_lambda", testServiceSerialization(app.Id, app.Name, caac.DeploymentTypes.AWSLambda, caac.ArtifactTypes.AWSLambda))
	t.Run("aws_ecs", testServiceSerialization(app.Id, app.Name, caac.DeploymentTypes.ECS, caac.ArtifactTypes.Docker))
	t.Run("pcf", testServiceSerialization(app.Id, app.Name, caac.DeploymentTypes.PCF, caac.ArtifactTypes.PCF))
	t.Run("winrm_iis_website", testServiceSerialization(app.Id, app.Name, caac.DeploymentTypes.WinRM, caac.ArtifactTypes.IISWebsite))
	t.Run("kubernetes_service", testServiceSerializationWithAdditionalTests(app.Id, app.Name, caac.DeploymentTypes.Kubernetes, caac.ArtifactTypes.Docker, func(t *testing.T, svc *caac.Service) {
		require.Equal(t, caac.HelmVersions.V2, svc.HelmVersion)
	}))
	t.Run("helm", testServiceSerializationWithAdditionalTests(app.Id, app.Name, caac.DeploymentTypes.Helm, caac.ArtifactTypes.Docker, func(t *testing.T, svc *caac.Service) {
		require.Equal(t, caac.HelmVersions.V2, svc.HelmVersion)
	}))

}

func testServiceSerialization(applicationId string, applicationName string, deploymentType string, artifactType string) func(t *testing.T) {
	return testServiceSerializationWithAdditionalTests(applicationId, applicationName, deploymentType, artifactType, nil)
}

func testServiceSerializationWithAdditionalTests(applicationId string, applicationName string, deploymentType string, artifactType string, additionalTests func(t *testing.T, serviceUnderTest *caac.Service)) func(t *testing.T) {
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
		require.Equal(t, caac.ObjectTypes.Service, svc.Type)

		if additionalTests != nil {
			additionalTests(t, svc)
		}

	}

}

func createService(applicationId string, serviceName string, deploymentType string, artifactType string) (*caac.Service, error) {
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
