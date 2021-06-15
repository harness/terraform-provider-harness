package caac

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"sync"
	"testing"
	"time"

	"github.com/micahlmartin/terraform-provider-harness/internal/client"
	"github.com/micahlmartin/terraform-provider-harness/internal/envvar"
	"github.com/micahlmartin/terraform-provider-harness/internal/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateKubernetesService(t *testing.T) {
	// Setup
	expectedName := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(5))
	service, err := ServiceFactory(expectedName, DeploymentTypes.Kubernetes, "")
	require.NoError(t, err)
	service.Description = "some description"
	application, err := createApplication(expectedName)
	require.NoError(t, err)
	expectedPath := fmt.Sprintf("Setup/Applications/%s/Services/%s/Index.yaml", application.Name, expectedName)

	// Create Service
	c := getClient()
	item, err := c.UpsertService(application.Name, expectedName, &service)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, item)
	require.Equal(t, expectedPath, item.YamlFilePath)

	c.ApiClient.Applications().DeleteApplication(application.Id)
}

func TestGetDirectoryTreeItem(t *testing.T) {

	// Setup
	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
	serviceName := fmt.Sprintf("service-%s-%s", t.Name(), utils.RandStringBytes(5))
	c := getClient()

	app, err := createApplication(appName)
	ks, _ := createService(appName, serviceName, DeploymentTypes.Kubernetes, "")
	rootItem, _ := c.GetDirectoryTree(app.Id)
	serviceItem := FindConfigAsCodeItemByPath(rootItem, ks.YamlFilePath)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, serviceItem)
	require.Equal(t, ks.YamlFilePath, serviceItem.DirectoryPath.Path)

	// Cleanup
	err = c.ApiClient.Applications().DeleteApplication(app.Id)
	require.NoError(t, err)
}

func TestGetDirectoryItemContent(t *testing.T) {

	// Setup
	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
	serviceName := fmt.Sprintf("service-%s-%s", t.Name(), utils.RandStringBytes(5))
	c := getClient()

	// Create objects for testing
	app, _ := createApplication(appName)
	require.NotNil(t, app)
	ks, _ := createService(appName, serviceName, DeploymentTypes.Kubernetes, "")
	require.NotNil(t, ks)

	rootItem, _ := c.GetDirectoryTree(app.Id)
	require.NotNil(t, rootItem)
	serviceItem := FindConfigAsCodeItemByPath(rootItem, ks.YamlFilePath)
	require.NotNil(t, serviceItem)

	// Test content
	content, err := c.GetDirectoryItemContent(serviceItem.RestName, serviceItem.UUID, app.Id)
	require.NoError(t, err)
	require.NotNil(t, content)

	// Verify
	require.Equal(t, ks.YamlFilePath, serviceItem.DirectoryPath.Path)

	// Cleanup
	err = c.ApiClient.Applications().DeleteApplication(app.Id)
	require.NoError(t, err)
}

