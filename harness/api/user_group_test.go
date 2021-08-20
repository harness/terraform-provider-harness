package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

var testUserGroup = graphql.UserGroup{
	Description: "description",
	Name:        "NOT_SET",
	NotificationSettings: &graphql.NotificationSettings{
		GroupEmailAddresses:      []string{"testuser@example.com"},
		MicrosoftTeamsWebhookUrl: "https://example.com/webhook",

		// BUG: https://harness.atlassian.net/browse/SWAT-5069
		// PagerDutyIntegrationKey:   "test-integration-key",

		SendMailToNewMembers:      true,
		SendNotificationToMembers: true,
		SlackNotificationSetting: &graphql.SlackNotificationSetting{
			SlackWebhookUrl:  "https://example.com/webhook",
			SlackChannelName: "test-channel",
		},
	},
	Permissions: &graphql.UserGroupPermissions{
		AccountPermissions: &graphql.AccountPermissions{
			AccountPermissionTypes: []graphql.AccountPermissionType{
				// graphql.AccountPermissionTypes.ADMINISTER_CE,
				graphql.AccountPermissionTypes.ADMINISTER_OTHER_ACCOUNT_FUNCTIONS,
				graphql.AccountPermissionTypes.CREATE_AND_DELETE_APPLICATION,
				graphql.AccountPermissionTypes.CREATE_CUSTOM_DASHBOARDS,
				graphql.AccountPermissionTypes.MANAGE_ALERT_NOTIFICATION_RULES,
				graphql.AccountPermissionTypes.MANAGE_API_KEYS,
				graphql.AccountPermissionTypes.MANAGE_APPLICATION_STACKS,
				graphql.AccountPermissionTypes.MANAGE_AUTHENTICATION_SETTINGS,
				graphql.AccountPermissionTypes.MANAGE_CLOUD_PROVIDERS,
				graphql.AccountPermissionTypes.MANAGE_CONFIG_AS_CODE,
				graphql.AccountPermissionTypes.MANAGE_CONNECTORS,
				graphql.AccountPermissionTypes.MANAGE_CUSTOM_DASHBOARDS,
				graphql.AccountPermissionTypes.MANAGE_DELEGATES,
				graphql.AccountPermissionTypes.MANAGE_DELEGATE_PROFILES,
				graphql.AccountPermissionTypes.MANAGE_DEPLOYMENT_FREEZES,
				graphql.AccountPermissionTypes.MANAGE_IP_WHITELIST,
				graphql.AccountPermissionTypes.MANAGE_PIPELINE_GOVERNANCE_STANDARDS,
				graphql.AccountPermissionTypes.MANAGE_RESTRICTED_ACCESS,
				graphql.AccountPermissionTypes.MANAGE_SECRETS,
				graphql.AccountPermissionTypes.MANAGE_SECRET_MANAGERS,
				graphql.AccountPermissionTypes.MANAGE_SSH_AND_WINRM,
				graphql.AccountPermissionTypes.MANAGE_TAGS,
				graphql.AccountPermissionTypes.MANAGE_TEMPLATE_LIBRARY,
				graphql.AccountPermissionTypes.MANAGE_USERS_AND_GROUPS,
				graphql.AccountPermissionTypes.MANAGE_USER_AND_USER_GROUPS_AND_API_KEYS,
				graphql.AccountPermissionTypes.READ_USERS_AND_GROUPS,
				graphql.AccountPermissionTypes.VIEW_AUDITS,
				// graphql.AccountPermissionTypes.VIEW_CE,
				graphql.AccountPermissionTypes.VIEW_USER_AND_USER_GROUPS_AND_API_KEYS,
			},
		},
		AppPermissions: []*graphql.AppPermission{
			{
				Actions: []graphql.Action{
					graphql.Actions.CREATE,
					graphql.Actions.DELETE,
					graphql.Actions.EXECUTE,
					graphql.Actions.EXECUTE_PIPELINE,
					graphql.Actions.EXECUTE_WORKFLOW,
					graphql.Actions.READ,
					graphql.Actions.ROLLBACK_WORKFLOW,
					graphql.Actions.UPDATE,
				},
				PermissionType: graphql.AppPermissionTypes.All,
				Applications: &graphql.AppFilter{
					FilterType: graphql.FilterTypes.All,
				},
			},
			{
				Actions: []graphql.Action{
					graphql.Actions.CREATE,
					graphql.Actions.DELETE,
					graphql.Actions.READ,
					graphql.Actions.UPDATE,
				},
				PermissionType: graphql.AppPermissionTypes.All,
				Applications: &graphql.AppFilter{
					FilterType: graphql.FilterTypes.All,
				},
			},
			{
				Actions: []graphql.Action{
					graphql.Actions.CREATE,
					graphql.Actions.DELETE,
					graphql.Actions.READ,
					graphql.Actions.UPDATE,
				},
				PermissionType: graphql.AppPermissionTypes.Workflow,
				Applications: &graphql.AppFilter{
					FilterType: graphql.FilterTypes.All,
				},
				Workflows: &graphql.WorkflowPermissionFilter{
					FilterTypes: []graphql.WorkflowPermissionFilterType{
						graphql.WorkflowPermissionFilterTypes.NonProductionWorkflows,
						graphql.WorkflowPermissionFilterTypes.ProductionWorkflows,
						graphql.WorkflowPermissionFilterTypes.WorkflowTemplates,
					},
				},
			},
			{
				Actions: []graphql.Action{
					graphql.Actions.CREATE,
					graphql.Actions.DELETE,
					graphql.Actions.READ,
					graphql.Actions.UPDATE,
				},
				PermissionType: graphql.AppPermissionTypes.Pipeline,
				Applications: &graphql.AppFilter{
					FilterType: graphql.FilterTypes.All,
				},
				Pipelines: &graphql.PipelinePermissionFilter{
					FilterTypes: []graphql.PipelinePermissionFilterType{
						graphql.PipelinePermissionFilterTypes.NonProductionPipelines,
						graphql.PipelinePermissionFilterTypes.ProductionPipelines,
					},
				},
			},
			{
				Actions: []graphql.Action{
					graphql.Actions.CREATE,
					graphql.Actions.DELETE,
					graphql.Actions.READ,
					graphql.Actions.UPDATE,
				},
				PermissionType: graphql.AppPermissionTypes.Env,
				Applications: &graphql.AppFilter{
					FilterType: graphql.FilterTypes.All,
				},
				Environments: &graphql.EnvPermissionFilter{
					FilterTypes: []graphql.EnvFilterType{
						graphql.EnvFilterTypes.NonProductionEnvironments,
						graphql.EnvFilterTypes.ProductionEnvironments,
					},
				},
			},
			{
				Actions: []graphql.Action{
					graphql.Actions.READ,
					graphql.Actions.ROLLBACK_WORKFLOW,
					graphql.Actions.EXECUTE_PIPELINE,
					graphql.Actions.EXECUTE_WORKFLOW,
				},
				PermissionType: graphql.AppPermissionTypes.Deployment,
				Applications: &graphql.AppFilter{
					FilterType: graphql.FilterTypes.All,
				},
				Deployments: &graphql.DeploymentPermissionFilter{
					FilterTypes: []graphql.DeploymentPermissionFilterType{
						graphql.DeploymentPermissionFilterTypes.NonProductionEnvironments,
						graphql.DeploymentPermissionFilterTypes.ProductionEnvironments,
					},
				},
			},
			{
				Actions: []graphql.Action{
					graphql.Actions.CREATE,
					graphql.Actions.DELETE,
					graphql.Actions.READ,
					graphql.Actions.UPDATE,
				},
				PermissionType: graphql.AppPermissionTypes.Service,
				Applications: &graphql.AppFilter{
					FilterType: graphql.FilterTypes.All,
				},
				Services: &graphql.ServicePermissionFilter{
					FilterType: graphql.FilterTypes.All,
				},
			},
			{
				Actions: []graphql.Action{
					graphql.Actions.CREATE,
					graphql.Actions.DELETE,
					graphql.Actions.READ,
					graphql.Actions.UPDATE,
				},
				PermissionType: graphql.AppPermissionTypes.Provisioner,
				Applications: &graphql.AppFilter{
					FilterType: graphql.FilterTypes.All,
				},
				Provisioners: &graphql.ProvisionerPermissionFilter{
					FilterType: graphql.FilterTypes.All,
				},
			},
		},
	},
}

