package fme_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/harness/terraform-provider-harness/internal/service/platform/fme"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceFMESplitDefinition(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	environmentID := os.Getenv("SPLIT_ENVIRONMENT_ID")
	if workspaceID == "" || environmentID == "" {
		t.Skip("SPLIT_WORKSPACE_ID and SPLIT_ENVIRONMENT_ID environment variables must be set for this test")
	}

	splitName := fmt.Sprintf("test-split-%s", utils.RandStringBytes(5))
	trafficTypeName := fmt.Sprintf("test-traffic-type-%s", utils.RandStringBytes(5))
	resourceName := "harness_fme_split_definition.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMESplitDefinitionDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMESplitDefinition(workspaceID, environmentID, trafficTypeName, splitName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttr(resourceName, "environment_id", environmentID),
					resource.TestCheckResourceAttr(resourceName, "split_name", splitName),
					resource.TestCheckResourceAttr(resourceName, "treatments.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "treatments.0.name", "on"),
					resource.TestCheckResourceAttr(resourceName, "treatments.1.name", "off"),
					resource.TestCheckResourceAttr(resourceName, "default_rule.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "default_rule.0.treatment", "off"),
					resource.TestCheckResourceAttr(resourceName, "baseline_treatment", "off"),
					resource.TestCheckResourceAttrSet(resourceName, "creation_time"),
				),
			},
			{
				Config: testAccResourceFMESplitDefinitionUpdate(workspaceID, environmentID, trafficTypeName, splitName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "treatments.0.description", "Updated feature enabled"),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.treatment", "on"),
					resource.TestCheckResourceAttr(resourceName, "rules.0.size", "75"),
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

func testAccFMESplitDefinitionDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// Get the resource from state
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no split definition ID is set")
		}

		// Get session and check if split definition still exists
		session := acctest.TestAccGetSession()
		if session.FMEClient == nil {
			return fmt.Errorf("FME client not configured")
		}

		workspaceID := rs.Primary.Attributes["workspace_id"]
		environmentID := rs.Primary.Attributes["environment_id"]
		splitName := rs.Primary.ID

		if workspaceID == "" || environmentID == "" {
			return fmt.Errorf("no workspace ID or environment ID found in state")
		}

		c := session.FMEClient.(*fme.FMEConfig)
		splitDefinition, err := c.APIClient.SplitDefinitions.Get(workspaceID, environmentID, splitName)
		if err == nil && splitDefinition != nil {
			// Split definitions might still exist but be reset to default state
			// We just check that the custom treatments are gone
			if len(splitDefinition.Treatments) > 0 {
				return fmt.Errorf("FME split definition still has custom treatments: %s", splitName)
			}
		}

		return nil
	}
}

func testAccResourceFMESplitDefinition(workspaceID, environmentID, trafficTypeName, splitName string) string {
	return fmt.Sprintf(`
		resource "harness_fme_traffic_type" "test" {
			workspace_id = "%[1]s"
			name         = "%[3]s"
		}

		resource "harness_fme_split" "test" {
			workspace_id     = "%[1]s"
			name             = "%[4]s"
			description      = "Test split for split definition"
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
`, workspaceID, environmentID, trafficTypeName, splitName)
}

func testAccResourceFMESplitDefinitionUpdate(workspaceID, environmentID, trafficTypeName, splitName string) string {
	return fmt.Sprintf(`
		resource "harness_fme_traffic_type" "test" {
			workspace_id = "%[1]s"
			name         = "%[3]s"
		}

		resource "harness_fme_split" "test" {
			workspace_id     = "%[1]s"
			name             = "%[4]s"
			description      = "Test split for split definition"
			traffic_type_id  = harness_fme_traffic_type.test.id
		}

		resource "harness_fme_split_definition" "test" {
			workspace_id   = "%[1]s"
			environment_id = "%[2]s"
			split_name     = harness_fme_split.test.name

			treatments {
				name        = "on"
				description = "Updated feature enabled"
				configurations {
					name  = "color"
					value = "green"
				}
			}

			treatments {
				name        = "off"
				description = "Feature disabled"
			}

			rules {
				treatment = "on"
				size      = 75
			}

			default_rule {
				treatment = "off"
				size      = 25
			}

			baseline_treatment = "off"
		}
`, workspaceID, environmentID, trafficTypeName, splitName)
}