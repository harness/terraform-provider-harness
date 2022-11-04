package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorAzureCloudCost(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_azure_cloud_cost.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnector_azureCloudCost(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tenant_id", "tenant_id"),
					resource.TestCheckResourceAttr(resourceName, "subscription_id", "subscription_id"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.storage_account_name", "storage_account_name"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.container_name", "container_name"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.directory_name", "directory_name"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.report_name", "report_name"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.subscription_id", "subscription_id"),
				),
			},
		},
	})

}

func testAccDataSourceConnector_azureCloudCost(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_azure_cloud_cost" "test" {
		identifier = "%[1]s"
		name = "%[1]s"
		description = "test"
		tags = ["foo:bar"]

		features_enabled = ["BILLING", "VISIBILITY", "OPTIMIZATION"]
		tenant_id ="tenant_id"
		subscription_id = "subscription_id"
		billing_export_spec {
			storage_account_name = "storage_account_name"
			container_name = "container_name"
			directory_name = "directory_name"
			report_name = "report_name"
			subscription_id = "subscription_id"
		}
	}

	data "harness_platform_connector_azure_cloud_cost" "test" {
		identifier = harness_platform_connector_azure_cloud_cost.test.identifier
	}
`, name)
}
