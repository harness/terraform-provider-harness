package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceConnectorKubernetsCloudCost(t *testing.T) {
	id := fmt.Sprintf("%s_%s", "KubernetsCloudCost", utils.RandStringBytes(5))
	name := id
	updatedName := fmt.Sprintf("%s_updated", name)
	resourceName := "harness_platform_connector_kubernetes_cloud_cost.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccConnectorDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConnectorKubernetsCloudCost(id, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connector_ref", id+"s"),
				),
			},
			{
				Config: testAccResourceConnectorKubernetsCloudCost(id, updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", id),
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connector_ref", id+"s"),
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

func testAccResourceConnectorKubernetsCloudCost(id string, name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_kubernetes" "test" {
		identifier = "%[1]ss"
		name = "%[2]ss"
		description = "test"
		tags = ["foo:bar"]

		inherit_from_delegate {
			delegate_selectors = ["harness-delegate"]
		}
	}

		resource "harness_platform_connector_kubernetes_cloud_cost" "test" {
			identifier = "%[1]s"
			name = "%[2]s"
			description = "test"
			tags = ["foo:bar"]

			features_enabled = ["VISIBILITY", "OPTIMIZATION"]
			connector_ref = harness_platform_connector_kubernetes.test.id
		}
`, id, name)
}
