package caac

import (
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/micahlmartin/terraform-provider-harness/harness/envvar"
	"github.com/micahlmartin/terraform-provider-harness/harness/graphql"
)

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

var safeTestNameRx = regexp.MustCompile("/")

func getSafeTestName(name string) string {
	return safeTestNameRx.ReplaceAllString(name, "_")
}
