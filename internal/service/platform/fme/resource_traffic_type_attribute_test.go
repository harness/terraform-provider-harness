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

func TestAccResourceFMETrafficTypeAttribute(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this test")
	}

	attributeName := fmt.Sprintf("test-attribute-%s", utils.RandStringBytes(5))
	trafficTypeName := fmt.Sprintf("test-traffic-type-%s", utils.RandStringBytes(5))
	resourceName := "harness_fme_traffic_type_attribute.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMETrafficTypeAttributeDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMETrafficTypeAttribute(workspaceID, trafficTypeName, attributeName, "STRING", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttr(resourceName, "display_name", attributeName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial test attribute description"),
					resource.TestCheckResourceAttr(resourceName, "data_type", "STRING"),
					resource.TestCheckResourceAttr(resourceName, "is_searchable", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "traffic_type_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				Config: testAccResourceFMETrafficTypeAttribute(workspaceID, trafficTypeName, attributeName, "NUMBER", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", attributeName),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated test attribute description"),
					resource.TestCheckResourceAttr(resourceName, "data_type", "NUMBER"),
					resource.TestCheckResourceAttr(resourceName, "is_searchable", "true"),
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

func testAccFMETrafficTypeAttributeDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// Get the resource from state
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no traffic type attribute ID is set")
		}

		// Get session and check if attribute still exists
		session := acctest.TestAccGetSession()
		if session.FMEClient == nil {
			return fmt.Errorf("FME client not configured")
		}

		workspaceID := rs.Primary.Attributes["workspace_id"]
		if workspaceID == "" {
			return fmt.Errorf("no workspace ID found in state")
		}

		trafficTypeID := rs.Primary.Attributes["traffic_type_id"]
		if trafficTypeID == "" {
			return fmt.Errorf("no traffic type ID found in state")
		}

		c := session.FMEClient.(*fme.FMEConfig)
		attribute, err := c.APIClient.TrafficTypeAttributes.Get(workspaceID, trafficTypeID, rs.Primary.ID)
		if err == nil && attribute != nil {
			return fmt.Errorf("FME traffic type attribute still exists: %s", rs.Primary.ID)
		}

		return nil
	}
}

func testAccResourceFMETrafficTypeAttribute(workspaceID, trafficTypeName, attributeName, dataType string, isSearchable bool) string {
	description := "Initial test attribute description"
	if dataType == "NUMBER" {
		description = "Updated test attribute description"
	}

	return fmt.Sprintf(`
		resource "harness_fme_traffic_type" "test" {
			workspace_id = "%[1]s"
			name         = "%[2]s"
		}

		resource "harness_fme_traffic_type_attribute" "test" {
			workspace_id      = "%[1]s"
			traffic_type_id   = harness_fme_traffic_type.test.id
			display_name      = "%[3]s"
			description       = "%[5]s"
			data_type         = "%[4]s"
			is_searchable     = %[6]t
		}
`, workspaceID, trafficTypeName, attributeName, dataType, description, isSearchable)
}