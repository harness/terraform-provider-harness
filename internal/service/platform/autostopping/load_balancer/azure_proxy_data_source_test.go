package load_balancer_test

import (
	"testing"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAzureProxyDataSource(t *testing.T) {
	resourceName := "data.harness_autostopping_azure_proxy_AzrProxyTFTest"
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testSampleDataSourceAzureProxy(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "AzrProxyTFTest"),
					resource.TestCheckResourceAttr(resourceName, "cloud_connector_id", "DoNotDelete_LightwingQA"),
					resource.TestCheckResourceAttr(resourceName, "host_name", "as-azr-proxy-tf-test.harness.io"),
					resource.TestCheckResourceAttr(resourceName, "region", "westus2"),
					resource.TestCheckResourceAttr(resourceName, "resource_group", "lightwing-r-and-d"),
					resource.TestCheckResourceAttr(resourceName, "vpc", "/subscriptions/abcdefgh-ijkl-mnop-qrst-uvwxyz123456/resourceGroups/lightwing-r-and-d/providers/Microsoft.Network/virtualNetworks/lightwing-r-and-d-vnet"),
					resource.TestCheckResourceAttr(resourceName, "subnet_id", "/subscriptions/abcdefgh-ijkl-mnop-qrst-uvwxyz123456/resourceGroups/lightwing-r-and-d/providers/Microsoft.Network/virtualNetworks/lightwing-r-and-d-vnet/subnets/lightwing-res-subnet"),
					resource.TestCheckResourceAttr(resourceName, "allocate_static_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "machine_type", "Standard_D2s_v3"),
					resource.TestCheckResourceAttr(resourceName, "keypair", "HarnessAzrWestUs2KeyPair"),
					resource.TestCheckResourceAttr(resourceName, "api_key", "don.teven-LOOKHERETHISISWR.ongbetterlookforsomething.thatiswidepublic"),
					resource.TestCheckResourceAttr(resourceName, "certificates.0.cert_secret_id", "https://lightwingrandd.vault.azure.net/secrets/secret-fullchain/thatisallreplaced6797nopointhere"),
					resource.TestCheckResourceAttr(resourceName, "certificates.0.key_secret_id", "https://lightwingrandd.vault.azure.net/secrets/secret--privkey/thatisallreplaced2nopointhered75"),
				),
			},
		},
	})
}

func testSampleDataSourceAzureProxy() string {
	return `
	resource "harness_autostopping_azure_proxy" "AzrProxyTFTest" {
		name = "AzrProxyTFTest"
		cloud_connector_id = "DoNotDelete_LightwingQA"
		host_name = "as-azr-proxy-tf-test.harness.io"
		region = "westus2"
		resource_group = "lightwing-r-and-d"
		vpc = "/subscriptions/abcdefgh-ijkl-mnop-qrst-uvwxyz123456/resourceGroups/lightwing-r-and-d/providers/Microsoft.Network/virtualNetworks/lightwing-r-and-d-vnet"
		subnet_id = "/subscriptions/abcdefgh-ijkl-mnop-qrst-uvwxyz123456/resourceGroups/lightwing-r-and-d/providers/Microsoft.Network/virtualNetworks/lightwing-r-and-d-vnet/subnets/lightwing-res-subnet"
		security_groups =["/subscriptions/abcdefgh-ijkl-mnop-qrst-uvwxyz123456/resourceGroups/lightwing-r-and-d/providers/Microsoft.Network/networkSecurityGroups/AutoStopNSG"]
		allocate_static_ip = true
		machine_type = "Standard_D2s_v3"
		keypair = "HarnessAzrWestUs2KeyPair"
		api_key = "don.teven-LOOKHERETHISISWR.ongbetterlookforsomething.thatiswidepublic"
		certificates {
			cert_secret_id = "https://lightwingrandd.vault.azure.net/secrets/secret-fullchain/thatisallreplaced6797nopointhere"
			key_secret_id = "https://lightwingrandd.vault.azure.net/secrets/secret--privkey/thatisallreplaced2nopointhered75"
		}
	}`
}
