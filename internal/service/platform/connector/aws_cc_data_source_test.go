package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorAwsCC(t *testing.T) {
	name := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	accountID := uniqueTestAwsAccountID()
	resourceName := "data.harness_platform_connector_awscc.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy("harness_platform_connector_awscc.test"),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsCC(name, accountID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "report_name", "test_report"),
					resource.TestCheckResourceAttr(resourceName, "report_type", "CUR2.0"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "s3_bucket", "s3bucket"),
					resource.TestCheckResourceAttr(resourceName, "s3_prefix", "cur"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "cross_account_access.0.role_arn", fmt.Sprintf("arn:aws:iam::%s:role/S3Access", accountID)),
					resource.TestCheckResourceAttr(resourceName, "cross_account_access.0.external_id", "harness:999999999999"),
					resource.TestCheckResourceAttr(resourceName, "features_enabled.#", "3"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorAwsCC(name string, accountID string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_awscc" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "%[2]s"
			report_name = "test_report"
			report_type = "CUR2.0"
			region = "us-east-1"
			s3_bucket = "s3bucket"
			s3_prefix = "cur"
			features_enabled = [
				"OPTIMIZATION",
				"VISIBILITY",
				"BILLING",
			]
			cross_account_access {
				role_arn = "arn:aws:iam::%[2]s:role/S3Access"
				external_id = "harness:999999999999"
			}
		}

		data "harness_platform_connector_awscc" "test" {
			identifier = harness_platform_connector_awscc.test.identifier
		}
	`, name, accountID)
}
