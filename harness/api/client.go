package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/harness-io/harness-go-sdk/harness"
	"github.com/harness-io/harness-go-sdk/harness/api/nextgen"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	AccountId   string
	APIKey      string
	BearerToken string
	Endpoint    string
	HTTPClient  *retryablehttp.Client
	UserAgent   string
	NGClient    *nextgen.APIClient
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
		AccountId:   helpers.EnvVars.HarnessAccountId.Get(),
		APIKey:      helpers.EnvVars.HarnessApiKey.Get(),
		BearerToken: helpers.EnvVars.HarnessBearerToken.Get(),
		Endpoint:    helpers.EnvVars.HarnessEndpoint.GetWithDefault(DefaultApiUrl),
		UserAgent:   userAgent,
		HTTPClient:  httpClient,
		NGClient: nextgen.NewAPIClient(&nextgen.Configuration{
			BasePath: helpers.EnvVars.HarnessNGEndpoint.GetWithDefault(DefaultNGApiUrl),
			DefaultHeader: map[string]string{
				helpers.EnvVars.HarnessNGApiKey.String(): helpers.EnvVars.HarnessNGApiKey.Get(),
			},
			UserAgent:  userAgent,
			HTTPClient: httpClient,
		}),
	}
}

func (client *Client) NewAuthorizedGetRequest(path string) (*retryablehttp.Request, error) {
	return client.NewAuthorizedRequest(path, http.MethodGet, nil)
}

func (client *Client) NewAuthorizedPostRequest(path string, rawBody interface{}) (*retryablehttp.Request, error) {
	return client.NewAuthorizedRequest(path, http.MethodPost, rawBody)
}

func (client *Client) NewAuthorizedDeleteRequest(path string) (*retryablehttp.Request, error) {
	return client.NewAuthorizedRequest(path, http.MethodDelete, nil)
}

func (client *Client) NewAuthorizedRequest(path string, method string, rawBody interface{}) (*retryablehttp.Request, error) {
	url := strings.Join([]string{client.Endpoint, path}, "")
	req, err := retryablehttp.NewRequest(method, url, rawBody)

	if err != nil {
		return nil, err
	}

	req.Header.Set(helpers.HTTPHeaders.UserAgent.String(), client.UserAgent)
	req.Header.Set(helpers.HTTPHeaders.ContentType.String(), helpers.HTTPHeaders.ApplicationJson.String())
	req.Header.Set(helpers.HTTPHeaders.Accept.String(), helpers.HTTPHeaders.ApplicationJson.String())
	req.Header.Set(helpers.HTTPHeaders.ApiKey.String(), client.APIKey)

	if client.BearerToken != "" {
		req.Header.Set(helpers.HTTPHeaders.Authorization.String(), fmt.Sprintf("Bearer %s", client.BearerToken))
	}

	return req, err
}

func getUserAgentString() string {
	return fmt.Sprintf("%s-%s", harness.SDKName, harness.SDKVersion)
}
