package caac

import (
	"net/http"
	"os"
	"regexp"
	"sync"
	"testing"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/micahlmartin/terraform-provider-harness/harness/envvar"
	"github.com/micahlmartin/terraform-provider-harness/harness/graphql"
	"github.com/stretchr/testify/require"
)

// func TestGetDirectoryTreeItem(t *testing.T) {

// 	// Setup
// 	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
// 	serviceName := fmt.Sprintf("service-%s-%s", t.Name(), utils.RandStringBytes(5))
// 	c := getClient()

// 	app, err := createApplication(appName)
// 	ks, _ := createService(app.Id, serviceName, DeploymentTypes.Kubernetes, "")
// 	rootItem, _ := c.GetDirectoryTree(app.Id)
// 	serviceItem := FindConfigAsCodeItemByPath(rootItem, ks.YamlFilePath)

// 	// Verify
// 	require.NoError(t, err)
// 	require.NotNil(t, serviceItem)
// 	require.Equal(t, ks.YamlFilePath, serviceItem.DirectoryPath.Path)

// 	// Cleanup
// 	err = c.ApiClient.Applications().DeleteApplication(app.Id)
// 	require.NoError(t, err)
// }

// func TestGetDirectoryItemContent(t *testing.T) {

// 	// Setup
// 	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
// 	serviceName := fmt.Sprintf("service-%s-%s", t.Name(), utils.RandStringBytes(5))
// 	c := getClient()

// 	// Create objects for testing
// 	app, _ := createApplication(appName)
// 	require.NotNil(t, app)
// 	ks, _ := createService(appName, serviceName, DeploymentTypes.Kubernetes, "")
// 	require.NotNil(t, ks)

// 	rootItem, _ := c.GetDirectoryTree(app.Id)
// 	require.NotNil(t, rootItem)
// 	serviceItem := FindConfigAsCodeItemByPath(rootItem, ks.YamlFilePath)
// 	require.NotNil(t, serviceItem)

// 	// Test content
// 	content, err := c.GetDirectoryItemContent(serviceItem.RestName, serviceItem.UUID, app.Id)
// 	require.NoError(t, err)
// 	require.NotNil(t, content)

// 	// Verify
// 	require.Equal(t, ks.YamlFilePath, serviceItem.DirectoryPath.Path)

// 	// Cleanup
// 	err = c.ApiClient.Applications().DeleteApplication(app.Id)
// 	require.NoError(t, err)
// }

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
			ApiClient: &graphql.ApiClient{
				UserAgent:   "micahlmartin-harness-go-sdk-0.0.1",
				Endpoint:    graphql.DefaultApiUrl,
				AccountId:   os.Getenv(envvar.HarnessAccountId),
				APIKey:      os.Getenv(envvar.HarnessApiKey),
				BearerToken: os.Getenv(envvar.HarnessBearerToken),
				HTTPClient: &retryablehttp.Client{
					RetryMax:     125,
					RetryWaitMin: 5 * time.Second,
					RetryWaitMax: 30 * time.Second,
					HTTPClient: &http.Client{
						Timeout: 30 * time.Second,
					},
					Backoff:    retryablehttp.DefaultBackoff,
					CheckRetry: retryablehttp.DefaultRetryPolicy,
				},
			},
		}
	})

	return c
}

func createApplication(name string) (*graphql.Application, error) {
	c := getClient()
	input := &graphql.Application{
		Name: name,
	}
	return c.ApiClient.Applications().CreateApplication(input)
}

func createService(applicationId string, serviceName string, deploymentType string, artifactType string) (*Service, error) {
	serviceInput, err := ServiceFactory(applicationId, serviceName, deploymentType, artifactType)
	if err != nil {
		return nil, err
	}

	serviceInput.ApplicationId = applicationId

	serviceInput.Description = "some description"
	return getClient().UpsertService(serviceInput)
}

// func getDirectoryItemContent(applicationId string, yamlFilePath string) (*ConfigAsCodeItem, error) {
// 	c := getClient()

// 	rootItem, err := c.GetDirectoryTree(applicationId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	serviceItem := FindConfigAsCodeItemByPath(rootItem, yamlFilePath)
// 	if serviceItem == nil {
// 		return nil, fmt.Errorf("could not find item at path '%s'", yamlFilePath)
// 	}

// 	// Test content
// 	return c.GetDirectoryItemContent(serviceItem.RestName, serviceItem.UUID, applicationId)
// }

var safeTestNameRx = regexp.MustCompile("/")

func getSafeTestName(name string) string {
	return safeTestNameRx.ReplaceAllString(name, "_")
}
