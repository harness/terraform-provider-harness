package fme_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceFMETrafficType(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this test")
	}

	trafficTypeName := fmt.Sprintf("test-traffic-type-ds-%s", utils.RandStringBytes(5))
	resourceName := "data.harness_fme_traffic_type.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMETrafficType(workspaceID, trafficTypeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", trafficTypeName),
					resource.TestCheckResourceAttr(resourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "type"),
				),
			},
		},
	})
}

func testAccDataSourceFMETrafficType(workspaceID, trafficTypeName string) string {
	return fmt.Sprintf(`
		resource "harness_fme_traffic_type" "test" {
			workspace_id = "%[1]s"
			name         = "%[2]s"
		}

		data "harness_fme_traffic_type" "test" {
			workspace_id = "%[1]s"
			id           = harness_fme_traffic_type.test.id
		}
`, workspaceID, trafficTypeName)
}