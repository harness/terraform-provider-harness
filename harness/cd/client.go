package cd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/helpers"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

type service struct {
	ApiClient *ApiClient
}

type Configuration struct {
	AccountId   string
	APIKey      string
	BearerToken string
	Endpoint    string
	HTTPClient  *retryablehttp.Client
	UserAgent   string
}

type ApiClient struct {
	common              service // Reuse a single struct instead of allocating one for each service on the heap.
	Configuration       *Configuration
	ApplicationClient   *ApplicationClient
	CloudProviderClient *CloudProviderClient
	ConfigAsCodeClient  *ConfigAsCodeClient
	ConnectorClient     *ConnectorClient
	DelegateClient      *DelegateClient
	SecretClient        *SecretClient
	SSOClient           *SSOClient
	UserClient          *UserClient
}

func NewClient(cfg *Configuration) *ApiClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = retryablehttp.NewClient()
	}

	c := &ApiClient{}
	c.Configuration = cfg
	c.common.ApiClient = c

	// API Services
	c.ApplicationClient = (*ApplicationClient)(&c.common)
	c.CloudProviderClient = (*CloudProviderClient)(&c.common)
	c.ConfigAsCodeClient = (*ConfigAsCodeClient)(&c.common)
	c.ConnectorClient = (*ConnectorClient)(&c.common)
	c.DelegateClient = (*DelegateClient)(&c.common)
	c.SecretClient = (*SecretClient)(&c.common)
	c.SSOClient = (*SSOClient)(&c.common)
	c.UserClient = (*UserClient)(&c.common)

	return c
}

func (client *ApiClient) NewAuthorizedGetRequest(path string) (*retryablehttp.Request, error) {
	return client.NewAuthorizedRequest(path, http.MethodGet, nil)
}

func (client *ApiClient) NewAuthorizedPostRequest(path string, rawBody interface{}) (*retryablehttp.Request, error) {
	return client.NewAuthorizedRequest(path, http.MethodPost, rawBody)
}

func (client *ApiClient) NewAuthorizedDeleteRequest(path string) (*retryablehttp.Request, error) {
	return client.NewAuthorizedRequest(path, http.MethodDelete, nil)
}

func (client *ApiClient) NewAuthorizedRequest(path string, method string, rawBody interface{}) (*retryablehttp.Request, error) {
	url := strings.Join([]string{client.Configuration.Endpoint, path}, "")
	req, err := retryablehttp.NewRequest(method, url, rawBody)

	if err != nil {
		return nil, err
	}

	req.Header.Set(helpers.HTTPHeaders.UserAgent.String(), client.Configuration.UserAgent)
	req.Header.Set(helpers.HTTPHeaders.ContentType.String(), helpers.HTTPHeaders.ApplicationJson.String())
	req.Header.Set(helpers.HTTPHeaders.Accept.String(), helpers.HTTPHeaders.ApplicationJson.String())
	req.Header.Set(helpers.HTTPHeaders.ApiKey.String(), client.Configuration.APIKey)

	if client.Configuration.BearerToken != "" {
		req.Header.Set(helpers.HTTPHeaders.Authorization.String(), fmt.Sprintf("Bearer %s", client.Configuration.BearerToken))
	}

	return req, err
}
