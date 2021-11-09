package cd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func createUser(email string) (*graphql.User, error) {
	return createUserWithGroups(email, nil)
}

func createUserWithGroups(email string, groupIds []string) (*graphql.User, error) {
	input := &graphql.CreateUserInput{
		Name:         "test",
		Email:        email,
		UserGroupIds: groupIds,
	}

	c := getClient()
	return c.UserClient.CreateUser(input)
}

func TestCreateUser(t *testing.T) {

	expectedEmail := fmt.Sprintf("%s_%s@example.com", t.Name(), utils.RandStringBytes(4))
	user, err := createUser(expectedEmail)

	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "test", user.Name)
	require.Equal(t, strings.ToLower(expectedEmail), user.Email)

	c := getClient()
	err = c.UserClient.DeleteUser(user.Id)
	require.NoError(t, err)
}

func TestCreateUserWithGroups(t *testing.T) {
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	expectedEmail := fmt.Sprintf("%s@example.com", expectedName)

	group, err := createUserGroup(expectedName)
	require.NoError(t, err)
	require.NotNil(t, group)

	user, err := createUserWithGroups(expectedEmail, []string{group.Id})

	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "test", user.Name)
	require.Equal(t, strings.ToLower(expectedEmail), user.Email)

	c := getClient()
	err = c.UserClient.DeleteUser(user.Id)
	require.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {

	expectedEmail := fmt.Sprintf("%s_%s@example.com", t.Name(), utils.RandStringBytes(4))
	user, err := createUser(expectedEmail)

	require.NoError(t, err)
	require.NotNil(t, user)

	c := getClient()

	_, err = c.UserClient.UpdateUser(&graphql.UpdateUserInput{
		Id:   user.Id,
		Name: "updated_name",
	})
	require.NoError(t, err)

	updateUser, err := c.UserClient.GetUserByEmail(expectedEmail)
	require.NoError(t, err)
	require.NotNil(t, updateUser)
	require.Equal(t, "updated_name", updateUser.Name)

	err = c.UserClient.DeleteUser(user.Id)
	require.NoError(t, err)
}

func TestGetUserByName(t *testing.T) {
	t.Skip("Looking up user by name currently has a bug https://harness.atlassian.net/browse/SWAT-5023")
	c := getClient()

	expectedEmail := fmt.Sprintf("%s_%s@example.com", t.Name(), utils.RandStringBytes(4))
	expectedUser, err := createUser(expectedEmail)
	require.NoError(t, err)
	require.NotNil(t, expectedUser)

	user, err := c.UserClient.GetUserByName(expectedUser.Name)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, expectedUser.Name, user.Name)
	require.Equal(t, expectedUser.Email, user.Email)

	err = c.UserClient.DeleteUser(expectedUser.Id)
	require.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	c := getClient()

	expectedEmail := fmt.Sprintf("%s_%s@example.com", t.Name(), utils.RandStringBytes(4))
	expectedUser, err := createUser(expectedEmail)
	require.NoError(t, err)
	require.NotNil(t, expectedUser)

	user, err := c.UserClient.GetUserByEmail(expectedUser.Email)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, expectedUser.Name, user.Name)
	require.Equal(t, expectedUser.Email, user.Email)

	err = c.UserClient.DeleteUser(expectedUser.Id)
	require.NoError(t, err)
}

func TestGetUserById(t *testing.T) {
	t.Skip("Looking up user by name currently has a bug https://harness.atlassian.net/browse/SWAT-5023")

	c := getClient()

	expectedEmail := fmt.Sprintf("%s_%s@example.com", t.Name(), utils.RandStringBytes(4))
	expectedUser, err := createUser(expectedEmail)
	require.NoError(t, err)
	require.NotNil(t, expectedUser)

	user, err := c.UserClient.GetUserById(expectedUser.Id)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, expectedUser.Name, user.Name)
	require.Equal(t, expectedUser.Email, user.Email)

	err = c.UserClient.DeleteUser(expectedUser.Id)
	require.NoError(t, err)
}

func TestListUsers(t *testing.T) {
	client := getClient()
	limit := 10
	offset := 0
	hasMore := true

	for hasMore {
		users, pagination, err := client.UserClient.ListUsers(limit, offset)
		require.NoError(t, err, "Failed to list users: %s", err)
		require.NotEmpty(t, users, "No users found")
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(users) == limit
		offset += 1
	}
}
