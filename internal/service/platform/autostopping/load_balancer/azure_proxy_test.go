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
	hostName := fmt.Sprintf("ab%s.com", name)
	resourceName := "harness_autostopping_azure_proxy.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testAzureProxyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAzureProxy(name, hostName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName),
				),
			},
			{
				Config: testAzureProxyUpdate(name, hostName),
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
		proxy, _ := testGetLoadBalancer(resourceName, state)
		if proxy != nil {
			return fmt.Errorf("Found azure proxy: %s", proxy.Id)
		}
		return nil
	}
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

func testAzureProxy(name string, hostName string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_azure_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = "cloud_connector_id"
			host_name = "%[2]s"
			region             = "eastus2"
			resource_group     = "resource_group"
			vpc                = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network"
			subnet_id          = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network/subnets/subnet_id"
			security_groups    = ["/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/networkSecurityGroups/network_security_group"]
			allocate_static_ip = true
            machine_type = "Standard_D2s_v3"
			keypair = "PLACE_HOLDER_VALUE"
            api_key = "PLACE_HOLDER_VALUE"
			delete_cloud_resources_on_destroy = true
		}
`, name, hostName)
}

func testAzureProxyUpdate(name string, hostName string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_azure_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = "cloud_connector_id"
			host_name = "%[2]s"
			region             = "eastus2"
			resource_group     = "resource_group"
			vpc                = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network"
			subnet_id          = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network/subnets/subnet_id"
			security_groups    = ["/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/networkSecurityGroups/network_security_group"]
			allocate_static_ip = true
            machine_type = "Standard_D2s_v3"
			keypair = "PLACE_HOLDER_VALUE"
            api_key = "PLACE_HOLDER_VALUE"
			delete_cloud_resources_on_destroy = false
		}
`, name, hostName)
}