func TestSerialization(t *testing.T) {
	// Setup
	c := getClient()
	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
	app, err := createApplication(appName)
	require.NotNil(t, app)
	require.NoError(t, err)

	t.Run("kubernetes_service", testServiceSerializationWithAdditionalTests(app.Id, app.Name, DeploymentTypes.Kubernetes, ArtifactTypes.Docker, func(t *testing.T, svc *Service) {
		require.Equal(t, HelmVersions.V3, svc.HelmVersion)
	}))
	time.Sleep(3 * time.Second)

	t.Run("ssh_service", testServiceSerialization(app.Id, app.Name, DeploymentTypes.SSH, ArtifactTypes.Tar))
	time.Sleep(3 * time.Second)

	t.Run("ami_service", testServiceSerialization(app.Id, app.Name, DeploymentTypes.AMI, ArtifactTypes.AMI))
	time.Sleep(3 * time.Second)

	t.Run("aws_codedeploy", testServiceSerialization(app.Id, app.Name, DeploymentTypes.AWSCodeDeploy, ArtifactTypes.AWSCodeDeploy))
	time.Sleep(3 * time.Second)

	t.Run("aws_lambda", testServiceSerialization(app.Id, app.Name, DeploymentTypes.AWSLambda, ArtifactTypes.AWSLambda))
	time.Sleep(3 * time.Second)

	t.Run("aws_ecs", testServiceSerialization(app.Id, app.Name, DeploymentTypes.ECS, ArtifactTypes.Docker))
	time.Sleep(3 * time.Second)

	t.Run("helm", testServiceSerializationWithAdditionalTests(app.Id, app.Name, DeploymentTypes.Helm, ArtifactTypes.Docker, func(t *testing.T, svc *Service) {
		require.Equal(t, HelmVersions.V2, svc.HelmVersion)
	}))
	time.Sleep(3 * time.Second)

	t.Run("pcf", testServiceSerialization(app.Id, app.Name, DeploymentTypes.PCF, ArtifactTypes.PCF))
	time.Sleep(3 * time.Second)

	t.Run("winrm_iis_website", testServiceSerialization(app.Id, app.Name, DeploymentTypes.WinRM, ArtifactTypes.IISWebsite))
	time.Sleep(3 * time.Second)

	t.Run("custom", testServiceSerialization(app.Id, app.Name, DeploymentTypes.Custom, ArtifactTypes.Tar))
	time.Sleep(3 * time.Second)

	// Cleanup
	c.ApiClient.Applications().DeleteApplication(app.Id)
}

func testServiceSerialization(applicationId string, applicationName string, deploymentType string, artifactType string) func(t *testing.T) {
	return testServiceSerializationWithAdditionalTests(applicationId, applicationName, deploymentType, artifactType, nil)
}

func testServiceSerializationWithAdditionalTests(applicationId string, applicationName string, deploymentType string, artifactType string, additionalTests func(t *testing.T, serviceUnderTest *Service)) func(t *testing.T) {
	return func(t *testing.T) {
		// Create service
		serviceName := fmt.Sprintf("%s-%s", getSafeTestName(t.Name()), utils.RandStringBytes(4))
		service, err := createService(applicationName, serviceName, deploymentType, artifactType)
		require.NoError(t, err)
		require.NotNil(t, service)

		// Get content
		content, err := getDirectoryItemContent(applicationId, service.YamlFilePath)
		require.NoError(t, err)
		require.NotNil(t, content)

		// Parse yaml config
		item, err := content.ParseYamlContent()
		require.NoError(t, err)
		require.NotNil(t, item)

		// Verify
		svc := item.(*Service)
		require.Equal(t, deploymentType, svc.DeploymentType)
		require.Equal(t, serviceName, svc.Name)
		require.Equal(t, artifactType, svc.ArtifactType)
		require.Equal(t, ObjectTypes.Service, svc.Type)

		if additionalTests != nil {
			additionalTests(t, svc)
		}

	}

}

func testAWSLambdaSerialization(applicationId string, applicationName string) func(t *testing.T) {
	return func(t *testing.T) {
		// Create service
		serviceName := fmt.Sprintf("%s-%s", getSafeTestName(t.Name()), utils.RandStringBytes(4))
		service, err := createService(applicationName, serviceName, DeploymentTypes.AWSLambda, "")
		require.NoError(t, err)
		require.NotNil(t, service)

		// Get content
		content, err := getDirectoryItemContent(applicationId, service.YamlFilePath)
		require.NoError(t, err)
		require.NotNil(t, content)

		// Parse yaml config
		item, err := content.ParseYamlContent()
		require.NoError(t, err)
		require.NotNil(t, item)

		// Verify
		svc := item.(*Service)
		require.Equal(t, DeploymentTypes.AWSLambda, svc.DeploymentType)
		require.Equal(t, serviceName, svc.Name)
		require.Equal(t, ArtifactTypes.AWSLambda, svc.ArtifactType)
		require.Equal(t, "", svc.HelmVersion)
		require.Equal(t, ObjectTypes.Service, svc.Type)
	}

}

