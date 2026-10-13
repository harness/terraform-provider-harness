package connector_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorAwsCC(t *testing.T) {
	//t.Skip("Skipping until account id issue is fixed https://harness.atlassian.net/browse/PL-20793")

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	accountID := uniqueTestAwsAccountID()
	resourceName := "harness_platform_connector_awscc.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAwsCC(id, name, accountID, "test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorAwsCC(id, name, accountID, "updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "description", "updated"),
					resource.TestCheckResourceAttr(resourceName, "report_name", "test_report"),
					resource.TestCheckResourceAttr(resourceName, "report_type", "CUR2.0"),
					resource.TestCheckResourceAttr(resourceName, "region", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "s3_bucket", "s3bucket"),
					resource.TestCheckResourceAttr(resourceName, "s3_prefix", "cur"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceConnectorAwsCC(id string, name string, accountID string, description string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_awscc" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "%[4]s"
			tags = ["foo:bar"]

			account_id = "%[3]s"
			report_name = "test_report"
			report_type = "CUR2.0"
			region = "us-east-1"
			s3_bucket = "s3bucket"
			s3_prefix = "cur"
			features_enabled = [
				"OPTIMIZATION",
				"VISIBILITY",
				"BILLING"
			]
			cross_account_access {
				role_arn = "arn:aws:iam::%[3]s:role/S3Access"
				external_id = "harness:999999999999"
			}
		}
`, id, name, accountID, description)
}

func uniqueTestAwsAccountID() string {
	return fmt.Sprintf("%012d", time.Now().UnixNano()%1_000_000_000_000)
}

func TestAccResourceConnectorAwsCCNoCUR(t *testing.T) {
	//t.Skip("Skipping until account id issue is fixed https://harness.atlassian.net/browse/PL-20793")

	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	// updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_awscc.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAwsCCNoCUR(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "cross_account_access.0.role_arn", "arn:aws:iam::123456789012:role/S3Access"),
					resource.TestCheckResourceAttr(resourceName, "cross_account_access.0.external_id", "harness:999999999999"),
					resource.TestCheckResourceAttr(resourceName, "features_enabled.#", "2"),
				),
			},
			// {
			// 	Config: testAccResourceConnector_artifactory_anonymous(id, updatedName),
			// 	Check: resource.ComposeTestCheckFunc(
			// 		resource.TestCheckResourceAttr(resourceName, "id", id),
			// 		resource.TestCheckResourceAttr(resourceName, "identifier", id),
			// 		resource.TestCheckResourceAttr(resourceName, "name", updatedName),
			// 		resource.TestCheckResourceAttr(resourceName, "description", "test"),
			// 		resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
			// 		// resource.TestCheckResourceAttr(resourceName, "artifactory.0.url", "https://artifactory.example.com"),
			// 		// resource.TestCheckResourceAttr(resourceName, "artifactory.0.delegate_selectors.#", "1"),
			// 	),
			// },
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceConnectorAwsCCNoCUR(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_awscc" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			account_id = "123456789012"
			features_enabled = [
				"OPTIMIZATION",
				"VISIBILITY"
			]
			cross_account_access {
				role_arn = "arn:aws:iam::123456789012:role/S3Access"
				external_id = "harness:999999999999"
			}
		}
`, id, name)
}
