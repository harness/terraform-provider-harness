package cd

import (
	"fmt"
	"strings"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
)

type UserClient struct {
	ApiClient *ApiClient
}

func (c *UserClient) CreateUser(input *graphql.CreateUserInput) (*graphql.User, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($input: CreateUserInput!) {
			createUser(input: $input) {
				user {
					%[1]s
				}
			}
		}`, userFields),
		Variables: map[string]interface{}{
			"input": &input,
		},
	}

	res := &struct {
		CreateUser struct {
			User graphql.User
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateUser.User, nil
}

func (c *UserClient) UpdateUser(input *graphql.UpdateUserInput) (*graphql.User, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($input: UpdateUserInput!) {
			updateUser(input: $input) {
				user {
					%[1]s
				}
			}
		}`, userFields),
		Variables: map[string]interface{}{
			"input": &input,
		},
	}

	res := &struct {
		UpdateUser struct {
			User graphql.User
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.UpdateUser.User, nil
}

func (c *UserClient) DeleteUser(id string) error {

	query := &GraphQLQuery{
		Query: `mutation($input: DeleteUserInput!) {
			deleteUser(input: $input) {
				clientMutationId
			}
		}`,
		Variables: map[string]interface{}{
			"input": struct {
				Id string `json:"id"`
			}{
				Id: id,
			},
		},
	}

	res := &struct {
		DeleteUser struct {
			ClientMutationId string
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return err
	}

	return nil
}

func (c *UserClient) GetUserByName(name string) (*graphql.User, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			userByName(name: "%[1]s") {
				%[2]s
			}
		}`, name, userFields),
	}

	res := struct {
		UserByName *graphql.User
	}{}

	if err := c.ApiClient.ExecuteGraphQLQuery(query, &res); err != nil {
		if strings.Contains(err.Error(), helpers.USER_NOT_FOUND) {
			return nil, nil
		}
		return nil, err
	}

	return res.UserByName, nil
}

func (c *UserClient) GetUserByEmail(email string) (*graphql.User, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			userByEmail(email: "%[1]s") {
				%[2]s
			}
		}`, strings.ToLower(email), userFields),
	}

	res := struct {
		UserByEmail *graphql.User
	}{}

	if err := c.ApiClient.ExecuteGraphQLQuery(query, &res); err != nil {
		if strings.Contains(err.Error(), helpers.USER_NOT_FOUND) {
			return nil, nil
		}
		return nil, err
	}

	return res.UserByEmail, nil
}

func (c *UserClient) GetUserById(id string) (*graphql.User, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			user(id: "%[1]s") {
				%[2]s
			}
		}`, id, userFields),
	}

	res := struct {
		User *graphql.User
	}{}

	if err := c.ApiClient.ExecuteGraphQLQuery(query, &res); err != nil {
		if strings.Contains(err.Error(), helpers.USER_NOT_FOUND) {
			return nil, nil
		}
		return nil, err
	}

	return res.User, nil
}

func (c *UserClient) ListUsers(limit int, offset int) ([]*graphql.User, *graphql.PageInfo, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($limit: Int!, $offset: Int) {
			users(limit: $limit, offset: $offset) {
				nodes {
					%[1]s
				}
				%[2]s
			}
		}`, userFields, paginationFields),
		Variables: map[string]interface{}{
			"limit":  limit,
			"offset": offset,
		},
	}

	res := struct {
		Users struct {
			Nodes    []*graphql.User
			PageInfo *graphql.PageInfo
		}
	}{}

	err := c.ApiClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, nil, err
	}

	return res.Users.Nodes, res.Users.PageInfo, nil
}

var userFields = `
	id
	email
	isEmailVerified
	isImportedFromIdentityProvider
	isPasswordExpired
	isTwoFactorAuthenticationEnabled
	isUserLocked
	name
`
