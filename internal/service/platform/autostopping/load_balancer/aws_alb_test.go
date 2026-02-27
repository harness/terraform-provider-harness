package load_balancer_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// AWS cloud connector ID in the test account.
const awsCloudConnectorID = "automation_aws_connector"

func TestResourceAwsALB(t *testing.T) {
	name := fmt.Sprintf("terr-awsalb-%s", strings.ToLower(utils.RandStringBytes(5)))
	resourceName := "harness_autostopping_aws_alb.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		//		CheckDestroy:      testAWSProxyDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAwsALB(name),
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

func testAwsALB(name string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_alb" "test" {
			name = "%[1]s"
			cloud_connector_id = %q
            region = "us-east-1"
			vpc = "vpc-0d47ab08fce6d8cc8"
			security_groups =["sg-0a2a6eaa3ad797636"]
			delete_cloud_resources_on_destroy = true
		}
`, name, awsCloudConnectorID)
}

func testAwsALBUpdate(name string) string {
	return fmt.Sprintf(`
		resource "harness_autostopping_aws_alb" "test" {
			name = "%[1]s"
			cloud_connector_id = %q
            region = "us-east-1"
            vpc = "vpc-0d47ab08fce6d8cc8"
            security_groups =["sg-0a2a6eaa3ad797636"]
			delete_cloud_resources_on_destroy = true
		}
`, name, awsCloudConnectorID)
}
