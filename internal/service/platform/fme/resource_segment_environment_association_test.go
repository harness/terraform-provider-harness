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

func TestAccResourceFMESegmentEnvironmentAssociation(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	environmentID := os.Getenv("SPLIT_ENVIRONMENT_ID")
	if workspaceID == "" || environmentID == "" {
		t.Skip("SPLIT_WORKSPACE_ID and SPLIT_ENVIRONMENT_ID environment variables must be set for this test")
	}

	segmentName := fmt.Sprintf("test-segment-%s", utils.RandStringBytes(5))
	trafficTypeName := fmt.Sprintf("test-traffic-type-%s", utils.RandStringBytes(5))
	resourceName := "harness_fme_segment_environment_association.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMESegmentEnvironmentAssociationDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMESegmentEnvironmentAssociation(workspaceID, environmentID, trafficTypeName, segmentName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttr(resourceName, "segment_name", segmentName),
					resource.TestCheckResourceAttr(resourceName, "environment_id", environmentID),
					resource.TestCheckResourceAttr(resourceName, "include_in_segment", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "environment_name"),
				),
			},
			{
				Config: testAccResourceFMESegmentEnvironmentAssociation(workspaceID, environmentID, trafficTypeName, segmentName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "include_in_segment", "false"),
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

func testAccFMESegmentEnvironmentAssociationDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// Get the resource from state
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no segment environment association ID is set")
		}

		// Get session and check if association still exists
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
		association, err := c.APIClient.SegmentEnvironmentAssociations.Get(workspaceID, segmentName, environmentID)
		if err == nil && association != nil && association.IncludeInSegment != nil && *association.IncludeInSegment {
			return fmt.Errorf("FME segment environment association still active: %s", rs.Primary.ID)
		}

		return nil
	}
}

func testAccResourceFMESegmentEnvironmentAssociation(workspaceID, environmentID, trafficTypeName, segmentName string, includeInSegment bool) string {
	return fmt.Sprintf(`
		resource "harness_fme_traffic_type" "test" {
			workspace_id = "%[1]s"
			name         = "%[3]s"
		}

		resource "harness_fme_segment" "test" {
			workspace_id     = "%[1]s"
			traffic_type_id  = harness_fme_traffic_type.test.id
			name             = "%[4]s"
			description      = "Test segment for environment association"
		}

		resource "harness_fme_segment_environment_association" "test" {
			workspace_id       = "%[1]s"
			segment_name       = harness_fme_segment.test.name
			environment_id     = "%[2]s"
			include_in_segment = %[5]t
		}
`, workspaceID, environmentID, trafficTypeName, segmentName, includeInSegment)
}