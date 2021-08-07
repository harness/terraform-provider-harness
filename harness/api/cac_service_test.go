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
	svc, _ := cac.NewEntity(cac.ObjectTypes.Service).(*cac.Service)
	svc.Name = serviceName
	svc.ApplicationId = app.Id
	svc.DeploymentType = cac.DeploymentTypes.Kubernetes
	svc.ArtifactType = cac.ArtifactTypes.Docker

	newService := &cac.Service{}
	err = c.ConfigAsCode().UpsertObject(svc, cac.GetServiceYamlPath(app.Name, serviceName), newService)
	require.NoError(t, err)
	require.NotEmpty(t, newService.Id)
	require.Equal(t, app.Id, newService.ApplicationId)
}

func TestGetService(t *testing.T) {

	// Create application
	c := getClient()
	appName := fmt.Sprintf("app_%s_%s", t.Name(), utils.RandStringBytes(4))
	app, err := createApplication(appName)
	require.NotNil(t, app)
	require.NoError(t, err)

	// Create service
	serviceName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	svcInput := cac.NewEntity(cac.ObjectTypes.Service).(*cac.Service)
	svcInput.Name = serviceName
	svcInput.ApplicationId = app.Id
	svcInput.DeploymentType = cac.DeploymentTypes.Kubernetes
	svcInput.ArtifactType = cac.ArtifactTypes.Docker

	require.NoError(t, err)
	require.NotNil(t, svcInput)

	svc := &cac.Service{}
	err = c.ConfigAsCode().UpsertObject(svcInput, cac.GetServiceYamlPath(app.Name, serviceName), svc)
	require.NoError(t, err)
	require.NotNil(t, svc)

	defer func() {
		err = c.Applications().DeleteApplication(app.Id)
		require.Nil(t, err, "Failed to delete application: %s", err)
	}()

	// Find service by id
	svcLookup := &cac.Service{}
	err = c.ConfigAsCode().FindObjectByPath(app.Id, cac.GetServiceYamlPath(app.Name, serviceName), svcLookup)
	require.NoError(t, err)
	require.NotNil(t, svcLookup)
	require.Equal(t, cac.ArtifactTypes.Docker, svcLookup.ArtifactType)
	require.Equal(t, cac.DeploymentTypes.Kubernetes, svcLookup.DeploymentType)
	require.Equal(t, serviceName, svcLookup.Name)
	require.Equal(t, cac.HelmVersions.V2, svcLookup.HelmVersion)
}

func TestGetServiceById(t *testing.T) {

	// Create application
	c := getClient()
	appName := fmt.Sprintf("app_%s_%s", t.Name(), utils.RandStringBytes(4))
	app, err := createApplication(appName)
	require.NotNil(t, app)
	require.NoError(t, err)

	// Create service
	serviceName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	svcInput := cac.NewEntity(cac.ObjectTypes.Service).(*cac.Service)
	svcInput.Name = serviceName
	svcInput.ApplicationId = app.Id
	svcInput.DeploymentType = cac.DeploymentTypes.Kubernetes
	svcInput.ArtifactType = cac.ArtifactTypes.Docker

	require.NoError(t, err)
	require.NotNil(t, svcInput)

	svc := &cac.Service{}
	err = c.ConfigAsCode().UpsertObject(svcInput, cac.GetServiceYamlPath(app.Name, serviceName), svc)
	require.NoError(t, err)
	require.NotNil(t, svc)

	defer func() {
		err := c.Applications().DeleteApplication(app.Id)
		require.NoError(t, err)
	}()

	// Find service by id
	svcLookup, err := c.ConfigAsCode().GetServiceById(app.Id, svc.Id)
	require.NoError(t, err)
	require.Equal(t, svc, svcLookup)
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

	svc, err := createService(app.Id, app.Name, expectedName, cac.DeploymentTypes.Kubernetes, cac.ArtifactTypes.Docker)
	require.NoError(t, err)
	require.NotNil(t, svc)

	svcYamlPath := cac.GetServiceYamlPath(app.Name, svc.Name)

	svcLookup := &cac.Service{}
	err = c.ConfigAsCode().FindObjectByPath(app.Id, svcYamlPath, svcLookup)
	require.NoError(t, err)
	require.NotNil(t, svcLookup)

	err = c.ConfigAsCode().DeleteEntity(svcYamlPath)
	require.NoError(t, err)

	svcLookup = &cac.Service{}
	err = c.ConfigAsCode().FindObjectByPath(app.Id, cac.GetServiceYamlPath(app.Name, svc.Name), svcLookup)
	require.Error(t, err, "received http status code '403'")
}

func testServiceSerialization(applicationId string, applicationName string, deploymentType cac.DeploymentType, artifactType cac.ArtifactType) func(t *testing.T) {
	return testServiceSerializationWithAdditionalTests(applicationId, applicationName, deploymentType, artifactType, nil)
}

func testServiceSerializationWithAdditionalTests(applicationId string, applicationName string, deploymentType cac.DeploymentType, artifactType cac.ArtifactType, additionalTests func(t *testing.T, serviceUnderTest *cac.Service)) func(t *testing.T) {
	return func(t *testing.T) {
		// Create service
		serviceName := fmt.Sprintf("%s-%s", getSafeTestName(t.Name()), utils.RandStringBytes(4))

		svc, err := createService(applicationId, applicationName, serviceName, deploymentType, artifactType)
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

func createService(applicationId string, applicationName string, serviceName string, deploymentType cac.DeploymentType, artifactType cac.ArtifactType) (*cac.Service, error) {
	serviceInput := cac.NewEntity(cac.ObjectTypes.Service).(*cac.Service)
	serviceInput.Name = serviceName
	serviceInput.DeploymentType = deploymentType
	serviceInput.ArtifactType = artifactType
	serviceInput.ApplicationId = applicationId

	serviceInput.Description = "some description"
	svc := &cac.Service{}
	filePath := cac.GetServiceYamlPath(applicationName, serviceName)
	err := getClient().ConfigAsCode().UpsertObject(serviceInput, filePath, svc)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

var safeTestNameRx = regexp.MustCompile("/")

func getSafeTestName(name string) string {
	return safeTestNameRx.ReplaceAllString(name, "_")
}
