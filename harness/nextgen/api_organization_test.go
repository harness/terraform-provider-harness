package nextgen

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetOrganizationByName(t *testing.T) {
	c, ctx := getClientWithContext()
	id := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(6))

	org := OrganizationRequest{
		Organization: &Organization{
			Identifier: id,
			Name:       id,
		},
	}

	resp, _, err := c.OrganizationApi.PostOrganization(ctx, org, c.AccountId)
	require.NoError(t, err)
	require.NotNil(t, resp.Data)

	defer func() {
		c.OrganizationApi.DeleteOrganization(ctx, id, c.AccountId, &OrganizationApiDeleteOrganizationOpts{})
	}()

	result, _, err := c.OrganizationApi.GetOrganizationByName(ctx, c.AccountId, id)
	require.NoError(t, err)
	require.NotNil(t, result.Organization)
	require.Equal(t, result.Organization.Name, id)
}
