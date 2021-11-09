package cd

import (
	"fmt"
	"testing"
	"time"

	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetApplicationById(t *testing.T) {

	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(5))
	newApp, err := createApplication(name)
	require.Nil(t, err, "Failed to create application: %s", err)

	// Lookup newly created app by ID
	client := getClient()
	app, err := client.ApplicationClient.GetApplicationById(newApp.Id)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, app, "App not found")
	require.Equal(t, newApp.Id, app.Id, "App Id doesn't match")
	require.NotEmpty(t, newApp.Description, "Description is empty")
	require.Equal(t, newApp.Description, app.Description, "Test application description")

	// cleanup
	err = client.ApplicationClient.DeleteApplication(newApp.Id)
	require.Nil(t, err, "Failed to delete application: %s", err)
}

func TestGetApplicationByName(t *testing.T) {
	// Create a new app
	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(5))
	newApp, err := createApplication(name)
	require.NoError(t, err, "Failed to create application: %s", err)

	client := getClient()
	app, err := client.ApplicationClient.GetApplicationByName(newApp.Name)

	// Verify
	require.NoError(t, err, "Could not look up application by name")
	require.NotNil(t, app, "App not found")
	require.Equal(t, newApp.Name, app.Name, "App name doesn't match")

	// cleanup
	err = client.ApplicationClient.DeleteApplication(newApp.Id)
	require.Nil(t, err, "Failed to delete application: %s", err)
}

func TestCreateApplication(t *testing.T) {

	// Create application
	client := getClient()
	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(5))
	input := &graphql.Application{
		ClientMutationId: time.Now().String(),
		Name:             name,
		Description:      "Test application description",
	}

	app, err := client.ApplicationClient.CreateApplication(input)

	// Verify
	require.NoError(t, err, "Failed to create application: %s", err)
	require.NotNil(t, app, "Application should not be nil")
	require.Equal(t, input.Name, app.Name, "Application name doesn't match")

	// Cleanup
	err = client.ApplicationClient.DeleteApplication(app.Id)
	require.Nil(t, err, "Failed to delete application: %s", err)
}

func TestGetDeletedApplication(t *testing.T) {

	// Create a new app
	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(5))
	newApp, err := createApplication(name)
	require.NoError(t, err, "Failed to create application: %s", err)

	client := getClient()

	// Delete application
	err = client.ApplicationClient.DeleteApplication(newApp.Id)
	require.Nil(t, err, "Failed to delete application: %s", err)

	// Verify
	app, err := client.ApplicationClient.GetApplicationById(newApp.Id)

	require.Nil(t, app, "Failed to delete app")
	require.NoError(t, err)

}

func TestUpdateApplication(t *testing.T) {

	// Setup
	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(5))
	expectedName := "test_name_change"

	// Create a new app
	newApp, err := createApplication(name)
	require.NoError(t, err, "Failed to create application: %s", err)
	require.Equal(t, name, newApp.Name)

	client := getClient()

	// Update application
	updatedApp, err := client.ApplicationClient.UpdateApplication(&graphql.UpdateApplicationInput{
		ApplicationId: newApp.Id,
		Name:          expectedName,
	})

	// Verify
	require.NoError(t, err)
	require.NotNil(t, updatedApp)
	require.Equal(t, expectedName, updatedApp.Name)

	// Cleanup
	err = client.ApplicationClient.DeleteApplication(updatedApp.Id)
	require.Nil(t, err, "Failed to delete application: %s", err)
}

// Helper function for creating application
func createApplication(name string) (*graphql.Application, error) {
	client := getClient()
	input := &graphql.Application{
		ClientMutationId: time.Now().String(),
		Name:             name,
		Description:      "Test application description",
	}

	return client.ApplicationClient.CreateApplication(input)
}

func TestListApplications(t *testing.T) {
	client := getClient()
	limit := 100
	offset := 0
	hasMore := true

	for hasMore {
		apps, pagination, err := client.ApplicationClient.ListApplications(limit, offset)
		require.NoError(t, err, "Failed to list applications: %s", err)
		require.NotEmpty(t, apps, "No applications found")
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(apps) == limit
		offset += limit
	}
}
