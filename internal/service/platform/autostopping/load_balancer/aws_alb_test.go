package load_balancer_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceAwsALB(t *testing.T) {
	name := utils.RandStringBytes(5)
	hostName := fmt.Sprintf("ab%s.lightwingtest.com", name)
	resourceName := "harness_autostopping_aws_alb.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testAWSProxyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAwsALB(name, hostName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName),
				),
			},
		},
	})
}

func testAwsALBDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		proxy, _ := testGetLoadBalancer(resourceName, state)
		if proxy != nil {
			return fmt.Errorf("Found aws alb: %s", proxy.Id)
		}
		return nil
	}
}

func testAwsALB(name string, hostName string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_alb" "test" {
			name = "%[1]s"
			cloud_connector_id = "LightwingNonProd"
			host_name = "%[2]s"
            region = "us-east-1"
			vpc = "vpc-2657db5c"
			security_groups =["sg-005ae65c1e4ef3227"]
			route53_hosted_zone_id = "/hostedzone/Z06070943NA512B2KHEHF"
		}
`, name, hostName)
}

func testAwsALBUpdate(name string, hostName string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_alb" "test" {
			name = "%[1]s"
			cloud_connector_id = "LightwingNonProd"
			host_name = "%[2]s"
            region = "us-east-1"
            vpc = "vpc-2657db5c"
			security_groups =["sg-005ae65c1e4ef3227","sg-005ae65c1e4ef3245"]
			route53_hosted_zone_id = "/hostedzone/Z06070943NA512B2KHEHF"
		}
`, name, hostName)
}
