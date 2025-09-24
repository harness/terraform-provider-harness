package fme_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceFMESplitDefinition(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	environmentID := os.Getenv("SPLIT_ENVIRONMENT_ID")
	if workspaceID == "" || environmentID == "" {
		t.Skip("SPLIT_WORKSPACE_ID and SPLIT_ENVIRONMENT_ID environment variables must be set for this test")
	}

	splitName := fmt.Sprintf("test-split-%s", utils.RandStringBytes(5))
	trafficTypeName := fmt.Sprintf("test-traffic-type-%s", utils.RandStringBytes(5))
	dataSourceName := "data.harness_fme_split_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFMESplitDefinition(workspaceID, environmentID, trafficTypeName, splitName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttr(dataSourceName, "environment_id", environmentID),
					resource.TestCheckResourceAttr(dataSourceName, "split_name", splitName),
					resource.TestCheckResourceAttr(dataSourceName, "treatments.#", "2"),
					resource.TestCheckResourceAttr(dataSourceName, "treatments.0.name", "on"),
					resource.TestCheckResourceAttr(dataSourceName, "treatments.1.name", "off"),
					resource.TestCheckResourceAttr(dataSourceName, "default_rule.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "default_rule.0.treatment", "off"),
					resource.TestCheckResourceAttr(dataSourceName, "baseline_treatment", "off"),
					resource.TestCheckResourceAttrSet(dataSourceName, "creation_time"),
				),
			},
		},
	})
}

func testAccDataSourceFMESplitDefinition(workspaceID, environmentID, trafficTypeName, splitName string) string {
	return fmt.Sprintf(`
		resource "harness_fme_traffic_type" "test" {
			workspace_id = "%[1]s"
			name         = "%[3]s"
		}

		resource "harness_fme_split" "test" {
			workspace_id     = "%[1]s"
			name             = "%[4]s"
			description      = "Test split for data source"
			traffic_type_id  = harness_fme_traffic_type.test.id
		}

		resource "harness_fme_split_definition" "test" {
			workspace_id   = "%[1]s"
			environment_id = "%[2]s"
			split_name     = harness_fme_split.test.name

			treatments {
				name        = "on"
				description = "Feature enabled"
				configurations {
					name  = "color"
					value = "blue"
				}
			}

			treatments {
				name        = "off"
				description = "Feature disabled"
			}

			default_rule {
				treatment = "off"
				size      = 100
			}

			baseline_treatment = "off"
		}

		data "harness_fme_split_definition" "test" {
			workspace_id   = "%[1]s"
			environment_id = "%[2]s"
			split_name     = harness_fme_split_definition.test.split_name

			depends_on = [harness_fme_split_definition.test]
		}
`, workspaceID, environmentID, trafficTypeName, splitName)
}