package delegate_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/cd/graphql"
	"github.com/harness-io/harness-go-sdk/harness/delegate"
	"github.com/harness-io/harness-go-sdk/harness/helpers"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

var defaultDelegateTimeout = time.Minute * 5

func getDelegateTimeout() time.Duration {
	if t := helpers.TestEnvVars.DelegateWaitTimeout.Get(); t != "" {
		timeout, err := time.ParseDuration(t)
		if err != nil {
			fmt.Printf("failed to parse delegate wait timeout: %s. Using default.", err)
			return defaultDelegateTimeout
		}
		return timeout
	}

	return defaultDelegateTimeout
}

func createDelegateContainer(t *testing.T, name string) *graphql.Delegate {
	ctx := context.Background()
	c := acctest.TestAccProvider.Meta().(*api.Client)

	cfg := &delegate.DockerDelegateConfig{
		AccountId:     c.AccountId,
		AccountSecret: helpers.TestEnvVars.DelegateSecret.Get(),
		DelegateName:  name,
		ContainerName: name,
		ProfileId:     helpers.TestEnvVars.DelegateProfileId.Get(),
	}

	t.Logf("Starting delegate %s", name)
	_, err := delegate.RunDelegateContainer(ctx, cfg)
	require.NoError(t, err, "failed to create delegate container: %s", err)

	delegate, err := c.CDClient.DelegateClient.WaitForDelegate(ctx, name, getDelegateTimeout())
	require.NoError(t, err, "failed to wait for delegate: %s", err)
	require.NotNil(t, delegate, "delegate should not be nil")

	return delegate
}

func deleteDelegate(t *testing.T, name string) {
	c := acctest.TestAccProvider.Meta().(*api.Client)
	delegate, err := c.CDClient.DelegateClient.GetDelegateByName(name)
	require.NoError(t, err, "Failed to get delegate: %s", err)
	require.NotNil(t, delegate, "Delegate should not be nil")

	cli, err := client.NewClientWithOpts(client.FromEnv)
	require.NoError(t, err, "failed to create docker client: %s", err)

	err = cli.ContainerStop(context.Background(), name, nil)
	require.NoError(t, err, "failed to stop delegate container: %s", err)

	err = cli.ContainerRemove(context.Background(), name, types.ContainerRemoveOptions{})
	require.NoError(t, err, "failed to remove delegate container: %s", err)
}

func TestAccApproveDelegate(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(7))
		resourceName = "harness_delegate_approval.test"
	)

	defer deleteDelegate(t, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			createDelegateContainer(t, name)
		},
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testaccDelegateApproval(name, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "status", graphql.DelegateStatusTypes.Enabled.String()),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					primary := s.RootModule().Resources[resourceName].Primary
					return primary.Attributes["name"], nil
				},
			},
			{
				Config:      testaccDelegateApproval(name, false),
				ExpectError: regexp.MustCompile(`.*has already been changed.*`),
			},
		},
	})
}

func testaccDelegateApproval(name string, approve bool) string {
	return fmt.Sprintf(`
		resource "harness_delegate_approval" "test" {
			name = "%[1]s"
			approve = %[2]t
		}
	`, name, approve)
}
