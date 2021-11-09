package cd

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestUpdateGitSyncConfig(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	conn, err := createGitConnector(name)
	require.NoError(t, err)
	require.NotNil(t, conn)

	defer func() {
		err := c.ConnectorClient.DeleteConnector(conn.Id)
		require.NoError(t, err)
	}()

	app, err := createApplication(name)
	require.NoError(t, err)
	require.NotNil(t, app)

	defer func() {
		err := c.ApplicationClient.DeleteApplication(app.Id)
		require.NoError(t, err)
	}()

	updateConfig := &graphql.UpdateApplicationGitSyncConfigInput{
		ApplicationId:  app.Id,
		GitConnectorId: conn.Id,
		SyncEnabled:    false,
		Branch:         "main",
	}

	config, err := c.ApplicationClient.UpdateGitSyncConfig(updateConfig)
	require.NoError(t, err)
	require.NotNil(t, config)

	err = c.ApplicationClient.RemoveGitSyncConfig(app.Id)
	require.NoError(t, err)
}
