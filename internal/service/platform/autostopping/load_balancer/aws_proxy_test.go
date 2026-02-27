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

// Env var that must be set to a valid AWS cloud connector ID to run this test.
const awsProxyCloudConnectorIDEnv = "automation_aws_connector"

func TestResourceAWSProxy(t *testing.T) {
	connectorID := os.Getenv(awsProxyCloudConnectorIDEnv)
	if connectorID == "" {
		t.Fatalf("%s must be set to a valid AWS cloud connector ID", awsProxyCloudConnectorIDEnv)
	}

	name := utils.RandStringBytes(5)
	resourceName := "harness_autostopping_aws_proxy.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testAWSProxyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAWSProxy(name, connectorID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
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

func testAWSProxy(name, cloudConnectorID string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = "%[2]s"
            region = "us-east-1"
			vpc = "vpc-2657db5c"
			security_groups =["sg-01"]
			machine_type = "t2.medium"
            api_key = ""
			allocate_static_ip = true
			delete_cloud_resources_on_destroy = false
		}
`, name, cloudConnectorID)
}

func testAWSProxyUpdate(name, cloudConnectorID string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_proxy" "test" {
			name = "%[1]s"
			cloud_connector_id = "%[2]s"
            region = "eastus2"
            vpc = "vpc-2657db5c"
			security_groups =["sg-01","sg-02"]
            machine_type = "t2.medium"
            api_key = ""
			allocate_static_ip = true
			delete_cloud_resources_on_destroy = true
		}
`, name, cloudConnectorID)
}
