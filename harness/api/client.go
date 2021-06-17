package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/harness-io/harness-go-sdk/harness"
	"github.com/harness-io/harness-go-sdk/harness/envvar"
	"github.com/harness-io/harness-go-sdk/harness/httphelpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	HTTPClient  *retryablehttp.Client
	Endpoint    string
	UserAgent   string
	APIKey      string
	ApiToken    string
	AccountId   string
	BearerToken string
}

func NewClient() *Client {
	return &Client{
		UserAgent:   getUserAgentString(),
		Endpoint:    utils.GetEnv(envvar.HarnessEndpoint, DefaultApiUrl),
		AccountId:   os.Getenv(envvar.HarnessAccountId),
		APIKey:      os.Getenv(envvar.HarnessApiKey),
		BearerToken: os.Getenv(envvar.HarnessBearerToken),
		HTTPClient: &retryablehttp.Client{
			RetryMax:     10,
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 10 * time.Second,
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
			Backoff:    retryablehttp.DefaultBackoff,
			CheckRetry: retryablehttp.DefaultRetryPolicy,
		},
	}
}

func (client *Client) NewHTTPRequest(method string, path string) (*retryablehttp.Request, error) {
	req, err := retryablehttp.NewRequest(method, fmt.Sprintf("%s/%s", client.Endpoint, path), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set(httphelpers.HeaderUserAgent, client.UserAgent)
	return req, err
}

// Creates an HTTP request using the bearer Token for authentication
func (client *Client) NewAuthorizedRequestWithBearerToken(path string) (*retryablehttp.Request, error) {
	req, err := client.NewHTTPRequest(http.MethodGet, path)

	if err != nil {
		return nil, err
	}

	req.Header.Set(httphelpers.HeaderAuthorization, fmt.Sprintf("Bearer %s", client.ApiToken))
	return req, nil
}

// Creates an HTTP request using an API key for authentication
func (client *Client) NewAuthorizedRequestWithApiKey(path string) (*retryablehttp.Request, error) {
	req, err := client.NewHTTPRequest(http.MethodGet, path)

	if err != nil {
		return nil, err
	}

	req.Header.Set(httphelpers.HeaderApiKey, client.APIKey)
	return req, nil
}

func getUserAgentString() string {
	return fmt.Sprintf("%s-%s", harness.SDKName, harness.SDKVersion)
}
