package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestCacGetApplicationById(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	c := getClient()

	testApp, err := createApplication(name)
	require.NoError(t, err)
	require.NotNil(t, testApp)

	app, err := c.ConfigAsCode().GetApplicationById(testApp.Id)
	require.NoError(t, err)
	require.NotNil(t, app)

	require.Equal(t, testApp.Id, app.Id)
	require.Equal(t, testApp.Name, app.Name)
}

func TestCacGetApplicationByName(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	testApp, err := createApplication(name)
	require.NoError(t, err)
	require.NotNil(t, testApp)

	c := getClient()

	app, err := c.ConfigAsCode().GetApplicationByName(testApp.Name)
	require.NoError(t, err)
	require.NotNil(t, app)
}

func TestCacGetApplicationById_DoesNotExist(t *testing.T) {
	c := getClient()

	app, err := c.ConfigAsCode().GetApplicationById("somenonexistentapp")
	require.True(t, app.IsEmpty())
	require.NoError(t, err)
}

func TestEmptyEntity_BothEmpty(t *testing.T) {
	app := &cac.Application{}
	empty := &cac.Application{}

	require.True(t, reflect.DeepEqual(app, empty))
}

func TestEmptyEntity_NotEmpty(t *testing.T) {
	app := &cac.Application{
		Id:                "id",
		HarnessApiVersion: cac.HarnessApiVersions.V1,
	}
	empty := &cac.Application{}

	require.False(t, reflect.DeepEqual(app, empty))
}
