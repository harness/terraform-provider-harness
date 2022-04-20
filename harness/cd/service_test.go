package cd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListServices(t *testing.T) {
	c := getClient()
	limit := 100
	offset := 0
	hasMore := true

	for hasMore {
		services, pagination, err := c.ServiceClient.ListServices(limit, offset, nil)
		require.NoError(t, err, "Failed to list services: %s", err)
		require.NotEmpty(t, services, "No services found")
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(services) == limit
		offset += limit
	}
}

func TestListServicesByApplicationId(t *testing.T) {
	c := getClient()

	app, err := c.ApplicationClient.GetApplicationByName("Harness Automation Example")
	require.NoError(t, err, "Failed to get application: %s", err)

	limit := 100
	offset := 0
	hasMore := true

	for hasMore {
		services, pagination, err := c.ServiceClient.ListServicesByApplicationId(app.Id, 100, 0)
		require.NoError(t, err, "Failed to list services: %s", err)
		require.NotEmpty(t, services, "No services found")
		require.NotNil(t, pagination, "Pagination should not be nil")

		hasMore = len(services) == limit
		offset += limit
	}
}

func TestGraphqlGetServiceById(t *testing.T) {
	c := getClient()

	app, err := c.ApplicationClient.GetApplicationByName("Harness Automation Example")
	require.NoError(t, err, "Failed to get application: %s", err)

	services, pagination, err := c.ServiceClient.ListServicesByApplicationId(app.Id, 1, 0)
	require.NoError(t, err, "Failed to list services: %s", err)
	require.NotEmpty(t, services, "No services found")
	require.NotNil(t, pagination, "Pagination should not be nil")

	svc := services[0]

	foundSvc, err := c.ServiceClient.GetServiceById(svc.Id)
	require.NoError(t, err, "Failed to get service: %s", err)
	require.NotNil(t, foundSvc, "Service should not be nil")
	require.Equal(t, svc.Id, foundSvc.Id)
}

func TestGraphqlGetServiceByName(t *testing.T) {
	c := getClient()

	svcName := "nginx"

	app, err := c.ApplicationClient.GetApplicationByName("Harness Automation Example")
	require.NoError(t, err, "Failed to get application: %s", err)

	foundSvc, err := c.ServiceClient.GetServiceByName(app.Id, svcName)
	require.NoError(t, err, "Failed to get service: %s", err)
	require.NotNil(t, foundSvc, "Service should not be nil")
	require.Equal(t, svcName, foundSvc.Name)
}
