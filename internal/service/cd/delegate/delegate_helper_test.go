package delegate_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/harness/harness-go-sdk/harness/cd"
	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/delegate"
	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/stretchr/testify/require"
)

var defaultDelegateTimeout = time.Minute * 5

func getDelegateTimeout() time.Duration {
	if t := helpers.TestEnvVars.DelegateWaitTimeout.Get(); t != "" {
		timeout, err := time.ParseDuration(t)
		if err != nil {
			fmt.Printf("failed to parse delegate wait timeout: %s. Using default.", err)
			return defaultDelegateTimeout
		}
		return timeout
	}

	return defaultDelegateTimeout
}

func createDelegateContainer(t *testing.T, name string) *graphql.Delegate {
	ctx := context.Background()
	c := acctest.TestAccProvider.Meta().(*cd.ApiClient)

	cfg := &delegate.DockerDelegateConfig{
		AccountId:     c.Configuration.AccountId,
		AccountSecret: helpers.TestEnvVars.DelegateSecret.Get(),
		DelegateName:  name,
		ContainerName: name,
		ProfileId:     helpers.TestEnvVars.DelegateProfileId.Get(),
	}

	t.Logf("Starting delegate %s", name)
	_, err := delegate.RunDelegateContainer(ctx, cfg)
	require.NoError(t, err, "failed to create delegate container: %s", err)

	delegate, err := c.DelegateClient.WaitForDelegate(ctx, name, getDelegateTimeout())
	require.NoError(t, err, "failed to wait for delegate: %s", err)
	require.NotNil(t, delegate, "delegate should not be nil")

	return delegate
}

func deleteDelegate(t *testing.T, name string) {
	c := acctest.TestAccProvider.Meta().(*cd.ApiClient)
	delegate, err := c.DelegateClient.GetDelegateByName(name)
	require.NoError(t, err, "Failed to get delegate: %s", err)
	require.NotNil(t, delegate, "Delegate should not be nil")

	cli, err := client.NewClientWithOpts(client.FromEnv)
	require.NoError(t, err, "failed to create docker client: %s", err)

	err = cli.ContainerStop(context.Background(), name, nil)
	require.NoError(t, err, "failed to stop delegate container: %s", err)

	err = cli.ContainerRemove(context.Background(), name, types.ContainerRemoveOptions{})
	require.NoError(t, err, "failed to remove delegate container: %s", err)
}
