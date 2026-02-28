package load_balancer_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// AWS cloud connector ID in the test account.
const awsProxyCloudConnectorID = "automation_aws_connector"

func TestResourceAWSProxy(t *testing.T) {
	apiKey := os.Getenv(platformAPIKeyEnv)

	name := fmt.Sprintf("terr-awsproxy-%s", strings.ToLower(utils.RandLowerString(5)))
	resourceName := "harness_autostopping_aws_proxy.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testAWSProxyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAWSProxy(name, apiKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			{
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
            vpc = "vpc-0d47ab08fce6d8cc8"
            security_groups =["sg-0a2a6eaa3ad797636"]
			machine_type = "t2.medium"
            api_key = %q
			allocate_static_ip = true
			delete_cloud_resources_on_destroy = false
		}
`, name, awsProxyCloudConnectorID, apiKey)
}

func testAWSProxyUpdate(name, apiKey string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = %q
            region = "eastus2"
            vpc = "vpc-0d47ab08fce6d8cc8"
            security_groups =["sg-0a2a6eaa3ad797636"]
            machine_type = "t2.medium"
            api_key = %q
			allocate_static_ip = true
			delete_cloud_resources_on_destroy = true
		}
`, name, awsProxyCloudConnectorID, apiKey)
}
