package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/micahlmartin/terraform-provider-harness/internal/utils"
	"github.com/stretchr/testify/require"
)

// Helper function for creating application
func createApplication(namePrefix string) (*Application, error) {
	client := getClient()
	name := fmt.Sprintf("%s-%s", namePrefix, utils.RandStringBytes(12))
	input := &Application{
		ClientMutationId: time.Now().String(),
		Name:             name,
		Description:      "Test application description",
	}

	return client.Applications().CreateApplication(input)
}

func TestGetApplicationById(t *testing.T) {

	// Create a new app
	newApp, err := createApplication(t.Name())
	require.Nil(t, err, "Failed to create application: %s", err)

	// Lookup newly created app by ID
	client := getClient()
	app, err := client.Applications().GetApplicationById(newApp.Id)

	// Verify
	require.NoError(t, err)
	require.NotNil(t, app, "App not found")
	require.Equal(t, newApp.Id, app.Id, "App Id doesn't match")
	require.NotEmpty(t, newApp.Description, "Description is empty")
	require.Equal(t, newApp.Description, app.Description, "Test application description")

	// cleanup
	client.Applications().DeleteApplication(newApp.Id)
}

func TestGetApplicationByName(t *testing.T) {
	// Create a new app
	newApp, err := createApplication(t.Name())
	require.NoError(t, err, "Failed to create application: %s", err)

	client := getClient()
	app, err := client.Applications().GetApplicationByName(newApp.Name)

	// Verify
	require.NoError(t, err, "Could not look up application by name")
	require.NotNil(t, app, "App not found")
	require.Equal(t, newApp.Name, app.Name, "App name doesn't match")

	// cleanup
	client.Applications().DeleteApplication(newApp.Id)
}

func TestCreateApplication(t *testing.T) {

	// Create application
	client := getClient()
	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(12))
	input := &Application{
		ClientMutationId: time.Now().String(),
		Name:             name,
		Description:      "Test application description",
	}

	app, err := client.Applications().CreateApplication(input)

	// Verify
	require.NoError(t, err, "Failed to create application: %s", err)
	require.NotNil(t, app, "Application should not be nil")
	require.Equal(t, input.Name, app.Name, "Application name doesn't match")

	// Cleanup
	client.Applications().DeleteApplication(app.Id)
}

func TestGetDeleteApplication(t *testing.T) {

	// Create a new app
	newApp, err := createApplication(t.Name())
	require.NoError(t, err, "Failed to create application: %s", err)

	client := getClient()

	// Delete application
	client.Applications().DeleteApplication(newApp.Id)

	// Verify
	app, err := client.Applications().GetApplicationById(newApp.Id)

	require.Nil(t, app, "Failed to delete app")
	require.NotNil(t, err)

}

func TestUpdateApplication(t *testing.T) {

	// Setup
	expectedName := "test_name_change"

	// Create a new app
	newApp, err := createApplication(t.Name())
	require.NoError(t, err, "Failed to create application: %s", err)
	require.NotEqual(t, expectedName, newApp.Name)

	client := getClient()

	// Update application
	updatedApp, err := client.Applications().UpdateApplication(&UpdateApplicationInput{
		ApplicationId: newApp.Id,
		Name:          expectedName,
	})

	// Verify
	require.NoError(t, err)
	require.NotNil(t, updatedApp)
	require.Equal(t, expectedName, updatedApp.Name)

	// Cleanup
	client.Applications().DeleteApplication(updatedApp.Id)
}
