package chaos

import "context"

func (c *APIClient) WithAuthContext(ctx context.Context) (*APIClient, context.Context) {
	authCtx := context.WithValue(ctx, ContextAPIKey, APIKey{Key: c.ApiKey})
	return c, authCtx
}
