package load_balancer_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceAzureProxy(t *testing.T) {
	apiKey := os.Getenv(platformAPIKeyEnv)

	name := fmt.Sprintf("terr-azure-proxy%s", strings.ToLower(utils.RandStringBytes(5)))
	resourceName := "harness_autostopping_azure_proxy.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testAzureProxyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAzureProxy(name, apiKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				Config: testAzureProxyUpdate(name, apiKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"api_key", "allocate_static_ip", "delete_cloud_resources_on_destroy", "keypair", "machine_type", "resource_group", "subnet_id"},
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

func testAzureProxy(name, apiKey string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_azure_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = "doNotDeleteAzureConnector"
			region             = "eastus2"
			resource_group     = "resource_group"
			vpc                = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network"
			subnet_id          = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network/subnets/subnet_id"
			security_groups    = ["/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/networkSecurityGroups/network_security_group"]
			allocate_static_ip = true
            machine_type = "Standard_D2s_v3"
			keypair = "PLACE_HOLDER_VALUE"
            api_key = %q
			delete_cloud_resources_on_destroy = true
		}
`, name, apiKey)
}

func testAzureProxyUpdate(name, apiKey string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_azure_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = "doNotDeleteAzureConnector"
			region             = "eastus2"
			resource_group     = "resource_group"
			vpc                = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network"
			subnet_id          = "/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/virtualNetworks/virtual_network/subnets/subnet_id"
			security_groups    = ["/subscriptions/subscription_id/resourceGroups/resource_group/providers/Microsoft.Network/networkSecurityGroups/network_security_group"]
			allocate_static_ip = true
            machine_type = "Standard_D2s_v3"
			keypair = "PLACE_HOLDER_VALUE"
            api_key = %q
			delete_cloud_resources_on_destroy = false
		}
`, name, apiKey)
}
