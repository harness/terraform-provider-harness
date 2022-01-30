package utils

import (
	"github.com/harness-io/harness-go-sdk/harness"
	"github.com/harness-io/harness-go-sdk/logging"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	log "github.com/sirupsen/logrus"
)

var (
	defaultRetryMax = 10
)

func GetDefaultHttpClient(logger *log.Logger) *retryablehttp.Client {
	httpClient := retryablehttp.NewClient()
	httpClient.Logger = retryablehttp.LeveledLogger(&logging.LeveledLogger{Logger: logger})
	httpClient.HTTPClient.Transport = logging.NewTransport(harness.SDKName, logger, cleanhttp.DefaultPooledClient().Transport)
	httpClient.RetryMax = defaultRetryMax
	return httpClient
}
