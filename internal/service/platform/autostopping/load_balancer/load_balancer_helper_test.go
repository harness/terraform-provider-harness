package load_balancer_test

import (
	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

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
