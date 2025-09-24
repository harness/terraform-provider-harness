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

func TestAccResourceFMESplit(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this test")
	}

	splitName := fmt.Sprintf("test-split-%s", utils.RandStringBytes(5))
	updatedDescription := "Updated test split description"
	resourceName := "harness_fme_split.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMESplitDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMESplit(workspaceID, splitName, "Initial test split description"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttr(resourceName, "name", splitName),
					resource.TestCheckResourceAttr(resourceName, "description", "Initial test split description"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "creation_time"),
				),
			},
			{
				Config: testAccResourceFMESplit(workspaceID, splitName, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workspace_id", workspaceID),
					resource.TestCheckResourceAttr(resourceName, "name", splitName),
					resource.TestCheckResourceAttr(resourceName, "description", updatedDescription),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
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

func TestAccResourceFMESplit_WithWorkspaceData(t *testing.T) {
	workspaceID := os.Getenv("SPLIT_WORKSPACE_ID")
	if workspaceID == "" {
		t.Skip("SPLIT_WORKSPACE_ID environment variable must be set for this test")
	}

	splitName := fmt.Sprintf("test-split-data-%s", utils.RandStringBytes(5))
	resourceName := "harness_fme_split.test"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheckFME(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      testAccFMESplitDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFMESplitWithWorkspaceData(workspaceID, splitName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", splitName),
					resource.TestCheckResourceAttr(resourceName, "description", "Split using workspace data source"),
					resource.TestCheckResourceAttrSet(resourceName, "workspace_id"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccFMESplitDestroy(resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		// Get the resource from state
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no split ID is set")
		}

		// Get session and check if split still exists
		session := acctest.TestAccGetSession()
		if session.FMEClient == nil {
			return fmt.Errorf("FME client not configured")
		}

		workspaceID := rs.Primary.Attributes["workspace_id"]
		if workspaceID == "" {
			return fmt.Errorf("no workspace ID found in state")
		}

		c := session.FMEClient.(*fme.FMEConfig)
		split, err := c.APIClient.Splits.Get(workspaceID, rs.Primary.ID)
		if err == nil && split != nil {
			return fmt.Errorf("FME split still exists: %s", rs.Primary.ID)
		}

		return nil
	}
}

func testAccResourceFMESplit(workspaceID, splitName, description string) string {
	return fmt.Sprintf(`
		resource "harness_fme_split" "test" {
			workspace_id = "%[1]s"
			name         = "%[2]s"
			description  = "%[3]s"
		}
`, workspaceID, splitName, description)
}

func testAccResourceFMESplitWithWorkspaceData(workspaceID, splitName string) string {
	return fmt.Sprintf(`
		data "harness_fme_workspace" "test" {
			id = "%[1]s"
		}

		resource "harness_fme_split" "test" {
			workspace_id = data.harness_fme_workspace.test.id
			name         = "%[2]s"
			description  = "Split using workspace data source"
		}
`, workspaceID, splitName)
}