package internal

import (
	"context"

	"github.com/harness/harness-go-sdk/harness/split"
)

// SplitAPIClient wraps split.APIClient so Session can call WithAuthContext like other Harness SDK clients.
// The Split SDK does not attach context to outbound HTTP requests yet; this preserves a consistent
// session API and non-nil context for callers.
type SplitAPIClient struct {
	*split.APIClient
}

// NewSplitAPIClient returns a wrapper for Session, or nil if c is nil.
func NewSplitAPIClient(c *split.APIClient) *SplitAPIClient {
	if c == nil {
		return nil
	}
	return &SplitAPIClient{APIClient: c}
}

// WithAuthContext returns the underlying Split client and request context.
func (w *SplitAPIClient) WithAuthContext(ctx context.Context) (*split.APIClient, context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	if w == nil || w.APIClient == nil {
		return nil, ctx
	}
	return w.APIClient, ctx
}
