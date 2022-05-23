package nextgen

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetProjectByName(t *testing.T) {
	name := fmt.Sprintf("%s%s", t.Name(), utils.RandStringBytes(6))
	c, ctx := getClientWithContext()

	p := ProjectRequest{
		Project: &Project{
			Name:       name,
			Identifier: name,
		},
	}
	resp, _, err := c.ProjectApi.PostProject(ctx, p, c.AccountId, &ProjectApiPostProjectOpts{})
	require.NoError(t, err)
	require.NotNil(t, resp.Data)

	defer func() {
		c.ProjectApi.DeleteProject(ctx, name, c.AccountId, &ProjectApiDeleteProjectOpts{})
	}()

	result, err := c.ProjectApi.GetProjectByName(ctx, c.AccountId, "default", name)
	require.NoError(t, err)
	require.NotNil(t, result.Project)
	require.Equal(t, result.Project.Name, name)
}
