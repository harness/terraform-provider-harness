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

type Session struct {
	AccountId  string
	Endpoint   string
	BaseURL    string
	HTTPClient *retryablehttp.Client
	NGClient   *nextgen.APIClient
	CDClient   *cd.ApiClient
}

type SessionOptions struct {
	AccountId    string
	Endpoint     string
	HTTPClient   *retryablehttp.Client
	UserAgent    string
	DebugLogging bool
	Logger       *log.Logger
}

func NewSession(opt *SessionOptions) *Session {
	opt.UserAgent = utils.CoalesceStr(opt.UserAgent, getUserAgentString())
	opt.AccountId = utils.CoalesceStr(opt.AccountId, helpers.EnvVars.AccountId.Get())
	opt.Endpoint = utils.CoalesceStr(opt.Endpoint, helpers.EnvVars.Endpoint.GetWithDefault(utils.BaseUrl))

	if opt.Logger == nil {
		opt.Logger = logging.GetDefaultLogger(opt.DebugLogging)
	}

	if opt.HTTPClient == nil {
		opt.HTTPClient = utils.GetDefaultHttpClient(opt.Logger)
	}

	return &Session{
		AccountId: opt.AccountId,
		Endpoint:  opt.Endpoint,
		CDClient: cd.NewClient(&cd.Configuration{
			AccountId:  opt.AccountId,
			APIKey:     helpers.EnvVars.ApiKey.Get(),
			Endpoint:   opt.Endpoint,
			UserAgent:  opt.UserAgent,
			HTTPClient: opt.HTTPClient,
			Logger:     opt.Logger,
		}),
		NGClient: nextgen.NewAPIClient(&nextgen.Configuration{
			BasePath: fmt.Sprintf("%s%s", opt.Endpoint, utils.DefaultNGApiUrl),
			DefaultHeader: map[string]string{
				helpers.HTTPHeaders.ApiKey.String(): helpers.EnvVars.NGApiKey.Get(),
			},
			UserAgent:  opt.UserAgent,
			HTTPClient: opt.HTTPClient,
			Logger:     opt.Logger,
		}),
	}
}

func getUserAgentString() string {
	return fmt.Sprintf("%s-%s", harness.SDKName, harness.SDKVersion)
}
