package client

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/micahlmartin/terraform-provider-harness/internal/envvar"
	"github.com/micahlmartin/terraform-provider-harness/internal/httphelpers"
	"github.com/stretchr/testify/require"
)

func TestNewRequest(t *testing.T) {
	client := getClient()
	req, err := client.NewRequest("some/path")

	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf("%s/some/path", DefaultApiUrl), fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path))
	require.Equal(t, client.UserAgent, req.Header.Get(httphelpers.HeaderUserAgent))
}

func getClient() *ApiClient {
	return &ApiClient{
		UserAgent: "micahlmartin-harness-go-sdk-0.0.1",
		Endpoint:  DefaultApiUrl,
		AccountId: os.Getenv(envvar.HarnessAccountId),
		APIKey:    os.Getenv(envvar.HarnessApiKey),
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func getUnauthorizedClient() *ApiClient {
	return &ApiClient{
		UserAgent: "micahlmartin-harness-go-sdk-0.0.1",
		Endpoint:  DefaultApiUrl,
		AccountId: os.Getenv(envvar.HarnessAccountId),
		APIKey:    "BAD_KEY",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
