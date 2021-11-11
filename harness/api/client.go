package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/harness-io/harness-go-sdk/harness"
	"github.com/harness-io/harness-go-sdk/harness/cd"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	AccountId string
	Endpoint  string
	NGClient  *nextgen.APIClient
	CDClient  *cd.ApiClient
}

func NewClient() *Client {

	httpClient := &retryablehttp.Client{
		RetryMax:     10,
		RetryWaitMin: 5 * time.Second,
		RetryWaitMax: 10 * time.Second,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		Logger:     log.New(os.Stderr, "", log.LstdFlags),
		Backoff:    retryablehttp.DefaultBackoff,
		CheckRetry: retryablehttp.DefaultRetryPolicy,
	}

	userAgent := getUserAgentString()

	return &Client{
		AccountId: helpers.EnvVars.AccountId.Get(),
		Endpoint:  helpers.EnvVars.Endpoint.GetWithDefault(utils.DefaultApiUrl),
		CDClient: cd.NewClient(&cd.Configuration{
			AccountId:   helpers.EnvVars.AccountId.Get(),
			APIKey:      helpers.EnvVars.ApiKey.Get(),
			BearerToken: helpers.EnvVars.BearerToken.Get(),
			Endpoint:    helpers.EnvVars.Endpoint.GetWithDefault(utils.DefaultApiUrl),
			UserAgent:   userAgent,
			HTTPClient:  httpClient,
		}),
		NGClient: nextgen.NewAPIClient(&nextgen.Configuration{
			BasePath: helpers.EnvVars.NGEndpoint.GetWithDefault(DefaultNGApiUrl),
			DefaultHeader: map[string]string{
				helpers.EnvVars.NGApiKey.String(): helpers.EnvVars.NGApiKey.Get(),
			},
			UserAgent:  userAgent,
			HTTPClient: httpClient,
		}),
	}
}

func getUserAgentString() string {
	return fmt.Sprintf("%s-%s", harness.SDKName, harness.SDKVersion)
}
