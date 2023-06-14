package load_balancer_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceAzureProxy(t *testing.T) {
	name := utils.RandStringBytes(5)
	hostName := fmt.Sprintf("%s.com", name)
	resourceName := "harness_autostopping_azure_proxy.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAzureProxyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAzureProxy(name, hostName),
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

func testAzureProxyDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		proxy, _ := testGetAzureProxy(resourceName, state)
		if proxy != nil {
			return fmt.Errorf("Found azure proxy: %s", proxy.Id)
		}
		return nil
	}
}

func testGetAzureProxy(resourceName string, state *terraform.State) (*nextgen.AccessPoint, error) {
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

func testAzureProxy(name string, hostName string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_azure_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = "Azure_SE"
			host_name = "%[2]s"
            region = "eastus2"
            resource_group = "tkouhsari-autostop-1_group"
            vpc = "/subscriptions/e8389fc5-0cb8-44ab-947b-c6cf62552be0/resourceGroups/tkouhsari-autostop-1_group/providers/Microsoft.Network/virtualNetworks/tkouhsari-autostop-1-vnet"
			subnet_id = "/subscriptions/e8389fc5-0cb8-44ab-947b-c6cf62552be0/resourceGroups/tkouhsari-autostop-1_group/providers/Microsoft.Network/virtualNetworks/tkouhsari-autostop-1-vnet/subnets/default"
			security_groups =["/subscriptions/e8389fc5-0cb8-44ab-947b-c6cf62552be0/resourceGroups/tkouhsari-autostop-1_group/providers/Microsoft.Network/networkSecurityGroups/tkouhsari-autostop-1-nsg"]
			allocate_static_ip = true
            machine_type = "Standard_D2s_v3"
			keypair = "tkouhsari-autostop-1_key"
            api_key = "pat.PL7d6h0LQP-O91d5j7Xgsg.645b976a9a97612476a2c987.n5z8la6q0Ji2FD37iaPY"
		}
`, name, hostName)
}
