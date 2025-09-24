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

func TestAccResourceFMEEnvironmentSegmentKeys(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	environmentID := os.Getenv("SPLIT_ENVIRONMENT_ID")
	if workspaceID == "" || environmentID == "" {
		t.Skip("SPLIT_WORKSPACE_ID and SPLIT_ENVIRONMENT_ID environment variables must be set for this test")
	}

	segmentName := fmt.Sprintf("test-segment-%s", utils.RandStringBytes(5))
	trafficTypeName := fmt.Sprintf("test-traffic-type-%s", utils.RandStringBytes(5))
	resourceName := "harness_fme_environment_segment_keys.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMEEnvironmentSegmentKeysDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMEEnvironmentSegmentKeys(workspaceID, environmentID, trafficTypeName, segmentName, []string{"user123", "user456"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttr(resourceName, "segment_name", segmentName),
					resource.TestCheckResourceAttr(resourceName, "environment_id", environmentID),
					resource.TestCheckResourceAttr(resourceName, "keys.#", "2"),
					resource.TestCheckTypeSetElemAttr(resourceName, "keys.*", "user123"),
					resource.TestCheckTypeSetElemAttr(resourceName, "keys.*", "user456"),
				),
			},
			{
				Config: testAccResourceFMEEnvironmentSegmentKeys(workspaceID, environmentID, trafficTypeName, segmentName, []string{"user123", "user456", "user789"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "keys.#", "3"),
					resource.TestCheckTypeSetElemAttr(resourceName, "keys.*", "user123"),
					resource.TestCheckTypeSetElemAttr(resourceName, "keys.*", "user456"),
					resource.TestCheckTypeSetElemAttr(resourceName, "keys.*", "user789"),
				),
			},
			{
				Config: testAccResourceFMEEnvironmentSegmentKeys(workspaceID, environmentID, trafficTypeName, segmentName, []string{"user456"}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "keys.#", "1"),
					resource.TestCheckTypeSetElemAttr(resourceName, "keys.*", "user456"),
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

func testAccFMEEnvironmentSegmentKeysDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// Get the resource from state
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no environment segment keys ID is set")
		}

		// Get session and check if keys still exist
		session := acctest.TestAccGetSession()
		if session.FMEClient == nil {
			return fmt.Errorf("FME client not configured")
		}

		workspaceID := rs.Primary.Attributes["workspace_id"]
		segmentName := rs.Primary.Attributes["segment_name"]
		environmentID := rs.Primary.Attributes["environment_id"]

		if workspaceID == "" || segmentName == "" || environmentID == "" {
			return fmt.Errorf("no workspace ID, segment name, or environment ID found in state")
		}

		c := session.FMEClient.(*fme.FMEConfig)
		segmentKeys, err := c.APIClient.EnvironmentSegmentKeys.Get(workspaceID, segmentName, environmentID)
		if err == nil && segmentKeys != nil && len(segmentKeys.Keys) > 0 {
			return fmt.Errorf("FME environment segment keys still exist: %s", rs.Primary.ID)
		}

		return nil
	}
}

func testAccResourceFMEEnvironmentSegmentKeys(workspaceID, environmentID, trafficTypeName, segmentName string, keys []string) string {
	keysConfig := ""
	for _, key := range keys {
		keysConfig += fmt.Sprintf(`"%s", `, key)
	}
	if len(keysConfig) > 0 {
		keysConfig = keysConfig[:len(keysConfig)-2] // Remove trailing comma and space
	}

	return fmt.Sprintf(`
		resource "harness_fme_traffic_type" "test" {
			workspace_id = "%[1]s"
			name         = "%[3]s"
		}

		resource "harness_fme_segment" "test" {
			workspace_id     = "%[1]s"
			traffic_type_id  = harness_fme_traffic_type.test.id
			name             = "%[4]s"
			description      = "Test segment for environment keys"
		}

		resource "harness_fme_environment_segment_keys" "test" {
			workspace_id   = "%[1]s"
			segment_name   = harness_fme_segment.test.name
			environment_id = "%[2]s"
			keys           = [%[5]s]
		}
`, workspaceID, environmentID, trafficTypeName, segmentName, keysConfig)
}