func TestThrottleRegex(t *testing.T) {
	message := "Too Many requests. Throttled. Max QPS: 2.5"
	found := throttledRegex.MatchString(message)
	require.True(t, found)
}

// Internal

var configureClientFunc sync.Once
var c *ConfigAsCodeClient

func getClient() *ConfigAsCodeClient {

	configureClientFunc.Do(func() {
		c = &ConfigAsCodeClient{
			ApiClient: &client.ApiClient{
				UserAgent:   "micahlmartin-harness-go-sdk-0.0.1",
				Endpoint:    client.DefaultApiUrl,
				AccountId:   os.Getenv(envvar.HarnessAccountId),
				APIKey:      os.Getenv(envvar.HarnessApiKey),
				BearerToken: os.Getenv(envvar.HarnessBearerToken),
				HTTPClient: &http.Client{
					Timeout: 10 * time.Second,
				},
			},
		}
	})

	return c
}

func createApplication(name string) (*client.Application, error) {
	c := getClient()
	input := &client.Application{
		Name: name,
	}
	return c.ApiClient.Applications().CreateApplication(input)
}

// func createKubernetesService(applicationName string, serviceName string) (*ConfigAsCodeItem, error) {
// 	service, err := ServiceFactory(serviceName, DeploymentTypes.Kubernetes, "")
// 	if err != nil {
// 		return nil, err
// 	}

// 	service.Description = "some description"
// 	return getClient().UpsertService(applicationName, serviceName, service)
// }

func createService(applicationName string, serviceName string, deploymentType string, artifactType string) (*ConfigAsCodeItem, error) {
	service, err := ServiceFactory(serviceName, deploymentType, artifactType)
	if err != nil {
		return nil, err
	}

	service.Description = "some description"
	return getClient().UpsertService(applicationName, serviceName, service)
}

// func createCodeDeployService(applicationName string, serviceName string) (*ConfigAsCodeItem, error) {
// 	return createService(applicationName, servi)
// 	service, err := ServiceFactory(serviceName, DeploymentTypes.AWSCodeDeploy, "")
// 	if err != nil {
// 		return nil, err
// 	}

// 	service.Description = "some description"
// 	return getClient().UpsertService(applicationName, serviceName, service)
// }

// func createAMIService(applicationName string, serviceName string) (*ConfigAsCodeItem, error) {
// 	service, err := ServiceFactory(serviceName, DeploymentTypes.AMI, "")
// 	if err != nil {
// 		return nil, err
// 	}

// 	service.Description = "some description"
// 	return getClient().UpsertService(applicationName, serviceName, service)
// }

// func createSSHService(applicationName string, serviceName string, artifactType string) (*ConfigAsCodeItem, error) {
// 	service, err := ServiceFactory(serviceName, DeploymentTypes.SSH, artifactType)
// 	if err != nil {
// 		return nil, err
// 	}

// 	service.Description = "some description"
// 	return getClient().UpsertService(applicationName, serviceName, service)
// }

func getDirectoryItemContent(applicationId string, yamlFilePath string) (*ConfigAsCodeItem, error) {
	c := getClient()

	rootItem, err := c.GetDirectoryTree(applicationId)
	if err != nil {
		return nil, err
	}

	serviceItem := FindConfigAsCodeItemByPath(rootItem, yamlFilePath)
	if serviceItem == nil {
		return nil, fmt.Errorf("could not find item at path '%s'", yamlFilePath)
	}

	// Test content
	return c.GetDirectoryItemContent(serviceItem.RestName, serviceItem.UUID, applicationId)
}

var safeTestNameRx = regexp.MustCompile("/")

func getSafeTestName(name string) string {
	return safeTestNameRx.ReplaceAllString(name, "_")
}
