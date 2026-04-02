package load_balancer_test

import (
	"strings"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// HARNESS_PLATFORM_API_KEY must be set for proxy tests (backend requires valid NG API key). If unset, tests are skipped.
const platformAPIKeyEnv = "HARNESS_PLATFORM_API_KEY"

// randAlnum returns a short random string with no underscore or space (for resource names).
func randAlnum(n int) string {
	s := strings.ToLower(utils.RandStringBytes(n))
	s = strings.ReplaceAll(s, "_", "")
	s = strings.ReplaceAll(s, " ", "")
	return s
}

func testGetLoadBalancer(resourceName string, state *terraform.State) (*nextgen.AccessPoint, error) {
	r := acctest.TestAccGetResource(resourceName, state)
	c, ctx := acctest.TestAccGetPlatformClientWithContext()
	id := r.Primary.ID

	resp, _, err := c.CloudCostAutoStoppingLoadBalancersApi.DescribeLoadBalancer(ctx, c.AccountId, id, c.AccountId)

	if err != nil {
		return nil, err
	}

	if resp.Response == nil {
		return nil, nil
	}

	return resp.Response, nil
}
