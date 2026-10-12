package provider

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type closeTracker struct {
	io.Reader
	closed bool
}

func (c *closeTracker) Close() error {
	c.closed = true
	return nil
}

func TestRetryHarnessNGUnavailable(t *testing.T) {
	ctx := context.Background()

	t.Run("retries 401 when NG is unreachable", func(t *testing.T) {
		body := `{"errors":["Error validating API key: Could not connect to NG. HTTP call failed with status code 503"]}`
		original := &closeTracker{Reader: strings.NewReader(body)}
		resp := &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       original,
		}

		retry, err := retryHarnessNGUnavailable(ctx, resp, nil)
		require.NoError(t, err)
		require.True(t, retry)
		require.True(t, original.closed)

		got, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, body, string(got))
	})

	t.Run("does not retry a real unauthorized response", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       io.NopCloser(strings.NewReader(`{"message":"Invalid credentials"}`)),
		}

		retry, err := retryHarnessNGUnavailable(ctx, resp, nil)
		require.NoError(t, err)
		require.False(t, retry)
	})

	t.Run("does not read or retry a successful response", func(t *testing.T) {
		body := `{"ok":true}`
		resp := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(body)),
		}

		retry, err := retryHarnessNGUnavailable(ctx, resp, nil)
		require.NoError(t, err)
		require.False(t, retry)

		got, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, body, string(got))
	})

	t.Run("still retries server errors", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: http.StatusServiceUnavailable,
			Body:       io.NopCloser(strings.NewReader(`unavailable`)),
		}

		retry, err := retryHarnessNGUnavailable(ctx, resp, nil)
		require.NoError(t, err)
		require.True(t, retry)
	})
}
