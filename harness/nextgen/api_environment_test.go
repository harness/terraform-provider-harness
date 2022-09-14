package nextgen

import (
	"fmt"
	"testing"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetEnvironmentByName(t *testing.T) {
	c, ctx := getClientWithContext()
	id := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(6))

	environment := EnvironmentRequest{
		Identifier: id,
		Name:       id,
		Type_:      EnvironmentTypes.PreProduction,
	}

	resp, _, err := c.EnvironmentsApi.CreateEnvironmentV2(ctx, c.AccountId, &EnvironmentsApiCreateEnvironmentV2Opts{
		Body: optional.NewInterface(environment),
	})
	require.NoError(t, err)
	require.NotNil(t, resp.Data)

	defer func() {
		c.EnvironmentsApi.DeleteEnvironmentV2(ctx, c.AccountId, id, &EnvironmentsApiDeleteEnvironmentV2Opts{})
	}()

	result, _, err := c.EnvironmentsApi.GetEnvironmentByName(ctx, c.AccountId, id, GetEnvironmentByNameOpts{})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, result.Name, id)
}
