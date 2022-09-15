package nextgen

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetServiceByName(t *testing.T) {
	c, ctx := getClientWithContext()
	id := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(6))

	service := ServiceRequest{
		Identifier: id,
		Name:       id,
	}

	resp, _, err := c.ServicesApi.CreateServiceV2(ctx, c.AccountId, &ServicesApiCreateServiceV2Opts{
		Body: optional.NewInterface(service),
	})
	require.NoError(t, err)
	require.NotNil(t, resp.Data)

	defer func() {
		c.ServicesApi.DeleteServiceV2(ctx, c.AccountId, id, &ServicesApiDeleteServiceV2Opts{})
	}()

	result, _, err := c.ServicesApi.GetServiceByName(ctx, c.AccountId, id, GetServiceByNameOpts{})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, result.Name, id)
}
