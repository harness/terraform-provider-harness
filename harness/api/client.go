package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/micahlmartin/terraform-provider-harness/harness/envvar"
	"github.com/micahlmartin/terraform-provider-harness/harness/httphelpers"
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

func New() *Client {
	return &Client{
		UserAgent:   "micahlmartin-harness-go-sdk-0.0.1",
		Endpoint:    DefaultApiUrl,
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

// // Creates a new unauthenticated HTTP request
// func (client *ApiClient) NewRequest(path string) (*http.Request, error) {
// 	return client.NewHTTPRequest(http.MethodGet, path)
// }

func (client *Client) NewHTTPRequest(method string, path string) (*retryablehttp.Request, error) {
	req, err := retryablehttp.NewRequest(method, fmt.Sprintf("%s/%s", client.Endpoint, path), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set(httphelpers.HeaderUserAgent, client.UserAgent)
	return req, err
}

// func (client *ApiClient) NewPostRequest(path string) (*http.Request, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", client.Endpoint, path), nil)
// 	if err = nil {}
// }

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
