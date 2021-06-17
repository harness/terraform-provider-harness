package api

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/harness-io/harness-go-sdk/harness/envvar"
	"github.com/harness-io/harness-go-sdk/harness/httphelpers"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/require"
)

func TestNewRequest(t *testing.T) {
	client := getClient()
	req, err := client.NewHTTPRequest(http.MethodGet, "some/path")

	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("%s/some/path", DefaultApiUrl), fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path))
	require.Equal(t, client.UserAgent, req.Header.Get(httphelpers.HeaderUserAgent))
}

func getClient() *Client {
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

func getUnauthorizedClient() *Client {
	return &Client{
		UserAgent: "micahlmartin-harness-go-sdk-0.0.1",
		Endpoint:  DefaultApiUrl,
		AccountId: os.Getenv(envvar.HarnessAccountId),
		APIKey:    "BAD_KEY",
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
