package load_balancer_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccResourceAwsALB_CCM32336_OutOfBandDeleteRecreates verifies that when an
// AutoStopping load balancer (AWS ALB) is deleted out-of-band (UI / direct API),
// the next refresh treats the GET as "not found" (HTTP 404 + ENTITY_NOT_FOUND)
// and re-plans a create instead of erroring out.
//
// Regression test for CCM-32336. The provider's helpers.HandleReadApiError clears
// state on 404 + ENTITY_NOT_FOUND so terraform plan reports a recreate.
func TestAccResourceAwsALB_CCM32336_OutOfBandDeleteRecreates(t *testing.T) {
	name := fmt.Sprintf("terr-c336alb-%s", randAlnum(5))
	resourceName := "harness_autostopping_aws_alb.test"

	var lbIDBefore string

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testCCM32336AwsALB(name),
				Check: resource.ComposeTestCheckFunc(
					extractAttr(resourceName, "identifier", &lbIDBefore),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					c, ctx := acctest.TestAccGetPlatformClientWithContext()
					if _, err := c.CloudCostAutoStoppingLoadBalancersApi.DeleteLoadBalancer(
						ctx,
						nextgen.DeleteAccessPointPayload{
							Ids:           []string{lbIDBefore},
							WithResources: true,
						},
						c.AccountId, c.AccountId,
					); err != nil {
						t.Fatalf("CCM-32336: out-of-band delete failed: %v", err)
					}
				},
				Config:             testCCM32336AwsALB(name),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testCCM32336AwsALB(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrWith(resourceName, "identifier", func(value string) error {
						if value == "" || value == lbIDBefore {
							return fmt.Errorf("expected new load balancer id after recreate, got %q (before %q)", value, lbIDBefore)
						}
						return nil
					}),
				),
			},
		},
	})
}

func testCCM32336AwsALB(name string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_alb" "test" {
			name = "%[1]s"
			cloud_connector_id = %q
			region = "us-east-1"
			vpc = %q
			security_groups = [%q]
			delete_cloud_resources_on_destroy = true
		}
`, name, awsCloudConnectorID, awsProxyVPC, awsProxySG)
}