func createUserGroup(name string) (*graphql.UserGroup, error) {
	input := testUserGroup
	input.Name = name

	c := getClient()
	return c.Users().CreateUserGroup(&input)
}

func TestCreateUserGroup_NotificationSettings(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	ug := testUserGroup
	ug.Name = expectedName
	ug.NotificationSettings = nil

	c := getClient()

	userGroup, err := c.Users().CreateUserGroup(&ug)

	require.NoError(t, err)
	require.NotNil(t, userGroup)
	require.Equal(t, expectedName, userGroup.Name)
	require.Nil(t, userGroup.NotificationSettings.SlackNotificationSetting)
	require.ElementsMatch(t, ug.Permissions.AccountPermissions.AccountPermissionTypes, userGroup.Permissions.AccountPermissions.AccountPermissionTypes)

	// c := getClient()
	err = c.Users().DeleteUserGroup(userGroup.Id)
	require.NoError(t, err)
}

func TestCreateUserGroup(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	userGroup, err := createUserGroup(expectedName)

	require.NoError(t, err)
	require.NotNil(t, userGroup)
	require.Equal(t, expectedName, userGroup.Name)
	require.Equal(t, userGroup.NotificationSettings, testUserGroup.NotificationSettings)

	c := getClient()
	err = c.Users().DeleteUserGroup(userGroup.Id)
	require.NoError(t, err)
}

