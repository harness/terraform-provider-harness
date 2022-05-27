package cd

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestAttachTagToApplication(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(4))

	app, err := createApplication(name)
	require.NoError(t, err)
	require.NotNil(t, app)

	defer func() {
		c.ApplicationClient.DeleteApplication(app.Id)
	}()

	input := &graphql.AttachTagInput{
		EntityId:   app.Id,
		EntityType: graphql.TagEntityTypes.Application,
		Name:       "test",
		Value:      "foo",
	}

	resp, err := c.TagClient.AttachTag(input)
	require.NoError(t, err)
	require.Equal(t, resp.TagLink.EntityId, input.EntityId)
	require.Equal(t, resp.TagLink.EntityType, input.EntityType)
	require.Equal(t, resp.TagLink.Name, input.Name)
	require.Equal(t, resp.TagLink.Value, input.Value)

	app, err = c.ApplicationClient.GetApplicationById(app.Id)
	require.NoError(t, err)
	require.Equal(t, 1, len(app.Tags))
	require.Equal(t, "test", app.Tags[0].Name)
	require.Equal(t, "foo", app.Tags[0].Value)
}

func TestDetachTagToApplication(t *testing.T) {
	c := getClient()

	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(4))

	app, err := createApplication(name)
	require.NoError(t, err)
	require.NotNil(t, app)

	defer func() {
		c.ApplicationClient.DeleteApplication(app.Id)
	}()

	input := &graphql.AttachTagInput{
		EntityId:   app.Id,
		EntityType: graphql.TagEntityTypes.Application,
		Name:       name,
		Value:      "foo",
	}

	resp, err := c.TagClient.AttachTag(input)
	require.NoError(t, err)
	require.Equal(t, resp.TagLink.EntityId, input.EntityId)
	require.Equal(t, resp.TagLink.EntityType, input.EntityType)
	require.Equal(t, resp.TagLink.Name, input.Name)
	require.Equal(t, resp.TagLink.Value, input.Value)

	detachInput := &graphql.DetachTagInput{
		EntityId:   app.Id,
		EntityType: graphql.TagEntityTypes.Application,
		Name:       name,
	}
	err = c.TagClient.DetachTag(detachInput)
	require.NoError(t, err)

	app, err = c.ApplicationClient.GetApplicationById(app.Id)
	require.NoError(t, err)
	require.Equal(t, 0, len(app.Tags))
}
