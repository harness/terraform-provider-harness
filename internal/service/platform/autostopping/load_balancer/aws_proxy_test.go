package load_balancer_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const awsProxyCloudConnectorID = "DoNotDelete_LightwingNonProd"
const awsProxyVPC = "vpc-08f63488a1e3c1bf1"
const awsProxySG = "sg-0e1f9ee9896d96583"

func TestResourceAWSProxy(t *testing.T) {
	apiKey := os.Getenv(platformAPIKeyEnv)

	name := fmt.Sprintf("terr-awsproxy-%s", randAlnum(5))
	resourceName := "harness_autostopping_aws_proxy.test"
	var proxyID string

	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
			cleanupStaleAWSProxies(t)
		},
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAWSProxy(name, apiKey),
				Check: resource.ComposeTestCheckFunc(
					extractAttr(resourceName, "identifier", &proxyID),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				PreConfig: func() {
					if err := waitForProxyReady(proxyID, 5*time.Minute); err != nil {
						t.Fatalf("Proxy not ready before update: %v", err)
					}
				},
				Config: testAWSProxyUpdate(name, apiKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"api_key", "allocate_static_ip", "delete_cloud_resources_on_destroy", "machine_type"},
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

func testAWSProxy(name, apiKey string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = %q
            region = "us-east-1"
            vpc = %q
            security_groups =[%q]
			machine_type = "t2.medium"
            api_key = %q
			allocate_static_ip = false
			delete_cloud_resources_on_destroy = false
		}
`, name, awsProxyCloudConnectorID, awsProxyVPC, awsProxySG, apiKey)
}

func testAWSProxyUpdate(name, apiKey string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = %q
            region = "us-east-1"
            vpc = %q
            security_groups =[%q]
            machine_type = "t2.medium"
            api_key = %q
			allocate_static_ip = false
			delete_cloud_resources_on_destroy = true
		}
`, name, awsProxyCloudConnectorID, awsProxyVPC, awsProxySG, apiKey)
}
