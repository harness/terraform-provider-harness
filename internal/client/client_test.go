package client

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/micahlmartin/terraform-provider-harness/internal/common"
	"github.com/micahlmartin/terraform-provider-harness/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T) {
	client := getClient()
	req, err := client.NewRequest("some/path")

	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%s/some/path", common.DEFAULT_API_URL), fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path))
	assert.Equal(t, client.UserAgent, req.Header.Get(common.HTTP_HEADER_USER_AGENT))
}

func getClient() *ApiClient {
	return &ApiClient{
		UserAgent: "micahlmartin-harness-go-sdk-0.0.1",
		Endpoint:  common.DEFAULT_API_URL,
		AccountId: testhelpers.APP_ID,
		APIKey:    testhelpers.API_KEY,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func getUnauthorizedClient() *ApiClient {
	return &ApiClient{
		UserAgent: "micahlmartin-harness-go-sdk-0.0.1",
		Endpoint:  common.DEFAULT_API_URL,
		AccountId: testhelpers.APP_ID,
		APIKey:    "AbcDEF$%^",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
