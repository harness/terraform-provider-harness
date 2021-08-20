package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/jinzhu/copier"
)

func (c *UserClient) CreateUserGroup(user *graphql.UserGroup) (*graphql.UserGroup, error) {

	input := &graphql.CreateUserGroupInput{}

	if err := copier.Copy(input, user); err != nil {
		return nil, err
	}

	if user.SAMLSettings != nil {
		input.SSOSetting = &graphql.SSOSettingInput{
			SAMLSettings: user.SAMLSettings,
		}
	}

	if user.LDAPSettings != nil {
		input.SSOSetting = &graphql.SSOSettingInput{
			LDAPSettings: user.LDAPSettings,
		}
	}

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($input: CreateUserGroupInput!) {
			createUserGroup(input: $input) {
				userGroup {
					%[1]s
				}
			}
		}`, userGroupFields),
		Variables: map[string]interface{}{
			"input": &input,
		},
	}

	res := &struct {
		CreateUserGroup struct {
			UserGroup graphql.UserGroup
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.CreateUserGroup.UserGroup, nil
}

func (c *UserClient) UpdateUserGroup(user *graphql.UserGroup) (*graphql.UserGroup, error) {

	input := &graphql.UpdateUserGroupInput{}

	if err := copier.Copy(input, user); err != nil {
		return nil, err
	}

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($input: UpdateUserGroupInput!) {
			updateUserGroup(input: $input) {
				userGroup {
					%[1]s
				}
			}
		}`, userGroupFields),
		Variables: map[string]interface{}{
			"input": &input,
		},
	}

	res := &struct {
		UpdateUserGroup struct {
			UserGroup graphql.UserGroup
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, err
	}

	return &res.UpdateUserGroup.UserGroup, nil
}

func (c *UserClient) GetUserGroupById(id string) (*graphql.UserGroup, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			userGroup(userGroupId: "%[1]s") {
				%[2]s
			}
		}`, id, userGroupFields),
	}

	res := struct {
		UserGroup *graphql.UserGroup
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, err
	}

	return res.UserGroup, nil
}

func (c *UserClient) GetUserGroupByName(name string) (*graphql.UserGroup, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`{
			userGroupByName(name: "%[1]s") {
				%[2]s
			}
		}`, name, userGroupFields),
	}

	res := struct {
		UserGroupByName *graphql.UserGroup
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, err
	}

	return res.UserGroupByName, nil
}

func (c *UserClient) DeleteUserGroup(id string) error {

	query := &GraphQLQuery{
		Query: `mutation($input: DeleteUserGroupInput!) {
			deleteUserGroup(input: $input) {
				clientMutationId
			}
		}`,
		Variables: map[string]interface{}{
			"input": struct {
				UserGroupId string `json:"userGroupId"`
			}{
				UserGroupId: id,
			},
		},
	}

	res := &struct {
		DeleteUserGroup struct {
			ClientMutationId string
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return err
	}

	return nil
}

func (c *UserClient) ListUserGroups(limit int, offset int) ([]*graphql.UserGroup, *graphql.PageInfo, error) {

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($limit: Int!, $offset: Int) {
			userGroups(limit: $limit, offset: $offset) {
				nodes {
					%[1]s
				}
				%[2]s
			}
		}`, userGroupFields, paginationFields),
		Variables: map[string]interface{}{
			"limit":  limit,
			"offset": offset,
		},
	}

	res := struct {
		UserGroups struct {
			Nodes    []*graphql.UserGroup
			PageInfo *graphql.PageInfo
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return nil, nil, err
	}

	return res.UserGroups.Nodes, res.UserGroups.PageInfo, nil
}

func (c *UserClient) IsUserInGroup(userId string, groupId string) (bool, error) {
	limit := 10
	offset := 0
	hasMore := true

	for hasMore {
		groups, _, err := c.ListGroupMembershipByUserId(userId, limit, offset)
		if err != nil {
			return false, err
		}

		for _, group := range groups {
			if group.Id == groupId {
				return true, nil
			}
		}

		hasMore = len(groups) == limit
		offset += limit
	}

	return false, nil
}

func (c *UserClient) ListGroupMembershipByUserId(userId string, limit int, offset int) ([]*graphql.UserGroup, *graphql.PageInfo, error) {
	query := &GraphQLQuery{
		Query: fmt.Sprintf(`query($limit: Int!, $offset: Int, $id: String!) {
				user(id: $id) {
					id
					email
					userGroups(limit: $limit, offset: $offset) {
						nodes {
							%[2]s
						}	
						%[3]s
					}
				}
			}`, userId, userGroupFields, paginationFields),
		Variables: map[string]interface{}{
			"id":     userId,
			"limit":  limit,
			"offset": offset,
		},
	}

	res := struct {
		User struct {
			Id         string
			Email      string
			UserGroups struct {
				Nodes    []*graphql.UserGroup
				PageInfo *graphql.PageInfo
			}
		}
	}{}

	if err := c.APIClient.ExecuteGraphQLQuery(query, &res); err != nil {
		return nil, nil, err
	}

	return res.User.UserGroups.Nodes, res.User.UserGroups.PageInfo, nil
}

func (c *UserClient) RemoveUserFromGroup(userId string, groupId string) (bool, error) {
	input := &graphql.AddUserToUserGroupInput{
		UserId:      userId,
		UserGroupId: groupId,
	}

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($input: RemoveUserFromUserGroupInput!) {
			removeUserFromUserGroup(input: $input) {
				userGroup {
					%[1]s
				}
			}
		}`, userGroupFields),
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	res := &struct {
		RemoveUserFromUserGroup struct {
			UserGroup graphql.UserGroup
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *UserClient) AddUserToGroup(userId string, groupId string) (bool, error) {

	input := &graphql.AddUserToUserGroupInput{
		UserId:      userId,
		UserGroupId: groupId,
	}

	query := &GraphQLQuery{
		Query: fmt.Sprintf(`mutation($input: AddUserToUserGroupInput!) {
			addUserToUserGroup(input: $input) {
				userGroup {
					%[1]s
				}
			}
		}`, userGroupFields),
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	res := &struct {
		AddUserToUserGroup struct {
			UserGroup graphql.UserGroup
		}
	}{}

	err := c.APIClient.ExecuteGraphQLQuery(query, &res)

	if err != nil {
		return false, err
	}

	return true, nil
}

var userGroupFields = `
	id
	description
	importedByScim
	isSSOLinked
	name
	notificationSettings {
		groupEmailAddresses
		microsoftTeamsWebhookUrl
		sendMailToNewMembers
		sendNotificationToMembers
		slackNotificationSetting {
			slackChannelName
			slackWebhookURL
		}
	}
	permissions {
		accountPermissions {
			accountPermissionTypes
		} 
		appPermissions {
			actions 
			permissionType
			applications {
				appIds
				filterType
			}
			deployments {
				envIds
				filterTypes
			}
			environments {
				envIds
				filterTypes
			}
			pipelines {
				envIds
				filterTypes
			}
			provisioners {
				filterType
				provisionerIds
			}
			services {
				filterType
				serviceIds
			}
			workflows {
				envIds
				filterTypes
			}
		}
	}
	ldapSettings: ssoSetting {
		... on LDAPSettings {
			groupDN
			groupName
			ssoProviderId
		} 
	}
	samlSettings: ssoSetting {
		... on SAMLSettings {
			groupName
			ssoProviderId
		} 
	}
`
