package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorAzureCloudCost(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_azure_cloud_cost.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorAzureCloudCost(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
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
			{
				Config: testAccResourceConnectorAzureCloudCost(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tenant_id", "tenant_id"),
					resource.TestCheckResourceAttr(resourceName, "subscription_id", "subscription_id"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.storage_account_name", "storage_account_name"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.container_name", "container_name"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.directory_name", "directory_name"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.report_name", "report_name"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.subscription_id", "subscription_id"),
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

func testAccResourceConnectorAzureCloudCost(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_azure_cloud_cost" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
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
`, id, name)
}
