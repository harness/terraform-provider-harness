package internal

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/nextgen"
	openapi_client_nextgen "github.com/harness/harness-openapi-go-client/nextgen"
)

type Session struct {
	AccountId string
	Endpoint  string
	CDClient  *cd.ApiClient
	PLClient  *nextgen.APIClient
	Client    *openapi_client_nextgen.APIClient
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