func TestUpdateUserGroup(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	userGroup, err := createUserGroup(expectedName)

	require.NoError(t, err)
	require.NotNil(t, userGroup)
	require.Equal(t, expectedName, userGroup.Name)

	userGroup.Description = "updated_description"

	c := getClient()

	updatedGroup, err := c.Users().UpdateUserGroup(userGroup)
	require.NoError(t, err)
	require.NotNil(t, updatedGroup)
	require.Equal(t, userGroup.Description, updatedGroup.Description)

	err = c.Users().DeleteUserGroup(userGroup.Id)
	require.NoError(t, err)
}

func TestGetUserGroupById(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	userGroup, err := createUserGroup(expectedName)

	require.NoError(t, err)
	require.NotNil(t, userGroup)

	c := getClient()

	foundUG, err := c.Users().GetUserGroupById(userGroup.Id)
	require.NoError(t, err)
	require.NotNil(t, foundUG)
	require.Equal(t, userGroup.Name, foundUG.Name)

	err = c.Users().DeleteUserGroup(userGroup.Id)
	require.NoError(t, err)
}

func TestGetUserGroupByName(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	userGroup, err := createUserGroup(expectedName)

	require.NoError(t, err)
	require.NotNil(t, userGroup)

	c := getClient()

	foundUG, err := c.Users().GetUserGroupByName(userGroup.Name)
	require.NoError(t, err)
	require.NotNil(t, foundUG)
	require.Equal(t, userGroup.Name, foundUG.Name)

	err = c.Users().DeleteUserGroup(userGroup.Id)
	require.NoError(t, err)
}

func TestAddUserToUserGroup(t *testing.T) {

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	userGroup, err := createUserGroup(expectedName)
	require.NoError(t, err)
	require.NotNil(t, userGroup)

	c := getClient()

	user, err := c.Users().GetUserByEmail("micahlmartin+testing@gmail.com")
	require.NoError(t, err)
	require.NotNil(t, user)

	ok, err := c.Users().AddUserToGroup(user.Id, userGroup.Id)
	require.NoError(t, err)
	require.True(t, ok)

	// User should already be in the group
	ok, err = c.Users().AddUserToGroup(user.Id, userGroup.Id)
	require.Error(t, err)
	require.False(t, ok)

	ok, err = c.Users().IsUserInGroup(user.Id, userGroup.Id)
	require.NoError(t, err)
	require.True(t, ok)

	ok, err = c.Users().RemoveUserFromGroup(user.Id, userGroup.Id)
	require.NoError(t, err)
	require.True(t, ok)

	ok, err = c.Users().IsUserInGroup(user.Id, userGroup.Id)
	require.NoError(t, err)
	require.False(t, ok)

	err = c.Users().DeleteUserGroup(userGroup.Id)
	require.NoError(t, err)
}

func TestListUserGroups(t *testing.T) {
	client := getClient()
	limit := 10
	offset := 0
	hasMore := true

	for hasMore {
		groups, pagination, err := client.Users().ListUsers(limit, offset)
		require.NoError(t, err, "Failed to list user groups: %s", err)
		require.NotEmpty(t, groups, "No user groups found")
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(groups) == limit
		offset += limit
	}
}

func TestListUsersGroupMembership(t *testing.T) {
	c := getClient()
	limit := 5
	offset := 0
	hasMore := true

	user, err := c.Users().GetUserByEmail("micahlmartin+testing@gmail.com")
	require.NoError(t, err)
	require.NotNil(t, user)

	for hasMore {
		groups, pagination, err := c.Users().ListGroupMembershipByUserId(user.Id, limit, offset)
		require.NoError(t, err, "Failed to list user group membership: %s", err)
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(groups) == limit
		offset += limit
	}

}
