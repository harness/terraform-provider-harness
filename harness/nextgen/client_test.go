package nextgen

import (
	"context"
	"sync"
)

var configureClient sync.Once
var client *APIClient

func getClientWithContext() (*APIClient, context.Context) {
	configureClient.Do(func() {
		cfg := NewConfiguration()
		client = NewAPIClient(cfg)
	})

	ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: client.ApiKey})
	return client, ctx
}
