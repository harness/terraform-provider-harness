package sdk

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness"
	"github.com/harness-io/harness-go-sdk/harness/cd"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/harness-go-sdk/logging"
	"github.com/hashicorp/go-retryablehttp"
	log "github.com/sirupsen/logrus"
)

var (
	defaultRetryMax = 10
)

type Session struct {
	AccountId  string
	BaseURL    string
	CDClient   *cd.ApiClient
	Endpoint   string
	HTTPClient *retryablehttp.Client
	NGClient   *nextgen.APIClient
}

type SessionOptions struct {
	AccountId    string
	ApiKey       string
	DebugLogging bool
	Endpoint     string
	HTTPClient   *retryablehttp.Client
	Logger       *log.Logger
	NGApiKey     string
	UserAgent    string
}

// NewSession creates a new session with default settings.
func NewSession(opt *SessionOptions) (*Session, error) {
	opt.AccountId = utils.CoalesceStr(opt.AccountId, helpers.EnvVars.AccountId.Get())
	opt.ApiKey = utils.CoalesceStr(opt.ApiKey, helpers.EnvVars.ApiKey.Get())
	opt.Endpoint = utils.CoalesceStr(opt.Endpoint, helpers.EnvVars.Endpoint.GetWithDefault(utils.BaseUrl))
	opt.NGApiKey = utils.CoalesceStr(opt.NGApiKey, helpers.EnvVars.NGApiKey.Get())
	opt.UserAgent = utils.CoalesceStr(opt.UserAgent, fmt.Sprintf("%s-%s", harness.SDKName, harness.SDKVersion))

	if opt.Logger == nil {
		logger := logging.NewLogger()

		if opt.DebugLogging {
			logger.SetLevel(log.DebugLevel)
		}

		opt.Logger = logger
	}

	if opt.HTTPClient == nil {
		opt.HTTPClient = utils.GetDefaultHttpClient(opt.Logger)
	}

	cdClient, err := opt.GetCDClientFromOptions()
	if err != nil {
		return nil, err
	}

	ngClient := opt.GetNGClientFromOptions()

	session := &Session{
		AccountId: opt.AccountId,
		Endpoint:  opt.Endpoint,
		CDClient:  cdClient,
		NGClient:  ngClient,
	}

	return session, nil
}

func (opts *SessionOptions) GetCDClientFromOptions() (*cd.ApiClient, error) {
	return cd.NewClient(&cd.Config{
		AccountId:  opts.AccountId,
		Endpoint:   opts.Endpoint,
		APIKey:     opts.ApiKey,
		UserAgent:  opts.UserAgent,
		HTTPClient: opts.HTTPClient,
		Logger:     opts.Logger,
	})
}

func (s *SessionOptions) GetNGClientFromOptions() *nextgen.APIClient {
	return nextgen.NewAPIClient(&nextgen.Configuration{
		BasePath: s.Endpoint,
		DefaultHeader: map[string]string{
			helpers.HTTPHeaders.ApiKey.String(): s.NGApiKey,
		},
		UserAgent:  s.UserAgent,
		HTTPClient: s.HTTPClient,
		Logger:     s.Logger,
	})
}
