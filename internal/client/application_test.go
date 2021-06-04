package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/micahlmartin/terraform-provider-harness/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

// Helper function for creating application
func createApplication(namePrefix string) (*Application, error) {
	client := getClient()
	name := fmt.Sprintf("%s-%s", namePrefix, testhelpers.RandStringBytes(12))
	input := &CreateApplicationInput{
		ClientMutationId: time.Now().String(),
		Name:             name,
		Description:      "Test application description",
	}

	return client.Applications().CreateApplication(input)
}

func TestGetApplicationById(t *testing.T) {

	// Create a new app
	newApp, err := createApplication(t.Name())
	assert.Nil(t, err, "Failed to create application: %s", err)

	// Lookup newly created app by ID
	client := getClient()
	app, err := client.Applications().GetApplicationById(newApp.Id)

	// Verify
	assert.Nil(t, err, "Could not look up application by id")
	assert.NotNil(t, app, "App not found")
	assert.Equal(t, newApp.Id, app.Id, "App Id doesn't match")

	// cleanup
	client.Applications().DeleteApplication(newApp.Id)
}

func TestGetApplicationByName(t *testing.T) {
	// Create a new app
	newApp, err := createApplication(t.Name())
	assert.Nil(t, err, "Failed to create application: %s", err)

	client := getClient()
	app, err := client.Applications().GetApplicationByName(newApp.Name)

	// Verify
	assert.Nil(t, err, "Could not look up application by name")
	assert.NotNil(t, app, "App not found")
	assert.Equal(t, newApp.Name, app.Name, "App name doesn't match")

	// cleanup
	client.Applications().DeleteApplication(newApp.Id)
}

func TestCreateApplication(t *testing.T) {

	// Create application
	client := getClient()
	name := fmt.Sprintf("%s-%s", t.Name(), testhelpers.RandStringBytes(12))
	input := &CreateApplicationInput{
		ClientMutationId: time.Now().String(),
		Name:             name,
		Description:      "Test application description",
	}

	app, err := client.Applications().CreateApplication(input)

	// Verify
	assert.Nil(t, err, "Failed to create application: %s", err)
	assert.NotNil(t, app, "Application should not be nil")
	assert.Equal(t, input.Name, app.Name, "Application name doesn't match")

	// Cleanup
	client.Applications().DeleteApplication(app.Id)
}

func TestGetDeleteApplication(t *testing.T) {

	// Create a new app
	newApp, err := createApplication(t.Name())
	assert.Nil(t, err, "Failed to create application: %s", err)

	client := getClient()

	// Delete application
	client.Applications().DeleteApplication(newApp.Id)

	// Verify
	app, err := client.Applications().GetApplicationById(newApp.Id)

	assert.Nil(t, app, "Failed to delete app")
	assert.NotNil(t, err)

}
