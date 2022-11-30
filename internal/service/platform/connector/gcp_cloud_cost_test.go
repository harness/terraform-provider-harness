package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorGCPCloudCost(t *testing.T) {
	id := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_gcp_cloud_cost.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorGCPCloudCost(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "gcp_project_id"),
					resource.TestCheckResourceAttr(resourceName, "service_account_email", "service_account_email"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.data_set_id", "data_set_id"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.table_id", "table_id"),
				),
			},
			{
				Config: testAccResourceConnectorGCPCloudCost(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "gcp_project_id", "gcp_project_id"),
					resource.TestCheckResourceAttr(resourceName, "service_account_email", "service_account_email"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.data_set_id", "data_set_id"),
					resource.TestCheckResourceAttr(resourceName, "billing_export_spec.0.table_id", "table_id"),
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

func testAccResourceConnectorGCPCloudCost(id string, name string) string {
	return fmt.Sprintf(`
		resource "harness_platform_connector_gcp_cloud_cost" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			features_enabled = ["BILLING", "VISIBILITY", "OPTIMIZATION"]
			gcp_project_id = "gcp_project_id"
			service_account_email = "service_account_email"
			billing_export_spec {
				data_set_id = "data_set_id"
				table_id = "table_id"
			}
		}
`, id, name)
}
