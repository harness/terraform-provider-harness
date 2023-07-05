package load_balancer_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceAWSProxy(t *testing.T) {
	name := utils.RandStringBytes(5)
	hostName := fmt.Sprintf("ab%s.lightwingtest.com", name)
	resourceName := "harness_autostopping_aws_proxy.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testAWSProxyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAWSProxy(name, hostName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"api_key"},
			},
		},
	})
}

func testAWSProxyDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		proxy, _ := testGetLoadBalancer(resourceName, state)
		if proxy != nil {
			return fmt.Errorf("Found aws proxy: %s", proxy.Id)
		}
		return nil
	}
}

func testAWSProxy(name string, hostName string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = "LightwingNonProd"
			host_name = "%[2]s"
            region = "us-east-1"
			vpc = "vpc-2657db5c"
			security_groups =["sg-005ae65c1e4ef3227"]
			route53_hosted_zone_id = "/hostedzone/Z06070943NA512B2KHEHF"
			machine_type = "t2.medium"
            api_key = ""
			allocate_static_ip = true
		}
`, name, hostName)
}

func testAWSProxyUpdate(name string, hostName string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = "LightwingNonProd"
			host_name = "%[2]s"
            region = "eastus2"
            resource_group = "tkouhsari-autostop-1_group"
            vpc = "vpc-2657db5c"
			security_groups =["sg-005ae65c1e4ef3227","sg-005ae65c1e4ef3245"]
            machine_type = "t2.medium"
            api_key = ""
		}
`, name, hostName)
}
