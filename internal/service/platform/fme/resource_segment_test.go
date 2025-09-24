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

func TestAccResourceFMESegment(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this test")
	}

	segmentName := fmt.Sprintf("test-segment-%s", utils.RandStringBytes(5))
	trafficTypeName := fmt.Sprintf("test-traffic-type-%s", utils.RandStringBytes(5))
	resourceName := "harness_fme_segment.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMESegmentDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMESegment(workspaceID, trafficTypeName, segmentName, "Initial test segment description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttr(resourceName, "name", segmentName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial test segment description"),
					resource.TestCheckResourceAttrSet(resourceName, "traffic_type_id"),
					resource.TestCheckResourceAttrSet(resourceName, "creation_time"),
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

func testAccFMESegmentDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// Get the resource from state
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no segment ID is set")
		}

		// Get session and check if segment still exists
		session := acctest.TestAccGetSession()
		if session.FMEClient == nil {
			return fmt.Errorf("FME client not configured")
		}

		workspaceID := rs.Primary.Attributes["workspace_id"]
		if workspaceID == "" {
			return fmt.Errorf("no workspace ID found in state")
		}

		c := session.FMEClient.(*fme.FMEConfig)
		segment, err := c.APIClient.Segments.Get(workspaceID, rs.Primary.ID)
		if err == nil && segment != nil {
			return fmt.Errorf("FME segment still exists: %s", rs.Primary.ID)
		}

		return nil
	}
}

func testAccResourceFMESegment(workspaceID, trafficTypeName, segmentName, description string) string {
	return fmt.Sprintf(`
		resource "harness_fme_traffic_type" "test" {
			workspace_id = "%[1]s"
			name         = "%[2]s"
		}

		resource "harness_fme_segment" "test" {
			workspace_id     = "%[1]s"
			traffic_type_id  = harness_fme_traffic_type.test.id
			name             = "%[3]s"
			description      = "%[4]s"
		}
`, workspaceID, trafficTypeName, segmentName, description)
}