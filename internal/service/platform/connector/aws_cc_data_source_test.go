package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorAwsCC(t *testing.T) {
	t.Skip("Skipping until account id issue is fixed https://harness.atlassian.net/browse/PL-20793")

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_awscc.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnectorAwsCC(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "cross_account_access.0.role_arn", "arn:aws:iam::123456789012:role/S3Access"),
					resource.TestCheckResourceAttr(resourceName, "cross_account_access.0.external_id", "harness:999999999999"),
					resource.TestCheckResourceAttr(resourceName, "features_enabled.#", "3"),
				),
			},
		},
	})
}

func testAccDataSourceConnectorAwsCC(name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_awscc" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "000000000001"
			report_name = "test_report"
			s3_bucket = "s3bucket"
			features_enabled = [
				"OPTIMIZATION",
				"VISIBILITY",
				"BILLING",
			]
			cross_account_access {
				role_arn = "arn:aws:iam::123456789012:role/S3Access"
				external_id = "harness:999999999999"
			}
		}

		data "harness_platform_connector_awscc" "test" {
			identifier = harness_platform_connector_awscc.test.identifier
		}
	`, name)
}
