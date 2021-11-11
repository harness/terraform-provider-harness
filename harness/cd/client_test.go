package cd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/require"
)

func TestNewRequest(t *testing.T) {
	client := getClient()
	req, err := client.NewAuthorizedGetRequest("/some/path")

	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("%s/some/path", utils.DefaultApiUrl), fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path))
	require.Equal(t, client.Configuration.UserAgent, req.Header.Get(helpers.HTTPHeaders.UserAgent.String()))
}

func getClient() *ApiClient {
	return NewClient(&Configuration{
		UserAgent:   "micahlmartin-harness-go-sdk-0.0.1",
		Endpoint:    utils.DefaultApiUrl,
		AccountId:   helpers.EnvVars.AccountId.Get(),
		APIKey:      helpers.EnvVars.ApiKey.Get(),
		BearerToken: helpers.EnvVars.BearerToken.Get(),
		HTTPClient: &retryablehttp.Client{
			RetryMax:     10,
			RetryWaitMin: 5 * time.Second,
			RetryWaitMax: 10 * time.Second,
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
			Logger:     log.New(os.Stderr, "", log.LstdFlags),
			Backoff:    retryablehttp.DefaultBackoff,
			CheckRetry: retryablehttp.DefaultRetryPolicy,
		},
	})
}

func GetUnauthorizedClient() *ApiClient {
	return NewClient(&Configuration{
		UserAgent: "micahlmartin-harness-go-sdk-0.0.1",
		Endpoint:  utils.DefaultApiUrl,
		AccountId: helpers.EnvVars.AccountId.Get(),
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
	})
}
