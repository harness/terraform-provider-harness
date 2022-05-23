package nextgen

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetUserGroupByName(t *testing.T) {
	c, ctx := getClientWithContext()
	id := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(6))

	ug := UserGroup{
		AccountIdentifier: c.AccountId,
		Identifier:        id,
		Name:              id,
	}

	resp, _, err := c.UserGroupApi.PostUserGroup(ctx, ug, c.AccountId, &UserGroupApiPostUserGroupOpts{})
	require.NoError(t, err)
	require.NotNil(t, resp.Data)

	defer func() {
		c.UserGroupApi.DeleteUserGroup(ctx, c.AccountId, id, &UserGroupApiDeleteUserGroupOpts{})
	}()

	result, err := c.UserGroupApi.GetUserGroupByName(ctx, c.AccountId, id, &UserGroupApiGetUserGroupByNameOpts{})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, result.Name, id)
}
