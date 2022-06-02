package cd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTriggerGetWebhookUrl(t *testing.T) {
	c := getClient()

	url, err := c.TriggerClient.GetWebhookUrl("J6E7fQBUQO6AKnOqRzagHA", "test")
	require.NoError(t, err)
	require.NotNil(t, url)
}

func TestTriggerGetWebhookUrl_NotFound(t *testing.T) {
	c := getClient()

	url, err := c.TriggerClient.GetWebhookUrl("J6E7fQBUQO6AKnOqRzagHA", "test2")
	require.NoError(t, err)
	require.Nil(t, url)
}
