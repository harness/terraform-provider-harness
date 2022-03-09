package cd

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/delegate"
	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/utils"
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
	c := getClient()

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
	c := getClient()
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

func TestCreateDelegate(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	delegate := createDelegateContainer(t, name)

	defer func() {
		deleteDelegate(t, name)
	}()

	require.NotNil(t, delegate)
}

func TestApproveDelegate(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	delegate := createDelegateContainer(t, name)

	defer func() {
		deleteDelegate(t, name)
	}()

	require.NotNil(t, delegate, "Delegate should not be nil")

	c := getClient()

	require.Equal(t, delegate.Status, graphql.DelegateStatusTypes.WaitingForApproval.String(), "Delegate status should be PENDING")

	delegate, err := c.DelegateClient.UpdateDelegateApprovalStatus(&graphql.DelegateApprovalRejectInput{
		AccountId:        c.Configuration.AccountId,
		DelegateApproval: graphql.DelegateApprovalTypes.Activate,
		DelegateId:       delegate.UUID,
	})

	require.NoError(t, err, "Failed to update delegate approval status: %s", err)
	require.NotNil(t, delegate, "Delegate should not be nil")
	require.Equal(t, delegate.Status, graphql.DelegateStatusTypes.Enabled.String(), "Delegate status should be enabled")
}

func TestRejectDelegate(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	delegate := createDelegateContainer(t, name)

	defer func() {
		deleteDelegate(t, name)
	}()

	require.NotNil(t, delegate, "Delegate should not be nil")

	c := getClient()

	require.Equal(t, delegate.Status, graphql.DelegateStatusTypes.WaitingForApproval.String(), "Delegate status should be PENDING")

	delegate, err := c.DelegateClient.UpdateDelegateApprovalStatus(&graphql.DelegateApprovalRejectInput{
		AccountId:        c.Configuration.AccountId,
		DelegateApproval: graphql.DelegateApprovalTypes.Reject,
		DelegateId:       delegate.UUID,
	})

	require.NoError(t, err, "Failed to update delegate approval status: %s", err)
	require.NotNil(t, delegate, "Delegate should not be nil")
	require.Equal(t, delegate.Status, graphql.DelegateStatusTypes.Deleted.String(), "Delegate status should be enabled")
}

func TestListDelegates(t *testing.T) {
	client := getClient()
	limit := 100
	offset := 0
	hasMore := true

	for hasMore {
		delegates, pagination, err := client.DelegateClient.ListDelegates(limit, offset)
		require.NoError(t, err, "Failed to list delegates: %s", err)
		require.NotEmpty(t, delegates, "No delegates found")
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(delegates) == limit
		offset += limit
	}
}

func TestGetDelegateByName(t *testing.T) {
	client := getClient()

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	_ = createDelegateContainer(t, name)

	defer func() {
		deleteDelegate(t, name)
	}()

	delegate, err := client.DelegateClient.GetDelegateByName(name)
	require.NoError(t, err, "Failed to get delegate: %s", err)
	require.NotNil(t, delegate, "Delegate should not be nil")
	require.Equal(t, name, delegate.DelegateName, "Delegate name should be %s", name)
}

func TestGetDelegateByName_NotFound(t *testing.T) {
	client := getClient()
	delegateName := "nodelegate"

	delegate, err := client.DelegateClient.GetDelegateByName(delegateName)
	require.NoError(t, err, "Failed to get delegate: %s", err)
	require.Nil(t, delegate, "Delegate should be nil")
}

func TestGetDelegateByStatus(t *testing.T) {
	client := getClient()
	delegateList, _, err := client.DelegateClient.ListDelegatesWithFilters(1, 0, "", graphql.DelegateStatusList.Enabled, "")
	require.NoError(t, err, "Failed to get delegate: %s", err)
	require.GreaterOrEqual(t, len(delegateList), 1, "Delegate list should have at least 1 delegate")
}

func TestGetDelegateByHostName(t *testing.T) {
	client := getClient()
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	delegate := createDelegateContainer(t, name)

	defer func() {
		deleteDelegate(t, name)
	}()

	delegateLookup, err := client.DelegateClient.GetDelegateByHostName(delegate.HostName)
	require.NoError(t, err, "Failed to get delegate: %s", err)
	require.Equal(t, delegate.HostName, delegateLookup.HostName, "Delegate hostname should be %s", delegate.HostName)
}

func TestGetDelegateById(t *testing.T) {
	client := getClient()
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	delegate := createDelegateContainer(t, name)

	defer func() {
		deleteDelegate(t, name)
	}()

	delegateLookup, err := client.DelegateClient.GetDelegateById(delegate.UUID)
	require.NoError(t, err, "Failed to get delegate: %s", err)
	require.Equal(t, delegate.UUID, delegateLookup.UUID, "Delegate id should be %s", delegate.UUID)
}
