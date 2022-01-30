package cd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/harness-go-sdk/logging"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	log "github.com/sirupsen/logrus"
)

type service struct {
	ApiClient *ApiClient
}

type Configuration struct {
	AccountId      string
	APIKey         string
	Endpoint       string
	HTTPClient     *retryablehttp.Client
	UserAgent      string
	DefaultHeaders map[string]string
	DebugLogging   bool
	Logger         *log.Logger
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
	Log                 *log.Logger
}

func NewClient(cfg *Configuration) *ApiClient {
	cfg.Endpoint = utils.CoalesceStr(cfg.Endpoint, helpers.EnvVars.Endpoint.GetWithDefault(utils.BaseUrl))

	validateConfig(cfg)

	if cfg.Logger == nil {
		cfg.Logger = logging.GetDefaultLogger(cfg.DebugLogging)
	}

	if cfg.HTTPClient == nil {
		cfg.HTTPClient = utils.GetDefaultHttpClient(cfg.Logger)
	}

	if cfg.DefaultHeaders == nil {
		cfg.DefaultHeaders = map[string]string{}
	}

	// Set default headers for all requests
	cfg.DefaultHeaders[helpers.HTTPHeaders.UserAgent.String()] = cfg.UserAgent
	cfg.DefaultHeaders[helpers.HTTPHeaders.Accept.String()] = helpers.HTTPHeaders.ApplicationJson.String()
	cfg.DefaultHeaders[helpers.HTTPHeaders.ContentType.String()] = helpers.HTTPHeaders.ApplicationJson.String()
	cfg.DefaultHeaders[helpers.HTTPHeaders.ApiKey.String()] = cfg.APIKey

	c := &ApiClient{}
	c.Log = cfg.Logger
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

func validateConfig(cfg *Configuration) {
	if cfg == nil {
		panic("Configuration is required")
	}

	if cfg.Endpoint == "" {
		panic("Endpoint must be set")
	}

	if cfg.APIKey == "" {
		panic("APIKey must be set")
	}
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

func (c *ApiClient) NewAuthorizedRequest(path string, method string, rawBody interface{}) (*retryablehttp.Request, error) {
	url := strings.Join([]string{c.Configuration.Endpoint, path}, "")
	req, err := retryablehttp.NewRequest(method, url, rawBody)

	if err != nil {
		return nil, err
	}

	for key, value := range c.Configuration.DefaultHeaders {
		req.Header.Set(key, value)
	}

	return req, err
}

func (c *ApiClient) getJson(req *retryablehttp.Request, obj interface{}) error {
	res, err := c.Configuration.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusUnauthorized {
		return errors.New("unauthorized")
	}

	defer res.Body.Close()

	// Unmarshal into our response object
	if err := json.NewDecoder(res.Body).Decode(obj); err != nil {
		return fmt.Errorf("error decoding response: %s", err)
	}

	if res.StatusCode == http.StatusUnauthorized {
		return errors.New("unauthorized")
	}

	return nil
}
