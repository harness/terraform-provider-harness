package client

import (
	"fmt"
	"net/http"

	"github.com/micahlmartin/terraform-provider-harness/internal/common"
)

type ApiClient struct {
	HTTPClient *http.Client
	Endpoint   string
	UserAgent  string
	APIKey     string
	ApiToken   string
	AccountId  string
}

// Creates a new unauthenticated HTTP request
func (client *ApiClient) NewRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", client.Endpoint, path), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set(common.HTTP_HEADER_USER_AGENT, client.UserAgent)
	return req, err
}

// Creates an HTTP request using the bearer Token for authentication
func (client *ApiClient) NewAuthorizedRequestWithBearerToken(path string) (*http.Request, error) {
	req, err := client.NewRequest(path)

	if err != nil {
		return nil, err
	}

	req.Header.Set(common.AUTHORIZATION_HEADER_FIELD, fmt.Sprintf("Bearer %s", client.ApiToken))
	return req, nil
}

// Creates an HTTP request using an API key for authentication
func (client *ApiClient) NewAuthorizedRequestWithApiKey(path string) (*http.Request, error) {
	req, err := client.NewRequest(path)

	if err != nil {
		return nil, err
	}

	req.Header.Set(common.HTTP_HEADER_X_API_KEY, client.APIKey)
	return req, nil
}
