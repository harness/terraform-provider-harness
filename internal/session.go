package internal

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/code"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/policymgmt"
	"github.com/harness/harness-go-sdk/harness/utils"
	openapi_client_nextgen "github.com/harness/harness-openapi-go-client/nextgen"
)

type Session struct {
	AccountId  string
	Endpoint   string
	CDClient   *cd.ApiClient
	PLClient   *nextgen.APIClient
	Client     *openapi_client_nextgen.APIClient
	CodeClient code.APIClient
}

func (s *Session) GetPlatformClient() (*nextgen.APIClient, context.Context) {
	return s.GetPlatformClientWithContext(nil)
}

func (s *Session) GetPlatformClientWithContext(ctx context.Context) (*nextgen.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	return s.PLClient.WithAuthContext(ctx)
}

func (s *Session) GetClientWithContext(ctx context.Context) (*openapi_client_nextgen.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	return s.Client.WithAuthContext(ctx)
}

func (s *Session) GetPolicyManagementClient() *policymgmt.APIClient {
	c := policymgmt.NewAPIClient(
		policymgmt.NewConfiguration(),
	)
	c.ChangeBasePath(s.Endpoint + "/pm")
	return c
}

func (s *Session) GetCodeClientWithContext(ctx context.Context) (*code.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	cfg := code.NewConfiguration()

	key := utils.GetEnv("HARNESS_PLATFORM_API_KEY", "pat.UmPeatlmRVec3U9Ii8VgVg.65c4c63973966a2ad865e208.rgsCeFWad3jn7sM09nr7")

	cfg.AddDefaultHeader("X-Api-Key", key)
	cfg.BasePath = utils.GetEnv("HARNESS_ENDPOINT", "https://app.harness.io/gateway") + "/code/api/v1"

	// cfg.AddDefaultHeader("Cookie", fmt.Sprintf("token=%s", key))
	// cfg.BasePath = utils.GetEnv("HARNESS_ENDPOINT", "https://app.harness.io/gateway") + "/api/v1"

	return code.NewAPIClient(cfg), ctx
}
