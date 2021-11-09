package ng_test

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/nextgen"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/harness-io/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnector(t *testing.T) {

	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_connector.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnector(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "type", nextgen.ConnectorTypes.K8sCluster.String()),
				),
			},
		},
	})
}

func testAccDataSourceConnector(name string) string {
	return fmt.Sprintf(`
		resource "harness_connector" "test" {
			identifier = "%[1]s"
			name = "%[1]s"

			k8s_cluster {
				inherit_from_delegate {
					delegate_selectors = ["harness-delegate"]
				}
			}
		}

		data "harness_connector" "test" {
			identifier = harness_connector.test.identifier
		}
	`, name)
}
