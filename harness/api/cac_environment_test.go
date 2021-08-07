package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestUpsertEnvironment(t *testing.T) {
	c := getClient()
	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
	envName := fmt.Sprintf("svc-%s-%s", t.Name(), utils.RandStringBytes(5))

	app, err := createApplication(appName)
	require.NoError(t, err)

	env, err := createEnvironment(app.Id, envName)
	require.NoError(t, err)
	require.NotNil(t, env)

	require.Equal(t, env.Name, envName)
	require.Equal(t, env.ApplicationId, app.Id)

	err = c.ConfigAsCode().DeleteEntity(cac.GetEnvironmentYamlPath(app.Name, envName))
	require.NoError(t, err)
}

func TestGetEnvironmentById(t *testing.T) {
	c := getClient()

	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
	envName := fmt.Sprintf("svc-%s-%s", t.Name(), utils.RandStringBytes(5))

	app, err := createApplication(appName)
	require.NoError(t, err)

	env, err := createEnvironment(app.Id, envName)
	require.NoError(t, err)
	require.NotNil(t, env)

	envLookup, err := c.ConfigAsCode().GetEnvironmentById(app.Id, env.Id)
	require.NoError(t, err)
	require.NotNil(t, envLookup)
	require.Equal(t, env, envLookup)
}

func TestGetEnvironmentByName(t *testing.T) {
	c := getClient()

	appName := fmt.Sprintf("app-%s-%s", t.Name(), utils.RandStringBytes(5))
	envName := fmt.Sprintf("svc-%s-%s", t.Name(), utils.RandStringBytes(5))

	app, err := createApplication(appName)
	require.NoError(t, err)

	env, err := createEnvironment(app.Id, envName)
	require.NoError(t, err)
	require.NotNil(t, env)

	envLookup, err := c.ConfigAsCode().GetEnvironmentByName(app.Id, env.Name)
	require.NoError(t, err)
	require.NotNil(t, envLookup)
	require.Equal(t, env, envLookup)
}

func createEnvironment(applicationId string, name string) (*cac.Environment, error) {
	input := cac.NewEntity(cac.ObjectTypes.Environment).(*cac.Environment)
	input.Name = name
	input.ApplicationId = applicationId
	input.EnvironmentType = cac.EnvironmentTypes.Prod
	input.ConfigMapYamlByServiceTemplateName = &map[string]interface{}{}

	c := getClient()
	return c.ConfigAsCode().UpsertEnvironment(input)

}
