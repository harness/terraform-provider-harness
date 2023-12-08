package load_balancer_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceAzureGateway(t *testing.T) {
	name := utils.RandStringBytes(5)
	hostName := fmt.Sprintf("ab%s.com", name)
	resourceName := "harness_autostopping_azure_gateway.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testAzureGatewayDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAzureGateway(name, hostName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "host_name", hostName),
				),
			},
		},
	})
}

func testAzureGatewayDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		proxy, _ := testGetLoadBalancer(resourceName, state)
		if proxy != nil {
			return fmt.Errorf("Found aws proxy: %s", proxy.Id)
		}
		return nil
	}
}

func testAzureGateway(name string, hostName string) string {
	return fmt.Sprintf(`
	resource "harness_autostopping_azure_gateway" "test" {
		name = "%[1]s"
		cloud_connector_id = "cloud_connector_id"
		host_name = "%[2]s"
		region = "eastus2"
		resource_group     = "resource_group"
		subnet_id          = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network/subnets/subnet_id"
		vpc                = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network"
		azure_func_region  = "westus2"
		frontend_ip        = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/publicIPAddresses/publicip"
		sku_size           = "sku2"
		delete_cloud_resources_on_destroy = true
	}
`, name, hostName)
}
