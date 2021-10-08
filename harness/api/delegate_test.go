package api

import (
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/stretchr/testify/require"
)

func TestListDelegates(t *testing.T) {
	client := getClient()
	limit := 100
	offset := 0
	hasMore := true

	for hasMore {
		delegates, pagination, err := client.Delegates().ListDelegates(limit, offset)
		require.NoError(t, err, "Failed to list delegates: %s", err)
		require.NotEmpty(t, delegates, "No delegates found")
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(delegates) == limit
		offset += limit
	}
}

func TestGetDelegateByName(t *testing.T) {
	client := getClient()
	delegateName := "harness-delegate"

	delegate, err := client.Delegates().GetDelegateByName(delegateName)
	require.NoError(t, err, "Failed to get delegate: %s", err)
	require.NotNil(t, delegate, "Delegate should not be nil")
	require.Equal(t, delegateName, delegate.DelegateName, "Delegate name should be %s", delegateName)
}

func TestGetDelegateByName_NotFound(t *testing.T) {
	client := getClient()
	delegateName := "nodelegate"

	delegate, err := client.Delegates().GetDelegateByName(delegateName)
	require.NoError(t, err, "Failed to get delegate: %s", err)
	require.Nil(t, delegate, "Delegate should be nil")
}

func TestGetDelegateByStatus(t *testing.T) {
	client := getClient()
	delegateList, _, err := client.Delegates().GetDelegateWithFilters(1, 0, "", graphql.DelegateStatusList.Enabled, "")
	require.NoError(t, err, "Failed to get delegate: %s", err)
	require.GreaterOrEqual(t, len(delegateList), 1, "Delegate list should have at least 1 delegate")
}
