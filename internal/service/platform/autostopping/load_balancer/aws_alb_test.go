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

// HARNESS_AWS_CLOUD_CONNECTOR_ID must be set to a valid AWS cloud connector ID in the test account to run this test.
const awsCloudConnectorIDEnv = "automation_aws_connector"

func TestResourceAwsALB(t *testing.T) {
	connectorID := os.Getenv(awsCloudConnectorIDEnv)
	if connectorID == "" {
		t.Fatalf("%s must be set to a valid AWS cloud connector ID", awsCloudConnectorIDEnv)
	}

	name := utils.RandStringBytes(5)
	resourceName := "harness_autostopping_aws_alb.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testAWSProxyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAwsALB(name, connectorID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
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

func testAwsALB(name, cloudConnectorID string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_alb" "test" {
			name = "%[1]s"
			cloud_connector_id = "%[2]s"
            region = "us-east-1"
			vpc = "vpc-2657db5c"
			security_groups =["sg-01","sg-02"]
			delete_cloud_resources_on_destroy = true
		}
`, name, cloudConnectorID)
}

func testAwsALBUpdate(name, cloudConnectorID string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_alb" "test" {
			name = "%[1]s"
			cloud_connector_id = "%[2]s"
            region = "us-east-1"
            vpc = "vpc-2657db5c"
			security_groups =["sg-01","sg-02"]
			delete_cloud_resources_on_destroy = true
		}
`, name, cloudConnectorID)
}
