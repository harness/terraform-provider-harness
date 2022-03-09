package cd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/harness/harness-go-sdk/harness"
	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/harness-go-sdk/logging"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	log "github.com/sirupsen/logrus"
)

type service struct {
	ApiClient *ApiClient
}

type Config struct {
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
	Configuration       *Config
	ApplicationClient   *ApplicationClient
	CloudProviderClient *CloudProviderClient
	ConfigAsCodeClient  *ConfigAsCodeClient
	ConnectorClient     *ConnectorClient
	DelegateClient      *DelegateClient
	ExecutionClient     *ExecutionClient
	SecretClient        *SecretClient
	SSOClient           *SSOClient
	UserClient          *UserClient
	Log                 *log.Logger
}

func DefaultConfig() *Config {
	logger := logging.NewLogger()
	if helpers.EnvVars.DebugEnabled.Get() == "true" {
		logger.SetLevel(log.DebugLevel)
	}

	cfg := &Config{
		AccountId:  helpers.EnvVars.AccountId.Get(),
		APIKey:     helpers.EnvVars.ApiKey.Get(),
		Endpoint:   helpers.EnvVars.Endpoint.GetWithDefault(utils.BaseUrl),
		Logger:     logger,
		HTTPClient: utils.GetDefaultHttpClient(logger),
		UserAgent:  fmt.Sprintf("%s-%s", harness.SDKName, harness.SDKVersion),
	}

	return cfg
}

func NewClient(cfg *Config) (*ApiClient, error) {
	if cfg == nil {
		return nil, errors.New("config is is required")
	}

	if cfg.AccountId == "" {
		return nil, cfg.NewInvalidConfigError("AccountId", nil)
	}

	if cfg.APIKey == "" {
		return nil, cfg.NewInvalidConfigError("ApiKey", nil)
	}

	if cfg.Endpoint == "" {
		return nil, cfg.NewInvalidConfigError("Endpoint", nil)
	}

	if cfg.HTTPClient == nil {
		return nil, cfg.NewInvalidConfigError("Endpoint", nil)
	}

	// defaultHeaders
	if cfg.DefaultHeaders == nil {
		cfg.DefaultHeaders = make(map[string]string)
	}

	// Set default headers for all requests
	cfg.DefaultHeaders[helpers.HTTPHeaders.UserAgent.String()] = utils.CoalesceStr(cfg.UserAgent, fmt.Sprintf("%s-%s", harness.SDKName, harness.SDKVersion))
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
	c.ExecutionClient = (*ExecutionClient)(&c.common)
	c.SecretClient = (*SecretClient)(&c.common)
	c.SSOClient = (*SSOClient)(&c.common)
	c.UserClient = (*UserClient)(&c.common)

	return c, nil
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
